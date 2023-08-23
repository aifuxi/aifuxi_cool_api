package mysql

import (
	"fmt"
	"time"

	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/internal"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
)

func GetArticles(data *dto.GetArticlesDTO) (*[]models.Article, int64, error) {
	articles := new([]models.Article)
	var total int64
	var titleLike, friendlyUrlLike string

	order := fmt.Sprintf("%s %s", data.OrderBy, data.Order)
	if len(data.Title) > 0 {
		titleLike = "%" + data.Title + "%"
	}

	if len(data.FriendlyUrl) > 0 {
		friendlyUrlLike = "%" + data.FriendlyUrl + "%"
	}

	if len(data.Title) == 0 && len(data.FriendlyUrl) == 0 {
		titleLike = "%" + data.Title + "%"
		friendlyUrlLike = "%" + data.FriendlyUrl + "%"
	}

	err := db.Order(order).Where("deleted_at is null").Where(
		db.Where("title LIKE ?", titleLike).Or("friendly_url LIKE ?", friendlyUrlLike),
	).Offset((data.Page - 1) * data.PageSize).Limit(data.PageSize).Find(articles).Error
	if err != nil {
		return nil, total, err
	}

	err = db.Model(models.Article{}).Where("deleted_at is null").Where(
		db.Where("title LIKE ?", titleLike).Or("friendly_url LIKE ?", friendlyUrlLike),
	).Count(&total).Error
	if err != nil {
		return nil, total, err
	}

	return articles, total, nil
}

func GetArticleByID(id int64) (*models.Article, error) {
	article := new(models.Article)
	err := db.Where("deleted_at is null").First(&article, id).Error
	if err != nil {
		return nil, err
	}

	return article, nil
}

func UpdateArticleByID(id int64, data *dto.UpdateArticleDTO) error {
	var article = &models.Article{
		ID: id,
	}
	err := db.Model(article).Where("deleted_at is null").Limit(1).Updates(
		models.Article{
			Title:       data.Title,
			Description: data.Description,
			Content:     data.Content,
			Cover:       data.Cover,
			IsTop:       data.IsTop,
			TopPriority: data.TopPriority,
			FriendlyUrl: data.FriendlyUrl,
		}).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteArticleByID(id int64) error {
	if !ArticleExistsByID(id) {
		return myerror.ErrorArticleNotFound
	}

	err := db.Model(models.Article{}).Where("id = ?", id).Limit(1).Update("deleted_at", time.Now().Local().Format(time.DateTime)).Error
	if err != nil {
		return err
	}

	return nil
}

func ArticleExistsByTitle(title string) bool {
	article := new(models.Article)
	db.Where("deleted_at is null and title = ?", title).First(&article)
	return article.ID != 0
}

func ArticleExistsByID(id int64) bool {
	article := new(models.Article)
	db.Where("deleted_at is null").First(&article, id)
	return article.ID != 0
}

func CreateArticle(data *dto.CreateArticleDTO) (*models.Article, error) {
	if exists := ArticleExistsByTitle(data.Title); exists {
		return nil, myerror.ErrorArticleExists
	}

	id, err := internal.GenSnowflakeID()
	if err != nil {
		return nil, err
	}

	article := &models.Article{
		ID:          id,
		Title:       data.Title,
		Description: data.Description,
		Content:     data.Content,
		Cover:       data.Cover,
		IsTop:       data.IsTop,
		TopPriority: data.TopPriority,
		FriendlyUrl: data.FriendlyUrl,
	}
	err = db.Model(&models.Article{}).Create(article).Error
	if err != nil {
		return nil, err
	}

	return article, nil
}
