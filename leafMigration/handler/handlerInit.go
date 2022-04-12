package handler

import (
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper"
)

func (h handler) Init(project string) error {
	if err := helper.Initialize(helper.InitializeProjectRequestDTO{
		ProjectName: project,
	}); err != nil {
		h.log.StandardLogger().Errorf("[%s] error initializing project: %+v", project, err.Error())
		return err
	}

	h.log.StandardLogger().Infof("[%s] finish initializing project", project)
	return nil
}
