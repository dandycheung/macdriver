// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// A set of methods used to associate a unique identifier with objects in your user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemidentification?language=objc
type PUserInterfaceItemIdentification interface {
	// optional
	SetIdentifier(value UserInterfaceItemIdentifier)
	HasSetIdentifier() bool

	// optional
	Identifier() UserInterfaceItemIdentifier
	HasIdentifier() bool
}

// ensure impl type implements protocol interface
var _ PUserInterfaceItemIdentification = (*UserInterfaceItemIdentificationObject)(nil)

// A concrete type for the [PUserInterfaceItemIdentification] protocol.
type UserInterfaceItemIdentificationObject struct {
	objc.Object
}

func (u_ UserInterfaceItemIdentificationObject) HasSetIdentifier() bool {
	return u_.RespondsToSelector(objc.Sel("setIdentifier:"))
}

// A string that identifies the user interface item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemidentification/1396829-identifier?language=objc
func (u_ UserInterfaceItemIdentificationObject) SetIdentifier(value UserInterfaceItemIdentifier) {
	objc.Call[objc.Void](u_, objc.Sel("setIdentifier:"), value)
}

func (u_ UserInterfaceItemIdentificationObject) HasIdentifier() bool {
	return u_.RespondsToSelector(objc.Sel("identifier"))
}

// A string that identifies the user interface item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemidentification/1396829-identifier?language=objc
func (u_ UserInterfaceItemIdentificationObject) Identifier() UserInterfaceItemIdentifier {
	rv := objc.Call[UserInterfaceItemIdentifier](u_, objc.Sel("identifier"))
	return rv
}
