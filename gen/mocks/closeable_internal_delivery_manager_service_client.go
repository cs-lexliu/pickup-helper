// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	deliverymanagerpb "github.com/cs-lexliu/pickup-helper/gen/pb/DeliveryManager"
	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// CloseableInternalDeliveryManagerServiceClient is an autogenerated mock type for the CloseableInternalDeliveryManagerServiceClient type
type CloseableInternalDeliveryManagerServiceClient struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *CloseableInternalDeliveryManagerServiceClient) Close() {
	_m.Called()
}

// CreateCategoryConfig provides a mock function with given fields: ctx, in, opts
func (_m *CloseableInternalDeliveryManagerServiceClient) CreateCategoryConfig(ctx context.Context, in *deliverymanagerpb.CreateCategoryConfigRequest, opts ...grpc.CallOption) (*deliverymanagerpb.CreateCategoryConfigResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *deliverymanagerpb.CreateCategoryConfigResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *deliverymanagerpb.CreateCategoryConfigRequest, ...grpc.CallOption) (*deliverymanagerpb.CreateCategoryConfigResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *deliverymanagerpb.CreateCategoryConfigRequest, ...grpc.CallOption) *deliverymanagerpb.CreateCategoryConfigResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deliverymanagerpb.CreateCategoryConfigResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *deliverymanagerpb.CreateCategoryConfigRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoryConfigByID provides a mock function with given fields: ctx, in, opts
func (_m *CloseableInternalDeliveryManagerServiceClient) GetCategoryConfigByID(ctx context.Context, in *deliverymanagerpb.GetCategoryConfigRequest, opts ...grpc.CallOption) (*deliverymanagerpb.GetCategoryConfigResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *deliverymanagerpb.GetCategoryConfigResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *deliverymanagerpb.GetCategoryConfigRequest, ...grpc.CallOption) (*deliverymanagerpb.GetCategoryConfigResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *deliverymanagerpb.GetCategoryConfigRequest, ...grpc.CallOption) *deliverymanagerpb.GetCategoryConfigResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deliverymanagerpb.GetCategoryConfigResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *deliverymanagerpb.GetCategoryConfigRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SyncDeliveryToLogistics provides a mock function with given fields: ctx, in, opts
func (_m *CloseableInternalDeliveryManagerServiceClient) SyncDeliveryToLogistics(ctx context.Context, in *deliverymanagerpb.SyncDeliveryToLogisticsRequest, opts ...grpc.CallOption) (*deliverymanagerpb.SyncDeliveryToLogisticsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *deliverymanagerpb.SyncDeliveryToLogisticsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *deliverymanagerpb.SyncDeliveryToLogisticsRequest, ...grpc.CallOption) (*deliverymanagerpb.SyncDeliveryToLogisticsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *deliverymanagerpb.SyncDeliveryToLogisticsRequest, ...grpc.CallOption) *deliverymanagerpb.SyncDeliveryToLogisticsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deliverymanagerpb.SyncDeliveryToLogisticsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *deliverymanagerpb.SyncDeliveryToLogisticsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCloseableInternalDeliveryManagerServiceClient creates a new instance of CloseableInternalDeliveryManagerServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCloseableInternalDeliveryManagerServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *CloseableInternalDeliveryManagerServiceClient {
	mock := &CloseableInternalDeliveryManagerServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
