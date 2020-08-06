/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

// TokenAuthorizationRequestBuilder contains the data and logic needed to build 'token_authorization_request' objects.
//
//
type TokenAuthorizationRequestBuilder struct {
	authorizationToken *string
}

// NewTokenAuthorizationRequest creates a new builder of 'token_authorization_request' objects.
func NewTokenAuthorizationRequest() *TokenAuthorizationRequestBuilder {
	return new(TokenAuthorizationRequestBuilder)
}

// AuthorizationToken sets the value of the 'authorization_token' attribute to the given value.
//
//
func (b *TokenAuthorizationRequestBuilder) AuthorizationToken(value string) *TokenAuthorizationRequestBuilder {
	b.authorizationToken = &value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *TokenAuthorizationRequestBuilder) Copy(object *TokenAuthorizationRequest) *TokenAuthorizationRequestBuilder {
	if object == nil {
		return b
	}
	b.authorizationToken = object.authorizationToken
	return b
}

// Build creates a 'token_authorization_request' object using the configuration stored in the builder.
func (b *TokenAuthorizationRequestBuilder) Build() (object *TokenAuthorizationRequest, err error) {
	object = new(TokenAuthorizationRequest)
	object.authorizationToken = b.authorizationToken
	return
}
