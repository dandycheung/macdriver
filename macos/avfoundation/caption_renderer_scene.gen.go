// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptionRendererScene] class.
var CaptionRendererSceneClass = _CaptionRendererSceneClass{objc.GetClass("AVCaptionRendererScene")}

type _CaptionRendererSceneClass struct {
	objc.Class
}

// An interface definition for the [CaptionRendererScene] class.
type ICaptionRendererScene interface {
	objc.IObject
	HasActiveCaptions() bool
	TimeRange() coremedia.TimeRange
	NeedsPeriodicRefresh() bool
}

// An object that holds a time range and an associated state which indicates when the renderer draws output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionrendererscene?language=objc
type CaptionRendererScene struct {
	objc.Object
}

func CaptionRendererSceneFrom(ptr unsafe.Pointer) CaptionRendererScene {
	return CaptionRendererScene{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptionRendererSceneClass) Alloc() CaptionRendererScene {
	rv := objc.Call[CaptionRendererScene](cc, objc.Sel("alloc"))
	return rv
}

func CaptionRendererScene_Alloc() CaptionRendererScene {
	return CaptionRendererSceneClass.Alloc()
}

func (cc _CaptionRendererSceneClass) New() CaptionRendererScene {
	rv := objc.Call[CaptionRendererScene](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptionRendererScene() CaptionRendererScene {
	return CaptionRendererSceneClass.New()
}

func (c_ CaptionRendererScene) Init() CaptionRendererScene {
	rv := objc.Call[CaptionRendererScene](c_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the scene contains one or more active captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionrendererscene/3752973-hasactivecaptions?language=objc
func (c_ CaptionRendererScene) HasActiveCaptions() bool {
	rv := objc.Call[bool](c_, objc.Sel("hasActiveCaptions"))
	return rv
}

// The time range during which the system doesn’t modify the scene. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionrendererscene/3752975-timerange?language=objc
func (c_ CaptionRendererScene) TimeRange() coremedia.TimeRange {
	rv := objc.Call[coremedia.TimeRange](c_, objc.Sel("timeRange"))
	return rv
}

// A Boolean value that indicates whether the scene requires redrawing while your app progresses through the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionrendererscene/3752974-needsperiodicrefresh?language=objc
func (c_ CaptionRendererScene) NeedsPeriodicRefresh() bool {
	rv := objc.Call[bool](c_, objc.Sel("needsPeriodicRefresh"))
	return rv
}