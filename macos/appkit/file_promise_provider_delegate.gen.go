// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that provides the name of the promised file and writes the file to the destination directory when the file promise is fulfilled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate?language=objc
type PFilePromiseProviderDelegate interface {
	// optional
	FilePromiseProviderFileNameForType(filePromiseProvider FilePromiseProvider, fileType string) string
	HasFilePromiseProviderFileNameForType() bool

	// optional
	OperationQueueForFilePromiseProvider(filePromiseProvider FilePromiseProvider) foundation.OperationQueue
	HasOperationQueueForFilePromiseProvider() bool

	// optional
	FilePromiseProviderWritePromiseToURLCompletionHandler(filePromiseProvider FilePromiseProvider, url foundation.URL, completionHandler func(errorOrNil foundation.Error))
	HasFilePromiseProviderWritePromiseToURLCompletionHandler() bool
}

// A delegate implementation builder for the [PFilePromiseProviderDelegate] protocol.
type FilePromiseProviderDelegate struct {
	_FilePromiseProviderFileNameForType                    func(filePromiseProvider FilePromiseProvider, fileType string) string
	_OperationQueueForFilePromiseProvider                  func(filePromiseProvider FilePromiseProvider) foundation.OperationQueue
	_FilePromiseProviderWritePromiseToURLCompletionHandler func(filePromiseProvider FilePromiseProvider, url foundation.URL, completionHandler func(errorOrNil foundation.Error))
}

func (di *FilePromiseProviderDelegate) HasFilePromiseProviderFileNameForType() bool {
	return di._FilePromiseProviderFileNameForType != nil
}

// Provides the drag destination file's name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate/2369278-filepromiseprovider?language=objc
func (di *FilePromiseProviderDelegate) SetFilePromiseProviderFileNameForType(f func(filePromiseProvider FilePromiseProvider, fileType string) string) {
	di._FilePromiseProviderFileNameForType = f
}

// Provides the drag destination file's name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate/2369278-filepromiseprovider?language=objc
func (di *FilePromiseProviderDelegate) FilePromiseProviderFileNameForType(filePromiseProvider FilePromiseProvider, fileType string) string {
	return di._FilePromiseProviderFileNameForType(filePromiseProvider, fileType)
}
func (di *FilePromiseProviderDelegate) HasOperationQueueForFilePromiseProvider() bool {
	return di._OperationQueueForFilePromiseProvider != nil
}

// Returns the operation queue from which to issue the write request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate/2369279-operationqueueforfilepromiseprov?language=objc
func (di *FilePromiseProviderDelegate) SetOperationQueueForFilePromiseProvider(f func(filePromiseProvider FilePromiseProvider) foundation.OperationQueue) {
	di._OperationQueueForFilePromiseProvider = f
}

// Returns the operation queue from which to issue the write request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate/2369279-operationqueueforfilepromiseprov?language=objc
func (di *FilePromiseProviderDelegate) OperationQueueForFilePromiseProvider(filePromiseProvider FilePromiseProvider) foundation.OperationQueue {
	return di._OperationQueueForFilePromiseProvider(filePromiseProvider)
}
func (di *FilePromiseProviderDelegate) HasFilePromiseProviderWritePromiseToURLCompletionHandler() bool {
	return di._FilePromiseProviderWritePromiseToURLCompletionHandler != nil
}

// Writes the contents of a promise to the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate/1644244-filepromiseprovider?language=objc
func (di *FilePromiseProviderDelegate) SetFilePromiseProviderWritePromiseToURLCompletionHandler(f func(filePromiseProvider FilePromiseProvider, url foundation.URL, completionHandler func(errorOrNil foundation.Error))) {
	di._FilePromiseProviderWritePromiseToURLCompletionHandler = f
}

// Writes the contents of a promise to the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate/1644244-filepromiseprovider?language=objc
func (di *FilePromiseProviderDelegate) FilePromiseProviderWritePromiseToURLCompletionHandler(filePromiseProvider FilePromiseProvider, url foundation.URL, completionHandler func(errorOrNil foundation.Error)) {
	di._FilePromiseProviderWritePromiseToURLCompletionHandler(filePromiseProvider, url, completionHandler)
}

// ensure impl type implements protocol interface
var _ PFilePromiseProviderDelegate = (*FilePromiseProviderDelegateObject)(nil)

// A concrete type for the [PFilePromiseProviderDelegate] protocol.
type FilePromiseProviderDelegateObject struct {
	objc.Object
}

func (f_ FilePromiseProviderDelegateObject) HasFilePromiseProviderFileNameForType() bool {
	return f_.RespondsToSelector(objc.Sel("filePromiseProvider:fileNameForType:"))
}

// Provides the drag destination file's name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate/2369278-filepromiseprovider?language=objc
func (f_ FilePromiseProviderDelegateObject) FilePromiseProviderFileNameForType(filePromiseProvider FilePromiseProvider, fileType string) string {
	rv := objc.Call[string](f_, objc.Sel("filePromiseProvider:fileNameForType:"), filePromiseProvider, fileType)
	return rv
}

func (f_ FilePromiseProviderDelegateObject) HasOperationQueueForFilePromiseProvider() bool {
	return f_.RespondsToSelector(objc.Sel("operationQueueForFilePromiseProvider:"))
}

// Returns the operation queue from which to issue the write request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate/2369279-operationqueueforfilepromiseprov?language=objc
func (f_ FilePromiseProviderDelegateObject) OperationQueueForFilePromiseProvider(filePromiseProvider FilePromiseProvider) foundation.OperationQueue {
	rv := objc.Call[foundation.OperationQueue](f_, objc.Sel("operationQueueForFilePromiseProvider:"), filePromiseProvider)
	return rv
}

func (f_ FilePromiseProviderDelegateObject) HasFilePromiseProviderWritePromiseToURLCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("filePromiseProvider:writePromiseToURL:completionHandler:"))
}

// Writes the contents of a promise to the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate/1644244-filepromiseprovider?language=objc
func (f_ FilePromiseProviderDelegateObject) FilePromiseProviderWritePromiseToURLCompletionHandler(filePromiseProvider FilePromiseProvider, url foundation.URL, completionHandler func(errorOrNil foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("filePromiseProvider:writePromiseToURL:completionHandler:"), filePromiseProvider, url, completionHandler)
}
