package service_test

import (
    "testing"
    "shop-app/3. 开发阶段/代码编写/sms/service"
    "shop-app/3. 开发阶段/代码编写/sms/repository"
    "github.com/stretchr/testify/assert"
)

type MockSmsRepository struct {
    repository SmsRepository
}

func (m *MockSmsRepository) SendSms(ctx context.Context, recipientPhone, message string) error {
    // 模拟发送短信
    return nil
}

func (m *MockSmsRepository) GetSmsTemplates(ctx context.Context) ([]*service.Template, error) {
    // 模拟获取短信模板
    return []*service.Template{
        {
            TemplateID: 1,
            TemplateName: "Welcome SMS",
            Message: "Welcome to our store!",
        },
    }, nil
}

func TestSmsService_SendSms(t *testing.T) {
    mockRepo := &MockSmsRepository{}
    smsService := service.NewSmsService(mockRepo)

    err := smsService.SendSms(context.Background(), "1234567890", "Welcome SMS")
    assert.NoError(t, err)
}

func TestSmsService_GetSmsTemplates(t *testing.T) {
    mockRepo := &MockSmsRepository{}
    smsService := service.NewSmsService(mockRepo)

    templates, err := smsService.GetSmsTemplates(context.Background())
    assert.NoError(t, err)
    assert.Len(t, templates, 1)
    assert.Equal(t, templates[0].TemplateID, int64(1))
    assert.Equal(t, templates[0].TemplateName, "Welcome SMS")
    assert.Equal(t, templates[0].Message, "Welcome to our store!")
}