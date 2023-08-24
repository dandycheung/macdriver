// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageCanny] class.
var ImageCannyClass = _ImageCannyClass{objc.GetClass("MPSImageCanny")}

type _ImageCannyClass struct {
	objc.Class
}

// An interface definition for the [ImageCanny] class.
type IImageCanny interface {
	IUnaryImageKernel
	LowThreshold() float64
	SetLowThreshold(value float64)
	ColorTransform() *float64
	UseFastMode() bool
	SetUseFastMode(value bool)
	Sigma() float64
	HighThreshold() float64
	SetHighThreshold(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny?language=objc
type ImageCanny struct {
	UnaryImageKernel
}

func ImageCannyFrom(ptr unsafe.Pointer) ImageCanny {
	return ImageCanny{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageCanny) InitWithDevice(device metal.PDevice) ImageCanny {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageCanny](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547972-initwithdevice?language=objc
func NewImageCannyWithDevice(device metal.PDevice) ImageCanny {
	instance := ImageCannyClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageCannyClass) Alloc() ImageCanny {
	rv := objc.Call[ImageCanny](ic, objc.Sel("alloc"))
	return rv
}

func ImageCanny_Alloc() ImageCanny {
	return ImageCannyClass.Alloc()
}

func (ic _ImageCannyClass) New() ImageCanny {
	rv := objc.Call[ImageCanny](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageCanny() ImageCanny {
	return ImageCannyClass.New()
}

func (i_ ImageCanny) Init() ImageCanny {
	rv := objc.Call[ImageCanny](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageCanny) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageCanny {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageCanny](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageCanny_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageCanny {
	instance := ImageCannyClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547974-lowthreshold?language=objc
func (i_ ImageCanny) LowThreshold() float64 {
	rv := objc.Call[float64](i_, objc.Sel("lowThreshold"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547974-lowthreshold?language=objc
func (i_ ImageCanny) SetLowThreshold(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setLowThreshold:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547969-colortransform?language=objc
func (i_ ImageCanny) ColorTransform() *float64 {
	rv := objc.Call[*float64](i_, objc.Sel("colorTransform"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547976-usefastmode?language=objc
func (i_ ImageCanny) UseFastMode() bool {
	rv := objc.Call[bool](i_, objc.Sel("useFastMode"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547976-usefastmode?language=objc
func (i_ ImageCanny) SetUseFastMode(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setUseFastMode:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547975-sigma?language=objc
func (i_ ImageCanny) Sigma() float64 {
	rv := objc.Call[float64](i_, objc.Sel("sigma"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547970-highthreshold?language=objc
func (i_ ImageCanny) HighThreshold() float64 {
	rv := objc.Call[float64](i_, objc.Sel("highThreshold"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547970-highthreshold?language=objc
func (i_ ImageCanny) SetHighThreshold(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setHighThreshold:"), value)
}