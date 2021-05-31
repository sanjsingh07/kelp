package backend

import (
	"encoding/json"
	"errors"
	// "fmt"
	"net/http"
	"strings"
	"log"

	"github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/stellar/kelp/configStruct"
	"github.com/stellar/kelp/support/kelpos"
	"github.com/stellar/kelp/support/toml"
)

type Response struct {
	Message string `json:"message"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

var CustomConfigVarJWT configStruct.CustomConfigStruct

var userIDfromjwt string
var DataPath, UsersSpecificBot, BotConfigsPath, BotLogsPath *kelpos.OSPath

func callFromJWTMiddlewareVar() {
	kos := kelpos.GetKelpOS()
	trimmedID := strings.TrimLeft(userIDfromjwt, "auth0|")

	UserIDGlobal := "user_"+trimmedID
	DataPath := kos.GetDotKelpWorkingDir().Join("bot_data")

	UsersSpecificBot = DataPath.Join(UserIDGlobal)
	BotConfigsPath = UsersSpecificBot.Join("configs")
	configFilePath := DataPath.Join("custom_config.cfg")
	// if e != nil {
	// 	log.Printf("error creating custom_config.cfg: %s\n", e)
	// }
	e := toml.WriteFile(configFilePath.Native(), &CustomConfigVarJWT)
	if e != nil {
		log.Printf("error writing Custom Config toml file: %s\n", e)
		return
	}
	BotLogsPath = UsersSpecificBot.Join("logs")
}



var JWTMiddlewareVar = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		// Verify 'iss' claim
		iss := "https://" + CustomConfigVarJWT.Domain + "/"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("Invalid issuer.")
		}

		// Verify 'aud' claim
		// audPass := CustomConfigVarJWT.Audience
		// checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(audPass, false)
		// if !checkAud {
		// 	return token, errors.New("Invalid audience.")
		// }

		// fmt.Println("IM PRINTING FROM JWT RIGHT HERE: ", CustomConfigVarJWT.Domain)
		// fmt.Println(token.Claims.(jwt.MapClaims)["aud"])
		// fmt.Println(token.Claims.(jwt.MapClaims)["sub"])
		// fmt.Printf("aud1st = %T\n", token.Claims.(jwt.MapClaims)["aud"])
		// fmt.Printf("aud1st = %T\n", token.Claims.(jwt.MapClaims)["aud"].([]string))
		// fmt.Printf("aud1st = %T\n", token.Claims.(jwt.MapClaims)["aud"].(string))

		
		if(checkIss){
			userIDfromjwt = token.Claims.(jwt.MapClaims)["sub"].(string)
		}

		callFromJWTMiddlewareVar()
		
		cert, err := getPemCert(token)
		if err != nil {
			panic(err.Error())
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	},
	SigningMethod: jwt.SigningMethodRS256,
})

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://" + CustomConfigVarJWT.Domain + "/.well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}
