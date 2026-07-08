package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/Masaaki618/insectfood-backend/internal/controllers"
	"github.com/Masaaki618/insectfood-backend/internal/dtos"
	"github.com/Masaaki618/insectfood-backend/internal/services"
	"github.com/Masaaki618/insectfood-backend/internal/services/mock"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("InsectController", func() {
	var (
		ctrl        *gomock.Controller
		mockService *mock.MockIInsectService
		controller  *controllers.InsectController
		router      *gin.Engine
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		mockService = mock.NewMockIInsectService(ctrl)
		controller = controllers.NewInsectController(mockService)
		router = gin.New()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("GetInsects", func() {
		BeforeEach(func() {
			router.GET("/api/v1/insects", controller.GetInsects)
		})

		Context("正常系", func() {
			It("200と昆虫一覧を返すこと", func() {
				mockService.EXPECT().GetInsects(gomock.Any()).Return([]dtos.InsectResponse{
					{ID: 1, Name: "コオロギ"},
				}, nil)

				req := httptest.NewRequest(http.MethodGet, "/api/v1/insects", nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})

		Context("サービスがエラーを返した場合", func() {
			It("500を返すこと", func() {
				mockService.EXPECT().GetInsects(gomock.Any()).Return(nil, fmt.Errorf("db error"))

				req := httptest.NewRequest(http.MethodGet, "/api/v1/insects", nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(w.Body.String()).To(ContainSubstring(`"code":500`))
				Expect(w.Body.String()).To(ContainSubstring(`"message":"internal server error"`))
			})
		})
	})

	Describe("GetInsectByID", func() {
		BeforeEach(func() {
			router.GET("/api/v1/insects/:id", controller.GetInsectByID)
		})

		Context("正常系", func() {
			It("200と昆虫詳細を返すこと", func() {
				mockService.EXPECT().GetInsectByID(gomock.Any(), uint(1)).Return(&dtos.InsectDetailResponse{
					InsectResponse: dtos.InsectResponse{ID: 1, Name: "コオロギ"},
					AIComment:      "テストコメント",
				}, nil)

				req := httptest.NewRequest(http.MethodGet, "/api/v1/insects/1", nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})

		Context("IDに数字以外が渡された場合", func() {
			It("400を返すこと", func() {
				req := httptest.NewRequest(http.MethodGet, "/api/v1/insects/abc", nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(w.Body.String()).To(ContainSubstring(`"code":400`))
				Expect(w.Body.String()).To(ContainSubstring(`"message":"invalid id"`))
			})
		})

		Context("昆虫が存在しない場合", func() {
			It("404を返すこと", func() {
				mockService.EXPECT().GetInsectByID(gomock.Any(), uint(999)).Return(nil, services.ErrNotFound)

				req := httptest.NewRequest(http.MethodGet, "/api/v1/insects/999", nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusNotFound))
				Expect(w.Body.String()).To(ContainSubstring(`"code":404`))
				Expect(w.Body.String()).To(ContainSubstring(`"message":"insect not found"`))
			})
		})

		Context("サービスがエラーを返した場合", func() {
			It("500を返すこと", func() {
				mockService.EXPECT().GetInsectByID(gomock.Any(), uint(1)).Return(nil, fmt.Errorf("db error"))

				req := httptest.NewRequest(http.MethodGet, "/api/v1/insects/1", nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(w.Body.String()).To(ContainSubstring(`"code":500`))
			})
		})
	})
})
