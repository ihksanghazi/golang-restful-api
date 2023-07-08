package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(writer http.ResponseWriter,request *http.Request, params httprouter.Param)
	Update(writer http.ResponseWriter,request *http.Request, params httprouter.Param)
	Delete(writer http.ResponseWriter,request *http.Request, params httprouter.Param)
	FindById(writer http.ResponseWriter,request *http.Request, params httprouter.Param)
	FindAll(writer http.ResponseWriter,request *http.Request, params httprouter.Param)
}