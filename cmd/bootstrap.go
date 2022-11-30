package main

import (
	"chatbotbasic/internal/domain/template"
	"chatbotbasic/internal/platform/databases"
	"chatbotbasic/internal/platform/webapi"
	"go.mongodb.org/mongo-driver/mongo"
)

func Bootstrap(db *mongo.Database) *webapi.AppModule {
	templateRepo := databases.NewTemplateRepoImpl(db)
	templateUseCase := template.NewTemplateUseCase(templateRepo)

	modules := &webapi.AppModule{
		TemplateModule: *templateUseCase,
	}

	return modules
}
