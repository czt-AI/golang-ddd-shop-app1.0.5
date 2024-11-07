阶段/代码编写/setting/handler/setting_handler.go
```
package handler

import (
    "net/http"
    "shop-app/2. 设计阶段/接口设计/系统设置管理接口设计规范"
    "shop-app/3. 开发阶段/代码编写/setting/service"
    "encoding/json"
)

type SettingHandler struct {
    settingService *service.SettingService
}

func (h *SettingHandler) GetSetting(w http.ResponseWriter, r *http.Request) {
    var req Setting规范.GetSettingRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    setting, err := h.settingService.GetSetting(r.Context(), req.SettingName)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(Setting规范.GetSettingResponse{
        Message: "Setting retrieved successfully",
        Data:    setting,
    })
}

func (h *SettingHandler) UpdateSetting(w http.ResponseWriter, r *http.Request) {
    var req Setting规范.UpdateSettingRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := h.settingService.UpdateSetting(r.Context(), req.SettingName, req.SettingValue)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{
        Message: "Setting updated successfully",
    })
}