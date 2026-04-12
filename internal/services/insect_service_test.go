package services_test

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"

	aimock "github.com/Masaaki618/insectfood-backend/internal/infrastructure/ai/mock"
	"github.com/Masaaki618/insectfood-backend/internal/models"
	"github.com/Masaaki618/insectfood-backend/internal/repositories/mock"
	"github.com/Masaaki618/insectfood-backend/internal/services"
)

var _ = Describe("InsectService", func() {
	var (
		ctrl       *gomock.Controller
		mockRepo   *mock.MockIInsectRepository
		mockClaude *aimock.MockIClaudeClient
		svc        services.IInsectService
		ctx        context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mock.NewMockIInsectRepository(ctrl)
		mockClaude = aimock.NewMockIClaudeClient(ctrl)
		svc = services.NewInsectService(mockRepo, mockClaude)
		ctx = context.Background()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("GetInsects", func() {
		Context("DBに昆虫が存在する場合", func() {
			It("昆虫一覧のDTOを返す", func() {
				mockRepo.EXPECT().GetInsects(ctx).Return([]models.Insect{
					{Name: "コオロギ", Difficulty: 1},
				}, nil)

				result, err := svc.GetInsects(ctx)

				Expect(err).To(BeNil())
				Expect(result).To(HaveLen(1))
				Expect(result[0].Name).To(Equal("コオロギ"))
			})
		})

		Context("DBに昆虫が0件の場合", func() {
			It("空のスライスを返す (エラーにはならない)", func() {
				mockRepo.EXPECT().GetInsects(ctx).Return([]models.Insect{}, nil)
				result, err := svc.GetInsects(ctx)
				Expect(err).To(BeNil())
				Expect(result).To(BeEmpty())
			})
		})

		Context("DBエラーが発生した場合", func() {
			It("エラーをラップして返す", func() {
				mockRepo.EXPECT().GetInsects(ctx).Return(nil, fmt.Errorf("db error"))
				result, err := svc.GetInsects(ctx)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("InsectService.GetInsects"))
				Expect(result).To(BeNil())
			})
		})
	})

	Describe("GetInsectByID", func() {
		var insect models.Insect
		BeforeEach(func() {
			insect = models.Insect{
				Model: gorm.Model{ID: 1},
				Name:  "コオロギ", Difficulty: 1,
			}
			mockRepo.EXPECT().GetInsectByID(ctx, insect.ID).Return(&insect, nil)
			mockRepo.EXPECT().GetRadarChartByInsectID(ctx, uint(1)).Return(nil, nil)
		})

		Context("DBに昆虫が存在する場合", func() {
			It("昆虫詳細のDTOを返す", func() {
				mockClaude.EXPECT().GenerateInsectComment(ctx, &insect).Return("テストコメント", nil)

				result, err := svc.GetInsectByID(ctx, 1)

				Expect(err).To(BeNil())
				Expect(result.Name).To(Equal("コオロギ"))
				Expect(result.AIComment).To(Equal("テストコメント"))
			})
		})

		Context("Claude APIが3回失敗した場合", func() {
			It("エラーを返す", func() {
				mockClaude.EXPECT().GenerateInsectComment(ctx, &insect).Return("", fmt.Errorf("api error"))
				result, err := svc.GetInsectByID(ctx, 1)

				Expect(err).To(HaveOccurred())
				Expect(result).To(BeNil())
			})
		})
	})
})
