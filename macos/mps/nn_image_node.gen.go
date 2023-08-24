// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNImageNode] class.
var NNImageNodeClass = _NNImageNodeClass{objc.GetClass("MPSNNImageNode")}

type _NNImageNodeClass struct {
	objc.Class
}

// An interface definition for the [NNImageNode] class.
type INNImageNode interface {
	objc.IObject
	Handle() HandleWrapper
	SetHandle(value PHandle)
	SetHandleObject(valueObject objc.IObject)
	ExportFromGraph() bool
	SetExportFromGraph(value bool)
	StopGradient() bool
	SetStopGradient(value bool)
	SynchronizeResource() bool
	SetSynchronizeResource(value bool)
	ImageAllocator() ImageAllocatorWrapper
	SetImageAllocator(value PImageAllocator)
	SetImageAllocatorObject(valueObject objc.IObject)
	Format() ImageFeatureChannelFormat
	SetFormat(value ImageFeatureChannelFormat)
}

// A placeholder node denoting the position of a neural network image in a graph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode?language=objc
type NNImageNode struct {
	objc.Object
}

func NNImageNodeFrom(ptr unsafe.Pointer) NNImageNode {
	return NNImageNode{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NNImageNodeClass) ExportedNodeWithHandle(handle objc.IObject) NNImageNode {
	rv := objc.Call[NNImageNode](nc, objc.Sel("exportedNodeWithHandle:"), objc.Ptr(handle))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866440-exportednodewithhandle?language=objc
func NNImageNode_ExportedNodeWithHandle(handle objc.IObject) NNImageNode {
	return NNImageNodeClass.ExportedNodeWithHandle(handle)
}

func (nc _NNImageNodeClass) NodeWithHandle(handle objc.IObject) NNImageNode {
	rv := objc.Call[NNImageNode](nc, objc.Sel("nodeWithHandle:"), objc.Ptr(handle))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866447-nodewithhandle?language=objc
func NNImageNode_NodeWithHandle(handle objc.IObject) NNImageNode {
	return NNImageNodeClass.NodeWithHandle(handle)
}

func (n_ NNImageNode) InitWithHandle(handle objc.IObject) NNImageNode {
	rv := objc.Call[NNImageNode](n_, objc.Sel("initWithHandle:"), objc.Ptr(handle))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866483-initwithhandle?language=objc
func NewNNImageNodeWithHandle(handle objc.IObject) NNImageNode {
	instance := NNImageNodeClass.Alloc().InitWithHandle(handle)
	instance.Autorelease()
	return instance
}

func (nc _NNImageNodeClass) Alloc() NNImageNode {
	rv := objc.Call[NNImageNode](nc, objc.Sel("alloc"))
	return rv
}

func NNImageNode_Alloc() NNImageNode {
	return NNImageNodeClass.Alloc()
}

func (nc _NNImageNodeClass) New() NNImageNode {
	rv := objc.Call[NNImageNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNImageNode() NNImageNode {
	return NNImageNodeClass.New()
}

func (n_ NNImageNode) Init() NNImageNode {
	rv := objc.Call[NNImageNode](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866406-handle?language=objc
func (n_ NNImageNode) Handle() HandleWrapper {
	rv := objc.Call[HandleWrapper](n_, objc.Sel("handle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866406-handle?language=objc
func (n_ NNImageNode) SetHandle(value PHandle) {
	po0 := objc.WrapAsProtocol("MPSHandle", value)
	objc.Call[objc.Void](n_, objc.Sel("setHandle:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866406-handle?language=objc
func (n_ NNImageNode) SetHandleObject(valueObject objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setHandle:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866478-exportfromgraph?language=objc
func (n_ NNImageNode) ExportFromGraph() bool {
	rv := objc.Call[bool](n_, objc.Sel("exportFromGraph"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866478-exportfromgraph?language=objc
func (n_ NNImageNode) SetExportFromGraph(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setExportFromGraph:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/3020689-stopgradient?language=objc
func (n_ NNImageNode) StopGradient() bool {
	rv := objc.Call[bool](n_, objc.Sel("stopGradient"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/3020689-stopgradient?language=objc
func (n_ NNImageNode) SetStopGradient(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setStopGradient:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2942638-synchronizeresource?language=objc
func (n_ NNImageNode) SynchronizeResource() bool {
	rv := objc.Call[bool](n_, objc.Sel("synchronizeResource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2942638-synchronizeresource?language=objc
func (n_ NNImageNode) SetSynchronizeResource(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setSynchronizeResource:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866490-imageallocator?language=objc
func (n_ NNImageNode) ImageAllocator() ImageAllocatorWrapper {
	rv := objc.Call[ImageAllocatorWrapper](n_, objc.Sel("imageAllocator"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866490-imageallocator?language=objc
func (n_ NNImageNode) SetImageAllocator(value PImageAllocator) {
	po0 := objc.WrapAsProtocol("MPSImageAllocator", value)
	objc.Call[objc.Void](n_, objc.Sel("setImageAllocator:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866490-imageallocator?language=objc
func (n_ NNImageNode) SetImageAllocatorObject(valueObject objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setImageAllocator:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866498-format?language=objc
func (n_ NNImageNode) Format() ImageFeatureChannelFormat {
	rv := objc.Call[ImageFeatureChannelFormat](n_, objc.Sel("format"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866498-format?language=objc
func (n_ NNImageNode) SetFormat(value ImageFeatureChannelFormat) {
	objc.Call[objc.Void](n_, objc.Sel("setFormat:"), value)
}