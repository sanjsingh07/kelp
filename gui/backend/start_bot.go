package backend

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/stellar/go/support/config"
	"github.com/stellar/kelp/gui/model2"
	"github.com/stellar/kelp/support/kelpos"
	"github.com/stellar/kelp/trader"
)

func (s *APIServer) startBot(w http.ResponseWriter, r *http.Request) {
	botNameBytes, e := ioutil.ReadAll(r.Body)
	if e != nil {
		s.writeErrorJson(w, fmt.Sprintf("error when reading request input: %s\n", e))
		return
	}
	botName := string(botNameBytes)

	e = s.doStartBot(botName, "buysell", nil, nil)
	if e != nil {
		s.writeKelpError(w, makeKelpErrorResponseWrapper(
			errorTypeBot,
			botName,
			time.Now().UTC(),
			errorLevelError,
			fmt.Sprintf("error starting bot: %s\n", e),
		))
		return
	}

	e = s.kos.AdvanceBotState(botName, kelpos.BotStateStopped)
	if e != nil {
		s.writeKelpError(w, makeKelpErrorResponseWrapper(
			errorTypeBot,
			botName,
			time.Now().UTC(),
			errorLevelError,
			fmt.Sprintf("error advancing bot state: %s\n", e),
		))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *APIServer) doStartBot(botName string, strategy string, iterations *uint8, maybeFinishCallback func()) error {
	filenamePair := model2.GetBotFilenames(botName, strategy)
	logPrefix := model2.GetLogPrefix(botName, strategy)

	// config files and log prefix under linux subsystem:
	// - unix relative paths did work on windows!
	// - native relative paths did not work on windows for config files but worked for log prefixes!
	// - native absolute paths did not work on windows
	// - unix absolute path did not work on windows
	//
	// config files and log prefix invoked without bash -c (i.e. not under linux system):
	// - unix relative paths did work on windows!
	// - native relative paths did work on windows!
	// - native absolute paths did work on windows!
	// - unix absolute path did not work on windows
	//
	// (see api_server.go#runKelpCommandBlocking and #runKelpCommandBackground for information that may be related to why
	// absolute paths did not work)
	//
	// The above experimentation makes unix relative paths the most common format so we will use that to start new bots
	//
	// Note that on windows it could use the native windows naming scheme (C:\ etc.) but in the linux subsystem on windows
	// there is no C:\ but instead is listed as /mnt/c/... so we need to convert from unix to windows (/mnt/c -> C:\) or
	// use relative paths, which is why it seems to work
	// Note that /mnt/c is unlikely to be valid in windows (but is valid in the linux subsystem) since it's usually prefixed by the
	// volume (C:\ etc.), which is why relative paths works so well here as it avoids this confusion.
	traderRelativeConfigPath, e := s.botConfigsPath.Join(filenamePair.Trader).RelFromPath(s.kos.GetDotKelpWorkingDir())
	if e != nil {
		return fmt.Errorf("unable to get relative path of trader config file from basepath: %s", e)
	}

	stratRelativeConfigPath, e := s.botConfigsPath.Join(filenamePair.Strategy).RelFromPath(s.kos.GetDotKelpWorkingDir())
	if e != nil {
		return fmt.Errorf("unable to get relative path of strategy config file from basepath: %s", e)
	}

	logRelativePrefixPath, e := s.botLogsPath.Join(logPrefix).RelFromPath(s.kos.GetDotKelpWorkingDir())
	if e != nil {
		return fmt.Errorf("unable to get relative path of log prefix path from basepath: %s", e)
	}

	// prevent starting pubnet bots if pubnet is disabled
	var botConfig trader.BotConfig
	traderLoadReadPath := s.botConfigsPath.Join(filenamePair.Trader)
	e = config.Read(traderLoadReadPath.Native(), &botConfig)
	if e != nil {
		return fmt.Errorf("cannot read bot config at path '%s': %s", traderLoadReadPath.Native(), e)
	}
	isPubnetBot := botConfig.IsTradingSdex() && strings.TrimSuffix(botConfig.HorizonURL, "/") == strings.TrimSuffix(s.apiPubNet.HorizonURL, "/")
	if s.disablePubnet && isPubnetBot {
		return fmt.Errorf("cannnot start pubnet bots when pubnet is disabled")
	}

	command := fmt.Sprintf("trade -c %s -s %s -f %s -l %s --ui",
		traderRelativeConfigPath.Unix(),
		strategy,
		stratRelativeConfigPath.Unix(),
		logRelativePrefixPath.Unix(),
	)
	if iterations != nil {
		command = fmt.Sprintf("%s --iter %d", command, *iterations)
	}
	if s.noHeaders {
		command = fmt.Sprintf("%s --no-headers", command)
	}
	if s.ccxtRestUrl != "" {
		command = fmt.Sprintf("%s --ccxt-rest-url %s", command, s.ccxtRestUrl)
	}
	log.Printf("run command for bot '%s': %s\n", botName, command)

	p, e := s.runKelpCommandBackground(botName, command)
	if e != nil {
		return fmt.Errorf("could not start bot %s: %s", botName, e)
	}

	if p.Cmd == nil {
		return fmt.Errorf("kelpCommand (p.Cmd) was nil for bot '%s' with strategy '%s'", botName, strategy)
	}

	go func(kelpCommand *exec.Cmd, name string) {
		defer s.kos.SafeUnregister(name)

		e := kelpCommand.Wait()
		if e != nil {
			if strings.Contains(e.Error(), "signal: killed") {
				log.Printf("bot '%s' with strategy '%s' was stopped (most likely from UI action)", name, strategy)
				return
			}

			s.addKelpErrorToMap(makeKelpErrorResponseWrapper(
				errorTypeBot,
				botName,
				time.Now().UTC(),
				errorLevelError,
				fmt.Sprintf("unknown error in start bot command for bot '%s' with strategy '%s': %s", name, strategy, e),
			).KelpError)

			// set state to stopped
			s.abruptStoppedState(botName)

			// we don't want to continue because the bot didn't finish correctly
			return
		}

		log.Printf("finished start bot command for bot '%s' with strategy '%s'\n", name, strategy)
		if maybeFinishCallback != nil {
			maybeFinishCallback()
		}
	}(p.Cmd, botName)

	return nil
}

func (s *APIServer) abruptStoppedState(botName string) {
	// advance state from running to stopping
	e := s.kos.AdvanceBotState(botName, kelpos.BotStateRunning)
	if e != nil {
		s.addKelpErrorToMap(makeKelpErrorResponseWrapper(
			errorTypeBot,
			botName,
			time.Now().UTC(),
			errorLevelWarning,
			fmt.Sprintf("could not advance state from running to stopping: %s", e),
		).KelpError)
		return
	}

	// advance state from stopping to stopped
	e = s.kos.AdvanceBotState(botName, kelpos.BotStateStopping)
	if e != nil {
		s.addKelpErrorToMap(makeKelpErrorResponseWrapper(
			errorTypeBot,
			botName,
			time.Now().UTC(),
			errorLevelWarning,
			fmt.Sprintf("could not advance state from stopping to stopped: %s", e),
		).KelpError)
		return
	}
}
