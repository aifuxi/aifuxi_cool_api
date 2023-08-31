package db

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

var (
	ErrTagExist    = errors.New("标签已存在")
	ErrTagNotFound = errors.New("标签不存在")
)

type ExistTagParams struct {
	ID          int64
	Name        string
	FriendlyURL string
}

func (q *Queries) ExistTag(arg ExistTagParams) (bool, error) {
	var tag Tag
	cond := Tag{
		ID:          arg.ID,
		Name:        arg.Name,
		FriendlyURL: arg.FriendlyURL,
	}

	err := q.db.Scopes(isDeleted).First(&tag, cond).Error
	if err != nil {
		// 如果 err 是 ErrRecordNotFound，只是记录没找到，不认为是出错了
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	if tag.ID != 0 {
		return true, nil
	}

	return false, nil
}

type ListTagsParams struct {
	Name        string
	FriendlyURL string

	Page     int
	PageSize int
	Order    string
	OrderBy  string
}

func (q *Queries) ListTags(arg ListTagsParams) ([]Tag, int64, error) {
	var tags []Tag
	var count int64

	queryDB := q.db.Model(Tag{}).Scopes(isDeleted)

	if len(arg.FriendlyURL) > 0 {
		queryDB.Where("friendly_url LIKE ?", "%"+arg.FriendlyURL+"%")
	}

	if len(arg.Name) > 0 {
		queryDB.Where("email LIKE ?", "%"+arg.Name+"%")
	}

	queryDB = queryDB.Count(&count)

	order := fmt.Sprintf("%s %s", arg.OrderBy, arg.Order)
	err := queryDB.Order(order).Scopes(paginate(arg.Page, arg.PageSize)).Find(&tags).Error
	if err != nil {
		return nil, count, err
	}

	for i, v := range tags {
		articleIDs, err := q.GetArticleIDsByTagID(v.ID)
		if err != nil {
			continue
		}

		tags[i].ArticleCount = len(articleIDs)
	}

	return tags, count, nil
}

type CreateTagParams struct {
	Name        string
	FriendlyURL string
}

func (q *Queries) CreateTag(arg CreateTagParams) (Tag, error) {
	tag := Tag{
		Name:        arg.Name,
		FriendlyURL: arg.FriendlyURL,
	}

	exitTagArg := ExistTagParams{Name: arg.Name}
	exist, err := q.ExistTag(exitTagArg)
	if err != nil {
		return Tag{}, err
	}

	if exist {
		return Tag{}, ErrTagExist
	}

	err = q.db.Scopes(isDeleted).Create(&tag).Error
	if err != nil {
		return Tag{}, err
	}

	return tag, nil
}

func (q *Queries) GetTagByID(id int64) (Tag, error) {
	var tag Tag

	err := q.db.Scopes(isDeleted).First(&tag, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Tag{}, ErrTagNotFound
		}

		return Tag{}, err
	}

	articleIDs, err := q.GetArticleIDsByTagID(tag.ID)
	if err != nil {
		return tag, nil
	}

	tag.ArticleCount = len(articleIDs)

	return tag, nil
}

func (q *Queries) GetTagsByIDs(ids []int64) ([]Tag, error) {
	var tags []Tag

	err := q.db.Scopes(isDeleted).Find(&tags, ids).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

type UpdateTagParams struct {
	Name        string
	FriendlyURL string
}

func (q *Queries) UpdateTag(id int64, arg UpdateTagParams) error {
	tag := Tag{
		ID: id,
	}
	cond := Tag{
		Name:        arg.Name,
		FriendlyURL: arg.FriendlyURL,
	}

	exitTagArg := ExistTagParams{ID: id}
	exist, err := q.ExistTag(exitTagArg)
	if err != nil {
		return err
	}

	if !exist {
		return ErrTagNotFound
	}

	err = q.db.Scopes(isDeleted).Model(&tag).Updates(cond).Error
	if err != nil {
		return err
	}

	return nil
}

func (q *Queries) DeleteTagByID(id int64) error {
	exitTagArg := ExistTagParams{ID: id}
	exist, err := q.ExistTag(exitTagArg)
	if err != nil {
		return err
	}
	if !exist {
		return ErrTagNotFound
	}

	now := time.Now()
	tag := Tag{
		ID: id,
	}
	cond := Tag{DeletedAt: &now}

	err = q.db.Scopes(isDeleted).Model(&tag).Updates(cond).Error
	if err != nil {
		return err
	}

	return nil
}
