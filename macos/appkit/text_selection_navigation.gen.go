// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextSelectionNavigation] class.
var TextSelectionNavigationClass = _TextSelectionNavigationClass{objc.GetClass("NSTextSelectionNavigation")}

type _TextSelectionNavigationClass struct {
	objc.Class
}

// An interface definition for the [TextSelectionNavigation] class.
type ITextSelectionNavigation interface {
	objc.IObject
	DestinationSelectionForTextSelectionDirectionDestinationExtendingConfined(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, extending bool, confined bool) TextSelection
	TextSelectionsInteractingAtPointInContainerAtLocationAnchorsModifiersSelectingBounds(point coregraphics.Point, containerLocation PTextLocation, anchors []ITextSelection, modifiers TextSelectionNavigationModifier, selecting bool, bounds coregraphics.Rect) []TextSelection
	TextSelectionsInteractingAtPointInContainerAtLocationObjectAnchorsModifiersSelectingBounds(point coregraphics.Point, containerLocationObject objc.IObject, anchors []ITextSelection, modifiers TextSelectionNavigationModifier, selecting bool, bounds coregraphics.Rect) []TextSelection
	TextSelectionForSelectionGranularityEnclosingPointInContainerAtLocation(selectionGranularity TextSelectionGranularity, point coregraphics.Point, location PTextLocation) TextSelection
	TextSelectionForSelectionGranularityEnclosingPointInContainerAtLocationObject(selectionGranularity TextSelectionGranularity, point coregraphics.Point, locationObject objc.IObject) TextSelection
	TextSelectionForSelectionGranularityEnclosingTextSelection(selectionGranularity TextSelectionGranularity, textSelection ITextSelection) TextSelection
	FlushLayoutCache()
	DeletionRangesForTextSelectionDirectionDestinationAllowsDecomposition(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, allowsDecomposition bool) []TextRange
	ResolvedInsertionLocationForTextSelectionWritingDirection(textSelection ITextSelection, writingDirection TextSelectionNavigationWritingDirection) TextLocationObject
	AllowsNonContiguousRanges() bool
	SetAllowsNonContiguousRanges(value bool)
	RotatesCoordinateSystemForLayoutOrientation() bool
	SetRotatesCoordinateSystemForLayoutOrientation(value bool)
	TextSelectionDataSource() TextSelectionDataSourceObject
}

// An interface used to expose methods for obtaining results from actions performed on text selections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation?language=objc
type TextSelectionNavigation struct {
	objc.Object
}

func TextSelectionNavigationFrom(ptr unsafe.Pointer) TextSelectionNavigation {
	return TextSelectionNavigation{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextSelectionNavigation) InitWithDataSource(dataSource PTextSelectionDataSource) TextSelectionNavigation {
	po0 := objc.WrapAsProtocol("NSTextSelectionDataSource", dataSource)
	rv := objc.Call[TextSelectionNavigation](t_, objc.Sel("initWithDataSource:"), po0)
	return rv
}

// Creates a new object using the text selection data source you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation/3801854-initwithdatasource?language=objc
func NewTextSelectionNavigationWithDataSource(dataSource PTextSelectionDataSource) TextSelectionNavigation {
	instance := TextSelectionNavigationClass.Alloc().InitWithDataSource(dataSource)
	instance.Autorelease()
	return instance
}

func (tc _TextSelectionNavigationClass) Alloc() TextSelectionNavigation {
	rv := objc.Call[TextSelectionNavigation](tc, objc.Sel("alloc"))
	return rv
}

func (tc _TextSelectionNavigationClass) New() TextSelectionNavigation {
	rv := objc.Call[TextSelectionNavigation](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextSelectionNavigation() TextSelectionNavigation {
	return TextSelectionNavigationClass.New()
}

func (t_ TextSelectionNavigation) Init() TextSelectionNavigation {
	rv := objc.Call[TextSelectionNavigation](t_, objc.Sel("init"))
	return rv
}

// Returns a new selection that results from applying the navigation operations you specify to the text selection you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation/3801852-destinationselectionfortextselec?language=objc
func (t_ TextSelectionNavigation) DestinationSelectionForTextSelectionDirectionDestinationExtendingConfined(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, extending bool, confined bool) TextSelection {
	rv := objc.Call[TextSelection](t_, objc.Sel("destinationSelectionForTextSelection:direction:destination:extending:confined:"), textSelection, direction, destination, extending, confined)
	return rv
}

// Returns an array of text selections produced by a tap or click at the point you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation/3801869-textselectionsinteractingatpoint?language=objc
func (t_ TextSelectionNavigation) TextSelectionsInteractingAtPointInContainerAtLocationAnchorsModifiersSelectingBounds(point coregraphics.Point, containerLocation PTextLocation, anchors []ITextSelection, modifiers TextSelectionNavigationModifier, selecting bool, bounds coregraphics.Rect) []TextSelection {
	po1 := objc.WrapAsProtocol("NSTextLocation", containerLocation)
	rv := objc.Call[[]TextSelection](t_, objc.Sel("textSelectionsInteractingAtPoint:inContainerAtLocation:anchors:modifiers:selecting:bounds:"), point, po1, anchors, modifiers, selecting, bounds)
	return rv
}

// Returns an array of text selections produced by a tap or click at the point you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation/3801869-textselectionsinteractingatpoint?language=objc
func (t_ TextSelectionNavigation) TextSelectionsInteractingAtPointInContainerAtLocationObjectAnchorsModifiersSelectingBounds(point coregraphics.Point, containerLocationObject objc.IObject, anchors []ITextSelection, modifiers TextSelectionNavigationModifier, selecting bool, bounds coregraphics.Rect) []TextSelection {
	rv := objc.Call[[]TextSelection](t_, objc.Sel("textSelectionsInteractingAtPoint:inContainerAtLocation:anchors:modifiers:selecting:bounds:"), point, containerLocationObject, anchors, modifiers, selecting, bounds)
	return rv
}

// Returns a text selection that expands to the nearest boundaries for selection granularity and an enclosing point you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation/3801858-textselectionforselectiongranula?language=objc
func (t_ TextSelectionNavigation) TextSelectionForSelectionGranularityEnclosingPointInContainerAtLocation(selectionGranularity TextSelectionGranularity, point coregraphics.Point, location PTextLocation) TextSelection {
	po2 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[TextSelection](t_, objc.Sel("textSelectionForSelectionGranularity:enclosingPoint:inContainerAtLocation:"), selectionGranularity, point, po2)
	return rv
}

// Returns a text selection that expands to the nearest boundaries for selection granularity and an enclosing point you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation/3801858-textselectionforselectiongranula?language=objc
func (t_ TextSelectionNavigation) TextSelectionForSelectionGranularityEnclosingPointInContainerAtLocationObject(selectionGranularity TextSelectionGranularity, point coregraphics.Point, locationObject objc.IObject) TextSelection {
	rv := objc.Call[TextSelection](t_, objc.Sel("textSelectionForSelectionGranularity:enclosingPoint:inContainerAtLocation:"), selectionGranularity, point, locationObject)
	return rv
}

// Returns a text selection expanded to the nearest boundaries for the selection granularity and enclosing text selection text ranges you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation/3801868-textselectionforselectiongranula?language=objc
func (t_ TextSelectionNavigation) TextSelectionForSelectionGranularityEnclosingTextSelection(selectionGranularity TextSelectionGranularity, textSelection ITextSelection) TextSelection {
	rv := objc.Call[TextSelection](t_, objc.Sel("textSelectionForSelectionGranularity:enclosingTextSelection:"), selectionGranularity, textSelection)
	return rv
}

// Flushes cached layout information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation/3801853-flushlayoutcache?language=objc
func (t_ TextSelectionNavigation) FlushLayoutCache() {
	objc.Call[objc.Void](t_, objc.Sel("flushLayoutCache"))
}

// Returns the ranges for deleting the text based on the current selection and movement arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation/3801851-deletionrangesfortextselection?language=objc
func (t_ TextSelectionNavigation) DeletionRangesForTextSelectionDirectionDestinationAllowsDecomposition(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, allowsDecomposition bool) []TextRange {
	rv := objc.Call[[]TextRange](t_, objc.Sel("deletionRangesForTextSelection:direction:destination:allowsDecomposition:"), textSelection, direction, destination, allowsDecomposition)
	return rv
}

// Returns the location for inserting the next input depending on the state of the current and secondary selections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation/3852581-resolvedinsertionlocationfortext?language=objc
func (t_ TextSelectionNavigation) ResolvedInsertionLocationForTextSelectionWritingDirection(textSelection ITextSelection, writingDirection TextSelectionNavigationWritingDirection) TextLocationObject {
	rv := objc.Call[TextLocationObject](t_, objc.Sel("resolvedInsertionLocationForTextSelection:writingDirection:"), textSelection, writingDirection)
	return rv
}

// Determines if the instance could produce selections with multiple noncontiguous selections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation/3801850-allowsnoncontiguousranges?language=objc
func (t_ TextSelectionNavigation) AllowsNonContiguousRanges() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsNonContiguousRanges"))
	return rv
}

// Determines if the instance could produce selections with multiple noncontiguous selections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation/3801850-allowsnoncontiguousranges?language=objc
func (t_ TextSelectionNavigation) SetAllowsNonContiguousRanges(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsNonContiguousRanges:"), value)
}

// Determines if the framework rotates the coordinate system to match the layout orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation/3801856-rotatescoordinatesystemforlayout?language=objc
func (t_ TextSelectionNavigation) RotatesCoordinateSystemForLayoutOrientation() bool {
	rv := objc.Call[bool](t_, objc.Sel("rotatesCoordinateSystemForLayoutOrientation"))
	return rv
}

// Determines if the framework rotates the coordinate system to match the layout orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation/3801856-rotatescoordinatesystemforlayout?language=objc
func (t_ TextSelectionNavigation) SetRotatesCoordinateSystemForLayoutOrientation(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setRotatesCoordinateSystemForLayoutOrientation:"), value)
}

// The data source associated with this selection navigation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigation/3801857-textselectiondatasource?language=objc
func (t_ TextSelectionNavigation) TextSelectionDataSource() TextSelectionDataSourceObject {
	rv := objc.Call[TextSelectionDataSourceObject](t_, objc.Sel("textSelectionDataSource"))
	return rv
}
