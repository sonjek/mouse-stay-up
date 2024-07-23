//go:build !darwin
// +build !darwin

package keyboard

import "C"
import "fmt"

func NewKeyboardController() *Controller {
	return &Controller{}
}

func (c *Controller) LockKeyboard() {
	fmt.Println("LockKeyboard called on non-macOS platform. No operation performed.")
}

func (c *Controller) UnlockKeyboard() {
	fmt.Println("UnlockKeyboard called on non-macOS platform. No operation performed.")
}
