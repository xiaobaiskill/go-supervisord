package supervisord

import (
	"errors"
	"net/http"

	"github.com/xiaobaiskill/xmlrpc"
)

type (
	Client struct {
		*xmlrpc.Client
	}
)

var (
	// Will be returned by a few functions not yet fully implemented.
	FIXMENotImplementedError error

	// Will be returned if the API endpoint returns false without further explanation.
	ReturnedFalseError error

	// Will be returned if the API endpoint returns a reply we don't know how to parse.
	ReturnedMalformedReply error
)

func init() {
	FIXMENotImplementedError = errors.New("Not implemented yet")
	ReturnedFalseError = errors.New("Call returned false")
	ReturnedMalformedReply = errors.New("Call returned malformed reply")
}

func (c *Client) stringCall(method string, args ...interface{}) (string, error) {
	var str string
	err := c.Call(method, args, &str)

	return str, err
}

func (c *Client) boolCall(method string, args ...interface{}) error {
	var result bool
	err := c.Call(method, args, &result)
	if err != nil {
		return err
	}

	if !result {
		return ReturnedFalseError
	}

	return nil
}

// Get a new client suitable for communicating with a supervisord.
// url must contain a real url to a supervisord RPC-service.
//
// Url for local supervisord should be http://127.0.0.1:9001/RPC2 by default.
func NewClient(url string, httpclient *http.Client) (*Client, error) {
	rpc, err := xmlrpc.NewClient(url, httpclient)
	if err != nil {
		return nil, err
	}

	return &Client{rpc}, nil
}

// NewUnixSocketClient returns a new client which connects to supervisord
// though a local unix socket
//func NewUnixSocketClient(path string) (*Client, error) {
//
//	// we inject this fake dialer, it will only connect
//	// to the path given, and does not care about what address
//	// is given to it.
//	dialer := func(_, _ string) (conn net.Conn, err error) {
//		return net.Dial("unix", path)
//	}
//
//	tr := &http.Transport{
//		Dial: dialer,
//	}
//
//	// we pass a valid url, as this is later url.Parse()'ed
//	// also we need to somehow specify "/RPC2"
//	rpc, err := xmlrpc.NewClient("http://127.0.0.1/RPC2", tr)
//	if err != nil {
//		return nil, err
//	}
//
//	return &Client{rpc}, nil
//
//}
