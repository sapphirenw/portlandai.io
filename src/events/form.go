package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

var emailRE = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

func EventForm(w http.ResponseWriter, r *http.Request) (string, int) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return "There was an issue reading the body", http.StatusInternalServerError
	}

	values, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		fmt.Println(err)
		return "There was an issue reading the body", http.StatusInternalServerError
	}

	// validate the form
	email := values.Get("email")
	if email == "" {
		return "Email cannot be empty", 400
	}

	if !emailRE.MatchString(email) {
		return "Invalid Email", 400
	}

	message := values.Get("message")
	if message == "" {
		return "Your message cannot be empty", 400
	}

	data := map[string]interface{}{
		"email":   email,
		"message": message,
	}

	encoded, _ := json.Marshal(data)

	resp, err := http.Post("https://pocketbase.sapphirenw.com/api/collections/portlandai_form/records", "application/json", bytes.NewBuffer(encoded))
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return "There was an issue sending the request", 500
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return "Successfully sent your message", 200
	} else {
		encoded, _ := json.Marshal(resp.Body)
		fmt.Println(resp.Status)
		fmt.Println(string(encoded))
		return "There was an issue sending the request", 500
	}
}
