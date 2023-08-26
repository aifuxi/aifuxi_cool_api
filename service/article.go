package service

import (
	"api.aifuxi.cool/dao/mysql"
	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
)

func GetArticles(arg dto.GetArticlesDTO) ([]models.Article, int64, error) {
	return mysql.GetArticles(arg)
}

func CreateArticle(arg dto.CreateArticleDTO) (models.Article, error) {
	return mysql.CreateArticle(arg)
}

func GetArticleByID(id int64) (models.Article, error) {
	return mysql.GetArticleByID(id)
}

func UpdateArticleByID(id int64, arg dto.UpdateArticleDTO) error {
	if !mysql.ArticleExistsByID(id) {
		return myerror.ErrorArticleNotFound
	}

	return mysql.UpdateArticleByID(id, arg)
}

func DeleteArticleByID(id int64) error {
	return mysql.DeleteArticleByID(id)
}
