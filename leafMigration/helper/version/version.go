package version

import (
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Version uint64

const MigrationTable = "leaf_migrations"
const LatestVersion Version = math.MaxUint64
const NoVersion Version = 0

type (
	DataVersion struct {
		// Note: concat Version and NeutralizedName
		// NeutralizedName is name replace space with _
		ObjectID    primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
		ID          string             `json:"id" bson:"id" gorm:"primaryKey;column=id"`
		Version     uint64             `json:"version" json:"version" gorm:"index;column=version"`
		Name        string             `json:"name" json:"name" gorm:"column=name"`
		ExecuteTime string             `json:"execute_time" json:"execute_time" gorm:"column=execute_time"`
	}
)

func (d *DataVersion) TableName() string {
	return MigrationTable
}

func (d *DataVersion) ToBson() bson.M {
	return bson.M{
		"id":           d.ID,
		"version":      d.Version,
		"name":         d.Name,
		"execute_time": d.ExecuteTime,
	}
}
