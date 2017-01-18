// Code generated by clubbygen.
// GENERATED FILE DO NOT EDIT
// +build clubby_strict

package gpio

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"cesanta.com/clubby"
	"cesanta.com/clubby/endpoint"
	"cesanta.com/clubby/frame"
	"cesanta.com/common/go/ourjson"
	"cesanta.com/common/go/ourtrace"
	"github.com/cesanta/errors"
	"golang.org/x/net/trace"

	"github.com/cesanta/ucl"
	"github.com/cesanta/validate-json/schema"
	"github.com/golang/glog"
)

var _ = bytes.MinRead
var _ = fmt.Errorf
var emptyMessage = ourjson.RawMessage{}
var _ = ourtrace.New
var _ = trace.New

const ServiceID = "http://mongoose-iot.com/fwGPIO"

type ReadArgs struct {
	Pin *int64 `json:"pin,omitempty"`
}

type ReadResult struct {
	Value *int64 `json:"value,omitempty"`
}

type RemoveIntHandlerArgs struct {
	Pin *int64 `json:"pin,omitempty"`
}

type SetIntHandlerArgs struct {
	Debounce_ms *int64  `json:"debounce_ms,omitempty"`
	Dst         *string `json:"dst,omitempty"`
	Edge        *string `json:"edge,omitempty"`
	Method      *string `json:"method,omitempty"`
	Pin         *int64  `json:"pin,omitempty"`
	Pull        *string `json:"pull,omitempty"`
}

type SetIntHandlerResult struct {
	Value *int64 `json:"value,omitempty"`
}

type ToggleArgs struct {
	Pin *int64 `json:"pin,omitempty"`
}

type ToggleResult struct {
	Value *int64 `json:"value,omitempty"`
}

type WriteArgs struct {
	Pin   *int64 `json:"pin,omitempty"`
	Value *int64 `json:"value,omitempty"`
}

type Service interface {
	Read(ctx context.Context, args *ReadArgs) (*ReadResult, error)
	RemoveIntHandler(ctx context.Context, args *RemoveIntHandlerArgs) error
	SetIntHandler(ctx context.Context, args *SetIntHandlerArgs) (*SetIntHandlerResult, error)
	Toggle(ctx context.Context, args *ToggleArgs) (*ToggleResult, error)
	Write(ctx context.Context, args *WriteArgs) error
}

type Instance interface {
	Call(context.Context, string, *frame.Command) (*frame.Response, error)
	TraceCall(context.Context, string, *frame.Command) (context.Context, trace.Trace, func(*error))
}

type _validators struct {
	// This comment prevents gofmt from aligning types in the struct.
	ReadArgs *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	ReadResult *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	RemoveIntHandlerArgs *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	SetIntHandlerArgs *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	SetIntHandlerResult *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	ToggleArgs *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	ToggleResult *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	WriteArgs *schema.Validator
}

var (
	validators     *_validators
	validatorsOnce sync.Once
)

func initValidators() {
	validators = &_validators{}

	loader := schema.NewLoader()

	service, err := ucl.Parse(bytes.NewBuffer(_ServiceDefinition))
	if err != nil {
		panic(err)
	}
	// Patch up shortcuts to be proper schemas.
	for _, v := range service.(*ucl.Object).Find("methods").(*ucl.Object).Value {
		if s, ok := v.(*ucl.Object).Find("result").(*ucl.String); ok {
			for kk := range v.(*ucl.Object).Value {
				if kk.Value == "result" {
					v.(*ucl.Object).Value[kk] = &ucl.Object{
						Value: map[ucl.Key]ucl.Value{
							ucl.Key{Value: "type"}: s,
						},
					}
				}
			}
		}
		if v.(*ucl.Object).Find("args") == nil {
			continue
		}
		args := v.(*ucl.Object).Find("args").(*ucl.Object)
		for kk, vv := range args.Value {
			if s, ok := vv.(*ucl.String); ok {
				args.Value[kk] = &ucl.Object{
					Value: map[ucl.Key]ucl.Value{
						ucl.Key{Value: "type"}: s,
					},
				}
			}
		}
	}
	var s *ucl.Object
	_ = s // avoid unused var error
	s = &ucl.Object{
		Value: map[ucl.Key]ucl.Value{
			ucl.Key{Value: "properties"}: service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Read").(*ucl.Object).Find("args"),
			ucl.Key{Value: "type"}:       &ucl.String{Value: "object"},
		},
	}
	if req, found := service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Read").(*ucl.Object).Lookup("required_args"); found {
		s.Value[ucl.Key{Value: "required"}] = req
	}
	validators.ReadArgs, err = schema.NewValidator(s, loader)
	if err != nil {
		panic(err)
	}
	validators.ReadResult, err = schema.NewValidator(service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Read").(*ucl.Object).Find("result"), loader)
	if err != nil {
		panic(err)
	}
	s = &ucl.Object{
		Value: map[ucl.Key]ucl.Value{
			ucl.Key{Value: "properties"}: service.(*ucl.Object).Find("methods").(*ucl.Object).Find("RemoveIntHandler").(*ucl.Object).Find("args"),
			ucl.Key{Value: "type"}:       &ucl.String{Value: "object"},
		},
	}
	if req, found := service.(*ucl.Object).Find("methods").(*ucl.Object).Find("RemoveIntHandler").(*ucl.Object).Lookup("required_args"); found {
		s.Value[ucl.Key{Value: "required"}] = req
	}
	validators.RemoveIntHandlerArgs, err = schema.NewValidator(s, loader)
	if err != nil {
		panic(err)
	}
	s = &ucl.Object{
		Value: map[ucl.Key]ucl.Value{
			ucl.Key{Value: "properties"}: service.(*ucl.Object).Find("methods").(*ucl.Object).Find("SetIntHandler").(*ucl.Object).Find("args"),
			ucl.Key{Value: "type"}:       &ucl.String{Value: "object"},
		},
	}
	if req, found := service.(*ucl.Object).Find("methods").(*ucl.Object).Find("SetIntHandler").(*ucl.Object).Lookup("required_args"); found {
		s.Value[ucl.Key{Value: "required"}] = req
	}
	validators.SetIntHandlerArgs, err = schema.NewValidator(s, loader)
	if err != nil {
		panic(err)
	}
	validators.SetIntHandlerResult, err = schema.NewValidator(service.(*ucl.Object).Find("methods").(*ucl.Object).Find("SetIntHandler").(*ucl.Object).Find("result"), loader)
	if err != nil {
		panic(err)
	}
	s = &ucl.Object{
		Value: map[ucl.Key]ucl.Value{
			ucl.Key{Value: "properties"}: service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Toggle").(*ucl.Object).Find("args"),
			ucl.Key{Value: "type"}:       &ucl.String{Value: "object"},
		},
	}
	if req, found := service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Toggle").(*ucl.Object).Lookup("required_args"); found {
		s.Value[ucl.Key{Value: "required"}] = req
	}
	validators.ToggleArgs, err = schema.NewValidator(s, loader)
	if err != nil {
		panic(err)
	}
	validators.ToggleResult, err = schema.NewValidator(service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Toggle").(*ucl.Object).Find("result"), loader)
	if err != nil {
		panic(err)
	}
	s = &ucl.Object{
		Value: map[ucl.Key]ucl.Value{
			ucl.Key{Value: "properties"}: service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Write").(*ucl.Object).Find("args"),
			ucl.Key{Value: "type"}:       &ucl.String{Value: "object"},
		},
	}
	if req, found := service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Write").(*ucl.Object).Lookup("required_args"); found {
		s.Value[ucl.Key{Value: "required"}] = req
	}
	validators.WriteArgs, err = schema.NewValidator(s, loader)
	if err != nil {
		panic(err)
	}
}

func NewClient(i Instance, addr string) Service {
	validatorsOnce.Do(initValidators)
	return &_Client{i: i, addr: addr}
}

type _Client struct {
	i    Instance
	addr string
}

func (c *_Client) Read(pctx context.Context, args *ReadArgs) (res *ReadResult, err error) {
	cmd := &frame.Command{
		Cmd: "GPIO.Read",
	}
	ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	defer finish(&err)
	_ = tr

	tr.LazyPrintf("args: %s", ourjson.LazyJSON(&args))
	cmd.Args = ourjson.DelayMarshaling(args)
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		v, err := ucl.Parse(bytes.NewReader(b))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.ReadArgs.Validate(v); err != nil {
				glog.Warningf("Sending invalid args for Read: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for Read")
			}
		}
	}
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&endpoint.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	tr.LazyPrintf("res: %s", ourjson.LazyJSON(&resp))

	bb, err := resp.Response.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal result as JSON: %+v", err)
	} else {
		rv, err := ucl.Parse(bytes.NewReader(bb))
		if err == nil {
			if err := validators.ReadResult.Validate(rv); err != nil {
				glog.Warningf("Got invalid result for Read: %+v", err)
				return nil, errors.Annotatef(err, "invalid response for Read")
			}
		}
	}
	var r *ReadResult
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) RemoveIntHandler(pctx context.Context, args *RemoveIntHandlerArgs) (err error) {
	cmd := &frame.Command{
		Cmd: "GPIO.RemoveIntHandler",
	}
	ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	defer finish(&err)
	_ = tr

	tr.LazyPrintf("args: %s", ourjson.LazyJSON(&args))
	cmd.Args = ourjson.DelayMarshaling(args)
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		v, err := ucl.Parse(bytes.NewReader(b))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.RemoveIntHandlerArgs.Validate(v); err != nil {
				glog.Warningf("Sending invalid args for RemoveIntHandler: %+v", err)
				return errors.Annotatef(err, "invalid args for RemoveIntHandler")
			}
		}
	}
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return errors.Trace(err)
	}
	if resp.Status != 0 {
		return errors.Trace(&endpoint.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}
	return nil
}

func (c *_Client) SetIntHandler(pctx context.Context, args *SetIntHandlerArgs) (res *SetIntHandlerResult, err error) {
	cmd := &frame.Command{
		Cmd: "GPIO.SetIntHandler",
	}
	ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	defer finish(&err)
	_ = tr

	tr.LazyPrintf("args: %s", ourjson.LazyJSON(&args))
	cmd.Args = ourjson.DelayMarshaling(args)
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		v, err := ucl.Parse(bytes.NewReader(b))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.SetIntHandlerArgs.Validate(v); err != nil {
				glog.Warningf("Sending invalid args for SetIntHandler: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for SetIntHandler")
			}
		}
	}
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&endpoint.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	tr.LazyPrintf("res: %s", ourjson.LazyJSON(&resp))

	bb, err := resp.Response.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal result as JSON: %+v", err)
	} else {
		rv, err := ucl.Parse(bytes.NewReader(bb))
		if err == nil {
			if err := validators.SetIntHandlerResult.Validate(rv); err != nil {
				glog.Warningf("Got invalid result for SetIntHandler: %+v", err)
				return nil, errors.Annotatef(err, "invalid response for SetIntHandler")
			}
		}
	}
	var r *SetIntHandlerResult
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) Toggle(pctx context.Context, args *ToggleArgs) (res *ToggleResult, err error) {
	cmd := &frame.Command{
		Cmd: "GPIO.Toggle",
	}
	ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	defer finish(&err)
	_ = tr

	tr.LazyPrintf("args: %s", ourjson.LazyJSON(&args))
	cmd.Args = ourjson.DelayMarshaling(args)
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		v, err := ucl.Parse(bytes.NewReader(b))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.ToggleArgs.Validate(v); err != nil {
				glog.Warningf("Sending invalid args for Toggle: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for Toggle")
			}
		}
	}
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&endpoint.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	tr.LazyPrintf("res: %s", ourjson.LazyJSON(&resp))

	bb, err := resp.Response.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal result as JSON: %+v", err)
	} else {
		rv, err := ucl.Parse(bytes.NewReader(bb))
		if err == nil {
			if err := validators.ToggleResult.Validate(rv); err != nil {
				glog.Warningf("Got invalid result for Toggle: %+v", err)
				return nil, errors.Annotatef(err, "invalid response for Toggle")
			}
		}
	}
	var r *ToggleResult
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) Write(pctx context.Context, args *WriteArgs) (err error) {
	cmd := &frame.Command{
		Cmd: "GPIO.Write",
	}
	ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	defer finish(&err)
	_ = tr

	tr.LazyPrintf("args: %s", ourjson.LazyJSON(&args))
	cmd.Args = ourjson.DelayMarshaling(args)
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		v, err := ucl.Parse(bytes.NewReader(b))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.WriteArgs.Validate(v); err != nil {
				glog.Warningf("Sending invalid args for Write: %+v", err)
				return errors.Annotatef(err, "invalid args for Write")
			}
		}
	}
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return errors.Trace(err)
	}
	if resp.Status != 0 {
		return errors.Trace(&endpoint.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}
	return nil
}

func RegisterService(i *clubby.Instance, impl Service) error {
	validatorsOnce.Do(initValidators)
	s := &_Server{impl}
	i.RegisterCommandHandler("GPIO.Read", s.Read)
	i.RegisterCommandHandler("GPIO.RemoveIntHandler", s.RemoveIntHandler)
	i.RegisterCommandHandler("GPIO.SetIntHandler", s.SetIntHandler)
	i.RegisterCommandHandler("GPIO.Toggle", s.Toggle)
	i.RegisterCommandHandler("GPIO.Write", s.Write)
	i.RegisterService(ServiceID, _ServiceDefinition)
	return nil
}

type _Server struct {
	impl Service
}

func (s *_Server) Read(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		if v, err := ucl.Parse(bytes.NewReader(b)); err != nil {
			glog.Errorf("Failed to parse valid JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.ReadArgs.Validate(v); err != nil {
				glog.Warningf("Got invalid args for Read: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for Read")
			}
		}
	}
	var args ReadArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	r, err := s.impl.Read(ctx, &args)
	if err != nil {
		return nil, errors.Trace(err)
	}
	bb, err := json.Marshal(r)
	if err == nil {
		v, err := ucl.Parse(bytes.NewBuffer(bb))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(bb), err)
		} else {
			if err := validators.ReadResult.Validate(v); err != nil {
				glog.Warningf("Returned invalid response for Read: %+v", err)
				return nil, errors.Annotatef(err, "server generated invalid responce for Read")
			}
		}
	}
	return r, nil
}

func (s *_Server) RemoveIntHandler(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		if v, err := ucl.Parse(bytes.NewReader(b)); err != nil {
			glog.Errorf("Failed to parse valid JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.RemoveIntHandlerArgs.Validate(v); err != nil {
				glog.Warningf("Got invalid args for RemoveIntHandler: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for RemoveIntHandler")
			}
		}
	}
	var args RemoveIntHandlerArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return nil, s.impl.RemoveIntHandler(ctx, &args)
}

func (s *_Server) SetIntHandler(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		if v, err := ucl.Parse(bytes.NewReader(b)); err != nil {
			glog.Errorf("Failed to parse valid JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.SetIntHandlerArgs.Validate(v); err != nil {
				glog.Warningf("Got invalid args for SetIntHandler: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for SetIntHandler")
			}
		}
	}
	var args SetIntHandlerArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	r, err := s.impl.SetIntHandler(ctx, &args)
	if err != nil {
		return nil, errors.Trace(err)
	}
	bb, err := json.Marshal(r)
	if err == nil {
		v, err := ucl.Parse(bytes.NewBuffer(bb))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(bb), err)
		} else {
			if err := validators.SetIntHandlerResult.Validate(v); err != nil {
				glog.Warningf("Returned invalid response for SetIntHandler: %+v", err)
				return nil, errors.Annotatef(err, "server generated invalid responce for SetIntHandler")
			}
		}
	}
	return r, nil
}

func (s *_Server) Toggle(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		if v, err := ucl.Parse(bytes.NewReader(b)); err != nil {
			glog.Errorf("Failed to parse valid JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.ToggleArgs.Validate(v); err != nil {
				glog.Warningf("Got invalid args for Toggle: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for Toggle")
			}
		}
	}
	var args ToggleArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	r, err := s.impl.Toggle(ctx, &args)
	if err != nil {
		return nil, errors.Trace(err)
	}
	bb, err := json.Marshal(r)
	if err == nil {
		v, err := ucl.Parse(bytes.NewBuffer(bb))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(bb), err)
		} else {
			if err := validators.ToggleResult.Validate(v); err != nil {
				glog.Warningf("Returned invalid response for Toggle: %+v", err)
				return nil, errors.Annotatef(err, "server generated invalid responce for Toggle")
			}
		}
	}
	return r, nil
}

func (s *_Server) Write(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		if v, err := ucl.Parse(bytes.NewReader(b)); err != nil {
			glog.Errorf("Failed to parse valid JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.WriteArgs.Validate(v); err != nil {
				glog.Warningf("Got invalid args for Write: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for Write")
			}
		}
	}
	var args WriteArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return nil, s.impl.Write(ctx, &args)
}

var _ServiceDefinition = json.RawMessage([]byte(`{
  "methods": {
    "Read": {
      "args": {
        "pin": {
          "doc": "Pin number.",
          "type": "integer"
        }
      },
      "doc": "Read value of a pin. Switches the pin to input mode if needed.",
      "result": {
        "properties": {
          "value": {
            "doc": "Value of the pin, 0 or 1.",
            "type": "integer"
          }
        },
        "type": "object"
      }
    },
    "RemoveIntHandler": {
      "args": {
        "pin": {
          "doc": "Pin number.",
          "type": "integer"
        }
      },
      "doc": "Removes interrupt handler previosuly set."
    },
    "SetIntHandler": {
      "args": {
        "debounce_ms": {
          "doc": "Optional debouncing delay.",
          "type": "integer"
        },
        "dst": {
          "doc": "Destination address for the RPC. Defaults to source of the request.",
          "type": "string"
        },
        "edge": {
          "doc": "Interrupt trigger edge. \"pos\" (0 -\u003e 1), \"neg\" (1 -\u003e 0) or \"any\".",
          "type": "string"
        },
        "method": {
          "doc": "Method for the request. Defaults to GPIO.Int.",
          "type": "string"
        },
        "pin": {
          "doc": "Pin number.",
          "type": "integer"
        },
        "pull": {
          "doc": "Pull setting. \"up\", \"down\" or \"none\" (default).",
          "type": "string"
        }
      },
      "doc": "Configures an interrupt handler on a pin. Switches the pin to input mode if needed. An RPC with the specified method is sent to the specified address and method when an interrupt happens. Request comes with the following arguments: \"pin\" - pin, \"value\" - read before sending, \"ts\" - timestamp. Response to these requests is not expected.\n",
      "result": {
        "properties": {
          "value": {
            "doc": "Value of the pin after toggle, 0 or 1.",
            "type": "integer"
          }
        },
        "type": "object"
      }
    },
    "Toggle": {
      "args": {
        "pin": {
          "doc": "Pin number.",
          "type": "integer"
        }
      },
      "doc": "Toggles pin value. Switches the pin to output mode if needed.",
      "result": {
        "properties": {
          "value": {
            "doc": "Value of the pin after toggle, 0 or 1.",
            "type": "integer"
          }
        },
        "type": "object"
      }
    },
    "Write": {
      "args": {
        "pin": {
          "doc": "Pin number.",
          "type": "integer"
        },
        "value": {
          "doc": "Value to set, 0 or 1.",
          "type": "integer"
        }
      },
      "doc": "Set value of a pin. Switches the pin to output mode if needed."
    }
  },
  "name": "GPIO",
  "namespace": "http://mongoose-iot.com/fw"
}`))
