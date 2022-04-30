// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package ipam

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
	"github.com/go-openapi/swag"
)

// NewPostIpamParams creates a new PostIpamParams object
// with the default values initialized.
func NewPostIpamParams() *PostIpamParams {
	var ()
	return &PostIpamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIpamParamsWithTimeout creates a new PostIpamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIpamParamsWithTimeout(timeout time.Duration) *PostIpamParams {
	var ()
	return &PostIpamParams{

		timeout: timeout,
	}
}

// NewPostIpamParamsWithContext creates a new PostIpamParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIpamParamsWithContext(ctx context.Context) *PostIpamParams {
	var ()
	return &PostIpamParams{

		Context: ctx,
	}
}

// NewPostIpamParamsWithHTTPClient creates a new PostIpamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIpamParamsWithHTTPClient(client *http.Client) *PostIpamParams {
	var ()
	return &PostIpamParams{
		HTTPClient: client,
	}
}

/*PostIpamParams contains all the parameters to send to the API endpoint
for the post ipam operation typically these are written to a http.Request
*/
type PostIpamParams struct {

	/*Expiration*/
	Expiration *bool
	/*Family*/
	Family *string
	/*Owner*/
	Owner *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ipam params
func (o *PostIpamParams) WithTimeout(timeout time.Duration) *PostIpamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ipam params
func (o *PostIpamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ipam params
func (o *PostIpamParams) WithContext(ctx context.Context) *PostIpamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ipam params
func (o *PostIpamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ipam params
func (o *PostIpamParams) WithHTTPClient(client *http.Client) *PostIpamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ipam params
func (o *PostIpamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpiration adds the expiration to the post ipam params
func (o *PostIpamParams) WithExpiration(expiration *bool) *PostIpamParams {
	o.SetExpiration(expiration)
	return o
}

// SetExpiration adds the expiration to the post ipam params
func (o *PostIpamParams) SetExpiration(expiration *bool) {
	o.Expiration = expiration
}

// WithFamily adds the family to the post ipam params
func (o *PostIpamParams) WithFamily(family *string) *PostIpamParams {
	o.SetFamily(family)
	return o
}

// SetFamily adds the family to the post ipam params
func (o *PostIpamParams) SetFamily(family *string) {
	o.Family = family
}

// WithOwner adds the owner to the post ipam params
func (o *PostIpamParams) WithOwner(owner *string) *PostIpamParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the post ipam params
func (o *PostIpamParams) SetOwner(owner *string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *PostIpamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expiration != nil {

		// header param expiration
		if err := r.SetHeaderParam("expiration", swag.FormatBool(*o.Expiration)); err != nil {
			return err
		}

	}

	if o.Family != nil {

		// query param family
		var qrFamily string
		if o.Family != nil {
			qrFamily = *o.Family
		}
		qFamily := qrFamily
		if qFamily != "" {
			if err := r.SetQueryParam("family", qFamily); err != nil {
				return err
			}
		}

	}

	if o.Owner != nil {

		// query param owner
		var qrOwner string
		if o.Owner != nil {
			qrOwner = *o.Owner
		}
		qOwner := qrOwner
		if qOwner != "" {
			if err := r.SetQueryParam("owner", qOwner); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
