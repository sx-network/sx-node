// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DataFeedOperatorClient is the client API for DataFeedOperator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataFeedOperatorClient interface {
	// AddProposal adds a datafeed proposal
	ReportOutcome(ctx context.Context, in *ReportOutcomeReq, opts ...grpc.CallOption) (*ReportOutcomeResp, error)
}

type dataFeedOperatorClient struct {
	cc grpc.ClientConnInterface
}

func NewDataFeedOperatorClient(cc grpc.ClientConnInterface) DataFeedOperatorClient {
	return &dataFeedOperatorClient{cc}
}

func (c *dataFeedOperatorClient) ReportOutcome(ctx context.Context, in *ReportOutcomeReq, opts ...grpc.CallOption) (*ReportOutcomeResp, error) {
	out := new(ReportOutcomeResp)
	err := c.cc.Invoke(ctx, "/v1.DataFeedOperator/ReportOutcome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataFeedOperatorServer is the server API for DataFeedOperator service.
// All implementations must embed UnimplementedDataFeedOperatorServer
// for forward compatibility
type DataFeedOperatorServer interface {
	// AddProposal adds a datafeed proposal
	ReportOutcome(context.Context, *ReportOutcomeReq) (*ReportOutcomeResp, error)
	mustEmbedUnimplementedDataFeedOperatorServer()
}

// UnimplementedDataFeedOperatorServer must be embedded to have forward compatible implementations.
type UnimplementedDataFeedOperatorServer struct {
}

func (UnimplementedDataFeedOperatorServer) ReportOutcome(context.Context, *ReportOutcomeReq) (*ReportOutcomeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportOutcome not implemented")
}
func (UnimplementedDataFeedOperatorServer) mustEmbedUnimplementedDataFeedOperatorServer() {}

// UnsafeDataFeedOperatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataFeedOperatorServer will
// result in compilation errors.
type UnsafeDataFeedOperatorServer interface {
	mustEmbedUnimplementedDataFeedOperatorServer()
}

func RegisterDataFeedOperatorServer(s grpc.ServiceRegistrar, srv DataFeedOperatorServer) {
	s.RegisterService(&DataFeedOperator_ServiceDesc, srv)
}

func _DataFeedOperator_ReportOutcome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportOutcomeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataFeedOperatorServer).ReportOutcome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DataFeedOperator/ReportOutcome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataFeedOperatorServer).ReportOutcome(ctx, req.(*ReportOutcomeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DataFeedOperator_ServiceDesc is the grpc.ServiceDesc for DataFeedOperator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataFeedOperator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.DataFeedOperator",
	HandlerType: (*DataFeedOperatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportOutcome",
			Handler:    _DataFeedOperator_ReportOutcome_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datafeed/proto/operator.proto",
}