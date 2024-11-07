package handler

import (
    "net/http"
    "shop-app/2. 设计阶段/接口设计/物流信息查询接口设计规范"
    "shop-app/3. 开发阶段/代码编写/logistics/service"
    "encoding/json"
)

type LogisticsHandler struct {
    logisticsService *service.LogisticsService
}

func (h *LogisticsHandler) QueryLogisticsInfo(w http.ResponseWriter, r *http.Request) {
    var req Logistics规范.QueryLogisticsInfoRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    logistics, err := h.logisticsService.QueryLogisticsInfo(r.Context(), req.OrderID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(Logistics规范.QueryLogisticsInfoResponse{
        Message: "Logistics information retrieved successfully",
        Data:    logistics,
    })
}

func (h *LogisticsHandler) UpdateLogisticsStatus(w http.ResponseWriter, r *http.Request) {
    var req Logistics规范.UpdateLogisticsStatusRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := h.logisticsService.UpdateLogisticsStatus(r.Context(), req.OrderID, req.Status)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{
        Message: "Logistics status updated successfully",
    })
}

func (h *LogisticsHandler) TrackLogistics(w http.ResponseWriter, r *http.Request) {
    var req Logistics规范.TrackLogisticsRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    logistics, err := h.logisticsService.TrackLogistics(r.Context(), req.OrderID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(Logistics规范.TrackLogisticsResponse{
        Message: "Logistics tracked successfully",
        Data:    logistics,
    })
}