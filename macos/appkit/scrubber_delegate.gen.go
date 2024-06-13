// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that a scrubber delegate implements to respond to user interactions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate?language=objc
type PScrubberDelegate interface {
	// optional
	ScrubberDidHighlightItemAtIndex(scrubber Scrubber, highlightedIndex int)
	HasScrubberDidHighlightItemAtIndex() bool

	// optional
	DidBeginInteractingWithScrubber(scrubber Scrubber)
	HasDidBeginInteractingWithScrubber() bool

	// optional
	ScrubberDidSelectItemAtIndex(scrubber Scrubber, selectedIndex int)
	HasScrubberDidSelectItemAtIndex() bool

	// optional
	DidCancelInteractingWithScrubber(scrubber Scrubber)
	HasDidCancelInteractingWithScrubber() bool

	// optional
	ScrubberDidChangeVisibleRange(scrubber Scrubber, visibleRange foundation.Range)
	HasScrubberDidChangeVisibleRange() bool

	// optional
	DidFinishInteractingWithScrubber(scrubber Scrubber)
	HasDidFinishInteractingWithScrubber() bool
}

// A delegate implementation builder for the [PScrubberDelegate] protocol.
type ScrubberDelegate struct {
	_ScrubberDidHighlightItemAtIndex  func(scrubber Scrubber, highlightedIndex int)
	_DidBeginInteractingWithScrubber  func(scrubber Scrubber)
	_ScrubberDidSelectItemAtIndex     func(scrubber Scrubber, selectedIndex int)
	_DidCancelInteractingWithScrubber func(scrubber Scrubber)
	_ScrubberDidChangeVisibleRange    func(scrubber Scrubber, visibleRange foundation.Range)
	_DidFinishInteractingWithScrubber func(scrubber Scrubber)
}

func (di *ScrubberDelegate) HasScrubberDidHighlightItemAtIndex() bool {
	return di._ScrubberDidHighlightItemAtIndex != nil
}

// Tells the delegate that the item at the specified index was highlighted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544788-scrubber?language=objc
func (di *ScrubberDelegate) SetScrubberDidHighlightItemAtIndex(f func(scrubber Scrubber, highlightedIndex int)) {
	di._ScrubberDidHighlightItemAtIndex = f
}

// Tells the delegate that the item at the specified index was highlighted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544788-scrubber?language=objc
func (di *ScrubberDelegate) ScrubberDidHighlightItemAtIndex(scrubber Scrubber, highlightedIndex int) {
	di._ScrubberDidHighlightItemAtIndex(scrubber, highlightedIndex)
}
func (di *ScrubberDelegate) HasDidBeginInteractingWithScrubber() bool {
	return di._DidBeginInteractingWithScrubber != nil
}

// Tells the delegate that the user is panning or scrolling the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544657-didbegininteractingwithscrubber?language=objc
func (di *ScrubberDelegate) SetDidBeginInteractingWithScrubber(f func(scrubber Scrubber)) {
	di._DidBeginInteractingWithScrubber = f
}

// Tells the delegate that the user is panning or scrolling the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544657-didbegininteractingwithscrubber?language=objc
func (di *ScrubberDelegate) DidBeginInteractingWithScrubber(scrubber Scrubber) {
	di._DidBeginInteractingWithScrubber(scrubber)
}
func (di *ScrubberDelegate) HasScrubberDidSelectItemAtIndex() bool {
	return di._ScrubberDidSelectItemAtIndex != nil
}

// Tells the delegate that the item at the specified index was selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544714-scrubber?language=objc
func (di *ScrubberDelegate) SetScrubberDidSelectItemAtIndex(f func(scrubber Scrubber, selectedIndex int)) {
	di._ScrubberDidSelectItemAtIndex = f
}

// Tells the delegate that the item at the specified index was selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544714-scrubber?language=objc
func (di *ScrubberDelegate) ScrubberDidSelectItemAtIndex(scrubber Scrubber, selectedIndex int) {
	di._ScrubberDidSelectItemAtIndex(scrubber, selectedIndex)
}
func (di *ScrubberDelegate) HasDidCancelInteractingWithScrubber() bool {
	return di._DidCancelInteractingWithScrubber != nil
}

// Tells the delegate that a user interaction with the scrubber has been canceled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2646905-didcancelinteractingwithscrubber?language=objc
func (di *ScrubberDelegate) SetDidCancelInteractingWithScrubber(f func(scrubber Scrubber)) {
	di._DidCancelInteractingWithScrubber = f
}

// Tells the delegate that a user interaction with the scrubber has been canceled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2646905-didcancelinteractingwithscrubber?language=objc
func (di *ScrubberDelegate) DidCancelInteractingWithScrubber(scrubber Scrubber) {
	di._DidCancelInteractingWithScrubber(scrubber)
}
func (di *ScrubberDelegate) HasScrubberDidChangeVisibleRange() bool {
	return di._ScrubberDidChangeVisibleRange != nil
}

// Tells the delegate that the range of items currently visible in the scrubber has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544817-scrubber?language=objc
func (di *ScrubberDelegate) SetScrubberDidChangeVisibleRange(f func(scrubber Scrubber, visibleRange foundation.Range)) {
	di._ScrubberDidChangeVisibleRange = f
}

// Tells the delegate that the range of items currently visible in the scrubber has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544817-scrubber?language=objc
func (di *ScrubberDelegate) ScrubberDidChangeVisibleRange(scrubber Scrubber, visibleRange foundation.Range) {
	di._ScrubberDidChangeVisibleRange(scrubber, visibleRange)
}
func (di *ScrubberDelegate) HasDidFinishInteractingWithScrubber() bool {
	return di._DidFinishInteractingWithScrubber != nil
}

// Tells the delegate that a pan or scroll interaction with the scrubber has ended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544653-didfinishinteractingwithscrubber?language=objc
func (di *ScrubberDelegate) SetDidFinishInteractingWithScrubber(f func(scrubber Scrubber)) {
	di._DidFinishInteractingWithScrubber = f
}

// Tells the delegate that a pan or scroll interaction with the scrubber has ended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544653-didfinishinteractingwithscrubber?language=objc
func (di *ScrubberDelegate) DidFinishInteractingWithScrubber(scrubber Scrubber) {
	di._DidFinishInteractingWithScrubber(scrubber)
}

// ensure impl type implements protocol interface
var _ PScrubberDelegate = (*ScrubberDelegateObject)(nil)

// A concrete type for the [PScrubberDelegate] protocol.
type ScrubberDelegateObject struct {
	objc.Object
}

func (s_ ScrubberDelegateObject) HasScrubberDidHighlightItemAtIndex() bool {
	return s_.RespondsToSelector(objc.Sel("scrubber:didHighlightItemAtIndex:"))
}

// Tells the delegate that the item at the specified index was highlighted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544788-scrubber?language=objc
func (s_ ScrubberDelegateObject) ScrubberDidHighlightItemAtIndex(scrubber Scrubber, highlightedIndex int) {
	objc.Call[objc.Void](s_, objc.Sel("scrubber:didHighlightItemAtIndex:"), scrubber, highlightedIndex)
}

func (s_ ScrubberDelegateObject) HasDidBeginInteractingWithScrubber() bool {
	return s_.RespondsToSelector(objc.Sel("didBeginInteractingWithScrubber:"))
}

// Tells the delegate that the user is panning or scrolling the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544657-didbegininteractingwithscrubber?language=objc
func (s_ ScrubberDelegateObject) DidBeginInteractingWithScrubber(scrubber Scrubber) {
	objc.Call[objc.Void](s_, objc.Sel("didBeginInteractingWithScrubber:"), scrubber)
}

func (s_ ScrubberDelegateObject) HasScrubberDidSelectItemAtIndex() bool {
	return s_.RespondsToSelector(objc.Sel("scrubber:didSelectItemAtIndex:"))
}

// Tells the delegate that the item at the specified index was selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544714-scrubber?language=objc
func (s_ ScrubberDelegateObject) ScrubberDidSelectItemAtIndex(scrubber Scrubber, selectedIndex int) {
	objc.Call[objc.Void](s_, objc.Sel("scrubber:didSelectItemAtIndex:"), scrubber, selectedIndex)
}

func (s_ ScrubberDelegateObject) HasDidCancelInteractingWithScrubber() bool {
	return s_.RespondsToSelector(objc.Sel("didCancelInteractingWithScrubber:"))
}

// Tells the delegate that a user interaction with the scrubber has been canceled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2646905-didcancelinteractingwithscrubber?language=objc
func (s_ ScrubberDelegateObject) DidCancelInteractingWithScrubber(scrubber Scrubber) {
	objc.Call[objc.Void](s_, objc.Sel("didCancelInteractingWithScrubber:"), scrubber)
}

func (s_ ScrubberDelegateObject) HasScrubberDidChangeVisibleRange() bool {
	return s_.RespondsToSelector(objc.Sel("scrubber:didChangeVisibleRange:"))
}

// Tells the delegate that the range of items currently visible in the scrubber has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544817-scrubber?language=objc
func (s_ ScrubberDelegateObject) ScrubberDidChangeVisibleRange(scrubber Scrubber, visibleRange foundation.Range) {
	objc.Call[objc.Void](s_, objc.Sel("scrubber:didChangeVisibleRange:"), scrubber, visibleRange)
}

func (s_ ScrubberDelegateObject) HasDidFinishInteractingWithScrubber() bool {
	return s_.RespondsToSelector(objc.Sel("didFinishInteractingWithScrubber:"))
}

// Tells the delegate that a pan or scroll interaction with the scrubber has ended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544653-didfinishinteractingwithscrubber?language=objc
func (s_ ScrubberDelegateObject) DidFinishInteractingWithScrubber(scrubber Scrubber) {
	objc.Call[objc.Void](s_, objc.Sel("didFinishInteractingWithScrubber:"), scrubber)
}
