package service_test

import (
    "testing"
    "shop-app/3. 开发阶段/代码编写/review/service"
    "shop-app/3. 开发阶段/代码编写/review/repository"
    "github.com/stretchr/testify/assert"
)

type MockReviewRepository struct {
    repository ReviewRepository
}

func (m *MockReviewRepository) SubmitReview(ctx context.Context, productID int, rating int, comment string) error {
    // 模拟提交评价
    return nil
}

func (m *MockReviewRepository) GetReviews(ctx context.Context, productID int) ([]*service.Review规范, error) {
    // 模拟获取评价列表
    return []*service.Review规范{
        {
            ReviewID: 1,
            ProductID: productID,
            UserID: 1,
            Rating: rating,
            Comment: comment,
        },
    }, nil
}

func (m *MockReviewRepository) GetReview(ctx context.Context, reviewID int) (*service.Review规范, error) {
    // 模拟获取单个评价
    return &service.Review规范{
        ReviewID: reviewID,
        ProductID: productID,
        UserID: 1,
        Rating: rating,
        Comment: comment,
    }, nil
}

func (m *MockReviewRepository) ReplyToReview(ctx context.Context, reviewID int, reply string) error {
    // 模拟回复评价
    return nil
}

func TestReviewService_SubmitReview(t *testing.T) {
    mockRepo := &MockReviewRepository{}
    reviewService := service.NewReviewService(mockRepo)

    err := reviewService.SubmitReview(context.Background(), 1, 5, "Great product!")
    assert.NoError(t, err)
}

func TestReviewService_GetReviews(t *testing.T) {
    mockRepo := &MockReviewRepository{}
    reviewService := service.NewReviewService(mockRepo)

    reviews, err := reviewService.GetReviews(context.Background(), 1)
    assert.NoError(t, err)
    assert.Len(t, reviews, 1)
    assert.Equal(t, reviews[0].ProductID, int64(1))
    assert.Equal(t, reviews[0].Rating, int(5))
    assert.Equal(t, reviews[0].Comment, "Great product!")
}

func TestReviewService_GetReview(t *testing.T) {
    mockRepo := &MockReviewRepository{}
    reviewService := service.NewReviewService(mockRepo)

    review, err := reviewService.GetReview(context.Background(), 1)
    assert.NoError(t, err)
    assert.NotNil(t, review)
    assert.Equal(t, review.ReviewID, int64(1))
    assert.Equal(t, review.ProductID, int64(1))
    assert.Equal(t, review.Rating, int(5))
    assert.Equal(t, review.Comment, "Great product!")
}

func TestReviewService_ReplyToReview(t *testing.T) {
    mockRepo := &MockReviewRepository{}
    reviewService := service.NewReviewService(mockRepo)

    err := reviewService.ReplyToReview(context.Background(), 1, "Thank you for your feedback!")
    assert.NoError(t, err)
}