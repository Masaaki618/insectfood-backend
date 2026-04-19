package services_test

import (
	"context"
	"fmt"

	"github.com/Masaaki618/insectfood-backend/internal/dtos"
	aimock "github.com/Masaaki618/insectfood-backend/internal/infrastructure/ai/mock"
	"github.com/Masaaki618/insectfood-backend/internal/models"
	"github.com/Masaaki618/insectfood-backend/internal/repositories/mock"
	"github.com/Masaaki618/insectfood-backend/internal/services"
	. "github.com/onsi/ginkgo/v2"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"

	. "github.com/onsi/gomega"
)

var _ = Describe("DiagnosisService", func() {
	var (
		ctrl       *gomock.Controller
		mockRepo   *mock.MockIInsectRepository
		mockClaude *aimock.MockIClaudeClient
		svc        services.IDiagnosisService
		ctx        context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mock.NewMockIInsectRepository(ctrl)
		mockClaude = aimock.NewMockIClaudeClient(ctrl)
		svc = services.NewDiagnosisService(mockRepo, mockClaude)
		ctx = context.Background()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Diagnose", func() {
		var insects []models.Insect
		var req dtos.DiagnosisRequest
		BeforeEach(func() {
			insects = []models.Insect{
				{Model: gorm.Model{ID: 1}, Name: "コオロギ"},
				{Model: gorm.Model{ID: 2}, Name: "ミールワーム"},
			}

			req = dtos.DiagnosisRequest{
				Scores: dtos.DiagnosisScores{
					Visual:   2,
					Physical: 0,
					Mental:   2,
				},
			}
		})
		Context("正常系", func() {
			It("診断結果のDTOを返すこと", func() {
				mockRepo.EXPECT().GetInsects(ctx).Return(insects, nil)
				mockClaude.EXPECT().GenerateDiagnosisResult(ctx, req.Scores.Visual, req.Scores.Physical, req.Scores.Mental, insects).Return(uint(1), "コメント", nil)

				res, err := svc.Diagnose(ctx, req)
				Expect(err).To(BeNil())
				Expect(res.AIComment).To(Equal("コメント"))
				Expect(res.Insect.Name).To(Equal("コオロギ"))

			})
		})

		Context("Claude APIが失敗した場合", func() {
			It("デフォルトの診断結果を返すこと", func() {
				mockRepo.EXPECT().GetInsects(ctx).Return(insects, nil)
				mockClaude.EXPECT().GenerateDiagnosisResult(ctx, req.Scores.Visual, req.Scores.Physical, req.Scores.Mental, insects).Return(uint(0), "", fmt.Errorf("GenerateDiagnosisResult JSONパース失敗"))

				res, err := svc.Diagnose(ctx, req)
				Expect(err).To(BeNil())
				Expect(res.AIComment).To(Equal(fmt.Sprintf("まずは%sから始めてみよう!", insects[0].Name)))
			})
		})

		Context("DBエラーが発生した時に", func() {
			It("エラーを返すこと", func() {
				mockRepo.EXPECT().GetInsects(ctx).Return(insects, fmt.Errorf("db error"))
				res, err := svc.Diagnose(ctx, req)
				Expect(err.Error()).To(ContainSubstring("diagnosisService.Diagnose"))
				Expect(res).To(BeNil())
			})
		})
	})
})
