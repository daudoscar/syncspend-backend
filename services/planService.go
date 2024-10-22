package services

import (
	"errors"
	"syncspend/dto"
	"syncspend/helpers"
	"syncspend/models"
	"syncspend/repositories"
)

type PlanService struct{}

func (s *PlanService) CreatePlan(data dto.CreatePlanDTO) (dto.PlanResponseDTO, error) {
	inviteCode, inviteCodeExpires := helpers.GenerateInviteCode()

	newPlan := &models.Plan{
		ID_Owner:          data.ID_Owner,
		Title:             data.Title,
		Description:       data.Description,
		InviteCode:        inviteCode,
		InviteCodeExpires: inviteCodeExpires,
	}

	if err := repositories.InsertPlan(newPlan); err != nil {
		return dto.PlanResponseDTO{}, errors.New("failed to create plan")
	}

	response := &dto.PlanResponseDTO{
		ID:                newPlan.ID,
		ID_Owner:          newPlan.ID_Owner,
		Title:             newPlan.Title,
		Description:       newPlan.Description,
		InviteCode:        newPlan.InviteCode,
		InviteCodeExpires: newPlan.InviteCodeExpires,
		CreatedAt:         newPlan.CreatedAt,
		UpdatedAt:         newPlan.UpdatedAt,
	}

	return *response, nil
}
