// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a keystone correction combined filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikeystonecorrectioncombined?language=objc
type PKeystoneCorrectionCombined interface {
	// optional
	SetFocalLength(value float32)
	HasSetFocalLength() bool

	// optional
	FocalLength() float32
	HasFocalLength() bool
}

// ensure impl type implements protocol interface
var _ PKeystoneCorrectionCombined = (*KeystoneCorrectionCombinedObject)(nil)

// A concrete type for the [PKeystoneCorrectionCombined] protocol.
type KeystoneCorrectionCombinedObject struct {
	objc.Object
}

func (k_ KeystoneCorrectionCombinedObject) HasSetFocalLength() bool {
	return k_.RespondsToSelector(objc.Sel("setFocalLength:"))
}

// The 35mm equivalent focal length of the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikeystonecorrectioncombined/3325518-focallength?language=objc
func (k_ KeystoneCorrectionCombinedObject) SetFocalLength(value float32) {
	objc.Call[objc.Void](k_, objc.Sel("setFocalLength:"), value)
}

func (k_ KeystoneCorrectionCombinedObject) HasFocalLength() bool {
	return k_.RespondsToSelector(objc.Sel("focalLength"))
}

// The 35mm equivalent focal length of the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikeystonecorrectioncombined/3325518-focallength?language=objc
func (k_ KeystoneCorrectionCombinedObject) FocalLength() float32 {
	rv := objc.Call[float32](k_, objc.Sel("focalLength"))
	return rv
}
