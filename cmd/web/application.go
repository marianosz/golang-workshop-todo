package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"github.com/rfinochi/golang-workshop-todo/pkg/models"
	"github.com/rfinochi/golang-workshop-todo/pkg/models/google"
	"github.com/rfinochi/golang-workshop-todo/pkg/models/memory"
	"github.com/rfinochi/golang-workshop-todo/pkg/models/mongo"
)

type application struct {
	errorLog  *log.Logger
	infoLog   *log.Logger
	router    *gin.Engine
	itemModel *models.ItemModel
}

func (app *application) initModels() {
	app.itemModel = &models.ItemModel{}

	repositoryType := os.Getenv("TODO_REPOSITORY_TYPE")

	if repositoryType == "Mongo" {
		app.itemModel.Repository = &mongo.ItemRepository{}
	} else if repositoryType == "Google" {
		app.itemModel.Repository = &google.ItemRepository{}
	} else {
		app.itemModel.Repository = &memory.ItemRepository{}
	}
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}