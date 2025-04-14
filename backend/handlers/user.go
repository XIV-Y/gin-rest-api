package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/XIV-Y/gin-rest-api/db"
	"github.com/XIV-Y/gin-rest-api/models"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Errors []Error     `json:"errors,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
}

type Error struct {
	Status string `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func GetUsers(c *gin.Context) {
	var users []models.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Errors: []Error{{
				Status: strconv.Itoa(http.StatusInternalServerError),
				Title:  "Database Error",
				Detail: result.Error.Error(),
			}},
		})
		return
	}

	var usersResponse []models.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, user.ToResponse())
	}

	c.JSON(http.StatusOK, Response{
		Data: usersResponse,
		Meta: map[string]interface{}{
			"total": len(usersResponse),
		},
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	result := db.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, Response{
			Errors: []Error{{
				Status: strconv.Itoa(http.StatusNotFound),
				Title:  "Resource Not Found",
				Detail: "User with specified ID not found",
			}},
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Data: user.ToResponse(),
	})
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Errors: []Error{{
				Status: strconv.Itoa(http.StatusBadRequest),
				Title:  "Validation Error",
				Detail: err.Error(),
			}},
		})
		return
	}

	result := db.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Errors: []Error{{
				Status: strconv.Itoa(http.StatusInternalServerError),
				Title:  "Database Error",
				Detail: result.Error.Error(),
			}},
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Data: user.ToResponse(),
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	result := db.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, Response{
			Errors: []Error{{
				Status: strconv.Itoa(http.StatusNotFound),
				Title:  "Resource Not Found",
				Detail: "User with specified ID not found",
			}},
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Errors: []Error{{
				Status: strconv.Itoa(http.StatusBadRequest),
				Title:  "Validation Error",
				Detail: err.Error(),
			}},
		})
		return
	}

	db.DB.Save(&user)

	c.JSON(http.StatusOK, Response{
		Data: user.ToResponse(),
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	result := db.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, Response{
			Errors: []Error{{
				Status: strconv.Itoa(http.StatusNotFound),
				Title:  "Resource Not Found",
				Detail: "User with specified ID not found",
			}},
		})
		return
	}

	db.DB.Delete(&user)

	c.JSON(http.StatusNoContent, nil)
}
