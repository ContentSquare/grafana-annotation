package poster

import (
	"github.com/contentsquare/grafana-annotation/pkg/config"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/contentsquare/grafana-annotation/pkg/models"
	elasticsearch7 "github.com/elastic/go-elasticsearch"
	"strings"
)

func PostAnnotation(tags []string, text ...string) {
	cfg := config.LoadConfig()
	hostname, error := os.Hostname()
	if error != nil {
		log.Panicf("unable to retrieve hostname. err=%v", error.Error())
	}
	a := models.NewAnnotation().
		WithRegion(cfg.Region).
		WithEnvironment(cfg.Env).
		WithRole(cfg.Role).
		WithText(strings.Join(text, " ")).
		WithTags(tags).
		WithTag(cfg.Region).
		WithTag(cfg.Role).
		WithTag(cfg.Env).
		WithTag(cfg.Provider).
		WithTag(hostname).
		WithHostname(hostname).
		Build()

	client, err := elasticsearch7.NewClient(elasticsearch7.Config{
		Addresses: cfg.Elasticsearch.BootstrapServers,
		Transport: nil,
	})
	if err != nil {
		log.Errorf("Unable to initiate connection to elasticsearch. err=%v", err.Error())
	}
	resp, err := a.PutDoc(client, cfg.Elasticsearch.IndexName)
	if err != nil {
		log.Errorf("Unable to post annotation to elasticsearch. err=%v", err.Error())
	} else {
		if resp.IsError() {
			log.Errorf("Issue while posting annotation to elasticsearch. err=%v", resp.Status())
		}
		log.Infof("Annotation successfully posted to elasticsearch. tags=%v", a.Tags)
	}
}
