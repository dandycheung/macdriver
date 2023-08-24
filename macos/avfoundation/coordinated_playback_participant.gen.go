// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CoordinatedPlaybackParticipant] class.
var CoordinatedPlaybackParticipantClass = _CoordinatedPlaybackParticipantClass{objc.GetClass("AVCoordinatedPlaybackParticipant")}

type _CoordinatedPlaybackParticipantClass struct {
	objc.Class
}

// An interface definition for the [CoordinatedPlaybackParticipant] class.
type ICoordinatedPlaybackParticipant interface {
	objc.IObject
	IsReadyToPlay() bool
	SuspensionReasons() []CoordinatedPlaybackSuspensionReason
	Identifier() foundation.UUID
}

// An object that represents a participant in a coordinated playback session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybackparticipant?language=objc
type CoordinatedPlaybackParticipant struct {
	objc.Object
}

func CoordinatedPlaybackParticipantFrom(ptr unsafe.Pointer) CoordinatedPlaybackParticipant {
	return CoordinatedPlaybackParticipant{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CoordinatedPlaybackParticipantClass) Alloc() CoordinatedPlaybackParticipant {
	rv := objc.Call[CoordinatedPlaybackParticipant](cc, objc.Sel("alloc"))
	return rv
}

func CoordinatedPlaybackParticipant_Alloc() CoordinatedPlaybackParticipant {
	return CoordinatedPlaybackParticipantClass.Alloc()
}

func (cc _CoordinatedPlaybackParticipantClass) New() CoordinatedPlaybackParticipant {
	rv := objc.Call[CoordinatedPlaybackParticipant](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCoordinatedPlaybackParticipant() CoordinatedPlaybackParticipant {
	return CoordinatedPlaybackParticipantClass.New()
}

func (c_ CoordinatedPlaybackParticipant) Init() CoordinatedPlaybackParticipant {
	rv := objc.Call[CoordinatedPlaybackParticipant](c_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the participant is ready to start coordinated playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybackparticipant/3750236-readytoplay?language=objc
func (c_ CoordinatedPlaybackParticipant) IsReadyToPlay() bool {
	rv := objc.Call[bool](c_, objc.Sel("isReadyToPlay"))
	return rv
}

// The reasons a participant isn’t currently participating in coordinated playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybackparticipant/3750237-suspensionreasons?language=objc
func (c_ CoordinatedPlaybackParticipant) SuspensionReasons() []CoordinatedPlaybackSuspensionReason {
	rv := objc.Call[[]CoordinatedPlaybackSuspensionReason](c_, objc.Sel("suspensionReasons"))
	return rv
}

// A unique identifier for the participant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybackparticipant/3750235-identifier?language=objc
func (c_ CoordinatedPlaybackParticipant) Identifier() foundation.UUID {
	rv := objc.Call[foundation.UUID](c_, objc.Sel("identifier"))
	return rv
}