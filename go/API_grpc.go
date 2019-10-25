//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: api

package proto

import (
  context "context"
  grpc "google.golang.org/grpc"
)

// Client API for API service
type APIClient interface{
  CreateStream(ctx context.Context, in *CreateStreamRequest, 
  	opts... grpc.CallOption) (* CreateStreamResponse, error)  
  Subscribe(ctx context.Context, in *SubscribeRequest, 
  	opts... grpc.CallOption) (API_SubscribeClient, error)  
  FetchMetadata(ctx context.Context, in *FetchMetadataRequest, 
  	opts... grpc.CallOption) (* FetchMetadataResponse, error)  
  Publish(ctx context.Context, in *PublishRequest, 
  	opts... grpc.CallOption) (* PublishResponse, error)  
}

type aPIClient struct {
  cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
  return &aPIClient{cc}
}

func (c *aPIClient) CreateStream(ctx context.Context, in *CreateStreamRequest, 
	opts... grpc.CallOption) (* CreateStreamResponse, error) {
  out := new(CreateStreamResponse)
  err := grpc.Invoke(ctx, "/proto.API/CreateStream", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

func (c *aPIClient) Subscribe(ctx context.Context, in *SubscribeRequest, 
	opts... grpc.CallOption) (API_SubscribeClient, error) {
  stream, err := grpc.NewClientStream(ctx, &_API_serviceDesc.Streams[0], c.cc, "/proto.API/Subscribe", opts...)
  if err != nil { return nil, err }
  x := &aPISubscribeClient{stream}
  if err := x.ClientStream.SendMsg(in); err != nil { return nil, err }
  if err := x.ClientStream.CloseSend(); err != nil { return nil, err }
  return x,nil
}

type API_SubscribeClient interface {
  Recv() (*Message, error)
  grpc.ClientStream
}

type aPISubscribeClient struct{
  grpc.ClientStream
}

func (x *aPISubscribeClient) Recv() (*Message, error) {
  m := new(Message)
  if err := x.ClientStream.RecvMsg(m); err != nil { return nil, err }
  return m, nil
}

func (c *aPIClient) FetchMetadata(ctx context.Context, in *FetchMetadataRequest, 
	opts... grpc.CallOption) (* FetchMetadataResponse, error) {
  out := new(FetchMetadataResponse)
  err := grpc.Invoke(ctx, "/proto.API/FetchMetadata", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

func (c *aPIClient) Publish(ctx context.Context, in *PublishRequest, 
	opts... grpc.CallOption) (* PublishResponse, error) {
  out := new(PublishResponse)
  err := grpc.Invoke(ctx, "/proto.API/Publish", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

// Server API for API service
type APIServer interface {
  CreateStream(context.Context, *CreateStreamRequest) (*CreateStreamResponse, error)  
  Subscribe(*SubscribeRequest, API_SubscribeServer) error  
  FetchMetadata(context.Context, *FetchMetadataRequest) (*FetchMetadataResponse, error)  
  Publish(context.Context, *PublishRequest) (*PublishResponse, error)  
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
  s.RegisterService(&_API_serviceDesc, srv)
}

func _API_CreateStream_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(CreateStreamRequest)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(APIServer).CreateStream(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/proto.API/CreateStream",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(APIServer).CreateStream(ctx, req.(* CreateStreamRequest))
  }
  return interceptor(ctx, in, info, handler)
}


func _API_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
  m := new(SubscribeRequest)
  if err := stream.RecvMsg(m); err != nil { return err }
  return srv.(APIServer).Subscribe(m, &aPISubscribeServer{stream})
}

type API_SubscribeServer interface { 
  Send(* Message) error
  grpc.ServerStream
}

type aPISubscribeServer struct {
  grpc.ServerStream
}

func (x *aPISubscribeServer) Send(m *Message) error {
  return x.ServerStream.SendMsg(m)
}


func _API_FetchMetadata_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(FetchMetadataRequest)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(APIServer).FetchMetadata(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/proto.API/FetchMetadata",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(APIServer).FetchMetadata(ctx, req.(* FetchMetadataRequest))
  }
  return interceptor(ctx, in, info, handler)
}


func _API_Publish_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(PublishRequest)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(APIServer).Publish(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/proto.API/Publish",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(APIServer).Publish(ctx, req.(* PublishRequest))
  }
  return interceptor(ctx, in, info, handler)
}


var _API_serviceDesc = grpc.ServiceDesc{
  ServiceName: "proto.API",
  HandlerType: (*APIServer)(nil),
  Methods: []grpc.MethodDesc{
    {
      MethodName: "CreateStream",
      Handler: _API_CreateStream_Handler, 
    },
    {
      MethodName: "FetchMetadata",
      Handler: _API_FetchMetadata_Handler, 
    },
    {
      MethodName: "Publish",
      Handler: _API_Publish_Handler, 
    },
  },
  Streams: []grpc.StreamDesc{
    {
      StreamName: "Subscribe",
      Handler: _API_Subscribe_Handler, 
      ServerStreams: true,
    },
  },
}

