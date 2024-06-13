// Code generated by DarwinKit. DO NOT EDIT.

package coremediaio

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioclassid?language=objc
type ClassID uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmiocontrolid?language=objc
type ControlID ObjectID

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmiodeviceid?language=objc
type DeviceID ObjectID

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmiodevicepropertyid?language=objc
type DevicePropertyID ObjectPropertySelector

// A structure that defines the properties that providers, devices, and streams support. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionproperty?language=objc
type ExtensionProperty string

const (
	ExtensionPropertyDeviceCanBeDefaultInputDevice       ExtensionProperty = "CMIOExtensionPropertyDeviceCanBeDefaultInputDevice"
	ExtensionPropertyDeviceCanBeDefaultOutputDevice      ExtensionProperty = "CMIOExtensionPropertyDeviceCanBeDefaultOutputDevice"
	ExtensionPropertyDeviceIsSuspended                   ExtensionProperty = "CMIOExtensionPropertyDeviceIsSuspended"
	ExtensionPropertyDeviceLinkedCoreAudioDeviceUID      ExtensionProperty = "CMIOExtensionPropertyDeviceLinkedCoreAudioDeviceUID"
	ExtensionPropertyDeviceModel                         ExtensionProperty = "CMIOExtensionPropertyDeviceModel"
	ExtensionPropertyDeviceTransportType                 ExtensionProperty = "CMIOExtensionPropertyDeviceTransportType"
	ExtensionPropertyProviderManufacturer                ExtensionProperty = "CMIOExtensionPropertyProviderManufacturer"
	ExtensionPropertyProviderName                        ExtensionProperty = "CMIOExtensionPropertyProviderName"
	ExtensionPropertyStreamActiveFormatIndex             ExtensionProperty = "CMIOExtensionPropertyStreamActiveFormatIndex"
	ExtensionPropertyStreamFrameDuration                 ExtensionProperty = "CMIOExtensionPropertyStreamFrameDuration"
	ExtensionPropertyStreamMaxFrameDuration              ExtensionProperty = "CMIOExtensionPropertyStreamMaxFrameDuration"
	ExtensionPropertyStreamSinkBufferQueueSize           ExtensionProperty = "CMIOExtensionPropertyStreamSinkBufferQueueSize"
	ExtensionPropertyStreamSinkBufferUnderrunCount       ExtensionProperty = "CMIOExtensionPropertyStreamSinkBufferUnderrunCount"
	ExtensionPropertyStreamSinkBuffersRequiredForStartup ExtensionProperty = "CMIOExtensionPropertyStreamSinkBuffersRequiredForStartup"
	ExtensionPropertyStreamSinkEndOfData                 ExtensionProperty = "CMIOExtensionPropertyStreamSinkEndOfData"
)

// Constants that indicate the clock type of a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamclocktype?language=objc
type ExtensionStreamClockType int

const (
	ExtensionStreamClockTypeCustom                   ExtensionStreamClockType = 2
	ExtensionStreamClockTypeHostTime                 ExtensionStreamClockType = 0
	ExtensionStreamClockTypeLinkedCoreAudioDeviceUID ExtensionStreamClockType = 1
)

// Constants that define the data-flow direction of the stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamdirection?language=objc
type ExtensionStreamDirection int

const (
	ExtensionStreamDirectionSink   ExtensionStreamDirection = 1
	ExtensionStreamDirectionSource ExtensionStreamDirection = 0
)

// Constants that specify the types of discontinuities that can occur in a media stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamdiscontinuityflags?language=objc
type ExtensionStreamDiscontinuityFlags uint32

const (
	ExtensionStreamDiscontinuityFlagNone          ExtensionStreamDiscontinuityFlags = 0
	ExtensionStreamDiscontinuityFlagSampleDropped ExtensionStreamDiscontinuityFlags = 64
	ExtensionStreamDiscontinuityFlagTime          ExtensionStreamDiscontinuityFlags = 2
	ExtensionStreamDiscontinuityFlagUnknown       ExtensionStreamDiscontinuityFlags = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmiohardwarepropertyid?language=objc
type HardwarePropertyID ObjectPropertySelector

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioobjectid?language=objc
type ObjectID uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioobjectpropertyelement?language=objc
type ObjectPropertyElement uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioobjectpropertyscope?language=objc
type ObjectPropertyScope uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioobjectpropertyselector?language=objc
type ObjectPropertySelector uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmiostreamid?language=objc
type StreamID ObjectID
