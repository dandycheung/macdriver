// Code generated by DarwinKit. DO NOT EDIT.

package quartz

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The IKScannerDeviceViewDelegate protocol defines the delegate protocol that the IKScannerDeviceView delegate must conform to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate?language=objc
type PScannerDeviceViewDelegate interface {
	// optional
	ScannerDeviceViewDidScanToURLError(scannerDeviceView ScannerDeviceView, url foundation.URL, error foundation.Error)
	HasScannerDeviceViewDidScanToURLError() bool

	// optional
	ScannerDeviceViewDidEncounterError(scannerDeviceView ScannerDeviceView, error foundation.Error)
	HasScannerDeviceViewDidEncounterError() bool

	// optional
	ScannerDeviceViewDidScanToURLFileDataError(scannerDeviceView ScannerDeviceView, url foundation.URL, data []byte, error foundation.Error)
	HasScannerDeviceViewDidScanToURLFileDataError() bool

	// optional
	ScannerDeviceViewDidScanToBandDataScanInfoError(scannerDeviceView ScannerDeviceView, data objc.Object, scanInfo foundation.Dictionary, error foundation.Error)
	HasScannerDeviceViewDidScanToBandDataScanInfoError() bool
}

// A delegate implementation builder for the [PScannerDeviceViewDelegate] protocol.
type ScannerDeviceViewDelegate struct {
	_ScannerDeviceViewDidScanToURLError              func(scannerDeviceView ScannerDeviceView, url foundation.URL, error foundation.Error)
	_ScannerDeviceViewDidEncounterError              func(scannerDeviceView ScannerDeviceView, error foundation.Error)
	_ScannerDeviceViewDidScanToURLFileDataError      func(scannerDeviceView ScannerDeviceView, url foundation.URL, data []byte, error foundation.Error)
	_ScannerDeviceViewDidScanToBandDataScanInfoError func(scannerDeviceView ScannerDeviceView, data objc.Object, scanInfo foundation.Dictionary, error foundation.Error)
}

func (di *ScannerDeviceViewDelegate) HasScannerDeviceViewDidScanToURLError() bool {
	return di._ScannerDeviceViewDidScanToURLError != nil
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate/1504768-scannerdeviceview?language=objc
func (di *ScannerDeviceViewDelegate) SetScannerDeviceViewDidScanToURLError(f func(scannerDeviceView ScannerDeviceView, url foundation.URL, error foundation.Error)) {
	di._ScannerDeviceViewDidScanToURLError = f
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate/1504768-scannerdeviceview?language=objc
func (di *ScannerDeviceViewDelegate) ScannerDeviceViewDidScanToURLError(scannerDeviceView ScannerDeviceView, url foundation.URL, error foundation.Error) {
	di._ScannerDeviceViewDidScanToURLError(scannerDeviceView, url, error)
}
func (di *ScannerDeviceViewDelegate) HasScannerDeviceViewDidEncounterError() bool {
	return di._ScannerDeviceViewDidEncounterError != nil
}

// Invoked whenever the scanner encounters an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate/1503476-scannerdeviceview?language=objc
func (di *ScannerDeviceViewDelegate) SetScannerDeviceViewDidEncounterError(f func(scannerDeviceView ScannerDeviceView, error foundation.Error)) {
	di._ScannerDeviceViewDidEncounterError = f
}

// Invoked whenever the scanner encounters an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate/1503476-scannerdeviceview?language=objc
func (di *ScannerDeviceViewDelegate) ScannerDeviceViewDidEncounterError(scannerDeviceView ScannerDeviceView, error foundation.Error) {
	di._ScannerDeviceViewDidEncounterError(scannerDeviceView, error)
}
func (di *ScannerDeviceViewDelegate) HasScannerDeviceViewDidScanToURLFileDataError() bool {
	return di._ScannerDeviceViewDidScanToURLFileDataError != nil
}

// Invoked when the scan has completed and the data is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate/1504462-scannerdeviceview?language=objc
func (di *ScannerDeviceViewDelegate) SetScannerDeviceViewDidScanToURLFileDataError(f func(scannerDeviceView ScannerDeviceView, url foundation.URL, data []byte, error foundation.Error)) {
	di._ScannerDeviceViewDidScanToURLFileDataError = f
}

// Invoked when the scan has completed and the data is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate/1504462-scannerdeviceview?language=objc
func (di *ScannerDeviceViewDelegate) ScannerDeviceViewDidScanToURLFileDataError(scannerDeviceView ScannerDeviceView, url foundation.URL, data []byte, error foundation.Error) {
	di._ScannerDeviceViewDidScanToURLFileDataError(scannerDeviceView, url, data, error)
}
func (di *ScannerDeviceViewDelegate) HasScannerDeviceViewDidScanToBandDataScanInfoError() bool {
	return di._ScannerDeviceViewDidScanToBandDataScanInfoError != nil
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate/1503867-scannerdeviceview?language=objc
func (di *ScannerDeviceViewDelegate) SetScannerDeviceViewDidScanToBandDataScanInfoError(f func(scannerDeviceView ScannerDeviceView, data objc.Object, scanInfo foundation.Dictionary, error foundation.Error)) {
	di._ScannerDeviceViewDidScanToBandDataScanInfoError = f
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate/1503867-scannerdeviceview?language=objc
func (di *ScannerDeviceViewDelegate) ScannerDeviceViewDidScanToBandDataScanInfoError(scannerDeviceView ScannerDeviceView, data objc.Object, scanInfo foundation.Dictionary, error foundation.Error) {
	di._ScannerDeviceViewDidScanToBandDataScanInfoError(scannerDeviceView, data, scanInfo, error)
}

// ensure impl type implements protocol interface
var _ PScannerDeviceViewDelegate = (*ScannerDeviceViewDelegateObject)(nil)

// A concrete type for the [PScannerDeviceViewDelegate] protocol.
type ScannerDeviceViewDelegateObject struct {
	objc.Object
}

func (s_ ScannerDeviceViewDelegateObject) HasScannerDeviceViewDidScanToURLError() bool {
	return s_.RespondsToSelector(objc.Sel("scannerDeviceView:didScanToURL:error:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate/1504768-scannerdeviceview?language=objc
func (s_ ScannerDeviceViewDelegateObject) ScannerDeviceViewDidScanToURLError(scannerDeviceView ScannerDeviceView, url foundation.URL, error foundation.Error) {
	objc.Call[objc.Void](s_, objc.Sel("scannerDeviceView:didScanToURL:error:"), scannerDeviceView, url, error)
}

func (s_ ScannerDeviceViewDelegateObject) HasScannerDeviceViewDidEncounterError() bool {
	return s_.RespondsToSelector(objc.Sel("scannerDeviceView:didEncounterError:"))
}

// Invoked whenever the scanner encounters an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate/1503476-scannerdeviceview?language=objc
func (s_ ScannerDeviceViewDelegateObject) ScannerDeviceViewDidEncounterError(scannerDeviceView ScannerDeviceView, error foundation.Error) {
	objc.Call[objc.Void](s_, objc.Sel("scannerDeviceView:didEncounterError:"), scannerDeviceView, error)
}

func (s_ ScannerDeviceViewDelegateObject) HasScannerDeviceViewDidScanToURLFileDataError() bool {
	return s_.RespondsToSelector(objc.Sel("scannerDeviceView:didScanToURL:fileData:error:"))
}

// Invoked when the scan has completed and the data is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate/1504462-scannerdeviceview?language=objc
func (s_ ScannerDeviceViewDelegateObject) ScannerDeviceViewDidScanToURLFileDataError(scannerDeviceView ScannerDeviceView, url foundation.URL, data []byte, error foundation.Error) {
	objc.Call[objc.Void](s_, objc.Sel("scannerDeviceView:didScanToURL:fileData:error:"), scannerDeviceView, url, data, error)
}

func (s_ ScannerDeviceViewDelegateObject) HasScannerDeviceViewDidScanToBandDataScanInfoError() bool {
	return s_.RespondsToSelector(objc.Sel("scannerDeviceView:didScanToBandData:scanInfo:error:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate/1503867-scannerdeviceview?language=objc
func (s_ ScannerDeviceViewDelegateObject) ScannerDeviceViewDidScanToBandDataScanInfoError(scannerDeviceView ScannerDeviceView, data objc.Object, scanInfo foundation.Dictionary, error foundation.Error) {
	objc.Call[objc.Void](s_, objc.Sel("scannerDeviceView:didScanToBandData:scanInfo:error:"), scannerDeviceView, data, scanInfo, error)
}
