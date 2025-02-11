package models

type SystemInfo struct {
	Data struct {
		Version    string `json:"version,omitempty"`
		APIVersion string `json:"api_version,omitempty"`
		PHPVersion string `json:"php_version,omitempty"`
		OS         string `json:"os,omitempty"`
		MySQL      string `json:"driver,omitempty"`
	} `json:"data"`
}
