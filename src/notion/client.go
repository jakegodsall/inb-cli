package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type NotionClient struct {
	ApiKey  string
	InboxId string
}

func NewNotionClient(apiKey string, inboxId string) *NotionClient {
	return &NotionClient{
		ApiKey:  apiKey,
		InboxId: inboxId,
	}
}

func (nc *NotionClient) GetDatabase() ([]byte, error) {
	basePath, err := url.Parse("https://api.notion.com/v1/databases/")
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}
	url := basePath.ResolveReference(&url.URL{Path: nc.InboxId}).String()
	log.Println("Request URL: " + url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer " + nc.ApiKey)
	req.Header.Set("Notion-Version", "2022-06-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("notion API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

func (nc *NotionClient) PostToInbox(task string) error {
	basePath, err := url.Parse("https://api.notion.com/v1/pages")
	if err != nil {
		return fmt.Errorf("failed to parse base URL: %w", err)
	}
	log.Println("Request URL: " + basePath.String())

	payload := map[string]interface{} {
		"parent": map[string]string {
			"database_id": nc.InboxId,
		},
		"icon": map[string]string {
			"emoji": "🔥",
		},
		"properties": map[string]interface{} {
			"Task": map[string]interface{} {
				"title": []map[string]interface{} {
					{
						"text": map[string]string {
							"content": task,
						},
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	req, err := http.NewRequest("POST", basePath.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer " + nc.ApiKey)
	req.Header.Set("Notion-Version", "2022-06-28")
	req.Header.Set("Content-Type", "application/json")

	log.Println("Sending task to Notion:", string(jsonData))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("notion API request failed with status: %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}