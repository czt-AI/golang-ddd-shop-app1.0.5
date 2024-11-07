package service_test

import (
    "testing"
    "shop-app/3. 开发阶段/代码编写/payment/service"
    "shop-app/3. 开发阶段/代码编写/payment/repository"
    "github.com/stretchr/testify/assert"
)

type MockPaymentRepository struct {
    repository PaymentRepository
}

func (m *MockPaymentRepository) InitiatePayment(ctx context.Context, orderID int, paymentMethod string) (*service.Payment规范, error) {
    // 模拟初始化支付
    return &service.Payment规范{
        PaymentID: 1,
        OrderID: orderID,
        PaymentMethod: paymentMethod,
        Status: "pending",
    }, nil
}

func (m *MockPaymentRepository) ExecutePayment(ctx context.Context, paymentID int) (*service.Payment规范, error) {
    // 模拟执行支付
    return &service.Payment规范{
        PaymentID: paymentID,
        OrderID: 1,
        PaymentMethod: "credit_card",
        Status: "completed",
    }, nil
}

func (m *MockPaymentRepository) QueryPayment(ctx context.Context, paymentID int) (*service.Payment规范, error) {
    // 模拟查询支付状态
    return &service.Payment规范{
        PaymentID: paymentID,
        OrderID: 1,
        PaymentMethod: "credit_card",
        Status: "completed",
    }, nil
}

func TestPaymentService_InitiatePayment(t *testing.T) {
    mockRepo := &MockPaymentRepository{}
    paymentService := service.NewPaymentService(mockRepo)

    payment, err := paymentService.InitiatePayment(context.Background(), 1, "credit_card")
    assert.NoError(t, err)
    assert.NotNil(t, payment)
    assert.Equal(t, payment.PaymentID, int64(1))
    assert.Equal(t, payment.OrderID, int64(1))
    assert.Equal(t, payment.PaymentMethod, "credit_card")
    assert.Equal(t, payment.Status, "pending")
}

func TestPaymentService_ExecutePayment(t *testing.T) {
    mockRepo := &MockPaymentRepository{}
    paymentService := service.NewPaymentService(mockRepo)

    payment, err := paymentService.ExecutePayment(context.Background(), 1)
    assert.NoError(t, err)
    assert.NotNil(t, payment)
    assert.Equal(t, payment.PaymentID, int64(1))
    assert.Equal(t, payment.OrderID, int64(1))
    assert.Equal(t, payment.PaymentMethod, "credit_card")
    assert.Equal(t, payment.Status, "completed")
}

func TestPaymentService_QueryPayment(t *testing.T) {
    mockRepo := &MockPaymentRepository{}
    paymentService := service.NewPaymentService(mockRepo)

    payment, err := paymentService.QueryPayment(context.Background(), 1)
    assert.NoError(t, err)
    assert.NotNil(t, payment)
    assert.Equal(t, payment.PaymentID, int64(1))
    assert.Equal(t, payment.OrderID, int64(1))
    assert.Equal(t, payment.PaymentMethod, "credit_card")
    assert.Equal(t, payment.Status, "completed")
}