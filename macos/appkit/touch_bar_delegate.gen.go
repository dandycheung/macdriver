// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// A protocol that allows you to provide the items for a bar dynamically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbardelegate?language=objc
type PTouchBarDelegate interface {
	// optional
	TouchBarMakeItemForIdentifier(touchBar TouchBar, identifier TouchBarItemIdentifier) TouchBarItem
	HasTouchBarMakeItemForIdentifier() bool
}

// A delegate implementation builder for the [PTouchBarDelegate] protocol.
type TouchBarDelegate struct {
	_TouchBarMakeItemForIdentifier func(touchBar TouchBar, identifier TouchBarItemIdentifier) TouchBarItem
}

func (di *TouchBarDelegate) HasTouchBarMakeItemForIdentifier() bool {
	return di._TouchBarMakeItemForIdentifier != nil
}

// Asks the delegate object for the bar item for the specified bar and item identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbardelegate/2544851-touchbar?language=objc
func (di *TouchBarDelegate) SetTouchBarMakeItemForIdentifier(f func(touchBar TouchBar, identifier TouchBarItemIdentifier) TouchBarItem) {
	di._TouchBarMakeItemForIdentifier = f
}

// Asks the delegate object for the bar item for the specified bar and item identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbardelegate/2544851-touchbar?language=objc
func (di *TouchBarDelegate) TouchBarMakeItemForIdentifier(touchBar TouchBar, identifier TouchBarItemIdentifier) TouchBarItem {
	return di._TouchBarMakeItemForIdentifier(touchBar, identifier)
}

// ensure impl type implements protocol interface
var _ PTouchBarDelegate = (*TouchBarDelegateObject)(nil)

// A concrete type for the [PTouchBarDelegate] protocol.
type TouchBarDelegateObject struct {
	objc.Object
}

func (t_ TouchBarDelegateObject) HasTouchBarMakeItemForIdentifier() bool {
	return t_.RespondsToSelector(objc.Sel("touchBar:makeItemForIdentifier:"))
}

// Asks the delegate object for the bar item for the specified bar and item identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbardelegate/2544851-touchbar?language=objc
func (t_ TouchBarDelegateObject) TouchBarMakeItemForIdentifier(touchBar TouchBar, identifier TouchBarItemIdentifier) TouchBarItem {
	rv := objc.Call[TouchBarItem](t_, objc.Sel("touchBar:makeItemForIdentifier:"), touchBar, identifier)
	return rv
}
