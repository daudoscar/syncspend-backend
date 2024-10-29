package repositories

import (
	"syncspend/config"
	"syncspend/models"
	"time"

	"gorm.io/gorm"
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

func GetPlanByInviteCode(inviteCode string) (*models.Plan, error) {
	var plan models.Plan
	if err := config.DB.Where("invite_code = ?", inviteCode).First(&plan).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

func GetMemberByPlanAndUser(planID, userID uint64) (*models.PlanMember, error) {
	var planMember models.PlanMember
	err := config.DB.Where("id_plan = ? AND id_user = ?", planID, userID).First(&planMember).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &planMember, nil
}

func InsertMember(planMember *models.PlanMember) error {
	return config.DB.Create(planMember).Error
}

func LeavePlan(planMemberID uint64) error {
	return config.DB.Model(&models.PlanMember{}).
		Where("id = ?", planMemberID).
		Updates(map[string]interface{}{
			"left_at":    time.Now(),
			"is_deleted": true,
		}).Error
}
