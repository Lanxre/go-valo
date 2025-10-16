package examples

import (
	"fmt"

	"github.com/lanxre/go-valo/client"
	"github.com/lanxre/go-valo/types"
)

func AccountExample() {
	cfg := types.Config{
		HenrikBaseURL: "https://api.henrikdev.xyz/valorant",
		HenrikAPIKey:  "<You token here>",
	}

	henrikClient := client.NewHenrikClient(cfg)
	data, err := henrikClient.GetAccountV1(map[string]string{
		"name": "lanore",
		"tag": "evil",
	})

	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("Name: %s, Tag: %s, LvL: %d\n", data.Name, data.Tag, data.AccountLevel)
}