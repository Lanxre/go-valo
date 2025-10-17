package examples

import (
	"fmt"

	"github.com/lanxre/go-valo/client"
	"github.com/lanxre/go-valo/types"
)

func AccountExample() {
	cfg := types.Config{
		HenrikAPIKey:  "<You token here>",
	}

	henrikClient := client.NewHenrikClient(cfg)
	data, err := henrikClient.GetAccountV1(types.ValorantAccountParams{
		Name: "lanore",
		Tag:  "evil",
	})

	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("Name: %s, Tag: %s, LvL: %d\n", data.Name, data.Tag, data.AccountLevel)
}
