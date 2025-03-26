package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"your_project/services"
)

func GitHubLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := services.GetOAuthURL("github", "random-state")
	http.Redirect(w, r, url, http.StatusFound)
}

func GitHubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := services.ExchangeOAuthCode("github", code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var userData map[string]interface{}
	json.Unmarshal(body, &userData)


	fmt.Fprintf(w, "GitHub Login Successful! User: %+v", userData)
}
