// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [SliderTouchBarItem] class.
var SliderTouchBarItemClass = _SliderTouchBarItemClass{objc.GetClass("NSSliderTouchBarItem")}

type _SliderTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [SliderTouchBarItem] class.
type ISliderTouchBarItem interface {
	ITouchBarItem
	MinimumSliderWidth() float64
	SetMinimumSliderWidth(value float64)
	Slider() Slider
	SetSlider(value ISlider)
	DoubleValue() float64
	SetDoubleValue(value float64)
	Label() string
	SetLabel(value string)
	MaximumValueAccessory() SliderAccessory
	SetMaximumValueAccessory(value ISliderAccessory)
	MinimumValueAccessory() SliderAccessory
	SetMinimumValueAccessory(value ISliderAccessory)
	SetCustomizationLabel(value string)
	Target() objc.Object
	SetTarget(value objc.IObject)
	MaximumSliderWidth() float64
	SetMaximumSliderWidth(value float64)
	ValueAccessoryWidth() SliderAccessoryWidth
	SetValueAccessoryWidth(value SliderAccessoryWidth)
	Action() objc.Selector
	SetAction(value objc.Selector)
}

// A bar item that provides a slider control for choosing a value in a range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem?language=objc
type SliderTouchBarItem struct {
	TouchBarItem
}

func SliderTouchBarItemFrom(ptr unsafe.Pointer) SliderTouchBarItem {
	return SliderTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

func (sc _SliderTouchBarItemClass) Alloc() SliderTouchBarItem {
	rv := objc.Call[SliderTouchBarItem](sc, objc.Sel("alloc"))
	return rv
}

func (sc _SliderTouchBarItemClass) New() SliderTouchBarItem {
	rv := objc.Call[SliderTouchBarItem](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSliderTouchBarItem() SliderTouchBarItem {
	return SliderTouchBarItemClass.New()
}

func (s_ SliderTouchBarItem) Init() SliderTouchBarItem {
	rv := objc.Call[SliderTouchBarItem](s_, objc.Sel("init"))
	return rv
}

func (s_ SliderTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) SliderTouchBarItem {
	rv := objc.Call[SliderTouchBarItem](s_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Creates a new item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544755-initwithidentifier?language=objc
func NewSliderTouchBarItemWithIdentifier(identifier TouchBarItemIdentifier) SliderTouchBarItem {
	instance := SliderTouchBarItemClass.Alloc().InitWithIdentifier(identifier)
	instance.Autorelease()
	return instance
}

// The minimum width of the slider’s track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/3237222-minimumsliderwidth?language=objc
func (s_ SliderTouchBarItem) MinimumSliderWidth() float64 {
	rv := objc.Call[float64](s_, objc.Sel("minimumSliderWidth"))
	return rv
}

// The minimum width of the slider’s track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/3237222-minimumsliderwidth?language=objc
func (s_ SliderTouchBarItem) SetMinimumSliderWidth(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMinimumSliderWidth:"), value)
}

// The slider displayed by the bar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/2544809-slider?language=objc
func (s_ SliderTouchBarItem) Slider() Slider {
	rv := objc.Call[Slider](s_, objc.Sel("slider"))
	return rv
}

// The slider displayed by the bar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/2544809-slider?language=objc
func (s_ SliderTouchBarItem) SetSlider(value ISlider) {
	objc.Call[objc.Void](s_, objc.Sel("setSlider:"), value)
}

// The double value of the slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/3237220-doublevalue?language=objc
func (s_ SliderTouchBarItem) DoubleValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("doubleValue"))
	return rv
}

// The double value of the slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/3237220-doublevalue?language=objc
func (s_ SliderTouchBarItem) SetDoubleValue(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setDoubleValue:"), value)
}

// The text displayed alongside the slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/2544741-label?language=objc
func (s_ SliderTouchBarItem) Label() string {
	rv := objc.Call[string](s_, objc.Sel("label"))
	return rv
}

// The text displayed alongside the slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/2544741-label?language=objc
func (s_ SliderTouchBarItem) SetLabel(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setLabel:"), value)
}

// The accessory that appears at the end of the slider with the maximum value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/2544672-maximumvalueaccessory?language=objc
func (s_ SliderTouchBarItem) MaximumValueAccessory() SliderAccessory {
	rv := objc.Call[SliderAccessory](s_, objc.Sel("maximumValueAccessory"))
	return rv
}

// The accessory that appears at the end of the slider with the maximum value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/2544672-maximumvalueaccessory?language=objc
func (s_ SliderTouchBarItem) SetMaximumValueAccessory(value ISliderAccessory) {
	objc.Call[objc.Void](s_, objc.Sel("setMaximumValueAccessory:"), value)
}

// The accessory that appears at the end of the slider with the minimum value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/2544846-minimumvalueaccessory?language=objc
func (s_ SliderTouchBarItem) MinimumValueAccessory() SliderAccessory {
	rv := objc.Call[SliderAccessory](s_, objc.Sel("minimumValueAccessory"))
	return rv
}

// The accessory that appears at the end of the slider with the minimum value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/2544846-minimumvalueaccessory?language=objc
func (s_ SliderTouchBarItem) SetMinimumValueAccessory(value ISliderAccessory) {
	objc.Call[objc.Void](s_, objc.Sel("setMinimumValueAccessory:"), value)
}

// The user-visible string identifying this item during bar customization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/2544793-customizationlabel?language=objc
func (s_ SliderTouchBarItem) SetCustomizationLabel(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setCustomizationLabel:"), value)
}

// An object that is notified when a user interacts with the slider or either of the accessories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/2544814-target?language=objc
func (s_ SliderTouchBarItem) Target() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("target"))
	return rv
}

// An object that is notified when a user interacts with the slider or either of the accessories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/2544814-target?language=objc
func (s_ SliderTouchBarItem) SetTarget(value objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setTarget:"), value)
}

// The maximum width of the slider’s track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/3237221-maximumsliderwidth?language=objc
func (s_ SliderTouchBarItem) MaximumSliderWidth() float64 {
	rv := objc.Call[float64](s_, objc.Sel("maximumSliderWidth"))
	return rv
}

// The maximum width of the slider’s track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/3237221-maximumsliderwidth?language=objc
func (s_ SliderTouchBarItem) SetMaximumSliderWidth(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMaximumSliderWidth:"), value)
}

// The width of the value accessories that appear at either end of the slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/2544811-valueaccessorywidth?language=objc
func (s_ SliderTouchBarItem) ValueAccessoryWidth() SliderAccessoryWidth {
	rv := objc.Call[SliderAccessoryWidth](s_, objc.Sel("valueAccessoryWidth"))
	return rv
}

// The width of the value accessories that appear at either end of the slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/2544811-valueaccessorywidth?language=objc
func (s_ SliderTouchBarItem) SetValueAccessoryWidth(value SliderAccessoryWidth) {
	objc.Call[objc.Void](s_, objc.Sel("setValueAccessoryWidth:"), value)
}

// The selector on the target object that is invoked when a user interacts with the slider or either of the accessories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/2544765-action?language=objc
func (s_ SliderTouchBarItem) Action() objc.Selector {
	rv := objc.Call[objc.Selector](s_, objc.Sel("action"))
	return rv
}

// The selector on the target object that is invoked when a user interacts with the slider or either of the accessories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/2544765-action?language=objc
func (s_ SliderTouchBarItem) SetAction(value objc.Selector) {
	objc.Call[objc.Void](s_, objc.Sel("setAction:"), value)
}
