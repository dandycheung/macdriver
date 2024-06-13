// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure an exposure adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciexposureadjust?language=objc
type PExposureAdjust interface {
	// optional
	SetEV(value float32)
	HasSetEV() bool

	// optional
	EV() float32
	HasEV() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool
}

// ensure impl type implements protocol interface
var _ PExposureAdjust = (*ExposureAdjustObject)(nil)

// A concrete type for the [PExposureAdjust] protocol.
type ExposureAdjustObject struct {
	objc.Object
}

func (e_ ExposureAdjustObject) HasSetEV() bool {
	return e_.RespondsToSelector(objc.Sel("setEV:"))
}

// The amount to adjust the exposure of the image by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciexposureadjust/3228253-ev?language=objc
func (e_ ExposureAdjustObject) SetEV(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setEV:"), value)
}

func (e_ ExposureAdjustObject) HasEV() bool {
	return e_.RespondsToSelector(objc.Sel("EV"))
}

// The amount to adjust the exposure of the image by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciexposureadjust/3228253-ev?language=objc
func (e_ ExposureAdjustObject) EV() float32 {
	rv := objc.Call[float32](e_, objc.Sel("EV"))
	return rv
}

func (e_ ExposureAdjustObject) HasSetInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciexposureadjust/3228254-inputimage?language=objc
func (e_ ExposureAdjustObject) SetInputImage(value Image) {
	objc.Call[objc.Void](e_, objc.Sel("setInputImage:"), value)
}

func (e_ ExposureAdjustObject) HasInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciexposureadjust/3228254-inputimage?language=objc
func (e_ ExposureAdjustObject) InputImage() Image {
	rv := objc.Call[Image](e_, objc.Sel("inputImage"))
	return rv
}
