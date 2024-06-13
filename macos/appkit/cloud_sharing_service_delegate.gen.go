// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods for responding to the life cycle events of the cloud-sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate?language=objc
type PCloudSharingServiceDelegate interface {
	// optional
	OptionsForSharingServiceShareProvider(cloudKitSharingService SharingService, provider foundation.ItemProvider) CloudKitSharingServiceOptions
	HasOptionsForSharingServiceShareProvider() bool

	// optional
	SharingServiceDidStopSharing(sharingService SharingService, share objc.Object)
	HasSharingServiceDidStopSharing() bool

	// optional
	SharingServiceDidCompleteForItemsError(sharingService SharingService, items []objc.Object, error foundation.Error)
	HasSharingServiceDidCompleteForItemsError() bool

	// optional
	SharingServiceDidSaveShare(sharingService SharingService, share objc.Object)
	HasSharingServiceDidSaveShare() bool
}

// A delegate implementation builder for the [PCloudSharingServiceDelegate] protocol.
type CloudSharingServiceDelegate struct {
	_OptionsForSharingServiceShareProvider  func(cloudKitSharingService SharingService, provider foundation.ItemProvider) CloudKitSharingServiceOptions
	_SharingServiceDidStopSharing           func(sharingService SharingService, share objc.Object)
	_SharingServiceDidCompleteForItemsError func(sharingService SharingService, items []objc.Object, error foundation.Error)
	_SharingServiceDidSaveShare             func(sharingService SharingService, share objc.Object)
}

func (di *CloudSharingServiceDelegate) HasOptionsForSharingServiceShareProvider() bool {
	return di._OptionsForSharingServiceShareProvider != nil
}

// Asks the delegate for the participant options for the cloud-sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644694-optionsforsharingservice?language=objc
func (di *CloudSharingServiceDelegate) SetOptionsForSharingServiceShareProvider(f func(cloudKitSharingService SharingService, provider foundation.ItemProvider) CloudKitSharingServiceOptions) {
	di._OptionsForSharingServiceShareProvider = f
}

// Asks the delegate for the participant options for the cloud-sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644694-optionsforsharingservice?language=objc
func (di *CloudSharingServiceDelegate) OptionsForSharingServiceShareProvider(cloudKitSharingService SharingService, provider foundation.ItemProvider) CloudKitSharingServiceOptions {
	return di._OptionsForSharingServiceShareProvider(cloudKitSharingService, provider)
}
func (di *CloudSharingServiceDelegate) HasSharingServiceDidStopSharing() bool {
	return di._SharingServiceDidStopSharing != nil
}

// Tells the delegate when the user stops sharing the CloudKit share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644709-sharingservice?language=objc
func (di *CloudSharingServiceDelegate) SetSharingServiceDidStopSharing(f func(sharingService SharingService, share objc.Object)) {
	di._SharingServiceDidStopSharing = f
}

// Tells the delegate when the user stops sharing the CloudKit share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644709-sharingservice?language=objc
func (di *CloudSharingServiceDelegate) SharingServiceDidStopSharing(sharingService SharingService, share objc.Object) {
	di._SharingServiceDidStopSharing(sharingService, share)
}
func (di *CloudSharingServiceDelegate) HasSharingServiceDidCompleteForItemsError() bool {
	return di._SharingServiceDidCompleteForItemsError != nil
}

// Tells the delegate when the cloud-sharing service completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644666-sharingservice?language=objc
func (di *CloudSharingServiceDelegate) SetSharingServiceDidCompleteForItemsError(f func(sharingService SharingService, items []objc.Object, error foundation.Error)) {
	di._SharingServiceDidCompleteForItemsError = f
}

// Tells the delegate when the cloud-sharing service completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644666-sharingservice?language=objc
func (di *CloudSharingServiceDelegate) SharingServiceDidCompleteForItemsError(sharingService SharingService, items []objc.Object, error foundation.Error) {
	di._SharingServiceDidCompleteForItemsError(sharingService, items, error)
}
func (di *CloudSharingServiceDelegate) HasSharingServiceDidSaveShare() bool {
	return di._SharingServiceDidSaveShare != nil
}

// Tells the delegate when the cloud-sharing service saves the CloudKit share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644712-sharingservice?language=objc
func (di *CloudSharingServiceDelegate) SetSharingServiceDidSaveShare(f func(sharingService SharingService, share objc.Object)) {
	di._SharingServiceDidSaveShare = f
}

// Tells the delegate when the cloud-sharing service saves the CloudKit share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644712-sharingservice?language=objc
func (di *CloudSharingServiceDelegate) SharingServiceDidSaveShare(sharingService SharingService, share objc.Object) {
	di._SharingServiceDidSaveShare(sharingService, share)
}

// ensure impl type implements protocol interface
var _ PCloudSharingServiceDelegate = (*CloudSharingServiceDelegateObject)(nil)

// A concrete type for the [PCloudSharingServiceDelegate] protocol.
type CloudSharingServiceDelegateObject struct {
	objc.Object
}

func (c_ CloudSharingServiceDelegateObject) HasOptionsForSharingServiceShareProvider() bool {
	return c_.RespondsToSelector(objc.Sel("optionsForSharingService:shareProvider:"))
}

// Asks the delegate for the participant options for the cloud-sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644694-optionsforsharingservice?language=objc
func (c_ CloudSharingServiceDelegateObject) OptionsForSharingServiceShareProvider(cloudKitSharingService SharingService, provider foundation.ItemProvider) CloudKitSharingServiceOptions {
	rv := objc.Call[CloudKitSharingServiceOptions](c_, objc.Sel("optionsForSharingService:shareProvider:"), cloudKitSharingService, provider)
	return rv
}

func (c_ CloudSharingServiceDelegateObject) HasSharingServiceDidStopSharing() bool {
	return c_.RespondsToSelector(objc.Sel("sharingService:didStopSharing:"))
}

// Tells the delegate when the user stops sharing the CloudKit share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644709-sharingservice?language=objc
func (c_ CloudSharingServiceDelegateObject) SharingServiceDidStopSharing(sharingService SharingService, share objc.Object) {
	objc.Call[objc.Void](c_, objc.Sel("sharingService:didStopSharing:"), sharingService, share)
}

func (c_ CloudSharingServiceDelegateObject) HasSharingServiceDidCompleteForItemsError() bool {
	return c_.RespondsToSelector(objc.Sel("sharingService:didCompleteForItems:error:"))
}

// Tells the delegate when the cloud-sharing service completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644666-sharingservice?language=objc
func (c_ CloudSharingServiceDelegateObject) SharingServiceDidCompleteForItemsError(sharingService SharingService, items []objc.Object, error foundation.Error) {
	objc.Call[objc.Void](c_, objc.Sel("sharingService:didCompleteForItems:error:"), sharingService, items, error)
}

func (c_ CloudSharingServiceDelegateObject) HasSharingServiceDidSaveShare() bool {
	return c_.RespondsToSelector(objc.Sel("sharingService:didSaveShare:"))
}

// Tells the delegate when the cloud-sharing service saves the CloudKit share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644712-sharingservice?language=objc
func (c_ CloudSharingServiceDelegateObject) SharingServiceDidSaveShare(sharingService SharingService, share objc.Object) {
	objc.Call[objc.Void](c_, objc.Sel("sharingService:didSaveShare:"), sharingService, share)
}
