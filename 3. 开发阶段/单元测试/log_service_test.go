package service_test

import (
    "testing"
    "shop-app/3. 开发阶段/代码编写/log/service"
    "shop-app/3. 开发阶段/代码编写/log/repository"
    "github.com/stretchr/testify/assert"
)

type MockLogRepository struct {
    repository LogRepository
}

func (m *MockLogRepository) RecordLog(ctx context.Context, logMessage string) error {
    // 模拟记录日志
    return nil
}

func (m *MockLogRepository) QueryLogs(ctx context.Context, startTime, endTime string, userID int, actionType string) ([]*service.Log规范, error) {
    // 模拟查询日志
    return []*service.Log规范{
        {
            LogID: 1,
            Timestamp: "2023-04-01T12:34:56Z",
            UserID: 1,
            ActionType: "login",
            Details: "User logged in successfully",
        },
    }, nil
}

func TestLogService_RecordLog(t *testing.T) {
    mockRepo := &MockLogRepository{}
    logService := service.NewLogService(mockRepo)

    err := logService.RecordLog(context.Background(), "User logged in successfully")
    assert.NoError(t, err)
}

func TestLogService_QueryLogs(t *testing.T) {
    mockRepo := &MockLogRepository{}
    logService := service.NewLogService(mockRepo)

    logs, err := logService.QueryLogs(context.Background(), "2023-04-01", "2023-04-02", 1, "login")
    assert.NoError(t, err)
    assert.Len(t, logs, 1)
    assert.Equal(t, logs[0].LogID, int64(1))
    assert.Equal(t, logs[0].Timestamp, "2023-04-01T12:34:56Z")
    assert.Equal(t, logs[0].UserID, int64(1))
    assert.Equal(t, logs[0].ActionType, "login")
    assert.Equal(t, logs[0].Details, "User logged in successfully")
}