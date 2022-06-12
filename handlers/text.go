package handler

type TextHandler struct {
}

// HandleLog implements log.Handler.
func (th *TextHandler) Log() error {
	return nil
}
