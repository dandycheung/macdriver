// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods that a delegate of a split view implements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate?language=objc
type PSplitViewDelegate interface {
	// optional
	SplitViewDidResizeSubviews(notification foundation.Notification)
	HasSplitViewDidResizeSubviews() bool

	// optional
	SplitViewConstrainMaxCoordinateOfSubviewAt(splitView SplitView, proposedMaximumPosition float64, dividerIndex int) float64
	HasSplitViewConstrainMaxCoordinateOfSubviewAt() bool

	// optional
	SplitViewWillResizeSubviews(notification foundation.Notification)
	HasSplitViewWillResizeSubviews() bool

	// optional
	SplitViewConstrainSplitPositionOfSubviewAt(splitView SplitView, proposedPosition float64, dividerIndex int) float64
	HasSplitViewConstrainSplitPositionOfSubviewAt() bool

	// optional
	SplitViewShouldAdjustSizeOfSubview(splitView SplitView, view View) bool
	HasSplitViewShouldAdjustSizeOfSubview() bool

	// optional
	SplitViewResizeSubviewsWithOldSize(splitView SplitView, oldSize foundation.Size)
	HasSplitViewResizeSubviewsWithOldSize() bool

	// optional
	SplitViewCanCollapseSubview(splitView SplitView, subview View) bool
	HasSplitViewCanCollapseSubview() bool

	// optional
	SplitViewShouldHideDividerAtIndex(splitView SplitView, dividerIndex int) bool
	HasSplitViewShouldHideDividerAtIndex() bool

	// optional
	SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView SplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect
	HasSplitViewEffectiveRectForDrawnRectOfDividerAtIndex() bool

	// optional
	SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView SplitView, dividerIndex int) foundation.Rect
	HasSplitViewAdditionalEffectiveRectOfDividerAtIndex() bool

	// optional
	SplitViewConstrainMinCoordinateOfSubviewAt(splitView SplitView, proposedMinimumPosition float64, dividerIndex int) float64
	HasSplitViewConstrainMinCoordinateOfSubviewAt() bool
}

// A delegate implementation builder for the [PSplitViewDelegate] protocol.
type SplitViewDelegate struct {
	_SplitViewDidResizeSubviews                         func(notification foundation.Notification)
	_SplitViewConstrainMaxCoordinateOfSubviewAt         func(splitView SplitView, proposedMaximumPosition float64, dividerIndex int) float64
	_SplitViewWillResizeSubviews                        func(notification foundation.Notification)
	_SplitViewConstrainSplitPositionOfSubviewAt         func(splitView SplitView, proposedPosition float64, dividerIndex int) float64
	_SplitViewShouldAdjustSizeOfSubview                 func(splitView SplitView, view View) bool
	_SplitViewResizeSubviewsWithOldSize                 func(splitView SplitView, oldSize foundation.Size)
	_SplitViewCanCollapseSubview                        func(splitView SplitView, subview View) bool
	_SplitViewShouldHideDividerAtIndex                  func(splitView SplitView, dividerIndex int) bool
	_SplitViewEffectiveRectForDrawnRectOfDividerAtIndex func(splitView SplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect
	_SplitViewAdditionalEffectiveRectOfDividerAtIndex   func(splitView SplitView, dividerIndex int) foundation.Rect
	_SplitViewConstrainMinCoordinateOfSubviewAt         func(splitView SplitView, proposedMinimumPosition float64, dividerIndex int) float64
}

func (di *SplitViewDelegate) HasSplitViewDidResizeSubviews() bool {
	return di._SplitViewDidResizeSubviews != nil
}

// Notifies the delegate when the split view resizes its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455314-splitviewdidresizesubviews?language=objc
func (di *SplitViewDelegate) SetSplitViewDidResizeSubviews(f func(notification foundation.Notification)) {
	di._SplitViewDidResizeSubviews = f
}

// Notifies the delegate when the split view resizes its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455314-splitviewdidresizesubviews?language=objc
func (di *SplitViewDelegate) SplitViewDidResizeSubviews(notification foundation.Notification) {
	di._SplitViewDidResizeSubviews(notification)
}
func (di *SplitViewDelegate) HasSplitViewConstrainMaxCoordinateOfSubviewAt() bool {
	return di._SplitViewConstrainMaxCoordinateOfSubviewAt != nil
}

// Allows the delegate to constrain the maximum coordinate limit of a divider when the user drags it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455300-splitview?language=objc
func (di *SplitViewDelegate) SetSplitViewConstrainMaxCoordinateOfSubviewAt(f func(splitView SplitView, proposedMaximumPosition float64, dividerIndex int) float64) {
	di._SplitViewConstrainMaxCoordinateOfSubviewAt = f
}

// Allows the delegate to constrain the maximum coordinate limit of a divider when the user drags it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455300-splitview?language=objc
func (di *SplitViewDelegate) SplitViewConstrainMaxCoordinateOfSubviewAt(splitView SplitView, proposedMaximumPosition float64, dividerIndex int) float64 {
	return di._SplitViewConstrainMaxCoordinateOfSubviewAt(splitView, proposedMaximumPosition, dividerIndex)
}
func (di *SplitViewDelegate) HasSplitViewWillResizeSubviews() bool {
	return di._SplitViewWillResizeSubviews != nil
}

// Notifies the delegate when the split view is about to resize its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455289-splitviewwillresizesubviews?language=objc
func (di *SplitViewDelegate) SetSplitViewWillResizeSubviews(f func(notification foundation.Notification)) {
	di._SplitViewWillResizeSubviews = f
}

// Notifies the delegate when the split view is about to resize its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455289-splitviewwillresizesubviews?language=objc
func (di *SplitViewDelegate) SplitViewWillResizeSubviews(notification foundation.Notification) {
	di._SplitViewWillResizeSubviews(notification)
}
func (di *SplitViewDelegate) HasSplitViewConstrainSplitPositionOfSubviewAt() bool {
	return di._SplitViewConstrainSplitPositionOfSubviewAt != nil
}

// Allows the delegate to constrain the divider to certain positions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455312-splitview?language=objc
func (di *SplitViewDelegate) SetSplitViewConstrainSplitPositionOfSubviewAt(f func(splitView SplitView, proposedPosition float64, dividerIndex int) float64) {
	di._SplitViewConstrainSplitPositionOfSubviewAt = f
}

// Allows the delegate to constrain the divider to certain positions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455312-splitview?language=objc
func (di *SplitViewDelegate) SplitViewConstrainSplitPositionOfSubviewAt(splitView SplitView, proposedPosition float64, dividerIndex int) float64 {
	return di._SplitViewConstrainSplitPositionOfSubviewAt(splitView, proposedPosition, dividerIndex)
}
func (di *SplitViewDelegate) HasSplitViewShouldAdjustSizeOfSubview() bool {
	return di._SplitViewShouldAdjustSizeOfSubview != nil
}

// Allows the delegate to specify whether to resize the subview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455269-splitview?language=objc
func (di *SplitViewDelegate) SetSplitViewShouldAdjustSizeOfSubview(f func(splitView SplitView, view View) bool) {
	di._SplitViewShouldAdjustSizeOfSubview = f
}

// Allows the delegate to specify whether to resize the subview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455269-splitview?language=objc
func (di *SplitViewDelegate) SplitViewShouldAdjustSizeOfSubview(splitView SplitView, view View) bool {
	return di._SplitViewShouldAdjustSizeOfSubview(splitView, view)
}
func (di *SplitViewDelegate) HasSplitViewResizeSubviewsWithOldSize() bool {
	return di._SplitViewResizeSubviewsWithOldSize != nil
}

// Allows the delegate to specify custom sizing behavior for the subviews of the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455273-splitview?language=objc
func (di *SplitViewDelegate) SetSplitViewResizeSubviewsWithOldSize(f func(splitView SplitView, oldSize foundation.Size)) {
	di._SplitViewResizeSubviewsWithOldSize = f
}

// Allows the delegate to specify custom sizing behavior for the subviews of the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455273-splitview?language=objc
func (di *SplitViewDelegate) SplitViewResizeSubviewsWithOldSize(splitView SplitView, oldSize foundation.Size) {
	di._SplitViewResizeSubviewsWithOldSize(splitView, oldSize)
}
func (di *SplitViewDelegate) HasSplitViewCanCollapseSubview() bool {
	return di._SplitViewCanCollapseSubview != nil
}

// Allows the delegate to determine whether the user can collapse and expand the specified subview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455304-splitview?language=objc
func (di *SplitViewDelegate) SetSplitViewCanCollapseSubview(f func(splitView SplitView, subview View) bool) {
	di._SplitViewCanCollapseSubview = f
}

// Allows the delegate to determine whether the user can collapse and expand the specified subview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455304-splitview?language=objc
func (di *SplitViewDelegate) SplitViewCanCollapseSubview(splitView SplitView, subview View) bool {
	return di._SplitViewCanCollapseSubview(splitView, subview)
}
func (di *SplitViewDelegate) HasSplitViewShouldHideDividerAtIndex() bool {
	return di._SplitViewShouldHideDividerAtIndex != nil
}

// Allows the delegate to determine whether the user can drag a divider or adjust it off the edge of the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455280-splitview?language=objc
func (di *SplitViewDelegate) SetSplitViewShouldHideDividerAtIndex(f func(splitView SplitView, dividerIndex int) bool) {
	di._SplitViewShouldHideDividerAtIndex = f
}

// Allows the delegate to determine whether the user can drag a divider or adjust it off the edge of the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455280-splitview?language=objc
func (di *SplitViewDelegate) SplitViewShouldHideDividerAtIndex(splitView SplitView, dividerIndex int) bool {
	return di._SplitViewShouldHideDividerAtIndex(splitView, dividerIndex)
}
func (di *SplitViewDelegate) HasSplitViewEffectiveRectForDrawnRectOfDividerAtIndex() bool {
	return di._SplitViewEffectiveRectForDrawnRectOfDividerAtIndex != nil
}

// Allows the delegate to modify the rectangle where mouse clicks initiate divider dragging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455288-splitview?language=objc
func (di *SplitViewDelegate) SetSplitViewEffectiveRectForDrawnRectOfDividerAtIndex(f func(splitView SplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect) {
	di._SplitViewEffectiveRectForDrawnRectOfDividerAtIndex = f
}

// Allows the delegate to modify the rectangle where mouse clicks initiate divider dragging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455288-splitview?language=objc
func (di *SplitViewDelegate) SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView SplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect {
	return di._SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView, proposedEffectiveRect, drawnRect, dividerIndex)
}
func (di *SplitViewDelegate) HasSplitViewAdditionalEffectiveRectOfDividerAtIndex() bool {
	return di._SplitViewAdditionalEffectiveRectOfDividerAtIndex != nil
}

// Allows the delegate to return an additional rectangle where mouse clicks can initiate divider dragging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455292-splitview?language=objc
func (di *SplitViewDelegate) SetSplitViewAdditionalEffectiveRectOfDividerAtIndex(f func(splitView SplitView, dividerIndex int) foundation.Rect) {
	di._SplitViewAdditionalEffectiveRectOfDividerAtIndex = f
}

// Allows the delegate to return an additional rectangle where mouse clicks can initiate divider dragging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455292-splitview?language=objc
func (di *SplitViewDelegate) SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView SplitView, dividerIndex int) foundation.Rect {
	return di._SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView, dividerIndex)
}
func (di *SplitViewDelegate) HasSplitViewConstrainMinCoordinateOfSubviewAt() bool {
	return di._SplitViewConstrainMinCoordinateOfSubviewAt != nil
}

// Allows the delegate to constrain the minimum coordinate limit of a divider when the user drags it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455302-splitview?language=objc
func (di *SplitViewDelegate) SetSplitViewConstrainMinCoordinateOfSubviewAt(f func(splitView SplitView, proposedMinimumPosition float64, dividerIndex int) float64) {
	di._SplitViewConstrainMinCoordinateOfSubviewAt = f
}

// Allows the delegate to constrain the minimum coordinate limit of a divider when the user drags it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455302-splitview?language=objc
func (di *SplitViewDelegate) SplitViewConstrainMinCoordinateOfSubviewAt(splitView SplitView, proposedMinimumPosition float64, dividerIndex int) float64 {
	return di._SplitViewConstrainMinCoordinateOfSubviewAt(splitView, proposedMinimumPosition, dividerIndex)
}

// ensure impl type implements protocol interface
var _ PSplitViewDelegate = (*SplitViewDelegateObject)(nil)

// A concrete type for the [PSplitViewDelegate] protocol.
type SplitViewDelegateObject struct {
	objc.Object
}

func (s_ SplitViewDelegateObject) HasSplitViewDidResizeSubviews() bool {
	return s_.RespondsToSelector(objc.Sel("splitViewDidResizeSubviews:"))
}

// Notifies the delegate when the split view resizes its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455314-splitviewdidresizesubviews?language=objc
func (s_ SplitViewDelegateObject) SplitViewDidResizeSubviews(notification foundation.Notification) {
	objc.Call[objc.Void](s_, objc.Sel("splitViewDidResizeSubviews:"), notification)
}

func (s_ SplitViewDelegateObject) HasSplitViewConstrainMaxCoordinateOfSubviewAt() bool {
	return s_.RespondsToSelector(objc.Sel("splitView:constrainMaxCoordinate:ofSubviewAt:"))
}

// Allows the delegate to constrain the maximum coordinate limit of a divider when the user drags it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455300-splitview?language=objc
func (s_ SplitViewDelegateObject) SplitViewConstrainMaxCoordinateOfSubviewAt(splitView SplitView, proposedMaximumPosition float64, dividerIndex int) float64 {
	rv := objc.Call[float64](s_, objc.Sel("splitView:constrainMaxCoordinate:ofSubviewAt:"), splitView, proposedMaximumPosition, dividerIndex)
	return rv
}

func (s_ SplitViewDelegateObject) HasSplitViewWillResizeSubviews() bool {
	return s_.RespondsToSelector(objc.Sel("splitViewWillResizeSubviews:"))
}

// Notifies the delegate when the split view is about to resize its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455289-splitviewwillresizesubviews?language=objc
func (s_ SplitViewDelegateObject) SplitViewWillResizeSubviews(notification foundation.Notification) {
	objc.Call[objc.Void](s_, objc.Sel("splitViewWillResizeSubviews:"), notification)
}

func (s_ SplitViewDelegateObject) HasSplitViewConstrainSplitPositionOfSubviewAt() bool {
	return s_.RespondsToSelector(objc.Sel("splitView:constrainSplitPosition:ofSubviewAt:"))
}

// Allows the delegate to constrain the divider to certain positions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455312-splitview?language=objc
func (s_ SplitViewDelegateObject) SplitViewConstrainSplitPositionOfSubviewAt(splitView SplitView, proposedPosition float64, dividerIndex int) float64 {
	rv := objc.Call[float64](s_, objc.Sel("splitView:constrainSplitPosition:ofSubviewAt:"), splitView, proposedPosition, dividerIndex)
	return rv
}

func (s_ SplitViewDelegateObject) HasSplitViewShouldAdjustSizeOfSubview() bool {
	return s_.RespondsToSelector(objc.Sel("splitView:shouldAdjustSizeOfSubview:"))
}

// Allows the delegate to specify whether to resize the subview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455269-splitview?language=objc
func (s_ SplitViewDelegateObject) SplitViewShouldAdjustSizeOfSubview(splitView SplitView, view View) bool {
	rv := objc.Call[bool](s_, objc.Sel("splitView:shouldAdjustSizeOfSubview:"), splitView, view)
	return rv
}

func (s_ SplitViewDelegateObject) HasSplitViewResizeSubviewsWithOldSize() bool {
	return s_.RespondsToSelector(objc.Sel("splitView:resizeSubviewsWithOldSize:"))
}

// Allows the delegate to specify custom sizing behavior for the subviews of the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455273-splitview?language=objc
func (s_ SplitViewDelegateObject) SplitViewResizeSubviewsWithOldSize(splitView SplitView, oldSize foundation.Size) {
	objc.Call[objc.Void](s_, objc.Sel("splitView:resizeSubviewsWithOldSize:"), splitView, oldSize)
}

func (s_ SplitViewDelegateObject) HasSplitViewCanCollapseSubview() bool {
	return s_.RespondsToSelector(objc.Sel("splitView:canCollapseSubview:"))
}

// Allows the delegate to determine whether the user can collapse and expand the specified subview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455304-splitview?language=objc
func (s_ SplitViewDelegateObject) SplitViewCanCollapseSubview(splitView SplitView, subview View) bool {
	rv := objc.Call[bool](s_, objc.Sel("splitView:canCollapseSubview:"), splitView, subview)
	return rv
}

func (s_ SplitViewDelegateObject) HasSplitViewShouldHideDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.Sel("splitView:shouldHideDividerAtIndex:"))
}

// Allows the delegate to determine whether the user can drag a divider or adjust it off the edge of the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455280-splitview?language=objc
func (s_ SplitViewDelegateObject) SplitViewShouldHideDividerAtIndex(splitView SplitView, dividerIndex int) bool {
	rv := objc.Call[bool](s_, objc.Sel("splitView:shouldHideDividerAtIndex:"), splitView, dividerIndex)
	return rv
}

func (s_ SplitViewDelegateObject) HasSplitViewEffectiveRectForDrawnRectOfDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.Sel("splitView:effectiveRect:forDrawnRect:ofDividerAtIndex:"))
}

// Allows the delegate to modify the rectangle where mouse clicks initiate divider dragging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455288-splitview?language=objc
func (s_ SplitViewDelegateObject) SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView SplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("splitView:effectiveRect:forDrawnRect:ofDividerAtIndex:"), splitView, proposedEffectiveRect, drawnRect, dividerIndex)
	return rv
}

func (s_ SplitViewDelegateObject) HasSplitViewAdditionalEffectiveRectOfDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.Sel("splitView:additionalEffectiveRectOfDividerAtIndex:"))
}

// Allows the delegate to return an additional rectangle where mouse clicks can initiate divider dragging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455292-splitview?language=objc
func (s_ SplitViewDelegateObject) SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView SplitView, dividerIndex int) foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("splitView:additionalEffectiveRectOfDividerAtIndex:"), splitView, dividerIndex)
	return rv
}

func (s_ SplitViewDelegateObject) HasSplitViewConstrainMinCoordinateOfSubviewAt() bool {
	return s_.RespondsToSelector(objc.Sel("splitView:constrainMinCoordinate:ofSubviewAt:"))
}

// Allows the delegate to constrain the minimum coordinate limit of a divider when the user drags it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455302-splitview?language=objc
func (s_ SplitViewDelegateObject) SplitViewConstrainMinCoordinateOfSubviewAt(splitView SplitView, proposedMinimumPosition float64, dividerIndex int) float64 {
	rv := objc.Call[float64](s_, objc.Sel("splitView:constrainMinCoordinate:ofSubviewAt:"), splitView, proposedMinimumPosition, dividerIndex)
	return rv
}
