package js

import (
	"syscall/js"

	"github.com/louisbuchbinder/core/lib/util"
	"github.com/louisbuchbinder/core/wasm"
)

func NewJsResponse(r *Response) js.Value {
	opts := js.Global().Get("Object").New()
	opts.Set("status", r.Status)
	return js.Global().Get("Response").New(r.Body, opts)
}

type Response struct {
	Body   string
	Status int
	// Headers ...
}

type jsRequest struct {
	method string
	url    string
}

var _ wasm.JsRequest = new(jsRequest)

func (r *jsRequest) Method() string {
	return r.method
}

func (r *jsRequest) Url() wasm.JsUrl {
	return newJsUrl(r.url)
}

type jsUrl struct {
	val js.Value
}

var _ wasm.JsUrl = new(jsUrl)

func newJsUrl(url string) wasm.JsUrl {
	return &jsUrl{js.Global().Get("URL").New(url)}
}

func (u *jsUrl) Hash() string     { return u.val.Get("hash").String() }
func (u *jsUrl) Host() string     { return u.val.Get("host").String() }
func (u *jsUrl) Hostname() string { return u.val.Get("hostname").String() }
func (u *jsUrl) Href() string     { return u.val.Get("href").String() }
func (u *jsUrl) Origin() string   { return u.val.Get("origin").String() }
func (u *jsUrl) Password() string { return u.val.Get("password").String() }
func (u *jsUrl) Pathname() string { return u.val.Get("pathname").String() }
func (u *jsUrl) Port() string     { return u.val.Get("port").String() }
func (u *jsUrl) Protocol() string { return u.val.Get("protocol").String() }
func (u *jsUrl) Search() string   { return u.val.Get("search").String() }
func (u *jsUrl) Username() string { return u.val.Get("username").String() }

func (u *jsUrl) GetSearchParam(k string) *string {
	// u.val.Get("searchParams").
	v := u.val.Get("searchParams").Call("get", k)
	if v.IsNull() {
		return nil
	}
	return util.Ptr(v.String())
}
