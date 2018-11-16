package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/luantranminh/team-management-app/models"

	"github.com/julienschmidt/httprouter"
	"github.com/luantranminh/team-management-app/project"
)

// ProjectInfo .
type ProjectInfo struct {
	ID      models.UUID     `json:"id"`
	Name    string          `json:"name"`
	Members []models.Member `json:"members"`
}

// GetProjectResponse .
type GetProjectResponse struct {
	Project ProjectInfo `json:"project"`
}

// CreateProjectRequest .
type CreateProjectRequest struct {
	Name string `json:"name"`
}

// ProjectHandler .
type ProjectHandler struct {
	projectUsecase project.Usecase
}

// GetByID .
func (p *ProjectHandler) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	projectID := ps.ByName("id")

	projectUUID, err := models.UUIDFromString(projectID)
	if err != nil {
		handleError(err, models.ErrorInvalidID, w)
		return
	}

	project, members, err := p.projectUsecase.GetByID(context.Background(), projectUUID)
	if err != nil {
		handleError(err, models.ErrorInvalidProject, w)
		return
	}

	info, err := json.Marshal(
		GetProjectResponse{
			Project: ProjectInfo{
				ID:      project.ID,
				Name:    project.Name,
				Members: members,
			},
		},
	)
	if err != nil {
		handleError(err, models.ErrorInvalidProject, w)
		return
	}

	w.Write(info)
}

// Create .
func (p *ProjectHandler) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		req = CreateProjectRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
	)

	if err != nil {
		handleError(err, models.ErrorInvalidProject, w)
		return
	}

	project, err := p.projectUsecase.Create(context.Background(), req.Name)
	if err != nil {
		handleError(err, models.ErrorInvalidProject, w)
		return
	}

	info, err := json.Marshal(project)
	if err != nil {
		handleError(err, models.ErrorInvalidProject, w)
		return
	}

	w.Write(info)
}

func handleError(err error, errorType error, w http.ResponseWriter) {
	log.Println(err)

	e := models.Error{
		Err:  errorType.Error(),
		Code: http.StatusBadRequest,
	}

	b, _ := json.Marshal(e)
	w.Write(b)
}

// NewProjectHTTPHandler .
func NewProjectHTTPHandler(pju project.Usecase) ProjectHandler {
	handler := ProjectHandler{projectUsecase: pju}

	return handler
}
