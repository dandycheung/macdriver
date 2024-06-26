// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [NotificationQueue] class.
var NotificationQueueClass = _NotificationQueueClass{objc.GetClass("NSNotificationQueue")}

type _NotificationQueueClass struct {
	objc.Class
}

// An interface definition for the [NotificationQueue] class.
type INotificationQueue interface {
	objc.IObject
	EnqueueNotificationPostingStyleCoalesceMaskForModes(notification INotification, postingStyle PostingStyle, coalesceMask NotificationCoalescing, modes []RunLoopMode)
	EnqueueNotificationPostingStyle(notification INotification, postingStyle PostingStyle)
	DequeueNotificationsMatchingCoalesceMask(notification INotification, coalesceMask uint)
}

// A notification center buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationqueue?language=objc
type NotificationQueue struct {
	objc.Object
}

func NotificationQueueFrom(ptr unsafe.Pointer) NotificationQueue {
	return NotificationQueue{
		Object: objc.ObjectFrom(ptr),
	}
}

func (n_ NotificationQueue) InitWithNotificationCenter(notificationCenter INotificationCenter) NotificationQueue {
	rv := objc.Call[NotificationQueue](n_, objc.Sel("initWithNotificationCenter:"), notificationCenter)
	return rv
}

// Initializes and returns a notification queue for the specified notification center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationqueue/1410062-initwithnotificationcenter?language=objc
func NewNotificationQueueWithNotificationCenter(notificationCenter INotificationCenter) NotificationQueue {
	instance := NotificationQueueClass.Alloc().InitWithNotificationCenter(notificationCenter)
	instance.Autorelease()
	return instance
}

func (nc _NotificationQueueClass) Alloc() NotificationQueue {
	rv := objc.Call[NotificationQueue](nc, objc.Sel("alloc"))
	return rv
}

func (nc _NotificationQueueClass) New() NotificationQueue {
	rv := objc.Call[NotificationQueue](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNotificationQueue() NotificationQueue {
	return NotificationQueueClass.New()
}

func (n_ NotificationQueue) Init() NotificationQueue {
	rv := objc.Call[NotificationQueue](n_, objc.Sel("init"))
	return rv
}

// Adds a notification to the notification queue with a specified posting style, criteria for coalescing, and run loop mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationqueue/1413873-enqueuenotification?language=objc
func (n_ NotificationQueue) EnqueueNotificationPostingStyleCoalesceMaskForModes(notification INotification, postingStyle PostingStyle, coalesceMask NotificationCoalescing, modes []RunLoopMode) {
	objc.Call[objc.Void](n_, objc.Sel("enqueueNotification:postingStyle:coalesceMask:forModes:"), notification, postingStyle, coalesceMask, modes)
}

// Adds a notification to the notification queue with a specified posting style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationqueue/1416340-enqueuenotification?language=objc
func (n_ NotificationQueue) EnqueueNotificationPostingStyle(notification INotification, postingStyle PostingStyle) {
	objc.Call[objc.Void](n_, objc.Sel("enqueueNotification:postingStyle:"), notification, postingStyle)
}

// Removes all notifications from the queue that match a provided notification using provided matching criteria. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationqueue/1416688-dequeuenotificationsmatching?language=objc
func (n_ NotificationQueue) DequeueNotificationsMatchingCoalesceMask(notification INotification, coalesceMask uint) {
	objc.Call[objc.Void](n_, objc.Sel("dequeueNotificationsMatching:coalesceMask:"), notification, coalesceMask)
}

// Returns the default notification queue for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationqueue/1412392-defaultqueue?language=objc
func (nc _NotificationQueueClass) DefaultQueue() NotificationQueue {
	rv := objc.Call[NotificationQueue](nc, objc.Sel("defaultQueue"))
	return rv
}

// Returns the default notification queue for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationqueue/1412392-defaultqueue?language=objc
func NotificationQueue_DefaultQueue() NotificationQueue {
	return NotificationQueueClass.DefaultQueue()
}
