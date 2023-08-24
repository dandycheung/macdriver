// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [State] class.
var StateClass = _StateClass{objc.GetClass("MPSState")}

type _StateClass struct {
	objc.Class
}

// An interface definition for the [State] class.
type IState interface {
	objc.IObject
	ResourceTypeAtIndex(index uint) StateResourceType
	DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(sourceImages []IImage, sourceStates []IState, kernel IKernel, inDescriptor IImageDescriptor) ImageDescriptor
	BufferSizeAtIndex(index uint) uint
	ResourceSize() uint
	TextureInfoAtIndex(index uint) StateTextureInfo
	ResourceAtIndexAllocateMemory(index uint, allocateMemory bool) metal.ResourceWrapper
	SynchronizeOnCommandBuffer(commandBuffer metal.PCommandBuffer)
	SynchronizeOnCommandBufferObject(commandBufferObject objc.IObject)
	ResourceCount() uint
	IsTemporary() bool
	ReadCount() uint
	SetReadCount(value uint)
	Label() string
	SetLabel(value string)
}

// An opaque data container for large storage in MPS CNN filters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate?language=objc
type State struct {
	objc.Object
}

func StateFrom(ptr unsafe.Pointer) State {
	return State{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ State) InitWithResources(resources []metal.PResource) State {
	rv := objc.Call[State](s_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewStateWithResources(resources []metal.PResource) State {
	instance := StateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (s_ State) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) State {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[State](s_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) State {
	instance := StateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (s_ State) InitWithResource(resource metal.PResource) State {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[State](s_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewStateWithResource(resource metal.PResource) State {
	instance := StateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (sc _StateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) State {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[State](sc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func State_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) State {
	return StateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

func (sc _StateClass) Alloc() State {
	rv := objc.Call[State](sc, objc.Sel("alloc"))
	return rv
}

func State_Alloc() State {
	return StateClass.Alloc()
}

func (sc _StateClass) New() State {
	rv := objc.Call[State](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewState() State {
	return StateClass.New()
}

func (s_ State) Init() State {
	rv := objc.Call[State](s_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947902-resourcetypeatindex?language=objc
func (s_ State) ResourceTypeAtIndex(index uint) StateResourceType {
	rv := objc.Call[StateResourceType](s_, objc.Sel("resourceTypeAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942394-destinationimagedescriptorforsou?language=objc
func (s_ State) DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(sourceImages []IImage, sourceStates []IState, kernel IKernel, inDescriptor IImageDescriptor) ImageDescriptor {
	rv := objc.Call[ImageDescriptor](s_, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:"), sourceImages, sourceStates, objc.Ptr(kernel), objc.Ptr(inDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947913-buffersizeatindex?language=objc
func (s_ State) BufferSizeAtIndex(index uint) uint {
	rv := objc.Call[uint](s_, objc.Sel("bufferSizeAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942397-resourcesize?language=objc
func (s_ State) ResourceSize() uint {
	rv := objc.Call[uint](s_, objc.Sel("resourceSize"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947899-textureinfoatindex?language=objc
func (s_ State) TextureInfoAtIndex(index uint) StateTextureInfo {
	rv := objc.Call[StateTextureInfo](s_, objc.Sel("textureInfoAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947916-resourceatindex?language=objc
func (s_ State) ResourceAtIndexAllocateMemory(index uint, allocateMemory bool) metal.ResourceWrapper {
	rv := objc.Call[metal.ResourceWrapper](s_, objc.Sel("resourceAtIndex:allocateMemory:"), index, allocateMemory)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942396-synchronizeoncommandbuffer?language=objc
func (s_ State) SynchronizeOnCommandBuffer(commandBuffer metal.PCommandBuffer) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](s_, objc.Sel("synchronizeOnCommandBuffer:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942396-synchronizeoncommandbuffer?language=objc
func (s_ State) SynchronizeOnCommandBufferObject(commandBufferObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("synchronizeOnCommandBuffer:"), objc.Ptr(commandBufferObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947900-resourcecount?language=objc
func (s_ State) ResourceCount() uint {
	rv := objc.Call[uint](s_, objc.Sel("resourceCount"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2867114-istemporary?language=objc
func (s_ State) IsTemporary() bool {
	rv := objc.Call[bool](s_, objc.Sel("isTemporary"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2867042-readcount?language=objc
func (s_ State) ReadCount() uint {
	rv := objc.Call[uint](s_, objc.Sel("readCount"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2867042-readcount?language=objc
func (s_ State) SetReadCount(value uint) {
	objc.Call[objc.Void](s_, objc.Sel("setReadCount:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2867179-label?language=objc
func (s_ State) Label() string {
	rv := objc.Call[string](s_, objc.Sel("label"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2867179-label?language=objc
func (s_ State) SetLabel(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setLabel:"), value)
}