// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// An interface that provides the content to share from the macOS share sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertoolbaritemdelegate?language=objc
type PSharingServicePickerToolbarItemDelegate interface {
	// optional
	ItemsForSharingServicePickerToolbarItem(pickerToolbarItem SharingServicePickerToolbarItem) []objc.Object
	HasItemsForSharingServicePickerToolbarItem() bool
}

// A delegate implementation builder for the [PSharingServicePickerToolbarItemDelegate] protocol.
type SharingServicePickerToolbarItemDelegate struct {
	_ItemsForSharingServicePickerToolbarItem func(pickerToolbarItem SharingServicePickerToolbarItem) []objc.Object
}

func (di *SharingServicePickerToolbarItemDelegate) HasItemsForSharingServicePickerToolbarItem() bool {
	return di._ItemsForSharingServicePickerToolbarItem != nil
}

// Asks the delegate for the items to share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertoolbaritemdelegate/3365983-itemsforsharingservicepickertool?language=objc
func (di *SharingServicePickerToolbarItemDelegate) SetItemsForSharingServicePickerToolbarItem(f func(pickerToolbarItem SharingServicePickerToolbarItem) []objc.Object) {
	di._ItemsForSharingServicePickerToolbarItem = f
}

// Asks the delegate for the items to share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertoolbaritemdelegate/3365983-itemsforsharingservicepickertool?language=objc
func (di *SharingServicePickerToolbarItemDelegate) ItemsForSharingServicePickerToolbarItem(pickerToolbarItem SharingServicePickerToolbarItem) []objc.Object {
	return di._ItemsForSharingServicePickerToolbarItem(pickerToolbarItem)
}

// ensure impl type implements protocol interface
var _ PSharingServicePickerToolbarItemDelegate = (*SharingServicePickerToolbarItemDelegateObject)(nil)

// A concrete type for the [PSharingServicePickerToolbarItemDelegate] protocol.
type SharingServicePickerToolbarItemDelegateObject struct {
	objc.Object
}

func (s_ SharingServicePickerToolbarItemDelegateObject) HasItemsForSharingServicePickerToolbarItem() bool {
	return s_.RespondsToSelector(objc.Sel("itemsForSharingServicePickerToolbarItem:"))
}

// Asks the delegate for the items to share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertoolbaritemdelegate/3365983-itemsforsharingservicepickertool?language=objc
func (s_ SharingServicePickerToolbarItemDelegateObject) ItemsForSharingServicePickerToolbarItem(pickerToolbarItem SharingServicePickerToolbarItem) []objc.Object {
	rv := objc.Call[[]objc.Object](s_, objc.Sel("itemsForSharingServicePickerToolbarItem:"), pickerToolbarItem)
	return rv
}
