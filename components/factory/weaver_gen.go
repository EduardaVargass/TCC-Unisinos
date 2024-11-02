// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package factory

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "TCC-UNISINOS/components/factory/FactoryComponent",
		Iface: reflect.TypeOf((*FactoryComponent)(nil)).Elem(),
		Impl:  reflect.TypeOf(factoryComponent{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return factoryComponent_local_stub{impl: impl.(FactoryComponent), tracer: tracer, calculateOEEPerMachineMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "TCC-UNISINOS/components/factory/FactoryComponent", Method: "CalculateOEEPerMachine", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return factoryComponent_client_stub{stub: stub, calculateOEEPerMachineMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "TCC-UNISINOS/components/factory/FactoryComponent", Method: "CalculateOEEPerMachine", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return factoryComponent_server_stub{impl: impl.(FactoryComponent), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return factoryComponent_reflect_stub{caller: caller}
		},
		RefData: "⟦5228c0a0:wEaVeReDgE:TCC-UNISINOS/components/factory/FactoryComponent→TCC-UNISINOS/components/machine/MachineComponent⟧\n",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[FactoryComponent] = (*factoryComponent)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*factoryComponent)(nil)

// Local stub implementations.

type factoryComponent_local_stub struct {
	impl                          FactoryComponent
	tracer                        trace.Tracer
	calculateOEEPerMachineMetrics *codegen.MethodMetrics
}

// Check that factoryComponent_local_stub implements the FactoryComponent interface.
var _ FactoryComponent = (*factoryComponent_local_stub)(nil)

func (s factoryComponent_local_stub) CalculateOEEPerMachine(ctx context.Context) (r0 []OEEPerMachine, err error) {
	// Update metrics.
	begin := s.calculateOEEPerMachineMetrics.Begin()
	defer func() { s.calculateOEEPerMachineMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "factory.FactoryComponent.CalculateOEEPerMachine", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.CalculateOEEPerMachine(ctx)
}

// Client stub implementations.

type factoryComponent_client_stub struct {
	stub                          codegen.Stub
	calculateOEEPerMachineMetrics *codegen.MethodMetrics
}

// Check that factoryComponent_client_stub implements the FactoryComponent interface.
var _ FactoryComponent = (*factoryComponent_client_stub)(nil)

func (s factoryComponent_client_stub) CalculateOEEPerMachine(ctx context.Context) (r0 []OEEPerMachine, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.calculateOEEPerMachineMetrics.Begin()
	defer func() { s.calculateOEEPerMachineMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "factory.FactoryComponent.CalculateOEEPerMachine", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_OEEPerMachine_cde04979(dec)
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][24]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.24.6 (codegen
version v0.24.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type factoryComponent_server_stub struct {
	impl    FactoryComponent
	addLoad func(key uint64, load float64)
}

// Check that factoryComponent_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*factoryComponent_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s factoryComponent_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "CalculateOEEPerMachine":
		return s.calculateOEEPerMachine
	default:
		return nil
	}
}

func (s factoryComponent_server_stub) calculateOEEPerMachine(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.CalculateOEEPerMachine(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_OEEPerMachine_cde04979(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type factoryComponent_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that factoryComponent_reflect_stub implements the FactoryComponent interface.
var _ FactoryComponent = (*factoryComponent_reflect_stub)(nil)

func (s factoryComponent_reflect_stub) CalculateOEEPerMachine(ctx context.Context) (r0 []OEEPerMachine, err error) {
	err = s.caller("CalculateOEEPerMachine", ctx, []any{}, []any{&r0})
	return
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*OEEPerMachine)(nil)

type __is_OEEPerMachine[T ~struct {
	MachineName string  "json:\"machineName\""
	OEE         float64 "json:\"oee\""
	weaver.AutoMarshal
}] struct{}

var _ __is_OEEPerMachine[OEEPerMachine]

func (x *OEEPerMachine) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("OEEPerMachine.WeaverMarshal: nil receiver"))
	}
	enc.String(x.MachineName)
	enc.Float64(x.OEE)
}

func (x *OEEPerMachine) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("OEEPerMachine.WeaverUnmarshal: nil receiver"))
	}
	x.MachineName = dec.String()
	x.OEE = dec.Float64()
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_OEEPerMachine_cde04979(enc *codegen.Encoder, arg []OEEPerMachine) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_OEEPerMachine_cde04979(dec *codegen.Decoder) []OEEPerMachine {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]OEEPerMachine, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}