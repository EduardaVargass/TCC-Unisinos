// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package oee

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
		Name:  "TCC-UNISINOS/components/oee/OEEComponent",
		Iface: reflect.TypeOf((*OEEComponent)(nil)).Elem(),
		Impl:  reflect.TypeOf(oeeComponent{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return oEEComponent_local_stub{impl: impl.(OEEComponent), tracer: tracer, calculateOEEMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "TCC-UNISINOS/components/oee/OEEComponent", Method: "CalculateOEE", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return oEEComponent_client_stub{stub: stub, calculateOEEMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "TCC-UNISINOS/components/oee/OEEComponent", Method: "CalculateOEE", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return oEEComponent_server_stub{impl: impl.(OEEComponent), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return oEEComponent_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[OEEComponent] = (*oeeComponent)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*oeeComponent)(nil)

// Local stub implementations.

type oEEComponent_local_stub struct {
	impl                OEEComponent
	tracer              trace.Tracer
	calculateOEEMetrics *codegen.MethodMetrics
}

// Check that oEEComponent_local_stub implements the OEEComponent interface.
var _ OEEComponent = (*oEEComponent_local_stub)(nil)

func (s oEEComponent_local_stub) CalculateOEE(ctx context.Context, a0 ProductionData) (r0 float64, err error) {
	// Update metrics.
	begin := s.calculateOEEMetrics.Begin()
	defer func() { s.calculateOEEMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "oee.OEEComponent.CalculateOEE", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.CalculateOEE(ctx, a0)
}

// Client stub implementations.

type oEEComponent_client_stub struct {
	stub                codegen.Stub
	calculateOEEMetrics *codegen.MethodMetrics
}

// Check that oEEComponent_client_stub implements the OEEComponent interface.
var _ OEEComponent = (*oEEComponent_client_stub)(nil)

func (s oEEComponent_client_stub) CalculateOEE(ctx context.Context, a0 ProductionData) (r0 float64, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.calculateOEEMetrics.Begin()
	defer func() { s.calculateOEEMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "oee.OEEComponent.CalculateOEE", trace.WithSpanKind(trace.SpanKindClient))
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

	// Preallocate a buffer of the right size.
	size := 0
	size += serviceweaver_size_ProductionData_5f5172ba(&a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	(a0).WeaverMarshal(enc)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Float64()
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

type oEEComponent_server_stub struct {
	impl    OEEComponent
	addLoad func(key uint64, load float64)
}

// Check that oEEComponent_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*oEEComponent_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s oEEComponent_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "CalculateOEE":
		return s.calculateOEE
	default:
		return nil
	}
}

func (s oEEComponent_server_stub) calculateOEE(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 ProductionData
	(&a0).WeaverUnmarshal(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.CalculateOEE(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Float64(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type oEEComponent_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that oEEComponent_reflect_stub implements the OEEComponent interface.
var _ OEEComponent = (*oEEComponent_reflect_stub)(nil)

func (s oEEComponent_reflect_stub) CalculateOEE(ctx context.Context, a0 ProductionData) (r0 float64, err error) {
	err = s.caller("CalculateOEE", ctx, []any{a0}, []any{&r0})
	return
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*ProductionData)(nil)

type __is_ProductionData[T ~struct {
	ProductionTime float64 "json:\"productionTime\""
	AvailableTime  float64 "json:\"availableTime\""
	IdealProdCount int     "json:\"idealProdCount\""
	RejCount       int     "json:\"rejCount\""
	GoodCount      int     "json:\"goodCount\""
	weaver.AutoMarshal
}] struct{}

var _ __is_ProductionData[ProductionData]

func (x *ProductionData) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("ProductionData.WeaverMarshal: nil receiver"))
	}
	enc.Float64(x.ProductionTime)
	enc.Float64(x.AvailableTime)
	enc.Int(x.IdealProdCount)
	enc.Int(x.RejCount)
	enc.Int(x.GoodCount)
}

func (x *ProductionData) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("ProductionData.WeaverUnmarshal: nil receiver"))
	}
	x.ProductionTime = dec.Float64()
	x.AvailableTime = dec.Float64()
	x.IdealProdCount = dec.Int()
	x.RejCount = dec.Int()
	x.GoodCount = dec.Int()
}

// Size implementations.

// serviceweaver_size_ProductionData_5f5172ba returns the size (in bytes) of the serialization
// of the provided type.
func serviceweaver_size_ProductionData_5f5172ba(x *ProductionData) int {
	size := 0
	size += 8
	size += 8
	size += 8
	size += 8
	size += 8
	size += 0
	return size
}
