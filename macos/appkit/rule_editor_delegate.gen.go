// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The NSRuleEditorDelegate protocol defines the optional methods implemented by delegates of NSRuleEditor objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate?language=objc
type PRuleEditorDelegate interface {
	// optional
	RuleEditorDisplayValueForCriterionInRow(editor RuleEditor, criterion objc.Object, row int) objc.Object
	HasRuleEditorDisplayValueForCriterionInRow() bool

	// optional
	RuleEditorPredicatePartsForCriterionWithDisplayValueInRow(editor RuleEditor, criterion objc.Object, value objc.Object, row int) map[RuleEditorPredicatePartKey]objc.Object
	HasRuleEditorPredicatePartsForCriterionWithDisplayValueInRow() bool

	// optional
	RuleEditorNumberOfChildrenForCriterionWithRowType(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int
	HasRuleEditorNumberOfChildrenForCriterionWithRowType() bool

	// optional
	RuleEditorChildForCriterionWithRowType(editor RuleEditor, index int, criterion objc.Object, rowType RuleEditorRowType) objc.Object
	HasRuleEditorChildForCriterionWithRowType() bool

	// optional
	RuleEditorRowsDidChange(notification foundation.Notification)
	HasRuleEditorRowsDidChange() bool
}

// A delegate implementation builder for the [PRuleEditorDelegate] protocol.
type RuleEditorDelegate struct {
	_RuleEditorDisplayValueForCriterionInRow                   func(editor RuleEditor, criterion objc.Object, row int) objc.Object
	_RuleEditorPredicatePartsForCriterionWithDisplayValueInRow func(editor RuleEditor, criterion objc.Object, value objc.Object, row int) map[RuleEditorPredicatePartKey]objc.Object
	_RuleEditorNumberOfChildrenForCriterionWithRowType         func(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int
	_RuleEditorChildForCriterionWithRowType                    func(editor RuleEditor, index int, criterion objc.Object, rowType RuleEditorRowType) objc.Object
	_RuleEditorRowsDidChange                                   func(notification foundation.Notification)
}

func (di *RuleEditorDelegate) HasRuleEditorDisplayValueForCriterionInRow() bool {
	return di._RuleEditorDisplayValueForCriterionInRow != nil
}

// Returns the value for a given criterion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1534788-ruleeditor?language=objc
func (di *RuleEditorDelegate) SetRuleEditorDisplayValueForCriterionInRow(f func(editor RuleEditor, criterion objc.Object, row int) objc.Object) {
	di._RuleEditorDisplayValueForCriterionInRow = f
}

// Returns the value for a given criterion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1534788-ruleeditor?language=objc
func (di *RuleEditorDelegate) RuleEditorDisplayValueForCriterionInRow(editor RuleEditor, criterion objc.Object, row int) objc.Object {
	return di._RuleEditorDisplayValueForCriterionInRow(editor, criterion, row)
}
func (di *RuleEditorDelegate) HasRuleEditorPredicatePartsForCriterionWithDisplayValueInRow() bool {
	return di._RuleEditorPredicatePartsForCriterionWithDisplayValueInRow != nil
}

// Returns a dictionary representing the parts of the predicate determined by the given criterion and value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1526667-ruleeditor?language=objc
func (di *RuleEditorDelegate) SetRuleEditorPredicatePartsForCriterionWithDisplayValueInRow(f func(editor RuleEditor, criterion objc.Object, value objc.Object, row int) map[RuleEditorPredicatePartKey]objc.Object) {
	di._RuleEditorPredicatePartsForCriterionWithDisplayValueInRow = f
}

// Returns a dictionary representing the parts of the predicate determined by the given criterion and value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1526667-ruleeditor?language=objc
func (di *RuleEditorDelegate) RuleEditorPredicatePartsForCriterionWithDisplayValueInRow(editor RuleEditor, criterion objc.Object, value objc.Object, row int) map[RuleEditorPredicatePartKey]objc.Object {
	return di._RuleEditorPredicatePartsForCriterionWithDisplayValueInRow(editor, criterion, value, row)
}
func (di *RuleEditorDelegate) HasRuleEditorNumberOfChildrenForCriterionWithRowType() bool {
	return di._RuleEditorNumberOfChildrenForCriterionWithRowType != nil
}

// Returns the number of child items of a given criterion or row type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1535329-ruleeditor?language=objc
func (di *RuleEditorDelegate) SetRuleEditorNumberOfChildrenForCriterionWithRowType(f func(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int) {
	di._RuleEditorNumberOfChildrenForCriterionWithRowType = f
}

// Returns the number of child items of a given criterion or row type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1535329-ruleeditor?language=objc
func (di *RuleEditorDelegate) RuleEditorNumberOfChildrenForCriterionWithRowType(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int {
	return di._RuleEditorNumberOfChildrenForCriterionWithRowType(editor, criterion, rowType)
}
func (di *RuleEditorDelegate) HasRuleEditorChildForCriterionWithRowType() bool {
	return di._RuleEditorChildForCriterionWithRowType != nil
}

// Returns the child of a given item at a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1527259-ruleeditor?language=objc
func (di *RuleEditorDelegate) SetRuleEditorChildForCriterionWithRowType(f func(editor RuleEditor, index int, criterion objc.Object, rowType RuleEditorRowType) objc.Object) {
	di._RuleEditorChildForCriterionWithRowType = f
}

// Returns the child of a given item at a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1527259-ruleeditor?language=objc
func (di *RuleEditorDelegate) RuleEditorChildForCriterionWithRowType(editor RuleEditor, index int, criterion objc.Object, rowType RuleEditorRowType) objc.Object {
	return di._RuleEditorChildForCriterionWithRowType(editor, index, criterion, rowType)
}
func (di *RuleEditorDelegate) HasRuleEditorRowsDidChange() bool {
	return di._RuleEditorRowsDidChange != nil
}

// Notifies the receiver that a rule editor’s rows changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1533292-ruleeditorrowsdidchange?language=objc
func (di *RuleEditorDelegate) SetRuleEditorRowsDidChange(f func(notification foundation.Notification)) {
	di._RuleEditorRowsDidChange = f
}

// Notifies the receiver that a rule editor’s rows changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1533292-ruleeditorrowsdidchange?language=objc
func (di *RuleEditorDelegate) RuleEditorRowsDidChange(notification foundation.Notification) {
	di._RuleEditorRowsDidChange(notification)
}

// ensure impl type implements protocol interface
var _ PRuleEditorDelegate = (*RuleEditorDelegateObject)(nil)

// A concrete type for the [PRuleEditorDelegate] protocol.
type RuleEditorDelegateObject struct {
	objc.Object
}

func (r_ RuleEditorDelegateObject) HasRuleEditorDisplayValueForCriterionInRow() bool {
	return r_.RespondsToSelector(objc.Sel("ruleEditor:displayValueForCriterion:inRow:"))
}

// Returns the value for a given criterion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1534788-ruleeditor?language=objc
func (r_ RuleEditorDelegateObject) RuleEditorDisplayValueForCriterionInRow(editor RuleEditor, criterion objc.Object, row int) objc.Object {
	rv := objc.Call[objc.Object](r_, objc.Sel("ruleEditor:displayValueForCriterion:inRow:"), editor, criterion, row)
	return rv
}

func (r_ RuleEditorDelegateObject) HasRuleEditorPredicatePartsForCriterionWithDisplayValueInRow() bool {
	return r_.RespondsToSelector(objc.Sel("ruleEditor:predicatePartsForCriterion:withDisplayValue:inRow:"))
}

// Returns a dictionary representing the parts of the predicate determined by the given criterion and value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1526667-ruleeditor?language=objc
func (r_ RuleEditorDelegateObject) RuleEditorPredicatePartsForCriterionWithDisplayValueInRow(editor RuleEditor, criterion objc.Object, value objc.Object, row int) map[RuleEditorPredicatePartKey]objc.Object {
	rv := objc.Call[map[RuleEditorPredicatePartKey]objc.Object](r_, objc.Sel("ruleEditor:predicatePartsForCriterion:withDisplayValue:inRow:"), editor, criterion, value, row)
	return rv
}

func (r_ RuleEditorDelegateObject) HasRuleEditorNumberOfChildrenForCriterionWithRowType() bool {
	return r_.RespondsToSelector(objc.Sel("ruleEditor:numberOfChildrenForCriterion:withRowType:"))
}

// Returns the number of child items of a given criterion or row type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1535329-ruleeditor?language=objc
func (r_ RuleEditorDelegateObject) RuleEditorNumberOfChildrenForCriterionWithRowType(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int {
	rv := objc.Call[int](r_, objc.Sel("ruleEditor:numberOfChildrenForCriterion:withRowType:"), editor, criterion, rowType)
	return rv
}

func (r_ RuleEditorDelegateObject) HasRuleEditorChildForCriterionWithRowType() bool {
	return r_.RespondsToSelector(objc.Sel("ruleEditor:child:forCriterion:withRowType:"))
}

// Returns the child of a given item at a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1527259-ruleeditor?language=objc
func (r_ RuleEditorDelegateObject) RuleEditorChildForCriterionWithRowType(editor RuleEditor, index int, criterion objc.Object, rowType RuleEditorRowType) objc.Object {
	rv := objc.Call[objc.Object](r_, objc.Sel("ruleEditor:child:forCriterion:withRowType:"), editor, index, criterion, rowType)
	return rv
}

func (r_ RuleEditorDelegateObject) HasRuleEditorRowsDidChange() bool {
	return r_.RespondsToSelector(objc.Sel("ruleEditorRowsDidChange:"))
}

// Notifies the receiver that a rule editor’s rows changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1533292-ruleeditorrowsdidchange?language=objc
func (r_ RuleEditorDelegateObject) RuleEditorRowsDidChange(notification foundation.Notification) {
	objc.Call[objc.Void](r_, objc.Sel("ruleEditorRowsDidChange:"), notification)
}
