package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"example/db/elasticsearch"
	"example/helper"
	"example/model"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"reflect"
	"text/template"
)

type (
	SearchRepository interface {
		List(ctx context.Context, param *model.NewSearchParam) ([]model.Satudata, int, error)
	}

	repo struct {
		*elasticsearch.Client
	}
)

func NewSearchRepository() SearchRepository {
	return &repo{
		elasticsearch.GetEs(),
	}
}

func (r repo) List(ctx context.Context, param *model.NewSearchParam) ([]model.Satudata, int, error) {
	var (
		result   map[string]interface{}
		buff     bytes.Buffer
		response []model.Satudata
		satudata model.Satudata
	)
	param.Offset, param.Limit = helper.GetLimitOffset(param.Page, param.Limit)

	fns := template.FuncMap{
		"last": func(x int, a interface{}) bool {
			fmt.Printf("datanya: %v\n", a)
			return x == reflect.ValueOf(a).Len()-1
		},
	}

	queryTemplate := `{
		"from": {{ .Offset }},
		"size": {{ .Limit }},
		{{ if .Sort }}
		"sort": [
			{{ if eq .Sort "NAMA_DATA" }}
			{
				"{{ .Sort }}.keyword": "{{ .Order }}"
			}
			{{ else }}
			{
				"{{ .Sort }}": "{{ .Order }}"
			}
			{{ end }}
		],
		{{ end }}
		"query": {
			"bool": {
				"must": [
					{{ if .Keyword }}
					{
						"query_string": {
							"query": "*{{ .Keyword }}*",
							"fields": [
								"INDIKATOR",
								"NAMA_DATA^3",
								"TAGS",
								"URUSAN^2"
							],
							"analyzer": "simple",
							"default_operator": "OR"
						}
					}
					{{ else }}
					{
						"query_string": {
							"query": "*"
						}
					}
					{{ end }}
				]
				{{ if or .Skpd .Urusan .Kategori }}
				,
				"should": [
					{{ if .Skpd }}
						{{ range $i, $e := .Skpd }}
							{{ if $i }},{{ end }} {{ if last $i $e }}{{ end }}
							{
								"match_phrase": {
									"PRODUSEN": "{{ . }}"
								}
							}
						{{ end }}
					{{ end }}
					{{ if .Urusan }}
					{{ if .Skpd }}
					,
					{{ end }}
						{{ range $i, $e := .Urusan }}
							{{ if $i }},{{ end }} {{ if last $i $e }}{{ end }}
							{
								"match_phrase": {
									"URUSAN": "{{ . }}"
								}
							}
						{{ end }}
					{{ end }}
					{{ if .Kategori }}
					{{ if or .Urusan .Skpd }}
					,
					{{ end }}
						{{ range $i, $e := .Kategori }}
							{{ if $i }},{{ end }} {{ if last $i $e }}{{ end }}
							{
								"match_phrase": {
									"KATEGORI": "{{ . }}"
								}
							}
						{{ end }}
					{{ end }}
				
				],
				"minimum_should_match": "1"
				{{end}}
			}
		}
	}`

	tmpl, err := template.New("test").Funcs(fns).Parse(queryTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&buff, param)
	if err != nil {
		panic(err)
	}

	fmt.Println(buff.String())
	//
	//query := map[string]interface{}{
	//	"from": param.Offset,
	//	"size": param.Limit,
	//	"query": map[string]interface{}{
	//		"bool": map[string]interface{}{
	//			"must": map[string]interface{}{
	//				"multi_match": map[string]interface{}{
	//					"query":    param.Keyword,
	//					"fields":   []string{"INDIKATOR", "NAMA_DATA^2"},
	//					"operator": "AND",
	//				},
	//			},
	//		},
	//	},
	//}
	//
	//if err := json.NewEncoder(&buff).Encode(query); err != nil {
	//	log.Error(err)
	//	return nil, 0, err
	//}

	res, err := r.Client.Search(
		r.Client.Search.WithContext(ctx),
		r.Client.Search.WithIndex(viper.GetString("ES_INDEX")),
		r.Client.Search.WithBody(&buff),
		r.Client.Search.WithPretty(),
	)
	if err != nil {
		log.Error(err)
		return nil, 0, err
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

		return nil, 0, err
	}

	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Error(err)
		return nil, 0, err
	}

	for _, item := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		doc := item.(map[string]interface{})
		satudata.ID = fmt.Sprintf("%v", doc["_id"])
		err = mapstructure.Decode(doc["_source"], &satudata)
		if err != nil {
			log.Error(err)
		}

		response = append(response, satudata)
	}

	total := result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]

	return response, int(total.(float64)), nil
}
