package funfact

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetFunFact(n int) (string, error) {
	url := fmt.Sprintf("http://numbersapi.com/%d?json", n)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch fun fact: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-Ok HTTP status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read fun fact response: %v", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", fmt.Errorf("failed to parse fun fact fact JSON: %v", err)
	}

	if text, ok := data["text"].(string); ok {
		return text, nil
	}

	return "", fmt.Errorf("fun fact not found in response")
}
