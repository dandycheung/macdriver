// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CNNNeuronReLUNNode] class.
var CNNNeuronReLUNNodeClass = _CNNNeuronReLUNNodeClass{objc.GetClass("MPSCNNNeuronReLUNNode")}

type _CNNNeuronReLUNNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronReLUNNode] class.
type ICNNNeuronReLUNNode interface {
	ICNNNeuronNode
}

// A representation a ReLUN neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunnode?language=objc
type CNNNeuronReLUNNode struct {
	CNNNeuronNode
}

func CNNNeuronReLUNNodeFrom(ptr unsafe.Pointer) CNNNeuronReLUNNode {
	return CNNNeuronReLUNNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (cc _CNNNeuronReLUNNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNeuronReLUNNode {
	rv := objc.Call[CNNNeuronReLUNNode](cc, objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunnode/2921594-nodewithsource?language=objc
func CNNNeuronReLUNNode_NodeWithSource(sourceNode INNImageNode) CNNNeuronReLUNNode {
	return CNNNeuronReLUNNodeClass.NodeWithSource(sourceNode)
}

func (cc _CNNNeuronReLUNNodeClass) NodeWithSourceAB(sourceNode INNImageNode, a float32, b float32) CNNNeuronReLUNNode {
	rv := objc.Call[CNNNeuronReLUNNode](cc, objc.Sel("nodeWithSource:a:b:"), sourceNode, a, b)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunnode/2921590-nodewithsource?language=objc
func CNNNeuronReLUNNode_NodeWithSourceAB(sourceNode INNImageNode, a float32, b float32) CNNNeuronReLUNNode {
	return CNNNeuronReLUNNodeClass.NodeWithSourceAB(sourceNode, a, b)
}

func (c_ CNNNeuronReLUNNode) InitWithSourceAB(sourceNode INNImageNode, a float32, b float32) CNNNeuronReLUNNode {
	rv := objc.Call[CNNNeuronReLUNNode](c_, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunnode/2921596-initwithsource?language=objc
func NewCNNNeuronReLUNNodeWithSourceAB(sourceNode INNImageNode, a float32, b float32) CNNNeuronReLUNNode {
	instance := CNNNeuronReLUNNodeClass.Alloc().InitWithSourceAB(sourceNode, a, b)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronReLUNNode) InitWithSource(sourceNode INNImageNode) CNNNeuronReLUNNode {
	rv := objc.Call[CNNNeuronReLUNNode](c_, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunnode/2921593-initwithsource?language=objc
func NewCNNNeuronReLUNNodeWithSource(sourceNode INNImageNode) CNNNeuronReLUNNode {
	instance := CNNNeuronReLUNNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronReLUNNodeClass) Alloc() CNNNeuronReLUNNode {
	rv := objc.Call[CNNNeuronReLUNNode](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNNeuronReLUNNodeClass) New() CNNNeuronReLUNNode {
	rv := objc.Call[CNNNeuronReLUNNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronReLUNNode() CNNNeuronReLUNNode {
	return CNNNeuronReLUNNodeClass.New()
}

func (c_ CNNNeuronReLUNNode) Init() CNNNeuronReLUNNode {
	rv := objc.Call[CNNNeuronReLUNNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronReLUNNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronReLUNNode {
	rv := objc.Call[CNNNeuronReLUNNode](cc, objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronReLUNNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronReLUNNode {
	return CNNNeuronReLUNNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}
