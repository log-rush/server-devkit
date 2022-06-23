package logRush

type Plugin interface {
	HandleLog(log Log)
}

type LogPlugin interface {
	HandleLog(log Log)
}

type HandleLog = func(log Log)
