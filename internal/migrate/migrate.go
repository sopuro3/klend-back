package migrate

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/sopuro3/klend-back/internal/model"
	"github.com/sopuro3/klend-back/pkg/password"
	"github.com/sopuro3/klend-back/pkg/password/argon2"
)

//nolint:wrapcheck
func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.Issue{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.LoanEntry{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.Equipment{}); err != nil {
		return err
	}

	return nil
}

//nolint:funlen
func Seed(db *gorm.DB) {
	var count int64

	db.Model(&model.Equipment{}).Count(&count)

	if count > 0 {
		return
	}

	//nolint:gomnd,lll
	equipments := []*model.Equipment{
		{Model: model.Model{ID: uuid.MustParse("018c7b9f8c55708f803527a5528e83ed")}, Name: "角スコップ", MaxQuantity: 20, Note: "てすとてすとてすと"},
		{Model: model.Model{ID: uuid.MustParse("018c7ba8d2df7adcaf3dbe411ce1cb60")}, Name: "バケツ", MaxQuantity: 99, Note: "てすとてすとてすと"},
	}

	//nolint:gomnd,lll
	loanEntries := []*model.LoanEntry{
		{Model: model.Model{ID: uuid.MustParse("018cf5eb-c686-75b7-8413-1d61612bd1b9")}, EquipmentID: uuid.MustParse("018c7b9f8c55708f803527a5528e83ed"), Quantity: 10},
		{Model: model.Model{ID: uuid.MustParse("018cf5ec-0faa-7378-9dea-e832670afdc7")}, EquipmentID: uuid.MustParse("018c7ba8d2df7adcaf3dbe411ce1cb60"), Quantity: 20},
		// {Model: model.Model{ID: uuid.MustParse("018cfd8b-ee64-71c2-929c-e8d1cca5c2f0")}, EquipmentID: uuid.MustParse("018c7b9f8c55708f803527a5528e83ed"), Quantity: 5},
	}

	//nolint:lll
	issues := []*model.Issue{
		{Model: model.Model{ID: uuid.MustParse("018c7765-ffd5-724f-aa7f-227175f54d3f")}, Address: "小森野1-1-1", Name: "久留米太郎", DisplayID: "0001", Status: "survey", Note: "テストデータ", LoanEntries: loanEntries[0:2]},
		// {Model: model.Model{ID: uuid.MustParse("018cfd89-67cd-77f2-955e-da5439bb8d7e")}, Address: "小森野1-1-2", Name: "久留米次郎", DisplayID: "0002", Status: "start", Note: "テストデータ", IsConfirmed: true, LoanEntries: loanEntries[2:]},
	}

	passwordEncoder := password.Encoder(argon2.NewArgon2Encoder())

	hashedPassword1, err := passwordEncoder.EncodePassword("password")
	if err != nil {
		slog.Warn(err.Error())
	}

	hashedPassword2, err := passwordEncoder.EncodePassword("test")
	if err != nil {
		slog.Warn(err.Error())
	}

	users := []*model.User{
		//nolint:lll
		{Model: model.Model{ID: uuid.MustParse("018d08be-febd-7763-b466-05174ab3f4d1")}, ExternalUserID: "admin", Email: "test@example.com", UserName: "久留米太郎", HashedPassword: string(hashedPassword1)},
		//nolint:lll
		{Model: model.Model{ID: uuid.MustParse("018d08c7-4169-7550-9f9b-bba288c03882")}, ExternalUserID: "user1", Email: "user1@example.com", UserName: "久留米次郎", HashedPassword: string(hashedPassword2)},
	}

	if err := db.Create(&equipments).Error; err != nil {
		slog.Warn("%+v", err)
	}

	if err := db.Create(&loanEntries).Error; err != nil {
		slog.Warn("%+v", err)
	}

	if err := db.Create(&issues).Error; err != nil {
		slog.Warn("%+v", err)
	}

	if err := db.Create(&users).Error; err != nil {
		slog.Warn("err", err)
	}
}
