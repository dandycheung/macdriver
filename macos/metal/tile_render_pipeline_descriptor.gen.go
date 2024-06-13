// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TileRenderPipelineDescriptor] class.
var TileRenderPipelineDescriptorClass = _TileRenderPipelineDescriptorClass{objc.GetClass("MTLTileRenderPipelineDescriptor")}

type _TileRenderPipelineDescriptorClass struct {
	objc.Class
}

// An interface definition for the [TileRenderPipelineDescriptor] class.
type ITileRenderPipelineDescriptor interface {
	objc.IObject
	Reset()
	MaxTotalThreadsPerThreadgroup() uint
	SetMaxTotalThreadsPerThreadgroup(value uint)
	SupportAddingBinaryFunctions() bool
	SetSupportAddingBinaryFunctions(value bool)
	ThreadgroupSizeMatchesTileSize() bool
	SetThreadgroupSizeMatchesTileSize(value bool)
	TileFunction() FunctionObject
	SetTileFunction(value PFunction)
	SetTileFunctionObject(valueObject objc.IObject)
	LinkedFunctions() LinkedFunctions
	SetLinkedFunctions(value ILinkedFunctions)
	ColorAttachments() TileRenderPipelineColorAttachmentDescriptorArray
	Label() string
	SetLabel(value string)
	BinaryArchives() []BinaryArchiveObject
	SetBinaryArchives(value []PBinaryArchive)
	MaxCallStackDepth() uint
	SetMaxCallStackDepth(value uint)
	RasterSampleCount() uint
	SetRasterSampleCount(value uint)
	TileBuffers() PipelineBufferDescriptorArray
	PreloadedLibraries() []DynamicLibraryObject
	SetPreloadedLibraries(value []PDynamicLibrary)
}

// An object that configures new render pipeline state objects for tile shading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor?language=objc
type TileRenderPipelineDescriptor struct {
	objc.Object
}

func TileRenderPipelineDescriptorFrom(ptr unsafe.Pointer) TileRenderPipelineDescriptor {
	return TileRenderPipelineDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (tc _TileRenderPipelineDescriptorClass) Alloc() TileRenderPipelineDescriptor {
	rv := objc.Call[TileRenderPipelineDescriptor](tc, objc.Sel("alloc"))
	return rv
}

func (tc _TileRenderPipelineDescriptorClass) New() TileRenderPipelineDescriptor {
	rv := objc.Call[TileRenderPipelineDescriptor](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTileRenderPipelineDescriptor() TileRenderPipelineDescriptor {
	return TileRenderPipelineDescriptorClass.New()
}

func (t_ TileRenderPipelineDescriptor) Init() TileRenderPipelineDescriptor {
	rv := objc.Call[TileRenderPipelineDescriptor](t_, objc.Sel("init"))
	return rv
}

// Specifies the default rendering pipeline state values for the descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/2866346-reset?language=objc
func (t_ TileRenderPipelineDescriptor) Reset() {
	objc.Call[objc.Void](t_, objc.Sel("reset"))
}

// The maximum number of threads in a threadgroup when dispatching a command using the pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/2990998-maxtotalthreadsperthreadgroup?language=objc
func (t_ TileRenderPipelineDescriptor) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Call[uint](t_, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}

// The maximum number of threads in a threadgroup when dispatching a command using the pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/2990998-maxtotalthreadsperthreadgroup?language=objc
func (t_ TileRenderPipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Call[objc.Void](t_, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}

// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/3750588-supportaddingbinaryfunctions?language=objc
func (t_ TileRenderPipelineDescriptor) SupportAddingBinaryFunctions() bool {
	rv := objc.Call[bool](t_, objc.Sel("supportAddingBinaryFunctions"))
	return rv
}

// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/3750588-supportaddingbinaryfunctions?language=objc
func (t_ TileRenderPipelineDescriptor) SetSupportAddingBinaryFunctions(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setSupportAddingBinaryFunctions:"), value)
}

// A Boolean value that indicates whether all threadgroups for this pipeline completely cover tiles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/2866359-threadgroupsizematchestilesize?language=objc
func (t_ TileRenderPipelineDescriptor) ThreadgroupSizeMatchesTileSize() bool {
	rv := objc.Call[bool](t_, objc.Sel("threadgroupSizeMatchesTileSize"))
	return rv
}

// A Boolean value that indicates whether all threadgroups for this pipeline completely cover tiles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/2866359-threadgroupsizematchestilesize?language=objc
func (t_ TileRenderPipelineDescriptor) SetThreadgroupSizeMatchesTileSize(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setThreadgroupSizeMatchesTileSize:"), value)
}

// The compute kernel or fragment function the pipeline calls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/2866354-tilefunction?language=objc
func (t_ TileRenderPipelineDescriptor) TileFunction() FunctionObject {
	rv := objc.Call[FunctionObject](t_, objc.Sel("tileFunction"))
	return rv
}

// The compute kernel or fragment function the pipeline calls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/2866354-tilefunction?language=objc
func (t_ TileRenderPipelineDescriptor) SetTileFunction(value PFunction) {
	po0 := objc.WrapAsProtocol("MTLFunction", value)
	objc.Call[objc.Void](t_, objc.Sel("setTileFunction:"), po0)
}

// The compute kernel or fragment function the pipeline calls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/2866354-tilefunction?language=objc
func (t_ TileRenderPipelineDescriptor) SetTileFunctionObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setTileFunction:"), valueObject)
}

// Functions that you can specify as function arguments for the tile shader when encoding commands that use the pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/3750586-linkedfunctions?language=objc
func (t_ TileRenderPipelineDescriptor) LinkedFunctions() LinkedFunctions {
	rv := objc.Call[LinkedFunctions](t_, objc.Sel("linkedFunctions"))
	return rv
}

// Functions that you can specify as function arguments for the tile shader when encoding commands that use the pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/3750586-linkedfunctions?language=objc
func (t_ TileRenderPipelineDescriptor) SetLinkedFunctions(value ILinkedFunctions) {
	objc.Call[objc.Void](t_, objc.Sel("setLinkedFunctions:"), value)
}

// An array of attachments that store color data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/2866554-colorattachments?language=objc
func (t_ TileRenderPipelineDescriptor) ColorAttachments() TileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Call[TileRenderPipelineColorAttachmentDescriptorArray](t_, objc.Sel("colorAttachments"))
	return rv
}

// A string that identifies the tile pipeline descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/2866523-label?language=objc
func (t_ TileRenderPipelineDescriptor) Label() string {
	rv := objc.Call[string](t_, objc.Sel("label"))
	return rv
}

// A string that identifies the tile pipeline descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/2866523-label?language=objc
func (t_ TileRenderPipelineDescriptor) SetLabel(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setLabel:"), value)
}

// An array of binary archives to search for precompiled versions of the shader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/3564464-binaryarchives?language=objc
func (t_ TileRenderPipelineDescriptor) BinaryArchives() []BinaryArchiveObject {
	rv := objc.Call[[]BinaryArchiveObject](t_, objc.Sel("binaryArchives"))
	return rv
}

// An array of binary archives to search for precompiled versions of the shader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/3564464-binaryarchives?language=objc
func (t_ TileRenderPipelineDescriptor) SetBinaryArchives(value []PBinaryArchive) {
	objc.Call[objc.Void](t_, objc.Sel("setBinaryArchives:"), value)
}

// The maximum function call depth from the top-most shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/3750587-maxcallstackdepth?language=objc
func (t_ TileRenderPipelineDescriptor) MaxCallStackDepth() uint {
	rv := objc.Call[uint](t_, objc.Sel("maxCallStackDepth"))
	return rv
}

// The maximum function call depth from the top-most shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/3750587-maxcallstackdepth?language=objc
func (t_ TileRenderPipelineDescriptor) SetMaxCallStackDepth(value uint) {
	objc.Call[objc.Void](t_, objc.Sel("setMaxCallStackDepth:"), value)
}

// The number of samples in each fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/2890265-rastersamplecount?language=objc
func (t_ TileRenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Call[uint](t_, objc.Sel("rasterSampleCount"))
	return rv
}

// The number of samples in each fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/2890265-rastersamplecount?language=objc
func (t_ TileRenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Call[objc.Void](t_, objc.Sel("setRasterSampleCount:"), value)
}

// An array that contains the buffer mutability options for a render pipeline’s tile function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/2928217-tilebuffers?language=objc
func (t_ TileRenderPipelineDescriptor) TileBuffers() PipelineBufferDescriptorArray {
	rv := objc.Call[PipelineBufferDescriptorArray](t_, objc.Sel("tileBuffers"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/3801725-preloadedlibraries?language=objc
func (t_ TileRenderPipelineDescriptor) PreloadedLibraries() []DynamicLibraryObject {
	rv := objc.Call[[]DynamicLibraryObject](t_, objc.Sel("preloadedLibraries"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinedescriptor/3801725-preloadedlibraries?language=objc
func (t_ TileRenderPipelineDescriptor) SetPreloadedLibraries(value []PDynamicLibrary) {
	objc.Call[objc.Void](t_, objc.Sel("setPreloadedLibraries:"), value)
}
