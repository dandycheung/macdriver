// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ValueTransformer] class.
var ValueTransformerClass = _ValueTransformerClass{objc.GetClass("NSValueTransformer")}

type _ValueTransformerClass struct {
	objc.Class
}

// An interface definition for the [ValueTransformer] class.
type IValueTransformer interface {
	objc.IObject
	TransformedValue(value objc.IObject) objc.Object
	ReverseTransformedValue(value objc.IObject) objc.Object
}

// An abstract class used to transform values from one representation to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvaluetransformer?language=objc
type ValueTransformer struct {
	objc.Object
}

func ValueTransformerFrom(ptr unsafe.Pointer) ValueTransformer {
	return ValueTransformer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _ValueTransformerClass) Alloc() ValueTransformer {
	rv := objc.Call[ValueTransformer](vc, objc.Sel("alloc"))
	return rv
}

func (vc _ValueTransformerClass) New() ValueTransformer {
	rv := objc.Call[ValueTransformer](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewValueTransformer() ValueTransformer {
	return ValueTransformerClass.New()
}

func (v_ ValueTransformer) Init() ValueTransformer {
	rv := objc.Call[ValueTransformer](v_, objc.Sel("init"))
	return rv
}

// Returns the class of the value returned by the receiver for a forward transformation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvaluetransformer/1401998-transformedvalueclass?language=objc
func (vc _ValueTransformerClass) TransformedValueClass() objc.Class {
	rv := objc.Call[objc.Class](vc, objc.Sel("transformedValueClass"))
	return rv
}

// Returns the class of the value returned by the receiver for a forward transformation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvaluetransformer/1401998-transformedvalueclass?language=objc
func ValueTransformer_TransformedValueClass() objc.Class {
	return ValueTransformerClass.TransformedValueClass()
}

// Returns an array of all the registered value transformers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvaluetransformer/1402012-valuetransformernames?language=objc
func (vc _ValueTransformerClass) ValueTransformerNames() []ValueTransformerName {
	rv := objc.Call[[]ValueTransformerName](vc, objc.Sel("valueTransformerNames"))
	return rv
}

// Returns an array of all the registered value transformers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvaluetransformer/1402012-valuetransformernames?language=objc
func ValueTransformer_ValueTransformerNames() []ValueTransformerName {
	return ValueTransformerClass.ValueTransformerNames()
}

// Returns the value transformer identified by a given identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvaluetransformer/1402010-valuetransformerforname?language=objc
func (vc _ValueTransformerClass) ValueTransformerForName(name ValueTransformerName) ValueTransformer {
	rv := objc.Call[ValueTransformer](vc, objc.Sel("valueTransformerForName:"), name)
	return rv
}

// Returns the value transformer identified by a given identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvaluetransformer/1402010-valuetransformerforname?language=objc
func ValueTransformer_ValueTransformerForName(name ValueTransformerName) ValueTransformer {
	return ValueTransformerClass.ValueTransformerForName(name)
}

// Returns a Boolean value that indicates whether the receiver can reverse a transformation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvaluetransformer/1402016-allowsreversetransformation?language=objc
func (vc _ValueTransformerClass) AllowsReverseTransformation() bool {
	rv := objc.Call[bool](vc, objc.Sel("allowsReverseTransformation"))
	return rv
}

// Returns a Boolean value that indicates whether the receiver can reverse a transformation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvaluetransformer/1402016-allowsreversetransformation?language=objc
func ValueTransformer_AllowsReverseTransformation() bool {
	return ValueTransformerClass.AllowsReverseTransformation()
}

// Returns the result of transforming a given value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvaluetransformer/1402004-transformedvalue?language=objc
func (v_ ValueTransformer) TransformedValue(value objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](v_, objc.Sel("transformedValue:"), value)
	return rv
}

// Returns the result of the reverse transformation of a given value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvaluetransformer/1402020-reversetransformedvalue?language=objc
func (v_ ValueTransformer) ReverseTransformedValue(value objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](v_, objc.Sel("reverseTransformedValue:"), value)
	return rv
}

// Registers the provided value transformer with a given identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvaluetransformer/1402018-setvaluetransformer?language=objc
func (vc _ValueTransformerClass) SetValueTransformerForName(transformer IValueTransformer, name ValueTransformerName) {
	objc.Call[objc.Void](vc, objc.Sel("setValueTransformer:forName:"), transformer, name)
}

// Registers the provided value transformer with a given identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvaluetransformer/1402018-setvaluetransformer?language=objc
func ValueTransformer_SetValueTransformerForName(transformer IValueTransformer, name ValueTransformerName) {
	ValueTransformerClass.SetValueTransformerForName(transformer, name)
}
