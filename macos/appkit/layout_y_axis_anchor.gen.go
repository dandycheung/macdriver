// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LayoutYAxisAnchor] class.
var LayoutYAxisAnchorClass = _LayoutYAxisAnchorClass{objc.GetClass("NSLayoutYAxisAnchor")}

type _LayoutYAxisAnchorClass struct {
	objc.Class
}

// An interface definition for the [LayoutYAxisAnchor] class.
type ILayoutYAxisAnchor interface {
	ILayoutAnchor
	ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintLessThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint
	AnchorWithOffsetToAnchor(otherAnchor ILayoutYAxisAnchor) LayoutDimension
	ConstraintEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint
}

// A factory class for creating vertical layout constraint objects using a fluent API. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutyaxisanchor?language=objc
type LayoutYAxisAnchor struct {
	LayoutAnchor
}

func LayoutYAxisAnchorFrom(ptr unsafe.Pointer) LayoutYAxisAnchor {
	return LayoutYAxisAnchor{
		LayoutAnchor: LayoutAnchorFrom(ptr),
	}
}

func (lc _LayoutYAxisAnchorClass) Alloc() LayoutYAxisAnchor {
	rv := objc.Call[LayoutYAxisAnchor](lc, objc.Sel("alloc"))
	return rv
}

func (lc _LayoutYAxisAnchorClass) New() LayoutYAxisAnchor {
	rv := objc.Call[LayoutYAxisAnchor](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutYAxisAnchor() LayoutYAxisAnchor {
	return LayoutYAxisAnchorClass.New()
}

func (l_ LayoutYAxisAnchor) Init() LayoutYAxisAnchor {
	rv := objc.Call[LayoutYAxisAnchor](l_, objc.Sel("init"))
	return rv
}

// Returns a constraint that defines the minimum distance by which the current anchor is positioned below the specified anchor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutyaxisanchor/2866086-constraintgreaterthanorequaltosy?language=objc
func (l_ LayoutYAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintGreaterThanOrEqualToSystemSpacingBelowAnchor:multiplier:"), anchor, multiplier)
	return rv
}

// Returns a constraint that defines the maximum distance by which the current anchor is positioned below the specified anchor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutyaxisanchor/2865829-constraintlessthanorequaltosyste?language=objc
func (l_ LayoutYAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintLessThanOrEqualToSystemSpacingBelowAnchor:multiplier:"), anchor, multiplier)
	return rv
}

// Creates a layout dimension object from two anchors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutyaxisanchor/2865935-anchorwithoffsettoanchor?language=objc
func (l_ LayoutYAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor ILayoutYAxisAnchor) LayoutDimension {
	rv := objc.Call[LayoutDimension](l_, objc.Sel("anchorWithOffsetToAnchor:"), otherAnchor)
	return rv
}

// Returns a constraint that defines the specific distance at which the current anchor is positioned below the specified anchor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutyaxisanchor/2866022-constraintequaltosystemspacingbe?language=objc
func (l_ LayoutYAxisAnchor) ConstraintEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintEqualToSystemSpacingBelowAnchor:multiplier:"), anchor, multiplier)
	return rv
}
