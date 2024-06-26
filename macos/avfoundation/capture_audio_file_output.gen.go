// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CaptureAudioFileOutput] class.
var CaptureAudioFileOutputClass = _CaptureAudioFileOutputClass{objc.GetClass("AVCaptureAudioFileOutput")}

type _CaptureAudioFileOutputClass struct {
	objc.Class
}

// An interface definition for the [CaptureAudioFileOutput] class.
type ICaptureAudioFileOutput interface {
	ICaptureFileOutput
	StartRecordingToOutputFileURLOutputFileTypeRecordingDelegate(outputFileURL foundation.IURL, fileType FileType, delegate PCaptureFileOutputRecordingDelegate)
	StartRecordingToOutputFileURLOutputFileTypeRecordingDelegateObject(outputFileURL foundation.IURL, fileType FileType, delegateObject objc.IObject)
	AudioSettings() map[string]objc.Object
	SetAudioSettings(value map[string]objc.IObject)
	Metadata() []MetadataItem
	SetMetadata(value []IMetadataItem)
}

// A capture output that records audio and saves the recorded audio to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiofileoutput?language=objc
type CaptureAudioFileOutput struct {
	CaptureFileOutput
}

func CaptureAudioFileOutputFrom(ptr unsafe.Pointer) CaptureAudioFileOutput {
	return CaptureAudioFileOutput{
		CaptureFileOutput: CaptureFileOutputFrom(ptr),
	}
}

func (cc _CaptureAudioFileOutputClass) New() CaptureAudioFileOutput {
	rv := objc.Call[CaptureAudioFileOutput](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureAudioFileOutput() CaptureAudioFileOutput {
	return CaptureAudioFileOutputClass.New()
}

func (c_ CaptureAudioFileOutput) Init() CaptureAudioFileOutput {
	rv := objc.Call[CaptureAudioFileOutput](c_, objc.Sel("init"))
	return rv
}

func (cc _CaptureAudioFileOutputClass) Alloc() CaptureAudioFileOutput {
	rv := objc.Call[CaptureAudioFileOutput](cc, objc.Sel("alloc"))
	return rv
}

// Returns an array containing UTIs identifying the file types AVCaptureAudioFileOutput can write. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiofileoutput/1390895-availableoutputfiletypes?language=objc
func (cc _CaptureAudioFileOutputClass) AvailableOutputFileTypes() []FileType {
	rv := objc.Call[[]FileType](cc, objc.Sel("availableOutputFileTypes"))
	return rv
}

// Returns an array containing UTIs identifying the file types AVCaptureAudioFileOutput can write. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiofileoutput/1390895-availableoutputfiletypes?language=objc
func CaptureAudioFileOutput_AvailableOutputFileTypes() []FileType {
	return CaptureAudioFileOutputClass.AvailableOutputFileTypes()
}

// Tells the receiver to start recording to a new file of the specified format, and specifies a delegate that will be notified when recording is finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiofileoutput/1387420-startrecordingtooutputfileurl?language=objc
func (c_ CaptureAudioFileOutput) StartRecordingToOutputFileURLOutputFileTypeRecordingDelegate(outputFileURL foundation.IURL, fileType FileType, delegate PCaptureFileOutputRecordingDelegate) {
	po2 := objc.WrapAsProtocol("AVCaptureFileOutputRecordingDelegate", delegate)
	objc.Call[objc.Void](c_, objc.Sel("startRecordingToOutputFileURL:outputFileType:recordingDelegate:"), outputFileURL, fileType, po2)
}

// Tells the receiver to start recording to a new file of the specified format, and specifies a delegate that will be notified when recording is finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiofileoutput/1387420-startrecordingtooutputfileurl?language=objc
func (c_ CaptureAudioFileOutput) StartRecordingToOutputFileURLOutputFileTypeRecordingDelegateObject(outputFileURL foundation.IURL, fileType FileType, delegateObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("startRecordingToOutputFileURL:outputFileType:recordingDelegate:"), outputFileURL, fileType, delegateObject)
}

// The settings used to decode or re-encode audio before it is output by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiofileoutput/1389958-audiosettings?language=objc
func (c_ CaptureAudioFileOutput) AudioSettings() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](c_, objc.Sel("audioSettings"))
	return rv
}

// The settings used to decode or re-encode audio before it is output by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiofileoutput/1389958-audiosettings?language=objc
func (c_ CaptureAudioFileOutput) SetAudioSettings(value map[string]objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setAudioSettings:"), value)
}

// A collection of metadata to be written to the receiver's output files. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiofileoutput/1389881-metadata?language=objc
func (c_ CaptureAudioFileOutput) Metadata() []MetadataItem {
	rv := objc.Call[[]MetadataItem](c_, objc.Sel("metadata"))
	return rv
}

// A collection of metadata to be written to the receiver's output files. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiofileoutput/1389881-metadata?language=objc
func (c_ CaptureAudioFileOutput) SetMetadata(value []IMetadataItem) {
	objc.Call[objc.Void](c_, objc.Sel("setMetadata:"), value)
}
