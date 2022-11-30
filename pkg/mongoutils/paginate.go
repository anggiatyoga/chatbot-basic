package mongoutils

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoPaginate struct {
	limit int64
	page  int64
}

func NewMongoPaginate(limit, page int) *mongoPaginate {
	return &mongoPaginate{
		limit: int64(limit),
		page:  int64(page),
	}
}

func (mp *mongoPaginate) GetPaginatedOpts() *options.FindOptions {
	l := mp.limit
	skip := mp.page*mp.limit - mp.limit
	fOpt := options.FindOptions{Limit: &l, Skip: &skip, Sort: bson.D{{"updated_at", -1}}}

	return &fOpt
}
