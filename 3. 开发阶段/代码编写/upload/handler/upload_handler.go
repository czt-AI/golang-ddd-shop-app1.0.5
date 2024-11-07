阶段/代码编写/upload/handler/upload_handler.go
```
package handler

import (
    "net/http"
    "shop-app/2. 设计阶段/接口设计/文件上传接口设计规范"
    "shop-app/3. 开发阶段/代码编写/upload/service"
    "encoding/json"
)

type UploadHandler struct {
    uploadService *service.UploadService
}

func (h *UploadHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
    var req File规范.UploadFileRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    response, err := h.uploadService.UploadFile(r.Context(), &req.File)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(File规范.UploadFileResponse{
        Message: "File uploaded successfully",
        Data:    response,
    })
}

func (h *UploadHandler) DownloadFile(w http.ResponseWriter, r *http.Request) {
    var req File规范.DownloadFileRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    file, err := h.uploadService.DownloadFile(r.Context(), req.FileID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Disposition", "attachment; filename=\"" + file.Name + "\"")
    w.Header().Set("Content-Type", file.ContentType)
    w.Write(file.Data)
}