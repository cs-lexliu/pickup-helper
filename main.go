package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	ordermanagerpb "github.com/cs-lexliu/pickup-helper/gen/pb/ordermanager"
	shoppingcheckoutcorepb "github.com/cs-lexliu/pickup-helper/gen/pb/shopping-checkout-core"
)

type Order = ordermanagerpb.FulfillmentOrder
type Address = ordermanagerpb.Address
type Checkout = shoppingcheckoutcorepb.Checkout
type BuyerInfo = shoppingcheckoutcorepb.Order_ReceiverInfo

const (
	orderManagerAddress     = "localhost:10001"
	shoppingCheckoutAddress = "localhost:10002"
)

func main() {
	if err := mainWithErr(); err != nil {
		log.Fatalf(fmt.Sprintf("main: %v", err))
	}
}

func mainWithErr() error {
	h, err := newHelper()
	if err != nil {
		log.Fatalf(fmt.Sprintf("new helper: %v", err))
	}
	defer h.closeHelper()

	nextAvailableTimeSlot := getNextAvailableTimeSlot()

	// TODO: find a way to get the orderID
	orderID := ""

	ctx := context.Background()
	if err := h.requestPickup(ctx, orderID, nextAvailableTimeSlot); err != nil {
		return fmt.Errorf("request pickup: %w", err)
	}
	return nil
}

type timeSlot struct {
	date string
	slot string
}

type helper struct {
	omConn    *grpc.ClientConn
	omClient  ordermanagerpb.OrderManagerClient
	sccConn   *grpc.ClientConn
	sccClient shoppingcheckoutcorepb.ShoppingCheckoutCoreServiceClient
}

func newHelper() (*helper, error) {
	omConn, err := grpc.Dial(orderManagerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("dail order manager: %w", err)
	}
	omClient := ordermanagerpb.NewOrderManagerClient(omConn)

	scConn, err := grpc.Dial(shoppingCheckoutAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("dail shopping checkout: %w", err)
	}
	scClient := shoppingcheckoutcorepb.NewShoppingCheckoutCoreServiceClient(scConn)

	return &helper{
		omConn:    omConn,
		omClient:  omClient,
		sccConn:   scConn,
		sccClient: scClient,
	}, nil
}

func (h *helper) closeHelper() {
	if h.omConn != nil {
		h.omConn.Close()
	}
	if h.sccConn != nil {
		h.sccConn.Close()
	}
}

func (h *helper) requestPickup(ctx context.Context, orderID string, ts *timeSlot) error {
	var (
		// fill in the Refrash pickup address as const
		pickupAddress = &Address{
			Country:    "",
			State:      "",
			City:       "",
			PostalCode: "",
			UnitNo:     "",
			Address1:   "",
			Address2:   "",
		}
	)

	order, err := h.getOrder(ctx, orderID)
	if err != nil {
		return fmt.Errorf("get order: %w", err)
	}

	checkoutID := order.GetCheckoutId()
	checkout, err := h.getCheckout(ctx, checkoutID)
	if err != nil {
		return fmt.Errorf("get checkout: %w", err)
	}
	buyerInfo := retrieveBuyerInfo(checkout)

	if err := h.startDelivery(ctx, orderID, buyerInfo, pickupAddress, ts); err != nil {
		return fmt.Errorf("start delivery: %w", err)
	}
	return nil
}

func (h *helper) getOrder(ctx context.Context, orderID string) (*Order, error) {
	req := &ordermanagerpb.GetOrdersRequest{
		Query: &ordermanagerpb.GetOrdersRequest_Query{
			ByIds: []string{orderID},
		},
	}
	resp, err := h.omClient.GetOrders(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("om get order %s: %w", orderID, err)
	}
	if len(resp.GetOrders()) == 0 {
		return nil, fmt.Errorf("no order found '%s'", orderID)
	}
	// assume the order is always existed
	return resp.GetOrders()[0], nil
}

func (h *helper) getCheckout(ctx context.Context, checkoutID string) (*Checkout, error) {
	req := &shoppingcheckoutcorepb.GetCheckoutByIDRequest{
		CheckoutId: checkoutID,
	}
	resp, err := h.sccClient.GetCheckoutByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("scc get checkout: %w", err)
	}
	return resp.GetCheckout(), nil
}

func (h *helper) startDelivery(ctx context.Context, orderID string, buyerInfo *BuyerInfo, pickupAddress *Address, timeslot *timeSlot) error {
	req := &ordermanagerpb.StartD2DDeliveryRequest{
		OrderId:    orderID,
		BuyerName:  buyerInfo.GetReceiverName(),
		BuyerPhone: buyerInfo.GetReceiverPhone(),
		BuyerEmail: buyerInfo.GetReceiverEmail(),
		Address:    pickupAddress,
		PickupDate: timeslot.date,
		PickupSlot: timeslot.slot,
	}
	if _, err := h.omClient.StartD2DDelivery(ctx, req); err != nil {
		return fmt.Errorf("om start d2d delivery: %w", err)
	}
	return nil
}

func retrieveBuyerInfo(checkout *shoppingcheckoutcorepb.Checkout) *BuyerInfo {
	if checkout == nil {
		return nil
	}
	return checkout.GetOrders()[0].GetDeliveryInfo().GetReceiverInfo()
}

func getNextAvailableTimeSlot() *timeSlot {
	now := time.Now()
	nextWorkingDay := getNextWorkingDay(now).Truncate(24 * time.Hour)
	date := strconv.Itoa(int(nextWorkingDay.Unix()))
	slot := getSlot(nextWorkingDay)
	return &timeSlot{
		date: date,
		slot: slot,
	}
}

func getNextWorkingDay(currentTime time.Time) time.Time {
	// add 1 day until the weekday is not Sunday
	for {
		currentTime = currentTime.AddDate(0, 0, 1)
		if currentTime.Weekday() != time.Sunday {
			return currentTime
		}
	}
}

func getSlot(currentTime time.Time) string {
	switch currentTime.Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		return "9AM-6PM"
	case time.Saturday:
		return "9AM-12PM"
	}
	return ""
}
