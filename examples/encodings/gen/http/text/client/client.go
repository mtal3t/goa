// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// text client HTTP transport
//
// Command:
// $ goa gen goa.design/goa/examples/encodings/design -o
// $(GOPATH)/src/goa.design/goa/examples/encodings

package client

import (
	"context"
	"net/http"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Client lists the text service endpoint HTTP clients.
type Client struct {
	// Concatstrings Doer is the HTTP client used to make requests to the
	// concatstrings endpoint.
	ConcatstringsDoer goahttp.Doer

	// Concatbytes Doer is the HTTP client used to make requests to the concatbytes
	// endpoint.
	ConcatbytesDoer goahttp.Doer

	// Concatstringfield Doer is the HTTP client used to make requests to the
	// concatstringfield endpoint.
	ConcatstringfieldDoer goahttp.Doer

	// Concatbytesfield Doer is the HTTP client used to make requests to the
	// concatbytesfield endpoint.
	ConcatbytesfieldDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the text service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ConcatstringsDoer:     doer,
		ConcatbytesDoer:       doer,
		ConcatstringfieldDoer: doer,
		ConcatbytesfieldDoer:  doer,
		RestoreResponseBody:   restoreBody,
		scheme:                scheme,
		host:                  host,
		decoder:               dec,
		encoder:               enc,
	}
}

// Concatstrings returns an endpoint that makes HTTP requests to the text
// service concatstrings server.
func (c *Client) Concatstrings() goa.Endpoint {
	var (
		decodeResponse = DecodeConcatstringsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildConcatstringsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ConcatstringsDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("text", "concatstrings", err)
		}
		return decodeResponse(resp)
	}
}

// Concatbytes returns an endpoint that makes HTTP requests to the text service
// concatbytes server.
func (c *Client) Concatbytes() goa.Endpoint {
	var (
		decodeResponse = DecodeConcatbytesResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildConcatbytesRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ConcatbytesDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("text", "concatbytes", err)
		}
		return decodeResponse(resp)
	}
}

// Concatstringfield returns an endpoint that makes HTTP requests to the text
// service concatstringfield server.
func (c *Client) Concatstringfield() goa.Endpoint {
	var (
		decodeResponse = DecodeConcatstringfieldResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildConcatstringfieldRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ConcatstringfieldDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("text", "concatstringfield", err)
		}
		return decodeResponse(resp)
	}
}

// Concatbytesfield returns an endpoint that makes HTTP requests to the text
// service concatbytesfield server.
func (c *Client) Concatbytesfield() goa.Endpoint {
	var (
		decodeResponse = DecodeConcatbytesfieldResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildConcatbytesfieldRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ConcatbytesfieldDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("text", "concatbytesfield", err)
		}
		return decodeResponse(resp)
	}
}
