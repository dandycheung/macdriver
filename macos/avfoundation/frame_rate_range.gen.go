// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [FrameRateRange] class.
var FrameRateRangeClass = _FrameRateRangeClass{objc.GetClass("AVFrameRateRange")}

type _FrameRateRangeClass struct {
	objc.Class
}

// An interface definition for the [FrameRateRange] class.
type IFrameRateRange interface {
	objc.IObject
	MinFrameRate() float64
	MaxFrameRate() float64
	MaxFrameDuration() coremedia.Time
	MinFrameDuration() coremedia.Time
}

// An immutable type that represents a range of valid frame rates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avframeraterange?language=objc
type FrameRateRange struct {
	objc.Object
}

func FrameRateRangeFrom(ptr unsafe.Pointer) FrameRateRange {
	return FrameRateRange{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FrameRateRangeClass) Alloc() FrameRateRange {
	rv := objc.Call[FrameRateRange](fc, objc.Sel("alloc"))
	return rv
}

func (fc _FrameRateRangeClass) New() FrameRateRange {
	rv := objc.Call[FrameRateRange](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFrameRateRange() FrameRateRange {
	return FrameRateRangeClass.New()
}

func (f_ FrameRateRange) Init() FrameRateRange {
	rv := objc.Call[FrameRateRange](f_, objc.Sel("init"))
	return rv
}

// The minimum frame rate supported by the range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avframeraterange/1389132-minframerate?language=objc
func (f_ FrameRateRange) MinFrameRate() float64 {
	rv := objc.Call[float64](f_, objc.Sel("minFrameRate"))
	return rv
}

// The maximum frame rate supported by the range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avframeraterange/1386988-maxframerate?language=objc
func (f_ FrameRateRange) MaxFrameRate() float64 {
	rv := objc.Call[float64](f_, objc.Sel("maxFrameRate"))
	return rv
}

// The maximum frame duration supported by the range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avframeraterange/1386786-maxframeduration?language=objc
func (f_ FrameRateRange) MaxFrameDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](f_, objc.Sel("maxFrameDuration"))
	return rv
}

// The minimum frame duration supported by the range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avframeraterange/1388420-minframeduration?language=objc
func (f_ FrameRateRange) MinFrameDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](f_, objc.Sel("minFrameDuration"))
	return rv
}
