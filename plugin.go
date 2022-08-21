package logRush

import "github.com/gofiber/fiber/v2"

type Plugin interface {
	Name() string
	HandleLog(log Log)
	SetupRouter(router fiber.Router)
}

type LogPlugin interface {
	Name() string
	HandleLog(log Log)
}

type HandleLog = func(log Log)

type RouterPlugin interface {
	Name() string
	SetupRouter(router fiber.Router)
}

type SetupRouter = func(router fiber.Router)
