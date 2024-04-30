package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func AuthenticateUser(email, password string) (string, error) {
	formData := url.Values{}
	formData.Set("email", email)
	formData.Set("password", password)
	resp, err := http.Post("http://localhost:4000/connect-user", "application/json", strings.NewReader(formData.Encode()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erreur lors de l'authentification: %s", resp.Status)
	}

	var response struct {
		Message string `json:"message"`
		ID      string `json:"id"`
		Code    int    `json:"code"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	return response.ID, nil
}
