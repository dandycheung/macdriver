// Code generated by DarwinKit. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [IncrementalStore] class.
var IncrementalStoreClass = _IncrementalStoreClass{objc.GetClass("NSIncrementalStore")}

type _IncrementalStoreClass struct {
	objc.Class
}

// An interface definition for the [IncrementalStore] class.
type IIncrementalStore interface {
	IPersistentStore
	ObtainPermanentIDsForObjectsError(array []IManagedObject, error unsafe.Pointer) []ManagedObjectID
	ManagedObjectContextDidRegisterObjectsWithIDs(objectIDs []IManagedObjectID)
	ReferenceObjectForObjectID(objectID IManagedObjectID) objc.Object
	ManagedObjectContextDidUnregisterObjectsWithIDs(objectIDs []IManagedObjectID)
	NewObjectIDForEntityReferenceObject(entity IEntityDescription, data objc.IObject) ManagedObjectID
	NewValuesForObjectWithIDWithContextError(objectID IManagedObjectID, context IManagedObjectContext, error unsafe.Pointer) IncrementalStoreNode
	NewValueForRelationshipForObjectWithIDWithContextError(relationship IRelationshipDescription, objectID IManagedObjectID, context IManagedObjectContext, error unsafe.Pointer) objc.Object
	ExecuteRequestWithContextError(request IPersistentStoreRequest, context IManagedObjectContext, error unsafe.Pointer) objc.Object
}

// An abstract superclass defining the API through which Core Data communicates with a store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstore?language=objc
type IncrementalStore struct {
	PersistentStore
}

func IncrementalStoreFrom(ptr unsafe.Pointer) IncrementalStore {
	return IncrementalStore{
		PersistentStore: PersistentStoreFrom(ptr),
	}
}

func (ic _IncrementalStoreClass) Alloc() IncrementalStore {
	rv := objc.Call[IncrementalStore](ic, objc.Sel("alloc"))
	return rv
}

func (ic _IncrementalStoreClass) New() IncrementalStore {
	rv := objc.Call[IncrementalStore](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewIncrementalStore() IncrementalStore {
	return IncrementalStoreClass.New()
}

func (i_ IncrementalStore) Init() IncrementalStore {
	rv := objc.Call[IncrementalStore](i_, objc.Sel("init"))
	return rv
}

func (i_ IncrementalStore) InitWithPersistentStoreCoordinatorConfigurationNameURLOptions(root IPersistentStoreCoordinator, name string, url foundation.IURL, options foundation.Dictionary) IncrementalStore {
	rv := objc.Call[IncrementalStore](i_, objc.Sel("initWithPersistentStoreCoordinator:configurationName:URL:options:"), root, name, url, options)
	return rv
}

// Returns a store initialized with the given arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506232-initwithpersistentstorecoordinat?language=objc
func NewIncrementalStoreWithPersistentStoreCoordinatorConfigurationNameURLOptions(root IPersistentStoreCoordinator, name string, url foundation.IURL, options foundation.Dictionary) IncrementalStore {
	instance := IncrementalStoreClass.Alloc().InitWithPersistentStoreCoordinatorConfigurationNameURLOptions(root, name, url, options)
	instance.Autorelease()
	return instance
}

// Returns an array containing the object IDs for a given array of newly-inserted objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstore/1506348-obtainpermanentidsforobjects?language=objc
func (i_ IncrementalStore) ObtainPermanentIDsForObjectsError(array []IManagedObject, error unsafe.Pointer) []ManagedObjectID {
	rv := objc.Call[[]ManagedObjectID](i_, objc.Sel("obtainPermanentIDsForObjects:error:"), array, error)
	return rv
}

// Indicates that objects identified by a given array of object IDs are in use in a managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstore/1506199-managedobjectcontextdidregistero?language=objc
func (i_ IncrementalStore) ManagedObjectContextDidRegisterObjectsWithIDs(objectIDs []IManagedObjectID) {
	objc.Call[objc.Void](i_, objc.Sel("managedObjectContextDidRegisterObjectsWithIDs:"), objectIDs)
}

// Returns the reference data used to construct a given object ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstore/1506828-referenceobjectforobjectid?language=objc
func (i_ IncrementalStore) ReferenceObjectForObjectID(objectID IManagedObjectID) objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("referenceObjectForObjectID:"), objectID)
	return rv
}

// Returns the identifier for the store at a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstore/1506781-identifierfornewstoreaturl?language=objc
func (ic _IncrementalStoreClass) IdentifierForNewStoreAtURL(storeURL foundation.IURL) objc.Object {
	rv := objc.Call[objc.Object](ic, objc.Sel("identifierForNewStoreAtURL:"), storeURL)
	return rv
}

// Returns the identifier for the store at a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstore/1506781-identifierfornewstoreaturl?language=objc
func IncrementalStore_IdentifierForNewStoreAtURL(storeURL foundation.IURL) objc.Object {
	return IncrementalStoreClass.IdentifierForNewStoreAtURL(storeURL)
}

// Indicates that objects identified by a given array of object IDs are no longer being used by a managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstore/1506878-managedobjectcontextdidunregiste?language=objc
func (i_ IncrementalStore) ManagedObjectContextDidUnregisterObjectsWithIDs(objectIDs []IManagedObjectID) {
	objc.Call[objc.Void](i_, objc.Sel("managedObjectContextDidUnregisterObjectsWithIDs:"), objectIDs)
}

// Returns a new object ID that uses given data as the key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstore/1506666-newobjectidforentity?language=objc
func (i_ IncrementalStore) NewObjectIDForEntityReferenceObject(entity IEntityDescription, data objc.IObject) ManagedObjectID {
	rv := objc.Call[ManagedObjectID](i_, objc.Sel("newObjectIDForEntity:referenceObject:"), entity, data)
	return rv
}

// Returns an incremental store node encapsulating the persistent external values of the object with a given object ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstore/1506729-newvaluesforobjectwithid?language=objc
func (i_ IncrementalStore) NewValuesForObjectWithIDWithContextError(objectID IManagedObjectID, context IManagedObjectContext, error unsafe.Pointer) IncrementalStoreNode {
	rv := objc.Call[IncrementalStoreNode](i_, objc.Sel("newValuesForObjectWithID:withContext:error:"), objectID, context, error)
	return rv
}

// Returns the relationship for the given relationship of the object with a given object ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstore/1506438-newvalueforrelationship?language=objc
func (i_ IncrementalStore) NewValueForRelationshipForObjectWithIDWithContextError(relationship IRelationshipDescription, objectID IManagedObjectID, context IManagedObjectContext, error unsafe.Pointer) objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("newValueForRelationship:forObjectWithID:withContext:error:"), relationship, objectID, context, error)
	return rv
}

// Returns a value as appropriate for the given request, or nil if the request cannot be completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstore/1506653-executerequest?language=objc
func (i_ IncrementalStore) ExecuteRequestWithContextError(request IPersistentStoreRequest, context IManagedObjectContext, error unsafe.Pointer) objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("executeRequest:withContext:error:"), request, context, error)
	return rv
}
