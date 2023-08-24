// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DelegatingPlaybackCoordinatorPauseCommand] class.
var DelegatingPlaybackCoordinatorPauseCommandClass = _DelegatingPlaybackCoordinatorPauseCommandClass{objc.GetClass("AVDelegatingPlaybackCoordinatorPauseCommand")}

type _DelegatingPlaybackCoordinatorPauseCommandClass struct {
	objc.Class
}

// An interface definition for the [DelegatingPlaybackCoordinatorPauseCommand] class.
type IDelegatingPlaybackCoordinatorPauseCommand interface {
	IDelegatingPlaybackCoordinatorPlaybackControlCommand
	ShouldBufferInAnticipationOfPlayback() bool
	AnticipatedPlaybackRate() float64
}

// A command that indicates to pause playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorpausecommand?language=objc
type DelegatingPlaybackCoordinatorPauseCommand struct {
	DelegatingPlaybackCoordinatorPlaybackControlCommand
}

func DelegatingPlaybackCoordinatorPauseCommandFrom(ptr unsafe.Pointer) DelegatingPlaybackCoordinatorPauseCommand {
	return DelegatingPlaybackCoordinatorPauseCommand{
		DelegatingPlaybackCoordinatorPlaybackControlCommand: DelegatingPlaybackCoordinatorPlaybackControlCommandFrom(ptr),
	}
}

func (dc _DelegatingPlaybackCoordinatorPauseCommandClass) Alloc() DelegatingPlaybackCoordinatorPauseCommand {
	rv := objc.Call[DelegatingPlaybackCoordinatorPauseCommand](dc, objc.Sel("alloc"))
	return rv
}

func DelegatingPlaybackCoordinatorPauseCommand_Alloc() DelegatingPlaybackCoordinatorPauseCommand {
	return DelegatingPlaybackCoordinatorPauseCommandClass.Alloc()
}

func (dc _DelegatingPlaybackCoordinatorPauseCommandClass) New() DelegatingPlaybackCoordinatorPauseCommand {
	rv := objc.Call[DelegatingPlaybackCoordinatorPauseCommand](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDelegatingPlaybackCoordinatorPauseCommand() DelegatingPlaybackCoordinatorPauseCommand {
	return DelegatingPlaybackCoordinatorPauseCommandClass.New()
}

func (d_ DelegatingPlaybackCoordinatorPauseCommand) Init() DelegatingPlaybackCoordinatorPauseCommand {
	rv := objc.Call[DelegatingPlaybackCoordinatorPauseCommand](d_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the player starts buffering in preparation for a request to begin playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorpausecommand/3750264-shouldbufferinanticipationofplay?language=objc
func (d_ DelegatingPlaybackCoordinatorPauseCommand) ShouldBufferInAnticipationOfPlayback() bool {
	rv := objc.Call[bool](d_, objc.Sel("shouldBufferInAnticipationOfPlayback"))
	return rv
}

// The rate at which the coordinator expects the current item to play. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorpausecommand/3750263-anticipatedplaybackrate?language=objc
func (d_ DelegatingPlaybackCoordinatorPauseCommand) AnticipatedPlaybackRate() float64 {
	rv := objc.Call[float64](d_, objc.Sel("anticipatedPlaybackRate"))
	return rv
}