// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorchanging?language=objc
type PColorChanging interface {
	// optional
	ChangeColor(sender ColorPanel)
	HasChangeColor() bool
}

// ensure impl type implements protocol interface
var _ PColorChanging = (*ColorChangingObject)(nil)

// A concrete type for the [PColorChanging] protocol.
type ColorChangingObject struct {
	objc.Object
}

func (c_ ColorChangingObject) HasChangeColor() bool {
	return c_.RespondsToSelector(objc.Sel("changeColor:"))
}

// Sent to the first responder when the user selects a color in an NSColorPanel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorchanging/3005175-changecolor?language=objc
func (c_ ColorChangingObject) ChangeColor(sender ColorPanel) {
	objc.Call[objc.Void](c_, objc.Sel("changeColor:"), sender)
}
