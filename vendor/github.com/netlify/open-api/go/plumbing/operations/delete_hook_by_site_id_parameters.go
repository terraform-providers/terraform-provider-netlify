// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteHookBySiteIDParams creates a new DeleteHookBySiteIDParams object
// with the default values initialized.
func NewDeleteHookBySiteIDParams() *DeleteHookBySiteIDParams {
	var ()
	return &DeleteHookBySiteIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteHookBySiteIDParamsWithTimeout creates a new DeleteHookBySiteIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteHookBySiteIDParamsWithTimeout(timeout time.Duration) *DeleteHookBySiteIDParams {
	var ()
	return &DeleteHookBySiteIDParams{

		timeout: timeout,
	}
}

// NewDeleteHookBySiteIDParamsWithContext creates a new DeleteHookBySiteIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteHookBySiteIDParamsWithContext(ctx context.Context) *DeleteHookBySiteIDParams {
	var ()
	return &DeleteHookBySiteIDParams{

		Context: ctx,
	}
}

// NewDeleteHookBySiteIDParamsWithHTTPClient creates a new DeleteHookBySiteIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteHookBySiteIDParamsWithHTTPClient(client *http.Client) *DeleteHookBySiteIDParams {
	var ()
	return &DeleteHookBySiteIDParams{
		HTTPClient: client,
	}
}

/*DeleteHookBySiteIDParams contains all the parameters to send to the API endpoint
for the delete hook by site Id operation typically these are written to a http.Request
*/
type DeleteHookBySiteIDParams struct {

	/*HookID*/
	HookID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete hook by site Id params
func (o *DeleteHookBySiteIDParams) WithTimeout(timeout time.Duration) *DeleteHookBySiteIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete hook by site Id params
func (o *DeleteHookBySiteIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete hook by site Id params
func (o *DeleteHookBySiteIDParams) WithContext(ctx context.Context) *DeleteHookBySiteIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete hook by site Id params
func (o *DeleteHookBySiteIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete hook by site Id params
func (o *DeleteHookBySiteIDParams) WithHTTPClient(client *http.Client) *DeleteHookBySiteIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete hook by site Id params
func (o *DeleteHookBySiteIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHookID adds the hookID to the delete hook by site Id params
func (o *DeleteHookBySiteIDParams) WithHookID(hookID string) *DeleteHookBySiteIDParams {
	o.SetHookID(hookID)
	return o
}

// SetHookID adds the hookId to the delete hook by site Id params
func (o *DeleteHookBySiteIDParams) SetHookID(hookID string) {
	o.HookID = hookID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteHookBySiteIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param hook_id
	if err := r.SetPathParam("hook_id", o.HookID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
