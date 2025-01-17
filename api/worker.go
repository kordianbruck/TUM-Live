package api

import (
	"TUM-Live/dao"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func configGinWorkerRouter(r gin.IRoutes) {
	r.GET("/api/worker/getJobs/:workerID", getJob)
}

func getJob(c *gin.Context) {
	_, err := dao.GetWorkerByID(context.Background(), c.Param("workerID"))
	if err != nil {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	job, err := dao.PickJob(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, &jobData{})
		return
	}
	c.JSON(http.StatusOK, &jobData{
		Id:   job.ID,
		Path: job.FilePath,
	})
}

type jobData struct {
	Id   uint   `json:"id"`
	Path string `json:"path"`
}
