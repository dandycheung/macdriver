// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a spotlight filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight?language=objc
type PSpotLight interface {
	// optional
	SetBrightness(value float32)
	HasSetBrightness() bool

	// optional
	Brightness() float32
	HasBrightness() bool

	// optional
	SetColor(value Color)
	HasSetColor() bool

	// optional
	Color() Color
	HasColor() bool

	// optional
	SetLightPosition(value Vector)
	HasSetLightPosition() bool

	// optional
	LightPosition() Vector
	HasLightPosition() bool

	// optional
	SetConcentration(value float32)
	HasSetConcentration() bool

	// optional
	Concentration() float32
	HasConcentration() bool

	// optional
	SetLightPointsAt(value Vector)
	HasSetLightPointsAt() bool

	// optional
	LightPointsAt() Vector
	HasLightPointsAt() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool
}

// ensure impl type implements protocol interface
var _ PSpotLight = (*SpotLightObject)(nil)

// A concrete type for the [PSpotLight] protocol.
type SpotLightObject struct {
	objc.Object
}

func (s_ SpotLightObject) HasSetBrightness() bool {
	return s_.RespondsToSelector(objc.Sel("setBrightness:"))
}

// The brightness of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228742-brightness?language=objc
func (s_ SpotLightObject) SetBrightness(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setBrightness:"), value)
}

func (s_ SpotLightObject) HasBrightness() bool {
	return s_.RespondsToSelector(objc.Sel("brightness"))
}

// The brightness of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228742-brightness?language=objc
func (s_ SpotLightObject) Brightness() float32 {
	rv := objc.Call[float32](s_, objc.Sel("brightness"))
	return rv
}

func (s_ SpotLightObject) HasSetColor() bool {
	return s_.RespondsToSelector(objc.Sel("setColor:"))
}

// The color of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228743-color?language=objc
func (s_ SpotLightObject) SetColor(value Color) {
	objc.Call[objc.Void](s_, objc.Sel("setColor:"), value)
}

func (s_ SpotLightObject) HasColor() bool {
	return s_.RespondsToSelector(objc.Sel("color"))
}

// The color of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228743-color?language=objc
func (s_ SpotLightObject) Color() Color {
	rv := objc.Call[Color](s_, objc.Sel("color"))
	return rv
}

func (s_ SpotLightObject) HasSetLightPosition() bool {
	return s_.RespondsToSelector(objc.Sel("setLightPosition:"))
}

// The x and y position of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228747-lightposition?language=objc
func (s_ SpotLightObject) SetLightPosition(value Vector) {
	objc.Call[objc.Void](s_, objc.Sel("setLightPosition:"), value)
}

func (s_ SpotLightObject) HasLightPosition() bool {
	return s_.RespondsToSelector(objc.Sel("lightPosition"))
}

// The x and y position of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228747-lightposition?language=objc
func (s_ SpotLightObject) LightPosition() Vector {
	rv := objc.Call[Vector](s_, objc.Sel("lightPosition"))
	return rv
}

func (s_ SpotLightObject) HasSetConcentration() bool {
	return s_.RespondsToSelector(objc.Sel("setConcentration:"))
}

// The size of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228744-concentration?language=objc
func (s_ SpotLightObject) SetConcentration(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setConcentration:"), value)
}

func (s_ SpotLightObject) HasConcentration() bool {
	return s_.RespondsToSelector(objc.Sel("concentration"))
}

// The size of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228744-concentration?language=objc
func (s_ SpotLightObject) Concentration() float32 {
	rv := objc.Call[float32](s_, objc.Sel("concentration"))
	return rv
}

func (s_ SpotLightObject) HasSetLightPointsAt() bool {
	return s_.RespondsToSelector(objc.Sel("setLightPointsAt:"))
}

// The x and y position that the spotlight points at. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228746-lightpointsat?language=objc
func (s_ SpotLightObject) SetLightPointsAt(value Vector) {
	objc.Call[objc.Void](s_, objc.Sel("setLightPointsAt:"), value)
}

func (s_ SpotLightObject) HasLightPointsAt() bool {
	return s_.RespondsToSelector(objc.Sel("lightPointsAt"))
}

// The x and y position that the spotlight points at. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228746-lightpointsat?language=objc
func (s_ SpotLightObject) LightPointsAt() Vector {
	rv := objc.Call[Vector](s_, objc.Sel("lightPointsAt"))
	return rv
}

func (s_ SpotLightObject) HasSetInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228745-inputimage?language=objc
func (s_ SpotLightObject) SetInputImage(value Image) {
	objc.Call[objc.Void](s_, objc.Sel("setInputImage:"), value)
}

func (s_ SpotLightObject) HasInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228745-inputimage?language=objc
func (s_ SpotLightObject) InputImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("inputImage"))
	return rv
}
