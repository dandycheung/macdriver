// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient?language=objc
type PTextCheckingClient interface {
	// optional
	SelectAndShowRange(range_ foundation.Range)
	HasSelectAndShowRange() bool

	// optional
	ViewForRangeFirstRectActualRange(range_ foundation.Range, firstRect foundation.RectPointer, actualRange foundation.RangePointer) View
	HasViewForRangeFirstRectActualRange() bool

	// optional
	AnnotatedSubstringForProposedRangeActualRange(range_ foundation.Range, actualRange foundation.RangePointer) foundation.AttributedString
	HasAnnotatedSubstringForProposedRangeActualRange() bool

	// optional
	ReplaceCharactersInRangeWithAnnotatedString(range_ foundation.Range, annotatedString foundation.AttributedString)
	HasReplaceCharactersInRangeWithAnnotatedString() bool

	// optional
	AddAnnotationsRange(annotations map[foundation.AttributedStringKey]string, range_ foundation.Range)
	HasAddAnnotationsRange() bool

	// optional
	SetAnnotationsRange(annotations map[foundation.AttributedStringKey]string, range_ foundation.Range)
	HasSetAnnotationsRange() bool

	// optional
	RemoveAnnotationRange(annotationName foundation.AttributedStringKey, range_ foundation.Range)
	HasRemoveAnnotationRange() bool

	// optional
	CandidateListTouchBarItem() CandidateListTouchBarItem
	HasCandidateListTouchBarItem() bool
}

// ensure impl type implements protocol interface
var _ PTextCheckingClient = (*TextCheckingClientObject)(nil)

// A concrete type for the [PTextCheckingClient] protocol.
type TextCheckingClientObject struct {
	objc.Object
}

func (t_ TextCheckingClientObject) HasSelectAndShowRange() bool {
	return t_.RespondsToSelector(objc.Sel("selectAndShowRange:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242724-selectandshowrange?language=objc
func (t_ TextCheckingClientObject) SelectAndShowRange(range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("selectAndShowRange:"), range_)
}

func (t_ TextCheckingClientObject) HasViewForRangeFirstRectActualRange() bool {
	return t_.RespondsToSelector(objc.Sel("viewForRange:firstRect:actualRange:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242726-viewforrange?language=objc
func (t_ TextCheckingClientObject) ViewForRangeFirstRectActualRange(range_ foundation.Range, firstRect foundation.RectPointer, actualRange foundation.RangePointer) View {
	rv := objc.Call[View](t_, objc.Sel("viewForRange:firstRect:actualRange:"), range_, firstRect, actualRange)
	return rv
}

func (t_ TextCheckingClientObject) HasAnnotatedSubstringForProposedRangeActualRange() bool {
	return t_.RespondsToSelector(objc.Sel("annotatedSubstringForProposedRange:actualRange:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242720-annotatedsubstringforproposedran?language=objc
func (t_ TextCheckingClientObject) AnnotatedSubstringForProposedRangeActualRange(range_ foundation.Range, actualRange foundation.RangePointer) foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](t_, objc.Sel("annotatedSubstringForProposedRange:actualRange:"), range_, actualRange)
	return rv
}

func (t_ TextCheckingClientObject) HasReplaceCharactersInRangeWithAnnotatedString() bool {
	return t_.RespondsToSelector(objc.Sel("replaceCharactersInRange:withAnnotatedString:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242723-replacecharactersinrange?language=objc
func (t_ TextCheckingClientObject) ReplaceCharactersInRangeWithAnnotatedString(range_ foundation.Range, annotatedString foundation.AttributedString) {
	objc.Call[objc.Void](t_, objc.Sel("replaceCharactersInRange:withAnnotatedString:"), range_, annotatedString)
}

func (t_ TextCheckingClientObject) HasAddAnnotationsRange() bool {
	return t_.RespondsToSelector(objc.Sel("addAnnotations:range:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242719-addannotations?language=objc
func (t_ TextCheckingClientObject) AddAnnotationsRange(annotations map[foundation.AttributedStringKey]string, range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("addAnnotations:range:"), annotations, range_)
}

func (t_ TextCheckingClientObject) HasSetAnnotationsRange() bool {
	return t_.RespondsToSelector(objc.Sel("setAnnotations:range:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242725-setannotations?language=objc
func (t_ TextCheckingClientObject) SetAnnotationsRange(annotations map[foundation.AttributedStringKey]string, range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setAnnotations:range:"), annotations, range_)
}

func (t_ TextCheckingClientObject) HasRemoveAnnotationRange() bool {
	return t_.RespondsToSelector(objc.Sel("removeAnnotation:range:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242722-removeannotation?language=objc
func (t_ TextCheckingClientObject) RemoveAnnotationRange(annotationName foundation.AttributedStringKey, range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("removeAnnotation:range:"), annotationName, range_)
}

func (t_ TextCheckingClientObject) HasCandidateListTouchBarItem() bool {
	return t_.RespondsToSelector(objc.Sel("candidateListTouchBarItem"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242721-candidatelisttouchbaritem?language=objc
func (t_ TextCheckingClientObject) CandidateListTouchBarItem() CandidateListTouchBarItem {
	rv := objc.Call[CandidateListTouchBarItem](t_, objc.Sel("candidateListTouchBarItem"))
	return rv
}
