package repositories

import (
	"syncspend/config"
	"syncspend/models"
)

type PlanRepository struct{}

func GetPlanByID(id uint64) (*models.Plan, error) {
	var plan models.Plan
	if err := config.DB.First(&plan, id).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

func InsertPlan(plan *models.Plan) error {
	if err := config.DB.Create(&plan).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePlan(plan *models.Plan) error {
	if err := config.DB.Save(plan).Error; err != nil {
		return err
	}
	return nil
}
