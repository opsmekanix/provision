package dhcp_subnets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/rackn/rocket-skates/models"
)

// GETDhcpSubnetHandlerFunc turns a function with the right signature into a g e t dhcp subnet handler
type GETDhcpSubnetHandlerFunc func(GETDhcpSubnetParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GETDhcpSubnetHandlerFunc) Handle(params GETDhcpSubnetParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GETDhcpSubnetHandler interface for that can handle valid g e t dhcp subnet params
type GETDhcpSubnetHandler interface {
	Handle(GETDhcpSubnetParams, *models.Principal) middleware.Responder
}

// NewGETDhcpSubnet creates a new http.Handler for the g e t dhcp subnet operation
func NewGETDhcpSubnet(ctx *middleware.Context, handler GETDhcpSubnetHandler) *GETDhcpSubnet {
	return &GETDhcpSubnet{Context: ctx, Handler: handler}
}

/*GETDhcpSubnet swagger:route GET /subnets/{id} Dhcp subnets gETDhcpSubnet

Get DHCP Subnet

*/
type GETDhcpSubnet struct {
	Context *middleware.Context
	Handler GETDhcpSubnetHandler
}

func (o *GETDhcpSubnet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGETDhcpSubnetParams()

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
