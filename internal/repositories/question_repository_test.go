package repositories_test

import (
	"context"

	"github.com/Masaaki618/insectfood-backend/internal/models"
	"github.com/Masaaki618/insectfood-backend/internal/repositories"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("QuestionRepository", func() {
	var (
		db  *gorm.DB
		ctx context.Context
	)

	BeforeEach(func() {
		var err error
		db, err = newTestDB()
		Expect(err).To(BeNil())
		ctx = context.Background()

		// テストデータを初期化
		db.Exec("DELETE FROM questions")
	})

	Describe("GetRandomQuestionsByCategory", func() {
		Context("質問が存在する場合", func() {
			It("指定カテゴリの質問をlimit件返すこと", func() {
				db.Create(&models.Question{Body: "質問1", Category: models.CategoryVisual})
				db.Create(&models.Question{Body: "質問2", Category: models.CategoryVisual})
				db.Create(&models.Question{Body: "質問3", Category: models.CategoryVisual})

				repo := repositories.NewQuestionRepository(db)
				result, err := repo.GetRandomQuestionsByCategory(ctx, models.CategoryVisual, 2)

				Expect(err).To(BeNil())
				Expect(result).To(HaveLen(2))
				Expect(result[0].Category).To(Equal(models.CategoryVisual))
			})
		})

		Context("指定カテゴリの質問が0件の場合", func() {
			It("空のスライスを返すこと", func() {
				repo := repositories.NewQuestionRepository(db)
				result, err := repo.GetRandomQuestionsByCategory(ctx, models.CategoryVisual, 2)

				Expect(err).To(BeNil())
				Expect(result).To(BeEmpty())
			})
		})
	})
})
