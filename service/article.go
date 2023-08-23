package service

import (
	"api.aifuxi.cool/dao/mysql"
	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
)

func GetArticles(data *dto.GetArticlesDTO) (*[]models.Article, int64, error) {
	return mysql.GetArticles(data)
}

func CreateArticle(data *dto.CreateArticleDTO) (*models.Article, error) {
	return mysql.CreateArticle(data)
}

func GetArticleByID(id int64) (*models.Article, error) {
	return mysql.GetArticleByID(id)
}

func UpdateArticleByID(id int64, data *dto.UpdateArticleDTO) error {
	if !mysql.ArticleExistsByID(id) {
		return myerror.ErrorArticleNotFound
	}

	return mysql.UpdateArticleByID(id, data)
}

func DeleteArticleByID(id int64) error {
	return mysql.DeleteArticleByID(id)
}
