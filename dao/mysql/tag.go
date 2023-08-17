package mysql

import (
	"errors"
	"time"

	"github.com/aifuxi/aifuxi_cool_api/dto"
	"github.com/aifuxi/aifuxi_cool_api/models"
)

func GetTags() (*[]models.Tag, error) {
	tags := new([]models.Tag)

	err := db.Where("deleted_at is null").Find(tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func GetTagByID(id any) (*models.Tag, error) {
	tag := new(models.Tag)
	err := db.Where("deleted_at is null").First(&tag, id).Error
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func DeleteTagByID(id any) error {
	if !TagExistsByID(id) {
		return ErrorTagNotFound
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

func TagExistsByID(id any) bool {
	tag := new(models.Tag)
	db.Where("deleted_at is null").First(&tag, id)
	return tag.ID != 0
}

func CreateTag(data *dto.CreateTagDTO) (*models.Tag, error) {
	if exists := TagExistsByName(data.Name); exists {
		return nil, errors.New("文章标签已存在")
	}

	tag := &models.Tag{
		ID:          time.Now().Unix(),
		Name:        data.Name,
		FriendlyUrl: data.FriendlyUrl,
	}
	err := db.Model(&models.Tag{}).Create(tag).Error

	if err != nil {
		return nil, err
	}

	return tag, nil
}
