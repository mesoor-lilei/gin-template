package user

import (
	"github.com/gin-gonic/gin"
	"gin-template/constant"
	"gin-template/model"
	"log"
	"net/http"
	"strconv"
)

func Save(c *gin.Context) {
	var entity model.User
	err := c.BindJSON(&entity)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, err.Error())
		return
	}
	if entity.Id == 0 {
		constant.Db.Create(&entity)
		log.Println("已创建", entity)
		c.JSON(http.StatusOK, "已创建")
		return
	}
	constant.Db.Model(&entity).Updates(map[string]interface{}{
		"name": entity.Name,
		"age":  entity.Age,
	})
	log.Println("已修改", entity)
	c.JSON(http.StatusOK, "已修改")
}

func GetById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, err.Error())
		return
	}
	entity := new(model.User)
	constant.Db.First(&entity, id)
	c.JSON(http.StatusOK, entity)
}
