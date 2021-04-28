package backend

import (
	"net/http"
	// "fmt"
	// "log"
	// "io/ioutil"
	// "encoding/json"
	// "github.com/stellar/kelp/plugins"
)

// type SignedXDRStruct struct {
// 	SignedXDR []string `json:"SignedXDR"`
// }

func (s *APIServer) signedCallback(w http.ResponseWriter, r *http.Request) {
	// signedTXBODY, e := ioutil.ReadAll(r.Body)
	// if e != nil {
	// 	s.writeErrorJson(w, fmt.Sprintf("error when reading request input: %s\n", e))
	// 	return
	// }
	// log.Printf("signedCallback requestJson: %s\n", string(signedTXBODY))

	// var signedXDRStruct SignedXDRStruct
	// e = json.Unmarshal(signedTXBODY, &signedXDRStruct)
	// if e != nil {
	// 	s.writeErrorJson(w, fmt.Sprintf("error unmarshaling json: %s; bodyString = %s", e, string(signedTXBODY)))
	// 	return
	// }

	// plugins.SubmitDelegatedTX("txeB64", true/*some augs*/);


	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("{}"))
}
