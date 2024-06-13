// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PageLayout] class.
var PageLayoutClass = _PageLayoutClass{objc.GetClass("NSPageLayout")}

type _PageLayoutClass struct {
	objc.Class
}

// An interface definition for the [PageLayout] class.
type IPageLayout interface {
	objc.IObject
	AddAccessoryController(accessoryController IViewController)
	RemoveAccessoryController(accessoryController IViewController)
	RunModalWithPrintInfo(printInfo IPrintInfo) int
	RunModal() int
	AccessoryControllers() []ViewController
	PrintInfo() PrintInfo
}

// A panel that queries the user for information such as paper type and orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagelayout?language=objc
type PageLayout struct {
	objc.Object
}

func PageLayoutFrom(ptr unsafe.Pointer) PageLayout {
	return PageLayout{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PageLayoutClass) Alloc() PageLayout {
	rv := objc.Call[PageLayout](pc, objc.Sel("alloc"))
	return rv
}

func (pc _PageLayoutClass) New() PageLayout {
	rv := objc.Call[PageLayout](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPageLayout() PageLayout {
	return PageLayoutClass.New()
}

func (p_ PageLayout) Init() PageLayout {
	rv := objc.Call[PageLayout](p_, objc.Sel("init"))
	return rv
}

// Adds the specified controller of an accessory view to be presented in the page setup panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagelayout/1397790-addaccessorycontroller?language=objc
func (p_ PageLayout) AddAccessoryController(accessoryController IViewController) {
	objc.Call[objc.Void](p_, objc.Sel("addAccessoryController:"), accessoryController)
}

// Removes the specified controller of an accessory view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagelayout/1397802-removeaccessorycontroller?language=objc
func (p_ PageLayout) RemoveAccessoryController(accessoryController IViewController) {
	objc.Call[objc.Void](p_, objc.Sel("removeAccessoryController:"), accessoryController)
}

// Displays the page layout panel and begins the modal loop using the specified print info object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagelayout/1397784-runmodalwithprintinfo?language=objc
func (p_ PageLayout) RunModalWithPrintInfo(printInfo IPrintInfo) int {
	rv := objc.Call[int](p_, objc.Sel("runModalWithPrintInfo:"), printInfo)
	return rv
}

// Returns a newly created page layout object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagelayout/1397806-pagelayout?language=objc
func (pc _PageLayoutClass) PageLayout() PageLayout {
	rv := objc.Call[PageLayout](pc, objc.Sel("pageLayout"))
	return rv
}

// Returns a newly created page layout object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagelayout/1397806-pagelayout?language=objc
func PageLayout_PageLayout() PageLayout {
	return PageLayoutClass.PageLayout()
}

// Displays the page layout panel and begins the modal loop using the shared print info object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagelayout/1397798-runmodal?language=objc
func (p_ PageLayout) RunModal() int {
	rv := objc.Call[int](p_, objc.Sel("runModal"))
	return rv
}

// An array of accessory view controllers belonging to the page layout panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagelayout/1397788-accessorycontrollers?language=objc
func (p_ PageLayout) AccessoryControllers() []ViewController {
	rv := objc.Call[[]ViewController](p_, objc.Sel("accessoryControllers"))
	return rv
}

// The printing information object used when the page layout panel is run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagelayout/1397804-printinfo?language=objc
func (p_ PageLayout) PrintInfo() PrintInfo {
	rv := objc.Call[PrintInfo](p_, objc.Sel("printInfo"))
	return rv
}
