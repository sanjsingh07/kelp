package backend

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"strings"
	// "encoding/json"
	// "github.com/stellar/kelp/plugins"
)

// type SignedURIStruct struct {
// 	SignedURI string `json:"uri"`
// }

func (s *APIServer) signedCallback(w http.ResponseWriter, r *http.Request) {
	signedTXBODY, e := ioutil.ReadAll(r.Body)
	if e != nil {
		s.writeErrorJson(w, fmt.Sprintf("error when reading request input: %s\n", e))
		return
	}
	log.Printf("signedCallback requestJson line 23: %s\n", string(signedTXBODY))

	stringToBeMani := string(signedTXBODY)
	strParts := strings.Split(stringToBeMani, "&")
	signedXDR := strings.Split(strParts[0], "xdr=")
	signedXDRTobeSubmitted := signedXDR[1]

	log.Printf("signedXDR only line 28: %s\n", signedXDRTobeSubmitted)
	

	/*var signedURIStruct SignedURIStruct
	e = json.Unmarshal(signedTXBODY, &signedURIStruct)
	if e != nil {
		s.writeErrorJson(w, fmt.Sprintf("error unmarshaling json: %s; bodyString = %s", e, string(signedTXBODY)))   //working correctly with above struct
		return
	}
	fmt.Println("signedCallback requestJson line 33: \n", signedURIStruct, reflect.TypeOf(signedURIStruct), "\n") */


	// plugins.SubmitDelegatedTX(signedXDRTobeSubmitted/*some augs*/);


	w.WriteHeader(http.StatusOK)
	// w.Write([]byte("{}"))
}
