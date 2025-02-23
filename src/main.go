package main

import (
	"fmt"

	"github.com/jakegodsall/inb-cli/config"
)

func main() {
	fmt.Println("Hello from CLI!")

	config, err := config.CreateConfigDir()
	if err != nil {
		err = fmt.Errorf("something went wrong: %v", err)
		fmt.Println(err)
		return
	}
	fmt.Println("config: ", config)
}