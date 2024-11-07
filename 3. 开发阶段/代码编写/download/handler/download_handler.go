package handler

import (
    "net/http"
    "shop-app/2. 设计阶段/接口设计/文件下载接口设计规范"
    "shop-app/3. 开发阶段/代码编写/download/service"
    "encoding/json"
)

type DownloadHandler struct {
    downloadService *service.DownloadService
}

func (h *DownloadHandler) DownloadFile(w http.ResponseWriter, r *http.Request) {
    var req File规范.DownloadFileRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    file, err := h.downloadService.DownloadFile(r.Context(), req.FileID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Disposition", "attachment; filename=\"" + file.Name + "\"")
    w.Header().Set("Content-Type", file.ContentType)
    w.Write(file.Data)
}