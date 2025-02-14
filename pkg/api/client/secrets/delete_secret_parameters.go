// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteSecretParams creates a new DeleteSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSecretParams() *DeleteSecretParams {
	return &DeleteSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSecretParamsWithTimeout creates a new DeleteSecretParams object
// with the ability to set a timeout on a request.
func NewDeleteSecretParamsWithTimeout(timeout time.Duration) *DeleteSecretParams {
	return &DeleteSecretParams{
		timeout: timeout,
	}
}

// NewDeleteSecretParamsWithContext creates a new DeleteSecretParams object
// with the ability to set a context for a request.
func NewDeleteSecretParamsWithContext(ctx context.Context) *DeleteSecretParams {
	return &DeleteSecretParams{
		Context: ctx,
	}
}

// NewDeleteSecretParamsWithHTTPClient creates a new DeleteSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSecretParamsWithHTTPClient(client *http.Client) *DeleteSecretParams {
	return &DeleteSecretParams{
		HTTPClient: client,
	}
}

/* DeleteSecretParams contains all the parameters to send to the API endpoint
   for the delete secret operation.

   Typically these are written to a http.Request.
*/
type DeleteSecretParams struct {

	/* Namespace.

	   target namespace
	*/
	Namespace string

	/* Secret.

	   target secret
	*/
	Secret string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSecretParams) WithDefaults() *DeleteSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete secret params
func (o *DeleteSecretParams) WithTimeout(timeout time.Duration) *DeleteSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete secret params
func (o *DeleteSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete secret params
func (o *DeleteSecretParams) WithContext(ctx context.Context) *DeleteSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete secret params
func (o *DeleteSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete secret params
func (o *DeleteSecretParams) WithHTTPClient(client *http.Client) *DeleteSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete secret params
func (o *DeleteSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the delete secret params
func (o *DeleteSecretParams) WithNamespace(namespace string) *DeleteSecretParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete secret params
func (o *DeleteSecretParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithSecret adds the secret to the delete secret params
func (o *DeleteSecretParams) WithSecret(secret string) *DeleteSecretParams {
	o.SetSecret(secret)
	return o
}

// SetSecret adds the secret to the delete secret params
func (o *DeleteSecretParams) SetSecret(secret string) {
	o.Secret = secret
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param secret
	if err := r.SetPathParam("secret", o.Secret); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
