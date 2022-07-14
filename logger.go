package logIt

type Logger struct {
	Handler Handler
}
type Handler interface {
	Log() error
}
