package mysql

import (
	"time"

	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/internal"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
)

func GetTags() (*[]models.Tag, error) {
	tags := new([]models.Tag)

	err := db.Where("deleted_at is null").Find(tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func GetTagByID(id int64) (*models.Tag, error) {
	tag := new(models.Tag)
	err := db.Where("deleted_at is null").First(&tag, id).Error
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func UpdateTagByID(id int64, data *dto.UpdateTagDTO) error {
	var tag = &models.Tag{
		ID: id,
	}
	err := db.Model(tag).Where("deleted_at is null").Limit(1).Updates(
		models.Tag{
			Name:        data.Name,
			FriendlyUrl: data.FriendlyUrl,
		}).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteTagByID(id int64) error {
	if !TagExistsByID(id) {
		return myerror.ErrorTagNotFound
	}

	err := db.Model(models.Tag{}).Where("id = ?", id).Limit(1).Update("deleted_at", time.Now().Local().Format(time.DateTime)).Error
	if err != nil {
		return err
	}

	return nil
}

func TagExistsByName(name string) bool {
	tag := new(models.Tag)
	db.Where("deleted_at is null and name = ?", name).First(&tag)
	return tag.ID != 0
}

func TagExistsByID(id int64) bool {
	tag := new(models.Tag)
	db.Where("deleted_at is null").First(&tag, id)
	return tag.ID != 0
}

func CreateTag(data *dto.CreateTagDTO) (*models.Tag, error) {
	if exists := TagExistsByName(data.Name); exists {
		return nil, myerror.ErrorTagExists
	}

	id, err := internal.GenSnowflakeID()
	if err != nil {
		return nil, err
	}

	tag := &models.Tag{
		ID:          id,
		Name:        data.Name,
		FriendlyUrl: data.FriendlyUrl,
	}
	err = db.Model(&models.Tag{}).Create(tag).Error
	if err != nil {
		return nil, err
	}

	return tag, nil
}
