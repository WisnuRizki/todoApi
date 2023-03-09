package activity

import (

	"net/http"
	"strconv"

	// "strings"
	"github.com/gin-gonic/gin"
);

func (activity Activity) CreateActivity(c *gin.Context){
	// Bind Json
	err := c.BindJSON(&activity)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status": "Bad Request",
			"message": err.Error(),
			"data": gin.H{},
		})
		return;
	}
	if len(activity.Title) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
			"message":"title cannot be null",
			"data":  gin.H{},
		})
        return
    }

	// Create New Activity
	res,err := activity.Create(&activity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
			"message":"title cannot be null",
			"data":  gin.H{},
		})
        return
	}

	// Send Response
	c.JSON(http.StatusCreated ,gin.H{
		"status": "Success",
		"message": "Success",
		"data": res,
	})	
}

func (activity Activity) GetAllActivity(c *gin.Context){
	res,err := activity.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
			"message":"Activity with ID ",
			"data":  gin.H{},
		})
        return
	}

	c.JSON(http.StatusOK,gin.H{
		"status": "Success",
		"message": "Success",
		"data": res,
	})
}

func (activity Activity) GetOneActivity(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		id = 0
	}

	res,err := activity.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Not Found",
			"message":"Activity with ID " + strconv.Itoa(id) + " Not Found",
			"data":  gin.H{},
		})
        return
	}

	if res == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Not Found",
			"message":"Activity with ID " + strconv.Itoa(id) + " Not Found",
			"data":  gin.H{},
		})
        return
	}

	c.JSON(http.StatusOK,gin.H{
		"status": "Success",
		"message": "Success",
		"data": res,
	})
}

func (activity Activity) UpdateActivity(c *gin.Context){
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
	err = c.BindJSON(&activity)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status": "Bad Request",
			"message": err.Error(),
			"data": gin.H{},
		})
		return;
	}
	if len(activity.Title) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
			"message":"title cannot be null",
			"data":  gin.H{},
		})
        return
    }

	// Check Activity is Exist
	_,err = activity.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Not Found",
			"message":"Activity with ID " + strconv.Itoa(id) + " Not Found",
			"data":  gin.H{},
		})
        return
	}

	// Update Activity
	res,err := activity.Update(uint(id),&activity)
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

func (acticity Activity) DeleteActivity(c *gin.Context){
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
	_,err = acticity.Get(uint(id))
	if err != nil {
		c.JSON(404, gin.H{
			"status": "Not Found",
			"message":"Activity with ID " + strconv.Itoa(id) + " Not Found",
			"data":  gin.H{},
		})
        return
	}

	// Delete Id
	err = acticity.Delete(uint(id))
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
		"data": gin.H{},
	})
}
