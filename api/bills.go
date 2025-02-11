package api

//func (h *APIHandler) AllBills() (*models.Bills, error) {
//	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/bills", h.Url), nil)
//	if err != nil {
//		return nil, fmt.Errorf("error creating request: %s", err.Error())
//	}
//
//	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.AccesToken))
//	req.Header.Set("Content-Type", "application/vnd.api+json")
//	req.Header.Set("Accept", "application/vnd.api+json")
//
//	// Send the request
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return nil, fmt.Errorf("error sending request: %s", err.Error())
//	}
//	defer resp.Body.Close()
//
//	// Read the response
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return nil, fmt.Errorf("error reading response: %s", err.Error())
//	}
//
//	// Print response
//	var data interface{}
//
//	err = json.Unmarshal(body, &data)
//	if err != nil {
//		return nil, fmt.Errorf("error reading response: %s", err.Error())
//	}
//	fmt.Println(123, data)
//	return nil, nil
//}
