package dhcp_subnets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/rackn/rocket-skates/models"
)

// DELETEDhcpSubnetHandlerFunc turns a function with the right signature into a d e l e t e dhcp subnet handler
type DELETEDhcpSubnetHandlerFunc func(DELETEDhcpSubnetParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DELETEDhcpSubnetHandlerFunc) Handle(params DELETEDhcpSubnetParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DELETEDhcpSubnetHandler interface for that can handle valid d e l e t e dhcp subnet params
type DELETEDhcpSubnetHandler interface {
	Handle(DELETEDhcpSubnetParams, *models.Principal) middleware.Responder
}

// NewDELETEDhcpSubnet creates a new http.Handler for the d e l e t e dhcp subnet operation
func NewDELETEDhcpSubnet(ctx *middleware.Context, handler DELETEDhcpSubnetHandler) *DELETEDhcpSubnet {
	return &DELETEDhcpSubnet{Context: ctx, Handler: handler}
}

/*DELETEDhcpSubnet swagger:route DELETE /subnets/{id} Dhcp subnets dELETEDhcpSubnet

Delete DHCP Subnet

*/
type DELETEDhcpSubnet struct {
	Context *middleware.Context
	Handler DELETEDhcpSubnetHandler
}

func (o *DELETEDhcpSubnet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDELETEDhcpSubnetParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
