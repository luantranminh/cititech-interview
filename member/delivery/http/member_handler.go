package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luantranminh/team-management-app/member"
	"github.com/luantranminh/team-management-app/models"
)

// MemberHandler .
type MemberHandler struct {
	memberUsecase member.Usecase
}

// CreateMemberRequest .
type CreateMemberRequest struct {
	Member struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"member"`
}

// CreateMemberResponse .
type CreateMemberResponse struct {
	Member models.Member `json:"member"`
}

// Create create new member by name and phone given
func (m *MemberHandler) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		req = CreateMemberRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
	)

	if err != nil {
		models.HandleError(err, models.ErrorInvalidMember, http.StatusBadRequest, w)
		return
	}

	mem, err := m.memberUsecase.Create(context.Background(), req.Member.Name, req.Member.Phone)
	if err != nil {
		models.HandleError(err, err, http.StatusBadRequest, w)
		return
	}

	member, err := json.Marshal(mem)
	if err != nil {
		models.HandleError(err, models.ErrorInvalidMember, http.StatusBadRequest, w)
		return
	}

	w.Write(member)
}

// AssignToProjectRequest .
type AssignToProjectRequest struct {
	Assignment struct {
		MemberID  models.UUID `json:"member_id"`
		ProjectID models.UUID `json:"project_id"`
	} `json:"assignment"`
}

// AssignToProject assign a member to a project
func (m *MemberHandler) AssignToProject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		req = AssignToProjectRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
	)

	if err != nil {
		models.HandleError(err, models.ErrorInvalidID, http.StatusBadRequest, w)
		return
	}

	err = m.memberUsecase.AssignToProject(context.Background(), req.Assignment.MemberID, req.Assignment.ProjectID)
	if err != nil {
		models.HandleError(err, models.ErrorInvalidAssignment, http.StatusBadRequest, w)
		return
	}

	w.Write([]byte(`{"success": true}`))
}

// NewMemberHTTPHandler .
func NewMemberHTTPHandler(mb member.Usecase) MemberHandler {
	handler := MemberHandler{memberUsecase: mb}
	return handler
}
