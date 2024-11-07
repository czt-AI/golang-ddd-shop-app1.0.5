阶段/代码编写/sms/handler/sms_handler.go
```
package handler

import (
    "net/http"
    "shop-app/2. 设计阶段/接口设计/短信发送接口设计规范"
    "shop-app/3. 开发阶段/代码编写/sms/service"
    "encoding/json"
)

type SmsHandler struct {
    smsService *service.SmsService
}

func (h *SmsHandler) SendSms(w http.ResponseWriter, r *http.Request) {
    var req Sms规范.SendSmsRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := h.smsService.SendSms(r.Context(), req.RecipientPhone, req.Message)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{
        Message: "SMS sent successfully",
    })
}

func (h *SmsHandler) GetSmsTemplates(w http.ResponseWriter, r *http.Request) {
    templates, err := h.smsService.GetSmsTemplates(r.Context())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(Sms规范.GetSmsTemplatesResponse{
        Message: "SMS templates retrieved successfully",
        Data:    templates,
    })
}