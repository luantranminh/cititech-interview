package main

import (
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/luantranminh/team-management-app/config/database/pg"
	projectHTTPDeliver "github.com/luantranminh/team-management-app/project/delivery/http"
	p "github.com/luantranminh/team-management-app/project/repository"
	pu "github.com/luantranminh/team-management-app/project/usecase"
)

func main() {

	pgDB, closeDB := pg.New(`user=postgres dbname=team sslmode=disable password=postgres host=localhost port=5432`)
	// memberRepo := m.NewMemberRepository(pgDB)
	projectRepo := p.NewProjectRepository(pgDB)

	projectUsecase := pu.NewProjectUsecase(projectRepo)
	pjHandler := projectHTTPDeliver.NewProjectHTTPHandler(projectUsecase)

	defer closeDB()

	router := httprouter.New()
	router.GET("/projects/:id", pjHandler.GetByID)
	router.POST("/projects/", pjHandler.Create)
	log.Fatal(http.ListenAndServe(":8000", router))
}
