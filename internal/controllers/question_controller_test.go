package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/Masaaki618/insectfood-backend/internal/controllers"
	"github.com/Masaaki618/insectfood-backend/internal/dtos"
	"github.com/Masaaki618/insectfood-backend/internal/services/mock"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("QuestionController", func() {
	var (
		ctrl        *gomock.Controller
		mockService *mock.MockIQuestionService
		controller  *controllers.QuestionController
		router      *gin.Engine
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		mockService = mock.NewMockIQuestionService(ctrl)
		controller = controllers.NewQuestionController(mockService)
		router = gin.New()
		router.GET("/api/v1/questions", controller.GetQuestions)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("GetQuestions", func() {
		Context("正常系", func() {
			It("200と質問一覧を返すこと", func() {
				mockService.EXPECT().GetQuestions(gomock.Any()).Return([]dtos.QuestionResponse{
					{Body: "テスト質問", Category: "visual"},
				}, nil)

				req := httptest.NewRequest(http.MethodGet, "/api/v1/questions", nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})

		Context("サービスがエラーを返した場合", func() {
			It("500を返すこと", func() {
				mockService.EXPECT().GetQuestions(gomock.Any()).Return(nil, fmt.Errorf("db error"))

				req := httptest.NewRequest(http.MethodGet, "/api/v1/questions", nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(w.Body.String()).To(ContainSubstring(`"code":500`))
				Expect(w.Body.String()).To(ContainSubstring(`"message":"internal server error"`))
			})
		})
	})
})
