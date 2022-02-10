package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Email struct {
	Email string
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./dist")))
	http.HandleFunc("/resetemail", emailReset)
	http.ListenAndServe(":8080", nil)
}

func emailReset(w http.ResponseWriter, r *http.Request) {
	var mail Email

	err := json.NewDecoder(r.Body).Decode(&mail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(mail)
	// There's a lot more that needs to be done (verification, max size, sending the email etc.), but this is just a quick prototype
}
