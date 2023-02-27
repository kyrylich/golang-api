package app

import (
	"github.com/gin-gonic/gin"
	"golangpet/internal/api/router"
	"golangpet/internal/config"
	"golangpet/internal/models"
	"golangpet/internal/translation"
)

type AppInterface interface {
	Run() error
	Boot() error
}

type App struct {
	router *gin.Engine
	Config *config.Config
}

func (a *App) Run() error {
	if err := a.Boot(); err != nil {
		return err
	}

	return a.router.Run()
}

func (a *App) Boot() error {
	cfg, err := config.InitConfiguration()
	if err != nil {
		return err
	}

	a.Config = cfg

	if err := models.ConnectDatabase(&cfg.Database); err != nil {
		return err
	}

	a.router = router.CreateRouter()

	if err := translation.RegisterValidationTranslations(); err != nil {
		return err
	}

	return nil
}
