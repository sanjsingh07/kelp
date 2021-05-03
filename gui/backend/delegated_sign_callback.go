package backend

import (
	"net/http"
	"net/url"
	"fmt"
	// "log"
	"io/ioutil"
	"strings"
	// "encoding/json"
	"github.com/stellar/kelp/plugins"
	// "github.com/stellar/kelp/api"

)

// type SignedURIStruct struct {
// 	SignedURI string `json:"uri"`
// }
// type delegatedSignSubmit struct {
// 	sdex                 *plugins.SDEX
// }

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

	// log.Printf("signedXDR only line 28: %s\n", decodedXDR)
	

	/*var signedURIStruct SignedURIStruct
	e = json.Unmarshal(signedTXBODY, &signedURIStruct)
	if e != nil {
		s.writeErrorJson(w, fmt.Sprintf("error unmarshaling json: %s; bodyString = %s", e, string(signedTXBODY)))   //working correctly with above struct
		return
	}
	fmt.Println("signedCallback requestJson line 33: \n", signedURIStruct, reflect.TypeOf(signedURIStruct), "\n") */

	// var dummy delegatedSignSubmit
	// (*plugins.SDEX).SubmitDelegatedTX(dummy.sdex, decodedXDR/*some augs*/); //asking for one more arguement --have (string)  --want (*plugins.SDEX, string)
	// t.sdex.SubmitDelegatedTX(signedXDRTobeSubmitted/*some augs*/);
	// dummy.sdex.SubmitDelegatedTX(signedXDRTobeSubmitted/*some augs*/);
	plugins.SubmitDelegatedTX(decodedXDR)



	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{ok}"))
}
