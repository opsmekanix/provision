package dhcp_subnets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

/*GETDhcpSubnetOK g e t dhcp subnet o k

swagger:response gETDhcpSubnetOK
*/
type GETDhcpSubnetOK struct {

	/*
	  In: Body
	*/
	Payload *models.DhcpSubnetInput `json:"body,omitempty"`
}

// NewGETDhcpSubnetOK creates GETDhcpSubnetOK with default headers values
func NewGETDhcpSubnetOK() *GETDhcpSubnetOK {
	return &GETDhcpSubnetOK{}
}

// WithPayload adds the payload to the g e t dhcp subnet o k response
func (o *GETDhcpSubnetOK) WithPayload(payload *models.DhcpSubnetInput) *GETDhcpSubnetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the g e t dhcp subnet o k response
func (o *GETDhcpSubnetOK) SetPayload(payload *models.DhcpSubnetInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GETDhcpSubnetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
