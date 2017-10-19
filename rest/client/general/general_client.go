// Code generated by go-swagger; DO NOT EDIT.

package general

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new general API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for general API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetConfig gets the current running configuration
*/
func (a *Client) GetConfig(params *GetConfigParams) (*GetConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetConfig",
		Method:             "GET",
		PathPattern:        "/config",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConfigOK), nil

}

/*
PatchConfig updates the current running configuration
*/
func (a *Client) PatchConfig(params *PatchConfigParams) (*PatchConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchConfig",
		Method:             "PATCH",
		PathPattern:        "/config",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchConfigOK), nil

}

/*
Login logs in and returns authentication type based on parameters
*/
func (a *Client) Login(params *LoginParams) (*LoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "login",
		Method:             "POST",
		PathPattern:        "/login",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LoginReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LoginOK), nil

}

/*
Root returns the default home page
*/
func (a *Client) Root(params *RootParams) (*RootOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRootParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "root",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RootReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RootOK), nil

}

/*
Version returns the current version running
*/
func (a *Client) Version(params *VersionParams) (*VersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "version",
		Method:             "GET",
		PathPattern:        "/version",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VersionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
