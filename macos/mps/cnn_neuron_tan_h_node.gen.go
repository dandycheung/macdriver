// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronTanHNode] class.
var CNNNeuronTanHNodeClass = _CNNNeuronTanHNodeClass{objc.GetClass("MPSCNNNeuronTanHNode")}

type _CNNNeuronTanHNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronTanHNode] class.
type ICNNNeuronTanHNode interface {
	ICNNNeuronNode
}

// A representation of a hyperbolic tangent neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontanhnode?language=objc
type CNNNeuronTanHNode struct {
	CNNNeuronNode
}

func CNNNeuronTanHNodeFrom(ptr unsafe.Pointer) CNNNeuronTanHNode {
	return CNNNeuronTanHNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (c_ CNNNeuronTanHNode) InitWithSource(sourceNode INNImageNode) CNNNeuronTanHNode {
	rv := objc.Call[CNNNeuronTanHNode](c_, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontanhnode/2921465-initwithsource?language=objc
func NewCNNNeuronTanHNodeWithSource(sourceNode INNImageNode) CNNNeuronTanHNode {
	instance := CNNNeuronTanHNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronTanHNode) InitWithSourceAB(sourceNode INNImageNode, a float32, b float32) CNNNeuronTanHNode {
	rv := objc.Call[CNNNeuronTanHNode](c_, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontanhnode/2866481-initwithsource?language=objc
func NewCNNNeuronTanHNodeWithSourceAB(sourceNode INNImageNode, a float32, b float32) CNNNeuronTanHNode {
	instance := CNNNeuronTanHNodeClass.Alloc().InitWithSourceAB(sourceNode, a, b)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronTanHNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNeuronTanHNode {
	rv := objc.Call[CNNNeuronTanHNode](cc, objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontanhnode/2921451-nodewithsource?language=objc
func CNNNeuronTanHNode_NodeWithSource(sourceNode INNImageNode) CNNNeuronTanHNode {
	return CNNNeuronTanHNodeClass.NodeWithSource(sourceNode)
}

func (cc _CNNNeuronTanHNodeClass) NodeWithSourceAB(sourceNode INNImageNode, a float32, b float32) CNNNeuronTanHNode {
	rv := objc.Call[CNNNeuronTanHNode](cc, objc.Sel("nodeWithSource:a:b:"), sourceNode, a, b)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontanhnode/2866416-nodewithsource?language=objc
func CNNNeuronTanHNode_NodeWithSourceAB(sourceNode INNImageNode, a float32, b float32) CNNNeuronTanHNode {
	return CNNNeuronTanHNodeClass.NodeWithSourceAB(sourceNode, a, b)
}

func (cc _CNNNeuronTanHNodeClass) Alloc() CNNNeuronTanHNode {
	rv := objc.Call[CNNNeuronTanHNode](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNNeuronTanHNodeClass) New() CNNNeuronTanHNode {
	rv := objc.Call[CNNNeuronTanHNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronTanHNode() CNNNeuronTanHNode {
	return CNNNeuronTanHNodeClass.New()
}

func (c_ CNNNeuronTanHNode) Init() CNNNeuronTanHNode {
	rv := objc.Call[CNNNeuronTanHNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronTanHNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronTanHNode {
	rv := objc.Call[CNNNeuronTanHNode](cc, objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronTanHNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronTanHNode {
	return CNNNeuronTanHNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}
