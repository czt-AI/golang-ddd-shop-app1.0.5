package service_test

import (
    "testing"
    "shop-app/3. 开发阶段/代码编写/email/service"
    "shop-app/3. 开发阶段/代码编写/email/repository"
    "github.com/stretchr/testify/assert"
)

type MockEmailRepository struct {
    repository EmailRepository
}

func (m *MockEmailRepository) SendEmail(ctx context.Context, recipientEmail, subject, body string) error {
    // 模拟发送邮件
    return nil
}

func (m *MockEmailRepository) GetEmailTemplates(ctx context.Context) ([]*service.Template, error) {
    // 模拟获取邮件模板
    return []*service.Template{
        {
            TemplateID: 1,
            TemplateName: "Welcome Email",
            Subject: "Welcome to our store",
            Body: "Thank you for joining us!",
        },
    }, nil
}

func TestEmailService_SendEmail(t *testing.T) {
    mockRepo := &MockEmailRepository{}
    emailService := service.NewEmailService(mockRepo)

    err := emailService.SendEmail(context.Background(), "test@example.com", "Welcome Email", "Thank you for joining us!")
    assert.NoError(t, err)
}

func TestEmailService_GetEmailTemplates(t *testing.T) {
    mockRepo := &MockEmailRepository{}
    emailService := service.NewEmailService(mockRepo)

    templates, err := emailService.GetEmailTemplates(context.Background())
    assert.NoError(t, err)
    assert.Len(t, templates, 1)
    assert.Equal(t, templates[0].TemplateID, int64(1))
    assert.Equal(t, templates[0].TemplateName, "Welcome Email")
    assert.Equal(t, templates[0].Subject, "Welcome to our store")
    assert.Equal(t, templates[0].Body, "Thank you for joining us!")
}