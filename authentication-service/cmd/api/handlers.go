package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	// authenticate user
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	fmt.Println("Start Logics!!")
	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	fmt.Println("Get Email!!", requestPayload)
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJson(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	fmt.Println("Match Password Start!!")
	valid, err := user.PasswordMatches(requestPayload.Password)
	fmt.Println("Match Password End!!")
	if err != nil || !valid {
		app.errorJson(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
	}

	app.writeJson(w, http.StatusAccepted, payload)
}
