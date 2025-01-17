/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version: 1.0.0
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// LoginViaApiResponse The Response for Login Flows via API
type LoginViaApiResponse struct {
	Session Session `json:"session"`
	// The Session Token  A session token is equivalent to a session cookie, but it can be sent in the HTTP Authorization Header:  Authorization: bearer ${session-token}  The session token is only issued for API flows, not for Browser flows!
	SessionToken string `json:"session_token"`
}

// NewLoginViaApiResponse instantiates a new LoginViaApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginViaApiResponse(session Session, sessionToken string) *LoginViaApiResponse {
	this := LoginViaApiResponse{}
	this.Session = session
	this.SessionToken = sessionToken
	return &this
}

// NewLoginViaApiResponseWithDefaults instantiates a new LoginViaApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginViaApiResponseWithDefaults() *LoginViaApiResponse {
	this := LoginViaApiResponse{}
	return &this
}

// GetSession returns the Session field value
func (o *LoginViaApiResponse) GetSession() Session {
	if o == nil {
		var ret Session
		return ret
	}

	return o.Session
}

// GetSessionOk returns a tuple with the Session field value
// and a boolean to check if the value has been set.
func (o *LoginViaApiResponse) GetSessionOk() (*Session, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Session, true
}

// SetSession sets field value
func (o *LoginViaApiResponse) SetSession(v Session) {
	o.Session = v
}

// GetSessionToken returns the SessionToken field value
func (o *LoginViaApiResponse) GetSessionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value
// and a boolean to check if the value has been set.
func (o *LoginViaApiResponse) GetSessionTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionToken, true
}

// SetSessionToken sets field value
func (o *LoginViaApiResponse) SetSessionToken(v string) {
	o.SessionToken = v
}

func (o LoginViaApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["session"] = o.Session
	}
	if true {
		toSerialize["session_token"] = o.SessionToken
	}
	return json.Marshal(toSerialize)
}

type NullableLoginViaApiResponse struct {
	value *LoginViaApiResponse
	isSet bool
}

func (v NullableLoginViaApiResponse) Get() *LoginViaApiResponse {
	return v.value
}

func (v *NullableLoginViaApiResponse) Set(val *LoginViaApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginViaApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginViaApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginViaApiResponse(val *LoginViaApiResponse) *NullableLoginViaApiResponse {
	return &NullableLoginViaApiResponse{value: val, isSet: true}
}

func (v NullableLoginViaApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginViaApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
