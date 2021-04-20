package backend

import (
	"fmt"
	"net/http"
)

func makeSep7URI(xdrBase64 string, pubkey string) string {
    return fmt.Sprintf("web+stellar:tx?xdr=%s&pubkey=%s", xdrBase64, pubkey)
}

func (s *APIServer) delegatedSigningCallback(w http.ResponseWriter, r *http.Request) {

	
	// w.Write([]byte("smth goes here, but idk what lol"))
}
