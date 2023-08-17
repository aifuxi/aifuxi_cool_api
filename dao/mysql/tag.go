package mysql

import (
	"github.com/aifuxi/aifuxi_cool_api/models"
)

func GetTags() (*[]models.Tag, error) {
	tags := new([]models.Tag)

	err := db.Model(models.Tag{}).Find(tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}
