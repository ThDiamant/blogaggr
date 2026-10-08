package main

import (
	"blogaggr/internal"
	"fmt"
)

func main() {
	config, err := internal.Read()
	if err != nil {
		fmt.Println(err)
	}

	config.CurrentUserName = "theodor"

	config.SetUser("theodor")

	config, err = internal.Read()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config)
}
