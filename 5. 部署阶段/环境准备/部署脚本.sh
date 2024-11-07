#!/bin/bash

# 部署脚本

# 设置项目目录
PROJECT_DIR="/path/to/shop-app"

# 创建项目目录
mkdir -p $PROJECT_DIR

# 克隆项目代码
git clone https://github.com/your-repository/shop-app.git $PROJECT_DIR

# 进入项目目录
cd $PROJECT_DIR

# 安装依赖
go mod tidy

# 构建项目
go build -o shop-app .

# 创建服务文件
cat << EOF > /etc/systemd/system/shop-app.service
[Unit]
Description=Shop App Service
After=network.target

[Service]
User=shop-app
Group=shop-app
WorkingDirectory=$PROJECT_DIR
ExecStart=/path/to/shop-app/shop-app
Restart=always
RestartSec=30

[Install]
WantedBy=multi-user.target
EOF

# 启动服务
systemctl start shop-app

# 设置服务开机自启
systemctl enable shop-app

# 输出部署成功信息
echo "Shop App deployment completed successfully."