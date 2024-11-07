package service_test

import (
    "testing"
    "shop-app/3. 开发阶段/代码编写/setting/service"
    "shop-app/3. 开发阶段/代码编写/setting/repository"
    "github.com/stretchr/testify/assert"
)

type MockSettingRepository struct {
    repository SettingRepository
}

func (m *MockSettingRepository) GetSetting(ctx context.Context, settingName string) (*service.Setting规范, error) {
    // 模拟获取系统设置
    return &service.Setting规范{
        SettingName: "site_name",
        SettingValue: "ShopApp",
    }, nil
}

func (m *MockSettingRepository) UpdateSetting(ctx context.Context, settingName, settingValue string) error {
    // 模拟更新系统设置
    return nil
}

func TestSettingService_GetSetting(t *testing.T) {
    mockRepo := &MockSettingRepository{}
    settingService := service.NewSettingService(mockRepo)

    setting, err := settingService.GetSetting(context.Background(), "site_name")
    assert.NoError(t, err)
    assert.NotNil(t, setting)
    assert.Equal(t, setting.SettingName, "site_name")
    assert.Equal(t, setting.SettingValue, "ShopApp")
}

func TestSettingService_UpdateSetting(t *testing.T) {
    mockRepo := &MockSettingRepository{}
    settingService := service.NewSettingService(mockRepo)

    err := settingService.UpdateSetting(context.Background(), "site_name", "NewShopApp")
    assert.NoError(t, err)
}