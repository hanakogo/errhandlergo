package errhandlergo

type ErrorHandler struct {
	handler func(error)
}

func NewDefaultErrorHandler() *ErrorHandler {
	return &ErrorHandler{
		handler: func(err error) {},
	}
}
func (e *ErrorHandler) SetHandler(handler func(error)) {
	e.handler = handler
}

func (e *ErrorHandler) H(err error, handlerPlugins ...HandlerPlugin) {
	for _, plugin := range handlerPlugins {
		plugins[plugin](err)
	}
	e.handler(err)
}
