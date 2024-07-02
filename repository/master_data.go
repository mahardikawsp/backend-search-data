package repository

import (
	"bytes"
	"encoding/json"
	"example/db/elasticsearch"
	"example/db/mysql"
	"example/model"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type (
	MasterDataRepository interface {
		ListSkpd() ([]model.Skpd, error)
		ListUrusan() ([]model.Urusan, error)
	}

	repoMasterData struct {
		db *mysql.Client
		*elasticsearch.Client
	}
)

func NewMasterDataRepository() MasterDataRepository {
	return &repoMasterData{
		mysql.GetDB(),
		elasticsearch.GetEs(),
	}
}

func (r repoMasterData) ListSkpd() ([]model.Skpd, error) {
	var (
		result   map[string]interface{}
		buff     bytes.Buffer
		response []model.Skpd
		skpd     model.Skpd
	)

	query := map[string]interface{}{
		"size": "1000",
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": "*",
			},
		},
	}

	if err := json.NewEncoder(&buff).Encode(query); err != nil {
		log.Error(err)
		return nil, err
	}

	res, err := r.Client.Search(
		r.Client.Search.WithIndex(viper.GetString("INDEX_SKPD")),
		r.Client.Search.WithBody(&buff),
		r.Client.Search.WithPretty(),
	)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	if res.IsError() {
		var e map[string]interface{}
		if err = json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Error(err)
		} else {
			log.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}

		return nil, err
	}

	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Error(err)
		return nil, err
	}

	for _, item := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		doc := item.(map[string]interface{})
		skpd.ID = fmt.Sprintf("%v", doc["_id"])
		err = mapstructure.Decode(doc["_source"], &skpd)
		if err != nil {
			log.Error(err)
		}

		response = append(response, skpd)
	}

	return response, nil
}

func (r repoMasterData) ListUrusan() ([]model.Urusan, error) {
	var (
		result   map[string]interface{}
		buff     bytes.Buffer
		response []model.Urusan
		urusan   model.Urusan
	)

	query := map[string]interface{}{
		"size": "1000",
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": "*",
			},
		},
	}

	if err := json.NewEncoder(&buff).Encode(query); err != nil {
		log.Error(err)
		return nil, err
	}

	fmt.Printf("query master_data: %v\n", buff.String())

	res, err := r.Client.Search(
		r.Client.Search.WithIndex(viper.GetString("INDEX_URUSAN")),
		r.Client.Search.WithBody(&buff),
		r.Client.Search.WithPretty(),
	)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	if res.IsError() {
		var e map[string]interface{}
		if err = json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Error(err)
		} else {
			log.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}

		return nil, err
	}

	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Error(err)
		return nil, err
	}

	for _, item := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		doc := item.(map[string]interface{})
		urusan.ID = fmt.Sprintf("%v", doc["_id"])
		err = mapstructure.Decode(doc["_source"], &urusan)
		if err != nil {
			log.Error(err)
		}

		response = append(response, urusan)
	}

	return response, nil
}
