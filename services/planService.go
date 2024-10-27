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

func (s *PlanService) UpdatePlan(planID uint64, data dto.UpdatePlanDTO) (*dto.PlanResponseDTO, error) {
	existingPlan, err := repositories.GetPlanByID(planID)
	if err != nil {
		return nil, errors.New("Plan not found")
	}

	if existingPlan.ID_Owner != data.ID_Owner {
		return nil, errors.New("User is not the owner of the plan")
	}

	existingPlan.Title = data.Title
	existingPlan.Description = data.Description

	if err := repositories.UpdatePlan(existingPlan); err != nil {
		return &dto.PlanResponseDTO{}, errors.New("failed to update plan")
	}

	response := &dto.PlanResponseDTO{
		ID:                existingPlan.ID,
		ID_Owner:          existingPlan.ID_Owner,
		Title:             existingPlan.Title,
		Description:       existingPlan.Description,
		InviteCode:        existingPlan.InviteCode,
		InviteCodeExpires: existingPlan.InviteCodeExpires,
		CreatedAt:         existingPlan.CreatedAt,
		UpdatedAt:         existingPlan.UpdatedAt,
	}

	return response, nil
}
