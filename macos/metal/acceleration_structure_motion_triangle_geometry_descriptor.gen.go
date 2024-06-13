// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AccelerationStructureMotionTriangleGeometryDescriptor] class.
var AccelerationStructureMotionTriangleGeometryDescriptorClass = _AccelerationStructureMotionTriangleGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureMotionTriangleGeometryDescriptor")}

type _AccelerationStructureMotionTriangleGeometryDescriptorClass struct {
	objc.Class
}

// An interface definition for the [AccelerationStructureMotionTriangleGeometryDescriptor] class.
type IAccelerationStructureMotionTriangleGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
	IndexType() IndexType
	SetIndexType(value IndexType)
	IndexBuffer() BufferObject
	SetIndexBuffer(value PBuffer)
	SetIndexBufferObject(valueObject objc.IObject)
	TriangleCount() uint
	SetTriangleCount(value uint)
	IndexBufferOffset() uint
	SetIndexBufferOffset(value uint)
	VertexBuffers() []MotionKeyframeData
	SetVertexBuffers(value []IMotionKeyframeData)
	VertexStride() uint
	SetVertexStride(value uint)
}

// A description of a list of triangle primitives, as motion keyframe data, to turn into an acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor?language=objc
type AccelerationStructureMotionTriangleGeometryDescriptor struct {
	AccelerationStructureGeometryDescriptor
}

func AccelerationStructureMotionTriangleGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureMotionTriangleGeometryDescriptor {
	return AccelerationStructureMotionTriangleGeometryDescriptor{
		AccelerationStructureGeometryDescriptor: AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

func (ac _AccelerationStructureMotionTriangleGeometryDescriptorClass) Descriptor() AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Call[AccelerationStructureMotionTriangleGeometryDescriptor](ac, objc.Sel("descriptor"))
	return rv
}

// Creates a new triangle descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/3750483-descriptor?language=objc
func AccelerationStructureMotionTriangleGeometryDescriptor_Descriptor() AccelerationStructureMotionTriangleGeometryDescriptor {
	return AccelerationStructureMotionTriangleGeometryDescriptorClass.Descriptor()
}

func (ac _AccelerationStructureMotionTriangleGeometryDescriptorClass) Alloc() AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Call[AccelerationStructureMotionTriangleGeometryDescriptor](ac, objc.Sel("alloc"))
	return rv
}

func (ac _AccelerationStructureMotionTriangleGeometryDescriptorClass) New() AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Call[AccelerationStructureMotionTriangleGeometryDescriptor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAccelerationStructureMotionTriangleGeometryDescriptor() AccelerationStructureMotionTriangleGeometryDescriptor {
	return AccelerationStructureMotionTriangleGeometryDescriptorClass.New()
}

func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) Init() AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Call[AccelerationStructureMotionTriangleGeometryDescriptor](a_, objc.Sel("init"))
	return rv
}

// The data type of indices in the index buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/3750486-indextype?language=objc
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) IndexType() IndexType {
	rv := objc.Call[IndexType](a_, objc.Sel("indexType"))
	return rv
}

// The data type of indices in the index buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/3750486-indextype?language=objc
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Call[objc.Void](a_, objc.Sel("setIndexType:"), value)
}

// A buffer that contains indices for the vertices that compose the triangle list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/3750484-indexbuffer?language=objc
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) IndexBuffer() BufferObject {
	rv := objc.Call[BufferObject](a_, objc.Sel("indexBuffer"))
	return rv
}

// A buffer that contains indices for the vertices that compose the triangle list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/3750484-indexbuffer?language=objc
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexBuffer(value PBuffer) {
	po0 := objc.WrapAsProtocol("MTLBuffer", value)
	objc.Call[objc.Void](a_, objc.Sel("setIndexBuffer:"), po0)
}

// A buffer that contains indices for the vertices that compose the triangle list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/3750484-indexbuffer?language=objc
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexBufferObject(valueObject objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setIndexBuffer:"), valueObject)
}

// The number of triangles in the buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/3750487-trianglecount?language=objc
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) TriangleCount() uint {
	rv := objc.Call[uint](a_, objc.Sel("triangleCount"))
	return rv
}

// The number of triangles in the buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/3750487-trianglecount?language=objc
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetTriangleCount(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setTriangleCount:"), value)
}

// The offset, in bytes, to the first index in the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/3750485-indexbufferoffset?language=objc
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) IndexBufferOffset() uint {
	rv := objc.Call[uint](a_, objc.Sel("indexBufferOffset"))
	return rv
}

// The offset, in bytes, to the first index in the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/3750485-indexbufferoffset?language=objc
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexBufferOffset(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setIndexBufferOffset:"), value)
}

// An array of motion keyframes, each containing triangle data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/3750488-vertexbuffers?language=objc
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) VertexBuffers() []MotionKeyframeData {
	rv := objc.Call[[]MotionKeyframeData](a_, objc.Sel("vertexBuffers"))
	return rv
}

// An array of motion keyframes, each containing triangle data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/3750488-vertexbuffers?language=objc
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexBuffers(value []IMotionKeyframeData) {
	objc.Call[objc.Void](a_, objc.Sel("setVertexBuffers:"), value)
}

// The stride, in bytes, between vertices in each vertex buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/3750489-vertexstride?language=objc
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) VertexStride() uint {
	rv := objc.Call[uint](a_, objc.Sel("vertexStride"))
	return rv
}

// The stride, in bytes, between vertices in each vertex buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/3750489-vertexstride?language=objc
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexStride(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setVertexStride:"), value)
}
