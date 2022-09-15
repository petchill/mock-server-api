package response

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResponseDB struct {
	ID            primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	Project       string                 `bson:"service" json:"service" form:"service" validate:"required"`
	Path          string                 `bson:"path" json:"path" form:"path" validate:"required"`
	StatusCode    int                    `bson:"status_code" json:"status_code" form:"status_code" validate:"required"`
	Body          map[string]interface{} `bson:"body" json:"body" form:"body" validate:"required"`
	RegexPath     primitive.Regex        `bson:"regex_path"`
	ChildResponse []ResponseDB
}

type ResponseDBQuery struct {
	Path string `bson:"path" `
}
