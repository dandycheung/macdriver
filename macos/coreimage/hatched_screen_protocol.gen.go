// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a hatched screen filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen?language=objc
type PHatchedScreen interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetAngle(value float32)
	HasSetAngle() bool

	// optional
	Angle() float32
	HasAngle() bool

	// optional
	SetWidth(value float32)
	HasSetWidth() bool

	// optional
	Width() float32
	HasWidth() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool

	// optional
	SetSharpness(value float32)
	HasSetSharpness() bool

	// optional
	Sharpness() float32
	HasSharpness() bool
}

// ensure impl type implements protocol interface
var _ PHatchedScreen = (*HatchedScreenObject)(nil)

// A concrete type for the [PHatchedScreen] protocol.
type HatchedScreenObject struct {
	objc.Object
}

func (h_ HatchedScreenObject) HasSetInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228483-inputimage?language=objc
func (h_ HatchedScreenObject) SetInputImage(value Image) {
	objc.Call[objc.Void](h_, objc.Sel("setInputImage:"), value)
}

func (h_ HatchedScreenObject) HasInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228483-inputimage?language=objc
func (h_ HatchedScreenObject) InputImage() Image {
	rv := objc.Call[Image](h_, objc.Sel("inputImage"))
	return rv
}

func (h_ HatchedScreenObject) HasSetAngle() bool {
	return h_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228481-angle?language=objc
func (h_ HatchedScreenObject) SetAngle(value float32) {
	objc.Call[objc.Void](h_, objc.Sel("setAngle:"), value)
}

func (h_ HatchedScreenObject) HasAngle() bool {
	return h_.RespondsToSelector(objc.Sel("angle"))
}

// The angle of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228481-angle?language=objc
func (h_ HatchedScreenObject) Angle() float32 {
	rv := objc.Call[float32](h_, objc.Sel("angle"))
	return rv
}

func (h_ HatchedScreenObject) HasSetWidth() bool {
	return h_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The distance between lines in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228485-width?language=objc
func (h_ HatchedScreenObject) SetWidth(value float32) {
	objc.Call[objc.Void](h_, objc.Sel("setWidth:"), value)
}

func (h_ HatchedScreenObject) HasWidth() bool {
	return h_.RespondsToSelector(objc.Sel("width"))
}

// The distance between lines in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228485-width?language=objc
func (h_ HatchedScreenObject) Width() float32 {
	rv := objc.Call[float32](h_, objc.Sel("width"))
	return rv
}

func (h_ HatchedScreenObject) HasSetCenter() bool {
	return h_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the hatched screen pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228482-center?language=objc
func (h_ HatchedScreenObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](h_, objc.Sel("setCenter:"), value)
}

func (h_ HatchedScreenObject) HasCenter() bool {
	return h_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the hatched screen pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228482-center?language=objc
func (h_ HatchedScreenObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](h_, objc.Sel("center"))
	return rv
}

func (h_ HatchedScreenObject) HasSetSharpness() bool {
	return h_.RespondsToSelector(objc.Sel("setSharpness:"))
}

// The amount of sharpening to apply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228484-sharpness?language=objc
func (h_ HatchedScreenObject) SetSharpness(value float32) {
	objc.Call[objc.Void](h_, objc.Sel("setSharpness:"), value)
}

func (h_ HatchedScreenObject) HasSharpness() bool {
	return h_.RespondsToSelector(objc.Sel("sharpness"))
}

// The amount of sharpening to apply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228484-sharpness?language=objc
func (h_ HatchedScreenObject) Sharpness() float32 {
	rv := objc.Call[float32](h_, objc.Sel("sharpness"))
	return rv
}
