// Code generated by DarwinKit. DO NOT EDIT.

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/mps"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ShapedType] class.
var ShapedTypeClass = _ShapedTypeClass{objc.GetClass("MPSGraphShapedType")}

type _ShapedTypeClass struct {
	objc.Class
}

// An interface definition for the [ShapedType] class.
type IShapedType interface {
	IType
	IsEqualTo(object IShapedType) bool
	DataType() mps.DataType
	SetDataType(value mps.DataType)
	Shape() *foundation.Array
	SetShape(value *foundation.Array)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphshapedtype?language=objc
type ShapedType struct {
	Type
}

func ShapedTypeFrom(ptr unsafe.Pointer) ShapedType {
	return ShapedType{
		Type: TypeFrom(ptr),
	}
}

func (s_ ShapedType) InitWithShapeDataType(shape *foundation.Array, dataType mps.DataType) ShapedType {
	rv := objc.Call[ShapedType](s_, objc.Sel("initWithShape:dataType:"), shape, dataType)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphshapedtype/3600280-initwithshape?language=objc
func NewShapedTypeWithShapeDataType(shape *foundation.Array, dataType mps.DataType) ShapedType {
	instance := ShapedTypeClass.Alloc().InitWithShapeDataType(shape, dataType)
	instance.Autorelease()
	return instance
}

func (sc _ShapedTypeClass) Alloc() ShapedType {
	rv := objc.Call[ShapedType](sc, objc.Sel("alloc"))
	return rv
}

func (sc _ShapedTypeClass) New() ShapedType {
	rv := objc.Call[ShapedType](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewShapedType() ShapedType {
	return ShapedTypeClass.New()
}

func (s_ ShapedType) Init() ShapedType {
	rv := objc.Call[ShapedType](s_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphshapedtype/3600281-isequalto?language=objc
func (s_ ShapedType) IsEqualTo(object IShapedType) bool {
	rv := objc.Call[bool](s_, objc.Sel("isEqualTo:"), object)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphshapedtype/3600279-datatype?language=objc
func (s_ ShapedType) DataType() mps.DataType {
	rv := objc.Call[mps.DataType](s_, objc.Sel("dataType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphshapedtype/3600279-datatype?language=objc
func (s_ ShapedType) SetDataType(value mps.DataType) {
	objc.Call[objc.Void](s_, objc.Sel("setDataType:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphshapedtype/3600282-shape?language=objc
func (s_ ShapedType) Shape() *foundation.Array {
	rv := objc.Call[*foundation.Array](s_, objc.Sel("shape"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphshapedtype/3600282-shape?language=objc
func (s_ ShapedType) SetShape(value *foundation.Array) {
	objc.Call[objc.Void](s_, objc.Sel("setShape:"), value)
}
