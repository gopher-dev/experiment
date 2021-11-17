package resolver

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"

	"github.com/graph-gophers/graphql-go"
	"github.com/kumparan/go-utils"
	log "github.com/sirupsen/logrus"
)

type (
	// Handler :nodoc:
	Handler struct {
		Schema *graphql.Schema
	}

	// PersistedQueryExtension :nodoc:
	PersistedQueryExtension struct {
		Version    int    `json:"version" url:"version"`
		Sha256Hash string `json:"sha256Hash" url:"sha256Hash"`
	}

	// Extensions :nodoc:
	Extensions struct {
		PersistedQuery *PersistedQueryExtension `json:"persistedQuery" url:"persistedQuery"`
	}

	// a workaround for getting `variables` as a JSON string
	requestOptionsCompatibility struct {
		Query         string `json:"query,omitempty" url:"query"`
		Variables     string `json:"variables" url:"variables"`
		OperationName string `json:"operationName" url:"operationName"`
		ID            string `json:"id,omitempty" url:"id"`
		PurgeCDN      string `json:"purgeCDN,omitempty" url:"purgeCDN"`
	}

	RequestOptions struct {
		Query         string                 `json:"query,omitempty" url:"query"`
		Variables     map[string]interface{} `json:"variables" url:"variables"`
		OperationName string                 `json:"operationName" url:"operationName"`
		Extensions    *Extensions            `json:"extensions,omitempty" url:"extensions"`

		// use to save purge cdn flag and complete url of request
		PurgeCDN    bool `json:"purgeCDN,omitempty" url:"purgeCDN"`
		PurgeCDNURL string
	}
)

// ServeEcho :nodoc:
func (h *Handler) ServeEcho(e *echo.Echo) {
	e.Any("/query/", h.QueryHandler)
}

// QueryHandler :nodoc:
func (h Handler) QueryHandler(ec echo.Context) error {
	var (
		requestOptions interface{}
	)

	// read body request
	body, err := ioutil.ReadAll(ec.Request().Body)
	if err != nil {
		log.WithField("context", utils.Dump(ec)).Error(err)
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	requestOptions = parseRequestOptionsFromBody(body)

	var output interface{}
	switch opts := requestOptions.(type) {
	case *RequestOptions:
		// TODO
	case []*RequestOptions:
		// TODO
	default:
		log.WithField("i", utils.Dump(opts))
		return ec.JSON(http.StatusBadRequest, opts)
	}

	return ec.JSON(http.StatusOK, output)
}

func (h *Handler) Initialize(schemaFile string, resolver *Resolver) {
	s, err := getSchema(schemaFile)
	if err != nil {
		panic(err)
	}
	schema := graphql.MustParseSchema(s, resolver)
	h.Schema = schema
}

// getSchema loads schema to a string
func getSchema(path string) (string, error) {
	filepaths, err := filePathWalkDir(path)
	if err != nil {
		log.WithField("path", utils.Dump(path)).Error(err)
		return "", err
	}

	var buffer bytes.Buffer
	for _, fp := range filepaths {
		b, err := ioutil.ReadFile(filepath.Clean(fp))
		if err != nil {
			log.WithField("path", utils.Dump(path)).Error(err)
			return "", err
		}
		_, _ = buffer.Write(b)
	}

	return buffer.String(), nil
}

func filePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func parseRequestOptionsFromBody(body []byte) interface{} {
	if bytes.HasPrefix(body, []byte("[")) {
		var queries []*RequestOptions
		err := json.Unmarshal(body, &queries)
		if err != nil {
			var optionsCompatible []*requestOptionsCompatibility
			_ = json.Unmarshal(body, &optionsCompatible)
			for i := range optionsCompatible {
				_ = json.Unmarshal([]byte(optionsCompatible[i].Variables), &queries[i].Variables)
			}
			return queries
		}
		return queries
	}

	var query RequestOptions
	err := json.Unmarshal(body, &query)
	if err != nil {
		var optionsCompatible requestOptionsCompatibility
		_ = json.Unmarshal(body, &optionsCompatible)
		_ = json.Unmarshal([]byte(optionsCompatible.Variables), &query.Variables)
		return &query
	}
	return &query
}
