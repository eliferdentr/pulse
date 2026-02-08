package api

import (
	"net/http"

	"github.com/eliferdentr/pulse/internal/jobs"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, manager *jobs.Manager) {
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
				"job":     job,
				"status": "job found",
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "job not found",
			})
		}

	})

}
