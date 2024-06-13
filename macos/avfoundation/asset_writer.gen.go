// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/uti"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetWriter] class.
var AssetWriterClass = _AssetWriterClass{objc.GetClass("AVAssetWriter")}

type _AssetWriterClass struct {
	objc.Class
}

// An interface definition for the [AssetWriter] class.
type IAssetWriter interface {
	objc.IObject
	CanAddInput(input IAssetWriterInput) bool
	CancelWriting()
	AddInputGroup(inputGroup IAssetWriterInputGroup)
	AddInput(input IAssetWriterInput)
	StartWriting() bool
	FinishWritingWithCompletionHandler(handler func())
	CanApplyOutputSettingsForMediaType(outputSettings map[string]objc.IObject, mediaType MediaType) bool
	FlushSegment()
	StartSessionAtSourceTime(startTime coremedia.Time)
	CanAddInputGroup(inputGroup IAssetWriterInputGroup) bool
	EndSessionAtSourceTime(endTime coremedia.Time)
	OutputURL() foundation.URL
	MovieTimeScale() coremedia.TimeScale
	SetMovieTimeScale(value coremedia.TimeScale)
	Inputs() []AssetWriterInput
	OutputFileTypeProfile() FileTypeProfile
	SetOutputFileTypeProfile(value FileTypeProfile)
	MovieFragmentInterval() coremedia.Time
	SetMovieFragmentInterval(value coremedia.Time)
	InitialSegmentStartTime() coremedia.Time
	SetInitialSegmentStartTime(value coremedia.Time)
	Metadata() []MetadataItem
	SetMetadata(value []IMetadataItem)
	PreferredOutputSegmentInterval() coremedia.Time
	SetPreferredOutputSegmentInterval(value coremedia.Time)
	OutputFileType() FileType
	AvailableMediaTypes() []MediaType
	ProducesCombinableFragments() bool
	SetProducesCombinableFragments(value bool)
	Error() foundation.Error
	InputGroups() []AssetWriterInputGroup
	Status() AssetWriterStatus
	Delegate() AssetWriterDelegateObject
	SetDelegate(value PAssetWriterDelegate)
	SetDelegateObject(valueObject objc.IObject)
	InitialMovieFragmentSequenceNumber() int
	SetInitialMovieFragmentSequenceNumber(value int)
	ShouldOptimizeForNetworkUse() bool
	SetShouldOptimizeForNetworkUse(value bool)
	OverallDurationHint() coremedia.Time
	SetOverallDurationHint(value coremedia.Time)
	DirectoryForTemporaryFiles() foundation.URL
	SetDirectoryForTemporaryFiles(value foundation.IURL)
}

// An object that writes media data to a container file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter?language=objc
type AssetWriter struct {
	objc.Object
}

func AssetWriterFrom(ptr unsafe.Pointer) AssetWriter {
	return AssetWriter{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ AssetWriter) InitWithURLFileTypeError(outputURL foundation.IURL, outputFileType FileType, outError unsafe.Pointer) AssetWriter {
	rv := objc.Call[AssetWriter](a_, objc.Sel("initWithURL:fileType:error:"), outputURL, outputFileType, outError)
	return rv
}

// Creates an object that writes media data to a container file at the output URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1389201-initwithurl?language=objc
func NewAssetWriterWithURLFileTypeError(outputURL foundation.IURL, outputFileType FileType, outError unsafe.Pointer) AssetWriter {
	instance := AssetWriterClass.Alloc().InitWithURLFileTypeError(outputURL, outputFileType, outError)
	instance.Autorelease()
	return instance
}

func (ac _AssetWriterClass) AssetWriterWithURLFileTypeError(outputURL foundation.IURL, outputFileType FileType, outError unsafe.Pointer) AssetWriter {
	rv := objc.Call[AssetWriter](ac, objc.Sel("assetWriterWithURL:fileType:error:"), outputURL, outputFileType, outError)
	return rv
}

// Returns a new object that writes media data to a container file at the output URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1426663-assetwriterwithurl?language=objc
func AssetWriter_AssetWriterWithURLFileTypeError(outputURL foundation.IURL, outputFileType FileType, outError unsafe.Pointer) AssetWriter {
	return AssetWriterClass.AssetWriterWithURLFileTypeError(outputURL, outputFileType, outError)
}

func (a_ AssetWriter) InitWithContentType(outputContentType uti.IType) AssetWriter {
	rv := objc.Call[AssetWriter](a_, objc.Sel("initWithContentType:"), outputContentType)
	return rv
}

// Creates an object that outputs segment data in a specified container format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/3563935-initwithcontenttype?language=objc
func NewAssetWriterWithContentType(outputContentType uti.IType) AssetWriter {
	instance := AssetWriterClass.Alloc().InitWithContentType(outputContentType)
	instance.Autorelease()
	return instance
}

func (ac _AssetWriterClass) Alloc() AssetWriter {
	rv := objc.Call[AssetWriter](ac, objc.Sel("alloc"))
	return rv
}

func (ac _AssetWriterClass) New() AssetWriter {
	rv := objc.Call[AssetWriter](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetWriter() AssetWriter {
	return AssetWriterClass.New()
}

func (a_ AssetWriter) Init() AssetWriter {
	rv := objc.Call[AssetWriter](a_, objc.Sel("init"))
	return rv
}

// Determines whether the asset writer supports adding the input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1387863-canaddinput?language=objc
func (a_ AssetWriter) CanAddInput(input IAssetWriterInput) bool {
	rv := objc.Call[bool](a_, objc.Sel("canAddInput:"), input)
	return rv
}

// Cancels the creation of the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1387234-cancelwriting?language=objc
func (a_ AssetWriter) CancelWriting() {
	objc.Call[objc.Void](a_, objc.Sel("cancelWriting"))
}

// Adds an input group to an asset writer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1385643-addinputgroup?language=objc
func (a_ AssetWriter) AddInputGroup(inputGroup IAssetWriterInputGroup) {
	objc.Call[objc.Void](a_, objc.Sel("addInputGroup:"), inputGroup)
}

// Adds an input to an asset writer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1390389-addinput?language=objc
func (a_ AssetWriter) AddInput(input IAssetWriterInput) {
	objc.Call[objc.Void](a_, objc.Sel("addInput:"), input)
}

// Tells the writer to start writing its output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1386724-startwriting?language=objc
func (a_ AssetWriter) StartWriting() bool {
	rv := objc.Call[bool](a_, objc.Sel("startWriting"))
	return rv
}

// Marks all unfinished inputs as finished and completes the writing of the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1390432-finishwritingwithcompletionhandl?language=objc
func (a_ AssetWriter) FinishWritingWithCompletionHandler(handler func()) {
	objc.Call[objc.Void](a_, objc.Sel("finishWritingWithCompletionHandler:"), handler)
}

// Determines whether the output file format supports the output settings for a specific media type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1388842-canapplyoutputsettings?language=objc
func (a_ AssetWriter) CanApplyOutputSettingsForMediaType(outputSettings map[string]objc.IObject, mediaType MediaType) bool {
	rv := objc.Call[bool](a_, objc.Sel("canApplyOutputSettings:forMediaType:"), outputSettings, mediaType)
	return rv
}

// Closes the current segment and outputs it to a delegate method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/3546586-flushsegment?language=objc
func (a_ AssetWriter) FlushSegment() {
	objc.Call[objc.Void](a_, objc.Sel("flushSegment"))
}

// Starts an asset-writing session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1389908-startsessionatsourcetime?language=objc
func (a_ AssetWriter) StartSessionAtSourceTime(startTime coremedia.Time) {
	objc.Call[objc.Void](a_, objc.Sel("startSessionAtSourceTime:"), startTime)
}

// Determines whether the asset writer supports adding the input group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1386698-canaddinputgroup?language=objc
func (a_ AssetWriter) CanAddInputGroup(inputGroup IAssetWriterInputGroup) bool {
	rv := objc.Call[bool](a_, objc.Sel("canAddInputGroup:"), inputGroup)
	return rv
}

// Finishes an asset-writing session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1389921-endsessionatsourcetime?language=objc
func (a_ AssetWriter) EndSessionAtSourceTime(endTime coremedia.Time) {
	objc.Call[objc.Void](a_, objc.Sel("endSessionAtSourceTime:"), endTime)
}

// The location of the container file that the writer outputs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1387731-outputurl?language=objc
func (a_ AssetWriter) OutputURL() foundation.URL {
	rv := objc.Call[foundation.URL](a_, objc.Sel("outputURL"))
	return rv
}

// The time scale of the movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1386762-movietimescale?language=objc
func (a_ AssetWriter) MovieTimeScale() coremedia.TimeScale {
	rv := objc.Call[coremedia.TimeScale](a_, objc.Sel("movieTimeScale"))
	return rv
}

// The time scale of the movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1386762-movietimescale?language=objc
func (a_ AssetWriter) SetMovieTimeScale(value coremedia.TimeScale) {
	objc.Call[objc.Void](a_, objc.Sel("setMovieTimeScale:"), value)
}

// The inputs an asset writer contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1388264-inputs?language=objc
func (a_ AssetWriter) Inputs() []AssetWriterInput {
	rv := objc.Call[[]AssetWriterInput](a_, objc.Sel("inputs"))
	return rv
}

// A profile for the output file type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/3546589-outputfiletypeprofile?language=objc
func (a_ AssetWriter) OutputFileTypeProfile() FileTypeProfile {
	rv := objc.Call[FileTypeProfile](a_, objc.Sel("outputFileTypeProfile"))
	return rv
}

// A profile for the output file type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/3546589-outputfiletypeprofile?language=objc
func (a_ AssetWriter) SetOutputFileTypeProfile(value FileTypeProfile) {
	objc.Call[objc.Void](a_, objc.Sel("setOutputFileTypeProfile:"), value)
}

// The interval at which to write movie fragments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1387469-moviefragmentinterval?language=objc
func (a_ AssetWriter) MovieFragmentInterval() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("movieFragmentInterval"))
	return rv
}

// The interval at which to write movie fragments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1387469-moviefragmentinterval?language=objc
func (a_ AssetWriter) SetMovieFragmentInterval(value coremedia.Time) {
	objc.Call[objc.Void](a_, objc.Sel("setMovieFragmentInterval:"), value)
}

// The start time of the initial segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/3546588-initialsegmentstarttime?language=objc
func (a_ AssetWriter) InitialSegmentStartTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("initialSegmentStartTime"))
	return rv
}

// The start time of the initial segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/3546588-initialsegmentstarttime?language=objc
func (a_ AssetWriter) SetInitialSegmentStartTime(value coremedia.Time) {
	objc.Call[objc.Void](a_, objc.Sel("setInitialSegmentStartTime:"), value)
}

// An array of metadata items to write to the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1387974-metadata?language=objc
func (a_ AssetWriter) Metadata() []MetadataItem {
	rv := objc.Call[[]MetadataItem](a_, objc.Sel("metadata"))
	return rv
}

// An array of metadata items to write to the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1387974-metadata?language=objc
func (a_ AssetWriter) SetMetadata(value []IMetadataItem) {
	objc.Call[objc.Void](a_, objc.Sel("setMetadata:"), value)
}

// The interval of output segments that you prefer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/3546590-preferredoutputsegmentinterval?language=objc
func (a_ AssetWriter) PreferredOutputSegmentInterval() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("preferredOutputSegmentInterval"))
	return rv
}

// The interval of output segments that you prefer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/3546590-preferredoutputsegmentinterval?language=objc
func (a_ AssetWriter) SetPreferredOutputSegmentInterval(value coremedia.Time) {
	objc.Call[objc.Void](a_, objc.Sel("setPreferredOutputSegmentInterval:"), value)
}

// The type of container file that the writer outputs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1387349-outputfiletype?language=objc
func (a_ AssetWriter) OutputFileType() FileType {
	rv := objc.Call[FileType](a_, objc.Sel("outputFileType"))
	return rv
}

// The media types the asset writer supports adding as inputs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1388730-availablemediatypes?language=objc
func (a_ AssetWriter) AvailableMediaTypes() []MediaType {
	rv := objc.Call[[]MediaType](a_, objc.Sel("availableMediaTypes"))
	return rv
}

// A Boolean value that indicates whether the asset writer outputs movie fragments suitable for combining with others. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/3626025-producescombinablefragments?language=objc
func (a_ AssetWriter) ProducesCombinableFragments() bool {
	rv := objc.Call[bool](a_, objc.Sel("producesCombinableFragments"))
	return rv
}

// A Boolean value that indicates whether the asset writer outputs movie fragments suitable for combining with others. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/3626025-producescombinablefragments?language=objc
func (a_ AssetWriter) SetProducesCombinableFragments(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setProducesCombinableFragments:"), value)
}

// An error object that describes an asset-writing failure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1390725-error?language=objc
func (a_ AssetWriter) Error() foundation.Error {
	rv := objc.Call[foundation.Error](a_, objc.Sel("error"))
	return rv
}

// The input groups an asset writer contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1388432-inputgroups?language=objc
func (a_ AssetWriter) InputGroups() []AssetWriterInputGroup {
	rv := objc.Call[[]AssetWriterInputGroup](a_, objc.Sel("inputGroups"))
	return rv
}

// The status of writing samples to the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1389335-status?language=objc
func (a_ AssetWriter) Status() AssetWriterStatus {
	rv := objc.Call[AssetWriterStatus](a_, objc.Sel("status"))
	return rv
}

// A delegate object that responds to asset-writing events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/3546585-delegate?language=objc
func (a_ AssetWriter) Delegate() AssetWriterDelegateObject {
	rv := objc.Call[AssetWriterDelegateObject](a_, objc.Sel("delegate"))
	return rv
}

// A delegate object that responds to asset-writing events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/3546585-delegate?language=objc
func (a_ AssetWriter) SetDelegate(value PAssetWriterDelegate) {
	po0 := objc.WrapAsProtocol("AVAssetWriterDelegate", value)
	objc.SetAssociatedObject(a_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](a_, objc.Sel("setDelegate:"), po0)
}

// A delegate object that responds to asset-writing events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/3546585-delegate?language=objc
func (a_ AssetWriter) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setDelegate:"), valueObject)
}

// The sequence number of the initial movie fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/3577532-initialmoviefragmentsequencenumb?language=objc
func (a_ AssetWriter) InitialMovieFragmentSequenceNumber() int {
	rv := objc.Call[int](a_, objc.Sel("initialMovieFragmentSequenceNumber"))
	return rv
}

// The sequence number of the initial movie fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/3577532-initialmoviefragmentsequencenumb?language=objc
func (a_ AssetWriter) SetInitialMovieFragmentSequenceNumber(value int) {
	objc.Call[objc.Void](a_, objc.Sel("setInitialMovieFragmentSequenceNumber:"), value)
}

// A Boolean value that indicates whether to write the output file to make it more suitable for playback over a network. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1389811-shouldoptimizefornetworkuse?language=objc
func (a_ AssetWriter) ShouldOptimizeForNetworkUse() bool {
	rv := objc.Call[bool](a_, objc.Sel("shouldOptimizeForNetworkUse"))
	return rv
}

// A Boolean value that indicates whether to write the output file to make it more suitable for playback over a network. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1389811-shouldoptimizefornetworkuse?language=objc
func (a_ AssetWriter) SetShouldOptimizeForNetworkUse(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setShouldOptimizeForNetworkUse:"), value)
}

// A hint of the final duration of the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1388408-overalldurationhint?language=objc
func (a_ AssetWriter) OverallDurationHint() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("overallDurationHint"))
	return rv
}

// A hint of the final duration of the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1388408-overalldurationhint?language=objc
func (a_ AssetWriter) SetOverallDurationHint(value coremedia.Time) {
	objc.Call[objc.Void](a_, objc.Sel("setOverallDurationHint:"), value)
}

// A directory to contain temporary files that the export process generates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1387445-directoryfortemporaryfiles?language=objc
func (a_ AssetWriter) DirectoryForTemporaryFiles() foundation.URL {
	rv := objc.Call[foundation.URL](a_, objc.Sel("directoryForTemporaryFiles"))
	return rv
}

// A directory to contain temporary files that the export process generates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/1387445-directoryfortemporaryfiles?language=objc
func (a_ AssetWriter) SetDirectoryForTemporaryFiles(value foundation.IURL) {
	objc.Call[objc.Void](a_, objc.Sel("setDirectoryForTemporaryFiles:"), value)
}
