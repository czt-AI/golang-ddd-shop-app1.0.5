package handler

import (
    "net/http"
    "shop-app/2. 设计阶段/接口设计/订单创建接口设计规范"
    "shop-app/2. 设计阶段/接口设计/订单查询接口设计规范"
    "shop-app/3. 开发阶段/代码编写/order/service"
    "encoding/json"
)

type OrderHandler struct {
    orderService *service.OrderService
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
    var req Order规范.CreateOrderRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    order, err := h.orderService.CreateOrder(r.Context(), &req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(Order规范.CreateOrderResponse{
        Message: "Order created successfully",
        Data:    order,
    })
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
    var req Order规范.GetOrderRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    order, err := h.orderService.GetOrderByID(r.Context(), req.OrderID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(Order规范.GetOrderResponse{
        Message: "Order retrieved successfully",
        Data:    order,
    })
}

func (h *OrderHandler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
    var req Order规范.UpdateOrderStatusRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := h.orderService.UpdateOrderStatus(r.Context(), req.OrderID, req.Status)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{
        Message: "Order status updated successfully",
    })
}

func (h *OrderHandler) CancelOrder(w http.ResponseWriter, r *http.Request) {
    var req Order规范.CancelOrderRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := h.orderService.CancelOrder(r.Context(), req.OrderID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{
        Message: "Order cancelled successfully",
    })
}