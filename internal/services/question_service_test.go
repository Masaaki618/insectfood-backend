package services_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/Masaaki618/insectfood-backend/internal/models"
	"github.com/Masaaki618/insectfood-backend/internal/repositories/mock"
	"github.com/Masaaki618/insectfood-backend/internal/services"
)

var _ = Describe("QuestionService", func() {
	var (
		ctrl     *gomock.Controller
		mockRepo *mock.MockIQuestionRepository
		svc      services.IQuestionService
		ctx      context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mock.NewMockIQuestionRepository(ctrl)
		svc = services.NewQuestionService(mockRepo)
		ctx = context.Background()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("GetQuestions", func() {
		Context("全カテゴリのデータが存在する場合", func() {
			It("3カテゴリ合計6件の質問を返す", func() {
				mockRepo.EXPECT().GetRandomQuestionsByCategory(ctx, models.CategoryVisual, 2).Return([]models.Question{
					{
						Body:     "visual",
						Category: models.CategoryVisual,
					},
					{
						Body:     "visual",
						Category: models.CategoryVisual,
					},
				}, nil)
				mockRepo.EXPECT().GetRandomQuestionsByCategory(ctx, models.CategoryPhysical, 2).Return([]models.Question{
					{
						Body:     "physical",
						Category: models.CategoryPhysical,
					},
					{
						Body:     "physical",
						Category: models.CategoryPhysical,
					},
				}, nil)
				mockRepo.EXPECT().GetRandomQuestionsByCategory(ctx, models.CategoryMental, 2).Return([]models.Question{
					{
						Body:     "mental",
						Category: models.CategoryMental,
					},
					{
						Body:     "mental",
						Category: models.CategoryMental,
					},
				}, nil)
				result, err := svc.GetQuestions(ctx)

				Expect(err).To(BeNil())
				Expect(result).To(HaveLen(6))
				Expect(result[0].Category).To(Equal("visual"))
				Expect(result[2].Category).To(Equal("physical"))
				Expect(result[4].Category).To(Equal("mental")) //nolint:typecheck
			})
		})
	})
})
