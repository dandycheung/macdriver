// Code generated by DarwinKit. DO NOT EDIT.

package quartz

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The IKDeviceBrowserViewDelegate defines the methods that the delegate of the IKDeviceBrowserView class can implement. All the methods are optional. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserviewdelegate?language=objc
type PDeviceBrowserViewDelegate interface {
	// optional
	DeviceBrowserViewDidEncounterError(deviceBrowserView DeviceBrowserView, error foundation.Error)
	HasDeviceBrowserViewDidEncounterError() bool
}

// A delegate implementation builder for the [PDeviceBrowserViewDelegate] protocol.
type DeviceBrowserViewDelegate struct {
	_DeviceBrowserViewDidEncounterError func(deviceBrowserView DeviceBrowserView, error foundation.Error)
}

func (di *DeviceBrowserViewDelegate) HasDeviceBrowserViewDidEncounterError() bool {
	return di._DeviceBrowserViewDidEncounterError != nil
}

// Invoked when the device browser encounters an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserviewdelegate/1443076-devicebrowserview?language=objc
func (di *DeviceBrowserViewDelegate) SetDeviceBrowserViewDidEncounterError(f func(deviceBrowserView DeviceBrowserView, error foundation.Error)) {
	di._DeviceBrowserViewDidEncounterError = f
}

// Invoked when the device browser encounters an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserviewdelegate/1443076-devicebrowserview?language=objc
func (di *DeviceBrowserViewDelegate) DeviceBrowserViewDidEncounterError(deviceBrowserView DeviceBrowserView, error foundation.Error) {
	di._DeviceBrowserViewDidEncounterError(deviceBrowserView, error)
}

// ensure impl type implements protocol interface
var _ PDeviceBrowserViewDelegate = (*DeviceBrowserViewDelegateObject)(nil)

// A concrete type for the [PDeviceBrowserViewDelegate] protocol.
type DeviceBrowserViewDelegateObject struct {
	objc.Object
}

func (d_ DeviceBrowserViewDelegateObject) HasDeviceBrowserViewDidEncounterError() bool {
	return d_.RespondsToSelector(objc.Sel("deviceBrowserView:didEncounterError:"))
}

// Invoked when the device browser encounters an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserviewdelegate/1443076-devicebrowserview?language=objc
func (d_ DeviceBrowserViewDelegateObject) DeviceBrowserViewDidEncounterError(deviceBrowserView DeviceBrowserView, error foundation.Error) {
	objc.Call[objc.Void](d_, objc.Sel("deviceBrowserView:didEncounterError:"), deviceBrowserView, error)
}
