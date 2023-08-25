package service

import (
	"api.aifuxi.cool/dao/mysql"
	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
)

func GetTags(arg dto.GetTagsDTO) ([]models.Tag, int64, error) {
	return mysql.GetTags(arg)
}

func CreateTag(arg dto.CreateTagDTO) (models.Tag, error) {
	return mysql.CreateTag(arg)
}

func GetTagByID(id int64) (models.Tag, error) {
	return mysql.GetTagByID(id)
}

func UpdateTagByID(id int64, arg dto.UpdateTagDTO) error {
	if !mysql.TagExistsByID(id) {
		return myerror.ErrorTagNotFound
	}

	return mysql.UpdateTagByID(id, arg)
}

func DeleteTagByID(id int64) error {
	return mysql.DeleteTagByID(id)
}
