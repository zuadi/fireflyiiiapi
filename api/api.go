package api

import (
	"fireflyiiiapi/api/models"
	"fireflyiiiapi/utils"
	"fmt"
	"io"
	"net/http"
)

// Secret key (must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256)
const key = "3c28tIg5JZGYK6cZzwqW6vCoDRWyItbK" // 32 bytes for AES-256

type APIHandler struct {
	Url             string
	AccesToken      string
	Deposits        models.TransactionStore
	Withdrawals     models.TransactionStore
	Transfers       models.TransactionStore
	Reconsiliations models.TransactionStore
	OpeningBalances models.TransactionStore
}

func NewAPIHandler(url, accesToken string) *APIHandler {
	return &APIHandler{
		Url:        url,
		AccesToken: accesToken,
	}
}

func SaveTokenToFile(file string, token string) error {
	// Encrypt the token
	encryptedData, err := utils.Encrypt([]byte(token), []byte(key))
	if err != nil {
		return fmt.Errorf("error encrypting: %s", err.Error())
	}

	// Save the encrypted data to a binary file
	err = utils.SaveToBinaryFile(file, encryptedData)
	if err != nil {
		return fmt.Errorf("error saving to file: %s", err.Error())
	}
	return nil
}

func ReadTokenToFile(file string) (string, error) {
	// Read the encrypted data from the binary file
	encryptedData, err := utils.ReadFromBinaryFile(file)
	if err != nil {
		return "", fmt.Errorf("error reading from file: %s", err.Error())
	}

	// Decrypt the data
	decryptedData, err := utils.Decrypt(encryptedData, []byte(key))
	if err != nil {
		return "", fmt.Errorf("error decrypting: %s", err.Error())
	}

	return string(decryptedData), nil
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
