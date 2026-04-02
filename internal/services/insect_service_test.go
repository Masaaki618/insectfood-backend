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

var _ = Describe("InsectService", func() {
	var (
		ctrl     *gomock.Controller
		mockRepo *mock.MockIInsectRepository
		svc      services.IInsectService
		ctx      context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mock.NewMockIInsectRepository(ctrl)
		svc = services.NewInsectService(mockRepo)
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
	})
})
