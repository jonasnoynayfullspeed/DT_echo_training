package usecase

import (
	"app/src/entities"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Interactor struct {
	Repository Repository
}

// アプリケーション固有のビジネスルール
// このファイルでは取得したデータを組み合わせたりしてユースケースを実現する

func (i Interactor) GetBlogList(page int, perPage int) (entities.Blogs, error) {
	return i.Repository.FetchAllBlogs(page, perPage)
}

func (i Interactor) GetBlog(ID string) (entities.Blog, error) {
	return i.Repository.FetchBlog(ID)
}

func (i Interactor) AddBlogPost(c echo.Context) error {
	var blog entities.Blog
	_ = c.Bind(&blog)
	return i.Repository.CreateBlog(&blog)
}

func (i Interactor) EditBlogPost(ID string, c echo.Context) error {
	idInt64, _ := strconv.ParseInt(ID, 10, 64)
	blog := entities.Blog{ID: idInt64}
	_ = c.Bind(&blog)
	return i.Repository.UpdateBlog(&blog)
}

func (i Interactor) DeleteBlogPost(ID string) error {
	return i.Repository.DeleteBlog(ID)
}
