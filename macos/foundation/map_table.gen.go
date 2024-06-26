// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [MapTable] class.
var MapTableClass = _MapTableClass{objc.GetClass("NSMapTable")}

type _MapTableClass struct {
	objc.Class
}

// An interface definition for the [MapTable] class.
type IMapTable interface {
	objc.IObject
	DictionaryRepresentation() Dictionary
	RemoveAllObjects()
	RemoveObjectForKey(aKey objc.IObject)
	KeyEnumerator() Enumerator
	ObjectForKey(aKey objc.IObject) objc.Object
	SetObjectForKey(anObject objc.IObject, aKey objc.IObject)
	ObjectEnumerator() Enumerator
	ValuePointerFunctions() PointerFunctions
	Count() uint
	KeyPointerFunctions() PointerFunctions
}

// A collection similar to a dictionary, but with a broader range of available memory semantics. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable?language=objc
type MapTable struct {
	objc.Object
}

func MapTableFrom(ptr unsafe.Pointer) MapTable {
	return MapTable{
		Object: objc.ObjectFrom(ptr),
	}
}

func (m_ MapTable) InitWithKeyOptionsValueOptionsCapacity(keyOptions PointerFunctionsOptions, valueOptions PointerFunctionsOptions, initialCapacity uint) MapTable {
	rv := objc.Call[MapTable](m_, objc.Sel("initWithKeyOptions:valueOptions:capacity:"), keyOptions, valueOptions, initialCapacity)
	return rv
}

// Returns a map table, initialized with the given options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391382-initwithkeyoptions?language=objc
func NewMapTableWithKeyOptionsValueOptionsCapacity(keyOptions PointerFunctionsOptions, valueOptions PointerFunctionsOptions, initialCapacity uint) MapTable {
	instance := MapTableClass.Alloc().InitWithKeyOptionsValueOptionsCapacity(keyOptions, valueOptions, initialCapacity)
	instance.Autorelease()
	return instance
}

func (m_ MapTable) InitWithKeyPointerFunctionsValuePointerFunctionsCapacity(keyFunctions IPointerFunctions, valueFunctions IPointerFunctions, initialCapacity uint) MapTable {
	rv := objc.Call[MapTable](m_, objc.Sel("initWithKeyPointerFunctions:valuePointerFunctions:capacity:"), keyFunctions, valueFunctions, initialCapacity)
	return rv
}

// Returns a map table, initialized with the given functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391429-initwithkeypointerfunctions?language=objc
func NewMapTableWithKeyPointerFunctionsValuePointerFunctionsCapacity(keyFunctions IPointerFunctions, valueFunctions IPointerFunctions, initialCapacity uint) MapTable {
	instance := MapTableClass.Alloc().InitWithKeyPointerFunctionsValuePointerFunctionsCapacity(keyFunctions, valueFunctions, initialCapacity)
	instance.Autorelease()
	return instance
}

func (mc _MapTableClass) Alloc() MapTable {
	rv := objc.Call[MapTable](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MapTableClass) New() MapTable {
	rv := objc.Call[MapTable](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMapTable() MapTable {
	return MapTableClass.New()
}

func (m_ MapTable) Init() MapTable {
	rv := objc.Call[MapTable](m_, objc.Sel("init"))
	return rv
}

// Returns a new map table object which has weak references to the keys and strong references to the values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391346-weaktostrongobjectsmaptable?language=objc
func (mc _MapTableClass) WeakToStrongObjectsMapTable() MapTable {
	rv := objc.Call[MapTable](mc, objc.Sel("weakToStrongObjectsMapTable"))
	return rv
}

// Returns a new map table object which has weak references to the keys and strong references to the values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391346-weaktostrongobjectsmaptable?language=objc
func MapTable_WeakToStrongObjectsMapTable() MapTable {
	return MapTableClass.WeakToStrongObjectsMapTable()
}

// Returns a dictionary representation of the map table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391402-dictionaryrepresentation?language=objc
func (m_ MapTable) DictionaryRepresentation() Dictionary {
	rv := objc.Call[Dictionary](m_, objc.Sel("dictionaryRepresentation"))
	return rv
}

// Empties the map table of its entries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391386-removeallobjects?language=objc
func (m_ MapTable) RemoveAllObjects() {
	objc.Call[objc.Void](m_, objc.Sel("removeAllObjects"))
}

// Returns a new map table, initialized with the given options [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391414-maptablewithkeyoptions?language=objc
func (mc _MapTableClass) MapTableWithKeyOptionsValueOptions(keyOptions PointerFunctionsOptions, valueOptions PointerFunctionsOptions) MapTable {
	rv := objc.Call[MapTable](mc, objc.Sel("mapTableWithKeyOptions:valueOptions:"), keyOptions, valueOptions)
	return rv
}

// Returns a new map table, initialized with the given options [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391414-maptablewithkeyoptions?language=objc
func MapTable_MapTableWithKeyOptionsValueOptions(keyOptions PointerFunctionsOptions, valueOptions PointerFunctionsOptions) MapTable {
	return MapTableClass.MapTableWithKeyOptionsValueOptions(keyOptions, valueOptions)
}

// Removes a given key and its associated value from the map table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391461-removeobjectforkey?language=objc
func (m_ MapTable) RemoveObjectForKey(aKey objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("removeObjectForKey:"), aKey)
}

// Returns a new map table object which has strong references to the keys and values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391440-strongtostrongobjectsmaptable?language=objc
func (mc _MapTableClass) StrongToStrongObjectsMapTable() MapTable {
	rv := objc.Call[MapTable](mc, objc.Sel("strongToStrongObjectsMapTable"))
	return rv
}

// Returns a new map table object which has strong references to the keys and values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391440-strongtostrongobjectsmaptable?language=objc
func MapTable_StrongToStrongObjectsMapTable() MapTable {
	return MapTableClass.StrongToStrongObjectsMapTable()
}

// Returns a new map table object which has weak references to the keys and values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391430-weaktoweakobjectsmaptable?language=objc
func (mc _MapTableClass) WeakToWeakObjectsMapTable() MapTable {
	rv := objc.Call[MapTable](mc, objc.Sel("weakToWeakObjectsMapTable"))
	return rv
}

// Returns a new map table object which has weak references to the keys and values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391430-weaktoweakobjectsmaptable?language=objc
func MapTable_WeakToWeakObjectsMapTable() MapTable {
	return MapTableClass.WeakToWeakObjectsMapTable()
}

// Returns an enumerator object that lets you access each key in the map table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391398-keyenumerator?language=objc
func (m_ MapTable) KeyEnumerator() Enumerator {
	rv := objc.Call[Enumerator](m_, objc.Sel("keyEnumerator"))
	return rv
}

// Returns a the value associated with a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391444-objectforkey?language=objc
func (m_ MapTable) ObjectForKey(aKey objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("objectForKey:"), aKey)
	return rv
}

// Adds a given key-value pair to the map table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391457-setobject?language=objc
func (m_ MapTable) SetObjectForKey(anObject objc.IObject, aKey objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setObject:forKey:"), anObject, aKey)
}

// Returns an enumerator object that lets you access each value in the map table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391400-objectenumerator?language=objc
func (m_ MapTable) ObjectEnumerator() Enumerator {
	rv := objc.Call[Enumerator](m_, objc.Sel("objectEnumerator"))
	return rv
}

// Returns a new map table object which has strong references to the keys and weak references to the values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391366-strongtoweakobjectsmaptable?language=objc
func (mc _MapTableClass) StrongToWeakObjectsMapTable() MapTable {
	rv := objc.Call[MapTable](mc, objc.Sel("strongToWeakObjectsMapTable"))
	return rv
}

// Returns a new map table object which has strong references to the keys and weak references to the values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391366-strongtoweakobjectsmaptable?language=objc
func MapTable_StrongToWeakObjectsMapTable() MapTable {
	return MapTableClass.StrongToWeakObjectsMapTable()
}

// The pointer functions the map table uses to manage values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391467-valuepointerfunctions?language=objc
func (m_ MapTable) ValuePointerFunctions() PointerFunctions {
	rv := objc.Call[PointerFunctions](m_, objc.Sel("valuePointerFunctions"))
	return rv
}

// The number of key-value pairs in the map table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391360-count?language=objc
func (m_ MapTable) Count() uint {
	rv := objc.Call[uint](m_, objc.Sel("count"))
	return rv
}

// The pointer functions the map table uses to manage keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/1391412-keypointerfunctions?language=objc
func (m_ MapTable) KeyPointerFunctions() PointerFunctions {
	rv := objc.Call[PointerFunctions](m_, objc.Sel("keyPointerFunctions"))
	return rv
}
