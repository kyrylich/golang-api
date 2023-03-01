package app

import (
	"github.com/gin-gonic/gin"
	"golangpet/internal/api/router"
	"golangpet/internal/config"
	factory "golangpet/internal/factory"
	"golangpet/internal/model"
	"golangpet/internal/translation"
	"gorm.io/gorm"
)

type AppInterface interface {
	Run() error
	Boot() error
	GetConfig() *config.Config
	GetDB() *gorm.DB
}

type App struct {
	router *gin.Engine
	config *config.Config
	db     *gorm.DB
}

func (a *App) GetConfig() *config.Config {
	return a.config
}

func (a *App) GetDB() *gorm.DB {
	return a.db
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

	a.config = cfg

	db, err := model.ConnectDatabase(&cfg.Database)
	if err != nil {
		return err
	}
	depFactory := factory.NewDependencyFactory(db)

	a.db = db
	a.router = router.CreateRouter(depFactory)

	if err = translation.RegisterValidationTranslations(); err != nil {
		return err
	}

	return nil
}
