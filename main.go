package main

import (
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/luantranminh/team-management-app/config/database/pg"
	memberHTTPDeliver "github.com/luantranminh/team-management-app/member/delivery/http"
	mb "github.com/luantranminh/team-management-app/member/repository"
	mbu "github.com/luantranminh/team-management-app/member/usecase"
	projectHTTPDeliver "github.com/luantranminh/team-management-app/project/delivery/http"
	p "github.com/luantranminh/team-management-app/project/repository"
	pu "github.com/luantranminh/team-management-app/project/usecase"
)

func main() {

	pgDB, closeDB := pg.New(`user=postgres dbname=team sslmode=disable password=postgres host=localhost port=5432`)

	projectRepo := p.NewProjectRepository(pgDB)
	projectUsecase := pu.NewProjectUsecase(projectRepo)
	pjHandler := projectHTTPDeliver.NewProjectHTTPHandler(projectUsecase)

	memberRepo := mb.NewMemberRepository(pgDB)
	memberUsecase := mbu.NewMemberUsecase(memberRepo)
	mbHandler := memberHTTPDeliver.NewMemberHTTPHandler(memberUsecase)

	defer closeDB()

	router := httprouter.New()
	router.GET("/projects/:id", pjHandler.GetByID)
	router.POST("/projects/", pjHandler.Create)

	router.POST("/members/", mbHandler.Create)
	router.POST("/members/assignments/", mbHandler.AssignToProject)

	log.Fatal(http.ListenAndServe(":8000", router))
}
