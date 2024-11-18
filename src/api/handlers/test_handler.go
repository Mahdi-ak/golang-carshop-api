package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// for HeaderBinder2
type header struct {
	UserId    string
	BrowserId string
}

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response": "Test!",
	})
}

func (h *TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response": "Users!",
	})
}
func (h *TestHandler) UserById(c *gin.Context) {

	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{
		"response": "Users by id!",
		"id":       id,
	})
}
func (h *TestHandler) UserByUsername(c *gin.Context) {

	username := c.Params.ByName("username")
	c.JSON(http.StatusOK, gin.H{
		"response": "User by username!",
		"username": username,
	})
}
func (h *TestHandler) Accounts(c *gin.Context) {

	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"response": "Accounts!",
		"id":       id,
	})
}
func (h *TestHandler) AddUser(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"response": "Add user!",
		"id":       "",
	})
}

// first method
func (h *TestHandler) HeaderBinder1(c *gin.Context) {
	userId := c.GetHeader("UserId")
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"userId": userId,
	})
}

// scend method
func (h *TestHandler) HeaderBinder2(c *gin.Context) {

	header := header{}
	c.BindHeader(&header)

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"userId": header,
	})
}
