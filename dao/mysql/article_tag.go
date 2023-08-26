package mysql

import (
	"errors"
	"fmt"
	"strconv"

	"api.aifuxi.cool/models"
)

func createArticleTagRecord(articleID int64, tagIDs []string) error {
	if articleID == 0 {
		return errors.New("invalid article id")
	}

	for _, tagID := range tagIDs {
		fmt.Printf("tagID: %v\n", tagID)
		if tagID != "" {
			// 手动把字符串类型的id值转换为int64
			id, _ := strconv.ParseInt(tagID, 10, 64)

			articleTag := models.ArticleTag{
				ArticleID: articleID,
				TagID:     id,
			}
			err := db.Model(models.ArticleTag{}).Create(&articleTag).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}
