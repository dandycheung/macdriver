// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// A set of optional methods implemented by delegates of NSAnimation objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate?language=objc
type PAnimationDelegate interface {
	// optional
	AnimationDidStop(animation Animation)
	HasAnimationDidStop() bool

	// optional
	AnimationValueForProgress(animation Animation, progress AnimationProgress) float32
	HasAnimationValueForProgress() bool

	// optional
	AnimationDidReachProgressMark(animation Animation, progress AnimationProgress)
	HasAnimationDidReachProgressMark() bool

	// optional
	AnimationDidEnd(animation Animation)
	HasAnimationDidEnd() bool

	// optional
	AnimationShouldStart(animation Animation) bool
	HasAnimationShouldStart() bool
}

// A delegate implementation builder for the [PAnimationDelegate] protocol.
type AnimationDelegate struct {
	_AnimationDidStop              func(animation Animation)
	_AnimationValueForProgress     func(animation Animation, progress AnimationProgress) float32
	_AnimationDidReachProgressMark func(animation Animation, progress AnimationProgress)
	_AnimationDidEnd               func(animation Animation)
	_AnimationShouldStart          func(animation Animation) bool
}

func (di *AnimationDelegate) HasAnimationDidStop() bool {
	return di._AnimationDidStop != nil
}

// Sent to the delegate when the specified animation is stopped before it completes its run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1534155-animationdidstop?language=objc
func (di *AnimationDelegate) SetAnimationDidStop(f func(animation Animation)) {
	di._AnimationDidStop = f
}

// Sent to the delegate when the specified animation is stopped before it completes its run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1534155-animationdidstop?language=objc
func (di *AnimationDelegate) AnimationDidStop(animation Animation) {
	di._AnimationDidStop(animation)
}
func (di *AnimationDelegate) HasAnimationValueForProgress() bool {
	return di._AnimationValueForProgress != nil
}

// Requests a custom curve value for the current progress value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1528965-animation?language=objc
func (di *AnimationDelegate) SetAnimationValueForProgress(f func(animation Animation, progress AnimationProgress) float32) {
	di._AnimationValueForProgress = f
}

// Requests a custom curve value for the current progress value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1528965-animation?language=objc
func (di *AnimationDelegate) AnimationValueForProgress(animation Animation, progress AnimationProgress) float32 {
	return di._AnimationValueForProgress(animation, progress)
}
func (di *AnimationDelegate) HasAnimationDidReachProgressMark() bool {
	return di._AnimationDidReachProgressMark != nil
}

// Sent to the delegate when an animation reaches a specific progress mark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1535100-animation?language=objc
func (di *AnimationDelegate) SetAnimationDidReachProgressMark(f func(animation Animation, progress AnimationProgress)) {
	di._AnimationDidReachProgressMark = f
}

// Sent to the delegate when an animation reaches a specific progress mark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1535100-animation?language=objc
func (di *AnimationDelegate) AnimationDidReachProgressMark(animation Animation, progress AnimationProgress) {
	di._AnimationDidReachProgressMark(animation, progress)
}
func (di *AnimationDelegate) HasAnimationDidEnd() bool {
	return di._AnimationDidEnd != nil
}

// Sent to the delegate when the specified animation completes its run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1535871-animationdidend?language=objc
func (di *AnimationDelegate) SetAnimationDidEnd(f func(animation Animation)) {
	di._AnimationDidEnd = f
}

// Sent to the delegate when the specified animation completes its run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1535871-animationdidend?language=objc
func (di *AnimationDelegate) AnimationDidEnd(animation Animation) {
	di._AnimationDidEnd(animation)
}
func (di *AnimationDelegate) HasAnimationShouldStart() bool {
	return di._AnimationShouldStart != nil
}

// Sent to the delegate just after an animation is started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1533279-animationshouldstart?language=objc
func (di *AnimationDelegate) SetAnimationShouldStart(f func(animation Animation) bool) {
	di._AnimationShouldStart = f
}

// Sent to the delegate just after an animation is started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1533279-animationshouldstart?language=objc
func (di *AnimationDelegate) AnimationShouldStart(animation Animation) bool {
	return di._AnimationShouldStart(animation)
}

// ensure impl type implements protocol interface
var _ PAnimationDelegate = (*AnimationDelegateObject)(nil)

// A concrete type for the [PAnimationDelegate] protocol.
type AnimationDelegateObject struct {
	objc.Object
}

func (a_ AnimationDelegateObject) HasAnimationDidStop() bool {
	return a_.RespondsToSelector(objc.Sel("animationDidStop:"))
}

// Sent to the delegate when the specified animation is stopped before it completes its run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1534155-animationdidstop?language=objc
func (a_ AnimationDelegateObject) AnimationDidStop(animation Animation) {
	objc.Call[objc.Void](a_, objc.Sel("animationDidStop:"), animation)
}

func (a_ AnimationDelegateObject) HasAnimationValueForProgress() bool {
	return a_.RespondsToSelector(objc.Sel("animation:valueForProgress:"))
}

// Requests a custom curve value for the current progress value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1528965-animation?language=objc
func (a_ AnimationDelegateObject) AnimationValueForProgress(animation Animation, progress AnimationProgress) float32 {
	rv := objc.Call[float32](a_, objc.Sel("animation:valueForProgress:"), animation, progress)
	return rv
}

func (a_ AnimationDelegateObject) HasAnimationDidReachProgressMark() bool {
	return a_.RespondsToSelector(objc.Sel("animation:didReachProgressMark:"))
}

// Sent to the delegate when an animation reaches a specific progress mark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1535100-animation?language=objc
func (a_ AnimationDelegateObject) AnimationDidReachProgressMark(animation Animation, progress AnimationProgress) {
	objc.Call[objc.Void](a_, objc.Sel("animation:didReachProgressMark:"), animation, progress)
}

func (a_ AnimationDelegateObject) HasAnimationDidEnd() bool {
	return a_.RespondsToSelector(objc.Sel("animationDidEnd:"))
}

// Sent to the delegate when the specified animation completes its run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1535871-animationdidend?language=objc
func (a_ AnimationDelegateObject) AnimationDidEnd(animation Animation) {
	objc.Call[objc.Void](a_, objc.Sel("animationDidEnd:"), animation)
}

func (a_ AnimationDelegateObject) HasAnimationShouldStart() bool {
	return a_.RespondsToSelector(objc.Sel("animationShouldStart:"))
}

// Sent to the delegate just after an animation is started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1533279-animationshouldstart?language=objc
func (a_ AnimationDelegateObject) AnimationShouldStart(animation Animation) bool {
	rv := objc.Call[bool](a_, objc.Sel("animationShouldStart:"), animation)
	return rv
}
