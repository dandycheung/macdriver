// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a morphology rectangle maximum filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectanglemaximum?language=objc
type PMorphologyRectangleMaximum interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetHeight(value float32)
	HasSetHeight() bool

	// optional
	Height() float32
	HasHeight() bool

	// optional
	SetWidth(value float32)
	HasSetWidth() bool

	// optional
	Width() float32
	HasWidth() bool
}

// ensure impl type implements protocol interface
var _ PMorphologyRectangleMaximum = (*MorphologyRectangleMaximumObject)(nil)

// A concrete type for the [PMorphologyRectangleMaximum] protocol.
type MorphologyRectangleMaximumObject struct {
	objc.Object
}

func (m_ MorphologyRectangleMaximumObject) HasSetInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectanglemaximum/3228584-inputimage?language=objc
func (m_ MorphologyRectangleMaximumObject) SetInputImage(value Image) {
	objc.Call[objc.Void](m_, objc.Sel("setInputImage:"), value)
}

func (m_ MorphologyRectangleMaximumObject) HasInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectanglemaximum/3228584-inputimage?language=objc
func (m_ MorphologyRectangleMaximumObject) InputImage() Image {
	rv := objc.Call[Image](m_, objc.Sel("inputImage"))
	return rv
}

func (m_ MorphologyRectangleMaximumObject) HasSetHeight() bool {
	return m_.RespondsToSelector(objc.Sel("setHeight:"))
}

// The height, in pixels, of the morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectanglemaximum/3228583-height?language=objc
func (m_ MorphologyRectangleMaximumObject) SetHeight(value float32) {
	objc.Call[objc.Void](m_, objc.Sel("setHeight:"), value)
}

func (m_ MorphologyRectangleMaximumObject) HasHeight() bool {
	return m_.RespondsToSelector(objc.Sel("height"))
}

// The height, in pixels, of the morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectanglemaximum/3228583-height?language=objc
func (m_ MorphologyRectangleMaximumObject) Height() float32 {
	rv := objc.Call[float32](m_, objc.Sel("height"))
	return rv
}

func (m_ MorphologyRectangleMaximumObject) HasSetWidth() bool {
	return m_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width, in pixels, of the morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectanglemaximum/3228585-width?language=objc
func (m_ MorphologyRectangleMaximumObject) SetWidth(value float32) {
	objc.Call[objc.Void](m_, objc.Sel("setWidth:"), value)
}

func (m_ MorphologyRectangleMaximumObject) HasWidth() bool {
	return m_.RespondsToSelector(objc.Sel("width"))
}

// The width, in pixels, of the morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectanglemaximum/3228585-width?language=objc
func (m_ MorphologyRectangleMaximumObject) Width() float32 {
	rv := objc.Call[float32](m_, objc.Sel("width"))
	return rv
}
