// Code generated by DarwinKit. DO NOT EDIT.

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TrackObjectRequest] class.
var TrackObjectRequestClass = _TrackObjectRequestClass{objc.GetClass("VNTrackObjectRequest")}

type _TrackObjectRequestClass struct {
	objc.Class
}

// An interface definition for the [TrackObjectRequest] class.
type ITrackObjectRequest interface {
	ITrackingRequest
}

// An image analysis request that tracks the movement of a previously identified object across multiple images or video frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackobjectrequest?language=objc
type TrackObjectRequest struct {
	TrackingRequest
}

func TrackObjectRequestFrom(ptr unsafe.Pointer) TrackObjectRequest {
	return TrackObjectRequest{
		TrackingRequest: TrackingRequestFrom(ptr),
	}
}

func (t_ TrackObjectRequest) InitWithDetectedObjectObservationCompletionHandler(observation IDetectedObjectObservation, completionHandler RequestCompletionHandler) TrackObjectRequest {
	rv := objc.Call[TrackObjectRequest](t_, objc.Sel("initWithDetectedObjectObservation:completionHandler:"), observation, completionHandler)
	return rv
}

// Creates a new object tracking request with a detected object observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackobjectrequest/2887295-initwithdetectedobjectobservatio?language=objc
func NewTrackObjectRequestWithDetectedObjectObservationCompletionHandler(observation IDetectedObjectObservation, completionHandler RequestCompletionHandler) TrackObjectRequest {
	instance := TrackObjectRequestClass.Alloc().InitWithDetectedObjectObservationCompletionHandler(observation, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TrackObjectRequest) InitWithDetectedObjectObservation(observation IDetectedObjectObservation) TrackObjectRequest {
	rv := objc.Call[TrackObjectRequest](t_, objc.Sel("initWithDetectedObjectObservation:"), observation)
	return rv
}

// Creates a new object tracking request with a detected object observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackobjectrequest/2887297-initwithdetectedobjectobservatio?language=objc
func NewTrackObjectRequestWithDetectedObjectObservation(observation IDetectedObjectObservation) TrackObjectRequest {
	instance := TrackObjectRequestClass.Alloc().InitWithDetectedObjectObservation(observation)
	instance.Autorelease()
	return instance
}

func (tc _TrackObjectRequestClass) Alloc() TrackObjectRequest {
	rv := objc.Call[TrackObjectRequest](tc, objc.Sel("alloc"))
	return rv
}

func (tc _TrackObjectRequestClass) New() TrackObjectRequest {
	rv := objc.Call[TrackObjectRequest](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTrackObjectRequest() TrackObjectRequest {
	return TrackObjectRequestClass.New()
}

func (t_ TrackObjectRequest) Init() TrackObjectRequest {
	rv := objc.Call[TrackObjectRequest](t_, objc.Sel("init"))
	return rv
}

func (t_ TrackObjectRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) TrackObjectRequest {
	rv := objc.Call[TrackObjectRequest](t_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewTrackObjectRequestWithCompletionHandler(completionHandler RequestCompletionHandler) TrackObjectRequest {
	instance := TrackObjectRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}
