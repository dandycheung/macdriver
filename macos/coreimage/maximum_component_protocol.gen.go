// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a maximum component filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimaximumcomponent?language=objc
type PMaximumComponent interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool
}

// ensure impl type implements protocol interface
var _ PMaximumComponent = (*MaximumComponentObject)(nil)

// A concrete type for the [PMaximumComponent] protocol.
type MaximumComponentObject struct {
	objc.Object
}

func (m_ MaximumComponentObject) HasSetInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimaximumcomponent/3228555-inputimage?language=objc
func (m_ MaximumComponentObject) SetInputImage(value Image) {
	objc.Call[objc.Void](m_, objc.Sel("setInputImage:"), value)
}

func (m_ MaximumComponentObject) HasInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimaximumcomponent/3228555-inputimage?language=objc
func (m_ MaximumComponentObject) InputImage() Image {
	rv := objc.Call[Image](m_, objc.Sel("inputImage"))
	return rv
}
