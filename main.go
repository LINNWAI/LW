package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"utils"
	"vro"
)

func main() {
	utils.InitLog()
	opts := plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return vro.Provider()
		},
	}

	plugin.Serve(&opts)
}
