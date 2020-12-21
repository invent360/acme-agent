// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/acme-agent/swagger/models"
)

// GetReader is a Reader for the Get structure.
type GetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOK creates a GetOK with default headers values
func NewGetOK() *GetOK {
	return &GetOK{}
}

/*GetOK handles this case with default header values.

Post get
*/
type GetOK struct {
	Payload *models.Post
}

func (o *GetOK) Error() string {
	return fmt.Sprintf("[GET /posts/{id}][%d] getOK  %+v", 200, o.Payload)
}

func (o *GetOK) GetPayload() *models.Post {
	return o.Payload
}

func (o *GetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Post)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBadRequest creates a GetBadRequest with default headers values
func NewGetBadRequest() *GetBadRequest {
	return &GetBadRequest{}
}

/*GetBadRequest handles this case with default header values.

Bad Request
*/
type GetBadRequest struct {
}

func (o *GetBadRequest) Error() string {
	return fmt.Sprintf("[GET /posts/{id}][%d] getBadRequest ", 400)
}

func (o *GetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNotFound creates a GetNotFound with default headers values
func NewGetNotFound() *GetNotFound {
	return &GetNotFound{}
}

/*GetNotFound handles this case with default header values.

Post Not Found
*/
type GetNotFound struct {
}

func (o *GetNotFound) Error() string {
	return fmt.Sprintf("[GET /posts/{id}][%d] getNotFound ", 404)
}

func (o *GetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
