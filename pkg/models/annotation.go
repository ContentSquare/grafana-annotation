package models

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	elasticsearch7 "github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

func NewAnnotation() ESAnnotationBuilder {
	return &esAnnotationBuilder{}
}

type EsAnnotation struct {
	// Elastic search grafana model
	Time *time.Time `json:"time"`
	Text string     `json:"text"`
	Tags []string   `json:"tags"`

	// Custom fields
	Environment string `json:"environment,omitempty"`
	Role        string `json:"role,omitempty"`
	Region      string `json:"region,omitempty"`
	Provider    string `json:"provider,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
}

type ESAnnotationBuilder interface {
	WithText(text string) ESAnnotationBuilder
	WithTag(tag string) ESAnnotationBuilder
	WithTags(tags []string) ESAnnotationBuilder
	WithRole(role string) ESAnnotationBuilder
	WithRegion(region string) ESAnnotationBuilder
	WithEnvironment(env string) ESAnnotationBuilder
	WithProvider(provider string) ESAnnotationBuilder
	WithHostname(hostname string) ESAnnotationBuilder
	Build() *EsAnnotation
}

type esAnnotationBuilder struct {
	// Elastic search grafana model
	text string
	tags []string

	// Custom fields
	environment string
	role        string
	region      string
	provider    string
	hostname    string
}

func (e *esAnnotationBuilder) WithText(text string) ESAnnotationBuilder {
	e.text = text
	return e
}

func (e *esAnnotationBuilder) WithTag(tag string) ESAnnotationBuilder {
	if e.tags == nil {
		e.tags = []string{tag}
	} else {
		e.tags = append(e.tags, tag)
	}
	return e
}

func (e *esAnnotationBuilder) WithTags(tags []string) ESAnnotationBuilder {
	if e.tags == nil {
		e.tags = tags
	} else {
		for _, x := range tags {
			e.tags = append(e.tags, x)
		}
	}
	return e
}

func (e *esAnnotationBuilder) WithRole(role string) ESAnnotationBuilder {
	e.role = role
	return e
}

func (e *esAnnotationBuilder) WithRegion(region string) ESAnnotationBuilder {
	e.region = region
	return e
}

func (e *esAnnotationBuilder) WithProvider(provider string) ESAnnotationBuilder {
	e.provider = provider
	return e
}

func (e *esAnnotationBuilder) WithEnvironment(env string) ESAnnotationBuilder {
	e.environment = env
	return e
}

func (e *esAnnotationBuilder) WithHostname(hostname string) ESAnnotationBuilder {
	e.hostname = hostname
	return e
}

func (e *esAnnotationBuilder) Build() *EsAnnotation {
	now := time.Now()
	if e.tags == nil {
		e.tags = []string{}
	}
	return &EsAnnotation{
		Time:        &now,
		Text:        e.text,
		Tags:        e.tags,
		Environment: e.environment,
		Role:        e.role,
		Region:      e.region,
		Provider:    e.provider,
		Hostname:    e.hostname,
	}
}

func (e *EsAnnotation) String() string {
	jsonDate, err := json.Marshal(e)
	if err != nil {
		log.Errorf("unable to marshal to json EsAnnotation")
		return ""
	}
	return string(jsonDate)
}

func (e *EsAnnotation) GetIndexName(indexPattern string) string {
	return e.Time.Format(indexPattern)
}

func (e *EsAnnotation) getDocumentID() string {
	docIDString := strconv.FormatInt(e.Time.Unix(), 10)
	hasher := md5.New()
	hasher.Write([]byte(docIDString))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (e *EsAnnotation) PutDoc(client *elasticsearch7.Client, indexPattern string) (*esapi.Response, error) {
	req := esapi.IndexRequest{
		Index:      e.GetIndexName(indexPattern),
		DocumentID: e.getDocumentID(),
		Body:       strings.NewReader(e.String()),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), client)
	if err == nil {
		log.Debugf("Put Status: %v", res.Status())
	}

	return res, err
}
