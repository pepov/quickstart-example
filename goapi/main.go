package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

const asdf = "asdf 2"

func main() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(fmt.Sprintf("hello I'm the api server %s", asdf)))
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	})

	http.HandleFunc("/api/vote", func(w http.ResponseWriter, r *http.Request) {

		voterId := fmt.Sprintf("voter-%d", rand.Int31()%100000+900000)

		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		vote := r.FormValue("vote")
		data := struct {
			VoterID string `json:voter_id`
			Vote    string `json:vote`
		}{
			VoterID: voterId,
			Vote:    vote,
		}
		fmt.Fprintf(w, "received vote request for '%s' from voter id: '%s'", vote, voterId)
		out, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
		}
		_, err = w.Write(out)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusOK)

	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
