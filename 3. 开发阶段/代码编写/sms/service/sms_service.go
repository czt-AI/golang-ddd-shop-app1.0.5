阶段/代码编写/sms/service/sms_service.go
```
package service

import (
    "context"
    "shop-app/2. 设计阶段/服务设计/短信发送服务设计规范"
    "shop-app/3. 开发阶段/环境搭建/开发环境配置指南"
    "shop-app/3. 开发阶段/代码编写/sms/repository"
)

type SmsService struct {
    smsRepository *repository.SmsRepository
}

func NewSmsService(smsRepository *repository.SmsRepository) *SmsService {
    return &SmsService{smsRepository: smsRepository}
}

func (s *SmsService) SendSms(ctx context.Context, recipientPhone, message string) error {
    // 实现发送短信逻辑
    // 使用 smsRepository 进行短信发送
    return nil
}

func (s *SmsService) GetSmsTemplates(ctx context.Context) ([]*Sms规范.Template, error) {
    // 实现获取短信模板逻辑
    // 使用 smsRepository 获取模板信息
    return nil, nil
}