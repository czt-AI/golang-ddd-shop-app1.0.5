阶段/代码编写/email/handler/email_handler.go
```
package handler

import (
    "net/http"
    "shop-app/2. 设计阶段/接口设计/邮件发送接口设计规范"
    "shop-app/3. 开发阶段/代码编写/email/service"
    "encoding/json"
)

type EmailHandler struct {
    emailService *service.EmailService
}

func (h *EmailHandler) SendEmail(w http.ResponseWriter, r *http.Request) {
    var req Email规范.SendEmailRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := h.emailService.SendEmail(r.Context(), req.RecipientEmail, req.Subject, req.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{
        Message: "Email sent successfully",
    })
}

func (h *EmailHandler) GetEmailTemplates(w http.ResponseWriter, r *http.Request) {
    templates, err := h.emailService.GetEmailTemplates(r.Context())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(Email规范.GetEmailTemplatesResponse{
        Message: "Email templates retrieved successfully",
        Data:    templates,
    })
}