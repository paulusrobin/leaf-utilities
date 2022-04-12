package handler

import (
	"context"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
)

func (h handler) Init(project string) error {
	if err := helper.Initialize(helper.InitializeProjectRequestDTO{
		ProjectName: project,
	}); err != nil {
		h.log.Error(leafLogger.BuildMessage(context.Background(), "[%s] error initializing project: %+v",
			leafLogger.WithAttr("project", project),
			leafLogger.WithAttr("error", err.Error())))
		return err
	}

	h.log.Info(leafLogger.BuildMessage(context.Background(), "[%s] finish initializing project",
		leafLogger.WithAttr("project", project)))
	return nil
}
