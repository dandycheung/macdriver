// AUTO-GENERATED CODE, DO NOT MODIFY

package coremediaio

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExtensionDeviceProperties] class.
var ExtensionDevicePropertiesClass = _ExtensionDevicePropertiesClass{objc.GetClass("CMIOExtensionDeviceProperties")}

type _ExtensionDevicePropertiesClass struct {
	objc.Class
}

// An interface definition for the [ExtensionDeviceProperties] class.
type IExtensionDeviceProperties interface {
	objc.IObject
	SetPropertyStateForProperty(propertyState IExtensionPropertyState, property ExtensionProperty)
	Model() string
	SetModel(value string)
	TransportType() foundation.Number
	SetTransportType(value foundation.INumber)
	LinkedCoreAudioDeviceUID() string
	SetLinkedCoreAudioDeviceUID(value string)
	Suspended() foundation.Number
	SetSuspended(value foundation.INumber)
	PropertiesDictionary() map[ExtensionProperty]ExtensionPropertyState
	SetPropertiesDictionary(value map[ExtensionProperty]IExtensionPropertyState)
}

// An object that defines the properties of a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondeviceproperties?language=objc
type ExtensionDeviceProperties struct {
	objc.Object
}

func ExtensionDevicePropertiesFrom(ptr unsafe.Pointer) ExtensionDeviceProperties {
	return ExtensionDeviceProperties{
		Object: objc.ObjectFrom(ptr),
	}
}

func (e_ ExtensionDeviceProperties) InitWithDictionary(propertiesDictionary map[ExtensionProperty]IExtensionPropertyState) ExtensionDeviceProperties {
	rv := objc.Call[ExtensionDeviceProperties](e_, objc.Sel("initWithDictionary:"), propertiesDictionary)
	return rv
}

// Creates a properties object with a dictionary of property states. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondeviceproperties/3915840-initwithdictionary?language=objc
func NewExtensionDevicePropertiesWithDictionary(propertiesDictionary map[ExtensionProperty]IExtensionPropertyState) ExtensionDeviceProperties {
	instance := ExtensionDevicePropertiesClass.Alloc().InitWithDictionary(propertiesDictionary)
	instance.Autorelease()
	return instance
}

func (ec _ExtensionDevicePropertiesClass) DevicePropertiesWithDictionary(propertiesDictionary map[ExtensionProperty]IExtensionPropertyState) ExtensionDeviceProperties {
	rv := objc.Call[ExtensionDeviceProperties](ec, objc.Sel("devicePropertiesWithDictionary:"), propertiesDictionary)
	return rv
}

// Returns a new properties object with a dictionary of property states. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondeviceproperties/3915839-devicepropertieswithdictionary?language=objc
func ExtensionDeviceProperties_DevicePropertiesWithDictionary(propertiesDictionary map[ExtensionProperty]IExtensionPropertyState) ExtensionDeviceProperties {
	return ExtensionDevicePropertiesClass.DevicePropertiesWithDictionary(propertiesDictionary)
}

func (ec _ExtensionDevicePropertiesClass) Alloc() ExtensionDeviceProperties {
	rv := objc.Call[ExtensionDeviceProperties](ec, objc.Sel("alloc"))
	return rv
}

func ExtensionDeviceProperties_Alloc() ExtensionDeviceProperties {
	return ExtensionDevicePropertiesClass.Alloc()
}

func (ec _ExtensionDevicePropertiesClass) New() ExtensionDeviceProperties {
	rv := objc.Call[ExtensionDeviceProperties](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionDeviceProperties() ExtensionDeviceProperties {
	return ExtensionDevicePropertiesClass.New()
}

func (e_ ExtensionDeviceProperties) Init() ExtensionDeviceProperties {
	rv := objc.Call[ExtensionDeviceProperties](e_, objc.Sel("init"))
	return rv
}

// Sets the value of a device property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondeviceproperties/3915844-setpropertystate?language=objc
func (e_ ExtensionDeviceProperties) SetPropertyStateForProperty(propertyState IExtensionPropertyState, property ExtensionProperty) {
	objc.Call[objc.Void](e_, objc.Sel("setPropertyState:forProperty:"), objc.Ptr(propertyState), property)
}

// A device model string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondeviceproperties/3915842-model?language=objc
func (e_ ExtensionDeviceProperties) Model() string {
	rv := objc.Call[string](e_, objc.Sel("model"))
	return rv
}

// A device model string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondeviceproperties/3915842-model?language=objc
func (e_ ExtensionDeviceProperties) SetModel(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setModel:"), value)
}

// The transport type of the device, such as USB or HDMI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondeviceproperties/3915846-transporttype?language=objc
func (e_ ExtensionDeviceProperties) TransportType() foundation.Number {
	rv := objc.Call[foundation.Number](e_, objc.Sel("transportType"))
	return rv
}

// The transport type of the device, such as USB or HDMI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondeviceproperties/3915846-transporttype?language=objc
func (e_ ExtensionDeviceProperties) SetTransportType(value foundation.INumber) {
	objc.Call[objc.Void](e_, objc.Sel("setTransportType:"), objc.Ptr(value))
}

// A universal identifier of the audio device linked to this device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondeviceproperties/3915841-linkedcoreaudiodeviceuid?language=objc
func (e_ ExtensionDeviceProperties) LinkedCoreAudioDeviceUID() string {
	rv := objc.Call[string](e_, objc.Sel("linkedCoreAudioDeviceUID"))
	return rv
}

// A universal identifier of the audio device linked to this device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondeviceproperties/3915841-linkedcoreaudiodeviceuid?language=objc
func (e_ ExtensionDeviceProperties) SetLinkedCoreAudioDeviceUID(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setLinkedCoreAudioDeviceUID:"), value)
}

// A Boolean value that indicates whether the device is in a suspended state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondeviceproperties/3915845-suspended?language=objc
func (e_ ExtensionDeviceProperties) Suspended() foundation.Number {
	rv := objc.Call[foundation.Number](e_, objc.Sel("suspended"))
	return rv
}

// A Boolean value that indicates whether the device is in a suspended state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondeviceproperties/3915845-suspended?language=objc
func (e_ ExtensionDeviceProperties) SetSuspended(value foundation.INumber) {
	objc.Call[objc.Void](e_, objc.Sel("setSuspended:"), objc.Ptr(value))
}

// A dictionary of properties for a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondeviceproperties/3915843-propertiesdictionary?language=objc
func (e_ ExtensionDeviceProperties) PropertiesDictionary() map[ExtensionProperty]ExtensionPropertyState {
	rv := objc.Call[map[ExtensionProperty]ExtensionPropertyState](e_, objc.Sel("propertiesDictionary"))
	return rv
}

// A dictionary of properties for a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondeviceproperties/3915843-propertiesdictionary?language=objc
func (e_ ExtensionDeviceProperties) SetPropertiesDictionary(value map[ExtensionProperty]IExtensionPropertyState) {
	objc.Call[objc.Void](e_, objc.Sel("setPropertiesDictionary:"), value)
}