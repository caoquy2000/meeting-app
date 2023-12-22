package controller

import (
	"net/http"

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
