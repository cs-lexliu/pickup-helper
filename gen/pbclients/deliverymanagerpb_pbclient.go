// Code generated by pbclient-gen. DO NOT EDIT.
// versions: v0.0.0
// source: gen/pb/DeliveryManager/deliverymanagerpb_grpc.pb.go

package pbclients

import (
	"fmt"

	"google.golang.org/grpc"

	"github.com/cs-lexliu/pickup-helper/gen/pb/DeliveryManager"
)

// CloseableDeliveryManagerServiceClient is a closable client for interacting with DeliveryManagerService service.
type CloseableDeliveryManagerServiceClient interface {
	deliverymanagerpb.DeliveryManagerServiceClient
	Close()
}

// closeableDeliveryManagerServiceClient is a closable client for interacting with DeliveryManagerService service.
type closeableDeliveryManagerServiceClient struct {
	deliverymanagerpb.DeliveryManagerServiceClient
	conn *grpc.ClientConn
}

func (c *closeableDeliveryManagerServiceClient) Close() {
	if c == nil || c.conn == nil {
		return
	}
	c.conn.Close()
}

func NewcloseableDeliveryManagerServiceClient(config Config) (*closeableDeliveryManagerServiceClient, error) {
	conn, err := GetGrpcClientConn(config)
	if err != nil {
		return nil, fmt.Errorf("init grpc conn: %w", err)
	}

	return &closeableDeliveryManagerServiceClient{
		DeliveryManagerServiceClient: deliverymanagerpb.NewDeliveryManagerServiceClient(conn),
		conn:                         conn,
	}, nil
}

// CloseableInternalDeliveryManagerServiceClient is a closable client for interacting with InternalDeliveryManagerService service.
type CloseableInternalDeliveryManagerServiceClient interface {
	deliverymanagerpb.InternalDeliveryManagerServiceClient
	Close()
}

// closeableInternalDeliveryManagerServiceClient is a closable client for interacting with InternalDeliveryManagerService service.
type closeableInternalDeliveryManagerServiceClient struct {
	deliverymanagerpb.InternalDeliveryManagerServiceClient
	conn *grpc.ClientConn
}

func (c *closeableInternalDeliveryManagerServiceClient) Close() {
	if c == nil || c.conn == nil {
		return
	}
	c.conn.Close()
}

func NewcloseableInternalDeliveryManagerServiceClient(config Config) (*closeableInternalDeliveryManagerServiceClient, error) {
	conn, err := GetGrpcClientConn(config)
	if err != nil {
		return nil, fmt.Errorf("init grpc conn: %w", err)
	}

	return &closeableInternalDeliveryManagerServiceClient{
		InternalDeliveryManagerServiceClient: deliverymanagerpb.NewInternalDeliveryManagerServiceClient(conn),
		conn:                                 conn,
	}, nil
}

// CloseableDeliveryViewerServiceClient is a closable client for interacting with DeliveryViewerService service.
type CloseableDeliveryViewerServiceClient interface {
	deliverymanagerpb.DeliveryViewerServiceClient
	Close()
}

// closeableDeliveryViewerServiceClient is a closable client for interacting with DeliveryViewerService service.
type closeableDeliveryViewerServiceClient struct {
	deliverymanagerpb.DeliveryViewerServiceClient
	conn *grpc.ClientConn
}

func (c *closeableDeliveryViewerServiceClient) Close() {
	if c == nil || c.conn == nil {
		return
	}
	c.conn.Close()
}

func NewcloseableDeliveryViewerServiceClient(config Config) (*closeableDeliveryViewerServiceClient, error) {
	conn, err := GetGrpcClientConn(config)
	if err != nil {
		return nil, fmt.Errorf("init grpc conn: %w", err)
	}

	return &closeableDeliveryViewerServiceClient{
		DeliveryViewerServiceClient: deliverymanagerpb.NewDeliveryViewerServiceClient(conn),
		conn:                        conn,
	}, nil
}