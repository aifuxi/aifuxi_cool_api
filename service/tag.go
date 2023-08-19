package service

import (
	"api.aifuxi.cool/dao/mysql"
	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
)

func GetTags() (*[]models.Tag, error) {
	return mysql.GetTags()
}

func CreateTag(data *dto.CreateTagDTO) (*models.Tag, error) {
	return mysql.CreateTag(data)
}

func GetTagByID(id int64) (*models.Tag, error) {
	return mysql.GetTagByID(id)
}

func UpdateTagByID(id int64, data *dto.UpdateTagDTO) error {
	if !mysql.TagExistsByID(id) {
		return myerror.ErrorTagNotFound
	}

	return mysql.UpdateTagByID(id, data)
}

func DeleteTagByID(id int64) error {
	return mysql.DeleteTagByID(id)
}
