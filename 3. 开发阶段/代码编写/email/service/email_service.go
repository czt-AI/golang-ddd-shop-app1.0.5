阶段/代码编写/email/service/email_service.go
```
package service

import (
    "context"
    "shop-app/2. 设计阶段/服务设计/邮件发送服务设计规范"
    "shop-app/3. 开发阶段/环境搭建/开发环境配置指南"
    "shop-app/3. 开发阶段/代码编写/email/repository"
)

type EmailService struct {
    emailRepository *repository.EmailRepository
}

func NewEmailService(emailRepository *repository.EmailRepository) *EmailService {
    return &EmailService{emailRepository: emailRepository}
}

func (s *EmailService) SendEmail(ctx context.Context, recipientEmail, subject, body string) error {
    // 实现发送邮件逻辑
    // 使用 emailRepository 进行邮件发送
    return nil
}

func (s *EmailService) GetEmailTemplates(ctx context.Context) ([]*Email规范.Template, error) {
    // 实现获取邮件模板逻辑
    // 使用 emailRepository 获取模板信息
    return nil, nil
}