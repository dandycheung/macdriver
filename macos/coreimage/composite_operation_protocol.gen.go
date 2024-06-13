// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a composite operation filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicompositeoperation?language=objc
type PCompositeOperation interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetBackgroundImage(value Image)
	HasSetBackgroundImage() bool

	// optional
	BackgroundImage() Image
	HasBackgroundImage() bool
}

// ensure impl type implements protocol interface
var _ PCompositeOperation = (*CompositeOperationObject)(nil)

// A concrete type for the [PCompositeOperation] protocol.
type CompositeOperationObject struct {
	objc.Object
}

func (c_ CompositeOperationObject) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as a foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicompositeoperation/3228183-inputimage?language=objc
func (c_ CompositeOperationObject) SetInputImage(value Image) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), value)
}

func (c_ CompositeOperationObject) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as a foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicompositeoperation/3228183-inputimage?language=objc
func (c_ CompositeOperationObject) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ CompositeOperationObject) HasSetBackgroundImage() bool {
	return c_.RespondsToSelector(objc.Sel("setBackgroundImage:"))
}

// The image to use as a background image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicompositeoperation/3228182-backgroundimage?language=objc
func (c_ CompositeOperationObject) SetBackgroundImage(value Image) {
	objc.Call[objc.Void](c_, objc.Sel("setBackgroundImage:"), value)
}

func (c_ CompositeOperationObject) HasBackgroundImage() bool {
	return c_.RespondsToSelector(objc.Sel("backgroundImage"))
}

// The image to use as a background image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicompositeoperation/3228182-backgroundimage?language=objc
func (c_ CompositeOperationObject) BackgroundImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("backgroundImage"))
	return rv
}
