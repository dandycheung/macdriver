// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that delegates of a URL connection implement to receive status about and provide feedback to the connection object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate?language=objc
type PURLConnectionDelegate interface {
	// optional
	ConnectionShouldUseCredentialStorage(connection URLConnection) bool
	HasConnectionShouldUseCredentialStorage() bool

	// optional
	ConnectionDidFailWithError(connection URLConnection, error Error)
	HasConnectionDidFailWithError() bool

	// optional
	ConnectionWillSendRequestForAuthenticationChallenge(connection URLConnection, challenge URLAuthenticationChallenge)
	HasConnectionWillSendRequestForAuthenticationChallenge() bool
}

// A delegate implementation builder for the [PURLConnectionDelegate] protocol.
type URLConnectionDelegate struct {
	_ConnectionShouldUseCredentialStorage                func(connection URLConnection) bool
	_ConnectionDidFailWithError                          func(connection URLConnection, error Error)
	_ConnectionWillSendRequestForAuthenticationChallenge func(connection URLConnection, challenge URLAuthenticationChallenge)
}

func (di *URLConnectionDelegate) HasConnectionShouldUseCredentialStorage() bool {
	return di._ConnectionShouldUseCredentialStorage != nil
}

// Sent to determine whether the URL loader should use the credential storage for authenticating the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate/1414890-connectionshouldusecredentialsto?language=objc
func (di *URLConnectionDelegate) SetConnectionShouldUseCredentialStorage(f func(connection URLConnection) bool) {
	di._ConnectionShouldUseCredentialStorage = f
}

// Sent to determine whether the URL loader should use the credential storage for authenticating the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate/1414890-connectionshouldusecredentialsto?language=objc
func (di *URLConnectionDelegate) ConnectionShouldUseCredentialStorage(connection URLConnection) bool {
	return di._ConnectionShouldUseCredentialStorage(connection)
}
func (di *URLConnectionDelegate) HasConnectionDidFailWithError() bool {
	return di._ConnectionDidFailWithError != nil
}

// Sent when a connection fails to load its request successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate/1418443-connection?language=objc
func (di *URLConnectionDelegate) SetConnectionDidFailWithError(f func(connection URLConnection, error Error)) {
	di._ConnectionDidFailWithError = f
}

// Sent when a connection fails to load its request successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate/1418443-connection?language=objc
func (di *URLConnectionDelegate) ConnectionDidFailWithError(connection URLConnection, error Error) {
	di._ConnectionDidFailWithError(connection, error)
}
func (di *URLConnectionDelegate) HasConnectionWillSendRequestForAuthenticationChallenge() bool {
	return di._ConnectionWillSendRequestForAuthenticationChallenge != nil
}

// Tells the delegate that the connection will send a request for an authentication challenge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate/1414078-connection?language=objc
func (di *URLConnectionDelegate) SetConnectionWillSendRequestForAuthenticationChallenge(f func(connection URLConnection, challenge URLAuthenticationChallenge)) {
	di._ConnectionWillSendRequestForAuthenticationChallenge = f
}

// Tells the delegate that the connection will send a request for an authentication challenge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate/1414078-connection?language=objc
func (di *URLConnectionDelegate) ConnectionWillSendRequestForAuthenticationChallenge(connection URLConnection, challenge URLAuthenticationChallenge) {
	di._ConnectionWillSendRequestForAuthenticationChallenge(connection, challenge)
}

// ensure impl type implements protocol interface
var _ PURLConnectionDelegate = (*URLConnectionDelegateObject)(nil)

// A concrete type for the [PURLConnectionDelegate] protocol.
type URLConnectionDelegateObject struct {
	objc.Object
}

func (u_ URLConnectionDelegateObject) HasConnectionShouldUseCredentialStorage() bool {
	return u_.RespondsToSelector(objc.Sel("connectionShouldUseCredentialStorage:"))
}

// Sent to determine whether the URL loader should use the credential storage for authenticating the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate/1414890-connectionshouldusecredentialsto?language=objc
func (u_ URLConnectionDelegateObject) ConnectionShouldUseCredentialStorage(connection URLConnection) bool {
	rv := objc.Call[bool](u_, objc.Sel("connectionShouldUseCredentialStorage:"), connection)
	return rv
}

func (u_ URLConnectionDelegateObject) HasConnectionDidFailWithError() bool {
	return u_.RespondsToSelector(objc.Sel("connection:didFailWithError:"))
}

// Sent when a connection fails to load its request successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate/1418443-connection?language=objc
func (u_ URLConnectionDelegateObject) ConnectionDidFailWithError(connection URLConnection, error Error) {
	objc.Call[objc.Void](u_, objc.Sel("connection:didFailWithError:"), connection, error)
}

func (u_ URLConnectionDelegateObject) HasConnectionWillSendRequestForAuthenticationChallenge() bool {
	return u_.RespondsToSelector(objc.Sel("connection:willSendRequestForAuthenticationChallenge:"))
}

// Tells the delegate that the connection will send a request for an authentication challenge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate/1414078-connection?language=objc
func (u_ URLConnectionDelegateObject) ConnectionWillSendRequestForAuthenticationChallenge(connection URLConnection, challenge URLAuthenticationChallenge) {
	objc.Call[objc.Void](u_, objc.Sel("connection:willSendRequestForAuthenticationChallenge:"), connection, challenge)
}
