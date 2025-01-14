package live

const (
	// ScopeGrafana contains builtin features of Grafana Core.
	ScopeGrafana = "grafana"
	// ScopePlugin passes control to a plugin.
	ScopePlugin = "plugin"
	// ScopeDatasource passes control to a datasource plugin.
	ScopeDatasource = "ds"
	// ScopeStream is a managed data frame stream
	ScopeStream = "stream"
	// ScopePush allows sending data into managed streams.  It does not support subscriptions
	ScopePush = "push"
)
