package service_test

import (
    "testing"
    "shop-app/3. 开发阶段/代码编写/upload/service"
    "shop-app/3. 开发阶段/代码编写/upload/repository"
    "github.com/stretchr/testify/assert"
)

type MockUploadRepository struct {
    repository UploadRepository
}

func (m *MockUploadRepository) UploadFile(ctx context.Context, file *service.File规范) (*service.UploadResponse, error) {
    // 模拟文件上传
    return &service.UploadResponse{
        FileID: 1,
        FileName: file.Name,
        ContentType: file.ContentType,
    }, nil
}

func (m *MockUploadRepository) DownloadFile(ctx context.Context, fileID int) (*service.File规范, error) {
    // 模拟文件下载
    return &service.File规范{
        Name: "test_file.txt",
        ContentType: "text/plain",
        Data: []byte("This is a test file."),
    }, nil
}

func TestUploadService_UploadFile(t *testing.T) {
    mockRepo := &MockUploadRepository{}
    uploadService := service.NewUploadService(mockRepo)

    file := &service.File规范{
        Name: "test_file.txt",
        ContentType: "text/plain",
        Data: []byte("This is a test file."),
    }

    response, err := uploadService.UploadFile(context.Background(), file)
    assert.NoError(t, err)
    assert.NotNil(t, response)
    assert.Equal(t, response.FileID, int64(1))
    assert.Equal(t, response.FileName, "test_file.txt")
    assert.Equal(t, response.ContentType, "text/plain")
}

func TestUploadService_DownloadFile(t *testing.T) {
    mockRepo := &MockUploadRepository{}
    uploadService := service.NewUploadService(mockRepo)

    file, err := uploadService.DownloadFile(context.Background(), 1)
    assert.NoError(t, err)
    assert.NotNil(t, file)
    assert.Equal(t, file.Name, "test_file.txt")
    assert.Equal(t, file.ContentType, "text/plain")
    assert.Equal(t, string(file.Data), "This is a test file.")
}