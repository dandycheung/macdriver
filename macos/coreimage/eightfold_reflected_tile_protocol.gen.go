// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure an eightfold reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile?language=objc
type PEightfoldReflectedTile interface {
	// optional
	SetWidth(value float32)
	HasSetWidth() bool

	// optional
	Width() float32
	HasWidth() bool

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
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// ensure impl type implements protocol interface
var _ PEightfoldReflectedTile = (*EightfoldReflectedTileObject)(nil)

// A concrete type for the [PEightfoldReflectedTile] protocol.
type EightfoldReflectedTileObject struct {
	objc.Object
}

func (e_ EightfoldReflectedTileObject) HasSetWidth() bool {
	return e_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228251-width?language=objc
func (e_ EightfoldReflectedTileObject) SetWidth(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setWidth:"), value)
}

func (e_ EightfoldReflectedTileObject) HasWidth() bool {
	return e_.RespondsToSelector(objc.Sel("width"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228251-width?language=objc
func (e_ EightfoldReflectedTileObject) Width() float32 {
	rv := objc.Call[float32](e_, objc.Sel("width"))
	return rv
}

func (e_ EightfoldReflectedTileObject) HasSetInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228250-inputimage?language=objc
func (e_ EightfoldReflectedTileObject) SetInputImage(value Image) {
	objc.Call[objc.Void](e_, objc.Sel("setInputImage:"), value)
}

func (e_ EightfoldReflectedTileObject) HasInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228250-inputimage?language=objc
func (e_ EightfoldReflectedTileObject) InputImage() Image {
	rv := objc.Call[Image](e_, objc.Sel("inputImage"))
	return rv
}

func (e_ EightfoldReflectedTileObject) HasSetAngle() bool {
	return e_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228248-angle?language=objc
func (e_ EightfoldReflectedTileObject) SetAngle(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setAngle:"), value)
}

func (e_ EightfoldReflectedTileObject) HasAngle() bool {
	return e_.RespondsToSelector(objc.Sel("angle"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228248-angle?language=objc
func (e_ EightfoldReflectedTileObject) Angle() float32 {
	rv := objc.Call[float32](e_, objc.Sel("angle"))
	return rv
}

func (e_ EightfoldReflectedTileObject) HasSetCenter() bool {
	return e_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228249-center?language=objc
func (e_ EightfoldReflectedTileObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](e_, objc.Sel("setCenter:"), value)
}

func (e_ EightfoldReflectedTileObject) HasCenter() bool {
	return e_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228249-center?language=objc
func (e_ EightfoldReflectedTileObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](e_, objc.Sel("center"))
	return rv
}
