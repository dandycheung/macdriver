// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a noise reduction filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinoisereduction?language=objc
type PNoiseReduction interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetNoiseLevel(value float32)
	HasSetNoiseLevel() bool

	// optional
	NoiseLevel() float32
	HasNoiseLevel() bool

	// optional
	SetSharpness(value float32)
	HasSetSharpness() bool

	// optional
	Sharpness() float32
	HasSharpness() bool
}

// ensure impl type implements protocol interface
var _ PNoiseReduction = (*NoiseReductionObject)(nil)

// A concrete type for the [PNoiseReduction] protocol.
type NoiseReductionObject struct {
	objc.Object
}

func (n_ NoiseReductionObject) HasSetInputImage() bool {
	return n_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinoisereduction/3228595-inputimage?language=objc
func (n_ NoiseReductionObject) SetInputImage(value Image) {
	objc.Call[objc.Void](n_, objc.Sel("setInputImage:"), value)
}

func (n_ NoiseReductionObject) HasInputImage() bool {
	return n_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinoisereduction/3228595-inputimage?language=objc
func (n_ NoiseReductionObject) InputImage() Image {
	rv := objc.Call[Image](n_, objc.Sel("inputImage"))
	return rv
}

func (n_ NoiseReductionObject) HasSetNoiseLevel() bool {
	return n_.RespondsToSelector(objc.Sel("setNoiseLevel:"))
}

// The amount of noise reduction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinoisereduction/3228596-noiselevel?language=objc
func (n_ NoiseReductionObject) SetNoiseLevel(value float32) {
	objc.Call[objc.Void](n_, objc.Sel("setNoiseLevel:"), value)
}

func (n_ NoiseReductionObject) HasNoiseLevel() bool {
	return n_.RespondsToSelector(objc.Sel("noiseLevel"))
}

// The amount of noise reduction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinoisereduction/3228596-noiselevel?language=objc
func (n_ NoiseReductionObject) NoiseLevel() float32 {
	rv := objc.Call[float32](n_, objc.Sel("noiseLevel"))
	return rv
}

func (n_ NoiseReductionObject) HasSetSharpness() bool {
	return n_.RespondsToSelector(objc.Sel("setSharpness:"))
}

// The sharpness of the final image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinoisereduction/3228597-sharpness?language=objc
func (n_ NoiseReductionObject) SetSharpness(value float32) {
	objc.Call[objc.Void](n_, objc.Sel("setSharpness:"), value)
}

func (n_ NoiseReductionObject) HasSharpness() bool {
	return n_.RespondsToSelector(objc.Sel("sharpness"))
}

// The sharpness of the final image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinoisereduction/3228597-sharpness?language=objc
func (n_ NoiseReductionObject) Sharpness() float32 {
	rv := objc.Call[float32](n_, objc.Sel("sharpness"))
	return rv
}
