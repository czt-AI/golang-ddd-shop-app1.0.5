阶段/代码编写/support/handler/support_handler.go
```
package handler

import (
    "net/http"
    "shop-app/2. 设计阶段/接口设计/客服咨询接口设计规范"
    "shop-app/3. 开发阶段/代码编写/support/service"
    "encoding/json"
)

type SupportHandler struct {
    supportService *service.SupportService
}

func (h *SupportHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
    var req Support规范.CreateTicketRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    ticket, err := h.supportService.CreateTicket(r.Context(), req.UserID, req.Message)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(Support规范.CreateTicketResponse{
        Message: "Support ticket created successfully",
        Data:    ticket,
    })
}

func (h *SupportHandler) GetTicket(w http.ResponseWriter, r *http.Request) {
    var req Support规范.GetTicketRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    ticket, err := h.supportService.GetTicket(r.Context(), req.TicketID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(Support规范.GetTicketResponse{
        Message: "Support ticket retrieved successfully",
        Data:    ticket,
    })
}

func (h *SupportHandler) UpdateTicketStatus(w http.ResponseWriter, r *http.Request) {
    var req Support规范.UpdateTicketStatusRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := h.supportService.UpdateTicketStatus(r.Context(), req.TicketID, req.Status)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{
        Message: "Support ticket status updated successfully",
    })
}

func (h *SupportHandler) GetTicketsByUser(w http.ResponseWriter, r *http.Request) {
    var req Support规范.GetTicketsByUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    tickets, err := h.supportService.GetTicketsByUser(r.Context(), req.UserID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(Support规范.GetTicketsByUserResponse{
        Message: "Support tickets retrieved successfully",
        Data:    tickets,
    })
}