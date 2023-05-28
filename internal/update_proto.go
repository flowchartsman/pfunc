//go:build ignore
// +build ignore

package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type shas struct {
	Files map[string]string `json:"files"`
}

func must[T any](f func() (T, error)) T {
	_, _, line, ok := runtime.Caller(1)
	if !ok {
		panic("could not get caller")
	}
	t, err := f()
	if err != nil {
		log.Fatalf("generate.go[%d]: %v", line, err)
	}
	return t
}

func calcSHA(path string) (string, error) {
  f, err := os.Open(path)
  if err != nil {
	  return "", err
  }
  defer f.Close()

  h := sha256.New()
  if _, err := io.Copy(h, f); err != nil {
	  return "", err
  }

  return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func getFileShas() (*shas, error) {
	shas := &shas{
		Files: map[string]string{},
	}
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return filepath.SkipDir
		}
		log.Println(path)
		filename := info.Name()
		if filename == "generate.go" {
			return nil
		}
		shas.Files[filename], err = calcSHA(filename)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return shas, nil
}

func loadShas() (*shas, error) {
	shas := &shas{}
	shafile, err := os.Open("checksums.json")
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(shafile).Decode(shas); err != nil {
		return nil, err
	}
	return shas, nil
}

func main() {
	log.SetFlags(0)
	shas := must(getFileShas)
	log.Printf("%#v", shas)
}
