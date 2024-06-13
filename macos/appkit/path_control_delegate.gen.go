// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that can be implemented by the delegate of a path control object to support dragging to and from the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate?language=objc
type PPathControlDelegate interface {
	// optional
	PathControlShouldDragItemWithPasteboard(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool
	HasPathControlShouldDragItemWithPasteboard() bool

	// optional
	PathControlShouldDragPathComponentCellWithPasteboard(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool
	HasPathControlShouldDragPathComponentCellWithPasteboard() bool

	// optional
	PathControlValidateDrop(pathControl PathControl, info DraggingInfoObject) DragOperation
	HasPathControlValidateDrop() bool

	// optional
	PathControlWillDisplayOpenPanel(pathControl PathControl, openPanel OpenPanel)
	HasPathControlWillDisplayOpenPanel() bool

	// optional
	PathControlAcceptDrop(pathControl PathControl, info DraggingInfoObject) bool
	HasPathControlAcceptDrop() bool

	// optional
	PathControlWillPopUpMenu(pathControl PathControl, menu Menu)
	HasPathControlWillPopUpMenu() bool
}

// A delegate implementation builder for the [PPathControlDelegate] protocol.
type PathControlDelegate struct {
	_PathControlShouldDragItemWithPasteboard              func(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool
	_PathControlShouldDragPathComponentCellWithPasteboard func(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool
	_PathControlValidateDrop                              func(pathControl PathControl, info DraggingInfoObject) DragOperation
	_PathControlWillDisplayOpenPanel                      func(pathControl PathControl, openPanel OpenPanel)
	_PathControlAcceptDrop                                func(pathControl PathControl, info DraggingInfoObject) bool
	_PathControlWillPopUpMenu                             func(pathControl PathControl, menu Menu)
}

func (di *PathControlDelegate) HasPathControlShouldDragItemWithPasteboard() bool {
	return di._PathControlShouldDragItemWithPasteboard != nil
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1526752-pathcontrol?language=objc
func (di *PathControlDelegate) SetPathControlShouldDragItemWithPasteboard(f func(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool) {
	di._PathControlShouldDragItemWithPasteboard = f
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1526752-pathcontrol?language=objc
func (di *PathControlDelegate) PathControlShouldDragItemWithPasteboard(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool {
	return di._PathControlShouldDragItemWithPasteboard(pathControl, pathItem, pasteboard)
}
func (di *PathControlDelegate) HasPathControlShouldDragPathComponentCellWithPasteboard() bool {
	return di._PathControlShouldDragPathComponentCellWithPasteboard != nil
}

// Implement this method to enable dragging from the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1533453-pathcontrol?language=objc
func (di *PathControlDelegate) SetPathControlShouldDragPathComponentCellWithPasteboard(f func(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool) {
	di._PathControlShouldDragPathComponentCellWithPasteboard = f
}

// Implement this method to enable dragging from the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1533453-pathcontrol?language=objc
func (di *PathControlDelegate) PathControlShouldDragPathComponentCellWithPasteboard(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool {
	return di._PathControlShouldDragPathComponentCellWithPasteboard(pathControl, pathComponentCell, pasteboard)
}
func (di *PathControlDelegate) HasPathControlValidateDrop() bool {
	return di._PathControlValidateDrop != nil
}

// Implement this method to enable dragging onto the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1528554-pathcontrol?language=objc
func (di *PathControlDelegate) SetPathControlValidateDrop(f func(pathControl PathControl, info DraggingInfoObject) DragOperation) {
	di._PathControlValidateDrop = f
}

// Implement this method to enable dragging onto the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1528554-pathcontrol?language=objc
func (di *PathControlDelegate) PathControlValidateDrop(pathControl PathControl, info DraggingInfoObject) DragOperation {
	return di._PathControlValidateDrop(pathControl, info)
}
func (di *PathControlDelegate) HasPathControlWillDisplayOpenPanel() bool {
	return di._PathControlWillDisplayOpenPanel != nil
}

// Implement this method to customize the Open panel shown by a pop-up–style path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1530012-pathcontrol?language=objc
func (di *PathControlDelegate) SetPathControlWillDisplayOpenPanel(f func(pathControl PathControl, openPanel OpenPanel)) {
	di._PathControlWillDisplayOpenPanel = f
}

// Implement this method to customize the Open panel shown by a pop-up–style path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1530012-pathcontrol?language=objc
func (di *PathControlDelegate) PathControlWillDisplayOpenPanel(pathControl PathControl, openPanel OpenPanel) {
	di._PathControlWillDisplayOpenPanel(pathControl, openPanel)
}
func (di *PathControlDelegate) HasPathControlAcceptDrop() bool {
	return di._PathControlAcceptDrop != nil
}

// Implement this method to accept previously validated contents dropped onto the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1528517-pathcontrol?language=objc
func (di *PathControlDelegate) SetPathControlAcceptDrop(f func(pathControl PathControl, info DraggingInfoObject) bool) {
	di._PathControlAcceptDrop = f
}

// Implement this method to accept previously validated contents dropped onto the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1528517-pathcontrol?language=objc
func (di *PathControlDelegate) PathControlAcceptDrop(pathControl PathControl, info DraggingInfoObject) bool {
	return di._PathControlAcceptDrop(pathControl, info)
}
func (di *PathControlDelegate) HasPathControlWillPopUpMenu() bool {
	return di._PathControlWillPopUpMenu != nil
}

// Implement this method to customize the menu of a pop-up–style path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1531724-pathcontrol?language=objc
func (di *PathControlDelegate) SetPathControlWillPopUpMenu(f func(pathControl PathControl, menu Menu)) {
	di._PathControlWillPopUpMenu = f
}

// Implement this method to customize the menu of a pop-up–style path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1531724-pathcontrol?language=objc
func (di *PathControlDelegate) PathControlWillPopUpMenu(pathControl PathControl, menu Menu) {
	di._PathControlWillPopUpMenu(pathControl, menu)
}

// ensure impl type implements protocol interface
var _ PPathControlDelegate = (*PathControlDelegateObject)(nil)

// A concrete type for the [PPathControlDelegate] protocol.
type PathControlDelegateObject struct {
	objc.Object
}

func (p_ PathControlDelegateObject) HasPathControlShouldDragItemWithPasteboard() bool {
	return p_.RespondsToSelector(objc.Sel("pathControl:shouldDragItem:withPasteboard:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1526752-pathcontrol?language=objc
func (p_ PathControlDelegateObject) PathControlShouldDragItemWithPasteboard(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool {
	rv := objc.Call[bool](p_, objc.Sel("pathControl:shouldDragItem:withPasteboard:"), pathControl, pathItem, pasteboard)
	return rv
}

func (p_ PathControlDelegateObject) HasPathControlShouldDragPathComponentCellWithPasteboard() bool {
	return p_.RespondsToSelector(objc.Sel("pathControl:shouldDragPathComponentCell:withPasteboard:"))
}

// Implement this method to enable dragging from the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1533453-pathcontrol?language=objc
func (p_ PathControlDelegateObject) PathControlShouldDragPathComponentCellWithPasteboard(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool {
	rv := objc.Call[bool](p_, objc.Sel("pathControl:shouldDragPathComponentCell:withPasteboard:"), pathControl, pathComponentCell, pasteboard)
	return rv
}

func (p_ PathControlDelegateObject) HasPathControlValidateDrop() bool {
	return p_.RespondsToSelector(objc.Sel("pathControl:validateDrop:"))
}

// Implement this method to enable dragging onto the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1528554-pathcontrol?language=objc
func (p_ PathControlDelegateObject) PathControlValidateDrop(pathControl PathControl, info DraggingInfoObject) DragOperation {
	po1 := objc.WrapAsProtocol("NSDraggingInfo", info)
	rv := objc.Call[DragOperation](p_, objc.Sel("pathControl:validateDrop:"), pathControl, po1)
	return rv
}

func (p_ PathControlDelegateObject) HasPathControlWillDisplayOpenPanel() bool {
	return p_.RespondsToSelector(objc.Sel("pathControl:willDisplayOpenPanel:"))
}

// Implement this method to customize the Open panel shown by a pop-up–style path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1530012-pathcontrol?language=objc
func (p_ PathControlDelegateObject) PathControlWillDisplayOpenPanel(pathControl PathControl, openPanel OpenPanel) {
	objc.Call[objc.Void](p_, objc.Sel("pathControl:willDisplayOpenPanel:"), pathControl, openPanel)
}

func (p_ PathControlDelegateObject) HasPathControlAcceptDrop() bool {
	return p_.RespondsToSelector(objc.Sel("pathControl:acceptDrop:"))
}

// Implement this method to accept previously validated contents dropped onto the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1528517-pathcontrol?language=objc
func (p_ PathControlDelegateObject) PathControlAcceptDrop(pathControl PathControl, info DraggingInfoObject) bool {
	po1 := objc.WrapAsProtocol("NSDraggingInfo", info)
	rv := objc.Call[bool](p_, objc.Sel("pathControl:acceptDrop:"), pathControl, po1)
	return rv
}

func (p_ PathControlDelegateObject) HasPathControlWillPopUpMenu() bool {
	return p_.RespondsToSelector(objc.Sel("pathControl:willPopUpMenu:"))
}

// Implement this method to customize the menu of a pop-up–style path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1531724-pathcontrol?language=objc
func (p_ PathControlDelegateObject) PathControlWillPopUpMenu(pathControl PathControl, menu Menu) {
	objc.Call[objc.Void](p_, objc.Sel("pathControl:willPopUpMenu:"), pathControl, menu)
}
