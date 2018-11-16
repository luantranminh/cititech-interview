package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/luantranminh/team-management-app/models"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/julienschmidt/httprouter"
	"github.com/luantranminh/team-management-app/project"
)

var validate = validator.New()

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
	Project struct {
		Name string `json:"name"`
	} `json:"project"`
}

// ProjectHandler .
type ProjectHandler struct {
	projectUsecase project.Usecase
}

// GetByID .
func (p *ProjectHandler) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	projectID := ps.ByName("id")

	projectUUID, err := models.UUIDFromString(projectID)
	if err != nil {
		models.HandleError(err, models.ErrorInvalidID, http.StatusBadRequest, w)
		return
	}

	project, members, err := p.projectUsecase.GetByID(context.Background(), projectUUID)
	if err != nil {
		models.HandleError(err, models.ErrorInvalidProject, http.StatusBadRequest, w)
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
		models.HandleError(err, models.ErrorInvalidProject, http.StatusBadRequest, w)
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

	if err := validate.Struct(req); err != nil {
		models.HandleError(err, models.ErrorInvalidProject, http.StatusBadRequest, w)
		return
	}

	if err != nil {
		models.HandleError(err, models.ErrorInvalidProject, http.StatusBadRequest, w)
		return
	}

	project, err := p.projectUsecase.Create(context.Background(), req.Project.Name)
	if err != nil {
		models.HandleError(err, models.ErrorInvalidProject, http.StatusBadRequest, w)
		return
	}

	info, err := json.Marshal(project)
	if err != nil {
		models.HandleError(err, models.ErrorInvalidProject, http.StatusBadRequest, w)
		return
	}

	w.Write(info)
}

// NewProjectHTTPHandler .
func NewProjectHTTPHandler(pju project.Usecase) ProjectHandler {
	handler := ProjectHandler{projectUsecase: pju}
	return handler
}
