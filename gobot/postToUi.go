package main

import (
	"net/http"
	"net/url"
)

func postToUi(message string) error {
	uiUrl := "http://steve-ui.azurewebsites.net/api/messages/v1/add"

	http.PostForm(uiUrl, url.Values{"message": {message}})

	return nil
}
