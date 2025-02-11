package api

import (
	"encoding/json"
	"fireflyiiiapi/api/models"
	"fmt"
	"net/http"
)

func (h *APIHandler) GetSystemInfo() (*models.SystemInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/about", h.Url), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err.Error())
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.AccesToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	body, err := SendRequest(req)
	if err != nil {
		return nil, err
	}
	systemInfo := models.SystemInfo{}

	err = json.Unmarshal(body, &systemInfo)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s", err.Error())
	}

	return &systemInfo, nil
}

func (h *APIHandler) GetUser() (*models.UserSingle, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/about/user", h.Url), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err.Error())
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.AccesToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	body, err := SendRequest(req)
	if err != nil {
		return nil, err
	}

	user := models.UserSingle{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s", err.Error())
	}
	return &user, nil
}
