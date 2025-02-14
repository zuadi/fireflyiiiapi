package api

import (
	"encoding/json"
	"fireflyiiiapi/api/models"
	"fmt"
	"net/http"
	"net/url"
)

func (h *APIHandler) GetAccounts() (*models.AccountArray, error) {
	baseURL := fmt.Sprintf("%s/api/v1/accounts", h.Url)

	params := url.Values{}
	//params.Add("limit", "10")         // Number of items per page
	//params.Add("page", "1")           // Page number
	//params.Add("start", "2018-09-17") // Start date (YYYY-MM-DD)
	//params.Add("end", "2018-09-17")   // End date (YYYY-MM-DD)
	//params.Add("type", string(typ)) // Filter by transaction type

	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", baseURL, params.Encode()), nil)
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
	accounts := models.AccountArray{}

	err = json.Unmarshal(body, &accounts)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s", err.Error())
	}
	return &accounts, nil
}
