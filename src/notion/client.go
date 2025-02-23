package notion

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

type NotionClient struct {
	ApiKey string
}

func (*NotionClient) GetDatabase() ([]byte, error) {
	inboxId := os.Getenv("NOTION_INBOX_ID")
	if inboxId == "" {
		return nil, fmt.Errorf("notion inbox database id not stored in the NOTION_INBOX_ID environment variable")
	}

	basePath := "https://api.notion.com/v1/databases"
	url := path.Join(basePath, inboxId)

	authToken := os.Getenv("NOTION_API_KEY")
	if authToken == "" {
		return nil, fmt.Errorf("api key not stored in the NOTION_API_KEY environment variable")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer " + url)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}