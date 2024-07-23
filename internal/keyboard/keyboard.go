package keyboard

type Controller struct {
	KeyboardLocked bool
	disableChan    chan bool
}

type Actions interface {
	LockKeyboard()
	UnlockKeyboard()
}

func NewController() *Controller {
	return NewKeyboardController()
}
