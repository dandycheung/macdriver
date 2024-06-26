// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorabsolutedifference?language=objc
type PColorAbsoluteDifference interface {
	// optional
	SetInputImage2(value Image)
	HasSetInputImage2() bool

	// optional
	InputImage2() Image
	HasInputImage2() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool
}

// ensure impl type implements protocol interface
var _ PColorAbsoluteDifference = (*ColorAbsoluteDifferenceObject)(nil)

// A concrete type for the [PColorAbsoluteDifference] protocol.
type ColorAbsoluteDifferenceObject struct {
	objc.Object
}

func (c_ ColorAbsoluteDifferenceObject) HasSetInputImage2() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage2:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorabsolutedifference/3547105-inputimage2?language=objc
func (c_ ColorAbsoluteDifferenceObject) SetInputImage2(value Image) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage2:"), value)
}

func (c_ ColorAbsoluteDifferenceObject) HasInputImage2() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage2"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorabsolutedifference/3547105-inputimage2?language=objc
func (c_ ColorAbsoluteDifferenceObject) InputImage2() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage2"))
	return rv
}

func (c_ ColorAbsoluteDifferenceObject) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorabsolutedifference/3547104-inputimage?language=objc
func (c_ ColorAbsoluteDifferenceObject) SetInputImage(value Image) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), value)
}

func (c_ ColorAbsoluteDifferenceObject) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorabsolutedifference/3547104-inputimage?language=objc
func (c_ ColorAbsoluteDifferenceObject) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}
