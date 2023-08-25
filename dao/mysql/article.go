package mysql

import (
	"fmt"
	"time"

	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/internal"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
)

func GetArticles(arg dto.GetArticlesDTO) ([]models.Article, int64, error) {
	var articles []models.Article
	var count int64
	var queryDB = db.Model(models.Article{}).Scopes(isDeletedRecord)

	if len(arg.Title) > 0 {
		queryDB.Where("title LIKE ?", "%"+arg.Title+"%")
	}

	if len(arg.FriendlyUrl) > 0 {
		queryDB.Where("friendly_url LIKE ?", "%"+arg.FriendlyUrl+"%")
	}

	queryDB = queryDB.Count(&count)

	order := fmt.Sprintf("%s %s", arg.OrderBy, arg.Order)
	err := queryDB.Order(order).Scopes(Paginate(arg.Page, arg.PageSize)).Find(&articles).Error
	if err != nil {
		return nil, count, err
	}

	return articles, count, nil
}

func GetArticleByID(id int64) (models.Article, error) {
	var article models.Article

	err := db.Scopes(isDeletedRecord).First(&article, id).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

func UpdateArticleByID(id int64, data dto.UpdateArticleDTO) error {
	err := db.Model(models.Article{}).Scopes(isDeletedRecord).Where("id = ?", id).Limit(1).Updates(
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
	var article models.Article

	db.Scopes(isDeletedRecord).Where("title = ?", title).First(&article)

	return article.ID != 0
}

func ArticleExistsByID(id int64) bool {
	var article models.Article

	db.Scopes(isDeletedRecord).First(&article, id)

	return article.ID != 0
}

func CreateArticle(data dto.CreateArticleDTO) (models.Article, error) {
	var article models.Article

	if exists := ArticleExistsByTitle(data.Title); exists {
		return article, myerror.ErrorArticleExists
	}

	id, err := internal.GenSnowflakeID()
	if err != nil {
		return article, err
	}

	article = models.Article{
		ID:          id,
		Title:       data.Title,
		Description: data.Description,
		Content:     data.Content,
		Cover:       data.Cover,
		IsTop:       data.IsTop,
		TopPriority: data.TopPriority,
		FriendlyUrl: data.FriendlyUrl,
	}

	err = db.Model(models.Article{}).Create(article).Error
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}
