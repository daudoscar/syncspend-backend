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
