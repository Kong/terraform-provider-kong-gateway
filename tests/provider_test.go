package tests

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/kong/terraform-provider-kong-gateway/internal/provider"
)

var (
	providerFactory = map[string]func() (tfprotov6.ProviderServer, error){
		"kong-gateway": providerserver.NewProtocol6WithError(provider.New("999.99.9")()),
	}
)
