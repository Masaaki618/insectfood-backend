package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/Masaaki618/insectfood-backend/internal/controllers"
	"github.com/Masaaki618/insectfood-backend/internal/dtos"
	"github.com/Masaaki618/insectfood-backend/internal/services/mock"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("DiagnosisController", func() {
	var (
		ctrl        *gomock.Controller
		mockService *mock.MockIDiagnosisService
		controller  *controllers.DiagnosisController
		router      *gin.Engine
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		mockService = mock.NewMockIDiagnosisService(ctrl)
		controller = controllers.NewDiagnosisController(mockService)
		router = gin.New()
		router.POST("/api/v1/diagnosis", controller.Diagnose)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Diagnose", func() {
		Context("正常系", func() {
			It("200と診断結果を返すこと", func() {
				mockService.EXPECT().Diagnose(gomock.Any(), gomock.Any()).Return(&dtos.DiagnosisResponse{
					Insect:    dtos.InsectResponse{ID: 1, Name: "コオロギ"},
					AIComment: "テストコメント",
				}, nil)

				body := `{"scores":{"visual":1,"physical":0,"mental":2}}`
				req := httptest.NewRequest(http.MethodPost, "/api/v1/diagnosis", strings.NewReader(body))
				req.Header.Set("Content-Type", "application/json")
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})

		Context("リクエストが不正な場合", func() {
			It("400を返すこと", func() {
				body := `{"scores":{"visual":3,"physical":0,"mental":2}}`
				req := httptest.NewRequest(http.MethodPost, "/api/v1/diagnosis", strings.NewReader(body))
				req.Header.Set("Content-Type", "application/json")
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(w.Body.String()).To(ContainSubstring(`"code":400`))
				Expect(w.Body.String()).To(ContainSubstring(`"message":"invalid request"`))
			})
		})

		Context("サービスがエラーを返した場合", func() {
			It("500を返すこと", func() {
				mockService.EXPECT().Diagnose(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("error"))

				body := `{"scores":{"visual":1,"physical":0,"mental":2}}`
				req := httptest.NewRequest(http.MethodPost, "/api/v1/diagnosis", strings.NewReader(body))
				req.Header.Set("Content-Type", "application/json")
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(w.Body.String()).To(ContainSubstring(`"code":500`))
			})
		})
	})
})
