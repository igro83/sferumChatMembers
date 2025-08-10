package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func splitString(s string, interval int) string {
	var result strings.Builder
	for i, r := range s {
		result.WriteRune(r)
		if (i+1)%interval == 0 {
			result.WriteString("\n")
		}
	}
	return strings.TrimSuffix(result.String(), "\n")
}
func fetchParseJson(client *http.Client, url string, structToParse interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("ошибка запроса: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ошибка чтения тела: %w", err)
	}

	if err := json.Unmarshal(data, structToParse); err != nil {
		return fmt.Errorf("ошибка десериализации: %w", err)
	}
	return nil
}
