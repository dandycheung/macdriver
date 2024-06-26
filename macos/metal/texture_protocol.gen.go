// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/iosurface"
	"github.com/progrium/darwinkit/objc"
)

// A resource that holds formatted image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture?language=objc
type PTexture interface {
	// optional
	GetBytesBytesPerRowBytesPerImageFromRegionMipmapLevelSlice(pixelBytes unsafe.Pointer, bytesPerRow uint, bytesPerImage uint, region Region, level uint, slice uint)
	HasGetBytesBytesPerRowBytesPerImageFromRegionMipmapLevelSlice() bool

	// optional
	GetBytesBytesPerRowFromRegionMipmapLevel(pixelBytes unsafe.Pointer, bytesPerRow uint, region Region, level uint)
	HasGetBytesBytesPerRowFromRegionMipmapLevel() bool

	// optional
	NewTextureViewWithPixelFormat(pixelFormat PixelFormat) TextureObject
	HasNewTextureViewWithPixelFormat() bool

	// optional
	NewTextureViewWithPixelFormatTextureTypeLevelsSlices(pixelFormat PixelFormat, textureType TextureType, levelRange foundation.Range, sliceRange foundation.Range) TextureObject
	HasNewTextureViewWithPixelFormatTextureTypeLevelsSlices() bool

	// optional
	ReplaceRegionMipmapLevelSliceWithBytesBytesPerRowBytesPerImage(region Region, level uint, slice uint, pixelBytes unsafe.Pointer, bytesPerRow uint, bytesPerImage uint)
	HasReplaceRegionMipmapLevelSliceWithBytesBytesPerRowBytesPerImage() bool

	// optional
	NewTextureViewWithPixelFormatTextureTypeLevelsSlicesSwizzle(pixelFormat PixelFormat, textureType TextureType, levelRange foundation.Range, sliceRange foundation.Range, swizzle TextureSwizzleChannels) TextureObject
	HasNewTextureViewWithPixelFormatTextureTypeLevelsSlicesSwizzle() bool

	// optional
	NewSharedTextureHandle() SharedTextureHandle
	HasNewSharedTextureHandle() bool

	// optional
	NewRemoteTextureViewForDevice(device DeviceObject) TextureObject
	HasNewRemoteTextureViewForDevice() bool

	// optional
	ReplaceRegionMipmapLevelWithBytesBytesPerRow(region Region, level uint, pixelBytes unsafe.Pointer, bytesPerRow uint)
	HasReplaceRegionMipmapLevelWithBytesBytesPerRow() bool

	// optional
	Height() uint
	HasHeight() bool

	// optional
	AllowGPUOptimizedContents() bool
	HasAllowGPUOptimizedContents() bool

	// optional
	SampleCount() uint
	HasSampleCount() bool

	// optional
	Swizzle() TextureSwizzleChannels
	HasSwizzle() bool

	// optional
	IsFramebufferOnly() bool
	HasIsFramebufferOnly() bool

	// optional
	Usage() TextureUsage
	HasUsage() bool

	// optional
	Depth() uint
	HasDepth() bool

	// optional
	Buffer() BufferObject
	HasBuffer() bool

	// optional
	ArrayLength() uint
	HasArrayLength() bool

	// optional
	FirstMipmapInTail() uint
	HasFirstMipmapInTail() bool

	// optional
	BufferOffset() uint
	HasBufferOffset() bool

	// optional
	ParentRelativeLevel() uint
	HasParentRelativeLevel() bool

	// optional
	RemoteStorageTexture() TextureObject
	HasRemoteStorageTexture() bool

	// optional
	ParentRelativeSlice() uint
	HasParentRelativeSlice() bool

	// optional
	CompressionType() TextureCompressionType
	HasCompressionType() bool

	// optional
	PixelFormat() PixelFormat
	HasPixelFormat() bool

	// optional
	BufferBytesPerRow() uint
	HasBufferBytesPerRow() bool

	// optional
	Width() uint
	HasWidth() bool

	// optional
	IsShareable() bool
	HasIsShareable() bool

	// optional
	TailSizeInBytes() uint
	HasTailSizeInBytes() bool

	// optional
	ParentTexture() TextureObject
	HasParentTexture() bool

	// optional
	TextureType() TextureType
	HasTextureType() bool

	// optional
	MipmapLevelCount() uint
	HasMipmapLevelCount() bool

	// optional
	Iosurface() iosurface.Ref
	HasIosurface() bool

	// optional
	IsSparse() bool
	HasIsSparse() bool

	// optional
	IosurfacePlane() uint
	HasIosurfacePlane() bool
}

// ensure impl type implements protocol interface
var _ PTexture = (*TextureObject)(nil)

// A concrete type for the [PTexture] protocol.
type TextureObject struct {
	objc.Object
}

func (t_ TextureObject) HasGetBytesBytesPerRowBytesPerImageFromRegionMipmapLevelSlice() bool {
	return t_.RespondsToSelector(objc.Sel("getBytes:bytesPerRow:bytesPerImage:fromRegion:mipmapLevel:slice:"))
}

// Copies pixel data from the texture to a buffer in system memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1516318-getbytes?language=objc
func (t_ TextureObject) GetBytesBytesPerRowBytesPerImageFromRegionMipmapLevelSlice(pixelBytes unsafe.Pointer, bytesPerRow uint, bytesPerImage uint, region Region, level uint, slice uint) {
	objc.Call[objc.Void](t_, objc.Sel("getBytes:bytesPerRow:bytesPerImage:fromRegion:mipmapLevel:slice:"), pixelBytes, bytesPerRow, bytesPerImage, region, level, slice)
}

func (t_ TextureObject) HasGetBytesBytesPerRowFromRegionMipmapLevel() bool {
	return t_.RespondsToSelector(objc.Sel("getBytes:bytesPerRow:fromRegion:mipmapLevel:"))
}

// Copies pixel data from the first slice of the texture to a buffer in system memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515751-getbytes?language=objc
func (t_ TextureObject) GetBytesBytesPerRowFromRegionMipmapLevel(pixelBytes unsafe.Pointer, bytesPerRow uint, region Region, level uint) {
	objc.Call[objc.Void](t_, objc.Sel("getBytes:bytesPerRow:fromRegion:mipmapLevel:"), pixelBytes, bytesPerRow, region, level)
}

func (t_ TextureObject) HasNewTextureViewWithPixelFormat() bool {
	return t_.RespondsToSelector(objc.Sel("newTextureViewWithPixelFormat:"))
}

// Creates a new view of the texture, reinterpreting its data using a different pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515598-newtextureviewwithpixelformat?language=objc
func (t_ TextureObject) NewTextureViewWithPixelFormat(pixelFormat PixelFormat) TextureObject {
	rv := objc.Call[TextureObject](t_, objc.Sel("newTextureViewWithPixelFormat:"), pixelFormat)
	return rv
}

func (t_ TextureObject) HasNewTextureViewWithPixelFormatTextureTypeLevelsSlices() bool {
	return t_.RespondsToSelector(objc.Sel("newTextureViewWithPixelFormat:textureType:levels:slices:"))
}

// Creates a new view of the texture, reinterpreting a subset of its data using a different type and pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515409-newtextureviewwithpixelformat?language=objc
func (t_ TextureObject) NewTextureViewWithPixelFormatTextureTypeLevelsSlices(pixelFormat PixelFormat, textureType TextureType, levelRange foundation.Range, sliceRange foundation.Range) TextureObject {
	rv := objc.Call[TextureObject](t_, objc.Sel("newTextureViewWithPixelFormat:textureType:levels:slices:"), pixelFormat, textureType, levelRange, sliceRange)
	return rv
}

func (t_ TextureObject) HasReplaceRegionMipmapLevelSliceWithBytesBytesPerRowBytesPerImage() bool {
	return t_.RespondsToSelector(objc.Sel("replaceRegion:mipmapLevel:slice:withBytes:bytesPerRow:bytesPerImage:"))
}

// Copies pixel data into a section of a texture slice. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515679-replaceregion?language=objc
func (t_ TextureObject) ReplaceRegionMipmapLevelSliceWithBytesBytesPerRowBytesPerImage(region Region, level uint, slice uint, pixelBytes unsafe.Pointer, bytesPerRow uint, bytesPerImage uint) {
	objc.Call[objc.Void](t_, objc.Sel("replaceRegion:mipmapLevel:slice:withBytes:bytesPerRow:bytesPerImage:"), region, level, slice, pixelBytes, bytesPerRow, bytesPerImage)
}

func (t_ TextureObject) HasNewTextureViewWithPixelFormatTextureTypeLevelsSlicesSwizzle() bool {
	return t_.RespondsToSelector(objc.Sel("newTextureViewWithPixelFormat:textureType:levels:slices:swizzle:"))
}

// Creates a new view of the texture, reinterpreting a subset of its data using a different type, pixel format, and swizzle pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/3114303-newtextureviewwithpixelformat?language=objc
func (t_ TextureObject) NewTextureViewWithPixelFormatTextureTypeLevelsSlicesSwizzle(pixelFormat PixelFormat, textureType TextureType, levelRange foundation.Range, sliceRange foundation.Range, swizzle TextureSwizzleChannels) TextureObject {
	rv := objc.Call[TextureObject](t_, objc.Sel("newTextureViewWithPixelFormat:textureType:levels:slices:swizzle:"), pixelFormat, textureType, levelRange, sliceRange, swizzle)
	return rv
}

func (t_ TextureObject) HasNewSharedTextureHandle() bool {
	return t_.RespondsToSelector(objc.Sel("newSharedTextureHandle"))
}

// Creates a new texture handle from a shareable texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/2981032-newsharedtexturehandle?language=objc
func (t_ TextureObject) NewSharedTextureHandle() SharedTextureHandle {
	rv := objc.Call[SharedTextureHandle](t_, objc.Sel("newSharedTextureHandle"))
	return rv
}

func (t_ TextureObject) HasNewRemoteTextureViewForDevice() bool {
	return t_.RespondsToSelector(objc.Sel("newRemoteTextureViewForDevice:"))
}

// Creates a remote texture view for another GPU in the same peer group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/2967449-newremotetextureviewfordevice?language=objc
func (t_ TextureObject) NewRemoteTextureViewForDevice(device DeviceObject) TextureObject {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[TextureObject](t_, objc.Sel("newRemoteTextureViewForDevice:"), po0)
	return rv
}

func (t_ TextureObject) HasReplaceRegionMipmapLevelWithBytesBytesPerRow() bool {
	return t_.RespondsToSelector(objc.Sel("replaceRegion:mipmapLevel:withBytes:bytesPerRow:"))
}

// Copies a block of pixels into a section of texture slice 0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515464-replaceregion?language=objc
func (t_ TextureObject) ReplaceRegionMipmapLevelWithBytesBytesPerRow(region Region, level uint, pixelBytes unsafe.Pointer, bytesPerRow uint) {
	objc.Call[objc.Void](t_, objc.Sel("replaceRegion:mipmapLevel:withBytes:bytesPerRow:"), region, level, pixelBytes, bytesPerRow)
}

func (t_ TextureObject) HasHeight() bool {
	return t_.RespondsToSelector(objc.Sel("height"))
}

// The height of the texture image for the base level mipmap, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515938-height?language=objc
func (t_ TextureObject) Height() uint {
	rv := objc.Call[uint](t_, objc.Sel("height"))
	return rv
}

func (t_ TextureObject) HasAllowGPUOptimizedContents() bool {
	return t_.RespondsToSelector(objc.Sel("allowGPUOptimizedContents"))
}

// A Boolean value indicating whether the GPU is allowed to adjust the contents of the texture to improve GPU performance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/2966640-allowgpuoptimizedcontents?language=objc
func (t_ TextureObject) AllowGPUOptimizedContents() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowGPUOptimizedContents"))
	return rv
}

func (t_ TextureObject) HasSampleCount() bool {
	return t_.RespondsToSelector(objc.Sel("sampleCount"))
}

// The number of samples in each pixel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515443-samplecount?language=objc
func (t_ TextureObject) SampleCount() uint {
	rv := objc.Call[uint](t_, objc.Sel("sampleCount"))
	return rv
}

func (t_ TextureObject) HasSwizzle() bool {
	return t_.RespondsToSelector(objc.Sel("swizzle"))
}

// The pattern that the GPU applies to pixels when you read or sample pixels from the texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/3114304-swizzle?language=objc
func (t_ TextureObject) Swizzle() TextureSwizzleChannels {
	rv := objc.Call[TextureSwizzleChannels](t_, objc.Sel("swizzle"))
	return rv
}

func (t_ TextureObject) HasIsFramebufferOnly() bool {
	return t_.RespondsToSelector(objc.Sel("isFramebufferOnly"))
}

// A Boolean value that indicates whether the texture can only be used as a render target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515749-framebufferonly?language=objc
func (t_ TextureObject) IsFramebufferOnly() bool {
	rv := objc.Call[bool](t_, objc.Sel("isFramebufferOnly"))
	return rv
}

func (t_ TextureObject) HasUsage() bool {
	return t_.RespondsToSelector(objc.Sel("usage"))
}

// Options that determine how you can use the texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515763-usage?language=objc
func (t_ TextureObject) Usage() TextureUsage {
	rv := objc.Call[TextureUsage](t_, objc.Sel("usage"))
	return rv
}

func (t_ TextureObject) HasDepth() bool {
	return t_.RespondsToSelector(objc.Sel("depth"))
}

// The depth of the texture image for the base level mipmap, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515942-depth?language=objc
func (t_ TextureObject) Depth() uint {
	rv := objc.Call[uint](t_, objc.Sel("depth"))
	return rv
}

func (t_ TextureObject) HasBuffer() bool {
	return t_.RespondsToSelector(objc.Sel("buffer"))
}

// The source buffer used to create this texture, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1619090-buffer?language=objc
func (t_ TextureObject) Buffer() BufferObject {
	rv := objc.Call[BufferObject](t_, objc.Sel("buffer"))
	return rv
}

func (t_ TextureObject) HasArrayLength() bool {
	return t_.RespondsToSelector(objc.Sel("arrayLength"))
}

// The number of slices in the texture array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515382-arraylength?language=objc
func (t_ TextureObject) ArrayLength() uint {
	rv := objc.Call[uint](t_, objc.Sel("arrayLength"))
	return rv
}

func (t_ TextureObject) HasFirstMipmapInTail() bool {
	return t_.RespondsToSelector(objc.Sel("firstMipmapInTail"))
}

// The index of the first mipmap in the tail. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/3043999-firstmipmapintail?language=objc
func (t_ TextureObject) FirstMipmapInTail() uint {
	rv := objc.Call[uint](t_, objc.Sel("firstMipmapInTail"))
	return rv
}

func (t_ TextureObject) HasBufferOffset() bool {
	return t_.RespondsToSelector(objc.Sel("bufferOffset"))
}

// The offset in the source buffer where the texture’s data comes from. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1619019-bufferoffset?language=objc
func (t_ TextureObject) BufferOffset() uint {
	rv := objc.Call[uint](t_, objc.Sel("bufferOffset"))
	return rv
}

func (t_ TextureObject) HasParentRelativeLevel() bool {
	return t_.RespondsToSelector(objc.Sel("parentRelativeLevel"))
}

// The base level of the parent texture used to create this texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1516265-parentrelativelevel?language=objc
func (t_ TextureObject) ParentRelativeLevel() uint {
	rv := objc.Call[uint](t_, objc.Sel("parentRelativeLevel"))
	return rv
}

func (t_ TextureObject) HasRemoteStorageTexture() bool {
	return t_.RespondsToSelector(objc.Sel("remoteStorageTexture"))
}

// The texture on another GPU that the texture was created from, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/2967451-remotestoragetexture?language=objc
func (t_ TextureObject) RemoteStorageTexture() TextureObject {
	rv := objc.Call[TextureObject](t_, objc.Sel("remoteStorageTexture"))
	return rv
}

func (t_ TextureObject) HasParentRelativeSlice() bool {
	return t_.RespondsToSelector(objc.Sel("parentRelativeSlice"))
}

// The base slice of the parent texture used to create this texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1516221-parentrelativeslice?language=objc
func (t_ TextureObject) ParentRelativeSlice() uint {
	rv := objc.Call[uint](t_, objc.Sel("parentRelativeSlice"))
	return rv
}

func (t_ TextureObject) HasCompressionType() bool {
	return t_.RespondsToSelector(objc.Sel("compressionType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/3875341-compressiontype?language=objc
func (t_ TextureObject) CompressionType() TextureCompressionType {
	rv := objc.Call[TextureCompressionType](t_, objc.Sel("compressionType"))
	return rv
}

func (t_ TextureObject) HasPixelFormat() bool {
	return t_.RespondsToSelector(objc.Sel("pixelFormat"))
}

// The format of pixels in the texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515344-pixelformat?language=objc
func (t_ TextureObject) PixelFormat() PixelFormat {
	rv := objc.Call[PixelFormat](t_, objc.Sel("pixelFormat"))
	return rv
}

func (t_ TextureObject) HasBufferBytesPerRow() bool {
	return t_.RespondsToSelector(objc.Sel("bufferBytesPerRow"))
}

// The number of bytes in each row of the texture’s source buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1619175-bufferbytesperrow?language=objc
func (t_ TextureObject) BufferBytesPerRow() uint {
	rv := objc.Call[uint](t_, objc.Sel("bufferBytesPerRow"))
	return rv
}

func (t_ TextureObject) HasWidth() bool {
	return t_.RespondsToSelector(objc.Sel("width"))
}

// The width of the texture image for the base level mipmap, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515339-width?language=objc
func (t_ TextureObject) Width() uint {
	rv := objc.Call[uint](t_, objc.Sel("width"))
	return rv
}

func (t_ TextureObject) HasIsShareable() bool {
	return t_.RespondsToSelector(objc.Sel("isShareable"))
}

// A Boolean indicating whether this texture can be shared with other processes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/2998889-shareable?language=objc
func (t_ TextureObject) IsShareable() bool {
	rv := objc.Call[bool](t_, objc.Sel("isShareable"))
	return rv
}

func (t_ TextureObject) HasTailSizeInBytes() bool {
	return t_.RespondsToSelector(objc.Sel("tailSizeInBytes"))
}

// The size of the sparse texture tail, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/3044002-tailsizeinbytes?language=objc
func (t_ TextureObject) TailSizeInBytes() uint {
	rv := objc.Call[uint](t_, objc.Sel("tailSizeInBytes"))
	return rv
}

func (t_ TextureObject) HasParentTexture() bool {
	return t_.RespondsToSelector(objc.Sel("parentTexture"))
}

// The parent texture used to create this texture, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515372-parenttexture?language=objc
func (t_ TextureObject) ParentTexture() TextureObject {
	rv := objc.Call[TextureObject](t_, objc.Sel("parentTexture"))
	return rv
}

func (t_ TextureObject) HasTextureType() bool {
	return t_.RespondsToSelector(objc.Sel("textureType"))
}

// The dimension and arrangement of the texture image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515517-texturetype?language=objc
func (t_ TextureObject) TextureType() TextureType {
	rv := objc.Call[TextureType](t_, objc.Sel("textureType"))
	return rv
}

func (t_ TextureObject) HasMipmapLevelCount() bool {
	return t_.RespondsToSelector(objc.Sel("mipmapLevelCount"))
}

// The number of mipmap levels in the texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515677-mipmaplevelcount?language=objc
func (t_ TextureObject) MipmapLevelCount() uint {
	rv := objc.Call[uint](t_, objc.Sel("mipmapLevelCount"))
	return rv
}

func (t_ TextureObject) HasIosurface() bool {
	return t_.RespondsToSelector(objc.Sel("iosurface"))
}

// A reference to the IOSurface the texture was created from, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1516104-iosurface?language=objc
func (t_ TextureObject) Iosurface() iosurface.Ref {
	rv := objc.Call[iosurface.Ref](t_, objc.Sel("iosurface"))
	return rv
}

func (t_ TextureObject) HasIsSparse() bool {
	return t_.RespondsToSelector(objc.Sel("isSparse"))
}

// A Boolean value that indicates whether this is a sparse texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/3153124-issparse?language=objc
func (t_ TextureObject) IsSparse() bool {
	rv := objc.Call[bool](t_, objc.Sel("isSparse"))
	return rv
}

func (t_ TextureObject) HasIosurfacePlane() bool {
	return t_.RespondsToSelector(objc.Sel("iosurfacePlane"))
}

// The plane of the IOSurface to reference if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515986-iosurfaceplane?language=objc
func (t_ TextureObject) IosurfacePlane() uint {
	rv := objc.Call[uint](t_, objc.Sel("iosurfacePlane"))
	return rv
}
