// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableMovie] class.
var MutableMovieClass = _MutableMovieClass{objc.GetClass("AVMutableMovie")}

type _MutableMovieClass struct {
	objc.Class
}

// An interface definition for the [MutableMovie] class.
type IMutableMovie interface {
	IMovie
	TracksWithMediaType(mediaType MediaType) []MutableMovieTrack
	RemoveTimeRange(timeRange coremedia.TimeRange)
	AddMutableTracksCopyingSettingsFromTracksOptions(existingTracks []IAssetTrack, options map[string]objc.IObject) []MutableMovieTrack
	TracksWithMediaCharacteristic(mediaCharacteristic MediaCharacteristic) []MutableMovieTrack
	MutableTrackCompatibleWithTrack(track IAssetTrack) MutableMovieTrack
	InsertTimeRangeOfAssetAtTimeCopySampleDataError(timeRange coremedia.TimeRange, asset IAsset, startTime coremedia.Time, copySampleData bool, outError foundation.IError) bool
	RemoveTrack(track IMovieTrack)
	AddMutableTrackWithMediaTypeCopySettingsFromTrackOptions(mediaType MediaType, track IAssetTrack, options map[string]objc.IObject) MutableMovieTrack
	ScaleTimeRangeToDuration(timeRange coremedia.TimeRange, duration coremedia.Time)
	TrackWithTrackID(trackID objc.IObject) MutableMovieTrack
	InsertEmptyTimeRange(timeRange coremedia.TimeRange)
	SetDefaultMediaDataStorage(value IMediaDataStorage)
	IsModified() bool
	SetModified(value bool)
	SetPreferredVolume(value float64)
	InterleavingPeriod() coremedia.Time
	SetInterleavingPeriod(value coremedia.Time)
	SetMetadata(value []IMetadataItem)
	SetPreferredRate(value float64)
	Timescale() coremedia.TimeScale
	SetTimescale(value coremedia.TimeScale)
	SetPreferredTransform(value coregraphics.AffineTransform)
}

// A mutable object that represents an audiovisual container that conforms to the QuickTime movie file format or a related format like MPEG-4. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie?language=objc
type MutableMovie struct {
	Movie
}

func MutableMovieFrom(ptr unsafe.Pointer) MutableMovie {
	return MutableMovie{
		Movie: MovieFrom(ptr),
	}
}

func (m_ MutableMovie) InitWithDataOptionsError(data []byte, options map[string]objc.IObject, outError foundation.IError) MutableMovie {
	rv := objc.Call[MutableMovie](m_, objc.Sel("initWithData:options:error:"), data, options, objc.Ptr(outError))
	return rv
}

// Creates a mutable movie object from a movie stored in a data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1388442-initwithdata?language=objc
func NewMutableMovieWithDataOptionsError(data []byte, options map[string]objc.IObject, outError foundation.IError) MutableMovie {
	instance := MutableMovieClass.Alloc().InitWithDataOptionsError(data, options, outError)
	instance.Autorelease()
	return instance
}

func (mc _MutableMovieClass) MovieWithDataOptionsError(data []byte, options map[string]objc.IObject, outError foundation.IError) MutableMovie {
	rv := objc.Call[MutableMovie](mc, objc.Sel("movieWithData:options:error:"), data, options, objc.Ptr(outError))
	return rv
}

// Returns a new mutable movie object from a movie stored in a data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1458234-moviewithdata?language=objc
func MutableMovie_MovieWithDataOptionsError(data []byte, options map[string]objc.IObject, outError foundation.IError) MutableMovie {
	return MutableMovieClass.MovieWithDataOptionsError(data, options, outError)
}

func (m_ MutableMovie) InitWithSettingsFromMovieOptionsError(movie IMovie, options map[string]objc.IObject, outError foundation.IError) MutableMovie {
	rv := objc.Call[MutableMovie](m_, objc.Sel("initWithSettingsFromMovie:options:error:"), objc.Ptr(movie), options, objc.Ptr(outError))
	return rv
}

// Creates a mutable movie object without tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1386408-initwithsettingsfrommovie?language=objc
func NewMutableMovieWithSettingsFromMovieOptionsError(movie IMovie, options map[string]objc.IObject, outError foundation.IError) MutableMovie {
	instance := MutableMovieClass.Alloc().InitWithSettingsFromMovieOptionsError(movie, options, outError)
	instance.Autorelease()
	return instance
}

func (m_ MutableMovie) InitWithURLOptionsError(URL foundation.IURL, options map[string]objc.IObject, outError foundation.IError) MutableMovie {
	rv := objc.Call[MutableMovie](m_, objc.Sel("initWithURL:options:error:"), objc.Ptr(URL), options, objc.Ptr(outError))
	return rv
}

// Creates a mutable movie object from a movie header stored in a QuickTime movie file of ISO base media file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1386052-initwithurl?language=objc
func NewMutableMovieWithURLOptionsError(URL foundation.IURL, options map[string]objc.IObject, outError foundation.IError) MutableMovie {
	instance := MutableMovieClass.Alloc().InitWithURLOptionsError(URL, options, outError)
	instance.Autorelease()
	return instance
}

func (mc _MutableMovieClass) MovieWithURLOptionsError(URL foundation.IURL, options map[string]objc.IObject, outError foundation.IError) MutableMovie {
	rv := objc.Call[MutableMovie](mc, objc.Sel("movieWithURL:options:error:"), objc.Ptr(URL), options, objc.Ptr(outError))
	return rv
}

// Returns a new mutable movie object from a movie header stored in a QuickTime movie file of ISO base media file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1458226-moviewithurl?language=objc
func MutableMovie_MovieWithURLOptionsError(URL foundation.IURL, options map[string]objc.IObject, outError foundation.IError) MutableMovie {
	return MutableMovieClass.MovieWithURLOptionsError(URL, options, outError)
}

func (mc _MutableMovieClass) MovieWithSettingsFromMovieOptionsError(movie IMovie, options map[string]objc.IObject, outError foundation.IError) MutableMovie {
	rv := objc.Call[MutableMovie](mc, objc.Sel("movieWithSettingsFromMovie:options:error:"), objc.Ptr(movie), options, objc.Ptr(outError))
	return rv
}

// Returns a new mutable movie object without tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1458238-moviewithsettingsfrommovie?language=objc
func MutableMovie_MovieWithSettingsFromMovieOptionsError(movie IMovie, options map[string]objc.IObject, outError foundation.IError) MutableMovie {
	return MutableMovieClass.MovieWithSettingsFromMovieOptionsError(movie, options, outError)
}

func (mc _MutableMovieClass) Alloc() MutableMovie {
	rv := objc.Call[MutableMovie](mc, objc.Sel("alloc"))
	return rv
}

func MutableMovie_Alloc() MutableMovie {
	return MutableMovieClass.Alloc()
}

func (mc _MutableMovieClass) New() MutableMovie {
	rv := objc.Call[MutableMovie](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableMovie() MutableMovie {
	return MutableMovieClass.New()
}

func (m_ MutableMovie) Init() MutableMovie {
	rv := objc.Call[MutableMovie](m_, objc.Sel("init"))
	return rv
}

func (m_ MutableMovie) InitWithDataOptions(data []byte, options map[string]objc.IObject) MutableMovie {
	rv := objc.Call[MutableMovie](m_, objc.Sel("initWithData:options:"), data, options)
	return rv
}

// Creates a movie object from a movie file’s data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1388090-initwithdata?language=objc
func NewMutableMovieWithDataOptions(data []byte, options map[string]objc.IObject) MutableMovie {
	instance := MutableMovieClass.Alloc().InitWithDataOptions(data, options)
	instance.Autorelease()
	return instance
}

func (mc _MutableMovieClass) MovieWithDataOptions(data []byte, options map[string]objc.IObject) MutableMovie {
	rv := objc.Call[MutableMovie](mc, objc.Sel("movieWithData:options:"), data, options)
	return rv
}

// Returns a new movie object from a movie file’s data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1458261-moviewithdata?language=objc
func MutableMovie_MovieWithDataOptions(data []byte, options map[string]objc.IObject) MutableMovie {
	return MutableMovieClass.MovieWithDataOptions(data, options)
}

func (m_ MutableMovie) InitWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) MutableMovie {
	rv := objc.Call[MutableMovie](m_, objc.Sel("initWithURL:options:"), objc.Ptr(URL), options)
	return rv
}

// Creates a movie object from a movie header stored in a QuickTime movie file of ISO base media file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1387923-initwithurl?language=objc
func NewMutableMovieWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) MutableMovie {
	instance := MutableMovieClass.Alloc().InitWithURLOptions(URL, options)
	instance.Autorelease()
	return instance
}

func (mc _MutableMovieClass) MovieWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) MutableMovie {
	rv := objc.Call[MutableMovie](mc, objc.Sel("movieWithURL:options:"), objc.Ptr(URL), options)
	return rv
}

// Returns a new movie object from a movie header stored in a QuickTime movie file of ISO base media file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1458223-moviewithurl?language=objc
func MutableMovie_MovieWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) MutableMovie {
	return MutableMovieClass.MovieWithURLOptions(URL, options)
}

func (mc _MutableMovieClass) AssetWithURL(URL foundation.IURL) MutableMovie {
	rv := objc.Call[MutableMovie](mc, objc.Sel("assetWithURL:"), objc.Ptr(URL))
	return rv
}

// Creates an asset that models the media at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1389943-assetwithurl?language=objc
func MutableMovie_AssetWithURL(URL foundation.IURL) MutableMovie {
	return MutableMovieClass.AssetWithURL(URL)
}

// Retrieves tracks in the movie that present media of the specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1390443-trackswithmediatype?language=objc
func (m_ MutableMovie) TracksWithMediaType(mediaType MediaType) []MutableMovieTrack {
	rv := objc.Call[[]MutableMovieTrack](m_, objc.Sel("tracksWithMediaType:"), mediaType)
	return rv
}

// Removes the specified time range from a movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1385605-removetimerange?language=objc
func (m_ MutableMovie) RemoveTimeRange(timeRange coremedia.TimeRange) {
	objc.Call[objc.Void](m_, objc.Sel("removeTimeRange:"), timeRange)
}

// Adds one or more empty tracks to the target movie and copies the track settings from the source tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1389215-addmutabletrackscopyingsettingsf?language=objc
func (m_ MutableMovie) AddMutableTracksCopyingSettingsFromTracksOptions(existingTracks []IAssetTrack, options map[string]objc.IObject) []MutableMovieTrack {
	rv := objc.Call[[]MutableMovieTrack](m_, objc.Sel("addMutableTracksCopyingSettingsFromTracks:options:"), existingTracks, options)
	return rv
}

// Retrieve tracks in the movie that present media of the specified characteristic. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1388547-trackswithmediacharacteristic?language=objc
func (m_ MutableMovie) TracksWithMediaCharacteristic(mediaCharacteristic MediaCharacteristic) []MutableMovieTrack {
	rv := objc.Call[[]MutableMovieTrack](m_, objc.Sel("tracksWithMediaCharacteristic:"), mediaCharacteristic)
	return rv
}

// Provides a reference to a track from a mutable movie into which you can insert any time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1388669-mutabletrackcompatiblewithtrack?language=objc
func (m_ MutableMovie) MutableTrackCompatibleWithTrack(track IAssetTrack) MutableMovieTrack {
	rv := objc.Call[MutableMovieTrack](m_, objc.Sel("mutableTrackCompatibleWithTrack:"), objc.Ptr(track))
	return rv
}

// Inserts all of the tracks in a specified time range of an asset into a movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1389598-inserttimerange?language=objc
func (m_ MutableMovie) InsertTimeRangeOfAssetAtTimeCopySampleDataError(timeRange coremedia.TimeRange, asset IAsset, startTime coremedia.Time, copySampleData bool, outError foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("insertTimeRange:ofAsset:atTime:copySampleData:error:"), timeRange, objc.Ptr(asset), startTime, copySampleData, objc.Ptr(outError))
	return rv
}

// Removes the specified track from the target movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1386735-removetrack?language=objc
func (m_ MutableMovie) RemoveTrack(track IMovieTrack) {
	objc.Call[objc.Void](m_, objc.Sel("removeTrack:"), objc.Ptr(track))
}

// Adds an empty track to the target movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1390063-addmutabletrackwithmediatype?language=objc
func (m_ MutableMovie) AddMutableTrackWithMediaTypeCopySettingsFromTrackOptions(mediaType MediaType, track IAssetTrack, options map[string]objc.IObject) MutableMovieTrack {
	rv := objc.Call[MutableMovieTrack](m_, objc.Sel("addMutableTrackWithMediaType:copySettingsFromTrack:options:"), mediaType, objc.Ptr(track), options)
	return rv
}

// Changes the duration of a time range in a movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1385653-scaletimerange?language=objc
func (m_ MutableMovie) ScaleTimeRangeToDuration(timeRange coremedia.TimeRange, duration coremedia.Time) {
	objc.Call[objc.Void](m_, objc.Sel("scaleTimeRange:toDuration:"), timeRange, duration)
}

// Retrieves a track in the movie that contains the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1389467-trackwithtrackid?language=objc
func (m_ MutableMovie) TrackWithTrackID(trackID objc.IObject) MutableMovieTrack {
	rv := objc.Call[MutableMovieTrack](m_, objc.Sel("trackWithTrackID:"), objc.Ptr(trackID))
	return rv
}

// Adds an empty time range to a movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1387515-insertemptytimerange?language=objc
func (m_ MutableMovie) InsertEmptyTimeRange(timeRange coremedia.TimeRange) {
	objc.Call[objc.Void](m_, objc.Sel("insertEmptyTimeRange:"), timeRange)
}

// The default storage container for media data that you add to a movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1389320-defaultmediadatastorage?language=objc
func (m_ MutableMovie) SetDefaultMediaDataStorage(value IMediaDataStorage) {
	objc.Call[objc.Void](m_, objc.Sel("setDefaultMediaDataStorage:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the movie is in a modified state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1389960-modified?language=objc
func (m_ MutableMovie) IsModified() bool {
	rv := objc.Call[bool](m_, objc.Sel("isModified"))
	return rv
}

// A Boolean value that indicates whether the movie is in a modified state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1389960-modified?language=objc
func (m_ MutableMovie) SetModified(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setModified:"), value)
}

// The preferred volume for the audible medata data of the movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1388614-preferredvolume?language=objc
func (m_ MutableMovie) SetPreferredVolume(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setPreferredVolume:"), value)
}

// A time period indicating the duration for interleaving runs of samples for each track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1386969-interleavingperiod?language=objc
func (m_ MutableMovie) InterleavingPeriod() coremedia.Time {
	rv := objc.Call[coremedia.Time](m_, objc.Sel("interleavingPeriod"))
	return rv
}

// A time period indicating the duration for interleaving runs of samples for each track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1386969-interleavingperiod?language=objc
func (m_ MutableMovie) SetInterleavingPeriod(value coremedia.Time) {
	objc.Call[objc.Void](m_, objc.Sel("setInterleavingPeriod:"), value)
}

// An array of metadata the movie stores. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1388742-metadata?language=objc
func (m_ MutableMovie) SetMetadata(value []IMetadataItem) {
	objc.Call[objc.Void](m_, objc.Sel("setMetadata:"), value)
}

// The natural rate for playing the movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1387335-preferredrate?language=objc
func (m_ MutableMovie) SetPreferredRate(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setPreferredRate:"), value)
}

// The time scale of the movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1390622-timescale?language=objc
func (m_ MutableMovie) Timescale() coremedia.TimeScale {
	rv := objc.Call[coremedia.TimeScale](m_, objc.Sel("timescale"))
	return rv
}

// The time scale of the movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1390622-timescale?language=objc
func (m_ MutableMovie) SetTimescale(value coremedia.TimeScale) {
	objc.Call[objc.Void](m_, objc.Sel("setTimescale:"), value)
}

// The transform performed on the visual media data of the movie for display purposes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/1388771-preferredtransform?language=objc
func (m_ MutableMovie) SetPreferredTransform(value coregraphics.AffineTransform) {
	objc.Call[objc.Void](m_, objc.Sel("setPreferredTransform:"), value)
}