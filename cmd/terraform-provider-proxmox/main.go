package main

import (
	"github.com/cracit/terraform-provider-proxmox/tree/master/cmd/terraform-provider-proxmox"
	"github.com/hashicorp/terraform/tree/master/plugin"
	"github.com/hashicorp/terraform/tree/master/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return proxmox.Provider()
		},
	})
}
