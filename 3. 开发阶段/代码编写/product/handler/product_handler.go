package handler

import (
    "net/http"
    "shop-app/2. 设计阶段/接口设计/商品列表接口设计规范"
    "shop-app/2. 设计阶段/接口设计/商品详情接口设计规范"
    "shop-app/3. 开发阶段/代码编写/product/service"
    "encoding/json"
)

type ProductHandler struct {
    productService *service.ProductService
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Page    int `json:"page"`
        Limit   int `json:"limit"`
        CategoryID int `json:"category_id"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    products, err := h.productService.ListProducts(r.Context(), req.Page, req.Limit)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
        Data    []*Product规范 `json:"data"`
    }{
        Message: "Product list retrieved successfully",
        Data:    products,
    })
}

func (h *ProductHandler) GetProductDetails(w http.ResponseWriter, r *http.Request) {
    var req struct {
        ProductID int `json:"product_id"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    product, err := h.productService.GetProductByID(r.Context(), req.ProductID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
        Data    *Product规范 `json:"data"`
    }{
        Message: "Product details retrieved successfully",
        Data:    product,
    })
}