package databases

import (
	"chatbotbasic/cmd/config"
	"chatbotbasic/internal/domain/template"
	"chatbotbasic/pkg/mongoutils"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	collection = "projects"
)

type TemplateRepoImpl struct {
	db *mongo.Database
}

func NewTemplateRepoImpl(db *mongo.Database) *TemplateRepoImpl {
	return &TemplateRepoImpl{db: db}
}

func (r *TemplateRepoImpl) Save(tt template.Template) error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	if len(tt.Stories) == 0 {
		tt.Stories = make([]template.Stories, 0)
	}

	if len(tt.Stories[0].Edges) == 0 {
		tt.Stories[0].Edges = make([]template.Edges, 0)
	}

	if len(tt.Stories[0].Blocks) == 0 {
		tt.Stories[0].Blocks = make([]template.Blocks, 0)
	}

	//for i, v := range tt.Stories[0].Blocks {
	//	for j, w := range v.Contents {
	//		tt.Stories[0].Blocks[i].Contents[j].Payload = w.Payload.ToEntityConditional()
	//	}
	//}

	_, err := r.db.Collection(collection).InsertOne(ctx, tt)
	if err != nil {
		log.Error().Str("template repo", "save template").Msg(err.Error())
		return err
	}
	return nil
}

func (r *TemplateRepoImpl) Update(template template.Template) error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	col := r.db.Collection(collection)
	filter := bson.M{"_id": template.ID}

	res := col.FindOneAndReplace(ctx, filter, template)

	if res.Err() != nil {
		log.Error().Str("template repo", "update template").Msg(res.Err().Error())
		return res.Err()
	}

	return nil
}

func (r *TemplateRepoImpl) ReadById(id string) (template.Template, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var result template.Template

	filter := bson.D{{"_id", id}}
	err := r.db.Collection(collection).FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return template.Template{}, err
	} else if err != nil {
		return template.Template{}, err
	}

	return result, nil
}

func (r *TemplateRepoImpl) Delete(id string) error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.D{{"$set", bson.D{
		{"deleted_at", time.Now()},
		{"updated_at", time.Now()},
	}}}

	res := r.db.Collection(collection).FindOneAndUpdate(ctx, filter, update)
	if res.Err() != nil {
		log.Error().Str("template repo", "Delete template").Msg(res.Err().Error())
		return res.Err()
	}

	return nil
}

func (r *TemplateRepoImpl) ReadAll(qPage int, qOwnerID string) ([]template.DataTemplate, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var templates []template.DataTemplate

	//fmt.Printf("type=%s sort=%s page=%s\n", qType, qSort, qPage)

	//filter := bson.D{{"$or", []interface{}{
	//	bson.D{{"deletedAt", nil}},
	//	bson.D{{"deletedAt", bson.M{"$exists": false}}},
	//}}}

	filter := bson.D{
		{"$and", []bson.M{
			{"type": "BASIC"},
			{"deleted_at": nil},
		}},
	}

	if qOwnerID != "" {
		filter = bson.D{
			{"$and", []bson.M{
				{"type": "BASIC"},
				{"deleted_at": nil},
				{"owner_id": qOwnerID},
			}},
		}
	}

	cursor, err := r.db.Collection(collection).Find(ctx, filter, mongoutils.NewMongoPaginate(10, qPage).GetPaginatedOpts())
	if err != nil {
		log.Error().Str("template repo", "ReadAll template").Msg(err.Error())
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &templates); err != nil {
		log.Error().Str("template repo", "ReadAll cursor template").Msg(err.Error())
		return nil, err
	}

	if len(templates) == 0 {
		templates = make([]template.DataTemplate, 0)
		return templates, nil
	}

	return templates, nil
}
