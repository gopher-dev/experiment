package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gopher-dev/experiment/graphql-go-simple/resolver"
	"github.com/kumparan/go-utils"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	log "github.com/sirupsen/logrus"
)

func main() {
	s, err := getSchema("./schema")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	schema := graphql.MustParseSchema(s, &resolver.Resolver{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// getSchema loads schema to a string
func getSchema(path string) (string, error) {
	filePaths, err := filePathWalkDir(path)
	fmt.Println(filePaths)
	if err != nil {
		log.WithField("path", utils.Dump(path)).Error(err)
		return "", err
	}

	var buffer bytes.Buffer
	for _, fp := range filePaths {
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
