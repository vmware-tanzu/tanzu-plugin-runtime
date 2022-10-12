// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAWSSubnetsParams creates a new GetAWSSubnetsParams object
// with the default values initialized.
func NewGetAWSSubnetsParams() *GetAWSSubnetsParams {
	var ()
	return &GetAWSSubnetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAWSSubnetsParamsWithTimeout creates a new GetAWSSubnetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAWSSubnetsParamsWithTimeout(timeout time.Duration) *GetAWSSubnetsParams {
	var ()
	return &GetAWSSubnetsParams{

		timeout: timeout,
	}
}

// NewGetAWSSubnetsParamsWithContext creates a new GetAWSSubnetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAWSSubnetsParamsWithContext(ctx context.Context) *GetAWSSubnetsParams {
	var ()
	return &GetAWSSubnetsParams{

		Context: ctx,
	}
}

// NewGetAWSSubnetsParamsWithHTTPClient creates a new GetAWSSubnetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAWSSubnetsParamsWithHTTPClient(client *http.Client) *GetAWSSubnetsParams {
	var ()
	return &GetAWSSubnetsParams{
		HTTPClient: client,
	}
}

/*
GetAWSSubnetsParams contains all the parameters to send to the API endpoint
for the get a w s subnets operation typically these are written to a http.Request
*/
type GetAWSSubnetsParams struct {

	/*VpcID
	  VPC Id

	*/
	VpcID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get a w s subnets params
func (o *GetAWSSubnetsParams) WithTimeout(timeout time.Duration) *GetAWSSubnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get a w s subnets params
func (o *GetAWSSubnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get a w s subnets params
func (o *GetAWSSubnetsParams) WithContext(ctx context.Context) *GetAWSSubnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get a w s subnets params
func (o *GetAWSSubnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get a w s subnets params
func (o *GetAWSSubnetsParams) WithHTTPClient(client *http.Client) *GetAWSSubnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get a w s subnets params
func (o *GetAWSSubnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVpcID adds the vpcID to the get a w s subnets params
func (o *GetAWSSubnetsParams) WithVpcID(vpcID string) *GetAWSSubnetsParams {
	o.SetVpcID(vpcID)
	return o
}

// SetVpcID adds the vpcId to the get a w s subnets params
func (o *GetAWSSubnetsParams) SetVpcID(vpcID string) {
	o.VpcID = vpcID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAWSSubnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param vpcId
	qrVpcID := o.VpcID
	qVpcID := qrVpcID
	if qVpcID != "" {
		if err := r.SetQueryParam("vpcId", qVpcID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
