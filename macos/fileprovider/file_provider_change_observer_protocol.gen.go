// Code generated by DarwinKit. DO NOT EDIT.

package fileprovider

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An observer that receives changes and deletions during enumeration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderchangeobserver?language=objc
type PFileProviderChangeObserver interface {
	// optional
	DidDeleteItemsWithIdentifiers(deletedItemIdentifiers []FileProviderItemIdentifier)
	HasDidDeleteItemsWithIdentifiers() bool

	// optional
	FinishEnumeratingChangesUpToSyncAnchorMoreComing(anchor FileProviderSyncAnchor, moreComing bool)
	HasFinishEnumeratingChangesUpToSyncAnchorMoreComing() bool

	// optional
	DidUpdateItems(updatedItems []objc.Object)
	HasDidUpdateItems() bool

	// optional
	FinishEnumeratingWithError(error foundation.Error)
	HasFinishEnumeratingWithError() bool

	// optional
	SuggestedBatchSize() int
	HasSuggestedBatchSize() bool
}

// ensure impl type implements protocol interface
var _ PFileProviderChangeObserver = (*FileProviderChangeObserverObject)(nil)

// A concrete type for the [PFileProviderChangeObserver] protocol.
type FileProviderChangeObserverObject struct {
	objc.Object
}

func (f_ FileProviderChangeObserverObject) HasDidDeleteItemsWithIdentifiers() bool {
	return f_.RespondsToSelector(objc.Sel("didDeleteItemsWithIdentifiers:"))
}

// Tells the observer that the specified items have been deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderchangeobserver/2879604-diddeleteitemswithidentifiers?language=objc
func (f_ FileProviderChangeObserverObject) DidDeleteItemsWithIdentifiers(deletedItemIdentifiers []FileProviderItemIdentifier) {
	objc.Call[objc.Void](f_, objc.Sel("didDeleteItemsWithIdentifiers:"), deletedItemIdentifiers)
}

func (f_ FileProviderChangeObserverObject) HasFinishEnumeratingChangesUpToSyncAnchorMoreComing() bool {
	return f_.RespondsToSelector(objc.Sel("finishEnumeratingChangesUpToSyncAnchor:moreComing:"))
}

// Tells the observer that all of the changes have been enumerated up to the specified sync anchor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderchangeobserver/2880220-finishenumeratingchangesuptosync?language=objc
func (f_ FileProviderChangeObserverObject) FinishEnumeratingChangesUpToSyncAnchorMoreComing(anchor FileProviderSyncAnchor, moreComing bool) {
	objc.Call[objc.Void](f_, objc.Sel("finishEnumeratingChangesUpToSyncAnchor:moreComing:"), anchor, moreComing)
}

func (f_ FileProviderChangeObserverObject) HasDidUpdateItems() bool {
	return f_.RespondsToSelector(objc.Sel("didUpdateItems:"))
}

// Tells the observer that the specified items have been updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderchangeobserver/2879607-didupdateitems?language=objc
func (f_ FileProviderChangeObserverObject) DidUpdateItems(updatedItems []objc.Object) {
	objc.Call[objc.Void](f_, objc.Sel("didUpdateItems:"), updatedItems)
}

func (f_ FileProviderChangeObserverObject) HasFinishEnumeratingWithError() bool {
	return f_.RespondsToSelector(objc.Sel("finishEnumeratingWithError:"))
}

// Tells the observer that an error occurred during change notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderchangeobserver/2879605-finishenumeratingwitherror?language=objc
func (f_ FileProviderChangeObserverObject) FinishEnumeratingWithError(error foundation.Error) {
	objc.Call[objc.Void](f_, objc.Sel("finishEnumeratingWithError:"), error)
}

func (f_ FileProviderChangeObserverObject) HasSuggestedBatchSize() bool {
	return f_.RespondsToSelector(objc.Sel("suggestedBatchSize"))
}

// The batch size that the system recommends. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderchangeobserver/3656525-suggestedbatchsize?language=objc
func (f_ FileProviderChangeObserverObject) SuggestedBatchSize() int {
	rv := objc.Call[int](f_, objc.Sel("suggestedBatchSize"))
	return rv
}
