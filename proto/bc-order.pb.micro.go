// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/bc-order.proto

package bcorder

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for BcOrder service

func NewBcOrderEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for BcOrder service

type BcOrderService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (BcOrder_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (BcOrder_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (BcOrder_BidiStreamService, error)
}

type bcOrderService struct {
	c    client.Client
	name string
}

func NewBcOrderService(name string, c client.Client) BcOrderService {
	return &bcOrderService{
		c:    c,
		name: name,
	}
}

func (c *bcOrderService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "BcOrder.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcOrderService) ClientStream(ctx context.Context, opts ...client.CallOption) (BcOrder_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "BcOrder.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &bcOrderServiceClientStream{stream}, nil
}

type BcOrder_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*ClientStreamRequest) error
}

type bcOrderServiceClientStream struct {
	stream client.Stream
}

func (x *bcOrderServiceClientStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *bcOrderServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *bcOrderServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bcOrderServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bcOrderServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bcOrderServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *bcOrderService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (BcOrder_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "BcOrder.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &bcOrderServiceServerStream{stream}, nil
}

type BcOrder_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type bcOrderServiceServerStream struct {
	stream client.Stream
}

func (x *bcOrderServiceServerStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *bcOrderServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *bcOrderServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bcOrderServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bcOrderServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bcOrderServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bcOrderService) BidiStream(ctx context.Context, opts ...client.CallOption) (BcOrder_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "BcOrder.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &bcOrderServiceBidiStream{stream}, nil
}

type BcOrder_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type bcOrderServiceBidiStream struct {
	stream client.Stream
}

func (x *bcOrderServiceBidiStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *bcOrderServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *bcOrderServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bcOrderServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bcOrderServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bcOrderServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *bcOrderServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for BcOrder service

type BcOrderHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, BcOrder_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, BcOrder_ServerStreamStream) error
	BidiStream(context.Context, BcOrder_BidiStreamStream) error
}

func RegisterBcOrderHandler(s server.Server, hdlr BcOrderHandler, opts ...server.HandlerOption) error {
	type bcOrder interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type BcOrder struct {
		bcOrder
	}
	h := &bcOrderHandler{hdlr}
	return s.Handle(s.NewHandler(&BcOrder{h}, opts...))
}

type bcOrderHandler struct {
	BcOrderHandler
}

func (h *bcOrderHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.BcOrderHandler.Call(ctx, in, out)
}

func (h *bcOrderHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.BcOrderHandler.ClientStream(ctx, &bcOrderClientStreamStream{stream})
}

type BcOrder_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type bcOrderClientStreamStream struct {
	stream server.Stream
}

func (x *bcOrderClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *bcOrderClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bcOrderClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bcOrderClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bcOrderClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *bcOrderHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.BcOrderHandler.ServerStream(ctx, m, &bcOrderServerStreamStream{stream})
}

type BcOrder_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type bcOrderServerStreamStream struct {
	stream server.Stream
}

func (x *bcOrderServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *bcOrderServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bcOrderServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bcOrderServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bcOrderServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *bcOrderHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.BcOrderHandler.BidiStream(ctx, &bcOrderBidiStreamStream{stream})
}

type BcOrder_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type bcOrderBidiStreamStream struct {
	stream server.Stream
}

func (x *bcOrderBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *bcOrderBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bcOrderBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bcOrderBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bcOrderBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *bcOrderBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
