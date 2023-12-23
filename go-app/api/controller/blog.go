package controller

import (
	"net/http"
	"strconv"

	"github.com/caoquy2000/meeting-app/api/service"
	"github.com/caoquy2000/meeting-app/models"
	"github.com/caoquy2000/meeting-app/utils"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	service service.PostService
}

func NewPostController(s service.PostService) PostController {
	return PostController{
		service: s,
	}
}

func (p PostController) GetPosts(ctx *gin.Context) {
	var posts models.Post

	keyword := ctx.Query("keyword")

	data, total, err := p.service.FindAll(posts, keyword)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
		return
	}

	resArr := make([]map[string]interface{}, 0, 0)

	for _, n := range *data {
		res := n.ResponseMap()
		resArr = append(resArr, res)
	}

	ctx.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "Posts Result",
		Data: map[string]interface{}{
			"rows":       resArr,
			"total_rows": total,
		},
	})
}

func (p *PostController) AddPost(ctx *gin.Context) {
	var post models.Post
	ctx.ShouldBindJSON(&post)

	if post.Title == "" {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}

	if post.Body == "" {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}

	err := p.service.Save(post)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create post")
		return
	}

	utils.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Post")
}

func (p *PostController) GetPost(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Invalid Id")
		return
	}

	var post models.Post
	post.ID = id
	foundPost, err := p.service.Find(post)
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding Post")
		return
	}
	response := foundPost.ResponseMap()

	ctx.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "Result set of Post",
		Data:    &response,
	})
}

func (p *PostController) DeletePost(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Invalid Id")
		return
	}

	er := p.service.Delete(id)

	if er != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to delete Post")
		return
	}

	response := &utils.Response{
		Success: true,
		Message: "Delete Successfully",
	}

	ctx.JSON(http.StatusOK, response)
}

func (p *PostController) UpdatePost(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Invalid Id")
		return
	}
	var post models.Post
	post.ID = id

	postRecord, err := p.service.Find(post)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Post with given id not found")
		return
	}
	ctx.ShouldBindJSON(&postRecord)

	if postRecord.Title == "" {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}

	if postRecord.Body == "" {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}

	if err := p.service.Update(postRecord); err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to Update Post")
		return
	}

	response := postRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "Update Success",
		Data:    response,
	})
}
