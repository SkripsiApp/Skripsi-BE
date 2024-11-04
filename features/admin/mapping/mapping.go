package mapping

import (
	"skripsi/features/admin/entity"
	"skripsi/features/admin/model"
)

// Mapping Core to Model
func AdminCoreToAdminModel(adminCore entity.AdminCore) model.Admin {
	adminModel := model.Admin{
		Id:       adminCore.Id,
		Username: adminCore.Username,
		Password: adminCore.Password,
		Role:     adminCore.Role,
	}
	return adminModel
}

func ListAdminCoreToAdminModel(adminCore []entity.AdminCore) []model.Admin {
	listAdminModel := []model.Admin{}
	for _, admin := range adminCore {
		adminModel := AdminCoreToAdminModel(admin)
		listAdminModel = append(listAdminModel, adminModel)
	}
	return listAdminModel
}

// Mapping Model to Core
func AdminModelToAdminCore(adminModel model.Admin) entity.AdminCore {
	adminCore := entity.AdminCore{
		Id:       adminModel.Id,
		Username: adminModel.Username,
		Password: adminModel.Password,
		Role:     adminModel.Role,
	}
	return adminCore
}

func ListAdminModelToAdminCore(adminModel []model.Admin) []entity.AdminCore {
	listAdminCore := []entity.AdminCore{}
	for _, admin := range adminModel {
		adminCore := AdminModelToAdminCore(admin)
		listAdminCore = append(listAdminCore, adminCore)
	}
	return listAdminCore
}
