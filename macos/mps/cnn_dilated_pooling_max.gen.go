// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CNNDilatedPoolingMax] class.
var CNNDilatedPoolingMaxClass = _CNNDilatedPoolingMaxClass{objc.GetClass("MPSCNNDilatedPoolingMax")}

type _CNNDilatedPoolingMaxClass struct {
	objc.Class
}

// An interface definition for the [CNNDilatedPoolingMax] class.
type ICNNDilatedPoolingMax interface {
	ICNNPooling
}

// A dilated max pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmax?language=objc
type CNNDilatedPoolingMax struct {
	CNNPooling
}

func CNNDilatedPoolingMaxFrom(ptr unsafe.Pointer) CNNDilatedPoolingMax {
	return CNNDilatedPoolingMax{
		CNNPooling: CNNPoolingFrom(ptr),
	}
}

func (c_ CNNDilatedPoolingMax) InitWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, dilationRateX uint, dilationRateY uint, strideInPixelsX uint, strideInPixelsY uint) CNNDilatedPoolingMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDilatedPoolingMax](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:dilationRateX:dilationRateY:strideInPixelsX:strideInPixelsY:"), po0, kernelWidth, kernelHeight, dilationRateX, dilationRateY, strideInPixelsX, strideInPixelsY)
	return rv
}

// Initializes a dilated max pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmax/2881192-initwithdevice?language=objc
func NewCNNDilatedPoolingMaxWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, dilationRateX uint, dilationRateY uint, strideInPixelsX uint, strideInPixelsY uint) CNNDilatedPoolingMax {
	instance := CNNDilatedPoolingMaxClass.Alloc().InitWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY(device, kernelWidth, kernelHeight, dilationRateX, dilationRateY, strideInPixelsX, strideInPixelsY)
	instance.Autorelease()
	return instance
}

func (cc _CNNDilatedPoolingMaxClass) Alloc() CNNDilatedPoolingMax {
	rv := objc.Call[CNNDilatedPoolingMax](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNDilatedPoolingMaxClass) New() CNNDilatedPoolingMax {
	rv := objc.Call[CNNDilatedPoolingMax](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNDilatedPoolingMax() CNNDilatedPoolingMax {
	return CNNDilatedPoolingMaxClass.New()
}

func (c_ CNNDilatedPoolingMax) Init() CNNDilatedPoolingMax {
	rv := objc.Call[CNNDilatedPoolingMax](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNDilatedPoolingMax) InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNDilatedPoolingMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDilatedPoolingMax](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), po0, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return rv
}

// Initializes a pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpooling/1648902-initwithdevice?language=objc
func NewCNNDilatedPoolingMaxWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNDilatedPoolingMax {
	instance := CNNDilatedPoolingMaxClass.Alloc().InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	instance.Autorelease()
	return instance
}

func (c_ CNNDilatedPoolingMax) InitWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNDilatedPoolingMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDilatedPoolingMax](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), po0, kernelWidth, kernelHeight)
	return rv
}

// Initializes a pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpooling/1648887-initwithdevice?language=objc
func NewCNNDilatedPoolingMaxWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNDilatedPoolingMax {
	instance := CNNDilatedPoolingMaxClass.Alloc().InitWithDeviceKernelWidthKernelHeight(device, kernelWidth, kernelHeight)
	instance.Autorelease()
	return instance
}

func (c_ CNNDilatedPoolingMax) InitWithDevice(device metal.PDevice) CNNDilatedPoolingMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDilatedPoolingMax](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNDilatedPoolingMaxWithDevice(device metal.PDevice) CNNDilatedPoolingMax {
	instance := CNNDilatedPoolingMaxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNDilatedPoolingMax) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNDilatedPoolingMax {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDilatedPoolingMax](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNDilatedPoolingMax_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNDilatedPoolingMax {
	instance := CNNDilatedPoolingMaxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
