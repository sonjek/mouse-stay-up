//go:build darwin
// +build darwin

package keyboard

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework Quartz
#include <Cocoa/Cocoa.h>
#include <Quartz/Quartz.h>

static CFMachPortRef eventTap;
static CFRunLoopSourceRef runLoopSource;

CGEventRef myCGEventCallback(CGEventTapProxy proxy, CGEventType type, CGEventRef event, void *refcon) {
    return NULL; // Block all keyboard events
}

void disableKeyboard() {
    eventTap = CGEventTapCreate(
        kCGHIDEventTap,
        kCGHeadInsertEventTap,
        kCGEventTapOptionDefault,
        CGEventMaskBit(kCGEventKeyDown) | CGEventMaskBit(kCGEventFlagsChanged),
        myCGEventCallback,
        NULL
    );

    if (!eventTap) {
        fprintf(stderr, "Failed to create event tap\n");
        exit(1);
    }

    runLoopSource = CFMachPortCreateRunLoopSource(kCFAllocatorDefault, eventTap, 0);
    CFRunLoopAddSource(CFRunLoopGetCurrent(), runLoopSource, kCFRunLoopCommonModes);
    CGEventTapEnable(eventTap, true);
    CFRunLoopRun();
}

void enableKeyboard() {
    if (eventTap) {
        CGEventTapEnable(eventTap, false);
        CFRunLoopRemoveSource(CFRunLoopGetCurrent(), runLoopSource, kCFRunLoopCommonModes);
        CFRelease(runLoopSource);
        CFRelease(eventTap);
    }
}
*/
import "C"

func NewKeyboardController() *Controller {
	return &Controller{
		KeyboardLocked: false,
		disableChan:    make(chan bool),
	}
}

func (c *Controller) LockKeyboard() {
	// If already locked, do nothing
	if c.KeyboardLocked {
		return
	}

	c.toggleLockUnlock()

	if c.disableChan == nil {
		c.disableChan = make(chan bool)
	}

	go func() {
		for {
			select {
			case <-c.disableChan:
				return
			default:
				C.disableKeyboard()
			}
		}
	}()
}

func (c *Controller) UnlockKeyboard() {
	// If already unlocked, do nothing
	if !c.KeyboardLocked {
		return
	}

	c.toggleLockUnlock()

	if c.disableChan != nil {
		close(c.disableChan)

		// Reset the channel to indicate it has been closed
		c.disableChan = nil
	}

	C.enableKeyboard()
}

func (c *Controller) toggleLockUnlock() {
	c.KeyboardLocked = !c.KeyboardLocked
}
