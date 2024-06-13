// Code generated by DarwinKit. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/contacts"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Geocoder] class.
var GeocoderClass = _GeocoderClass{objc.GetClass("CLGeocoder")}

type _GeocoderClass struct {
	objc.Class
}

// An interface definition for the [Geocoder] class.
type IGeocoder interface {
	objc.IObject
	GeocodeAddressStringInRegionCompletionHandler(addressString string, region IRegion, completionHandler GeocodeCompletionHandler)
	GeocodePostalAddressCompletionHandler(postalAddress contacts.IPostalAddress, completionHandler GeocodeCompletionHandler)
	GeocodeAddressStringInRegionPreferredLocaleCompletionHandler(addressString string, region IRegion, locale foundation.ILocale, completionHandler GeocodeCompletionHandler)
	ReverseGeocodeLocationPreferredLocaleCompletionHandler(location ILocation, locale foundation.ILocale, completionHandler GeocodeCompletionHandler)
	GeocodePostalAddressPreferredLocaleCompletionHandler(postalAddress contacts.IPostalAddress, locale foundation.ILocale, completionHandler GeocodeCompletionHandler)
	CancelGeocode()
	GeocodeAddressStringCompletionHandler(addressString string, completionHandler GeocodeCompletionHandler)
	ReverseGeocodeLocationCompletionHandler(location ILocation, completionHandler GeocodeCompletionHandler)
	IsGeocoding() bool
}

// An interface for converting between geographic coordinates and place names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder?language=objc
type Geocoder struct {
	objc.Object
}

func GeocoderFrom(ptr unsafe.Pointer) Geocoder {
	return Geocoder{
		Object: objc.ObjectFrom(ptr),
	}
}

func (gc _GeocoderClass) Alloc() Geocoder {
	rv := objc.Call[Geocoder](gc, objc.Sel("alloc"))
	return rv
}

func (gc _GeocoderClass) New() Geocoder {
	rv := objc.Call[Geocoder](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGeocoder() Geocoder {
	return GeocoderClass.New()
}

func (g_ Geocoder) Init() Geocoder {
	rv := objc.Call[Geocoder](g_, objc.Sel("init"))
	return rv
}

// Submits a forward-geocoding request using the specified string and region information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/1423591-geocodeaddressstring?language=objc
func (g_ Geocoder) GeocodeAddressStringInRegionCompletionHandler(addressString string, region IRegion, completionHandler GeocodeCompletionHandler) {
	objc.Call[objc.Void](g_, objc.Sel("geocodeAddressString:inRegion:completionHandler:"), addressString, region, completionHandler)
}

// Submits a forward-geocoding requesting using the specified Contacts framework information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/2890752-geocodepostaladdress?language=objc
func (g_ Geocoder) GeocodePostalAddressCompletionHandler(postalAddress contacts.IPostalAddress, completionHandler GeocodeCompletionHandler) {
	objc.Call[objc.Void](g_, objc.Sel("geocodePostalAddress:completionHandler:"), postalAddress, completionHandler)
}

// Submits a forward-geocoding requesting using the specified address string and locale information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/2890753-geocodeaddressstring?language=objc
func (g_ Geocoder) GeocodeAddressStringInRegionPreferredLocaleCompletionHandler(addressString string, region IRegion, locale foundation.ILocale, completionHandler GeocodeCompletionHandler) {
	objc.Call[objc.Void](g_, objc.Sel("geocodeAddressString:inRegion:preferredLocale:completionHandler:"), addressString, region, locale, completionHandler)
}

// Submits a reverse-geocoding request for the specified location and locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/2908779-reversegeocodelocation?language=objc
func (g_ Geocoder) ReverseGeocodeLocationPreferredLocaleCompletionHandler(location ILocation, locale foundation.ILocale, completionHandler GeocodeCompletionHandler) {
	objc.Call[objc.Void](g_, objc.Sel("reverseGeocodeLocation:preferredLocale:completionHandler:"), location, locale, completionHandler)
}

// Submits a forward-geocoding requesting using the specified locale and Contacts framework information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/2890750-geocodepostaladdress?language=objc
func (g_ Geocoder) GeocodePostalAddressPreferredLocaleCompletionHandler(postalAddress contacts.IPostalAddress, locale foundation.ILocale, completionHandler GeocodeCompletionHandler) {
	objc.Call[objc.Void](g_, objc.Sel("geocodePostalAddress:preferredLocale:completionHandler:"), postalAddress, locale, completionHandler)
}

// Cancels a pending geocoding request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/1423562-cancelgeocode?language=objc
func (g_ Geocoder) CancelGeocode() {
	objc.Call[objc.Void](g_, objc.Sel("cancelGeocode"))
}

// Submits a forward-geocoding request using the specified string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/1423509-geocodeaddressstring?language=objc
func (g_ Geocoder) GeocodeAddressStringCompletionHandler(addressString string, completionHandler GeocodeCompletionHandler) {
	objc.Call[objc.Void](g_, objc.Sel("geocodeAddressString:completionHandler:"), addressString, completionHandler)
}

// Submits a reverse-geocoding request for the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/1423621-reversegeocodelocation?language=objc
func (g_ Geocoder) ReverseGeocodeLocationCompletionHandler(location ILocation, completionHandler GeocodeCompletionHandler) {
	objc.Call[objc.Void](g_, objc.Sel("reverseGeocodeLocation:completionHandler:"), location, completionHandler)
}

// A Boolean value indicating whether the receiver is in the middle of geocoding its value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/1423765-geocoding?language=objc
func (g_ Geocoder) IsGeocoding() bool {
	rv := objc.Call[bool](g_, objc.Sel("isGeocoding"))
	return rv
}
