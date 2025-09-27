package gowordstat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func post(url string, jsonData []byte, token string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	return body, nil
}

func request[responseType any](
	url string,
	payload any,
	token string,
) (responseType, error) {
	var result responseType
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return result, fmt.Errorf("encode: %w", err)
	}

	response, err := post(url, jsonData, token)
	if err != nil {
		return result, fmt.Errorf("post: %w", err)
	}

	if err := json.Unmarshal(response, &result); err != nil {
		return result, fmt.Errorf("decode: %w", err)
	}
	return result, nil
}
