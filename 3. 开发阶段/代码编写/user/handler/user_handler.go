package handler

import (
    "net/http"
    "shop-app/2. 设计阶段/接口设计/用户注册接口设计规范"
    "shop-app/2. 设计阶段/接口设计/用户登录接口设计规范"
    "shop-app/3. 开发阶段/代码编写/user/service"
    "encoding/json"
)

type UserHandler struct {
    userService *service.UserService
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
    var req user规范.UserRegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user, err := h.userService.Register(r.Context(), &req.User)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user规范.UserRegisterResponse{
        Message: "User registered successfully",
        Data:    user,
    })
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    var req user规范.UserLoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user, err := h.userService.Login(r.Context(), req.Username, req.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user规范.UserLoginResponse{
        Message: "Login successful",
        Data:    user,
    })
}