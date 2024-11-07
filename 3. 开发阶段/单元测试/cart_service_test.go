package service_test

import (
    "testing"
    "shop-app/3. 开发阶段/代码编写/cart/service"
    "shop-app/3. 开发阶段/代码编写/cart/repository"
    "github.com/stretchr/testify/assert"
)

type MockCartRepository struct {
    repository CartRepository
}

func (m *MockCartRepository) AddToCart(ctx context.Context, userID int, productID int, quantity int) error {
    // 模拟添加商品到购物车
    return nil
}

func (m *MockCartRepository) GetCart(ctx context.Context, userID int) ([]*service.Cart规范, error) {
    // 模拟获取购物车列表
    return []*service.Cart规范{
        {
            UserID: 1,
            ProductID: 1,
            Quantity: 1,
        },
    }, nil
}

func (m *MockCartRepository) UpdateCart(ctx context.Context, userID int, productID int, quantity int) error {
    // 模拟更新购物车商品数量
    return nil
}

func (m *MockCartRepository) RemoveFromCart(ctx context.Context, userID int, productID int) error {
    // 模拟从购物车中移除商品
    return nil
}

func (m *MockCartRepository) ClearCart(ctx context.Context, userID int) error {
    // 模拟清空购物车
    return nil
}

func TestCartService_AddToCart(t *testing.T) {
    mockRepo := &MockCartRepository{}
    cartService := service.NewCartService(mockRepo)

    err := cartService.AddToCart(context.Background(), 1, 1, 1)
    assert.NoError(t, err)
}

func TestCartService_GetCart(t *testing.T) {
    mockRepo := &MockCartRepository{}
    cartService := service.NewCartService(mockRepo)

    carts, err := cartService.GetCart(context.Background(), 1)
    assert.NoError(t, err)
    assert.Len(t, carts, 1)
    assert.Equal(t, carts[0].UserID, int64(1))
    assert.Equal(t, carts[0].ProductID, int64(1))
    assert.Equal(t, carts[0].Quantity, int64(1))
}

func TestCartService_UpdateCart(t *testing.T) {
    mockRepo := &MockCartRepository{}
    cartService := service.NewCartService(mockRepo)

    err := cartService.UpdateCart(context.Background(), 1, 1, 2)
    assert.NoError(t, err)
}

func TestCartService_RemoveFromCart(t *testing.T) {
    mockRepo := &MockCartRepository{}
    cartService := service.NewCartService(mockRepo)

    err := cartService.RemoveFromCart(context.Background(), 1, 1)
    assert.NoError(t, err)
}

func TestCartService_ClearCart(t *testing.T) {
    mockRepo := &MockCartRepository{}
    cartService := service.NewCartService(mockRepo)

    err := cartService.ClearCart(context.Background(), 1)
    assert.NoError(t, err)
}