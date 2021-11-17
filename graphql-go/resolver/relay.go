package resolver

import (
	"bytes"
	"io/ioutil"
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

	// a workaround for getting`variables` as a JSON string
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
	e.Any("/query/", func(ec echo.Context) error {
		return nil
	})
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
