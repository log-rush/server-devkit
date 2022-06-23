package devkit

import logRush "github.com/log-rush/server-devkit"

type Plugin struct {
	LogHandler logRush.HandleLog
}

func (p Plugin) HandleLog(log logRush.Log) {
	p.LogHandler(log)
}

func NewPlugin(
	logHandler logRush.HandleLog,
) logRush.Plugin {
	p := Plugin{
		LogHandler: logHandler,
	}

	return p
}

func NewLogPlugin(
	logHandler logRush.HandleLog,
) logRush.LogPlugin {
	p := Plugin{
		LogHandler: logHandler,
	}

	return p
}
