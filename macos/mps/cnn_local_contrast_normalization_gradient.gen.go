// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNLocalContrastNormalizationGradient] class.
var CNNLocalContrastNormalizationGradientClass = _CNNLocalContrastNormalizationGradientClass{objc.GetClass("MPSCNNLocalContrastNormalizationGradient")}

type _CNNLocalContrastNormalizationGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNLocalContrastNormalizationGradient] class.
type ICNNLocalContrastNormalizationGradient interface {
	ICNNGradientKernel
	Pm() float32
	SetPm(value float32)
	P0() float32
	SetP0(value float32)
	Alpha() float32
	SetAlpha(value float32)
	Beta() float32
	SetBeta(value float32)
	Ps() float32
	SetPs(value float32)
	Delta() float32
	SetDelta(value float32)
}

// A gradient local-contrast normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient?language=objc
type CNNLocalContrastNormalizationGradient struct {
	CNNGradientKernel
}

func CNNLocalContrastNormalizationGradientFrom(ptr unsafe.Pointer) CNNLocalContrastNormalizationGradient {
	return CNNLocalContrastNormalizationGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (c_ CNNLocalContrastNormalizationGradient) InitWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNLocalContrastNormalizationGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNLocalContrastNormalizationGradient](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), po0, kernelWidth, kernelHeight)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942462-initwithdevice?language=objc
func NewCNNLocalContrastNormalizationGradientWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNLocalContrastNormalizationGradient {
	instance := CNNLocalContrastNormalizationGradientClass.Alloc().InitWithDeviceKernelWidthKernelHeight(device, kernelWidth, kernelHeight)
	instance.Autorelease()
	return instance
}

func (cc _CNNLocalContrastNormalizationGradientClass) Alloc() CNNLocalContrastNormalizationGradient {
	rv := objc.Call[CNNLocalContrastNormalizationGradient](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNLocalContrastNormalizationGradientClass) New() CNNLocalContrastNormalizationGradient {
	rv := objc.Call[CNNLocalContrastNormalizationGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNLocalContrastNormalizationGradient() CNNLocalContrastNormalizationGradient {
	return CNNLocalContrastNormalizationGradientClass.New()
}

func (c_ CNNLocalContrastNormalizationGradient) Init() CNNLocalContrastNormalizationGradient {
	rv := objc.Call[CNNLocalContrastNormalizationGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNLocalContrastNormalizationGradient) InitWithDevice(device metal.PDevice) CNNLocalContrastNormalizationGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNLocalContrastNormalizationGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNLocalContrastNormalizationGradientWithDevice(device metal.PDevice) CNNLocalContrastNormalizationGradient {
	instance := CNNLocalContrastNormalizationGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNLocalContrastNormalizationGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNLocalContrastNormalizationGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNLocalContrastNormalizationGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNLocalContrastNormalizationGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNLocalContrastNormalizationGradient {
	instance := CNNLocalContrastNormalizationGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942485-pm?language=objc
func (c_ CNNLocalContrastNormalizationGradient) Pm() float32 {
	rv := objc.Call[float32](c_, objc.Sel("pm"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942485-pm?language=objc
func (c_ CNNLocalContrastNormalizationGradient) SetPm(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setPm:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942472-p0?language=objc
func (c_ CNNLocalContrastNormalizationGradient) P0() float32 {
	rv := objc.Call[float32](c_, objc.Sel("p0"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942472-p0?language=objc
func (c_ CNNLocalContrastNormalizationGradient) SetP0(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setP0:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942471-alpha?language=objc
func (c_ CNNLocalContrastNormalizationGradient) Alpha() float32 {
	rv := objc.Call[float32](c_, objc.Sel("alpha"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942471-alpha?language=objc
func (c_ CNNLocalContrastNormalizationGradient) SetAlpha(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setAlpha:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942484-beta?language=objc
func (c_ CNNLocalContrastNormalizationGradient) Beta() float32 {
	rv := objc.Call[float32](c_, objc.Sel("beta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942484-beta?language=objc
func (c_ CNNLocalContrastNormalizationGradient) SetBeta(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setBeta:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942466-ps?language=objc
func (c_ CNNLocalContrastNormalizationGradient) Ps() float32 {
	rv := objc.Call[float32](c_, objc.Sel("ps"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942466-ps?language=objc
func (c_ CNNLocalContrastNormalizationGradient) SetPs(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setPs:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942482-delta?language=objc
func (c_ CNNLocalContrastNormalizationGradient) Delta() float32 {
	rv := objc.Call[float32](c_, objc.Sel("delta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942482-delta?language=objc
func (c_ CNNLocalContrastNormalizationGradient) SetDelta(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setDelta:"), value)
}
