package main

import (
        "bytes"
        "encoding/json"
        "flag"
        "fmt"
        "github.com/mbndr/logo"
        "gopkg.in/yaml.v2"
        "io/ioutil"
        "net/http"
        "os"
        "time"
        "path/filepath"
        "os/user"
)

const (
        apiPath string = "api/annotations/graphite"
)

var (
        flagConfig FlagsConfig
        config     Config
        log        = logo.NewSimpleLogger(os.Stderr, logo.INFO, "grafana-annotation ", true)
)

type arrayFlags []string

func (i *arrayFlags) String() string {
        return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
        *i = append(*i, value)
        return nil
}

type FlagsConfig struct {
        FilePath string
        Verbose  bool
        Type     string
        What     string
        Data     string
        Tags     arrayFlags
}

func (c *FlagsConfig) Setup() {
        hostname, _ := os.Hostname()
        usr, _ := user.Current()
        dir := usr.HomeDir
        flag.StringVar(&c.FilePath, "config-file", filepath.Join(dir, ".grafana-anotation-poster.yml") , "Configuration File")
        flag.Var(&c.Tags, "tag", "Tags. may be repeated multiple times")
        flag.StringVar(&c.What, "what", hostname, "The What item to post.")
        flag.StringVar(&c.Data, "data", "", "Additional data.")
        flag.BoolVar(&c.Verbose, "verbose", false, "Be Verbose.")
        flag.Parse()
}

type Config struct {
        GrafanaUri  string `yaml:"grafanaUri"`
        BearerToken string `yaml:"bearerToken"`
}

func (c *Config) loadConfig(filePath string) {

        yamlFile, err := ioutil.ReadFile(filePath)
        if err != nil {
                log.Fatalf("Unable to open configuration file.Get err #%v ", err)
        }
        err = yaml.Unmarshal(yamlFile, c)
        if err != nil {
                log.Fatalf("Unmarshal: %v", err)
        }

        return
}

type GraphiteAnnotation struct {
        What string   `json:"what"`
        Tags []string `json:"tags"`
        When int64    `json:"when"`
        Data string   `json:"data"`
}

func (a *GraphiteAnnotation) toJson() []byte {
        payload, _ := json.Marshal(a)
        return payload
}

type jsonAnnotationResponse struct {
        Id      int    `json:id`
        Message string `json:message`
}

func (a *GraphiteAnnotation) post(url string, token string) {

        completeUrl := fmt.Sprintf("%v/%v", url, apiPath)
        payload := a.toJson()
        log.Debug(string(payload))
        log.Debug(completeUrl)
        req, err := http.NewRequest("POST", completeUrl, bytes.NewBuffer(payload))
        req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
        req.Header.Set("Content-Type", "application/json")

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
                log.Fatalf("unable to post to url %v. err=%v", completeUrl, err.Error())
        }
        defer resp.Body.Close()

        log.Debugf("response Status: %v", resp.Status)
        log.Debugf("response Headers: %v", resp.Header)
        body, _ := ioutil.ReadAll(resp.Body)
        var response jsonAnnotationResponse
        err = json.Unmarshal(body, &response)
        log.Debugf("response body: %v", response)
        if err != nil {
                log.Fatalf("Unable to parse response. response is %v. err=%v", string(body), err.Error())
        }
        if response.Id == 0 {
                // no id in response
                log.Info("Sending annotation. Message is ", response.Message)
        } else {
                log.Fatalf("response Body:", string(body))
        }
}

func NewGraphiteAnnotation(what string, tags arrayFlags, data string) GraphiteAnnotation {
        now := time.Now()
        when := now.Unix()
        log.Debugf("new anotation with %v %v %v %v", what, when, tags, data)
        return GraphiteAnnotation{What: what, When: when, Tags: tags, Data: data}
}

func main() {
        flagConfig.Setup()
        config.loadConfig(flagConfig.FilePath)
        if flagConfig.Verbose {
                log.SetLevel(logo.DEBUG)
        }
        if _, err := os.Stat(flagConfig.FilePath); os.IsNotExist(err) {
                log.Fatalf("Config file not found. %v", flagConfig.FilePath)
        }
        annotation := NewGraphiteAnnotation(flagConfig.What, flagConfig.Tags, flagConfig.Data)
        annotation.post(config.GrafanaUri, config.BearerToken)
}
