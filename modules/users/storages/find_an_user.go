package userstorage

import (
	"context"
	"local-app/common"
	usermodel "local-app/modules/users/models"

	"gorm.io/gorm"
)

func (s *dbStorage) FindUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string,
) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
	}
	return &user, nil
}
