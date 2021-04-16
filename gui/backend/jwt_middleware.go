package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/stellar/kelp/configStruct"
	"github.com/stellar/kelp/support/kelpos"
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

var CustomConfigVar configStruct.CustomConfigStruct

var userIDfromjwt string
var UsersSpecificBot, BotConfigsPath, BotLogsPath *kelpos.OSPath

func callFromJWTMiddlewareVar() {
	kos := kelpos.GetKelpOS()
	trimmedID := strings.TrimLeft(userIDfromjwt, "auth0|")
	// fmt.Println("from list_bots file: ", trimmedID)

	UserIDGlobal := "user_"+trimmedID
	dataPath := kos.GetDotKelpWorkingDir().Join("bot_data")
	// fmt.Println("Printing from list_bots.go file:", UserIDGlobal)

	UsersSpecificBot = dataPath.Join(UserIDGlobal)
	BotConfigsPath = UsersSpecificBot.Join("configs")
	BotLogsPath = UsersSpecificBot.Join("logs")

	fmt.Println("Printing from jwtMiddleware file: line 47-48", UsersSpecificBot.AsString())
	fmt.Println("Printing from jwtMiddleware file: line 47-48", BotConfigsPath.AsString())
	fmt.Println("Printing from jwtMiddleware file: line 47-48", BotLogsPath.AsString())
}



var JWTMiddlewareVar = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		// Verify 'iss' claim
		iss := "https://" + CustomConfigVar.Domain + "/"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("Invalid issuer.")
		}

		// Verify 'aud' claim
		// audPass := CustomConfigVar.Audience
		// checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(audPass, false)
		// if !checkAud {
		// 	return token, errors.New("Invalid audience.")
		// }

		// fmt.Println("IM PRINTING FROM JWT RIGHT HERE: ", CustomConfigVar.Domain)
		// fmt.Println(token.Claims.(jwt.MapClaims)["aud"])
		// fmt.Println(token.Claims.(jwt.MapClaims)["sub"])
		// fmt.Printf("aud1st = %T\n", token.Claims.(jwt.MapClaims)["aud"])
		// fmt.Printf("aud1st = %T\n", token.Claims.(jwt.MapClaims)["aud"].([]string))
		// fmt.Printf("aud1st = %T\n", token.Claims.(jwt.MapClaims)["aud"].(string))

		
		if(checkIss){
			// User_ID_struct_var.User_id_token = token.Claims.(jwt.MapClaims)["sub"].(string)
			userIDfromjwt = token.Claims.(jwt.MapClaims)["sub"].(string)
			fmt.Println("Printing from jwtMiddleware file: line 85", userIDfromjwt)
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
	resp, err := http.Get("https://" + CustomConfigVar.Domain + "/.well-known/jwks.json")

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
