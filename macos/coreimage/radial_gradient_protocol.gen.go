// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a radial gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient?language=objc
type PRadialGradient interface {
	// optional
	SetColor1(value Color)
	HasSetColor1() bool

	// optional
	Color1() Color
	HasColor1() bool

	// optional
	SetRadius1(value float32)
	HasSetRadius1() bool

	// optional
	Radius1() float32
	HasRadius1() bool

	// optional
	SetRadius0(value float32)
	HasSetRadius0() bool

	// optional
	Radius0() float32
	HasRadius0() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool

	// optional
	SetColor0(value Color)
	HasSetColor0() bool

	// optional
	Color0() Color
	HasColor0() bool
}

// ensure impl type implements protocol interface
var _ PRadialGradient = (*RadialGradientObject)(nil)

// A concrete type for the [PRadialGradient] protocol.
type RadialGradientObject struct {
	objc.Object
}

func (r_ RadialGradientObject) HasSetColor1() bool {
	return r_.RespondsToSelector(objc.Sel("setColor1:"))
}

// The second color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228687-color1?language=objc
func (r_ RadialGradientObject) SetColor1(value Color) {
	objc.Call[objc.Void](r_, objc.Sel("setColor1:"), value)
}

func (r_ RadialGradientObject) HasColor1() bool {
	return r_.RespondsToSelector(objc.Sel("color1"))
}

// The second color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228687-color1?language=objc
func (r_ RadialGradientObject) Color1() Color {
	rv := objc.Call[Color](r_, objc.Sel("color1"))
	return rv
}

func (r_ RadialGradientObject) HasSetRadius1() bool {
	return r_.RespondsToSelector(objc.Sel("setRadius1:"))
}

// The radius of the ending circle to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228689-radius1?language=objc
func (r_ RadialGradientObject) SetRadius1(value float32) {
	objc.Call[objc.Void](r_, objc.Sel("setRadius1:"), value)
}

func (r_ RadialGradientObject) HasRadius1() bool {
	return r_.RespondsToSelector(objc.Sel("radius1"))
}

// The radius of the ending circle to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228689-radius1?language=objc
func (r_ RadialGradientObject) Radius1() float32 {
	rv := objc.Call[float32](r_, objc.Sel("radius1"))
	return rv
}

func (r_ RadialGradientObject) HasSetRadius0() bool {
	return r_.RespondsToSelector(objc.Sel("setRadius0:"))
}

// The radius of the starting circle to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228688-radius0?language=objc
func (r_ RadialGradientObject) SetRadius0(value float32) {
	objc.Call[objc.Void](r_, objc.Sel("setRadius0:"), value)
}

func (r_ RadialGradientObject) HasRadius0() bool {
	return r_.RespondsToSelector(objc.Sel("radius0"))
}

// The radius of the starting circle to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228688-radius0?language=objc
func (r_ RadialGradientObject) Radius0() float32 {
	rv := objc.Call[float32](r_, objc.Sel("radius0"))
	return rv
}

func (r_ RadialGradientObject) HasSetCenter() bool {
	return r_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228685-center?language=objc
func (r_ RadialGradientObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](r_, objc.Sel("setCenter:"), value)
}

func (r_ RadialGradientObject) HasCenter() bool {
	return r_.RespondsToSelector(objc.Sel("center"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228685-center?language=objc
func (r_ RadialGradientObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](r_, objc.Sel("center"))
	return rv
}

func (r_ RadialGradientObject) HasSetColor0() bool {
	return r_.RespondsToSelector(objc.Sel("setColor0:"))
}

// The first color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228686-color0?language=objc
func (r_ RadialGradientObject) SetColor0(value Color) {
	objc.Call[objc.Void](r_, objc.Sel("setColor0:"), value)
}

func (r_ RadialGradientObject) HasColor0() bool {
	return r_.RespondsToSelector(objc.Sel("color0"))
}

// The first color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228686-color0?language=objc
func (r_ RadialGradientObject) Color0() Color {
	rv := objc.Call[Color](r_, objc.Sel("color0"))
	return rv
}
