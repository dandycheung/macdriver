// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// A set of methods that drawer delegates implement to open, close, and resize the drawer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawerdelegate?language=objc
type PDrawerDelegate interface {
}

// A delegate implementation builder for the [PDrawerDelegate] protocol.
type DrawerDelegate struct {
}

// ensure impl type implements protocol interface
var _ PDrawerDelegate = (*DrawerDelegateObject)(nil)

// A concrete type for the [PDrawerDelegate] protocol.
type DrawerDelegateObject struct {
	objc.Object
}
