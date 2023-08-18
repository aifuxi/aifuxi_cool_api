package service

import (
	"github.com/aifuxi/aifuxi_cool_api/dao/mysql"
	"github.com/aifuxi/aifuxi_cool_api/dto"
	"github.com/aifuxi/aifuxi_cool_api/models"
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
		return mysql.ErrorTagNotFound
	}

	return mysql.UpdateTagByID(id, data)
}

func DeleteTagByID(id int64) error {
	return mysql.DeleteTagByID(id)
}
