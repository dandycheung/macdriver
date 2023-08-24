// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetWriterInputPixelBufferAdaptor] class.
var AssetWriterInputPixelBufferAdaptorClass = _AssetWriterInputPixelBufferAdaptorClass{objc.GetClass("AVAssetWriterInputPixelBufferAdaptor")}

type _AssetWriterInputPixelBufferAdaptorClass struct {
	objc.Class
}

// An interface definition for the [AssetWriterInputPixelBufferAdaptor] class.
type IAssetWriterInputPixelBufferAdaptor interface {
	objc.IObject
	AppendPixelBufferWithPresentationTime(pixelBuffer corevideo.PixelBufferRef, presentationTime coremedia.Time) bool
	AssetWriterInput() AssetWriterInput
	SourcePixelBufferAttributes() map[string]objc.Object
	PixelBufferPool() corevideo.PixelBufferPoolRef
}

// An object that appends video samples to an asset writer input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputpixelbufferadaptor?language=objc
type AssetWriterInputPixelBufferAdaptor struct {
	objc.Object
}

func AssetWriterInputPixelBufferAdaptorFrom(ptr unsafe.Pointer) AssetWriterInputPixelBufferAdaptor {
	return AssetWriterInputPixelBufferAdaptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetWriterInputPixelBufferAdaptorClass) AssetWriterInputPixelBufferAdaptorWithAssetWriterInputSourcePixelBufferAttributes(input IAssetWriterInput, sourcePixelBufferAttributes map[string]objc.IObject) AssetWriterInputPixelBufferAdaptor {
	rv := objc.Call[AssetWriterInputPixelBufferAdaptor](ac, objc.Sel("assetWriterInputPixelBufferAdaptorWithAssetWriterInput:sourcePixelBufferAttributes:"), objc.Ptr(input), sourcePixelBufferAttributes)
	return rv
}

// Returns a new pixel buffer adaptor that appends pixel buffers to write to the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputpixelbufferadaptor/1449096-assetwriterinputpixelbufferadapt?language=objc
func AssetWriterInputPixelBufferAdaptor_AssetWriterInputPixelBufferAdaptorWithAssetWriterInputSourcePixelBufferAttributes(input IAssetWriterInput, sourcePixelBufferAttributes map[string]objc.IObject) AssetWriterInputPixelBufferAdaptor {
	return AssetWriterInputPixelBufferAdaptorClass.AssetWriterInputPixelBufferAdaptorWithAssetWriterInputSourcePixelBufferAttributes(input, sourcePixelBufferAttributes)
}

func (a_ AssetWriterInputPixelBufferAdaptor) InitWithAssetWriterInputSourcePixelBufferAttributes(input IAssetWriterInput, sourcePixelBufferAttributes map[string]objc.IObject) AssetWriterInputPixelBufferAdaptor {
	rv := objc.Call[AssetWriterInputPixelBufferAdaptor](a_, objc.Sel("initWithAssetWriterInput:sourcePixelBufferAttributes:"), objc.Ptr(input), sourcePixelBufferAttributes)
	return rv
}

// Creates a new pixel buffer adaptor to receive pixel buffers for writing to the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputpixelbufferadaptor/1390639-initwithassetwriterinput?language=objc
func NewAssetWriterInputPixelBufferAdaptorWithAssetWriterInputSourcePixelBufferAttributes(input IAssetWriterInput, sourcePixelBufferAttributes map[string]objc.IObject) AssetWriterInputPixelBufferAdaptor {
	instance := AssetWriterInputPixelBufferAdaptorClass.Alloc().InitWithAssetWriterInputSourcePixelBufferAttributes(input, sourcePixelBufferAttributes)
	instance.Autorelease()
	return instance
}

func (ac _AssetWriterInputPixelBufferAdaptorClass) Alloc() AssetWriterInputPixelBufferAdaptor {
	rv := objc.Call[AssetWriterInputPixelBufferAdaptor](ac, objc.Sel("alloc"))
	return rv
}

func AssetWriterInputPixelBufferAdaptor_Alloc() AssetWriterInputPixelBufferAdaptor {
	return AssetWriterInputPixelBufferAdaptorClass.Alloc()
}

func (ac _AssetWriterInputPixelBufferAdaptorClass) New() AssetWriterInputPixelBufferAdaptor {
	rv := objc.Call[AssetWriterInputPixelBufferAdaptor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetWriterInputPixelBufferAdaptor() AssetWriterInputPixelBufferAdaptor {
	return AssetWriterInputPixelBufferAdaptorClass.New()
}

func (a_ AssetWriterInputPixelBufferAdaptor) Init() AssetWriterInputPixelBufferAdaptor {
	rv := objc.Call[AssetWriterInputPixelBufferAdaptor](a_, objc.Sel("init"))
	return rv
}

// Appends a pixel buffer to the adaptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputpixelbufferadaptor/1388102-appendpixelbuffer?language=objc
func (a_ AssetWriterInputPixelBufferAdaptor) AppendPixelBufferWithPresentationTime(pixelBuffer corevideo.PixelBufferRef, presentationTime coremedia.Time) bool {
	rv := objc.Call[bool](a_, objc.Sel("appendPixelBuffer:withPresentationTime:"), pixelBuffer, presentationTime)
	return rv
}

// The asset writer input to which the adaptor appends pixel buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputpixelbufferadaptor/1387565-assetwriterinput?language=objc
func (a_ AssetWriterInputPixelBufferAdaptor) AssetWriterInput() AssetWriterInput {
	rv := objc.Call[AssetWriterInput](a_, objc.Sel("assetWriterInput"))
	return rv
}

// The attributes of the pixel buffers that the pool contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputpixelbufferadaptor/1387829-sourcepixelbufferattributes?language=objc
func (a_ AssetWriterInputPixelBufferAdaptor) SourcePixelBufferAttributes() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](a_, objc.Sel("sourcePixelBufferAttributes"))
	return rv
}

// A pool of pixel buffers to append to the adaptor’s input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputpixelbufferadaptor/1389662-pixelbufferpool?language=objc
func (a_ AssetWriterInputPixelBufferAdaptor) PixelBufferPool() corevideo.PixelBufferPoolRef {
	rv := objc.Call[corevideo.PixelBufferPoolRef](a_, objc.Sel("pixelBufferPool"))
	return rv
}