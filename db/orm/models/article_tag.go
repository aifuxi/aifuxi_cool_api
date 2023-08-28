package models

type ArticleTag struct {
	ArticleID int64 `gorm:"column:article_id;type:bigint;" json:"article_id,string"`
	TagID     int64 `gorm:"column:tag_id;type:bigint;" json:"tag_id,string"`
}

// GORM 自定义表名
func (ArticleTag) TableName() string {
	return "article_tag"
}
