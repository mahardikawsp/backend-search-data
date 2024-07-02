package elasticsearch

import (
	"example/config"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"strconv"
	"strings"
)

type Client struct {
	*elasticsearch.Client
}

var client *Client

func Initialize() {
	esInfo := config.GetElasticSearchInfo()
	useSSL, _ := strconv.ParseBool(esInfo.UseSSL)
	cfg := elasticsearch.Config{}

	if useSSL {
		panic("implement me")
	} else {
		cfg.Addresses = strings.Split(esInfo.Host, ",")
		cfg.Username = esInfo.Username
		cfg.Password = esInfo.Password
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic("Failed to connect elasticsearch \n")
	}

	res, err := es.Info()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("es response: %v\n", res.Status())
	defer res.Body.Close()

	client = &Client{es}
}

func GetEs() *Client {
	if client == nil {
		Initialize()
	}

	return client
}
