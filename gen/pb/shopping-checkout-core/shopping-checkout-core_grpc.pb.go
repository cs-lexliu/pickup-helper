// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: shopping-checkout-core/shopping-checkout-core.proto

package shoppingcheckoutcorepb

import (
	context "context"
	messaging_proto "github.com/carousell/messaging/messaging/messaging_proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ShoppingCheckoutCoreService_CancelCheckout_FullMethodName               = "/shoppingcheckoutcorepb.ShoppingCheckoutCoreService/CancelCheckout"
	ShoppingCheckoutCoreService_CompleteCheckout_FullMethodName             = "/shoppingcheckoutcorepb.ShoppingCheckoutCoreService/CompleteCheckout"
	ShoppingCheckoutCoreService_GetCheckoutByID_FullMethodName              = "/shoppingcheckoutcorepb.ShoppingCheckoutCoreService/GetCheckoutByID"
	ShoppingCheckoutCoreService_GetCheckoutList_FullMethodName              = "/shoppingcheckoutcorepb.ShoppingCheckoutCoreService/GetCheckoutList"
	ShoppingCheckoutCoreService_GetCheckoutListByBuyerID_FullMethodName     = "/shoppingcheckoutcorepb.ShoppingCheckoutCoreService/GetCheckoutListByBuyerID"
	ShoppingCheckoutCoreService_InitCheckout_FullMethodName                 = "/shoppingcheckoutcorepb.ShoppingCheckoutCoreService/InitCheckout"
	ShoppingCheckoutCoreService_InvalidateCheckout_FullMethodName           = "/shoppingcheckoutcorepb.ShoppingCheckoutCoreService/InvalidateCheckout"
	ShoppingCheckoutCoreService_ReceiveMessage_FullMethodName               = "/shoppingcheckoutcorepb.ShoppingCheckoutCoreService/ReceiveMessage"
	ShoppingCheckoutCoreService_UpdateCheckoutPendingPayment_FullMethodName = "/shoppingcheckoutcorepb.ShoppingCheckoutCoreService/UpdateCheckoutPendingPayment"
	ShoppingCheckoutCoreService_MarkListingSoftReserved_FullMethodName      = "/shoppingcheckoutcorepb.ShoppingCheckoutCoreService/MarkListingSoftReserved"
	ShoppingCheckoutCoreService_UnmarkListingSoftReserved_FullMethodName    = "/shoppingcheckoutcorepb.ShoppingCheckoutCoreService/UnmarkListingSoftReserved"
)

// ShoppingCheckoutCoreServiceClient is the client API for ShoppingCheckoutCoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShoppingCheckoutCoreServiceClient interface {
	CancelCheckout(ctx context.Context, in *CancelCheckoutRequest, opts ...grpc.CallOption) (*CancelCheckoutResponse, error)
	CompleteCheckout(ctx context.Context, in *CompleteCheckoutRequest, opts ...grpc.CallOption) (*CompleteCheckoutResponse, error)
	GetCheckoutByID(ctx context.Context, in *GetCheckoutByIDRequest, opts ...grpc.CallOption) (*GetCheckoutByIDResponse, error)
	GetCheckoutList(ctx context.Context, in *GetCheckoutListRequest, opts ...grpc.CallOption) (*GetCheckoutListResponse, error)
	GetCheckoutListByBuyerID(ctx context.Context, in *GetCheckoutListByBuyerIDRequest, opts ...grpc.CallOption) (*GetCheckoutListByBuyerIDResponse, error)
	InitCheckout(ctx context.Context, in *InitCheckoutRequest, opts ...grpc.CallOption) (*InitCheckoutResponse, error)
	InvalidateCheckout(ctx context.Context, in *InvalidateCheckoutRequest, opts ...grpc.CallOption) (*InvalidateCheckoutResponse, error)
	ReceiveMessage(ctx context.Context, in *messaging_proto.MessageRequest, opts ...grpc.CallOption) (*messaging_proto.MessageResponse, error)
	UpdateCheckoutPendingPayment(ctx context.Context, in *UpdateCheckoutPendingPaymentRequest, opts ...grpc.CallOption) (*UpdateCheckoutPendingPaymentResponse, error)
	MarkListingSoftReserved(ctx context.Context, in *MarkListingSoftReservedRequest, opts ...grpc.CallOption) (*MarkListingSoftReservedResponse, error)
	UnmarkListingSoftReserved(ctx context.Context, in *UnmarkListingSoftReservedRequest, opts ...grpc.CallOption) (*UnmarkListingSoftReservedResponse, error)
}

type shoppingCheckoutCoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShoppingCheckoutCoreServiceClient(cc grpc.ClientConnInterface) ShoppingCheckoutCoreServiceClient {
	return &shoppingCheckoutCoreServiceClient{cc}
}

func (c *shoppingCheckoutCoreServiceClient) CancelCheckout(ctx context.Context, in *CancelCheckoutRequest, opts ...grpc.CallOption) (*CancelCheckoutResponse, error) {
	out := new(CancelCheckoutResponse)
	err := c.cc.Invoke(ctx, ShoppingCheckoutCoreService_CancelCheckout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCheckoutCoreServiceClient) CompleteCheckout(ctx context.Context, in *CompleteCheckoutRequest, opts ...grpc.CallOption) (*CompleteCheckoutResponse, error) {
	out := new(CompleteCheckoutResponse)
	err := c.cc.Invoke(ctx, ShoppingCheckoutCoreService_CompleteCheckout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCheckoutCoreServiceClient) GetCheckoutByID(ctx context.Context, in *GetCheckoutByIDRequest, opts ...grpc.CallOption) (*GetCheckoutByIDResponse, error) {
	out := new(GetCheckoutByIDResponse)
	err := c.cc.Invoke(ctx, ShoppingCheckoutCoreService_GetCheckoutByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCheckoutCoreServiceClient) GetCheckoutList(ctx context.Context, in *GetCheckoutListRequest, opts ...grpc.CallOption) (*GetCheckoutListResponse, error) {
	out := new(GetCheckoutListResponse)
	err := c.cc.Invoke(ctx, ShoppingCheckoutCoreService_GetCheckoutList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCheckoutCoreServiceClient) GetCheckoutListByBuyerID(ctx context.Context, in *GetCheckoutListByBuyerIDRequest, opts ...grpc.CallOption) (*GetCheckoutListByBuyerIDResponse, error) {
	out := new(GetCheckoutListByBuyerIDResponse)
	err := c.cc.Invoke(ctx, ShoppingCheckoutCoreService_GetCheckoutListByBuyerID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCheckoutCoreServiceClient) InitCheckout(ctx context.Context, in *InitCheckoutRequest, opts ...grpc.CallOption) (*InitCheckoutResponse, error) {
	out := new(InitCheckoutResponse)
	err := c.cc.Invoke(ctx, ShoppingCheckoutCoreService_InitCheckout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCheckoutCoreServiceClient) InvalidateCheckout(ctx context.Context, in *InvalidateCheckoutRequest, opts ...grpc.CallOption) (*InvalidateCheckoutResponse, error) {
	out := new(InvalidateCheckoutResponse)
	err := c.cc.Invoke(ctx, ShoppingCheckoutCoreService_InvalidateCheckout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCheckoutCoreServiceClient) ReceiveMessage(ctx context.Context, in *messaging_proto.MessageRequest, opts ...grpc.CallOption) (*messaging_proto.MessageResponse, error) {
	out := new(messaging_proto.MessageResponse)
	err := c.cc.Invoke(ctx, ShoppingCheckoutCoreService_ReceiveMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCheckoutCoreServiceClient) UpdateCheckoutPendingPayment(ctx context.Context, in *UpdateCheckoutPendingPaymentRequest, opts ...grpc.CallOption) (*UpdateCheckoutPendingPaymentResponse, error) {
	out := new(UpdateCheckoutPendingPaymentResponse)
	err := c.cc.Invoke(ctx, ShoppingCheckoutCoreService_UpdateCheckoutPendingPayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCheckoutCoreServiceClient) MarkListingSoftReserved(ctx context.Context, in *MarkListingSoftReservedRequest, opts ...grpc.CallOption) (*MarkListingSoftReservedResponse, error) {
	out := new(MarkListingSoftReservedResponse)
	err := c.cc.Invoke(ctx, ShoppingCheckoutCoreService_MarkListingSoftReserved_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCheckoutCoreServiceClient) UnmarkListingSoftReserved(ctx context.Context, in *UnmarkListingSoftReservedRequest, opts ...grpc.CallOption) (*UnmarkListingSoftReservedResponse, error) {
	out := new(UnmarkListingSoftReservedResponse)
	err := c.cc.Invoke(ctx, ShoppingCheckoutCoreService_UnmarkListingSoftReserved_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShoppingCheckoutCoreServiceServer is the server API for ShoppingCheckoutCoreService service.
// All implementations must embed UnimplementedShoppingCheckoutCoreServiceServer
// for forward compatibility
type ShoppingCheckoutCoreServiceServer interface {
	CancelCheckout(context.Context, *CancelCheckoutRequest) (*CancelCheckoutResponse, error)
	CompleteCheckout(context.Context, *CompleteCheckoutRequest) (*CompleteCheckoutResponse, error)
	GetCheckoutByID(context.Context, *GetCheckoutByIDRequest) (*GetCheckoutByIDResponse, error)
	GetCheckoutList(context.Context, *GetCheckoutListRequest) (*GetCheckoutListResponse, error)
	GetCheckoutListByBuyerID(context.Context, *GetCheckoutListByBuyerIDRequest) (*GetCheckoutListByBuyerIDResponse, error)
	InitCheckout(context.Context, *InitCheckoutRequest) (*InitCheckoutResponse, error)
	InvalidateCheckout(context.Context, *InvalidateCheckoutRequest) (*InvalidateCheckoutResponse, error)
	ReceiveMessage(context.Context, *messaging_proto.MessageRequest) (*messaging_proto.MessageResponse, error)
	UpdateCheckoutPendingPayment(context.Context, *UpdateCheckoutPendingPaymentRequest) (*UpdateCheckoutPendingPaymentResponse, error)
	MarkListingSoftReserved(context.Context, *MarkListingSoftReservedRequest) (*MarkListingSoftReservedResponse, error)
	UnmarkListingSoftReserved(context.Context, *UnmarkListingSoftReservedRequest) (*UnmarkListingSoftReservedResponse, error)
	mustEmbedUnimplementedShoppingCheckoutCoreServiceServer()
}

// UnimplementedShoppingCheckoutCoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShoppingCheckoutCoreServiceServer struct {
}

func (UnimplementedShoppingCheckoutCoreServiceServer) CancelCheckout(context.Context, *CancelCheckoutRequest) (*CancelCheckoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelCheckout not implemented")
}
func (UnimplementedShoppingCheckoutCoreServiceServer) CompleteCheckout(context.Context, *CompleteCheckoutRequest) (*CompleteCheckoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteCheckout not implemented")
}
func (UnimplementedShoppingCheckoutCoreServiceServer) GetCheckoutByID(context.Context, *GetCheckoutByIDRequest) (*GetCheckoutByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckoutByID not implemented")
}
func (UnimplementedShoppingCheckoutCoreServiceServer) GetCheckoutList(context.Context, *GetCheckoutListRequest) (*GetCheckoutListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckoutList not implemented")
}
func (UnimplementedShoppingCheckoutCoreServiceServer) GetCheckoutListByBuyerID(context.Context, *GetCheckoutListByBuyerIDRequest) (*GetCheckoutListByBuyerIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckoutListByBuyerID not implemented")
}
func (UnimplementedShoppingCheckoutCoreServiceServer) InitCheckout(context.Context, *InitCheckoutRequest) (*InitCheckoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitCheckout not implemented")
}
func (UnimplementedShoppingCheckoutCoreServiceServer) InvalidateCheckout(context.Context, *InvalidateCheckoutRequest) (*InvalidateCheckoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateCheckout not implemented")
}
func (UnimplementedShoppingCheckoutCoreServiceServer) ReceiveMessage(context.Context, *messaging_proto.MessageRequest) (*messaging_proto.MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveMessage not implemented")
}
func (UnimplementedShoppingCheckoutCoreServiceServer) UpdateCheckoutPendingPayment(context.Context, *UpdateCheckoutPendingPaymentRequest) (*UpdateCheckoutPendingPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCheckoutPendingPayment not implemented")
}
func (UnimplementedShoppingCheckoutCoreServiceServer) MarkListingSoftReserved(context.Context, *MarkListingSoftReservedRequest) (*MarkListingSoftReservedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkListingSoftReserved not implemented")
}
func (UnimplementedShoppingCheckoutCoreServiceServer) UnmarkListingSoftReserved(context.Context, *UnmarkListingSoftReservedRequest) (*UnmarkListingSoftReservedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnmarkListingSoftReserved not implemented")
}
func (UnimplementedShoppingCheckoutCoreServiceServer) mustEmbedUnimplementedShoppingCheckoutCoreServiceServer() {
}

// UnsafeShoppingCheckoutCoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShoppingCheckoutCoreServiceServer will
// result in compilation errors.
type UnsafeShoppingCheckoutCoreServiceServer interface {
	mustEmbedUnimplementedShoppingCheckoutCoreServiceServer()
}

func RegisterShoppingCheckoutCoreServiceServer(s grpc.ServiceRegistrar, srv ShoppingCheckoutCoreServiceServer) {
	s.RegisterService(&ShoppingCheckoutCoreService_ServiceDesc, srv)
}

func _ShoppingCheckoutCoreService_CancelCheckout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelCheckoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCheckoutCoreServiceServer).CancelCheckout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShoppingCheckoutCoreService_CancelCheckout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCheckoutCoreServiceServer).CancelCheckout(ctx, req.(*CancelCheckoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCheckoutCoreService_CompleteCheckout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteCheckoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCheckoutCoreServiceServer).CompleteCheckout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShoppingCheckoutCoreService_CompleteCheckout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCheckoutCoreServiceServer).CompleteCheckout(ctx, req.(*CompleteCheckoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCheckoutCoreService_GetCheckoutByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCheckoutByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCheckoutCoreServiceServer).GetCheckoutByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShoppingCheckoutCoreService_GetCheckoutByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCheckoutCoreServiceServer).GetCheckoutByID(ctx, req.(*GetCheckoutByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCheckoutCoreService_GetCheckoutList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCheckoutListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCheckoutCoreServiceServer).GetCheckoutList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShoppingCheckoutCoreService_GetCheckoutList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCheckoutCoreServiceServer).GetCheckoutList(ctx, req.(*GetCheckoutListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCheckoutCoreService_GetCheckoutListByBuyerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCheckoutListByBuyerIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCheckoutCoreServiceServer).GetCheckoutListByBuyerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShoppingCheckoutCoreService_GetCheckoutListByBuyerID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCheckoutCoreServiceServer).GetCheckoutListByBuyerID(ctx, req.(*GetCheckoutListByBuyerIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCheckoutCoreService_InitCheckout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitCheckoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCheckoutCoreServiceServer).InitCheckout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShoppingCheckoutCoreService_InitCheckout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCheckoutCoreServiceServer).InitCheckout(ctx, req.(*InitCheckoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCheckoutCoreService_InvalidateCheckout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvalidateCheckoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCheckoutCoreServiceServer).InvalidateCheckout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShoppingCheckoutCoreService_InvalidateCheckout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCheckoutCoreServiceServer).InvalidateCheckout(ctx, req.(*InvalidateCheckoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCheckoutCoreService_ReceiveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(messaging_proto.MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCheckoutCoreServiceServer).ReceiveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShoppingCheckoutCoreService_ReceiveMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCheckoutCoreServiceServer).ReceiveMessage(ctx, req.(*messaging_proto.MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCheckoutCoreService_UpdateCheckoutPendingPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCheckoutPendingPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCheckoutCoreServiceServer).UpdateCheckoutPendingPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShoppingCheckoutCoreService_UpdateCheckoutPendingPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCheckoutCoreServiceServer).UpdateCheckoutPendingPayment(ctx, req.(*UpdateCheckoutPendingPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCheckoutCoreService_MarkListingSoftReserved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkListingSoftReservedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCheckoutCoreServiceServer).MarkListingSoftReserved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShoppingCheckoutCoreService_MarkListingSoftReserved_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCheckoutCoreServiceServer).MarkListingSoftReserved(ctx, req.(*MarkListingSoftReservedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCheckoutCoreService_UnmarkListingSoftReserved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnmarkListingSoftReservedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCheckoutCoreServiceServer).UnmarkListingSoftReserved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShoppingCheckoutCoreService_UnmarkListingSoftReserved_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCheckoutCoreServiceServer).UnmarkListingSoftReserved(ctx, req.(*UnmarkListingSoftReservedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShoppingCheckoutCoreService_ServiceDesc is the grpc.ServiceDesc for ShoppingCheckoutCoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShoppingCheckoutCoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shoppingcheckoutcorepb.ShoppingCheckoutCoreService",
	HandlerType: (*ShoppingCheckoutCoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CancelCheckout",
			Handler:    _ShoppingCheckoutCoreService_CancelCheckout_Handler,
		},
		{
			MethodName: "CompleteCheckout",
			Handler:    _ShoppingCheckoutCoreService_CompleteCheckout_Handler,
		},
		{
			MethodName: "GetCheckoutByID",
			Handler:    _ShoppingCheckoutCoreService_GetCheckoutByID_Handler,
		},
		{
			MethodName: "GetCheckoutList",
			Handler:    _ShoppingCheckoutCoreService_GetCheckoutList_Handler,
		},
		{
			MethodName: "GetCheckoutListByBuyerID",
			Handler:    _ShoppingCheckoutCoreService_GetCheckoutListByBuyerID_Handler,
		},
		{
			MethodName: "InitCheckout",
			Handler:    _ShoppingCheckoutCoreService_InitCheckout_Handler,
		},
		{
			MethodName: "InvalidateCheckout",
			Handler:    _ShoppingCheckoutCoreService_InvalidateCheckout_Handler,
		},
		{
			MethodName: "ReceiveMessage",
			Handler:    _ShoppingCheckoutCoreService_ReceiveMessage_Handler,
		},
		{
			MethodName: "UpdateCheckoutPendingPayment",
			Handler:    _ShoppingCheckoutCoreService_UpdateCheckoutPendingPayment_Handler,
		},
		{
			MethodName: "MarkListingSoftReserved",
			Handler:    _ShoppingCheckoutCoreService_MarkListingSoftReserved_Handler,
		},
		{
			MethodName: "UnmarkListingSoftReserved",
			Handler:    _ShoppingCheckoutCoreService_UnmarkListingSoftReserved_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shopping-checkout-core/shopping-checkout-core.proto",
}
