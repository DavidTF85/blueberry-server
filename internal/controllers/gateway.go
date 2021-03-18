package controllers

import (
	"encoding/json"
	"net/http"

	"GitHub/DavidTF85/blueberry-server/utils"
	"GitHub/DavidTF85/blueberry-server/utils/models"

	cache "github.com/patrickmn/go-cache"
)

func (c *Controller) postRegister(w http.ResponseWriter, r *http.Request) {
	var requestData models.RegisterRequest
	data := r.Body
	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if requestData.Email == "" || requestData.Password == "" {
		http.Error(w, "Fields are not properly you dumdum", http.StatusBadRequest)
		return
	}

	c.cache.Set("e-mail", requestData.Email, cache.NoExpiration)

	var responseData = &models.RegisterResponse{
		Message: "You have successfully registered. Please login to continue!",
	}
	err = json.NewEncoder(w).Encode(&responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) postLogin(w http.ResponseWriter, r *http.Request) {
	var (
		hashedPassword string
		email          string
	)
	data := r.Body

	var requestData models.LoginRequest

	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if requestData.Email == "" || requestData.Password == "" {
		http.Error(w, "Fields are not properly you dumdum", http.StatusBadRequest)
		return
	}

	if cachedPasskey, found := c.cache.Get("hashed_password"); found {
		hashedPassword = cachedPasskey.(string)
	}
	cachedEmail, found := c.cache.Get("e-mail")
	if found {
		email = cachedEmail.(string)
	}

	if requestData.Email != email {
		http.Error(w, "This user does not match our records", http.StatusBadRequest)
		return
	}

	if utils.ComparePassword(w, r, []byte(Password), []byte(requestData.Password)) == false {
		http.Error(w, "The password incorrect--U idiot", http.StatusBadRequest)
		return
	}

	var responseData = &models.LoginResponse{
		Message: "Niceee U Logging in--welcomo Borat",
	}
	err = json.NewEncoder(w).Encode(&responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
