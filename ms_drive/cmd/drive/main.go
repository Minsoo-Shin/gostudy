package main

import (
	"github.com/Minsoo-Shin/ms_drive/internal/drive/api"
	"github.com/Minsoo-Shin/ms_drive/internal/drive/repository"
	"github.com/Minsoo-Shin/ms_drive/migration"
	"github.com/Minsoo-Shin/ms_drive/pkg/gorm"
	"github.com/Minsoo-Shin/ms_drive/pkg/router/gin"
	"github.com/Minsoo-Shin/ms_drive/pkg/storage_cloud"
)

func main() {
	r := gin.New()
	db := gorm.New()
	sc := storage_cloud.New()
	migration.NewMigration(db)
	repo := repository.New(db)
	api.New(r, sc, repo)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
