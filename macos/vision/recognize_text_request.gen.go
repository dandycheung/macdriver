// Code generated by DarwinKit. DO NOT EDIT.

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RecognizeTextRequest] class.
var RecognizeTextRequestClass = _RecognizeTextRequestClass{objc.GetClass("VNRecognizeTextRequest")}

type _RecognizeTextRequestClass struct {
	objc.Class
}

// An interface definition for the [RecognizeTextRequest] class.
type IRecognizeTextRequest interface {
	IImageBasedRequest
	SupportedRecognitionLanguagesAndReturnError(error unsafe.Pointer) []string
	RecognitionLanguages() []string
	SetRecognitionLanguages(value []string)
	RecognitionLevel() RequestTextRecognitionLevel
	SetRecognitionLevel(value RequestTextRecognitionLevel)
	MinimumTextHeight() float32
	SetMinimumTextHeight(value float32)
	CustomWords() []string
	SetCustomWords(value []string)
	UsesLanguageCorrection() bool
	SetUsesLanguageCorrection(value bool)
}

// An image analysis request that finds and recognizes text in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest?language=objc
type RecognizeTextRequest struct {
	ImageBasedRequest
}

func RecognizeTextRequestFrom(ptr unsafe.Pointer) RecognizeTextRequest {
	return RecognizeTextRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (rc _RecognizeTextRequestClass) Alloc() RecognizeTextRequest {
	rv := objc.Call[RecognizeTextRequest](rc, objc.Sel("alloc"))
	return rv
}

func (rc _RecognizeTextRequestClass) New() RecognizeTextRequest {
	rv := objc.Call[RecognizeTextRequest](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRecognizeTextRequest() RecognizeTextRequest {
	return RecognizeTextRequestClass.New()
}

func (r_ RecognizeTextRequest) Init() RecognizeTextRequest {
	rv := objc.Call[RecognizeTextRequest](r_, objc.Sel("init"))
	return rv
}

func (r_ RecognizeTextRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) RecognizeTextRequest {
	rv := objc.Call[RecognizeTextRequest](r_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewRecognizeTextRequestWithCompletionHandler(completionHandler RequestCompletionHandler) RecognizeTextRequest {
	instance := RecognizeTextRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// Returns the identifiers of the languages that the request supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/3751006-supportedrecognitionlanguagesand?language=objc
func (r_ RecognizeTextRequest) SupportedRecognitionLanguagesAndReturnError(error unsafe.Pointer) []string {
	rv := objc.Call[[]string](r_, objc.Sel("supportedRecognitionLanguagesAndReturnError:"), error)
	return rv
}

// An array of languages to detect, in priority order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/3152642-recognitionlanguages?language=objc
func (r_ RecognizeTextRequest) RecognitionLanguages() []string {
	rv := objc.Call[[]string](r_, objc.Sel("recognitionLanguages"))
	return rv
}

// An array of languages to detect, in priority order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/3152642-recognitionlanguages?language=objc
func (r_ RecognizeTextRequest) SetRecognitionLanguages(value []string) {
	objc.Call[objc.Void](r_, objc.Sel("setRecognitionLanguages:"), value)
}

// A value that determines whether the request prioritizes accuracy or speed in text recognition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/3152643-recognitionlevel?language=objc
func (r_ RecognizeTextRequest) RecognitionLevel() RequestTextRecognitionLevel {
	rv := objc.Call[RequestTextRecognitionLevel](r_, objc.Sel("recognitionLevel"))
	return rv
}

// A value that determines whether the request prioritizes accuracy or speed in text recognition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/3152643-recognitionlevel?language=objc
func (r_ RecognizeTextRequest) SetRecognitionLevel(value RequestTextRecognitionLevel) {
	objc.Call[objc.Void](r_, objc.Sel("setRecognitionLevel:"), value)
}

// The minimum height, relative to the image height, of the text to recognize. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/3152641-minimumtextheight?language=objc
func (r_ RecognizeTextRequest) MinimumTextHeight() float32 {
	rv := objc.Call[float32](r_, objc.Sel("minimumTextHeight"))
	return rv
}

// The minimum height, relative to the image height, of the text to recognize. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/3152641-minimumtextheight?language=objc
func (r_ RecognizeTextRequest) SetMinimumTextHeight(value float32) {
	objc.Call[objc.Void](r_, objc.Sel("setMinimumTextHeight:"), value)
}

// An array of strings to supplement the recognized languages at the word-recognition stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/3152640-customwords?language=objc
func (r_ RecognizeTextRequest) CustomWords() []string {
	rv := objc.Call[[]string](r_, objc.Sel("customWords"))
	return rv
}

// An array of strings to supplement the recognized languages at the word-recognition stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/3152640-customwords?language=objc
func (r_ RecognizeTextRequest) SetCustomWords(value []string) {
	objc.Call[objc.Void](r_, objc.Sel("setCustomWords:"), value)
}

// A Boolean value that indicates whether the request applies language correction during the recognition process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/3166773-useslanguagecorrection?language=objc
func (r_ RecognizeTextRequest) UsesLanguageCorrection() bool {
	rv := objc.Call[bool](r_, objc.Sel("usesLanguageCorrection"))
	return rv
}

// A Boolean value that indicates whether the request applies language correction during the recognition process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/3166773-useslanguagecorrection?language=objc
func (r_ RecognizeTextRequest) SetUsesLanguageCorrection(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setUsesLanguageCorrection:"), value)
}
