package mysql

import (
	"fmt"
	"time"

	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/internal"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
)

func GetTags(data *dto.GetTagsDTO) (*[]models.Tag, int64, error) {
	tags := new([]models.Tag)
	var total int64
	var nameLike, friendlyUrlLike string

	order := fmt.Sprintf("%s %s", data.OrderBy, data.Order)
	if len(data.Name) > 0 {
		nameLike = "%" + data.Name + "%"
	}

	if len(data.FriendlyUrl) > 0 {
		friendlyUrlLike = "%" + data.FriendlyUrl + "%"
	}

	if len(data.Name) == 0 && len(data.FriendlyUrl) == 0 {
		nameLike = "%" + data.Name + "%"
		friendlyUrlLike = "%" + data.FriendlyUrl + "%"
	}

	err := db.Order(order).Where("deleted_at is null").Where(
		db.Where("name LIKE ?", nameLike).Or("friendly_url LIKE ?", friendlyUrlLike),
	).Offset((data.Page - 1) * data.PageSize).Limit(data.PageSize).Find(tags).Error
	if err != nil {
		return nil, total, err
	}

	err = db.Model(models.Tag{}).Where("deleted_at is null").Where(
		db.Where("name LIKE ?", nameLike).Or("friendly_url LIKE ?", friendlyUrlLike),
	).Count(&total).Error
	if err != nil {
		return nil, total, err
	}

	return tags, total, nil
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
