// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/netlify/open-api/go/models"
)

// ListSiteDeploysReader is a Reader for the ListSiteDeploys structure.
type ListSiteDeploysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSiteDeploysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSiteDeploysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListSiteDeploysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSiteDeploysOK creates a ListSiteDeploysOK with default headers values
func NewListSiteDeploysOK() *ListSiteDeploysOK {
	return &ListSiteDeploysOK{}
}

/*ListSiteDeploysOK handles this case with default header values.

OK
*/
type ListSiteDeploysOK struct {
	Payload []*models.Deploy
}

func (o *ListSiteDeploysOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/deploys][%d] listSiteDeploysOK  %+v", 200, o.Payload)
}

func (o *ListSiteDeploysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSiteDeploysDefault creates a ListSiteDeploysDefault with default headers values
func NewListSiteDeploysDefault(code int) *ListSiteDeploysDefault {
	return &ListSiteDeploysDefault{
		_statusCode: code,
	}
}

/*ListSiteDeploysDefault handles this case with default header values.

error
*/
type ListSiteDeploysDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list site deploys default response
func (o *ListSiteDeploysDefault) Code() int {
	return o._statusCode
}

func (o *ListSiteDeploysDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/deploys][%d] listSiteDeploys default  %+v", o._statusCode, o.Payload)
}

func (o *ListSiteDeploysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
