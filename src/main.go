package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jakegodsall/inb-cli/config"
	"github.com/jakegodsall/inb-cli/notion"
)

func main() {
	fmt.Println("Hello from CLI!")

	config, err := config.GetOrCreateConfigDir()
	if err != nil {
		err = fmt.Errorf("something went wrong: %v", err)
		fmt.Println(err)
		return
	}
	fmt.Println("config: ", config)

	notionApiKey := os.Getenv("NOTION_API_KEY")
	if notionApiKey == "" {
		err = fmt.Errorf("notion api key not set in the NOTION_API_KEY environment variable")
		log.Fatal(err)
		return
	}

	inboxId := os.Getenv("NOTION_INBOX_DATABASE_ID")
	if inboxId == "" {
		err = fmt.Errorf("notion inbox database id not set in the NOTION_INBOX_DATABASE_ID environment variable")
		log.Fatal(err)
		return
	}

	client := notion.NewNotionClient(notionApiKey, inboxId)

	data, err := client.GetDatabase()
	if err != nil {
		err = fmt.Errorf("something went wrong: %v", err)
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}