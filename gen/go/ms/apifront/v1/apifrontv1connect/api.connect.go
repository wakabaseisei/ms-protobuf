// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: ms/apifront/v1/api.proto

package apifrontv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/wakabaseisei/ms-protobuf/gen/go/ms/apifront/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// GreetServiceName is the fully-qualified name of the GreetService service.
	GreetServiceName = "ms.apifront.v1.GreetService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GreetServicePingProcedure is the fully-qualified name of the GreetService's Ping RPC.
	GreetServicePingProcedure = "/ms.apifront.v1.GreetService/Ping"
	// GreetServiceGreetProcedure is the fully-qualified name of the GreetService's Greet RPC.
	GreetServiceGreetProcedure = "/ms.apifront.v1.GreetService/Greet"
)

// GreetServiceClient is a client for the ms.apifront.v1.GreetService service.
type GreetServiceClient interface {
	Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error)
	Greet(context.Context, *connect_go.Request[v1.GreetRequest]) (*connect_go.Response[v1.GreetResponse], error)
}

// NewGreetServiceClient constructs a client for the ms.apifront.v1.GreetService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGreetServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) GreetServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &greetServiceClient{
		ping: connect_go.NewClient[v1.PingRequest, v1.PingResponse](
			httpClient,
			baseURL+GreetServicePingProcedure,
			opts...,
		),
		greet: connect_go.NewClient[v1.GreetRequest, v1.GreetResponse](
			httpClient,
			baseURL+GreetServiceGreetProcedure,
			opts...,
		),
	}
}

// greetServiceClient implements GreetServiceClient.
type greetServiceClient struct {
	ping  *connect_go.Client[v1.PingRequest, v1.PingResponse]
	greet *connect_go.Client[v1.GreetRequest, v1.GreetResponse]
}

// Ping calls ms.apifront.v1.GreetService.Ping.
func (c *greetServiceClient) Ping(ctx context.Context, req *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// Greet calls ms.apifront.v1.GreetService.Greet.
func (c *greetServiceClient) Greet(ctx context.Context, req *connect_go.Request[v1.GreetRequest]) (*connect_go.Response[v1.GreetResponse], error) {
	return c.greet.CallUnary(ctx, req)
}

// GreetServiceHandler is an implementation of the ms.apifront.v1.GreetService service.
type GreetServiceHandler interface {
	Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error)
	Greet(context.Context, *connect_go.Request[v1.GreetRequest]) (*connect_go.Response[v1.GreetResponse], error)
}

// NewGreetServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGreetServiceHandler(svc GreetServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	greetServicePingHandler := connect_go.NewUnaryHandler(
		GreetServicePingProcedure,
		svc.Ping,
		opts...,
	)
	greetServiceGreetHandler := connect_go.NewUnaryHandler(
		GreetServiceGreetProcedure,
		svc.Greet,
		opts...,
	)
	return "/ms.apifront.v1.GreetService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GreetServicePingProcedure:
			greetServicePingHandler.ServeHTTP(w, r)
		case GreetServiceGreetProcedure:
			greetServiceGreetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGreetServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGreetServiceHandler struct{}

func (UnimplementedGreetServiceHandler) Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ms.apifront.v1.GreetService.Ping is not implemented"))
}

func (UnimplementedGreetServiceHandler) Greet(context.Context, *connect_go.Request[v1.GreetRequest]) (*connect_go.Response[v1.GreetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ms.apifront.v1.GreetService.Greet is not implemented"))
}
