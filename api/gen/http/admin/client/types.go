// Code generated by goa v3.7.10, DO NOT EDIT.
//
// admin HTTP client types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	admin "github.com/tektoncd/hub/api/gen/admin"
	goa "goa.design/goa/v3/pkg"
)

// UpdateAgentRequestBody is the type of the "admin" service "UpdateAgent"
// endpoint HTTP request body.
type UpdateAgentRequestBody struct {
	// Name of Agent
	Name string `form:"name" json:"name" xml:"name"`
	// Scopes required for Agent
	Scopes []string `form:"scopes" json:"scopes" xml:"scopes"`
}

// UpdateAgentResponseBody is the type of the "admin" service "UpdateAgent"
// endpoint HTTP response body.
type UpdateAgentResponseBody struct {
	// Agent JWT
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

// RefreshConfigResponseBody is the type of the "admin" service "RefreshConfig"
// endpoint HTTP response body.
type RefreshConfigResponseBody struct {
	// Config file checksum
	Checksum *string `form:"checksum,omitempty" json:"checksum,omitempty" xml:"checksum,omitempty"`
}

// UpdateAgentInvalidPayloadResponseBody is the type of the "admin" service
// "UpdateAgent" endpoint HTTP response body for the "invalid-payload" error.
type UpdateAgentInvalidPayloadResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateAgentInvalidTokenResponseBody is the type of the "admin" service
// "UpdateAgent" endpoint HTTP response body for the "invalid-token" error.
type UpdateAgentInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateAgentInvalidScopesResponseBody is the type of the "admin" service
// "UpdateAgent" endpoint HTTP response body for the "invalid-scopes" error.
type UpdateAgentInvalidScopesResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateAgentInternalErrorResponseBody is the type of the "admin" service
// "UpdateAgent" endpoint HTTP response body for the "internal-error" error.
type UpdateAgentInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// RefreshConfigInvalidTokenResponseBody is the type of the "admin" service
// "RefreshConfig" endpoint HTTP response body for the "invalid-token" error.
type RefreshConfigInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// RefreshConfigInvalidScopesResponseBody is the type of the "admin" service
// "RefreshConfig" endpoint HTTP response body for the "invalid-scopes" error.
type RefreshConfigInvalidScopesResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// RefreshConfigInternalErrorResponseBody is the type of the "admin" service
// "RefreshConfig" endpoint HTTP response body for the "internal-error" error.
type RefreshConfigInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// NewUpdateAgentRequestBody builds the HTTP request body from the payload of
// the "UpdateAgent" endpoint of the "admin" service.
func NewUpdateAgentRequestBody(p *admin.UpdateAgentPayload) *UpdateAgentRequestBody {
	body := &UpdateAgentRequestBody{
		Name: p.Name,
	}
	if p.Scopes != nil {
		body.Scopes = make([]string, len(p.Scopes))
		for i, val := range p.Scopes {
			body.Scopes[i] = val
		}
	}
	return body
}

// NewUpdateAgentResultOK builds a "admin" service "UpdateAgent" endpoint
// result from a HTTP "OK" response.
func NewUpdateAgentResultOK(body *UpdateAgentResponseBody) *admin.UpdateAgentResult {
	v := &admin.UpdateAgentResult{
		Token: *body.Token,
	}

	return v
}

// NewUpdateAgentInvalidPayload builds a admin service UpdateAgent endpoint
// invalid-payload error.
func NewUpdateAgentInvalidPayload(body *UpdateAgentInvalidPayloadResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateAgentInvalidToken builds a admin service UpdateAgent endpoint
// invalid-token error.
func NewUpdateAgentInvalidToken(body *UpdateAgentInvalidTokenResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateAgentInvalidScopes builds a admin service UpdateAgent endpoint
// invalid-scopes error.
func NewUpdateAgentInvalidScopes(body *UpdateAgentInvalidScopesResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateAgentInternalError builds a admin service UpdateAgent endpoint
// internal-error error.
func NewUpdateAgentInternalError(body *UpdateAgentInternalErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewRefreshConfigResultOK builds a "admin" service "RefreshConfig" endpoint
// result from a HTTP "OK" response.
func NewRefreshConfigResultOK(body *RefreshConfigResponseBody) *admin.RefreshConfigResult {
	v := &admin.RefreshConfigResult{
		Checksum: *body.Checksum,
	}

	return v
}

// NewRefreshConfigInvalidToken builds a admin service RefreshConfig endpoint
// invalid-token error.
func NewRefreshConfigInvalidToken(body *RefreshConfigInvalidTokenResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewRefreshConfigInvalidScopes builds a admin service RefreshConfig endpoint
// invalid-scopes error.
func NewRefreshConfigInvalidScopes(body *RefreshConfigInvalidScopesResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewRefreshConfigInternalError builds a admin service RefreshConfig endpoint
// internal-error error.
func NewRefreshConfigInternalError(body *RefreshConfigInternalErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateUpdateAgentResponseBody runs the validations defined on
// UpdateAgentResponseBody
func ValidateUpdateAgentResponseBody(body *UpdateAgentResponseBody) (err error) {
	if body.Token == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("token", "body"))
	}
	return
}

// ValidateRefreshConfigResponseBody runs the validations defined on
// RefreshConfigResponseBody
func ValidateRefreshConfigResponseBody(body *RefreshConfigResponseBody) (err error) {
	if body.Checksum == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("checksum", "body"))
	}
	return
}

// ValidateUpdateAgentInvalidPayloadResponseBody runs the validations defined
// on UpdateAgent_invalid-payload_Response_Body
func ValidateUpdateAgentInvalidPayloadResponseBody(body *UpdateAgentInvalidPayloadResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateAgentInvalidTokenResponseBody runs the validations defined on
// UpdateAgent_invalid-token_Response_Body
func ValidateUpdateAgentInvalidTokenResponseBody(body *UpdateAgentInvalidTokenResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateAgentInvalidScopesResponseBody runs the validations defined on
// UpdateAgent_invalid-scopes_Response_Body
func ValidateUpdateAgentInvalidScopesResponseBody(body *UpdateAgentInvalidScopesResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateAgentInternalErrorResponseBody runs the validations defined on
// UpdateAgent_internal-error_Response_Body
func ValidateUpdateAgentInternalErrorResponseBody(body *UpdateAgentInternalErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateRefreshConfigInvalidTokenResponseBody runs the validations defined
// on RefreshConfig_invalid-token_Response_Body
func ValidateRefreshConfigInvalidTokenResponseBody(body *RefreshConfigInvalidTokenResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateRefreshConfigInvalidScopesResponseBody runs the validations defined
// on RefreshConfig_invalid-scopes_Response_Body
func ValidateRefreshConfigInvalidScopesResponseBody(body *RefreshConfigInvalidScopesResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateRefreshConfigInternalErrorResponseBody runs the validations defined
// on RefreshConfig_internal-error_Response_Body
func ValidateRefreshConfigInternalErrorResponseBody(body *RefreshConfigInternalErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}
