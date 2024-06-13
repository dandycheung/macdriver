// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a Gabor gradients filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaborgradients?language=objc
type PGaborGradients interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool
}

// ensure impl type implements protocol interface
var _ PGaborGradients = (*GaborGradientsObject)(nil)

// A concrete type for the [PGaborGradients] protocol.
type GaborGradientsObject struct {
	objc.Object
}

func (g_ GaborGradientsObject) HasSetInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaborgradients/3325514-inputimage?language=objc
func (g_ GaborGradientsObject) SetInputImage(value Image) {
	objc.Call[objc.Void](g_, objc.Sel("setInputImage:"), value)
}

func (g_ GaborGradientsObject) HasInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaborgradients/3325514-inputimage?language=objc
func (g_ GaborGradientsObject) InputImage() Image {
	rv := objc.Call[Image](g_, objc.Sel("inputImage"))
	return rv
}
