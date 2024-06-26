// Code generated by DarwinKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [DOMXPathExpression] class.
var DOMXPathExpressionClass = _DOMXPathExpressionClass{objc.GetClass("DOMXPathExpression")}

type _DOMXPathExpressionClass struct {
	objc.Class
}

// An interface definition for the [DOMXPathExpression] class.
type IDOMXPathExpression interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domxpathexpression?language=objc
type DOMXPathExpression struct {
	DOMObject
}

func DOMXPathExpressionFrom(ptr unsafe.Pointer) DOMXPathExpression {
	return DOMXPathExpression{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMXPathExpressionClass) Alloc() DOMXPathExpression {
	rv := objc.Call[DOMXPathExpression](dc, objc.Sel("alloc"))
	return rv
}

func (dc _DOMXPathExpressionClass) New() DOMXPathExpression {
	rv := objc.Call[DOMXPathExpression](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMXPathExpression() DOMXPathExpression {
	return DOMXPathExpressionClass.New()
}

func (d_ DOMXPathExpression) Init() DOMXPathExpression {
	rv := objc.Call[DOMXPathExpression](d_, objc.Sel("init"))
	return rv
}
