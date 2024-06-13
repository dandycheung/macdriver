// Code generated by DarwinKit. DO NOT EDIT.

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [HumanBodyPoseObservation] class.
var HumanBodyPoseObservationClass = _HumanBodyPoseObservationClass{objc.GetClass("VNHumanBodyPoseObservation")}

type _HumanBodyPoseObservationClass struct {
	objc.Class
}

// An interface definition for the [HumanBodyPoseObservation] class.
type IHumanBodyPoseObservation interface {
	IRecognizedPointsObservation
	RecognizedPointForJointNameError(jointName HumanBodyPoseObservationJointName, error unsafe.Pointer) RecognizedPoint
	RecognizedPointsForJointsGroupNameError(jointsGroupName HumanBodyPoseObservationJointsGroupName, error unsafe.Pointer) map[HumanBodyPoseObservationJointName]RecognizedPoint
	AvailableJointsGroupNames() []HumanBodyPoseObservationJointsGroupName
	AvailableJointNames() []HumanBodyPoseObservationJointName
}

// An observation that provides the body points the analysis recognized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservation?language=objc
type HumanBodyPoseObservation struct {
	RecognizedPointsObservation
}

func HumanBodyPoseObservationFrom(ptr unsafe.Pointer) HumanBodyPoseObservation {
	return HumanBodyPoseObservation{
		RecognizedPointsObservation: RecognizedPointsObservationFrom(ptr),
	}
}

func (hc _HumanBodyPoseObservationClass) Alloc() HumanBodyPoseObservation {
	rv := objc.Call[HumanBodyPoseObservation](hc, objc.Sel("alloc"))
	return rv
}

func (hc _HumanBodyPoseObservationClass) New() HumanBodyPoseObservation {
	rv := objc.Call[HumanBodyPoseObservation](hc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewHumanBodyPoseObservation() HumanBodyPoseObservation {
	return HumanBodyPoseObservationClass.New()
}

func (h_ HumanBodyPoseObservation) Init() HumanBodyPoseObservation {
	rv := objc.Call[HumanBodyPoseObservation](h_, objc.Sel("init"))
	return rv
}

// Retrieves the recognized point for a joint name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservation/3675603-recognizedpointforjointname?language=objc
func (h_ HumanBodyPoseObservation) RecognizedPointForJointNameError(jointName HumanBodyPoseObservationJointName, error unsafe.Pointer) RecognizedPoint {
	rv := objc.Call[RecognizedPoint](h_, objc.Sel("recognizedPointForJointName:error:"), jointName, error)
	return rv
}

// Retrieves the recognized points associated with the joint group name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservation/3675604-recognizedpointsforjointsgroupna?language=objc
func (h_ HumanBodyPoseObservation) RecognizedPointsForJointsGroupNameError(jointsGroupName HumanBodyPoseObservationJointsGroupName, error unsafe.Pointer) map[HumanBodyPoseObservationJointName]RecognizedPoint {
	rv := objc.Call[map[HumanBodyPoseObservationJointName]RecognizedPoint](h_, objc.Sel("recognizedPointsForJointsGroupName:error:"), jointsGroupName, error)
	return rv
}

// The available joint group names in the observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservation/3675602-availablejointsgroupnames?language=objc
func (h_ HumanBodyPoseObservation) AvailableJointsGroupNames() []HumanBodyPoseObservationJointsGroupName {
	rv := objc.Call[[]HumanBodyPoseObservationJointsGroupName](h_, objc.Sel("availableJointsGroupNames"))
	return rv
}

// The names of the available joints in the observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservation/3675601-availablejointnames?language=objc
func (h_ HumanBodyPoseObservation) AvailableJointNames() []HumanBodyPoseObservationJointName {
	rv := objc.Call[[]HumanBodyPoseObservationJointName](h_, objc.Sel("availableJointNames"))
	return rv
}
