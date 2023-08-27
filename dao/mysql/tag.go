package mysql

import (
	"fmt"
	"time"

	"api.aifuxi.cool/dto"
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

	for i, tag := range tags {
		articleCount, _ := GetTagArticleCount(tag.ID)
		tags[i].ArticleCount = articleCount
	}

	return tags, count, nil
}

func GetTagArticleCount(id int64) (int, error) {
	var articleTagIDs []models.ArticleTag

	err := db.Model(models.ArticleTag{}).Where("tag_id = ?", id).Find(&articleTagIDs).Error
	if err != nil {
		return 0, err
	}

	return len(articleTagIDs), nil
}

func GetTagByID(id int64) (models.Tag, error) {
	var tag models.Tag

	err := db.Scopes(isDeletedRecord).First(&tag, id).Error
	if err != nil {
		return tag, err
	}

	articleCount, err := GetTagArticleCount(tag.ID)
	if err != nil {
		return tag, err
	}

	tag.ArticleCount = articleCount

	return tag, nil
}

func UpdateTagByID(id int64, arg dto.UpdateTagDTO) error {
	err := db.Model(models.Tag{}).Scopes(isDeletedRecord).Where("id = ?", id).Limit(1).Updates(
		models.Tag{
			Name:        arg.Name,
			FriendlyUrl: arg.FriendlyUrl,
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

func CreateTag(arg dto.CreateTagDTO) (models.Tag, error) {
	var tag models.Tag

	if exists := TagExistsByName(arg.Name); exists {
		return tag, myerror.ErrorTagExists
	}

	tag = models.Tag{
		Name:        arg.Name,
		FriendlyUrl: arg.FriendlyUrl,
	}
	err := db.Model(&models.Tag{}).Create(&tag).Error
	if err != nil {
		return tag, err
	}

	return tag, nil
}
