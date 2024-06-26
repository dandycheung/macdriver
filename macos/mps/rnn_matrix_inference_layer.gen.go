// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [RNNMatrixInferenceLayer] class.
var RNNMatrixInferenceLayerClass = _RNNMatrixInferenceLayerClass{objc.GetClass("MPSRNNMatrixInferenceLayer")}

type _RNNMatrixInferenceLayerClass struct {
	objc.Class
}

// An interface definition for the [RNNMatrixInferenceLayer] class.
type IRNNMatrixInferenceLayer interface {
	IKernel
	EncodeSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsRecurrentInputStateRecurrentOutputStates(commandBuffer metal.PCommandBuffer, sourceMatrices []IMatrix, sourceOffsets *uint, destinationMatrices []IMatrix, destinationOffsets *uint, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates foundation.IMutableArray)
	EncodeSequenceToCommandBufferObjectSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsRecurrentInputStateRecurrentOutputStates(commandBufferObject objc.IObject, sourceMatrices []IMatrix, sourceOffsets *uint, destinationMatrices []IMatrix, destinationOffsets *uint, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates foundation.IMutableArray)
	EncodeSequenceToCommandBufferSourceMatricesDestinationMatricesRecurrentInputStateRecurrentOutputStates(commandBuffer metal.PCommandBuffer, sourceMatrices []IMatrix, destinationMatrices []IMatrix, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates foundation.IMutableArray)
	EncodeSequenceToCommandBufferObjectSourceMatricesDestinationMatricesRecurrentInputStateRecurrentOutputStates(commandBufferObject objc.IObject, sourceMatrices []IMatrix, destinationMatrices []IMatrix, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates foundation.IMutableArray)
	EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardMatricesDestinationBackwardMatrices(commandBuffer metal.PCommandBuffer, sourceSequence []IMatrix, destinationForwardMatrices []IMatrix, destinationBackwardMatrices []IMatrix)
	EncodeBidirectionalSequenceToCommandBufferObjectSourceSequenceDestinationForwardMatricesDestinationBackwardMatrices(commandBufferObject objc.IObject, sourceSequence []IMatrix, destinationForwardMatrices []IMatrix, destinationBackwardMatrices []IMatrix)
	BidirectionalCombineMode() RNNBidirectionalCombineMode
	SetBidirectionalCombineMode(value RNNBidirectionalCombineMode)
	RecurrentOutputIsTemporary() bool
	SetRecurrentOutputIsTemporary(value bool)
	NumberOfLayers() uint
	OutputFeatureChannels() uint
	InputFeatureChannels() uint
	StoreAllIntermediateStates() bool
	SetStoreAllIntermediateStates(value bool)
}

// A recurrent neural network layer for inference on Metal Performance Shaders matrices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer?language=objc
type RNNMatrixInferenceLayer struct {
	Kernel
}

func RNNMatrixInferenceLayerFrom(ptr unsafe.Pointer) RNNMatrixInferenceLayer {
	return RNNMatrixInferenceLayer{
		Kernel: KernelFrom(ptr),
	}
}

func (r_ RNNMatrixInferenceLayer) InitWithDeviceRnnDescriptor(device metal.PDevice, rnnDescriptor IRNNDescriptor) RNNMatrixInferenceLayer {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[RNNMatrixInferenceLayer](r_, objc.Sel("initWithDevice:rnnDescriptor:"), po0, rnnDescriptor)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865704-initwithdevice?language=objc
func NewRNNMatrixInferenceLayerWithDeviceRnnDescriptor(device metal.PDevice, rnnDescriptor IRNNDescriptor) RNNMatrixInferenceLayer {
	instance := RNNMatrixInferenceLayerClass.Alloc().InitWithDeviceRnnDescriptor(device, rnnDescriptor)
	instance.Autorelease()
	return instance
}

func (r_ RNNMatrixInferenceLayer) InitWithDeviceRnnDescriptors(device metal.PDevice, rnnDescriptors []IRNNDescriptor) RNNMatrixInferenceLayer {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[RNNMatrixInferenceLayer](r_, objc.Sel("initWithDevice:rnnDescriptors:"), po0, rnnDescriptors)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865751-initwithdevice?language=objc
func NewRNNMatrixInferenceLayerWithDeviceRnnDescriptors(device metal.PDevice, rnnDescriptors []IRNNDescriptor) RNNMatrixInferenceLayer {
	instance := RNNMatrixInferenceLayerClass.Alloc().InitWithDeviceRnnDescriptors(device, rnnDescriptors)
	instance.Autorelease()
	return instance
}

func (r_ RNNMatrixInferenceLayer) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) RNNMatrixInferenceLayer {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[RNNMatrixInferenceLayer](r_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865746-copywithzone?language=objc
func RNNMatrixInferenceLayer_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) RNNMatrixInferenceLayer {
	instance := RNNMatrixInferenceLayerClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (rc _RNNMatrixInferenceLayerClass) Alloc() RNNMatrixInferenceLayer {
	rv := objc.Call[RNNMatrixInferenceLayer](rc, objc.Sel("alloc"))
	return rv
}

func (rc _RNNMatrixInferenceLayerClass) New() RNNMatrixInferenceLayer {
	rv := objc.Call[RNNMatrixInferenceLayer](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRNNMatrixInferenceLayer() RNNMatrixInferenceLayer {
	return RNNMatrixInferenceLayerClass.New()
}

func (r_ RNNMatrixInferenceLayer) Init() RNNMatrixInferenceLayer {
	rv := objc.Call[RNNMatrixInferenceLayer](r_, objc.Sel("init"))
	return rv
}

func (r_ RNNMatrixInferenceLayer) InitWithDevice(device metal.PDevice) RNNMatrixInferenceLayer {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[RNNMatrixInferenceLayer](r_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewRNNMatrixInferenceLayerWithDevice(device metal.PDevice) RNNMatrixInferenceLayer {
	instance := RNNMatrixInferenceLayerClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2966781-encodesequencetocommandbuffer?language=objc
func (r_ RNNMatrixInferenceLayer) EncodeSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsRecurrentInputStateRecurrentOutputStates(commandBuffer metal.PCommandBuffer, sourceMatrices []IMatrix, sourceOffsets *uint, destinationMatrices []IMatrix, destinationOffsets *uint, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates foundation.IMutableArray) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](r_, objc.Sel("encodeSequenceToCommandBuffer:sourceMatrices:sourceOffsets:destinationMatrices:destinationOffsets:recurrentInputState:recurrentOutputStates:"), po0, sourceMatrices, sourceOffsets, destinationMatrices, destinationOffsets, recurrentInputState, recurrentOutputStates)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2966781-encodesequencetocommandbuffer?language=objc
func (r_ RNNMatrixInferenceLayer) EncodeSequenceToCommandBufferObjectSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsRecurrentInputStateRecurrentOutputStates(commandBufferObject objc.IObject, sourceMatrices []IMatrix, sourceOffsets *uint, destinationMatrices []IMatrix, destinationOffsets *uint, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates foundation.IMutableArray) {
	objc.Call[objc.Void](r_, objc.Sel("encodeSequenceToCommandBuffer:sourceMatrices:sourceOffsets:destinationMatrices:destinationOffsets:recurrentInputState:recurrentOutputStates:"), commandBufferObject, sourceMatrices, sourceOffsets, destinationMatrices, destinationOffsets, recurrentInputState, recurrentOutputStates)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865705-encodesequencetocommandbuffer?language=objc
func (r_ RNNMatrixInferenceLayer) EncodeSequenceToCommandBufferSourceMatricesDestinationMatricesRecurrentInputStateRecurrentOutputStates(commandBuffer metal.PCommandBuffer, sourceMatrices []IMatrix, destinationMatrices []IMatrix, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates foundation.IMutableArray) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](r_, objc.Sel("encodeSequenceToCommandBuffer:sourceMatrices:destinationMatrices:recurrentInputState:recurrentOutputStates:"), po0, sourceMatrices, destinationMatrices, recurrentInputState, recurrentOutputStates)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865705-encodesequencetocommandbuffer?language=objc
func (r_ RNNMatrixInferenceLayer) EncodeSequenceToCommandBufferObjectSourceMatricesDestinationMatricesRecurrentInputStateRecurrentOutputStates(commandBufferObject objc.IObject, sourceMatrices []IMatrix, destinationMatrices []IMatrix, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates foundation.IMutableArray) {
	objc.Call[objc.Void](r_, objc.Sel("encodeSequenceToCommandBuffer:sourceMatrices:destinationMatrices:recurrentInputState:recurrentOutputStates:"), commandBufferObject, sourceMatrices, destinationMatrices, recurrentInputState, recurrentOutputStates)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865698-encodebidirectionalsequencetocom?language=objc
func (r_ RNNMatrixInferenceLayer) EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardMatricesDestinationBackwardMatrices(commandBuffer metal.PCommandBuffer, sourceSequence []IMatrix, destinationForwardMatrices []IMatrix, destinationBackwardMatrices []IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](r_, objc.Sel("encodeBidirectionalSequenceToCommandBuffer:sourceSequence:destinationForwardMatrices:destinationBackwardMatrices:"), po0, sourceSequence, destinationForwardMatrices, destinationBackwardMatrices)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865698-encodebidirectionalsequencetocom?language=objc
func (r_ RNNMatrixInferenceLayer) EncodeBidirectionalSequenceToCommandBufferObjectSourceSequenceDestinationForwardMatricesDestinationBackwardMatrices(commandBufferObject objc.IObject, sourceSequence []IMatrix, destinationForwardMatrices []IMatrix, destinationBackwardMatrices []IMatrix) {
	objc.Call[objc.Void](r_, objc.Sel("encodeBidirectionalSequenceToCommandBuffer:sourceSequence:destinationForwardMatrices:destinationBackwardMatrices:"), commandBufferObject, sourceSequence, destinationForwardMatrices, destinationBackwardMatrices)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865739-bidirectionalcombinemode?language=objc
func (r_ RNNMatrixInferenceLayer) BidirectionalCombineMode() RNNBidirectionalCombineMode {
	rv := objc.Call[RNNBidirectionalCombineMode](r_, objc.Sel("bidirectionalCombineMode"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865739-bidirectionalcombinemode?language=objc
func (r_ RNNMatrixInferenceLayer) SetBidirectionalCombineMode(value RNNBidirectionalCombineMode) {
	objc.Call[objc.Void](r_, objc.Sel("setBidirectionalCombineMode:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865714-recurrentoutputistemporary?language=objc
func (r_ RNNMatrixInferenceLayer) RecurrentOutputIsTemporary() bool {
	rv := objc.Call[bool](r_, objc.Sel("recurrentOutputIsTemporary"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865714-recurrentoutputistemporary?language=objc
func (r_ RNNMatrixInferenceLayer) SetRecurrentOutputIsTemporary(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setRecurrentOutputIsTemporary:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2873347-numberoflayers?language=objc
func (r_ RNNMatrixInferenceLayer) NumberOfLayers() uint {
	rv := objc.Call[uint](r_, objc.Sel("numberOfLayers"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2890142-outputfeaturechannels?language=objc
func (r_ RNNMatrixInferenceLayer) OutputFeatureChannels() uint {
	rv := objc.Call[uint](r_, objc.Sel("outputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2890143-inputfeaturechannels?language=objc
func (r_ RNNMatrixInferenceLayer) InputFeatureChannels() uint {
	rv := objc.Call[uint](r_, objc.Sel("inputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865729-storeallintermediatestates?language=objc
func (r_ RNNMatrixInferenceLayer) StoreAllIntermediateStates() bool {
	rv := objc.Call[bool](r_, objc.Sel("storeAllIntermediateStates"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865729-storeallintermediatestates?language=objc
func (r_ RNNMatrixInferenceLayer) SetStoreAllIntermediateStates(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setStoreAllIntermediateStates:"), value)
}
