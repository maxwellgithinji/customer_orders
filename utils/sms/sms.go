package sms

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type SMS interface {
	SendMessage(phonenumbers string, message string) (string, error)
}

type message struct{}

func NewSMS() SMS {
	return &message{}
}

func (*message) SendMessage(phonenumbers string, message string) (string, error) {
	// Init env vars
	callbackUrl := os.Getenv("AFRICAS_TALKING_CALLBACK_URL")
	callbackResource := os.Getenv("AFRICAS_TALKING_CALLBACK_RESOURCE")
	apiKey := os.Getenv("AFRICAS_TALKING_API_KEY")
	username := os.Getenv("AFRICAS_TALKING_USERNAME")
	from := os.Getenv("AFRICAS_TALKING_SENDER_ID")

	// Set form data
	formdata := url.Values{}
	formdata.Set("username", username)
	formdata.Set("to", phonenumbers)
	formdata.Set("message", message)
	formdata.Set("from", from)

	// Setup Url
	u, _ := url.ParseRequestURI(callbackUrl)
	u.Path = callbackResource
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(formdata.Encode()))
	r.Header.Add("apiKey", apiKey)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)

	return resp.Status, nil
}
