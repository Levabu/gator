package state

import (
	"github.com/levabu/gator/internal/config"
	"github.com/levabu/gator/internal/database"
)

type State struct {
	Config *config.Config
	DB *database.Queries
}