// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetReader] class.
var AssetReaderClass = _AssetReaderClass{objc.GetClass("AVAssetReader")}

type _AssetReaderClass struct {
	objc.Class
}

// An interface definition for the [AssetReader] class.
type IAssetReader interface {
	objc.IObject
	StartReading() bool
	AddOutput(output IAssetReaderOutput)
	CancelReading()
	CanAddOutput(output IAssetReaderOutput) bool
	Status() AssetReaderStatus
	TimeRange() coremedia.TimeRange
	SetTimeRange(value coremedia.TimeRange)
	Asset() Asset
	Error() foundation.Error
	Outputs() []AssetReaderOutput
}

// An object that reads media data from an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader?language=objc
type AssetReader struct {
	objc.Object
}

func AssetReaderFrom(ptr unsafe.Pointer) AssetReader {
	return AssetReader{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ AssetReader) InitWithAssetError(asset IAsset, outError unsafe.Pointer) AssetReader {
	rv := objc.Call[AssetReader](a_, objc.Sel("initWithAsset:error:"), asset, outError)
	return rv
}

// Creates an object to read media data from an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/1385593-initwithasset?language=objc
func NewAssetReaderWithAssetError(asset IAsset, outError unsafe.Pointer) AssetReader {
	instance := AssetReaderClass.Alloc().InitWithAssetError(asset, outError)
	instance.Autorelease()
	return instance
}

func (ac _AssetReaderClass) AssetReaderWithAssetError(asset IAsset, outError unsafe.Pointer) AssetReader {
	rv := objc.Call[AssetReader](ac, objc.Sel("assetReaderWithAsset:error:"), asset, outError)
	return rv
}

// Returns a new object to read media data from an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/1420148-assetreaderwithasset?language=objc
func AssetReader_AssetReaderWithAssetError(asset IAsset, outError unsafe.Pointer) AssetReader {
	return AssetReaderClass.AssetReaderWithAssetError(asset, outError)
}

func (ac _AssetReaderClass) Alloc() AssetReader {
	rv := objc.Call[AssetReader](ac, objc.Sel("alloc"))
	return rv
}

func (ac _AssetReaderClass) New() AssetReader {
	rv := objc.Call[AssetReader](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetReader() AssetReader {
	return AssetReaderClass.New()
}

func (a_ AssetReader) Init() AssetReader {
	rv := objc.Call[AssetReader](a_, objc.Sel("init"))
	return rv
}

// Prepares the asset reader to start reading sample buffers from the asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/1390286-startreading?language=objc
func (a_ AssetReader) StartReading() bool {
	rv := objc.Call[bool](a_, objc.Sel("startReading"))
	return rv
}

// Adds an output to the reader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/1390110-addoutput?language=objc
func (a_ AssetReader) AddOutput(output IAssetReaderOutput) {
	objc.Call[objc.Void](a_, objc.Sel("addOutput:"), output)
}

// Cancels any background work and stops the reader’s outputs from reading more samples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/1390258-cancelreading?language=objc
func (a_ AssetReader) CancelReading() {
	objc.Call[objc.Void](a_, objc.Sel("cancelReading"))
}

// Determines whether you can add the output to the asset reader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/1387485-canaddoutput?language=objc
func (a_ AssetReader) CanAddOutput(output IAssetReaderOutput) bool {
	rv := objc.Call[bool](a_, objc.Sel("canAddOutput:"), output)
	return rv
}

// The status of reading sample buffers from the asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/1390211-status?language=objc
func (a_ AssetReader) Status() AssetReaderStatus {
	rv := objc.Call[AssetReaderStatus](a_, objc.Sel("status"))
	return rv
}

// The time range within the asset to read. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/1386400-timerange?language=objc
func (a_ AssetReader) TimeRange() coremedia.TimeRange {
	rv := objc.Call[coremedia.TimeRange](a_, objc.Sel("timeRange"))
	return rv
}

// The time range within the asset to read. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/1386400-timerange?language=objc
func (a_ AssetReader) SetTimeRange(value coremedia.TimeRange) {
	objc.Call[objc.Void](a_, objc.Sel("setTimeRange:"), value)
}

// The asset from which to read media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/1389128-asset?language=objc
func (a_ AssetReader) Asset() Asset {
	rv := objc.Call[Asset](a_, objc.Sel("asset"))
	return rv
}

// An error that describes the reason for a failure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/1388114-error?language=objc
func (a_ AssetReader) Error() foundation.Error {
	rv := objc.Call[foundation.Error](a_, objc.Sel("error"))
	return rv
}

// The outputs from which you read media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/1387132-outputs?language=objc
func (a_ AssetReader) Outputs() []AssetReaderOutput {
	rv := objc.Call[[]AssetReaderOutput](a_, objc.Sel("outputs"))
	return rv
}
