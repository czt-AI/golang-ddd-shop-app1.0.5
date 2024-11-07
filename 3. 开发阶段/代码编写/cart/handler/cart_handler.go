package handler

import (
    "net/http"
    "shop-app/2. 设计阶段/接口设计/购物车接口设计规范"
    "shop-app/3. 开发阶段/代码编写/cart/service"
    "encoding/json"
)

type CartHandler struct {
    cartService *service.CartService
}

func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
    var req struct {
        UserID   int `json:"user_id"`
        ProductID int `json:"product_id"`
        Quantity int `json:"quantity"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := h.cartService.AddToCart(r.Context(), req.UserID, req.ProductID, req.Quantity)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{
        Message: "Product added to cart successfully",
    })
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
    var req struct {
        UserID int `json:"user_id"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    carts, err := h.cartService.GetCart(r.Context(), req.UserID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
        Data    []*Cart规范 `json:"data"`
    }{
        Message: "Cart retrieved successfully",
        Data:    carts,
    })
}

func (h *CartHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {
    var req struct {
        UserID   int `json:"user_id"`
        ProductID int `json:"product_id"`
        Quantity int `json:"quantity"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := h.cartService.UpdateCart(r.Context(), req.UserID, req.ProductID, req.Quantity)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{
        Message: "Cart updated successfully",
    })
}

func (h *CartHandler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
    var req struct {
        UserID   int `json:"user_id"`
        ProductID int `json:"product_id"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := h.cartService.RemoveFromCart(r.Context(), req.UserID, req.ProductID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{
        Message: "Product removed from cart successfully",
    })
}

func (h *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
    var req struct {
        UserID int `json:"user_id"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := h.cartService.ClearCart(r.Context(), req.UserID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{
        Message: "Cart cleared successfully",
    })
}