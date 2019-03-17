package article_service

import (
	"Gin/learnGin/golangDemo/models"
	"Gin/learnGin/golangDemo/pkg/gredis"
	"Gin/learnGin/golangDemo/service/cache_service"
	"encoding/json"
	"github.com/gpmgo/gopm/modules/log"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int

}

func (a *Article) Add() error{
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_by":      a.CreatedBy,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	}

	if err := models.AddArticle(article); err != nil  {
		return err
	}

	return nil
}

func (a *Article) Get() (*models.Article, error){
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		data , err := gredis.Get(key)
		if err != nil {
			log.Info("", err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}
	gredis.Set(key, article, 3600)
	return article, nil
}

func (a *Article) ExistByID() (bool, error) {
	return false, nil
}

func (a *Article) Edit() error {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"modifiedBy":      a.ModifiedBy,
		"content":         a.Content,
		"created_by":      a.CreatedBy,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	}

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		gredis.Set(key, article, 3600)
	}

	return models.EditArticle(a.ID, article)
}

func (a *Article) Delete() error{
	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		gredis.Delete(key)
	}

	return models.DeleteArticle(a.ID)
}