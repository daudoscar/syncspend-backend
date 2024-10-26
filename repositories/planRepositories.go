package repositories

import (
	"syncspend/config"
	"syncspend/models"
)

type PlanRepository struct{}

func InsertPlan(plan *models.Plan) error {
	if err := config.DB.Create(&plan).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePlan(plan *models.Plan) error {
	if err := config.DB.Save(&plan).Error; err != nil {
		return err
	}
	return nil
}

func DeletePlanByID(plan *models.Plan) error {
	if err := config.DB.Delete(&plan).Error; err != nil {
		return err
	}
	return nil
}

func GetPlan(planID uint64) *models.Plan {
	var plan models.Plan

	if err := config.DB.First(&plan, planID).Error; err != nil {
		return nil
	}

	return &plan
}

func FindUserPlans(userID uint64) ([]models.Plan, error) {
	var plans []models.Plan

	if err := config.DB.Where("id_owner = ?", userID).Find(&plans).Error; err != nil {
		return nil, err
	}

	return plans, nil
}
