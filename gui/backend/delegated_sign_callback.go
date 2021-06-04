package backend

import (
	"net/http"
	"net/url"
	"fmt"
	// "log"
	"io/ioutil"
	"strings"
	"github.com/stellar/kelp/plugins"

)

func (s *APIServer) /*(t *delegatedSignSubmit)*/ signedCallback(w http.ResponseWriter, r *http.Request) {
	signedTXBODY, e := ioutil.ReadAll(r.Body)
	if e != nil {
		s.writeErrorJson(w, fmt.Sprintf("error when reading request input: %s\n", e))
		return
	}
	// log.Printf("signedCallback requestJson line 23: %s\n", string(signedTXBODY))

	stringToBeMani := string(signedTXBODY)
	strParts := strings.Split(stringToBeMani, "&")
	signedXDR := strings.Split(strParts[0], "xdr=")
	signedXDRTobeSubmitted := signedXDR[1]
	decodedXDR, err := url.QueryUnescape(signedXDRTobeSubmitted)
	if err != nil {
		s.writeErrorJson(w, fmt.Sprintf("error when decoding encodedSignedXDR: %s\n", e))
		return
	}
	network_passphrase := strings.Split(strParts[1], "+")
	HorizonUrl := strings.Split(network_passphrase[0], "network_passphrase=")
	decodedHorizonUrl, err := url.QueryUnescape(HorizonUrl[1])
	// decodedHorizonUrl := string(HorizonUrl[1])
	if err != nil {
		s.writeErrorJson(w, fmt.Sprintf("error when decoding encodedSignedXDR: %s\n", e))
		return
	}

	if(decodedHorizonUrl == "Test"){
		decodedHorizonUrl = "https://horizon-testnet.stellar.org"
	} else {
		decodedHorizonUrl = "https://horizon.stellar.org"
	}

	plugins.SubmitDelegatedTX(decodedXDR, decodedHorizonUrl)

	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Status: 200 OK"))
}