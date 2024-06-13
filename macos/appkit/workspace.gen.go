// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/uti"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Workspace] class.
var WorkspaceClass = _WorkspaceClass{objc.GetClass("NSWorkspace")}

type _WorkspaceClass struct {
	objc.Class
}

// An interface definition for the [Workspace] class.
type IWorkspace interface {
	objc.IObject
	SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler func(error foundation.Error))
	DuplicateURLsCompletionHandler(URLs []foundation.IURL, handler func(newURLs foundation.Dictionary, error foundation.Error))
	URLForApplicationToOpenContentType(contentType uti.IType) foundation.URL
	OpenApplicationAtURLConfigurationCompletionHandler(applicationURL foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler func(app RunningApplication, error foundation.Error))
	UnmountAndEjectDeviceAtPath(path string) bool
	DesktopImageURLForScreen(screen IScreen) foundation.URL
	OpenURLConfigurationCompletionHandler(url foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler func(app RunningApplication, error foundation.Error))
	URLsForApplicationsToOpenURL(url foundation.IURL) []foundation.URL
	NoteFileSystemChanged(path string)
	SetDesktopImageURLForScreenOptionsError(url foundation.IURL, screen IScreen, options map[WorkspaceDesktopImageOptionKey]objc.IObject, error unsafe.Pointer) bool
	URLsForApplicationsToOpenContentType(contentType uti.IType) []foundation.URL
	ActivateFileViewerSelectingURLs(fileURLs []foundation.IURL)
	ExtendPowerOffBy(requested int) int
	RequestAuthorizationOfTypeCompletionHandler(type_ WorkspaceAuthorizationType, completionHandler func(authorization WorkspaceAuthorization, error foundation.Error))
	RecycleURLsCompletionHandler(URLs []foundation.IURL, handler func(newURLs foundation.Dictionary, error foundation.Error))
	SelectFileInFileViewerRootedAtPath(fullPath string, rootFullPath string) bool
	SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler(applicationURL foundation.IURL, contentType uti.IType, completionHandler func(error foundation.Error))
	OpenURL(url foundation.IURL) bool
	IconForFiles(fullPaths []string) Image
	URLForApplicationWithBundleIdentifier(bundleIdentifier string) foundation.URL
	IsFilePackageAtPath(fullPath string) bool
	IconForFile(fullPath string) Image
	IconForContentType(contentType uti.IType) Image
	SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler func(error foundation.Error))
	DesktopImageOptionsForScreen(screen IScreen) map[WorkspaceDesktopImageOptionKey]objc.Object
	OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls []foundation.IURL, applicationURL foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler func(app RunningApplication, error foundation.Error))
	UnmountAndEjectDeviceAtURLError(url foundation.IURL, error unsafe.Pointer) bool
	HideOtherApplications()
	ShowSearchResultsForQueryString(queryString string) bool
	URLsForApplicationsWithBundleIdentifier(bundleIdentifier string) []foundation.URL
	SetIconForFileOptions(image IImage, fullPath string, options WorkspaceIconCreationOptions) bool
	SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler(applicationURL foundation.IURL, urlScheme string, completionHandler func(error foundation.Error))
	URLForApplicationToOpenURL(url foundation.IURL) foundation.URL
	GetFileSystemInfoForPathIsRemovableIsWritableIsUnmountableDescriptionType(fullPath string, removableFlag *bool, writableFlag *bool, unmountableFlag *bool, description unsafe.Pointer, fileSystemType unsafe.Pointer) bool
	AccessibilityDisplayShouldIncreaseContrast() bool
	AccessibilityDisplayShouldDifferentiateWithoutColor() bool
	AccessibilityDisplayShouldInvertColors() bool
	RunningApplications() []RunningApplication
	IsSwitchControlEnabled() bool
	MenuBarOwningApplication() RunningApplication
	IsVoiceOverEnabled() bool
	NotificationCenter() foundation.NotificationCenter
	AccessibilityDisplayShouldReduceMotion() bool
	FrontmostApplication() RunningApplication
	FileLabels() []string
	FileLabelColors() []Color
	AccessibilityDisplayShouldReduceTransparency() bool
}

// A workspace that can launch other apps and perform a variety of file-handling services. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace?language=objc
type Workspace struct {
	objc.Object
}

func WorkspaceFrom(ptr unsafe.Pointer) Workspace {
	return Workspace{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WorkspaceClass) Alloc() Workspace {
	rv := objc.Call[Workspace](wc, objc.Sel("alloc"))
	return rv
}

func (wc _WorkspaceClass) New() Workspace {
	rv := objc.Call[Workspace](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWorkspace() Workspace {
	return WorkspaceClass.New()
}

func (w_ Workspace) Init() Workspace {
	rv := objc.Call[Workspace](w_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/3753004-setdefaultapplicationaturl?language=objc
func (w_ Workspace) SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](w_, objc.Sel("setDefaultApplicationAtURL:toOpenFileAtURL:completionHandler:"), applicationURL, url, completionHandler)
}

// Duplicates the specified URLS asynchronously in the same manner as the Finder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1524490-duplicateurls?language=objc
func (w_ Workspace) DuplicateURLsCompletionHandler(URLs []foundation.IURL, handler func(newURLs foundation.Dictionary, error foundation.Error)) {
	objc.Call[objc.Void](w_, objc.Sel("duplicateURLs:completionHandler:"), URLs, handler)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/3752998-urlforapplicationtoopencontentty?language=objc
func (w_ Workspace) URLForApplicationToOpenContentType(contentType uti.IType) foundation.URL {
	rv := objc.Call[foundation.URL](w_, objc.Sel("URLForApplicationToOpenContentType:"), contentType)
	return rv
}

// Launches the app at the specified URL and asynchronously reports back on the app's status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/3172700-openapplicationaturl?language=objc
func (w_ Workspace) OpenApplicationAtURLConfigurationCompletionHandler(applicationURL foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler func(app RunningApplication, error foundation.Error)) {
	objc.Call[objc.Void](w_, objc.Sel("openApplicationAtURL:configuration:completionHandler:"), applicationURL, configuration, completionHandler)
}

// Unmounts and ejects the device at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1527741-unmountandejectdeviceatpath?language=objc
func (w_ Workspace) UnmountAndEjectDeviceAtPath(path string) bool {
	rv := objc.Call[bool](w_, objc.Sel("unmountAndEjectDeviceAtPath:"), path)
	return rv
}

// Returns the URL for the desktop image for the given screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1530635-desktopimageurlforscreen?language=objc
func (w_ Workspace) DesktopImageURLForScreen(screen IScreen) foundation.URL {
	rv := objc.Call[foundation.URL](w_, objc.Sel("desktopImageURLForScreen:"), screen)
	return rv
}

// Opens a URL asynchronously using the provided options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/3172701-openurl?language=objc
func (w_ Workspace) OpenURLConfigurationCompletionHandler(url foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler func(app RunningApplication, error foundation.Error)) {
	objc.Call[objc.Void](w_, objc.Sel("openURL:configuration:completionHandler:"), url, configuration, completionHandler)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/3753000-urlsforapplicationstoopenurl?language=objc
func (w_ Workspace) URLsForApplicationsToOpenURL(url foundation.IURL) []foundation.URL {
	rv := objc.Call[[]foundation.URL](w_, objc.Sel("URLsForApplicationsToOpenURL:"), url)
	return rv
}

// Informs the workspace object that the file system changed at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1525376-notefilesystemchanged?language=objc
func (w_ Workspace) NoteFileSystemChanged(path string) {
	objc.Call[objc.Void](w_, objc.Sel("noteFileSystemChanged:"), path)
}

// Sets the desktop image for the given screen to the image at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1527228-setdesktopimageurl?language=objc
func (w_ Workspace) SetDesktopImageURLForScreenOptionsError(url foundation.IURL, screen IScreen, options map[WorkspaceDesktopImageOptionKey]objc.IObject, error unsafe.Pointer) bool {
	rv := objc.Call[bool](w_, objc.Sel("setDesktopImageURL:forScreen:options:error:"), url, screen, options, error)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/3752999-urlsforapplicationstoopencontent?language=objc
func (w_ Workspace) URLsForApplicationsToOpenContentType(contentType uti.IType) []foundation.URL {
	rv := objc.Call[[]foundation.URL](w_, objc.Sel("URLsForApplicationsToOpenContentType:"), contentType)
	return rv
}

// Activates the Finder, and opens one or more windows selecting the specified files. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1524549-activatefileviewerselectingurls?language=objc
func (w_ Workspace) ActivateFileViewerSelectingURLs(fileURLs []foundation.IURL) {
	objc.Call[objc.Void](w_, objc.Sel("activateFileViewerSelectingURLs:"), fileURLs)
}

// Requests the system wait for the specified amount of time before turning off the power or logging out the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1533106-extendpoweroffby?language=objc
func (w_ Workspace) ExtendPowerOffBy(requested int) int {
	rv := objc.Call[int](w_, objc.Sel("extendPowerOffBy:"), requested)
	return rv
}

// Requests authorization to perform a privileged file operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/3025774-requestauthorizationoftype?language=objc
func (w_ Workspace) RequestAuthorizationOfTypeCompletionHandler(type_ WorkspaceAuthorizationType, completionHandler func(authorization WorkspaceAuthorization, error foundation.Error)) {
	objc.Call[objc.Void](w_, objc.Sel("requestAuthorizationOfType:completionHandler:"), type_, completionHandler)
}

// Moves the specified URLs to the trash in the same manner as the Finder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1530465-recycleurls?language=objc
func (w_ Workspace) RecycleURLsCompletionHandler(URLs []foundation.IURL, handler func(newURLs foundation.Dictionary, error foundation.Error)) {
	objc.Call[objc.Void](w_, objc.Sel("recycleURLs:completionHandler:"), URLs, handler)
}

// Selects the file at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1524399-selectfile?language=objc
func (w_ Workspace) SelectFileInFileViewerRootedAtPath(fullPath string, rootFullPath string) bool {
	rv := objc.Call[bool](w_, objc.Sel("selectFile:inFileViewerRootedAtPath:"), fullPath, rootFullPath)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/3753002-setdefaultapplicationaturl?language=objc
func (w_ Workspace) SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler(applicationURL foundation.IURL, contentType uti.IType, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](w_, objc.Sel("setDefaultApplicationAtURL:toOpenContentType:completionHandler:"), applicationURL, contentType, completionHandler)
}

// Opens the location at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1533463-openurl?language=objc
func (w_ Workspace) OpenURL(url foundation.IURL) bool {
	rv := objc.Call[bool](w_, objc.Sel("openURL:"), url)
	return rv
}

// Returns an image containing the icon for the specified files. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1525487-iconforfiles?language=objc
func (w_ Workspace) IconForFiles(fullPaths []string) Image {
	rv := objc.Call[Image](w_, objc.Sel("iconForFiles:"), fullPaths)
	return rv
}

// Returns the URL for the app with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1534053-urlforapplicationwithbundleident?language=objc
func (w_ Workspace) URLForApplicationWithBundleIdentifier(bundleIdentifier string) foundation.URL {
	rv := objc.Call[foundation.URL](w_, objc.Sel("URLForApplicationWithBundleIdentifier:"), bundleIdentifier)
	return rv
}

// Determines whether the specified path is a file package. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1529991-isfilepackageatpath?language=objc
func (w_ Workspace) IsFilePackageAtPath(fullPath string) bool {
	rv := objc.Call[bool](w_, objc.Sel("isFilePackageAtPath:"), fullPath)
	return rv
}

// Returns an image containing the icon for the specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1528158-iconforfile?language=objc
func (w_ Workspace) IconForFile(fullPath string) Image {
	rv := objc.Call[Image](w_, objc.Sel("iconForFile:"), fullPath)
	return rv
}

// Returns an image containing the icon for the specified content type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/3553230-iconforcontenttype?language=objc
func (w_ Workspace) IconForContentType(contentType uti.IType) Image {
	rv := objc.Call[Image](w_, objc.Sel("iconForContentType:"), contentType)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/3753003-setdefaultapplicationaturl?language=objc
func (w_ Workspace) SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](w_, objc.Sel("setDefaultApplicationAtURL:toOpenContentTypeOfFileAtURL:completionHandler:"), applicationURL, url, completionHandler)
}

// Returns the desktop image options for the given screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1530855-desktopimageoptionsforscreen?language=objc
func (w_ Workspace) DesktopImageOptionsForScreen(screen IScreen) map[WorkspaceDesktopImageOptionKey]objc.Object {
	rv := objc.Call[map[WorkspaceDesktopImageOptionKey]objc.Object](w_, objc.Sel("desktopImageOptionsForScreen:"), screen)
	return rv
}

// Opens one or more URLs asynchronously in the specified app using the provided options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/3172702-openurls?language=objc
func (w_ Workspace) OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls []foundation.IURL, applicationURL foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler func(app RunningApplication, error foundation.Error)) {
	objc.Call[objc.Void](w_, objc.Sel("openURLs:withApplicationAtURL:configuration:completionHandler:"), urls, applicationURL, configuration, completionHandler)
}

// Attempts to eject the volume mounted at the given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1530469-unmountandejectdeviceaturl?language=objc
func (w_ Workspace) UnmountAndEjectDeviceAtURLError(url foundation.IURL, error unsafe.Pointer) bool {
	rv := objc.Call[bool](w_, objc.Sel("unmountAndEjectDeviceAtURL:error:"), url, error)
	return rv
}

// Hides all applications other than the sender. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1530417-hideotherapplications?language=objc
func (w_ Workspace) HideOtherApplications() {
	objc.Call[objc.Void](w_, objc.Sel("hideOtherApplications"))
}

// Displays a Spotlight search results window in Finder for the specified query string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1532131-showsearchresultsforquerystring?language=objc
func (w_ Workspace) ShowSearchResultsForQueryString(queryString string) bool {
	rv := objc.Call[bool](w_, objc.Sel("showSearchResultsForQueryString:"), queryString)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/3753001-urlsforapplicationswithbundleide?language=objc
func (w_ Workspace) URLsForApplicationsWithBundleIdentifier(bundleIdentifier string) []foundation.URL {
	rv := objc.Call[[]foundation.URL](w_, objc.Sel("URLsForApplicationsWithBundleIdentifier:"), bundleIdentifier)
	return rv
}

// Sets the icon for the file or directory at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1529882-seticon?language=objc
func (w_ Workspace) SetIconForFileOptions(image IImage, fullPath string, options WorkspaceIconCreationOptions) bool {
	rv := objc.Call[bool](w_, objc.Sel("setIcon:forFile:options:"), image, fullPath, options)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/3753005-setdefaultapplicationaturl?language=objc
func (w_ Workspace) SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler(applicationURL foundation.IURL, urlScheme string, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](w_, objc.Sel("setDefaultApplicationAtURL:toOpenURLsWithScheme:completionHandler:"), applicationURL, urlScheme, completionHandler)
}

// Returns the URL to the default app that would be opened. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1533391-urlforapplicationtoopenurl?language=objc
func (w_ Workspace) URLForApplicationToOpenURL(url foundation.IURL) foundation.URL {
	rv := objc.Call[foundation.URL](w_, objc.Sel("URLForApplicationToOpenURL:"), url)
	return rv
}

// Returns information about the file system at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1526987-getfilesysteminfoforpath?language=objc
func (w_ Workspace) GetFileSystemInfoForPathIsRemovableIsWritableIsUnmountableDescriptionType(fullPath string, removableFlag *bool, writableFlag *bool, unmountableFlag *bool, description unsafe.Pointer, fileSystemType unsafe.Pointer) bool {
	rv := objc.Call[bool](w_, objc.Sel("getFileSystemInfoForPath:isRemovable:isWritable:isUnmountable:description:type:"), fullPath, removableFlag, writableFlag, unmountableFlag, description, fileSystemType)
	return rv
}

// A Boolean value that indicates whether the app presents a high-contrast user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1526290-accessibilitydisplayshouldincrea?language=objc
func (w_ Workspace) AccessibilityDisplayShouldIncreaseContrast() bool {
	rv := objc.Call[bool](w_, objc.Sel("accessibilityDisplayShouldIncreaseContrast"))
	return rv
}

// A Boolean value that indicates whether the app avoids conveying information through color alone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1524656-accessibilitydisplayshoulddiffer?language=objc
func (w_ Workspace) AccessibilityDisplayShouldDifferentiateWithoutColor() bool {
	rv := objc.Call[bool](w_, objc.Sel("accessibilityDisplayShouldDifferentiateWithoutColor"))
	return rv
}

// A Boolean value that indicates whether the accessibility option to invert colors is in an enabled state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1644068-accessibilitydisplayshouldinvert?language=objc
func (w_ Workspace) AccessibilityDisplayShouldInvertColors() bool {
	rv := objc.Call[bool](w_, objc.Sel("accessibilityDisplayShouldInvertColors"))
	return rv
}

// Returns an array of running apps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1534059-runningapplications?language=objc
func (w_ Workspace) RunningApplications() []RunningApplication {
	rv := objc.Call[[]RunningApplication](w_, objc.Sel("runningApplications"))
	return rv
}

// A Boolean value that indicates whether Switch Control is currently running. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/2880322-switchcontrolenabled?language=objc
func (w_ Workspace) IsSwitchControlEnabled() bool {
	rv := objc.Call[bool](w_, objc.Sel("isSwitchControlEnabled"))
	return rv
}

// Returns the app that owns the currently displayed menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1525848-menubarowningapplication?language=objc
func (w_ Workspace) MenuBarOwningApplication() RunningApplication {
	rv := objc.Call[RunningApplication](w_, objc.Sel("menuBarOwningApplication"))
	return rv
}

// A Boolean value that indicates whether VoiceOver is currently running. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/2880317-voiceoverenabled?language=objc
func (w_ Workspace) IsVoiceOverEnabled() bool {
	rv := objc.Call[bool](w_, objc.Sel("isVoiceOverEnabled"))
	return rv
}

// The notification center for workspace notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1525071-notificationcenter?language=objc
func (w_ Workspace) NotificationCenter() foundation.NotificationCenter {
	rv := objc.Call[foundation.NotificationCenter](w_, objc.Sel("notificationCenter"))
	return rv
}

// A Boolean value that indicates whether the accessibility option to reduce motion is in an enabled state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1644069-accessibilitydisplayshouldreduce?language=objc
func (w_ Workspace) AccessibilityDisplayShouldReduceMotion() bool {
	rv := objc.Call[bool](w_, objc.Sel("accessibilityDisplayShouldReduceMotion"))
	return rv
}

// Returns the frontmost app, which is the app that receives key events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1532097-frontmostapplication?language=objc
func (w_ Workspace) FrontmostApplication() RunningApplication {
	rv := objc.Call[RunningApplication](w_, objc.Sel("frontmostApplication"))
	return rv
}

// The shared workspace object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1530344-sharedworkspace?language=objc
func (wc _WorkspaceClass) SharedWorkspace() Workspace {
	rv := objc.Call[Workspace](wc, objc.Sel("sharedWorkspace"))
	return rv
}

// The shared workspace object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1530344-sharedworkspace?language=objc
func Workspace_SharedWorkspace() Workspace {
	return WorkspaceClass.SharedWorkspace()
}

// The array of file labels, returned as strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1533953-filelabels?language=objc
func (w_ Workspace) FileLabels() []string {
	rv := objc.Call[[]string](w_, objc.Sel("fileLabels"))
	return rv
}

// The array of colors for the file labels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1527553-filelabelcolors?language=objc
func (w_ Workspace) FileLabelColors() []Color {
	rv := objc.Call[[]Color](w_, objc.Sel("fileLabelColors"))
	return rv
}

// A Boolean value that indicates whether the app avoids using semitransparent backgrounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/1533006-accessibilitydisplayshouldreduce?language=objc
func (w_ Workspace) AccessibilityDisplayShouldReduceTransparency() bool {
	rv := objc.Call[bool](w_, objc.Sel("accessibilityDisplayShouldReduceTransparency"))
	return rv
}
