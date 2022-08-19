package devkit

import (
	"github.com/gofiber/fiber/v2"
	logRush "github.com/log-rush/server-devkit"
)

type Plugin struct {
	name          string
	LogHandler    logRush.HandleLog
	RouterHandler logRush.SetupRouter
}

func (p *Plugin) HandleLog(log logRush.Log) {
	p.LogHandler(log)
}

func (p *Plugin) SetupRouter(router fiber.Router) {
	p.RouterHandler(router)
}

func (p *Plugin) Name() string {
	return p.name
}

func NewPlugin(
	name string,
	logHandler logRush.HandleLog,
	routerHandler logRush.SetupRouter,
) logRush.Plugin {
	p := Plugin{
		name:          name,
		LogHandler:    logHandler,
		RouterHandler: routerHandler,
	}

	return &p
}

func NewLogPlugin(
	name string,
	logHandler logRush.HandleLog,
) logRush.LogPlugin {
	p := Plugin{
		name:       name,
		LogHandler: logHandler,
	}

	return &p
}

func NewRouterPlugin(
	name string,
	routerHandler logRush.SetupRouter,
) logRush.LogPlugin {
	p := Plugin{
		name:          name,
		RouterHandler: routerHandler,
	}

	return &p
}
