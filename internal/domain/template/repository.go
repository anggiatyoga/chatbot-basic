package template

type TemplateRepo interface {
	Save(template Template) error
	Update(template Template) error
	Delete(id string) error
	ReadById(id string) (Template, error)
	ReadAll(qPage int, qOwnerID string) ([]DataTemplate, error)
}
