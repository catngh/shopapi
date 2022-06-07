package controllers

import (
	"github.com/BerIincat/shopapi/database"
)

type Handler struct {
	db *database.DbControllers
}

func NewHandler(db *database.DbControllers) *Handler {
	c := &Handler{db: db}
	return c
}
