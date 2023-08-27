package mysql

import (
	"fmt"
	"strconv"
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

	for index, article := range articles {
		tags, _ := GetTagsByArticleID(article.ID)
		fmt.Printf("当前文章id: %d, tags: %v\n", article.ID, tags)
		articles[index].Tags = tags
	}

	return articles, count, nil
}

func GetTagsByArticleID(id int64) ([]models.Tag, error) {
	var articleTagIDs []models.ArticleTag
	var tagIDs []int64
	var tags []models.Tag

	err := db.Model(models.ArticleTag{}).Where("article_id = ?", id).Find(&articleTagIDs).Error
	if err != nil {
		return nil, err
	}

	if len(articleTagIDs) == 0 {
		return tags, nil
	}

	for _, v := range articleTagIDs {
		tagIDs = append(tagIDs, v.TagID)
	}

	err = db.Model(models.Tag{}).Find(&tags, tagIDs).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func GetArticleByID(id int64) (models.Article, error) {
	var article models.Article

	err := db.Scopes(isDeletedRecord).First(&article, id).Error
	if err != nil {
		return article, err
	}

	tags, err := GetTagsByArticleID(article.ID)
	if err != nil {
		return article, err
	}

	article.Tags = tags

	return article, nil
}

func GetArticleByFriendlyUrl(friendlyUrl string) (models.Article, error) {
	var article models.Article

	err := db.Scopes(isDeletedRecord).First(&article, models.Article{FriendlyUrl: friendlyUrl}).Error
	if err != nil {
		return article, err
	}

	tags, err := GetTagsByArticleID(article.ID)
	if err != nil {
		return article, err
	}

	article.Tags = tags

	return article, nil
}

func UpdateArticleByID(id int64, arg dto.UpdateArticleDTO) error {
	err := db.Model(models.Article{}).Scopes(isDeletedRecord).Where("id = ?", id).Limit(1).Updates(
		models.Article{
			Title:       arg.Title,
			Description: arg.Description,
			Content:     arg.Content,
			Cover:       arg.Cover,
			IsTop:       arg.IsTop,
			TopPriority: arg.TopPriority,
			FriendlyUrl: arg.FriendlyUrl,
			IsPublished: arg.IsPublished,
		}).Error
	if err != nil {
		return err
	}

	// 先找出文章下所有的tag id
	var articleTagIDs []models.ArticleTag
	err = db.Model(models.ArticleTag{}).Where("article_id = ?", id).Find(&articleTagIDs).Error
	if err != nil {
		return err
	}
	var currentTagIDs []int64
	fmt.Printf("找到的是articleTagIDs: %v\n", articleTagIDs)
	for _, v := range articleTagIDs {
		currentTagIDs = append(currentTagIDs, v.TagID)
	}

	tagIDs := []int64{}

	for _, v := range arg.TagIDs {
		tagID, _ := strconv.ParseInt(v, 10, 64)
		tagIDs = append(tagIDs, tagID)
	}

	// 判断有没有取消关联tag
	for _, v := range currentTagIDs {
		// [1,2,3]  [3]
		if !internal.FindInt64(tagIDs, v) {
			deleteArticleTag(id, v)
		}
	}

	// 判断有没有新关联tag
	for _, v := range tagIDs {
		if !internal.FindInt64(currentTagIDs, v) {
			createArticleTag(id, v)
		}
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

	err = deleteArticleTagByArticleID(id)
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

func CreateArticle(arg dto.CreateArticleDTO) (models.Article, error) {
	var article models.Article

	if exists := ArticleExistsByTitle(arg.Title); exists {
		return article, myerror.ErrorArticleExists
	}

	article = models.Article{
		Title:       arg.Title,
		Description: arg.Description,
		Content:     arg.Content,
		Cover:       arg.Cover,
		IsTop:       arg.IsTop,
		TopPriority: arg.TopPriority,
		FriendlyUrl: arg.FriendlyUrl,
		IsPublished: arg.IsPublished,
	}

	err := db.Model(models.Article{}).Create(&article).Error
	if err != nil {
		return models.Article{}, err
	}

	fmt.Printf("arg.TagIDs: %v\n", arg.TagIDs)
	// 把文章和标签关联起来
	err = createArticleTagRecord(article.ID, arg.TagIDs)
	if err != nil {
		return models.Article{}, err
	}

	var tags []models.Tag
	for _, v := range arg.TagIDs {
		id, _ := strconv.ParseInt(v, 10, 64)
		tag, _ := GetTagByID(id)
		tags = append(tags, tag)
	}

	article.Tags = tags
	return article, nil
}
