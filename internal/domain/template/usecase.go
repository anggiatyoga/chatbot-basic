package template

import (
	"time"
)

type UseCase interface {
	SaveTemplate(id string, template Template) error
	UpdateTemplate(template Template) error
	RemoveTemplate(id string) error
	GetTemplateById(id string) (Template, error)
	GetAllTemplate(qPage int, qOwnerID string) ([]DataTemplate, error)
	//GetTemplateForCoster(id string) (coster.TemplateCoster, error)
	//GetTemplateForDev(id string) (coster.TemplateCoster, error)
}

type TemplateUseCase struct {
	repo TemplateRepo
}

func NewTemplateUseCase(r TemplateRepo) *TemplateUseCase {
	return &TemplateUseCase{repo: r}
}

func (u *TemplateUseCase) SaveTemplate(id string, template Template) error {
	// generated id
	//id, err := uuid.NewUUID()
	//if err != nil {
	//	return err
	//}
	//template.ID = id.String()
	template.ID = id
	// generated time create, update,
	template.CreatedAt = time.Now()
	template.UpdatedAt = time.Now()

	err := u.repo.Save(template)
	return err
}

func (u *TemplateUseCase) UpdateTemplate(template Template) error {
	template.UpdatedAt = time.Now()
	return u.repo.Update(template)
}

func (u *TemplateUseCase) RemoveTemplate(id string) error {
	return u.repo.Delete(id)
}

func (u *TemplateUseCase) GetTemplateById(id string) (Template, error) {
	return u.repo.ReadById(id)
}

func (u *TemplateUseCase) GetAllTemplate(qPage int, qOwnerID string) ([]DataTemplate, error) {
	return u.repo.ReadAll(qPage, qOwnerID)
}

//func (u *TemplateUseCase) GetTemplateForDev(id string) (coster.TemplateCoster, error) {
//	template, err := u.repo.ReadById(id)
//	if err != nil {
//		return coster.TemplateCoster{}, err
//	}
//
//	var mapBloc map[string]Blocks
//	for _, b := range template.Stories[0].Blocks {
//		mapBloc[b.ID] = b
//	}
//
//	var parentId []Blocks
//	var payloads []Blocks
//
//	for _, e := range template.Stories[0].Edges {
//		parentId = append(parentId, mapBloc[e.SourceBlockID])
//		payloads = append(payloads, mapBloc[e.TargetBlockID])
//	}
//
//	var tc []coster.TemplateContents
//	for i, t := range parentId {
//		var ps []coster.Payloads
//		for _, p := range t.Contents {
//			ps = append(ps, p.ToPayloadTemplateCoster())
//		}
//		tc = append(tc, t.ToTemplateCosterDev(parentId[i].ToParentIdsCoster(), ps))
//	}
//
//	templateCoster := coster.TemplateCoster{
//		ID:                  template.ID,
//		TemplateName:        template.Title,
//		TemplateDescription: template.Description,
//		ClientID:            "",
//		ChannelID:           "",
//		AccountID:           "",
//		AccountAlias:        "",
//		DivisionID:          "",
//		IsDeleted:           false,
//		OwnerID:             template.OwnerID,
//		UpdatedAt:           template.UpdatedAt,
//		CreatedAt:           template.CreatedAt,
//		TemplateContents:    tc,
//	}
//	return templateCoster, nil
//}

//func (u *TemplateUseCase) GetTemplateForCoster(id string) (coster.TemplateCoster, error) {
//	template, err := u.repo.ReadById(id)
//	if err != nil {
//		return coster.TemplateCoster{}, err
//	}
//
//	var tc []coster.TemplateContents
//	for _, b := range template.Stories[0].Blocks {
//		tc = append(tc, b.ToTemplateCoster(template.Stories[0].Blocks))
//	}
//
//	templateCoster := coster.TemplateCoster{
//		ID:                  template.ID,
//		TemplateName:        template.Title,
//		TemplateDescription: template.Description,
//		ClientID:            "",
//		ChannelID:           "",
//		AccountID:           "",
//		AccountAlias:        "",
//		DivisionID:          "",
//		IsDeleted:           false,
//		OwnerID:             template.OwnerID,
//		UpdatedAt:           template.UpdatedAt,
//		CreatedAt:           template.CreatedAt,
//		TemplateContents:    tc,
//	}
//
//	return templateCoster, nil
//}
