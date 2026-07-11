package repositories_test

import (
	"context"

	"github.com/Masaaki618/insectfood-backend/internal/models"
	"github.com/Masaaki618/insectfood-backend/internal/repositories"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("InsectRepository", func() {
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
		db.Exec("DELETE FROM radar_charts")
		db.Exec("DELETE FROM insects")
	})

	Describe("GetInsects", func() {
		Context("昆虫が存在する場合", func() {
			It("昆虫一覧を返すこと", func() {
				db.Create(&models.Insect{Name: "コオロギ", Difficulty: 1, Introduction: "テスト", Taste: "ナッツ", Texture: "サクサク"})
				db.Create(&models.Insect{Name: "ミールワーム", Difficulty: 1, Introduction: "テスト", Taste: "淡白", Texture: "サクサク"})

				repo := repositories.NewInsectRepository(db)
				result, err := repo.GetInsects(ctx)

				Expect(err).To(BeNil())
				Expect(result).To(HaveLen(2))
			})
		})

		Context("昆虫が0件の場合", func() {
			It("空のスライスを返すこと", func() {
				repo := repositories.NewInsectRepository(db)
				result, err := repo.GetInsects(ctx)

				Expect(err).To(BeNil())
				Expect(result).To(BeEmpty())
			})
		})
	})

	Describe("GetInsectByID", func() {
		Context("昆虫が存在する場合", func() {
			It("昆虫詳細を返すこと", func() {
				insect := models.Insect{Name: "コオロギ", Difficulty: 1, Introduction: "テスト", Taste: "ナッツ", Texture: "サクサク"}
				db.Create(&insect)

				repo := repositories.NewInsectRepository(db)
				result, err := repo.GetInsectByID(ctx, insect.ID)

				Expect(err).To(BeNil())
				Expect(result.Name).To(Equal("コオロギ"))
			})
		})

		Context("存在しないIDを指定した場合", func() {
			It("ErrRecordNotFoundを含むエラーを返すこと", func() {
				repo := repositories.NewInsectRepository(db)
				result, err := repo.GetInsectByID(ctx, 99999)

				Expect(err).To(HaveOccurred())
				Expect(result).To(BeNil())
			})
		})
	})

	Describe("GetRadarChartByInsectID", func() {
		Context("レーダーチャートが存在する場合", func() {
			It("レーダーチャートを返すこと", func() {
				insect := models.Insect{Name: "コオロギ", Difficulty: 1, Introduction: "テスト", Taste: "ナッツ", Texture: "サクサク"}
				db.Create(&insect)
				db.Create(&models.RadarChart{InsectID: insect.ID, UmamiScore: 3, BitterScore: 1, EguScore: 1, FlavorScore: 4, KimoScore: 1})

				repo := repositories.NewInsectRepository(db)
				result, err := repo.GetRadarChartByInsectID(ctx, insect.ID)

				Expect(err).To(BeNil())
				Expect(result).NotTo(BeNil())
				Expect(result.UmamiScore).To(Equal(uint8(3)))
			})
		})

		Context("レーダーチャートが存在しない場合", func() {
			It("nilを返すこと（エラーにはならない）", func() {
				repo := repositories.NewInsectRepository(db)
				result, err := repo.GetRadarChartByInsectID(ctx, 99999)

				Expect(err).To(BeNil())
				Expect(result).To(BeNil())
			})
		})
	})
})
