# GO-VALO

GO-VALO is API a go-based wrapper for the following [Henrik Valorant API](https://github.com/Henrik-3/unofficial-valorant-api)

This API is free and freely accessible for everyone. This is the first version. There could be some bugs, unexpected exceptions or similar.

### API key

You can request an API key on [Henrik's discord server](https://discord.com/invite/X3GaVkX2YN) <br>


### Example

Get Account Information
```go
package main

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
```