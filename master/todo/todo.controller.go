package todo

import (

	"net/http"
	"strconv"

	// "strings"

	"github.com/gin-gonic/gin"
);

func (todo Todo) CreateTodo(c *gin.Context){
	// Bind Json
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status": "Bad Request",
			"message": err.Error(),
			"data": gin.H{},
		})
		return;
	}
	if len(todo.Title) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
			"message":"title cannot be null",
			"data":  gin.H{},
		})
        return
    }

	if todo.ActivityId == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
			"message":"activity_group_id cannot be null",
			"data":  gin.H{},
		})
        return
    }

	// Create New Todo
	res,err := todo.Create(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
			"message":"title cannot be null",
			"data":  gin.H{},
		})
        return
	}

	// Send Response
	c.JSON(http.StatusCreated,gin.H{
		"status": "Success",
		"message": "Success",
		"data": res,
	})	
}

func (todo Todo) GetOneTodo(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		id = 0
	}

	res,err :=todo.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Not Found",
			"message":"Todo with ID " + strconv.Itoa(id) + " Not Found",
		})
        return
	}

	c.JSON(http.StatusOK,gin.H{
		"status": "Success",
		"message": "Success",
		"data": res,
	})
}

func (todo Todo) GetAllTodo(c *gin.Context){
	activityIdQuery := c.Query("activity_group_id")

	activityId,err := strconv.Atoi(activityIdQuery)
	if err != nil {
		activityId = 0
	}
	todos := []Todo{}
	res,err := todo.GetAll(uint(activityId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "Success",
			"message":"Todo with ID ",
			"data":  todos,
		})
        return
	}

	c.JSON(http.StatusOK,gin.H{
		"status": "Success",
		"message": "Success",
		"data": res,
	})
}

func (todo Todo) DeleteTodo(c *gin.Context){
	// Get Id
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status": "Bad Request",
			"message": "Wrong format Id",
			"data": gin.H{},
		})
		return;
	}


	// Check Id
	_,err = todo.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Not Found",
			"message":"Todo with ID " + strconv.Itoa(id) + " Not Found",
			"data":  gin.H{},
		})
        return
	}

	// Delete Id
	err = todo.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{
			"status":  "Not Found",
			"message": err.Error(),
			"data": gin.H{},
		})
		return;
	}

	// Send Response
	c.JSON(http.StatusOK,gin.H{
		"status": "Success",
		"message": "Success",
		"data": gin.H{},
	})
}

func (todo Todo) UpdateTodo(c *gin.Context){
	// Get Params
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status": "Bad Request",
			"message": "Wrong format Id",
			"data": gin.H{},
		})
		return;
	}

	// Bind Json
	err = c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status": "Bad Request",
			"message": err.Error(),
			"data": gin.H{},
		})
		return;
	}

	// Check Todo is Exist
	_,err = todo.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Not Found",
			"message":"Todo with ID " + strconv.Itoa(id) + " Not Found",
			"data":  gin.H{},
		})
        return
	}

	// Update Todo
	res,err := todo.Update(uint(id),&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status": "Bad Request",
			"message": err.Error(),
			"data": gin.H{},
		})
		return;
	}

	// Send Response
	c.JSON(http.StatusOK,gin.H{
		"status": "Success",
		"message": "Success",
		"data": res,
	})
}
