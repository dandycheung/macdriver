// Code generated by DarwinKit. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ReplicatorLayer] class.
var ReplicatorLayerClass = _ReplicatorLayerClass{objc.GetClass("CAReplicatorLayer")}

type _ReplicatorLayerClass struct {
	objc.Class
}

// An interface definition for the [ReplicatorLayer] class.
type IReplicatorLayer interface {
	ILayer
	InstanceRedOffset() float32
	SetInstanceRedOffset(value float32)
	InstanceBlueOffset() float32
	SetInstanceBlueOffset(value float32)
	InstanceDelay() corefoundation.TimeInterval
	SetInstanceDelay(value corefoundation.TimeInterval)
	InstanceCount() int
	SetInstanceCount(value int)
	InstanceAlphaOffset() float32
	SetInstanceAlphaOffset(value float32)
	InstanceColor() coregraphics.ColorRef
	SetInstanceColor(value coregraphics.ColorRef)
	PreservesDepth() bool
	SetPreservesDepth(value bool)
	InstanceGreenOffset() float32
	SetInstanceGreenOffset(value float32)
	InstanceTransform() Transform3D
	SetInstanceTransform(value Transform3D)
}

// A layer that creates a specified number of sublayer copies with varying geometric, temporal, and color transformations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer?language=objc
type ReplicatorLayer struct {
	Layer
}

func ReplicatorLayerFrom(ptr unsafe.Pointer) ReplicatorLayer {
	return ReplicatorLayer{
		Layer: LayerFrom(ptr),
	}
}

func (rc _ReplicatorLayerClass) Alloc() ReplicatorLayer {
	rv := objc.Call[ReplicatorLayer](rc, objc.Sel("alloc"))
	return rv
}

func (rc _ReplicatorLayerClass) New() ReplicatorLayer {
	rv := objc.Call[ReplicatorLayer](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewReplicatorLayer() ReplicatorLayer {
	return ReplicatorLayerClass.New()
}

func (r_ ReplicatorLayer) Init() ReplicatorLayer {
	rv := objc.Call[ReplicatorLayer](r_, objc.Sel("init"))
	return rv
}

func (r_ ReplicatorLayer) ModelLayer() ReplicatorLayer {
	rv := objc.Call[ReplicatorLayer](r_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func ReplicatorLayer_ModelLayer() ReplicatorLayer {
	instance := ReplicatorLayerClass.Alloc().ModelLayer()
	instance.Autorelease()
	return instance
}

func (r_ ReplicatorLayer) PresentationLayer() ReplicatorLayer {
	rv := objc.Call[ReplicatorLayer](r_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func ReplicatorLayer_PresentationLayer() ReplicatorLayer {
	instance := ReplicatorLayerClass.Alloc().PresentationLayer()
	instance.Autorelease()
	return instance
}

func (r_ ReplicatorLayer) InitWithLayer(layer objc.IObject) ReplicatorLayer {
	rv := objc.Call[ReplicatorLayer](r_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func NewReplicatorLayerWithLayer(layer objc.IObject) ReplicatorLayer {
	instance := ReplicatorLayerClass.Alloc().InitWithLayer(layer)
	instance.Autorelease()
	return instance
}

func (rc _ReplicatorLayerClass) Layer() ReplicatorLayer {
	rv := objc.Call[ReplicatorLayer](rc, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func ReplicatorLayer_Layer() ReplicatorLayer {
	return ReplicatorLayerClass.Layer()
}

// Defines the offset added to the red component of the color for each replicated instance. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1522090-instanceredoffset?language=objc
func (r_ ReplicatorLayer) InstanceRedOffset() float32 {
	rv := objc.Call[float32](r_, objc.Sel("instanceRedOffset"))
	return rv
}

// Defines the offset added to the red component of the color for each replicated instance. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1522090-instanceredoffset?language=objc
func (r_ ReplicatorLayer) SetInstanceRedOffset(value float32) {
	objc.Call[objc.Void](r_, objc.Sel("setInstanceRedOffset:"), value)
}

// Defines the offset added to the blue component of the color for each replicated instance. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1522267-instanceblueoffset?language=objc
func (r_ ReplicatorLayer) InstanceBlueOffset() float32 {
	rv := objc.Call[float32](r_, objc.Sel("instanceBlueOffset"))
	return rv
}

// Defines the offset added to the blue component of the color for each replicated instance. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1522267-instanceblueoffset?language=objc
func (r_ ReplicatorLayer) SetInstanceBlueOffset(value float32) {
	objc.Call[objc.Void](r_, objc.Sel("setInstanceBlueOffset:"), value)
}

// Specifies the delay, in seconds, between replicated copies. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1522391-instancedelay?language=objc
func (r_ ReplicatorLayer) InstanceDelay() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](r_, objc.Sel("instanceDelay"))
	return rv
}

// Specifies the delay, in seconds, between replicated copies. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1522391-instancedelay?language=objc
func (r_ ReplicatorLayer) SetInstanceDelay(value corefoundation.TimeInterval) {
	objc.Call[objc.Void](r_, objc.Sel("setInstanceDelay:"), value)
}

// The number of copies to create, including the source layers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1521883-instancecount?language=objc
func (r_ ReplicatorLayer) InstanceCount() int {
	rv := objc.Call[int](r_, objc.Sel("instanceCount"))
	return rv
}

// The number of copies to create, including the source layers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1521883-instancecount?language=objc
func (r_ ReplicatorLayer) SetInstanceCount(value int) {
	objc.Call[objc.Void](r_, objc.Sel("setInstanceCount:"), value)
}

// Defines the offset added to the alpha component of the color for each replicated instance. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1521898-instancealphaoffset?language=objc
func (r_ ReplicatorLayer) InstanceAlphaOffset() float32 {
	rv := objc.Call[float32](r_, objc.Sel("instanceAlphaOffset"))
	return rv
}

// Defines the offset added to the alpha component of the color for each replicated instance. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1521898-instancealphaoffset?language=objc
func (r_ ReplicatorLayer) SetInstanceAlphaOffset(value float32) {
	objc.Call[objc.Void](r_, objc.Sel("setInstanceAlphaOffset:"), value)
}

// Defines the color used to multiply the source object. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1522154-instancecolor?language=objc
func (r_ ReplicatorLayer) InstanceColor() coregraphics.ColorRef {
	rv := objc.Call[coregraphics.ColorRef](r_, objc.Sel("instanceColor"))
	return rv
}

// Defines the color used to multiply the source object. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1522154-instancecolor?language=objc
func (r_ ReplicatorLayer) SetInstanceColor(value coregraphics.ColorRef) {
	objc.Call[objc.Void](r_, objc.Sel("setInstanceColor:"), value)
}

// Defines whether this layer flattens its sublayers into its plane. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1522095-preservesdepth?language=objc
func (r_ ReplicatorLayer) PreservesDepth() bool {
	rv := objc.Call[bool](r_, objc.Sel("preservesDepth"))
	return rv
}

// Defines whether this layer flattens its sublayers into its plane. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1522095-preservesdepth?language=objc
func (r_ ReplicatorLayer) SetPreservesDepth(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setPreservesDepth:"), value)
}

// Defines the offset added to the green component of the color for each replicated instance. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1522032-instancegreenoffset?language=objc
func (r_ ReplicatorLayer) InstanceGreenOffset() float32 {
	rv := objc.Call[float32](r_, objc.Sel("instanceGreenOffset"))
	return rv
}

// Defines the offset added to the green component of the color for each replicated instance. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1522032-instancegreenoffset?language=objc
func (r_ ReplicatorLayer) SetInstanceGreenOffset(value float32) {
	objc.Call[objc.Void](r_, objc.Sel("setInstanceGreenOffset:"), value)
}

// The transform matrix applied to the previous instance to produce the current instance. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1522312-instancetransform?language=objc
func (r_ ReplicatorLayer) InstanceTransform() Transform3D {
	rv := objc.Call[Transform3D](r_, objc.Sel("instanceTransform"))
	return rv
}

// The transform matrix applied to the previous instance to produce the current instance. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/1522312-instancetransform?language=objc
func (r_ ReplicatorLayer) SetInstanceTransform(value Transform3D) {
	objc.Call[objc.Void](r_, objc.Sel("setInstanceTransform:"), value)
}
