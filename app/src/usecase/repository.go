package usecase

import (
	"app/src/entities"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) FetchAllBlogs(page int, perPage int) (entities.Blogs, error) {
	var blogs = entities.Blogs{}
	offset := (page - 1) * perPage
	tx := r.DB.Order("created_at DESC").Offset(offset).Limit(perPage).Find(&blogs)
	return blogs, tx.Error
}

func (r *Repository) FetchBlog(ID string) (entities.Blog, error) {
	var blog = entities.Blog{}
	tx := r.DB.First(&blog, ID)
	return blog, tx.Error
}

func (r *Repository) CreateBlog(blog *entities.Blog) error {
	res := r.DB.Create(&blog)
	return errorHandler(res)
}

func (r *Repository) UpdateBlog(blog *entities.Blog) error {
	res := r.DB.Model(&blog).Updates(&blog)
	return errorHandler(res)
}

func (r *Repository) DeleteBlog(ID string) error {
	res := r.DB.Delete(&entities.Blog{}, ID)
	return errorHandler(res)
}

func errorHandler(tx *gorm.DB) error {
	if tx.Error == nil {
		return nil
	}

	return tx.Error
}
