// AUTO-GENERATED CODE, DO NOT MODIFY

package contactsui

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContactPicker] class.
var ContactPickerClass = _ContactPickerClass{objc.GetClass("CNContactPicker")}

type _ContactPickerClass struct {
	objc.Class
}

// An interface definition for the [ContactPicker] class.
type IContactPicker interface {
	objc.IObject
	Close()
	ShowRelativeToRectOfViewPreferredEdge(positioningRect foundation.Rect, positioningView appkit.IView, preferredEdge foundation.RectEdge)
	Delegate() ContactPickerDelegateWrapper
	SetDelegate(value PContactPickerDelegate)
	SetDelegateObject(valueObject objc.IObject)
	DisplayedKeys() []string
	SetDisplayedKeys(value []string)
}

// A popover-based interface for selecting a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpicker?language=objc
type ContactPicker struct {
	objc.Object
}

func ContactPickerFrom(ptr unsafe.Pointer) ContactPicker {
	return ContactPicker{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContactPickerClass) Alloc() ContactPicker {
	rv := objc.Call[ContactPicker](cc, objc.Sel("alloc"))
	return rv
}

func ContactPicker_Alloc() ContactPicker {
	return ContactPickerClass.Alloc()
}

func (cc _ContactPickerClass) New() ContactPicker {
	rv := objc.Call[ContactPicker](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContactPicker() ContactPicker {
	return ContactPickerClass.New()
}

func (c_ ContactPicker) Init() ContactPicker {
	rv := objc.Call[ContactPicker](c_, objc.Sel("init"))
	return rv
}

// Closes the popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpicker/1522592-close?language=objc
func (c_ ContactPicker) Close() {
	objc.Call[objc.Void](c_, objc.Sel("close"))
}

// Shows the picker popover anchored to the specified view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpicker/1522591-showrelativetorect?language=objc
func (c_ ContactPicker) ShowRelativeToRectOfViewPreferredEdge(positioningRect foundation.Rect, positioningView appkit.IView, preferredEdge foundation.RectEdge) {
	objc.Call[objc.Void](c_, objc.Sel("showRelativeToRect:ofView:preferredEdge:"), positioningRect, objc.Ptr(positioningView), preferredEdge)
}

// The picker delegate to be notified when the user chooses a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpicker/1522588-delegate?language=objc
func (c_ ContactPicker) Delegate() ContactPickerDelegateWrapper {
	rv := objc.Call[ContactPickerDelegateWrapper](c_, objc.Sel("delegate"))
	return rv
}

// The picker delegate to be notified when the user chooses a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpicker/1522588-delegate?language=objc
func (c_ ContactPicker) SetDelegate(value PContactPickerDelegate) {
	po0 := objc.WrapAsProtocol("CNContactPickerDelegate", value)
	objc.SetAssociatedObject(c_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:"), po0)
}

// The picker delegate to be notified when the user chooses a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpicker/1522588-delegate?language=objc
func (c_ ContactPicker) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The keys to be displayed when a contact is expanded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpicker/1522585-displayedkeys?language=objc
func (c_ ContactPicker) DisplayedKeys() []string {
	rv := objc.Call[[]string](c_, objc.Sel("displayedKeys"))
	return rv
}

// The keys to be displayed when a contact is expanded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpicker/1522585-displayedkeys?language=objc
func (c_ ContactPicker) SetDisplayedKeys(value []string) {
	objc.Call[objc.Void](c_, objc.Sel("setDisplayedKeys:"), value)
}