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

	planResponse := &dto.PlanResponseDTO{
		ID:                newPlan.ID,
		ID_Owner:          newPlan.ID_Owner,
		Title:             newPlan.Title,
		Description:       newPlan.Description,
		InviteCode:        newPlan.InviteCode,
		InviteCodeExpires: newPlan.InviteCodeExpires,
		CreatedAt:         newPlan.CreatedAt,
		UpdatedAt:         newPlan.UpdatedAt,
	}

	return *planResponse, nil
}

func (s *PlanService) UpdatePlan(data dto.PlanResponseDTO) (dto.PlanResponseDTO, error) {

	inviteCode, inviteCodeExpires := helpers.GenerateInviteCode()

	updatedPlan := &models.Plan{
		ID_Owner:          data.ID_Owner,
		Title:             data.Title,
		Description:       data.Description,
		InviteCode:        inviteCode,
		InviteCodeExpires: inviteCodeExpires,
	}

	if err := repositories.UpdatePlan(updatedPlan); err != nil {
		return dto.PlanResponseDTO{}, errors.New("failed to update plan")
	}

	planResponse := dto.PlanResponseDTO{
		ID:                updatedPlan.ID,
		ID_Owner:          updatedPlan.ID_Owner,
		Title:             updatedPlan.Title,
		Description:       updatedPlan.Description,
		InviteCode:        updatedPlan.InviteCode,
		InviteCodeExpires: updatedPlan.InviteCodeExpires,
		CreatedAt:         updatedPlan.CreatedAt,
		UpdatedAt:         updatedPlan.UpdatedAt,
	}

	return planResponse, nil
}

func (s *PlanService) DeletePlan(data dto.GetPlanDTO) error {
	plan := &models.Plan{ID: data.ID}

	if err := repositories.DeletePlanByID(plan); err != nil {
		return errors.New("failed to delete plan")
	}

	return nil
}
