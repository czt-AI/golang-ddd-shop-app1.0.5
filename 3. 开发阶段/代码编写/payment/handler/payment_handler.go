package handler

import (
    "net/http"
    "shop-app/2. 设计阶段/接口设计/支付接口设计规范"
    "shop-app/3. 开发阶段/代码编写/payment/service"
    "encoding/json"
)

type PaymentHandler struct {
    paymentService *service.PaymentService
}

func (h *PaymentHandler) InitiatePayment(w http.ResponseWriter, r *http.Request) {
    var req Payment规范.InitiatePaymentRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    payment, err := h.paymentService.InitiatePayment(r.Context(), req.OrderID, req.PaymentMethod)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(Payment规范.InitiatePaymentResponse{
        Message: "Payment initiated successfully",
        Data:    payment,
    })
}

func (h *PaymentHandler) ExecutePayment(w http.ResponseWriter, r *http.Request) {
    var req Payment规范.ExecutePaymentRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    payment, err := h.paymentService.ExecutePayment(r.Context(), req.PaymentID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(Payment规范.ExecutePaymentResponse{
        Message: "Payment executed successfully",
        Data:    payment,
    })
}

func (h *PaymentHandler) QueryPayment(w http.ResponseWriter, r *http.Request) {
    var req Payment规范.QueryPaymentRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    payment, err := h.paymentService.QueryPayment(r.Context(), req.PaymentID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(Payment规范.QueryPaymentResponse{
        Message: "Payment queried successfully",
        Data:    payment,
    })
}