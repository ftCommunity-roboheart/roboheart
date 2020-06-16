package servicemanager

import (
	"github.com/ftCommunity/roboheart/internal/service"
	"github.com/ftCommunity/roboheart/internal/services/core/acm"
	"github.com/ftCommunity/roboheart/internal/services/core/config"
	"github.com/ftCommunity/roboheart/internal/services/core/fwver"
	"github.com/ftCommunity/roboheart/internal/services/core/power"
	"github.com/ftCommunity/roboheart/internal/services/core/releasever"
	"github.com/ftCommunity/roboheart/internal/services/core/web"
)

var (
	services = []service.Service{
		config.Service,
		acm.Service,
		fwver.Service,
		relver.Service,
		power.Service,
		web.Service,
	}
)
