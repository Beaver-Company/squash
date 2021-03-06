// Code generated by go-swagger; DO NOT EDIT.

package debugattachment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/solo-io/squash/pkg/models"
)

// AddDebugAttachmentReader is a Reader for the AddDebugAttachment structure.
type AddDebugAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDebugAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddDebugAttachmentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddDebugAttachmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAddDebugAttachmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewAddDebugAttachmentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewAddDebugAttachmentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddDebugAttachmentCreated creates a AddDebugAttachmentCreated with default headers values
func NewAddDebugAttachmentCreated() *AddDebugAttachmentCreated {
	return &AddDebugAttachmentCreated{}
}

/*AddDebugAttachmentCreated handles this case with default header values.

Debug attachment created
*/
type AddDebugAttachmentCreated struct {
	Payload *models.DebugAttachment
}

func (o *AddDebugAttachmentCreated) Error() string {
	return fmt.Sprintf("[POST /debugattachment][%d] addDebugAttachmentCreated  %+v", 201, o.Payload)
}

func (o *AddDebugAttachmentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DebugAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDebugAttachmentBadRequest creates a AddDebugAttachmentBadRequest with default headers values
func NewAddDebugAttachmentBadRequest() *AddDebugAttachmentBadRequest {
	return &AddDebugAttachmentBadRequest{}
}

/*AddDebugAttachmentBadRequest handles this case with default header values.

Bad request
*/
type AddDebugAttachmentBadRequest struct {
}

func (o *AddDebugAttachmentBadRequest) Error() string {
	return fmt.Sprintf("[POST /debugattachment][%d] addDebugAttachmentBadRequest ", 400)
}

func (o *AddDebugAttachmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddDebugAttachmentNotFound creates a AddDebugAttachmentNotFound with default headers values
func NewAddDebugAttachmentNotFound() *AddDebugAttachmentNotFound {
	return &AddDebugAttachmentNotFound{}
}

/*AddDebugAttachmentNotFound handles this case with default header values.

Not found
*/
type AddDebugAttachmentNotFound struct {
}

func (o *AddDebugAttachmentNotFound) Error() string {
	return fmt.Sprintf("[POST /debugattachment][%d] addDebugAttachmentNotFound ", 404)
}

func (o *AddDebugAttachmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddDebugAttachmentUnprocessableEntity creates a AddDebugAttachmentUnprocessableEntity with default headers values
func NewAddDebugAttachmentUnprocessableEntity() *AddDebugAttachmentUnprocessableEntity {
	return &AddDebugAttachmentUnprocessableEntity{}
}

/*AddDebugAttachmentUnprocessableEntity handles this case with default header values.

Invalid input
*/
type AddDebugAttachmentUnprocessableEntity struct {
}

func (o *AddDebugAttachmentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /debugattachment][%d] addDebugAttachmentUnprocessableEntity ", 422)
}

func (o *AddDebugAttachmentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddDebugAttachmentServiceUnavailable creates a AddDebugAttachmentServiceUnavailable with default headers values
func NewAddDebugAttachmentServiceUnavailable() *AddDebugAttachmentServiceUnavailable {
	return &AddDebugAttachmentServiceUnavailable{}
}

/*AddDebugAttachmentServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type AddDebugAttachmentServiceUnavailable struct {
}

func (o *AddDebugAttachmentServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /debugattachment][%d] addDebugAttachmentServiceUnavailable ", 503)
}

func (o *AddDebugAttachmentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
