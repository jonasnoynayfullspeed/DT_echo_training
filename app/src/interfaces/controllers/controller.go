package controllers

import (
	"app/src/entities"
	"app/src/infrastructure/sqlhandler"
	"app/src/usecase"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Controller struct {
	Interactor usecase.Interactor
}

type HTMLContent struct {
	SiteTitle  string
	SiteHeader string
	Blog       entities.Blog
	Blogs      entities.Blogs
}

type JSONResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

/*
このファイルには外部からのリクエストで受け取ったデータをusecaseで使えるように変形したり、
内部からのデータを外部機能に向けて便利な形式に変換したりする
例)　外部からのデータをArticleエンティティに変換
*/

func NewController(sqlhandler *sqlhandler.SqlHandler) *Controller {
	return &Controller{
		Interactor: usecase.Interactor{
			Repository: usecase.Repository{
				DB: sqlhandler.DB,
			},
		},
	}
}

func (c Controller) Index(ctx echo.Context) error {
	blogs, err := c.Interactor.GetBlogList()
	if err != nil {
		log.Print(err)
		return ctx.Render(500, "index.html", nil)
	}

	return ctx.Render(http.StatusOK, "index.html", HTMLContent{
		SiteTitle:  "Go Blog - A Sample Echo App",
		SiteHeader: "Go Blog",
		Blogs:      blogs,
	})
}

func (c Controller) ShowBlogListPage(ctx echo.Context) error {
	blogs, err := c.Interactor.GetBlogList()
	if err != nil {
		return ctx.Render(500, "blog_list.html", nil)
	}

	return ctx.Render(http.StatusOK, "blog_list.html", HTMLContent{
		SiteTitle:  "Blogs - Go Blog",
		SiteHeader: "Blogs",
		Blogs:      blogs,
	})
}

func (c Controller) ShowBlogDetailsPage(ctx echo.Context) error {
	ID := ctx.Param("id")
	blog, err := c.Interactor.GetBlog(ID)
	if err != nil {
		log.Print(err)
		return ctx.Render(500, "blog_show.html", nil)
	}
	return ctx.Render(http.StatusOK, "blog_show.html", HTMLContent{
		SiteTitle: blog.Title,
		Blog:      blog,
	})
}

func (c Controller) EditBlogDetailsPage(ctx echo.Context) error {
	ID := ctx.Param("id")
	blog, err := c.Interactor.GetBlog(ID)
	if err != nil {
		log.Print(err)
		return ctx.Render(500, "blog_edit.html", nil)
	}
	return ctx.Render(http.StatusOK, "blog_edit.html", blog)
}

func (c Controller) ShowBlogCreatePage(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "blog_create.html", nil)
}

func (c Controller) AddBlogPost(ctx echo.Context) error {
	err := c.Interactor.AddBlogPost(ctx)
	return redirectToPage(ctx, "/blog/list", err)
}

func (c Controller) EditBlogPost(ctx echo.Context) error {
	ID := ctx.Param("id")
	err := c.Interactor.EditBlogPost(ID, ctx)
	return redirectToPage(ctx, "/blog/show/"+ID, err)
}

func (c Controller) DeleteBlogPost(ctx echo.Context) error {
	ID := ctx.Param("id")
	err := c.Interactor.DeleteBlogPost(ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, JSONResponse{Error: err.Error()})
	}

	return ctx.JSON(http.StatusOK, JSONResponse{Message: "Blog successfully deleted."})
}

func redirectToPage(c echo.Context, page string, err error) error {
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	} else {
		return c.Redirect(http.StatusFound, page)
	}
}