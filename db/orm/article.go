package db

import (
	"api.aifuxi.cool/util"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

var (
	ErrArticleExist    = errors.New("文章已存在")
	ErrArticleNotFound = errors.New("文章不存在")
)

type ExistArticleParams struct {
	ID          int64
	Title       string
	FriendlyURL string
}

func (q *Queries) ExistArticle(arg ExistArticleParams) (bool, error) {
	var article Article
	cond := Article{
		ID:          arg.ID,
		Title:       arg.Title,
		FriendlyURL: arg.FriendlyURL,
	}

	err := q.db.Scopes(isDeleted).First(&article, cond).Error
	if err != nil {
		// 如果 err 是 ErrRecordNotFound，只是记录没找到，不认为是出错了
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	if article.ID != 0 {
		return true, nil
	}

	return false, nil
}

type ListArticlesParams struct {
	Title       string
	FriendlyURL string

	Page     int
	PageSize int
	Order    string
	OrderBy  string
}

func (q *Queries) ListArticles(arg ListArticlesParams) ([]Article, int64, error) {
	var articles []Article
	var count int64

	queryDB := q.db.Model(Article{}).Scopes(isDeleted)

	if len(arg.FriendlyURL) > 0 {
		queryDB.Where("friendly_url LIKE ?", "%"+arg.FriendlyURL+"%")
	}

	if len(arg.Title) > 0 {
		queryDB.Where("title LIKE ?", "%"+arg.Title+"%")
	}

	queryDB = queryDB.Count(&count)

	order := fmt.Sprintf("%s %s", arg.OrderBy, arg.Order)
	err := queryDB.Order(order).Scopes(paginate(arg.Page, arg.PageSize)).Find(&articles).Error
	if err != nil {
		return nil, count, err
	}

	for i, article := range articles {
		var tagIDs []int64
		var tags []Tag
		tagIDs, err = q.GetArticleTagIDs(article.ID)
		if err != nil {
			err = fmt.Errorf("get article_tag ids error %d: %w", i, err)
		}

		tags, err = q.GetTagsByIDs(tagIDs)
		if err != nil {
			err = fmt.Errorf("get tags by ids error %d: %w", i, err)
		}

		if err != nil {
			continue
		}

		articles[i].Tags = tags
	}

	return articles, count, nil
}

type CreateArticleParams struct {
	Title       string
	Description string
	Cover       string
	Content     string
	FriendlyURL string
	IsTop       bool
	TopPriority int
	IsPublished bool
	TagIDs      []int64
}

func (q *Queries) CreateArticle(arg CreateArticleParams) (Article, error) {
	article := Article{
		Title:       arg.Title,
		Description: arg.Description,
		Cover:       arg.Cover,
		Content:     arg.Content,
		FriendlyURL: arg.FriendlyURL,
		IsTop:       arg.IsTop,
		TopPriority: arg.TopPriority,
		IsPublished: arg.IsPublished,
	}

	exitArticleArg := ExistArticleParams{Title: arg.Title}
	exist, err := q.ExistArticle(exitArticleArg)
	if err != nil {
		return Article{}, err
	}

	if exist {
		return Article{}, ErrArticleExist
	}

	err = q.db.Scopes(isDeleted).Create(&article).Error
	if err != nil {
		return Article{}, err
	}

	err = q.BatchCreateArticleTag(article.ID, arg.TagIDs)
	if err != nil {
		return article, err
	}

	var tags []Tag
	tags, err = q.GetTagsByIDs(arg.TagIDs)
	if err != nil {
		return article, err
	}

	article.Tags = tags

	return article, nil
}

func (q *Queries) GetArticleByID(id int64) (Article, error) {
	var article Article
	var tagIDs []int64
	var tags []Tag

	err := q.db.Scopes(isDeleted).First(&article, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Article{}, ErrArticleNotFound
		}

		return Article{}, err
	}

	tagIDs, err = q.GetArticleTagIDs(article.ID)
	if err != nil {
		return article, err
	}

	tags, err = q.GetTagsByIDs(tagIDs)
	if err != nil {
		return article, err
	}

	article.Tags = tags

	return article, nil
}

type UpdateArticleParams struct {
	Title       string
	Description string
	Cover       string
	Content     string
	FriendlyURL string
	IsTop       bool
	TopPriority int
	IsPublished bool
	TagIDs      []int64
}

func (q *Queries) UpdateArticle(id int64, arg UpdateArticleParams) error {
	article := Article{
		ID: id,
	}
	cond := Article{
		Title:       arg.Title,
		Description: arg.Description,
		Cover:       arg.Cover,
		Content:     arg.Content,
		FriendlyURL: arg.FriendlyURL,
		IsTop:       arg.IsTop,
		TopPriority: arg.TopPriority,
		IsPublished: arg.IsPublished,
	}

	exitArticleArg := ExistArticleParams{ID: id}
	exist, err := q.ExistArticle(exitArticleArg)
	if err != nil {
		return err
	}

	if !exist {
		return ErrArticleNotFound
	}

	err = q.db.Scopes(isDeleted).Model(&article).Updates(cond).Error
	if err != nil {
		return err
	}

	// 先找出文章下所有的tag id
	var articleTagIDs []int64
	articleTagIDs, err = q.GetArticleTagIDs(article.ID)
	if err != nil {
		return err
	}

	var tagErr error
	// 判断有没有取消关联tag
	for i, v := range articleTagIDs {
		// [1,2,3]  [3]
		if !util.FindInt64(arg.TagIDs, v) {
			err := q.DeleteArticleTag(id, v)
			if err != nil {
				tagErr = fmt.Errorf("delete article tag error %d: %w", i, err)
				continue
			}
		}
	}

	// 判断有没有新关联tag
	for i, v := range arg.TagIDs {
		if !util.FindInt64(articleTagIDs, v) {
			err := q.CreateArticleTag(id, v)
			if err != nil {
				tagErr = fmt.Errorf("delete article tag error %d: %w", i, err)
				continue
			}
		}
	}

	if tagErr != nil {
		return tagErr
	}

	return nil
}

func (q *Queries) DeleteArticleByID(id int64) error {
	exitArticleArg := ExistArticleParams{ID: id}
	exist, err := q.ExistArticle(exitArticleArg)
	if err != nil {
		return err
	}
	if !exist {
		return ErrArticleNotFound
	}

	now := time.Now()
	article := Article{
		ID: id,
	}
	cond := Article{DeletedAt: &now}

	err = q.db.Scopes(isDeleted).Model(&article).Updates(cond).Error
	if err != nil {
		return err
	}

	err = q.DeleteArticleTagByArticleID(article.ID)
	if err != nil {
		return err
	}

	return nil
}
