// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ProgressIndicator] class.
var ProgressIndicatorClass = _ProgressIndicatorClass{objc.GetClass("NSProgressIndicator")}

type _ProgressIndicatorClass struct {
	objc.Class
}

// An interface definition for the [ProgressIndicator] class.
type IProgressIndicator interface {
	IView
	IncrementBy(delta float64)
	StopAnimation(sender objc.IObject)
	StartAnimation(sender objc.IObject)
	SizeToFit()
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	Style() ProgressIndicatorStyle
	SetStyle(value ProgressIndicatorStyle)
	UsesThreadedAnimation() bool
	SetUsesThreadedAnimation(value bool)
	MinValue() float64
	SetMinValue(value float64)
	DoubleValue() float64
	SetDoubleValue(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	IsDisplayedWhenStopped() bool
	SetDisplayedWhenStopped(value bool)
	IsIndeterminate() bool
	SetIndeterminate(value bool)
}

// An interface that provides visual feedback to the user about the status of an ongoing task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator?language=objc
type ProgressIndicator struct {
	View
}

func ProgressIndicatorFrom(ptr unsafe.Pointer) ProgressIndicator {
	return ProgressIndicator{
		View: ViewFrom(ptr),
	}
}

func (pc _ProgressIndicatorClass) Alloc() ProgressIndicator {
	rv := objc.Call[ProgressIndicator](pc, objc.Sel("alloc"))
	return rv
}

func ProgressIndicator_Alloc() ProgressIndicator {
	return ProgressIndicatorClass.Alloc()
}

func (pc _ProgressIndicatorClass) New() ProgressIndicator {
	rv := objc.Call[ProgressIndicator](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewProgressIndicator() ProgressIndicator {
	return ProgressIndicatorClass.New()
}

func (p_ ProgressIndicator) Init() ProgressIndicator {
	rv := objc.Call[ProgressIndicator](p_, objc.Sel("init"))
	return rv
}

func (p_ ProgressIndicator) InitWithFrame(frameRect foundation.Rect) ProgressIndicator {
	rv := objc.Call[ProgressIndicator](p_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func ProgressIndicator_InitWithFrame(frameRect foundation.Rect) ProgressIndicator {
	return ProgressIndicatorClass.Alloc().InitWithFrame(frameRect)
}

// Advances the progress bar of a determinate progress indicator by the specified amount. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501155-incrementby?language=objc
func (p_ ProgressIndicator) IncrementBy(delta float64) {
	objc.Call[objc.Void](p_, objc.Sel("incrementBy:"), delta)
}

// Stops the animation of an indeterminate progress indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501141-stopanimation?language=objc
func (p_ ProgressIndicator) StopAnimation(sender objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("stopAnimation:"), sender)
}

// Starts the animation of an indeterminate progress indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501167-startanimation?language=objc
func (p_ ProgressIndicator) StartAnimation(sender objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("startAnimation:"), sender)
}

// This action method resizes the progress indicator to an appropriate size depending on the value of style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501144-sizetofit?language=objc
func (p_ ProgressIndicator) SizeToFit() {
	objc.Call[objc.Void](p_, objc.Sel("sizeToFit"))
}

// The size of the progress indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501157-controlsize?language=objc
func (p_ ProgressIndicator) ControlSize() ControlSize {
	rv := objc.Call[ControlSize](p_, objc.Sel("controlSize"))
	return rv
}

// The size of the progress indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501157-controlsize?language=objc
func (p_ ProgressIndicator) SetControlSize(value ControlSize) {
	objc.Call[objc.Void](p_, objc.Sel("setControlSize:"), value)
}

// The style of the progress indicator (bar or spinning). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501158-style?language=objc
func (p_ ProgressIndicator) Style() ProgressIndicatorStyle {
	rv := objc.Call[ProgressIndicatorStyle](p_, objc.Sel("style"))
	return rv
}

// The style of the progress indicator (bar or spinning). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501158-style?language=objc
func (p_ ProgressIndicator) SetStyle(value ProgressIndicatorStyle) {
	objc.Call[objc.Void](p_, objc.Sel("setStyle:"), value)
}

// A Boolean that indicates whether the progress indicator implements animation in a separate thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501160-usesthreadedanimation?language=objc
func (p_ ProgressIndicator) UsesThreadedAnimation() bool {
	rv := objc.Call[bool](p_, objc.Sel("usesThreadedAnimation"))
	return rv
}

// A Boolean that indicates whether the progress indicator implements animation in a separate thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501160-usesthreadedanimation?language=objc
func (p_ ProgressIndicator) SetUsesThreadedAnimation(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setUsesThreadedAnimation:"), value)
}

// The minimum value for the progress indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501169-minvalue?language=objc
func (p_ ProgressIndicator) MinValue() float64 {
	rv := objc.Call[float64](p_, objc.Sel("minValue"))
	return rv
}

// The minimum value for the progress indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501169-minvalue?language=objc
func (p_ ProgressIndicator) SetMinValue(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setMinValue:"), value)
}

// The value that indicates the current extent of the progress indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501173-doublevalue?language=objc
func (p_ ProgressIndicator) DoubleValue() float64 {
	rv := objc.Call[float64](p_, objc.Sel("doubleValue"))
	return rv
}

// The value that indicates the current extent of the progress indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501173-doublevalue?language=objc
func (p_ ProgressIndicator) SetDoubleValue(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setDoubleValue:"), value)
}

// The maximum value for the progress indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501165-maxvalue?language=objc
func (p_ ProgressIndicator) MaxValue() float64 {
	rv := objc.Call[float64](p_, objc.Sel("maxValue"))
	return rv
}

// The maximum value for the progress indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501165-maxvalue?language=objc
func (p_ ProgressIndicator) SetMaxValue(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setMaxValue:"), value)
}

// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501171-displayedwhenstopped?language=objc
func (p_ ProgressIndicator) IsDisplayedWhenStopped() bool {
	rv := objc.Call[bool](p_, objc.Sel("isDisplayedWhenStopped"))
	return rv
}

// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501171-displayedwhenstopped?language=objc
func (p_ ProgressIndicator) SetDisplayedWhenStopped(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setDisplayedWhenStopped:"), value)
}

// A Boolean that indicates whether the progress indicator is indeterminate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501146-indeterminate?language=objc
func (p_ ProgressIndicator) IsIndeterminate() bool {
	rv := objc.Call[bool](p_, objc.Sel("isIndeterminate"))
	return rv
}

// A Boolean that indicates whether the progress indicator is indeterminate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/1501146-indeterminate?language=objc
func (p_ ProgressIndicator) SetIndeterminate(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setIndeterminate:"), value)
}