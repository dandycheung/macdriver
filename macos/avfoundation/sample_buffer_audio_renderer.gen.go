// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SampleBufferAudioRenderer] class.
var SampleBufferAudioRendererClass = _SampleBufferAudioRendererClass{objc.GetClass("AVSampleBufferAudioRenderer")}

type _SampleBufferAudioRendererClass struct {
	objc.Class
}

// An interface definition for the [SampleBufferAudioRenderer] class.
type ISampleBufferAudioRenderer interface {
	objc.IObject
	FlushFromSourceTimeCompletionHandler(time coremedia.Time, completionHandler func(flushSucceeded bool))
	AllowedAudioSpatializationFormats() AudioSpatializationFormats
	SetAllowedAudioSpatializationFormats(value AudioSpatializationFormats)
	Volume() float64
	SetVolume(value float64)
	Error() foundation.Error
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm)
	IsMuted() bool
	SetMuted(value bool)
	AudioOutputDeviceUniqueID() string
	SetAudioOutputDeviceUniqueID(value string)
	Status() QueuedSampleBufferRenderingStatus
}

// An object used to decompress audio and play compressed or uncompressed audio. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer?language=objc
type SampleBufferAudioRenderer struct {
	objc.Object
}

func SampleBufferAudioRendererFrom(ptr unsafe.Pointer) SampleBufferAudioRenderer {
	return SampleBufferAudioRenderer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SampleBufferAudioRendererClass) Alloc() SampleBufferAudioRenderer {
	rv := objc.Call[SampleBufferAudioRenderer](sc, objc.Sel("alloc"))
	return rv
}

func SampleBufferAudioRenderer_Alloc() SampleBufferAudioRenderer {
	return SampleBufferAudioRendererClass.Alloc()
}

func (sc _SampleBufferAudioRendererClass) New() SampleBufferAudioRenderer {
	rv := objc.Call[SampleBufferAudioRenderer](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSampleBufferAudioRenderer() SampleBufferAudioRenderer {
	return SampleBufferAudioRendererClass.New()
}

func (s_ SampleBufferAudioRenderer) Init() SampleBufferAudioRenderer {
	rv := objc.Call[SampleBufferAudioRenderer](s_, objc.Sel("init"))
	return rv
}

// Flushes queued sample buffers with presentation time stamps later than or equal to the specified time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/2866181-flushfromsourcetime?language=objc
func (s_ SampleBufferAudioRenderer) FlushFromSourceTimeCompletionHandler(time coremedia.Time, completionHandler func(flushSucceeded bool)) {
	objc.Call[objc.Void](s_, objc.Sel("flushFromSourceTime:completionHandler:"), time, completionHandler)
}

// The source audio channel layouts the audio renderer supports for spatialization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/3750310-allowedaudiospatializationformat?language=objc
func (s_ SampleBufferAudioRenderer) AllowedAudioSpatializationFormats() AudioSpatializationFormats {
	rv := objc.Call[AudioSpatializationFormats](s_, objc.Sel("allowedAudioSpatializationFormats"))
	return rv
}

// The source audio channel layouts the audio renderer supports for spatialization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/3750310-allowedaudiospatializationformat?language=objc
func (s_ SampleBufferAudioRenderer) SetAllowedAudioSpatializationFormats(value AudioSpatializationFormats) {
	objc.Call[objc.Void](s_, objc.Sel("setAllowedAudioSpatializationFormats:"), value)
}

// The current audio volume for the audio renderer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/2866179-volume?language=objc
func (s_ SampleBufferAudioRenderer) Volume() float64 {
	rv := objc.Call[float64](s_, objc.Sel("volume"))
	return rv
}

// The current audio volume for the audio renderer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/2866179-volume?language=objc
func (s_ SampleBufferAudioRenderer) SetVolume(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setVolume:"), value)
}

// The error that caused the renderer to no longer render sample buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/2866178-error?language=objc
func (s_ SampleBufferAudioRenderer) Error() foundation.Error {
	rv := objc.Call[foundation.Error](s_, objc.Sel("error"))
	return rv
}

// The processing algorithm used to manage audio pitch at different rates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/2866180-audiotimepitchalgorithm?language=objc
func (s_ SampleBufferAudioRenderer) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm {
	rv := objc.Call[AudioTimePitchAlgorithm](s_, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}

// The processing algorithm used to manage audio pitch at different rates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/2866180-audiotimepitchalgorithm?language=objc
func (s_ SampleBufferAudioRenderer) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm) {
	objc.Call[objc.Void](s_, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}

// A Boolean value that indicates whether audio for the renderer is in a muted state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/2866177-muted?language=objc
func (s_ SampleBufferAudioRenderer) IsMuted() bool {
	rv := objc.Call[bool](s_, objc.Sel("isMuted"))
	return rv
}

// A Boolean value that indicates whether audio for the renderer is in a muted state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/2866177-muted?language=objc
func (s_ SampleBufferAudioRenderer) SetMuted(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setMuted:"), value)
}

// The unique identifier of the output device used to play audio. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/2866182-audiooutputdeviceuniqueid?language=objc
func (s_ SampleBufferAudioRenderer) AudioOutputDeviceUniqueID() string {
	rv := objc.Call[string](s_, objc.Sel("audioOutputDeviceUniqueID"))
	return rv
}

// The unique identifier of the output device used to play audio. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/2866182-audiooutputdeviceuniqueid?language=objc
func (s_ SampleBufferAudioRenderer) SetAudioOutputDeviceUniqueID(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setAudioOutputDeviceUniqueID:"), value)
}

// The status of the audio renderer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/2866183-status?language=objc
func (s_ SampleBufferAudioRenderer) Status() QueuedSampleBufferRenderingStatus {
	rv := objc.Call[QueuedSampleBufferRenderingStatus](s_, objc.Sel("status"))
	return rv
}