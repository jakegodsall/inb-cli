package notion

import (
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

func (*NotionClient) PostToInbox(task string) error {
	return fmt.Errorf("testing")
}