package db

import (
	"example/db/elasticsearch"
)

func RegisterDB() {
	elasticsearch.Initialize()
}
