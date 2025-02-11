package api

import (
	"fmt"
	"io"
	"net/http"
)

type APIHandler struct {
	Url        string
	AccesToken string
}

func NewAPIHandler(url, accesToken string) *APIHandler {
	return &APIHandler{
		Url:        url,
		AccesToken: accesToken,
	}
}

func SendRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s", err.Error())
	}
	defer resp.Body.Close()

	// Read the response
	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return nil, fmt.Errorf("error reading response: %s", errRead.Error())
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error %s %s", resp.Status, body)
	}
	return body, nil
}
