package db

func (q *Queries) CreateArticleTag(articleID int64, tagID int64) error {
	articleTag := ArticleTag{
		ArticleID: articleID,
		TagID:     tagID,
	}
	err := q.db.Model(ArticleTag{}).Create(&articleTag).Error
	if err != nil {
		return err
	}

	return nil
}

func (q *Queries) BatchCreateArticleTag(articleID int64, tagIDs []int64) error {
	if len(tagIDs) == 0 {
		return nil
	}

	var articleTagList []ArticleTag
	for _, tagID := range tagIDs {
		articleTagList = append(articleTagList, ArticleTag{
			ArticleID: articleID,
			TagID:     tagID,
		})
	}

	err := q.db.Create(&articleTagList).Error
	if err != nil {
		return err
	}

	return nil
}

func (q *Queries) DeleteArticleTag(articleID int64, tagID int64) error {
	err := q.db.Delete(ArticleTag{}, ArticleTag{
		ArticleID: articleID,
		TagID:     tagID,
	}).Limit(1).Error
	if err != nil {
		return err
	}

	return nil
}

func (q *Queries) DeleteArticleTagByArticleID(articleID int64) error {
	err := q.db.Delete(ArticleTag{}, ArticleTag{
		ArticleID: articleID,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (q *Queries) DeleteArticleTagByTagID(tagID int64) error {
	err := q.db.Delete(ArticleTag{}, ArticleTag{
		TagID: tagID,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (q *Queries) GetArticleTagIDs(articleID int64) ([]int64, error) {
	var tagIDs []int64
	var articleTagList []ArticleTag
	err := q.db.Model(ArticleTag{}).Find(&articleTagList, ArticleTag{
		ArticleID: articleID,
	}).Error
	if err != nil {
		return nil, err
	}

	for _, v := range articleTagList {
		tagIDs = append(tagIDs, v.TagID)
	}

	return tagIDs, nil
}

func (q *Queries) GetArticleIDsByTagID(tagID int64) ([]int64, error) {
	var articleIDs []int64
	var articleTagList []ArticleTag
	err := q.db.Model(ArticleTag{}).Find(&articleTagList, ArticleTag{
		TagID: tagID,
	}).Error
	if err != nil {
		return nil, err
	}

	for _, v := range articleTagList {
		articleIDs = append(articleIDs, v.TagID)
	}

	return articleIDs, nil
}
