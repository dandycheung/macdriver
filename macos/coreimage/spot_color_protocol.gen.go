// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a spot color filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor?language=objc
type PSpotColor interface {
	// optional
	SetReplacementColor3(value Color)
	HasSetReplacementColor3() bool

	// optional
	ReplacementColor3() Color
	HasReplacementColor3() bool

	// optional
	SetContrast2(value float32)
	HasSetContrast2() bool

	// optional
	Contrast2() float32
	HasContrast2() bool

	// optional
	SetCenterColor1(value Color)
	HasSetCenterColor1() bool

	// optional
	CenterColor1() Color
	HasCenterColor1() bool

	// optional
	SetCloseness1(value float32)
	HasSetCloseness1() bool

	// optional
	Closeness1() float32
	HasCloseness1() bool

	// optional
	SetCloseness3(value float32)
	HasSetCloseness3() bool

	// optional
	Closeness3() float32
	HasCloseness3() bool

	// optional
	SetReplacementColor2(value Color)
	HasSetReplacementColor2() bool

	// optional
	ReplacementColor2() Color
	HasReplacementColor2() bool

	// optional
	SetContrast3(value float32)
	HasSetContrast3() bool

	// optional
	Contrast3() float32
	HasContrast3() bool

	// optional
	SetCloseness2(value float32)
	HasSetCloseness2() bool

	// optional
	Closeness2() float32
	HasCloseness2() bool

	// optional
	SetCenterColor2(value Color)
	HasSetCenterColor2() bool

	// optional
	CenterColor2() Color
	HasCenterColor2() bool

	// optional
	SetReplacementColor1(value Color)
	HasSetReplacementColor1() bool

	// optional
	ReplacementColor1() Color
	HasReplacementColor1() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetContrast1(value float32)
	HasSetContrast1() bool

	// optional
	Contrast1() float32
	HasContrast1() bool

	// optional
	SetCenterColor3(value Color)
	HasSetCenterColor3() bool

	// optional
	CenterColor3() Color
	HasCenterColor3() bool
}

// ensure impl type implements protocol interface
var _ PSpotColor = (*SpotColorObject)(nil)

// A concrete type for the [PSpotColor] protocol.
type SpotColorObject struct {
	objc.Object
}

func (s_ SpotColorObject) HasSetReplacementColor3() bool {
	return s_.RespondsToSelector(objc.Sel("setReplacementColor3:"))
}

// A replacement color for the third color range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228740-replacementcolor3?language=objc
func (s_ SpotColorObject) SetReplacementColor3(value Color) {
	objc.Call[objc.Void](s_, objc.Sel("setReplacementColor3:"), value)
}

func (s_ SpotColorObject) HasReplacementColor3() bool {
	return s_.RespondsToSelector(objc.Sel("replacementColor3"))
}

// A replacement color for the third color range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228740-replacementcolor3?language=objc
func (s_ SpotColorObject) ReplacementColor3() Color {
	rv := objc.Call[Color](s_, objc.Sel("replacementColor3"))
	return rv
}

func (s_ SpotColorObject) HasSetContrast2() bool {
	return s_.RespondsToSelector(objc.Sel("setContrast2:"))
}

// The contrast of the second replacement color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228735-contrast2?language=objc
func (s_ SpotColorObject) SetContrast2(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setContrast2:"), value)
}

func (s_ SpotColorObject) HasContrast2() bool {
	return s_.RespondsToSelector(objc.Sel("contrast2"))
}

// The contrast of the second replacement color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228735-contrast2?language=objc
func (s_ SpotColorObject) Contrast2() float32 {
	rv := objc.Call[float32](s_, objc.Sel("contrast2"))
	return rv
}

func (s_ SpotColorObject) HasSetCenterColor1() bool {
	return s_.RespondsToSelector(objc.Sel("setCenterColor1:"))
}

// The center value of the first color range to replace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228728-centercolor1?language=objc
func (s_ SpotColorObject) SetCenterColor1(value Color) {
	objc.Call[objc.Void](s_, objc.Sel("setCenterColor1:"), value)
}

func (s_ SpotColorObject) HasCenterColor1() bool {
	return s_.RespondsToSelector(objc.Sel("centerColor1"))
}

// The center value of the first color range to replace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228728-centercolor1?language=objc
func (s_ SpotColorObject) CenterColor1() Color {
	rv := objc.Call[Color](s_, objc.Sel("centerColor1"))
	return rv
}

func (s_ SpotColorObject) HasSetCloseness1() bool {
	return s_.RespondsToSelector(objc.Sel("setCloseness1:"))
}

// A value that indicates how closely the first color must match before it’s replaced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228731-closeness1?language=objc
func (s_ SpotColorObject) SetCloseness1(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setCloseness1:"), value)
}

func (s_ SpotColorObject) HasCloseness1() bool {
	return s_.RespondsToSelector(objc.Sel("closeness1"))
}

// A value that indicates how closely the first color must match before it’s replaced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228731-closeness1?language=objc
func (s_ SpotColorObject) Closeness1() float32 {
	rv := objc.Call[float32](s_, objc.Sel("closeness1"))
	return rv
}

func (s_ SpotColorObject) HasSetCloseness3() bool {
	return s_.RespondsToSelector(objc.Sel("setCloseness3:"))
}

// A value that indicates how closely the third color must match before it’s replaced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228733-closeness3?language=objc
func (s_ SpotColorObject) SetCloseness3(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setCloseness3:"), value)
}

func (s_ SpotColorObject) HasCloseness3() bool {
	return s_.RespondsToSelector(objc.Sel("closeness3"))
}

// A value that indicates how closely the third color must match before it’s replaced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228733-closeness3?language=objc
func (s_ SpotColorObject) Closeness3() float32 {
	rv := objc.Call[float32](s_, objc.Sel("closeness3"))
	return rv
}

func (s_ SpotColorObject) HasSetReplacementColor2() bool {
	return s_.RespondsToSelector(objc.Sel("setReplacementColor2:"))
}

// A replacement color for the second color range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228739-replacementcolor2?language=objc
func (s_ SpotColorObject) SetReplacementColor2(value Color) {
	objc.Call[objc.Void](s_, objc.Sel("setReplacementColor2:"), value)
}

func (s_ SpotColorObject) HasReplacementColor2() bool {
	return s_.RespondsToSelector(objc.Sel("replacementColor2"))
}

// A replacement color for the second color range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228739-replacementcolor2?language=objc
func (s_ SpotColorObject) ReplacementColor2() Color {
	rv := objc.Call[Color](s_, objc.Sel("replacementColor2"))
	return rv
}

func (s_ SpotColorObject) HasSetContrast3() bool {
	return s_.RespondsToSelector(objc.Sel("setContrast3:"))
}

// The contrast of the third replacement color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228736-contrast3?language=objc
func (s_ SpotColorObject) SetContrast3(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setContrast3:"), value)
}

func (s_ SpotColorObject) HasContrast3() bool {
	return s_.RespondsToSelector(objc.Sel("contrast3"))
}

// The contrast of the third replacement color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228736-contrast3?language=objc
func (s_ SpotColorObject) Contrast3() float32 {
	rv := objc.Call[float32](s_, objc.Sel("contrast3"))
	return rv
}

func (s_ SpotColorObject) HasSetCloseness2() bool {
	return s_.RespondsToSelector(objc.Sel("setCloseness2:"))
}

// A value that indicates how closely the second color must match before it’s replaced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228732-closeness2?language=objc
func (s_ SpotColorObject) SetCloseness2(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setCloseness2:"), value)
}

func (s_ SpotColorObject) HasCloseness2() bool {
	return s_.RespondsToSelector(objc.Sel("closeness2"))
}

// A value that indicates how closely the second color must match before it’s replaced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228732-closeness2?language=objc
func (s_ SpotColorObject) Closeness2() float32 {
	rv := objc.Call[float32](s_, objc.Sel("closeness2"))
	return rv
}

func (s_ SpotColorObject) HasSetCenterColor2() bool {
	return s_.RespondsToSelector(objc.Sel("setCenterColor2:"))
}

// The center value of the second color range to replace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228729-centercolor2?language=objc
func (s_ SpotColorObject) SetCenterColor2(value Color) {
	objc.Call[objc.Void](s_, objc.Sel("setCenterColor2:"), value)
}

func (s_ SpotColorObject) HasCenterColor2() bool {
	return s_.RespondsToSelector(objc.Sel("centerColor2"))
}

// The center value of the second color range to replace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228729-centercolor2?language=objc
func (s_ SpotColorObject) CenterColor2() Color {
	rv := objc.Call[Color](s_, objc.Sel("centerColor2"))
	return rv
}

func (s_ SpotColorObject) HasSetReplacementColor1() bool {
	return s_.RespondsToSelector(objc.Sel("setReplacementColor1:"))
}

// A replacement color for the first color range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228738-replacementcolor1?language=objc
func (s_ SpotColorObject) SetReplacementColor1(value Color) {
	objc.Call[objc.Void](s_, objc.Sel("setReplacementColor1:"), value)
}

func (s_ SpotColorObject) HasReplacementColor1() bool {
	return s_.RespondsToSelector(objc.Sel("replacementColor1"))
}

// A replacement color for the first color range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228738-replacementcolor1?language=objc
func (s_ SpotColorObject) ReplacementColor1() Color {
	rv := objc.Call[Color](s_, objc.Sel("replacementColor1"))
	return rv
}

func (s_ SpotColorObject) HasSetInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228737-inputimage?language=objc
func (s_ SpotColorObject) SetInputImage(value Image) {
	objc.Call[objc.Void](s_, objc.Sel("setInputImage:"), value)
}

func (s_ SpotColorObject) HasInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228737-inputimage?language=objc
func (s_ SpotColorObject) InputImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("inputImage"))
	return rv
}

func (s_ SpotColorObject) HasSetContrast1() bool {
	return s_.RespondsToSelector(objc.Sel("setContrast1:"))
}

// The contrast of the first replacement color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228734-contrast1?language=objc
func (s_ SpotColorObject) SetContrast1(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setContrast1:"), value)
}

func (s_ SpotColorObject) HasContrast1() bool {
	return s_.RespondsToSelector(objc.Sel("contrast1"))
}

// The contrast of the first replacement color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228734-contrast1?language=objc
func (s_ SpotColorObject) Contrast1() float32 {
	rv := objc.Call[float32](s_, objc.Sel("contrast1"))
	return rv
}

func (s_ SpotColorObject) HasSetCenterColor3() bool {
	return s_.RespondsToSelector(objc.Sel("setCenterColor3:"))
}

// The center value of the third color range to replace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228730-centercolor3?language=objc
func (s_ SpotColorObject) SetCenterColor3(value Color) {
	objc.Call[objc.Void](s_, objc.Sel("setCenterColor3:"), value)
}

func (s_ SpotColorObject) HasCenterColor3() bool {
	return s_.RespondsToSelector(objc.Sel("centerColor3"))
}

// The center value of the third color range to replace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228730-centercolor3?language=objc
func (s_ SpotColorObject) CenterColor3() Color {
	rv := objc.Call[Color](s_, objc.Sel("centerColor3"))
	return rv
}
