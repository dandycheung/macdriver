// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A dynamically linkable representation of compiled shader code for a specific Metal device object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldynamiclibrary?language=objc
type PDynamicLibrary interface {
	// optional
	SerializeToURLError(url foundation.URL, error unsafe.Pointer) bool
	HasSerializeToURLError() bool

	// optional
	InstallName() string
	HasInstallName() bool

	// optional
	Device() DeviceObject
	HasDevice() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool
}

// ensure impl type implements protocol interface
var _ PDynamicLibrary = (*DynamicLibraryObject)(nil)

// A concrete type for the [PDynamicLibrary] protocol.
type DynamicLibraryObject struct {
	objc.Object
}

func (d_ DynamicLibraryObject) HasSerializeToURLError() bool {
	return d_.RespondsToSelector(objc.Sel("serializeToURL:error:"))
}

// Writes the contents of the dynamic library to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldynamiclibrary/3553982-serializetourl?language=objc
func (d_ DynamicLibraryObject) SerializeToURLError(url foundation.URL, error unsafe.Pointer) bool {
	rv := objc.Call[bool](d_, objc.Sel("serializeToURL:error:"), url, error)
	return rv
}

func (d_ DynamicLibraryObject) HasInstallName() bool {
	return d_.RespondsToSelector(objc.Sel("installName"))
}

// A file path for this dynamic library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldynamiclibrary/3553980-installname?language=objc
func (d_ DynamicLibraryObject) InstallName() string {
	rv := objc.Call[string](d_, objc.Sel("installName"))
	return rv
}

func (d_ DynamicLibraryObject) HasDevice() bool {
	return d_.RespondsToSelector(objc.Sel("device"))
}

// The Metal device object that created the dynamic library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldynamiclibrary/3553979-device?language=objc
func (d_ DynamicLibraryObject) Device() DeviceObject {
	rv := objc.Call[DeviceObject](d_, objc.Sel("device"))
	return rv
}

func (d_ DynamicLibraryObject) HasSetLabel() bool {
	return d_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that identifies the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldynamiclibrary/3553981-label?language=objc
func (d_ DynamicLibraryObject) SetLabel(value string) {
	objc.Call[objc.Void](d_, objc.Sel("setLabel:"), value)
}

func (d_ DynamicLibraryObject) HasLabel() bool {
	return d_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldynamiclibrary/3553981-label?language=objc
func (d_ DynamicLibraryObject) Label() string {
	rv := objc.Call[string](d_, objc.Sel("label"))
	return rv
}
