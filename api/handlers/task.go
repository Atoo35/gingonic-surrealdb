package handlers

import (
	"fmt"
	"net/http"

	"github.com/Atoo35/gingonic-surrealdb/database"
	"github.com/Atoo35/gingonic-surrealdb/models"
	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

type TaskHandler struct {
	DB *surrealdb.DB
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := database.DB.Select("tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error while selecting tasks: %s", err.Error()),
		})
		return
	}

	tasksSlice := new([]models.Task)
	err = surrealdb.Unmarshal(tasks, &tasksSlice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error while unmarshalling tasks: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"tasks": tasksSlice,
	})
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	task := new(models.Task)
	if err := c.ShouldBindJSON(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error while binding json: %s", err.Error()),
		})
		return
	}

	_, err := database.DB.Create("tasks", task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error while creating task: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"task": task,
	})
}

func getSingletask(id string) (interface{}, error) {
	task, err := database.DB.Select(id)
	if err != nil {
		return nil, err
	}
	fmt.Println(task)
	return task, nil
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := getSingletask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error while getting task: %s", err.Error()),
		})
		return
	}

	taskModel := new(models.Task)
	err = surrealdb.Unmarshal(task, &taskModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error while unmarshalling task: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"task": taskModel,
	})
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	task := new(models.Task)
	if err := c.ShouldBindJSON(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error while binding json: %s", err.Error()),
		})
		return
	}

	_, err := getSingletask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error while getting task: %s", err.Error()),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error while unmarshalling task: %s", err.Error()),
		})
		return
	}

	_, err = database.DB.Update(id, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error while updating task: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"task": task,
	})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	_, err := database.DB.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error while deleting task: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": fmt.Sprintf("Task %s deleted", id),
	})
}
