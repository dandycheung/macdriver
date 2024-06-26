// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CNNNeuronGradientNode] class.
var CNNNeuronGradientNodeClass = _CNNNeuronGradientNodeClass{objc.GetClass("MPSCNNNeuronGradientNode")}

type _CNNNeuronGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronGradientNode] class.
type ICNNNeuronGradientNode interface {
	INNGradientFilterNode
	Descriptor() NNNeuronDescriptor
}

// A representation of a gradient exponential neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradientnode?language=objc
type CNNNeuronGradientNode struct {
	NNGradientFilterNode
}

func CNNNeuronGradientNodeFrom(ptr unsafe.Pointer) CNNNeuronGradientNode {
	return CNNNeuronGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (cc _CNNNeuronGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateDescriptor(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, descriptor INNNeuronDescriptor) CNNNeuronGradientNode {
	rv := objc.Call[CNNNeuronGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:descriptor:"), sourceGradient, sourceImage, gradientState, descriptor)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradientnode/2948028-nodewithsourcegradient?language=objc
func CNNNeuronGradientNode_NodeWithSourceGradientSourceImageGradientStateDescriptor(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, descriptor INNNeuronDescriptor) CNNNeuronGradientNode {
	return CNNNeuronGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateDescriptor(sourceGradient, sourceImage, gradientState, descriptor)
}

func (c_ CNNNeuronGradientNode) InitWithSourceGradientSourceImageGradientStateDescriptor(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, descriptor INNNeuronDescriptor) CNNNeuronGradientNode {
	rv := objc.Call[CNNNeuronGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:descriptor:"), sourceGradient, sourceImage, gradientState, descriptor)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradientnode/2948031-initwithsourcegradient?language=objc
func NewCNNNeuronGradientNodeWithSourceGradientSourceImageGradientStateDescriptor(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, descriptor INNNeuronDescriptor) CNNNeuronGradientNode {
	instance := CNNNeuronGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateDescriptor(sourceGradient, sourceImage, gradientState, descriptor)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronGradientNodeClass) Alloc() CNNNeuronGradientNode {
	rv := objc.Call[CNNNeuronGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNNeuronGradientNodeClass) New() CNNNeuronGradientNode {
	rv := objc.Call[CNNNeuronGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronGradientNode() CNNNeuronGradientNode {
	return CNNNeuronGradientNodeClass.New()
}

func (c_ CNNNeuronGradientNode) Init() CNNNeuronGradientNode {
	rv := objc.Call[CNNNeuronGradientNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradientnode/2948040-descriptor?language=objc
func (c_ CNNNeuronGradientNode) Descriptor() NNNeuronDescriptor {
	rv := objc.Call[NNNeuronDescriptor](c_, objc.Sel("descriptor"))
	return rv
}
