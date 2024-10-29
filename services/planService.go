package services

import (
	"errors"
	"strconv"
	"syncspend/dto"
	"syncspend/helpers"
	"syncspend/models"
	"syncspend/repositories"
)

type PlanService struct{}

func (s *PlanService) GetPlanByID(planID uint64) (*dto.PlanResponseDTO, error) {
	plan, err := repositories.GetPlanByID(planID)
	if err != nil {
		return nil, errors.New("plan not found")
	}

	response := &dto.PlanResponseDTO{
		ID:                plan.ID,
		ID_Owner:          plan.ID_Owner,
		Title:             plan.Title,
		Description:       plan.Description,
		InviteCode:        plan.InviteCode,
		InviteCodeExpires: plan.InviteCodeExpires,
		CreatedAt:         plan.CreatedAt,
		UpdatedAt:         plan.UpdatedAt,
	}

	return response, nil
}

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
		return nil, errors.New("plan not found")
	}

	if existingPlan.ID_Owner != data.ID_Owner {
		return nil, errors.New("user is not the owner of the plan")
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

func (s *PlanService) DeletePlan(planID, userID uint64) error {
	existingPlan, err := repositories.GetPlanByID(planID)
	if err != nil {
		return errors.New("plan not found")
	}

	if existingPlan.ID_Owner != strconv.FormatUint(userID, 10) {
		return errors.New("user is not the owner of the plan")
	}

	existingPlan.IsDeleted = true

	if err := repositories.UpdatePlan(existingPlan); err != nil {
		return errors.New("failed to delete plan")
	}

	return nil
}

func (s *PlanService) RecoverPlan(planID, userID uint64) error {
	existingPlan, err := repositories.GetPlanByID(planID)
	if err != nil {
		return errors.New("plan not found")
	}

	if existingPlan.ID_Owner != strconv.FormatUint(userID, 10) {
		return errors.New("user is not the owner of the plan")
	}

	if !existingPlan.IsDeleted {
		return errors.New("plan is active")
	}

	existingPlan.IsDeleted = false

	if err := repositories.UpdatePlan(existingPlan); err != nil {
		return errors.New("failed to delete plan")
	}

	return nil
}
