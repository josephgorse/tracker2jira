// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/king-jam/tracker2jira/rest/models"
)

// GetTasksReader is a Reader for the GetTasks structure.
type GetTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTasksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTasksOK creates a GetTasksOK with default headers values
func NewGetTasksOK() *GetTasksOK {
	return &GetTasksOK{}
}

/*GetTasksOK handles this case with default header values.

The list of current tasks
*/
type GetTasksOK struct {
	Payload models.Tasks
}

func (o *GetTasksOK) Error() string {
	return fmt.Sprintf("[GET /tasks][%d] getTasksOK  %+v", 200, o.Payload)
}

func (o *GetTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksBadRequest creates a GetTasksBadRequest with default headers values
func NewGetTasksBadRequest() *GetTasksBadRequest {
	return &GetTasksBadRequest{}
}

/*GetTasksBadRequest handles this case with default header values.

Bad Request
*/
type GetTasksBadRequest struct {
}

func (o *GetTasksBadRequest) Error() string {
	return fmt.Sprintf("[GET /tasks][%d] getTasksBadRequest ", 400)
}

func (o *GetTasksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}