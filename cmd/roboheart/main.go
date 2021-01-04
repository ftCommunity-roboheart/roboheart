package main

import (
	"github.com/servicemngr/core/package/manifest"
	"github.com/servicemngr/core/package/servicemngr"
)

var ServiceProviders = [][]manifest.ServiceManifest{
	//svcs_core.Services,
}

func main() {
	servicemngr.Run("roboheart", "roboheart", ServiceProviders)
}
