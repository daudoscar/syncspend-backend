package services

import (
	"errors"
	"syncspend/dto"
	"syncspend/helpers"
	"syncspend/models"
	"syncspend/repositories"
	"time"
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

	if existingPlan.ID_Owner != userID {
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

	if existingPlan.ID_Owner != userID {
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

func (s *PlanService) JoinPlan(data *dto.JoinPlanDTO) error {
	plan, err := repositories.GetPlanByInviteCode(data.InviteCode)
	if err != nil {
		return errors.New("invalid invite code or plan not found")
	}
	existingMember, err := repositories.GetMemberByPlanAndUser(plan.ID, data.UserID)
	if err != nil {
		return errors.New("unable to verify membership status at this time")
	}

	if existingMember != nil && !existingMember.IsDeleted {
		return errors.New("user is already a member of this plan")
	}

	if existingMember != nil && existingMember.IsDeleted {
		return repositories.ReactivateMember(existingMember.ID)
	}

	planMember := &models.PlanMember{
		ID_Plan:  plan.ID,
		ID_User:  data.UserID,
		IsAdmin:  false,
		JoinedAt: time.Now(),
	}

	return repositories.InsertMember(planMember)
}

func (s *PlanService) LeavePlan(data *dto.LeavePlanDTO) error {
	existingMember, err := repositories.GetMemberByPlanAndUser(data.PlanID, data.UserID)
	if err != nil {
		return errors.New("unable to verify membership status at this time")
	}
	if existingMember == nil {
		return errors.New("user is not a member of this plan")
	}

	return repositories.LeavePlan(existingMember.ID)
}

func (s *PlanService) PromoteMemberPlan(PlanID uint64, data dto.PromoteMemberDTO) error {
	PlanCredentials, err := repositories.GetPlanMemberByID(PlanID)
	if err != nil {
		return errors.New("plan not found")
	}

	if PlanCredentials.IsAdmin {
		return errors.New("Targeted user is already an admin")
	}
	PlanCredentials.IsAdmin = true

	PlanInfo, err := repositories.GetPlanByID(PlanID)
	if err != nil {
		return err
	}
	if PlanInfo.ID_Owner != data.OwnerID {
		return errors.New("user does not have access to promote")
	}

	existingMember, err := repositories.GetMemberByPlanAndUser(PlanID, data.UserID)
	if err != nil {
		return errors.New("unable to verify membership status at this time")
	}
	if existingMember == nil {
		return errors.New("user is not a member of this plan")
	}

	err = repositories.UpdatePlanMember(PlanCredentials)
	if err != nil {
		return err
	}

	return nil
}

func (s *PlanService) DemoteAdminPlan(PlanID uint64, data dto.PromoteMemberDTO) error {
	PlanCredentials, err := repositories.GetPlanMemberByID(PlanID)
	if err != nil {
		return errors.New("plan not found")
	}

	if !PlanCredentials.IsAdmin {
		return errors.New("Targeted user is not an admin")
	}
	PlanCredentials.IsAdmin = false

	PlanInfo, err := repositories.GetPlanByID(PlanID)
	if err != nil {
		return err
	}
	if PlanInfo.ID_Owner != data.OwnerID {
		return errors.New("user does not have access to promote")
	}

	existingMember, err := repositories.GetMemberByPlanAndUser(PlanID, data.UserID)
	if err != nil {
		return errors.New("unable to verify membership status at this time")
	}
	if existingMember == nil {
		return errors.New("user is not a member of this plan")
	}

	err = repositories.UpdatePlanMember(PlanCredentials)
	if err != nil {
		return err
	}

	return nil
}
