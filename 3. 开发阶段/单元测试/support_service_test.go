package service_test

import (
    "testing"
    "shop-app/3. 开发阶段/代码编写/support/service"
    "shop-app/3. 开发阶段/代码编写/support/repository"
    "github.com/stretchr/testify/assert"
)

type MockSupportRepository struct {
    repository SupportRepository
}

func (m *MockSupportRepository) CreateTicket(ctx context.Context, userID int, message string) (*service.Ticket, error) {
    // 模拟创建工单
    return &service.Ticket{
        TicketID: 1,
        UserID: userID,
        Message: message,
        Status: "open",
    }, nil
}

func (m *MockSupportRepository) GetTicket(ctx context.Context, ticketID int) (*service.Ticket, error) {
    // 模拟获取工单
    return &service.Ticket{
        TicketID: ticketID,
        UserID: 1,
        Message: "Support request",
        Status: "open",
    }, nil
}

func (m *MockSupportRepository) UpdateTicketStatus(ctx context.Context, ticketID int, status service.TicketStatus) error {
    // 模拟更新工单状态
    return nil
}

func (m *MockSupportRepository) GetTicketsByUser(ctx context.Context, userID int) ([]*service.Ticket, error) {
    // 模拟获取用户所有工单
    return []*service.Ticket{
        {
            TicketID: 1,
            UserID: userID,
            Message: "Support request",
            Status: "open",
        },
    }, nil
}

func TestSupportService_CreateTicket(t *testing.T) {
    mockRepo := &MockSupportRepository{}
    supportService := service.NewSupportService(mockRepo)

    ticket, err := supportService.CreateTicket(context.Background(), 1, "I need help!")
    assert.NoError(t, err)
    assert.NotNil(t, ticket)
    assert.Equal(t, ticket.TicketID, int64(1))
    assert.Equal(t, ticket.UserID, int64(1))
    assert.Equal(t, ticket.Message, "I need help!")
    assert.Equal(t, ticket.Status, "open")
}

func TestSupportService_GetTicket(t *testing.T) {
    mockRepo := &MockSupportRepository{}
    supportService := service.NewSupportService(mockRepo)

    ticket, err := supportService.GetTicket(context.Background(), 1)
    assert.NoError(t, err)
    assert.NotNil(t, ticket)
    assert.Equal(t, ticket.TicketID, int64(1))
    assert.Equal(t, ticket.UserID, int64(1))
    assert.Equal(t, ticket.Message, "Support request")
    assert.Equal(t, ticket.Status, "open")
}

func TestSupportService_UpdateTicketStatus(t *testing.T) {
    mockRepo := &MockSupportRepository{}
    supportService := service.NewSupportService(mockRepo)

    err := supportService.UpdateTicketStatus(context.Background(), 1, "closed")
    assert.NoError(t, err)
}

func TestSupportService_GetTicketsByUser(t *testing.T) {
    mockRepo := &MockSupportRepository{}
    supportService := service.NewSupportService(mockRepo)

    tickets, err := supportService.GetTicketsByUser(context.Background(), 1)
    assert.NoError(t, err)
    assert.Len(t, tickets, 1)
    assert.Equal(t, tickets[0].TicketID, int64(1))
    assert.Equal(t, tickets[0].UserID, int64(1))
    assert.Equal(t, tickets[0].Message, "Support request")
    assert.Equal(t, tickets[0].Status, "open")
}