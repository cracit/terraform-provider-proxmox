package main

import (
	"github.com/cracit/terraform-provider-proxmox/tree/master/proxmox"
	"github.com/hashicorp/terraform/tree/master/plugin"
	"github.com/hashicorp/terraform/tree/master/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProvisionerFunc: func() terraform.ResourceProvisioner {
			return proxmox.Provisioner()
		},
	})
}
