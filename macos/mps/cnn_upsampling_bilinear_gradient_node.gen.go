// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNUpsamplingBilinearGradientNode] class.
var CNNUpsamplingBilinearGradientNodeClass = _CNNUpsamplingBilinearGradientNodeClass{objc.GetClass("MPSCNNUpsamplingBilinearGradientNode")}

type _CNNUpsamplingBilinearGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNUpsamplingBilinearGradientNode] class.
type ICNNUpsamplingBilinearGradientNode interface {
	INNGradientFilterNode
	ScaleFactorY() float64
	ScaleFactorX() float64
}

// A representation of a gradient bilinear spatial upsampling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilineargradientnode?language=objc
type CNNUpsamplingBilinearGradientNode struct {
	NNGradientFilterNode
}

func CNNUpsamplingBilinearGradientNodeFrom(ptr unsafe.Pointer) CNNUpsamplingBilinearGradientNode {
	return CNNUpsamplingBilinearGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (c_ CNNUpsamplingBilinearGradientNode) InitWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, scaleFactorX float64, scaleFactorY float64) CNNUpsamplingBilinearGradientNode {
	rv := objc.Call[CNNUpsamplingBilinearGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), scaleFactorX, scaleFactorY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilineargradientnode/2947991-initwithsourcegradient?language=objc
func NewCNNUpsamplingBilinearGradientNodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, scaleFactorX float64, scaleFactorY float64) CNNUpsamplingBilinearGradientNode {
	instance := CNNUpsamplingBilinearGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
	instance.Autorelease()
	return instance
}

func (cc _CNNUpsamplingBilinearGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, scaleFactorX float64, scaleFactorY float64) CNNUpsamplingBilinearGradientNode {
	rv := objc.Call[CNNUpsamplingBilinearGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), scaleFactorX, scaleFactorY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilineargradientnode/2948025-nodewithsourcegradient?language=objc
func CNNUpsamplingBilinearGradientNode_NodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, scaleFactorX float64, scaleFactorY float64) CNNUpsamplingBilinearGradientNode {
	return CNNUpsamplingBilinearGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
}

func (cc _CNNUpsamplingBilinearGradientNodeClass) Alloc() CNNUpsamplingBilinearGradientNode {
	rv := objc.Call[CNNUpsamplingBilinearGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNUpsamplingBilinearGradientNode_Alloc() CNNUpsamplingBilinearGradientNode {
	return CNNUpsamplingBilinearGradientNodeClass.Alloc()
}

func (cc _CNNUpsamplingBilinearGradientNodeClass) New() CNNUpsamplingBilinearGradientNode {
	rv := objc.Call[CNNUpsamplingBilinearGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNUpsamplingBilinearGradientNode() CNNUpsamplingBilinearGradientNode {
	return CNNUpsamplingBilinearGradientNodeClass.New()
}

func (c_ CNNUpsamplingBilinearGradientNode) Init() CNNUpsamplingBilinearGradientNode {
	rv := objc.Call[CNNUpsamplingBilinearGradientNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilineargradientnode/2948054-scalefactory?language=objc
func (c_ CNNUpsamplingBilinearGradientNode) ScaleFactorY() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleFactorY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilineargradientnode/2948051-scalefactorx?language=objc
func (c_ CNNUpsamplingBilinearGradientNode) ScaleFactorX() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleFactorX"))
	return rv
}