// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WindowTabGroup] class.
var WindowTabGroupClass = _WindowTabGroupClass{objc.GetClass("NSWindowTabGroup")}

type _WindowTabGroupClass struct {
	objc.Class
}

// An interface definition for the [WindowTabGroup] class.
type IWindowTabGroup interface {
	objc.IObject
	RemoveWindow(window IWindow)
	InsertWindowAtIndex(window IWindow, index int)
	AddWindow(window IWindow)
	Windows() []Window
	Identifier() WindowTabbingIdentifier
	IsTabBarVisible() bool
	IsOverviewVisible() bool
	SetOverviewVisible(value bool)
	SelectedWindow() Window
	SetSelectedWindow(value IWindow)
}

// A group of windows that display together as a single tabbed window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup?language=objc
type WindowTabGroup struct {
	objc.Object
}

func WindowTabGroupFrom(ptr unsafe.Pointer) WindowTabGroup {
	return WindowTabGroup{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WindowTabGroupClass) Alloc() WindowTabGroup {
	rv := objc.Call[WindowTabGroup](wc, objc.Sel("alloc"))
	return rv
}

func (wc _WindowTabGroupClass) New() WindowTabGroup {
	rv := objc.Call[WindowTabGroup](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWindowTabGroup() WindowTabGroup {
	return WindowTabGroupClass.New()
}

func (w_ WindowTabGroup) Init() WindowTabGroup {
	rv := objc.Call[WindowTabGroup](w_, objc.Sel("init"))
	return rv
}

// Removes a window from the tab group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/2879459-removewindow?language=objc
func (w_ WindowTabGroup) RemoveWindow(window IWindow) {
	objc.Call[objc.Void](w_, objc.Sel("removeWindow:"), window)
}

// Inserts a window at a specific location within the tab group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/2879455-insertwindow?language=objc
func (w_ WindowTabGroup) InsertWindowAtIndex(window IWindow, index int) {
	objc.Call[objc.Void](w_, objc.Sel("insertWindow:atIndex:"), window, index)
}

// Adds a window to the tab group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/2879450-addwindow?language=objc
func (w_ WindowTabGroup) AddWindow(window IWindow) {
	objc.Call[objc.Void](w_, objc.Sel("addWindow:"), window)
}

// A collection of the windows that are currently grouped together by this window tab group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/2879458-windows?language=objc
func (w_ WindowTabGroup) Windows() []Window {
	rv := objc.Call[[]Window](w_, objc.Sel("windows"))
	return rv
}

// The unique identifier for a tabbed window group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/2879445-identifier?language=objc
func (w_ WindowTabGroup) Identifier() WindowTabbingIdentifier {
	rv := objc.Call[WindowTabbingIdentifier](w_, objc.Sel("identifier"))
	return rv
}

// A Boolean value indicating whether the tabbed window group currently displays a tab bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/2879451-tabbarvisible?language=objc
func (w_ WindowTabGroup) IsTabBarVisible() bool {
	rv := objc.Call[bool](w_, objc.Sel("isTabBarVisible"))
	return rv
}

// A Boolean value indicating if the tab overview is currently displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/2879448-overviewvisible?language=objc
func (w_ WindowTabGroup) IsOverviewVisible() bool {
	rv := objc.Call[bool](w_, objc.Sel("isOverviewVisible"))
	return rv
}

// A Boolean value indicating if the tab overview is currently displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/2879448-overviewvisible?language=objc
func (w_ WindowTabGroup) SetOverviewVisible(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setOverviewVisible:"), value)
}

// The selected, or frontmost, window in the tab group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/2879457-selectedwindow?language=objc
func (w_ WindowTabGroup) SelectedWindow() Window {
	rv := objc.Call[Window](w_, objc.Sel("selectedWindow"))
	return rv
}

// The selected, or frontmost, window in the tab group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/2879457-selectedwindow?language=objc
func (w_ WindowTabGroup) SetSelectedWindow(value IWindow) {
	objc.Call[objc.Void](w_, objc.Sel("setSelectedWindow:"), value)
}
