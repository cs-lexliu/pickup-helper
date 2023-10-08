// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	messaging_proto "github.com/carousell/messaging/messaging/messaging_proto"

	mock "github.com/stretchr/testify/mock"

	ordermanagerpb "github.com/cs-lexliu/pickup-helper/gen/pb/OrderManager"
)

// CloseableOrderManagerClient is an autogenerated mock type for the CloseableOrderManagerClient type
type CloseableOrderManagerClient struct {
	mock.Mock
}

// CancelOrder provides a mock function with given fields: ctx, in, opts
func (_m *CloseableOrderManagerClient) CancelOrder(ctx context.Context, in *ordermanagerpb.CancelOrderRequest, opts ...grpc.CallOption) (*ordermanagerpb.CancelOrderResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ordermanagerpb.CancelOrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.CancelOrderRequest, ...grpc.CallOption) (*ordermanagerpb.CancelOrderResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.CancelOrderRequest, ...grpc.CallOption) *ordermanagerpb.CancelOrderResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ordermanagerpb.CancelOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ordermanagerpb.CancelOrderRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *CloseableOrderManagerClient) Close() {
	_m.Called()
}

// CreateOrder provides a mock function with given fields: ctx, in, opts
func (_m *CloseableOrderManagerClient) CreateOrder(ctx context.Context, in *ordermanagerpb.CreateOrderRequest, opts ...grpc.CallOption) (*ordermanagerpb.CreateOrderResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ordermanagerpb.CreateOrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.CreateOrderRequest, ...grpc.CallOption) (*ordermanagerpb.CreateOrderResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.CreateOrderRequest, ...grpc.CallOption) *ordermanagerpb.CreateOrderResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ordermanagerpb.CreateOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ordermanagerpb.CreateOrderRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FinishOrder provides a mock function with given fields: ctx, in, opts
func (_m *CloseableOrderManagerClient) FinishOrder(ctx context.Context, in *ordermanagerpb.FinishOrderRequest, opts ...grpc.CallOption) (*ordermanagerpb.FinishOrderResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ordermanagerpb.FinishOrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.FinishOrderRequest, ...grpc.CallOption) (*ordermanagerpb.FinishOrderResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.FinishOrderRequest, ...grpc.CallOption) *ordermanagerpb.FinishOrderResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ordermanagerpb.FinishOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ordermanagerpb.FinishOrderRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrders provides a mock function with given fields: ctx, in, opts
func (_m *CloseableOrderManagerClient) GetOrders(ctx context.Context, in *ordermanagerpb.GetOrdersRequest, opts ...grpc.CallOption) (*ordermanagerpb.GetOrdersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ordermanagerpb.GetOrdersResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.GetOrdersRequest, ...grpc.CallOption) (*ordermanagerpb.GetOrdersResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.GetOrdersRequest, ...grpc.CallOption) *ordermanagerpb.GetOrdersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ordermanagerpb.GetOrdersResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ordermanagerpb.GetOrdersRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotifyDisputeRequestClosed provides a mock function with given fields: ctx, in, opts
func (_m *CloseableOrderManagerClient) NotifyDisputeRequestClosed(ctx context.Context, in *ordermanagerpb.NotifyDisputeRequestClosedRequest, opts ...grpc.CallOption) (*ordermanagerpb.NotifyDisputeRequestClosedResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ordermanagerpb.NotifyDisputeRequestClosedResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.NotifyDisputeRequestClosedRequest, ...grpc.CallOption) (*ordermanagerpb.NotifyDisputeRequestClosedResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.NotifyDisputeRequestClosedRequest, ...grpc.CallOption) *ordermanagerpb.NotifyDisputeRequestClosedResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ordermanagerpb.NotifyDisputeRequestClosedResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ordermanagerpb.NotifyDisputeRequestClosedRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotifyDisputeRequestCreated provides a mock function with given fields: ctx, in, opts
func (_m *CloseableOrderManagerClient) NotifyDisputeRequestCreated(ctx context.Context, in *ordermanagerpb.NotifyDisputeRequestCreatedRequest, opts ...grpc.CallOption) (*ordermanagerpb.NotifyDisputeRequestCreatedResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ordermanagerpb.NotifyDisputeRequestCreatedResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.NotifyDisputeRequestCreatedRequest, ...grpc.CallOption) (*ordermanagerpb.NotifyDisputeRequestCreatedResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.NotifyDisputeRequestCreatedRequest, ...grpc.CallOption) *ordermanagerpb.NotifyDisputeRequestCreatedResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ordermanagerpb.NotifyDisputeRequestCreatedResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ordermanagerpb.NotifyDisputeRequestCreatedRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReceiveMessage provides a mock function with given fields: ctx, in, opts
func (_m *CloseableOrderManagerClient) ReceiveMessage(ctx context.Context, in *messaging_proto.MessageRequest, opts ...grpc.CallOption) (*messaging_proto.MessageResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *messaging_proto.MessageResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *messaging_proto.MessageRequest, ...grpc.CallOption) (*messaging_proto.MessageResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *messaging_proto.MessageRequest, ...grpc.CallOption) *messaging_proto.MessageResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*messaging_proto.MessageResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *messaging_proto.MessageRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScheduleMeetup provides a mock function with given fields: ctx, in, opts
func (_m *CloseableOrderManagerClient) ScheduleMeetup(ctx context.Context, in *ordermanagerpb.ScheduleMeetupRequest, opts ...grpc.CallOption) (*ordermanagerpb.ScheduleMeetupResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ordermanagerpb.ScheduleMeetupResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.ScheduleMeetupRequest, ...grpc.CallOption) (*ordermanagerpb.ScheduleMeetupResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.ScheduleMeetupRequest, ...grpc.CallOption) *ordermanagerpb.ScheduleMeetupResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ordermanagerpb.ScheduleMeetupResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ordermanagerpb.ScheduleMeetupRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartD2DDelivery provides a mock function with given fields: ctx, in, opts
func (_m *CloseableOrderManagerClient) StartD2DDelivery(ctx context.Context, in *ordermanagerpb.StartD2DDeliveryRequest, opts ...grpc.CallOption) (*ordermanagerpb.StartD2DDeliveryResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ordermanagerpb.StartD2DDeliveryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.StartD2DDeliveryRequest, ...grpc.CallOption) (*ordermanagerpb.StartD2DDeliveryResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ordermanagerpb.StartD2DDeliveryRequest, ...grpc.CallOption) *ordermanagerpb.StartD2DDeliveryResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ordermanagerpb.StartD2DDeliveryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ordermanagerpb.StartD2DDeliveryRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCloseableOrderManagerClient creates a new instance of CloseableOrderManagerClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCloseableOrderManagerClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *CloseableOrderManagerClient {
	mock := &CloseableOrderManagerClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
