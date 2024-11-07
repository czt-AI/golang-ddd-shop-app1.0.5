package service_test

import (
    "testing"
    "shop-app/3. 开发阶段/代码编写/user/service"
    "shop-app/3. 开发阶段/代码编写/user/repository"
    "github.com/stretchr/testify/assert"
)

type MockUserRepository struct {
    userRepository repository.UserRepository
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, userID int) (*service.User规范, error) {
    // 模拟获取用户ID
    return &service.User规范{
        UserID: 1,
        Username: "test_user",
        Email: "test@example.com",
        Password: "password",
    }, nil
}

func (m *MockUserRepository) GetUserByUsername(ctx context.Context, username string) (*service.User规范, error) {
    // 模拟获取用户名
    return &service.User规范{
        UserID: 1,
        Username: "test_user",
        Email: "test@example.com",
        Password: "password",
    }, nil
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*service.User规范, error) {
    // 模拟获取邮箱
    return &service.User规范{
        UserID: 1,
        Username: "test_user",
        Email: "test@example.com",
        Password: "password",
    }, nil
}

func (m *MockUserRepository) GetUserByPhone(ctx context.Context, phone string) (*service.User规范, error) {
    // 模拟获取手机号
    return &service.User规范{
        UserID: 1,
        Username: "test_user",
        Email: "test@example.com",
        Password: "password",
    }, nil
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *service.User规范) error {
    // 模拟创建用户
    return nil
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, user *service.User规范) error {
    // 模拟更新用户
    return nil
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, userID int) error {
    // 模拟删除用户
    return nil
}

func TestUserService_GetUserByID(t *testing.T) {
    mockRepo := &MockUserRepository{}
    userService := service.NewUserService(mockRepo)

    user, err := userService.GetUserByID(context.Background(), 1)
    assert.NoError(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, user.UserID, int64(1))
    assert.Equal(t, user.Username, "test_user")
    assert.Equal(t, user.Email, "test@example.com")
    assert.Equal(t, user.Password, "password")
}

func TestUserService_GetUserByUsername(t *testing.T) {
    mockRepo := &MockUserRepository{}
    userService := service.NewUserService(mockRepo)

    user, err := userService.GetUserByUsername(context.Background(), "test_user")
    assert.NoError(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, user.UserID, int64(1))
    assert.Equal(t, user.Username, "test_user")
    assert.Equal(t, user.Email, "test@example.com")
    assert.Equal(t, user.Password, "password")
}

func TestUserService_GetUserByEmail(t *testing.T) {
    mockRepo := &MockUserRepository{}
    userService := service.NewUserService(mockRepo)

    user, err := userService.GetUserByEmail(context.Background(), "test@example.com")
    assert.NoError(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, user.UserID, int64(1))
    assert.Equal(t, user.Username, "test_user")
    assert.Equal(t, user.Email, "test@example.com")
    assert.Equal(t, user.Password, "password")
}

func TestUserService_GetUserByPhone(t *testing.T) {
    mockRepo := &MockUserRepository{}
    userService := service.NewUserService(mockRepo)

    user, err := userService.GetUserByPhone(context.Background(), "1234567890")
    assert.NoError(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, user.UserID, int64(1))
    assert.Equal(t, user.Username, "test_user")
    assert.Equal(t, user.Email, "test@example.com")
    assert.Equal(t, user.Password, "password")
}

func TestUserService_CreateUser(t *testing.T) {
    mockRepo := &MockUserRepository{}
    userService := service.NewUserService(mockRepo)

    user := &service.User规范{
        UserID: 2,
        Username: "new_user",
        Email: "new@example.com",
        Password: "new_secure_password",
    }

    err := userService.CreateUser(context.Background(), user)
    assert.NoError(t, err)
}

func TestUserService_UpdateUser(t *testing.T) {
    mockRepo := &MockUserRepository{}
    userService := service.NewUserService(mockRepo)

    user := &service.User规范{
        UserID: 2,
        Username: "new_user",
        Email: "new@example.com",
        Password: "new_secure_password",
    }

    err := userService.UpdateUser(context.Background(), user)
    assert.NoError(t, err)
}

func TestUserService_DeleteUser(t *testing.T) {
    mockRepo := &MockUserRepository{}
    userService := service.NewUserService(mockRepo)

    err := userService.DeleteUser(context.Background(), 2)
    assert.NoError(t, err)
}