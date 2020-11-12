package v1

import (
	fmt "fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OAuthAccessToken describes an OAuth access token
type OAuthAccessToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// ClientName references the client that created this token.
	ClientName string `json:"clientName,omitempty" protobuf:"bytes,2,opt,name=clientName"`

	// ExpiresIn is the seconds from CreationTime before this token expires.
	ExpiresIn int64 `json:"expiresIn,omitempty" protobuf:"varint,3,opt,name=expiresIn"`

	// Scopes is an array of the requested scopes.
	Scopes []string `json:"scopes,omitempty" protobuf:"bytes,4,rep,name=scopes"`

	// RedirectURI is the redirection associated with the token.
	RedirectURI string `json:"redirectURI,omitempty" protobuf:"bytes,5,opt,name=redirectURI"`

	// UserName is the user name associated with this token
	UserName string `json:"userName,omitempty" protobuf:"bytes,6,opt,name=userName"`

	// UserUID is the unique UID associated with this token
	UserUID string `json:"userUID,omitempty" protobuf:"bytes,7,opt,name=userUID"`

	// AuthorizeToken contains the token that authorized this token
	AuthorizeToken string `json:"authorizeToken,omitempty" protobuf:"bytes,8,opt,name=authorizeToken"`

	// RefreshToken is the value by which this token can be renewed. Can be blank.
	RefreshToken string `json:"refreshToken,omitempty" protobuf:"bytes,9,opt,name=refreshToken"`

	// InactivityTimeoutSeconds is the value in seconds, from the
	// CreationTimestamp, after which this token can no longer be used.
	// The value is automatically incremented when the token is used.
	InactivityTimeoutSeconds int32 `json:"inactivityTimeoutSeconds,omitempty" protobuf:"varint,10,opt,name=inactivityTimeoutSeconds"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OAuthAuthorizeToken describes an OAuth authorization token
type OAuthAuthorizeToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// ClientName references the client that created this token.
	ClientName string `json:"clientName,omitempty" protobuf:"bytes,2,opt,name=clientName"`

	// ExpiresIn is the seconds from CreationTime before this token expires.
	ExpiresIn int64 `json:"expiresIn,omitempty" protobuf:"varint,3,opt,name=expiresIn"`

	// Scopes is an array of the requested scopes.
	Scopes []string `json:"scopes,omitempty" protobuf:"bytes,4,rep,name=scopes"`

	// RedirectURI is the redirection associated with the token.
	RedirectURI string `json:"redirectURI,omitempty" protobuf:"bytes,5,opt,name=redirectURI"`

	// State data from request
	State string `json:"state,omitempty" protobuf:"bytes,6,opt,name=state"`

	// UserName is the user name associated with this token
	UserName string `json:"userName,omitempty" protobuf:"bytes,7,opt,name=userName"`

	// UserUID is the unique UID associated with this token. UserUID and UserName must both match
	// for this token to be valid.
	UserUID string `json:"userUID,omitempty" protobuf:"bytes,8,opt,name=userUID"`

	// CodeChallenge is the optional code_challenge associated with this authorization code, as described in rfc7636
	CodeChallenge string `json:"codeChallenge,omitempty" protobuf:"bytes,9,opt,name=codeChallenge"`

	// CodeChallengeMethod is the optional code_challenge_method associated with this authorization code, as described in rfc7636
	CodeChallengeMethod string `json:"codeChallengeMethod,omitempty" protobuf:"bytes,10,opt,name=codeChallengeMethod"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OAuthClient describes an OAuth client
type OAuthClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Secret is the unique secret associated with a client
	Secret string `json:"secret,omitempty" protobuf:"bytes,2,opt,name=secret"`

	// AdditionalSecrets holds other secrets that may be used to identify the client.  This is useful for rotation
	// and for service account token validation
	AdditionalSecrets []string `json:"additionalSecrets,omitempty" protobuf:"bytes,3,rep,name=additionalSecrets"`

	// RespondWithChallenges indicates whether the client wants authentication needed responses made in the form of challenges instead of redirects
	RespondWithChallenges bool `json:"respondWithChallenges,omitempty" protobuf:"varint,4,opt,name=respondWithChallenges"`

	// RedirectURIs is the valid redirection URIs associated with a client
	// +patchStrategy=merge
	RedirectURIs []string `json:"redirectURIs,omitempty" patchStrategy:"merge" protobuf:"bytes,5,rep,name=redirectURIs"`

	// GrantMethod is a required field which determines how to handle grants for this client.
	// Valid grant handling methods are:
	//  - auto:   always approves grant requests, useful for trusted clients
	//  - prompt: prompts the end user for approval of grant requests, useful for third-party clients
	GrantMethod GrantHandlerType `json:"grantMethod,omitempty" protobuf:"bytes,6,opt,name=grantMethod,casttype=GrantHandlerType"`

	// ScopeRestrictions describes which scopes this client can request.  Each requested scope
	// is checked against each restriction.  If any restriction matches, then the scope is allowed.
	// If no restriction matches, then the scope is denied.
	ScopeRestrictions []ScopeRestriction `json:"scopeRestrictions,omitempty" protobuf:"bytes,7,rep,name=scopeRestrictions"`

	// AccessTokenMaxAgeSeconds overrides the default access token max age for tokens granted to this client.
	// 0 means no expiration.
	AccessTokenMaxAgeSeconds *int32 `json:"accessTokenMaxAgeSeconds,omitempty" protobuf:"varint,8,opt,name=accessTokenMaxAgeSeconds"`

	// AccessTokenInactivityTimeoutSeconds overrides the default token
	// inactivity timeout for tokens granted to this client.
	// The value represents the maximum amount of time that can occur between
	// consecutive uses of the token. Tokens become invalid if they are not
	// used within this temporal window. The user will need to acquire a new
	// token to regain access once a token times out.
	// This value needs to be set only if the default set in configuration is
	// not appropriate for this client. Valid values are:
	// - 0: Tokens for this client never time out
	// - X: Tokens time out if there is no activity for X seconds
	// The current minimum allowed value for X is 300 (5 minutes)
	AccessTokenInactivityTimeoutSeconds *int32 `json:"accessTokenInactivityTimeoutSeconds,omitempty" protobuf:"varint,9,opt,name=accessTokenInactivityTimeoutSeconds"`
}

type GrantHandlerType string

const (
	// GrantHandlerAuto auto-approves client authorization grant requests
	GrantHandlerAuto GrantHandlerType = "auto"
	// GrantHandlerPrompt prompts the user to approve new client authorization grant requests
	GrantHandlerPrompt GrantHandlerType = "prompt"
	// GrantHandlerDeny auto-denies client authorization grant requests
	GrantHandlerDeny GrantHandlerType = "deny"
)

// ScopeRestriction describe one restriction on scopes.  Exactly one option must be non-nil.
type ScopeRestriction struct {
	// ExactValues means the scope has to match a particular set of strings exactly
	ExactValues []string `json:"literals,omitempty" protobuf:"bytes,1,rep,name=literals"`

	// ClusterRole describes a set of restrictions for cluster role scoping.
	ClusterRole *ClusterRoleScopeRestriction `json:"clusterRole,omitempty" protobuf:"bytes,2,opt,name=clusterRole"`
}

// ClusterRoleScopeRestriction describes restrictions on cluster role scopes
type ClusterRoleScopeRestriction struct {
	// RoleNames is the list of cluster roles that can referenced.  * means anything
	RoleNames []string `json:"roleNames" protobuf:"bytes,1,rep,name=roleNames"`
	// Namespaces is the list of namespaces that can be referenced.  * means any of them (including *)
	Namespaces []string `json:"namespaces" protobuf:"bytes,2,rep,name=namespaces"`
	// AllowEscalation indicates whether you can request roles and their escalating resources
	AllowEscalation bool `json:"allowEscalation" protobuf:"varint,3,opt,name=allowEscalation"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OAuthClientAuthorization describes an authorization created by an OAuth client
type OAuthClientAuthorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// ClientName references the client that created this authorization
	ClientName string `json:"clientName,omitempty" protobuf:"bytes,2,opt,name=clientName"`

	// UserName is the user name that authorized this client
	UserName string `json:"userName,omitempty" protobuf:"bytes,3,opt,name=userName"`

	// UserUID is the unique UID associated with this authorization. UserUID and UserName
	// must both match for this authorization to be valid.
	UserUID string `json:"userUID,omitempty" protobuf:"bytes,4,opt,name=userUID"`

	// Scopes is an array of the granted scopes.
	Scopes []string `json:"scopes,omitempty" protobuf:"bytes,5,rep,name=scopes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OAuthAccessTokenList is a collection of OAuth access tokens
type OAuthAccessTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of OAuth access tokens
	Items []OAuthAccessToken `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OAuthAuthorizeTokenList is a collection of OAuth authorization tokens
type OAuthAuthorizeTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of OAuth authorization tokens
	Items []OAuthAuthorizeToken `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OAuthClientList is a collection of OAuth clients
type OAuthClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of OAuth clients
	Items []OAuthClient `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OAuthClientAuthorizationList is a collection of OAuth client authorizations
type OAuthClientAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of OAuth client authorizations
	Items []OAuthClientAuthorization `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OAuthRedirectReference is a reference to an OAuth redirect object.
type OAuthRedirectReference struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// The reference to an redirect object in the current namespace.
	Reference RedirectReference `json:"reference,omitempty" protobuf:"bytes,2,opt,name=reference"`
}

// RedirectReference specifies the target in the current namespace that resolves into redirect URIs.  Only the 'Route' kind is currently allowed.
type RedirectReference struct {
	// The group of the target that is being referred to.
	Group string `json:"group" protobuf:"bytes,1,opt,name=group"`

	// The kind of the target that is being referred to.  Currently, only 'Route' is allowed.
	Kind string `json:"kind" protobuf:"bytes,2,opt,name=kind"`

	// The name of the target that is being referred to. e.g. name of the Route.
	Name string `json:"name" protobuf:"bytes,3,opt,name=name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserOAuthAccessToken is a virtual resource to mirror OAuthAccessTokens to
// the user the access token was issued for
type UserOAuthAccessToken OAuthAccessToken

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserOAuthAccessTokenList is a collection of access tokens issued on behalf of
// the requesting user
type UserOAuthAccessTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []UserOAuthAccessToken `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type OAuthTokenReview struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec holds information about the request being evaluated
	Spec TokenReviewSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`

	// Status is filled in by the server and indicates whether the request can be authenticated.
	// +optional
	Status TokenReviewStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// TokenReviewSpec is a description of the token authentication request.
type TokenReviewSpec struct {
	// Token is the opaque bearer token.
	// +optional
	Token string `json:"token,omitempty" protobuf:"bytes,1,opt,name=token"`
	// Audiences is a list of the identifiers that the resource server presented
	// with the token identifies as. Audience-aware token authenticators will
	// verify that the token was intended for at least one of the audiences in
	// this list. If no audiences are provided, the audience will default to the
	// audience of the Kubernetes apiserver.
	// +optional
	Audiences []string `json:"audiences,omitempty" protobuf:"bytes,2,rep,name=audiences"`
}

// TokenReviewStatus is the result of the token authentication request.
type TokenReviewStatus struct {
	// Authenticated indicates that the token was associated with a known user.
	// +optional
	Authenticated bool `json:"authenticated,omitempty" protobuf:"varint,1,opt,name=authenticated"`
	// User is the UserInfo associated with the provided token.
	// +optional
	User UserInfo `json:"user,omitempty" protobuf:"bytes,2,opt,name=user"`
	// Audiences are audience identifiers chosen by the authenticator that are
	// compatible with both the TokenReview and token. An identifier is any
	// identifier in the intersection of the TokenReviewSpec audiences and the
	// token's audiences. A client of the TokenReview API that sets the
	// spec.audiences field should validate that a compatible audience identifier
	// is returned in the status.audiences field to ensure that the TokenReview
	// server is audience aware. If a TokenReview returns an empty
	// status.audience field where status.authenticated is "true", the token is
	// valid against the audience of the Kubernetes API server.
	// +optional
	Audiences []string `json:"audiences,omitempty" protobuf:"bytes,4,rep,name=audiences"`
	// Error indicates that the token couldn't be checked
	// +optional
	Error string `json:"error,omitempty" protobuf:"bytes,3,opt,name=error"`
}

// UserInfo holds the information about the user needed to implement the
// user.Info interface.
type UserInfo struct {
	// The name that uniquely identifies this user among all active users.
	// +optional
	Username string `json:"username,omitempty" protobuf:"bytes,1,opt,name=username"`
	// A unique value that identifies this user across time. If this user is
	// deleted and another user by the same name is added, they will have
	// different UIDs.
	// +optional
	UID string `json:"uid,omitempty" protobuf:"bytes,2,opt,name=uid"`
	// The names of groups this user is a part of.
	// +optional
	Groups []string `json:"groups,omitempty" protobuf:"bytes,3,rep,name=groups"`
	// Any additional information provided by the authenticator.
	// +optional
	Extra map[string]ExtraValue `json:"extra,omitempty" protobuf:"bytes,4,rep,name=extra"`
}

// ExtraValue masks the value so protobuf can generate
// +protobuf.nullable=true
// +protobuf.options.(gogoproto.goproto_stringer)=false
type ExtraValue []string

func (t ExtraValue) String() string {
	return fmt.Sprintf("%v", []string(t))
}
