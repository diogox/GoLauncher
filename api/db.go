package api

import "github.com/diogox/GoLauncher/api/models"

type DB interface {
	AddApp(exec string, name string, descr string, iconName string) error
	FindAppByID(exec string) (string, error)
	FindAppByName(name string) ([]models.AppInfo, error)
}