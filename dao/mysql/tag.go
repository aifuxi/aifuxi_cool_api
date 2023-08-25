package mysql

import (
	"fmt"
	"time"

	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/internal"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
)

func GetTags(arg dto.GetTagsDTO) ([]models.Tag, int64, error) {
	var tags []models.Tag
	var count int64
	var queryDB = db.Model(models.Tag{}).Scopes(isDeletedRecord)

	if len(arg.Name) > 0 {
		queryDB.Where("name LIKE ?", "%"+arg.Name+"%")
	}

	if len(arg.FriendlyUrl) > 0 {
		queryDB.Where("friendly_url LIKE ?", "%"+arg.FriendlyUrl+"%")
	}

	queryDB = queryDB.Count(&count)

	order := fmt.Sprintf("%s %s", arg.OrderBy, arg.Order)
	err := queryDB.Order(order).Scopes(Paginate(arg.Page, arg.PageSize)).Find(&tags).Error
	if err != nil {
		return nil, count, err
	}

	return tags, count, nil
}

func GetTagByID(id int64) (models.Tag, error) {
	var tag models.Tag

	err := db.Scopes(isDeletedRecord).First(&tag, id).Error
	if err != nil {
		return tag, err
	}

	return tag, nil
}

func UpdateTagByID(id int64, data dto.UpdateTagDTO) error {
	err := db.Model(models.Tag{}).Scopes(isDeletedRecord).Where("id = ?", id).Limit(1).Updates(
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
	var tag models.Tag
	db.Scopes(isDeletedRecord).Where("name = ?", name).First(&tag)

	return tag.ID != 0
}

func TagExistsByID(id int64) bool {
	var tag models.Tag
	db.Scopes(isDeletedRecord).First(&tag, id)

	return tag.ID != 0
}

func CreateTag(data dto.CreateTagDTO) (models.Tag, error) {
	var tag models.Tag

	if exists := TagExistsByName(data.Name); exists {
		return tag, myerror.ErrorTagExists
	}

	id, err := internal.GenSnowflakeID()
	if err != nil {
		return tag, err
	}

	tag = models.Tag{
		ID:          id,
		Name:        data.Name,
		FriendlyUrl: data.FriendlyUrl,
	}
	err = db.Model(&models.Tag{}).Create(&tag).Error
	if err != nil {
		return tag, err
	}

	return tag, nil
}
