package api

import (
	"net/http"

	"github.com/eliferdentr/pulse/internal/jobs"
	"github.com/gin-gonic/gin"
)

func NewRouter(manager *jobs.Manager) *gin.Engine{
	r := gin.Default()
	r.Use(RequestLogger())
	r.POST("/jobs", func(c *gin.Context) {
		var req jobs.JobRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := manager.SubmitJob(req)
		c.JSON(http.StatusAccepted, gin.H{
			"id":     id,
			"status": "queued",
		})

	})

	r.GET("/jobs/:id", func(c *gin.Context) {
		idString := c.Param("id")
		job, ok := manager.Store.Get(idString)
		if ok {
			c.JSON(http.StatusOK, gin.H{
				"job": job,
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "job not found",
			})
		}

	})
	r.POST("/jobs/:id/cancel", func(c *gin.Context) {
		id := c.Param("id")

		ok := manager.CancelJob(id)
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "job not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "cancel requested"})
	})
	return r

}
