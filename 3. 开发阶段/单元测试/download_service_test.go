package service_test

import (
    "testing"
    "shop-app/3. 开发阶段/代码编写/download/service"
    "shop-app/3. 开发阶段/代码编写/download/repository"
    "github.com/stretchr/testify/assert"
)

type MockDownloadRepository struct {
    repository DownloadRepository
}

func (m *MockDownloadRepository) DownloadFile(ctx context.Context, fileID int) (*service.File规范, error) {
    // 模拟文件下载
    return &service.File规范{
        Name: "test_file.txt",
        ContentType: "text/plain",
        Data: []byte("This is a test file."),
    }, nil
}

func TestDownloadService_DownloadFile(t *testing.T) {
    mockRepo := &MockDownloadRepository{}
    downloadService := service.NewDownloadService(mockRepo)

    file, err := downloadService.DownloadFile(context.Background(), 1)
    assert.NoError(t, err)
    assert.NotNil(t, file)
    assert.Equal(t, file.Name, "test_file.txt")
    assert.Equal(t, file.ContentType, "text/plain")
    assert.Equal(t, string(file.Data), "This is a test file.")
}