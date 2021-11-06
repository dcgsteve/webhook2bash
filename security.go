package main

import "net/http"

func isAuthorised(r *http.Request) bool {
	authStatus := false
	token := ""

	// Grab token if it exists
	if len(r.Header[envTag]) > 0 {
		token = r.Header[envTag][0]
	}

	// Does token match acceptance criteria
	if token == envValidToken {
		authStatus = true
	}

	return authStatus
}
