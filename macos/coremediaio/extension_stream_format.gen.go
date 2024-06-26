// Code generated by DarwinKit. DO NOT EDIT.

package coremediaio

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [ExtensionStreamFormat] class.
var ExtensionStreamFormatClass = _ExtensionStreamFormatClass{objc.GetClass("CMIOExtensionStreamFormat")}

type _ExtensionStreamFormatClass struct {
	objc.Class
}

// An interface definition for the [ExtensionStreamFormat] class.
type IExtensionStreamFormat interface {
	objc.IObject
	MinFrameDuration() coremedia.Time
	FormatDescription() coremedia.FormatDescriptionRef
	ValidFrameDurations() []foundation.Dictionary
	MaxFrameDuration() coremedia.Time
}

// An object that describes the format of a media stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamformat?language=objc
type ExtensionStreamFormat struct {
	objc.Object
}

func ExtensionStreamFormatFrom(ptr unsafe.Pointer) ExtensionStreamFormat {
	return ExtensionStreamFormat{
		Object: objc.ObjectFrom(ptr),
	}
}

func (e_ ExtensionStreamFormat) InitWithFormatDescriptionMaxFrameDurationMinFrameDurationValidFrameDurations(formatDescription coremedia.FormatDescriptionRef, maxFrameDuration coremedia.Time, minFrameDuration coremedia.Time, validFrameDurations []foundation.Dictionary) ExtensionStreamFormat {
	rv := objc.Call[ExtensionStreamFormat](e_, objc.Sel("initWithFormatDescription:maxFrameDuration:minFrameDuration:validFrameDurations:"), formatDescription, maxFrameDuration, minFrameDuration, validFrameDurations)
	return rv
}

// Creates a stream format with a format description and frame durations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamformat/3915898-initwithformatdescription?language=objc
func NewExtensionStreamFormatWithFormatDescriptionMaxFrameDurationMinFrameDurationValidFrameDurations(formatDescription coremedia.FormatDescriptionRef, maxFrameDuration coremedia.Time, minFrameDuration coremedia.Time, validFrameDurations []foundation.Dictionary) ExtensionStreamFormat {
	instance := ExtensionStreamFormatClass.Alloc().InitWithFormatDescriptionMaxFrameDurationMinFrameDurationValidFrameDurations(formatDescription, maxFrameDuration, minFrameDuration, validFrameDurations)
	instance.Autorelease()
	return instance
}

func (ec _ExtensionStreamFormatClass) StreamFormatWithFormatDescriptionMaxFrameDurationMinFrameDurationValidFrameDurations(formatDescription coremedia.FormatDescriptionRef, maxFrameDuration coremedia.Time, minFrameDuration coremedia.Time, validFrameDurations []foundation.Dictionary) ExtensionStreamFormat {
	rv := objc.Call[ExtensionStreamFormat](ec, objc.Sel("streamFormatWithFormatDescription:maxFrameDuration:minFrameDuration:validFrameDurations:"), formatDescription, maxFrameDuration, minFrameDuration, validFrameDurations)
	return rv
}

// Returns a new stream format with a format description and frame durations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamformat/3915901-streamformatwithformatdescriptio?language=objc
func ExtensionStreamFormat_StreamFormatWithFormatDescriptionMaxFrameDurationMinFrameDurationValidFrameDurations(formatDescription coremedia.FormatDescriptionRef, maxFrameDuration coremedia.Time, minFrameDuration coremedia.Time, validFrameDurations []foundation.Dictionary) ExtensionStreamFormat {
	return ExtensionStreamFormatClass.StreamFormatWithFormatDescriptionMaxFrameDurationMinFrameDurationValidFrameDurations(formatDescription, maxFrameDuration, minFrameDuration, validFrameDurations)
}

func (ec _ExtensionStreamFormatClass) Alloc() ExtensionStreamFormat {
	rv := objc.Call[ExtensionStreamFormat](ec, objc.Sel("alloc"))
	return rv
}

func (ec _ExtensionStreamFormatClass) New() ExtensionStreamFormat {
	rv := objc.Call[ExtensionStreamFormat](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionStreamFormat() ExtensionStreamFormat {
	return ExtensionStreamFormatClass.New()
}

func (e_ ExtensionStreamFormat) Init() ExtensionStreamFormat {
	rv := objc.Call[ExtensionStreamFormat](e_, objc.Sel("init"))
	return rv
}

// The minimum frame duration a stream supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamformat/3915900-minframeduration?language=objc
func (e_ ExtensionStreamFormat) MinFrameDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](e_, objc.Sel("minFrameDuration"))
	return rv
}

// A description of the format of the stream’s media samples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamformat/3915897-formatdescription?language=objc
func (e_ ExtensionStreamFormat) FormatDescription() coremedia.FormatDescriptionRef {
	rv := objc.Call[coremedia.FormatDescriptionRef](e_, objc.Sel("formatDescription"))
	return rv
}

// An array of frame durations the stream supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamformat/3915902-validframedurations?language=objc
func (e_ ExtensionStreamFormat) ValidFrameDurations() []foundation.Dictionary {
	rv := objc.Call[[]foundation.Dictionary](e_, objc.Sel("validFrameDurations"))
	return rv
}

// The maximum duration a stream supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamformat/3915899-maxframeduration?language=objc
func (e_ ExtensionStreamFormat) MaxFrameDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](e_, objc.Sel("maxFrameDuration"))
	return rv
}
