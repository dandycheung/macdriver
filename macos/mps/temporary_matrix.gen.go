// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TemporaryMatrix] class.
var TemporaryMatrixClass = _TemporaryMatrixClass{objc.GetClass("MPSTemporaryMatrix")}

type _TemporaryMatrixClass struct {
	objc.Class
}

// An interface definition for the [TemporaryMatrix] class.
type ITemporaryMatrix interface {
	IMatrix
	ReadCount() uint
	SetReadCount(value uint)
}

// A matrix allocated on GPU private memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix?language=objc
type TemporaryMatrix struct {
	Matrix
}

func TemporaryMatrixFrom(ptr unsafe.Pointer) TemporaryMatrix {
	return TemporaryMatrix{
		Matrix: MatrixFrom(ptr),
	}
}

func (tc _TemporaryMatrixClass) TemporaryMatrixWithCommandBufferMatrixDescriptor(commandBuffer metal.PCommandBuffer, matrixDescriptor IMatrixDescriptor) TemporaryMatrix {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[TemporaryMatrix](tc, objc.Sel("temporaryMatrixWithCommandBuffer:matrixDescriptor:"), po0, matrixDescriptor)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix/2867180-temporarymatrixwithcommandbuffer?language=objc
func TemporaryMatrix_TemporaryMatrixWithCommandBufferMatrixDescriptor(commandBuffer metal.PCommandBuffer, matrixDescriptor IMatrixDescriptor) TemporaryMatrix {
	return TemporaryMatrixClass.TemporaryMatrixWithCommandBufferMatrixDescriptor(commandBuffer, matrixDescriptor)
}

func (tc _TemporaryMatrixClass) Alloc() TemporaryMatrix {
	rv := objc.Call[TemporaryMatrix](tc, objc.Sel("alloc"))
	return rv
}

func (tc _TemporaryMatrixClass) New() TemporaryMatrix {
	rv := objc.Call[TemporaryMatrix](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTemporaryMatrix() TemporaryMatrix {
	return TemporaryMatrixClass.New()
}

func (t_ TemporaryMatrix) Init() TemporaryMatrix {
	rv := objc.Call[TemporaryMatrix](t_, objc.Sel("init"))
	return rv
}

func (t_ TemporaryMatrix) InitWithBufferOffsetDescriptor(buffer metal.PBuffer, offset uint, descriptor IMatrixDescriptor) TemporaryMatrix {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	rv := objc.Call[TemporaryMatrix](t_, objc.Sel("initWithBuffer:offset:descriptor:"), po0, offset, descriptor)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/3229863-initwithbuffer?language=objc
func NewTemporaryMatrixWithBufferOffsetDescriptor(buffer metal.PBuffer, offset uint, descriptor IMatrixDescriptor) TemporaryMatrix {
	instance := TemporaryMatrixClass.Alloc().InitWithBufferOffsetDescriptor(buffer, offset, descriptor)
	instance.Autorelease()
	return instance
}

func (t_ TemporaryMatrix) InitWithBufferDescriptor(buffer metal.PBuffer, descriptor IMatrixDescriptor) TemporaryMatrix {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	rv := objc.Call[TemporaryMatrix](t_, objc.Sel("initWithBuffer:descriptor:"), po0, descriptor)
	return rv
}

// Initializes a matrix with a buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143201-initwithbuffer?language=objc
func NewTemporaryMatrixWithBufferDescriptor(buffer metal.PBuffer, descriptor IMatrixDescriptor) TemporaryMatrix {
	instance := TemporaryMatrixClass.Alloc().InitWithBufferDescriptor(buffer, descriptor)
	instance.Autorelease()
	return instance
}

func (t_ TemporaryMatrix) InitWithDeviceDescriptor(device metal.PDevice, descriptor IMatrixDescriptor) TemporaryMatrix {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[TemporaryMatrix](t_, objc.Sel("initWithDevice:descriptor:"), po0, descriptor)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2942567-initwithdevice?language=objc
func NewTemporaryMatrixWithDeviceDescriptor(device metal.PDevice, descriptor IMatrixDescriptor) TemporaryMatrix {
	instance := TemporaryMatrixClass.Alloc().InitWithDeviceDescriptor(device, descriptor)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix/2867073-prefetchstoragewithcommandbuffer?language=objc
func (tc _TemporaryMatrixClass) PrefetchStorageWithCommandBufferMatrixDescriptorList(commandBuffer metal.PCommandBuffer, descriptorList []IMatrixDescriptor) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](tc, objc.Sel("prefetchStorageWithCommandBuffer:matrixDescriptorList:"), po0, descriptorList)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix/2867073-prefetchstoragewithcommandbuffer?language=objc
func TemporaryMatrix_PrefetchStorageWithCommandBufferMatrixDescriptorList(commandBuffer metal.PCommandBuffer, descriptorList []IMatrixDescriptor) {
	TemporaryMatrixClass.PrefetchStorageWithCommandBufferMatrixDescriptorList(commandBuffer, descriptorList)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix/2867151-readcount?language=objc
func (t_ TemporaryMatrix) ReadCount() uint {
	rv := objc.Call[uint](t_, objc.Sel("readCount"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix/2867151-readcount?language=objc
func (t_ TemporaryMatrix) SetReadCount(value uint) {
	objc.Call[objc.Void](t_, objc.Sel("setReadCount:"), value)
}
