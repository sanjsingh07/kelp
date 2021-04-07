package backend

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"fmt"
	"path/filepath"

	"github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
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

type AuthConfiguration struct {
    auth0enabled bool
	domain string
	audience string
}

var AuthConfig AuthConfiguration

func init() {
	absPath, _ := filepath.Abs("../gui/web/src/auth_config.json")
	file, _ := os.Open(absPath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	AuthConfig = AuthConfiguration{}
	err := decoder.Decode(&AuthConfig)
	if err != nil {
		fmt.Println("error:", err)
}
}

var JWTMiddlewareVar = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		// Verify 'iss' claim
		iss := "https://" +AuthConfig.domain+"/"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("Invalid issuer.")
		}
	  // Verify 'aud' claim
	  aud := AuthConfig.audience
	  checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
	  if !checkAud {
		  return token, errors.New("Invalid audience.")
	  }
	  

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
    resp, err := http.Get("https://"+AuthConfig.domain+"/.well-known/jwks.json")

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
