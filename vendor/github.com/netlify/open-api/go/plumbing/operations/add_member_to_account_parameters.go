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

// NewAddMemberToAccountParams creates a new AddMemberToAccountParams object
// with the default values initialized.
func NewAddMemberToAccountParams() *AddMemberToAccountParams {
	var ()
	return &AddMemberToAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddMemberToAccountParamsWithTimeout creates a new AddMemberToAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddMemberToAccountParamsWithTimeout(timeout time.Duration) *AddMemberToAccountParams {
	var ()
	return &AddMemberToAccountParams{

		timeout: timeout,
	}
}

// NewAddMemberToAccountParamsWithContext creates a new AddMemberToAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddMemberToAccountParamsWithContext(ctx context.Context) *AddMemberToAccountParams {
	var ()
	return &AddMemberToAccountParams{

		Context: ctx,
	}
}

// NewAddMemberToAccountParamsWithHTTPClient creates a new AddMemberToAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddMemberToAccountParamsWithHTTPClient(client *http.Client) *AddMemberToAccountParams {
	var ()
	return &AddMemberToAccountParams{
		HTTPClient: client,
	}
}

/*AddMemberToAccountParams contains all the parameters to send to the API endpoint
for the add member to account operation typically these are written to a http.Request
*/
type AddMemberToAccountParams struct {

	/*AccountSlug*/
	AccountSlug string
	/*Email*/
	Email string
	/*Role*/
	Role *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add member to account params
func (o *AddMemberToAccountParams) WithTimeout(timeout time.Duration) *AddMemberToAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add member to account params
func (o *AddMemberToAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add member to account params
func (o *AddMemberToAccountParams) WithContext(ctx context.Context) *AddMemberToAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add member to account params
func (o *AddMemberToAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add member to account params
func (o *AddMemberToAccountParams) WithHTTPClient(client *http.Client) *AddMemberToAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add member to account params
func (o *AddMemberToAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountSlug adds the accountSlug to the add member to account params
func (o *AddMemberToAccountParams) WithAccountSlug(accountSlug string) *AddMemberToAccountParams {
	o.SetAccountSlug(accountSlug)
	return o
}

// SetAccountSlug adds the accountSlug to the add member to account params
func (o *AddMemberToAccountParams) SetAccountSlug(accountSlug string) {
	o.AccountSlug = accountSlug
}

// WithEmail adds the email to the add member to account params
func (o *AddMemberToAccountParams) WithEmail(email string) *AddMemberToAccountParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the add member to account params
func (o *AddMemberToAccountParams) SetEmail(email string) {
	o.Email = email
}

// WithRole adds the role to the add member to account params
func (o *AddMemberToAccountParams) WithRole(role *string) *AddMemberToAccountParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the add member to account params
func (o *AddMemberToAccountParams) SetRole(role *string) {
	o.Role = role
}

// WriteToRequest writes these params to a swagger request
func (o *AddMemberToAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_slug
	if err := r.SetPathParam("account_slug", o.AccountSlug); err != nil {
		return err
	}

	// query param email
	qrEmail := o.Email
	qEmail := qrEmail
	if qEmail != "" {
		if err := r.SetQueryParam("email", qEmail); err != nil {
			return err
		}
	}

	if o.Role != nil {

		// query param role
		var qrRole string
		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {
			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
