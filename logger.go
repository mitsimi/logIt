package logIt

type Handler interface {
	Log() error
}
