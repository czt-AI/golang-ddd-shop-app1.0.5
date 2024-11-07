package service_test

import (
    "testing"
    "shop-app/3. 开发阶段/代码编写/product/service"
    "shop-app/3. 开发阶段/代码编写/product/repository"
    "github.com/stretchr/testify/assert"
)

type MockProductRepository struct {
    repository ProductRepository
}

func (m *MockProductRepository) GetProductByID(ctx context.Context, productID int) (*service.Product规范, error) {
    // 模拟获取商品ID
    return &service.Product规范{
        ProductID: 1,
        Name: "Test Product",
        Description: "This is a test product",
        Price: 10.99,
        Stock: 100,
        CategoryID: 1,
    }, nil
}

func (m *MockProductRepository) ListProducts(ctx context.Context, page, limit int) ([]*service.Product规范, error) {
    // 模拟获取商品列表
    return []*service.Product规范{
        {
            ProductID: 1,
            Name: "Test Product",
            Description: "This is a test product",
            Price: 10.99,
            Stock: 100,
            CategoryID: 1,
        },
    }, nil
}

func (m *MockProductRepository) CreateProduct(ctx context.Context, product *service.Product规范) error {
    // 模拟创建商品
    return nil
}

func (m *MockProductRepository) UpdateProduct(ctx context.Context, productID int, product *service.Product规范) error {
    // 模拟更新商品
    return nil
}

func (m *MockProductRepository) DeleteProduct(ctx context.Context, productID int) error {
    // 模拟删除商品
    return nil
}

func TestProductService_GetProductByID(t *testing.T) {
    mockRepo := &MockProductRepository{}
    productService := service.New ProductService(mockRepo)

    product, err := productService.GetProductByID(context.Background(), 1)
    assert.NoError(t, err)
    assert.NotNil(t, product)
    assert.Equal(t, product.ProductID, int64(1))
    assert.Equal(t, product.Name, "Test Product")
    assert.Equal(t, product.Description, "This is a test product")
    assert.Equal(t, product.Price, 10.99)
    assert.Equal(t, product.Stock, int64(100))
    assert.Equal(t, product.CategoryID, int64(1))
}

func TestProductService_ListProducts(t *testing.T) {
    mockRepo := &MockProductRepository{}
    productService := service.New ProductService(mockRepo)

    products, err := productService.ListProducts(context.Background(), 1, 10)
    assert.NoError(t, err)
    assert.Len(t, products, 1)
    assert.Equal(t, products[0].ProductID, int64(1))
    assert.Equal(t, products[0].Name, "Test Product")
    assert.Equal(t, products[0].Description, "This is a test product")
    assert.Equal(t, products[0].Price, 10.99)
    assert.Equal(t, products[0].Stock, int64(100))
    assert.Equal(t, products[0].CategoryID, int64(1))
}

func TestProductService_CreateProduct(t *testing.T) {
    mockRepo := &MockProductRepository{}
    productService := service.New ProductService(mockRepo)

    product := &service.Product规范{
        ProductID: 2,
        Name: "New Product",
        Description: "This is a new product",
        Price: 20.99,
        Stock: 200,
        CategoryID: 1,
    }

    err := productService.CreateProduct(context.Background(), product)
    assert.NoError(t, err)
}

func TestProductService_UpdateProduct(t *testing.T) {
    mockRepo := &MockProductRepository{}
    productService := service.New ProductService(mockRepo)

    product := &service.Product规范{
        ProductID: 2,
        Name: "Updated Product",
        Description: "This is an updated product",
        Price: 25.99,
        Stock: 250,
        CategoryID: 1,
    }

    err := productService.UpdateProduct(context.Background(), 2, product)
    assert.NoError(t, err)
}

func TestProductService_DeleteProduct(t *testing.T) {
    mockRepo := &MockProductRepository{}
    productService := service.New ProductService(mockRepo)

    err := productService.DeleteProduct(context.Background(), 2)
    assert.NoError(t, err)
}