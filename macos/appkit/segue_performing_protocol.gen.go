// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that support the mediation of a custom segue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegueperforming?language=objc
type PSeguePerforming interface {
	// optional
	PerformSegueWithIdentifierSender(identifier StoryboardSegueIdentifier, sender objc.Object)
	HasPerformSegueWithIdentifierSender() bool

	// optional
	PrepareForSegueSender(segue StoryboardSegue, sender objc.Object)
	HasPrepareForSegueSender() bool

	// optional
	ShouldPerformSegueWithIdentifierSender(identifier StoryboardSegueIdentifier, sender objc.Object) bool
	HasShouldPerformSegueWithIdentifierSender() bool
}

// ensure impl type implements protocol interface
var _ PSeguePerforming = (*SeguePerformingObject)(nil)

// A concrete type for the [PSeguePerforming] protocol.
type SeguePerformingObject struct {
	objc.Object
}

func (s_ SeguePerformingObject) HasPerformSegueWithIdentifierSender() bool {
	return s_.RespondsToSelector(objc.Sel("performSegueWithIdentifier:sender:"))
}

// Performs the specified segue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegueperforming/1409583-performseguewithidentifier?language=objc
func (s_ SeguePerformingObject) PerformSegueWithIdentifierSender(identifier StoryboardSegueIdentifier, sender objc.Object) {
	objc.Call[objc.Void](s_, objc.Sel("performSegueWithIdentifier:sender:"), identifier, sender)
}

func (s_ SeguePerformingObject) HasPrepareForSegueSender() bool {
	return s_.RespondsToSelector(objc.Sel("prepareForSegue:sender:"))
}

// Called when a segue is about to be performed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegueperforming/1409580-prepareforsegue?language=objc
func (s_ SeguePerformingObject) PrepareForSegueSender(segue StoryboardSegue, sender objc.Object) {
	objc.Call[objc.Void](s_, objc.Sel("prepareForSegue:sender:"), segue, sender)
}

func (s_ SeguePerformingObject) HasShouldPerformSegueWithIdentifierSender() bool {
	return s_.RespondsToSelector(objc.Sel("shouldPerformSegueWithIdentifier:sender:"))
}

// Called immediately prior to the performance of a storyboard segue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegueperforming/1409574-shouldperformseguewithidentifier?language=objc
func (s_ SeguePerformingObject) ShouldPerformSegueWithIdentifierSender(identifier StoryboardSegueIdentifier, sender objc.Object) bool {
	rv := objc.Call[bool](s_, objc.Sel("shouldPerformSegueWithIdentifier:sender:"), identifier, sender)
	return rv
}
