阶段/代码编写/log/handler/log_handler.go
```
package handler

import (
    "net/http"
    "shop-app/2. 设计阶段/接口设计/日志管理接口设计规范"
    "shop-app/3. 开发阶段/代码编写/log/service"
    "encoding/json"
)

type LogHandler struct {
    logService *service.LogService
}

func (h *LogHandler) RecordLog(w http.ResponseWriter, r *http.Request) {
    var req Log规范.RecordLogRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := h.logService.RecordLog(r.Context(), req.LogMessage)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{
        Message: "Log recorded successfully",
    })
}

func (h *LogHandler) QueryLogs(w http.ResponseWriter, r *http.Request) {
    var req Log规范.QueryLogsRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    logs, err := h.logService.QueryLogs(r.Context(), req.StartTime, req.EndTime, req.UserID, req.ActionType)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(Log规范.QueryLogsResponse{
        Message: "Logs retrieved successfully",
        Data:    logs,
    })
}