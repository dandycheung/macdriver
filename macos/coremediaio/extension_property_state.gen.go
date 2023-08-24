// AUTO-GENERATED CODE, DO NOT MODIFY

package coremediaio

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExtensionPropertyState] class.
var ExtensionPropertyStateClass = _ExtensionPropertyStateClass{objc.GetClass("CMIOExtensionPropertyState")}

type _ExtensionPropertyStateClass struct {
	objc.Class
}

// An interface definition for the [ExtensionPropertyState] class.
type IExtensionPropertyState interface {
	objc.IObject
	Value() objc.Object
	Attributes() ExtensionPropertyAttributes
}

// An object that describes the state of a property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionpropertystate?language=objc
type ExtensionPropertyState struct {
	objc.Object
}

func ExtensionPropertyStateFrom(ptr unsafe.Pointer) ExtensionPropertyState {
	return ExtensionPropertyState{
		Object: objc.ObjectFrom(ptr),
	}
}

func (e_ ExtensionPropertyState) InitWithValueAttributes(value objc.IObject, attributes IExtensionPropertyAttributes) ExtensionPropertyState {
	rv := objc.Call[ExtensionPropertyState](e_, objc.Sel("initWithValue:attributes:"), objc.Ptr(value), objc.Ptr(attributes))
	return rv
}

// Creates a property state with a value and attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionpropertystate/3915872-initwithvalue?language=objc
func NewExtensionPropertyStateWithValueAttributes(value objc.IObject, attributes IExtensionPropertyAttributes) ExtensionPropertyState {
	instance := ExtensionPropertyStateClass.Alloc().InitWithValueAttributes(value, attributes)
	instance.Autorelease()
	return instance
}

func (ec _ExtensionPropertyStateClass) PropertyStateWithValue(value objc.IObject) ExtensionPropertyState {
	rv := objc.Call[ExtensionPropertyState](ec, objc.Sel("propertyStateWithValue:"), objc.Ptr(value))
	return rv
}

// Returns a new property state with a value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionpropertystate/3915873-propertystatewithvalue?language=objc
func ExtensionPropertyState_PropertyStateWithValue(value objc.IObject) ExtensionPropertyState {
	return ExtensionPropertyStateClass.PropertyStateWithValue(value)
}

func (ec _ExtensionPropertyStateClass) Alloc() ExtensionPropertyState {
	rv := objc.Call[ExtensionPropertyState](ec, objc.Sel("alloc"))
	return rv
}

func ExtensionPropertyState_Alloc() ExtensionPropertyState {
	return ExtensionPropertyStateClass.Alloc()
}

func (ec _ExtensionPropertyStateClass) New() ExtensionPropertyState {
	rv := objc.Call[ExtensionPropertyState](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionPropertyState() ExtensionPropertyState {
	return ExtensionPropertyStateClass.New()
}

func (e_ ExtensionPropertyState) Init() ExtensionPropertyState {
	rv := objc.Call[ExtensionPropertyState](e_, objc.Sel("init"))
	return rv
}

// The value for a property state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionpropertystate/3915875-value?language=objc
func (e_ ExtensionPropertyState) Value() objc.Object {
	rv := objc.Call[objc.Object](e_, objc.Sel("value"))
	return rv
}

// The attributes for a property state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionpropertystate/3915870-attributes?language=objc
func (e_ ExtensionPropertyState) Attributes() ExtensionPropertyAttributes {
	rv := objc.Call[ExtensionPropertyAttributes](e_, objc.Sel("attributes"))
	return rv
}