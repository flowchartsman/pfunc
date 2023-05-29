//go:build ignore
// +build ignore

package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

type repoTag struct {
	Tag     string `json:"name"`
	Tarball string `json:"tarball_url"`
	Commit  struct {
		SHA string `json:"sha"`
	} `json:"commit"`
}

func requireTools(binNames ...string) {
	for _, n := range binNames {
		_, err := exec.LookPath(n)
		if err != nil {
			log.Fatalf("couldn't find %s: %v", n, err)
		}
	}
}

func fileSHA(path string) (string, error) {
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

const fnAPIPath = `./internal/fnapi`

func getProtoGoSHAs() (map[string]string, error) {
	shas := map[string]string{}
	goProtoFiles, err := filepath.Glob(filepath.Join(fnAPIPath, "*.go"))
	if err != nil {
		return nil, fmt.Errorf("listing generated files in %s: %v", fnAPIPath, err)
	}
	for _, path := range goProtoFiles {
		filename := filepath.Base(path)
		if filename == "doc.go" {
			continue
		}
		shas[filename], err = fileSHA(path)
		if err != nil {
			return nil, fmt.Errorf("calculating sha for %s: %v", filename, err)
		}
	}
	if len(shas) == 0 {
		return nil, fmt.Errorf("no generated files found in %s", fnAPIPath)
	}
	return shas, nil
}

const pulsarTagsURL = "https://api.github.com/repos/apache/pulsar/tags"

var isRelease = regexp.MustCompile(`^v\d+\.\d+.\d+`)

func getPulsarTags() (map[string]*repoTag, error) {
	var tags []*repoTag
	page := 1
	for {
		pageURL := fmt.Sprintf("%s?page=%d", pulsarTagsURL, page)
		resp, err := http.Get(pageURL)
		if err != nil {
			return nil, fmt.Errorf("fetching pulsar tags: %v", err)
		}
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("fetching pulsar tags: non-200 code %d", resp.StatusCode)
		}
		var pageTags []*repoTag
		if err := json.NewDecoder(resp.Body).Decode(&pageTags); err != nil {
			return nil, fmt.Errorf("page %d json decode: %v", page, err)
		}
		if len(pageTags) == 0 {
			break
		}
		tags = append(tags, pageTags...)
		page++
	}

	if len(tags) == 0 {
		return nil, fmt.Errorf("no tags found")
	}

	out := map[string]*repoTag{}
	for _, tag := range tags {
		if isRelease.MatchString(tag.Tag) {
			out[tag.Tag] = tag
		}
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no release tags found")
	}
	return out, nil
}

const protoPrefix = `pulsar-functions/proto/src/main/proto`

func getProtosToTmpdir(releaseTag *repoTag) (string, error) {
	tmpdir, err := os.MkdirTemp("", "pfunc_gen")
	if err != nil {
		return "", fmt.Errorf("creating temporary directory: %v", err)
	}
	defer func() {
		if err != nil {
			log.Println("removing temporary directory")
			os.RemoveAll(tmpdir)
		}
	}()
	resp, err := http.Get(releaseTag.Tarball)
	if err != nil {
		return "", fmt.Errorf("fetching release tarball: %v", err)
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("fetching release tarball: non-200 code %d", resp.StatusCode)
	}
	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("creating gzip reader: %v", err)
	}
	defer gz.Close()
	tarReader := tar.NewReader(gz)
	for {
		header, err := tarReader.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return "", fmt.Errorf("untar: %v", err)
		}

		if err != nil {
			return "", fmt.Errorf("tar.Next(): %v", err)
		}

		switch header.Typeflag {
		case tar.TypeReg:
			tarPath := strings.SplitAfterN(header.Name, `/`, 2)[1]
			if !strings.HasPrefix(tarPath, protoPrefix) {
				break
			}
			protoFile := filepath.Base(header.Name)
			log.Printf("\t%s", protoFile)
			protoOut, err := os.Create(filepath.Join(tmpdir, protoFile))
			if err != nil {
				return "", fmt.Errorf("extracting %s: %v", protoFile, err)
			}
			if _, err := io.Copy(protoOut, tarReader); err != nil {
				return "", fmt.Errorf("extracting %s: %v", protoFile, err)
			}
			protoOut.Close()
		}
	}
	return tmpdir, nil
}

var docTmpl = template.Must(template.New("doc").Parse(docTmplSrc))

func updateProto(release string) error {
	currentFiles, err := getProtoGoSHAs()
	if err != nil {
		return err
	}

	log.Println("fetching pulsar release tags")
	pulsarTags, err := getPulsarTags()
	if err != nil {
		return err
	}
	releaseTag, found := pulsarTags[release]
	if !found {
		return fmt.Errorf("release %s not found", release)
	}
	log.Println("extracting release protos")
	protoLoc, err := getProtosToTmpdir(releaseTag)
	if err != nil {
		return err
	}
	defer os.RemoveAll(protoLoc)

	protoFiles, err := filepath.Glob(protoLoc + `/*.proto`)
	if err != nil {
		return fmt.Errorf("getting list of proto files in temp dir: %v", err)
	}
	protoCmd := []string{
		"protoc",
		"--proto_path", protoLoc,
		"--go_out=./internal",
		"--go-grpc_out=./internal",
	}
	for _, protoPath := range protoFiles {
		protoBase := filepath.Base(protoPath)
		protoCmd = append(protoCmd, fmt.Sprintf("--go_opt=M%s=/fnapi", protoBase))
		protoCmd = append(protoCmd, fmt.Sprintf("--go-grpc_opt=M%s=/fnapi", protoBase))
	}
	protoCmd = append(protoCmd, protoFiles...)
	if _, err := runCmd(protoCmd...); err != nil {
		return err
	}

	changes := false
	toDelete := []string{}

	newFiles, err := getProtoGoSHAs()
	for newFile, newSHA := range newFiles {
		currentSHA, found := currentFiles[newFile]
		if currentSHA == newSHA {
			continue
		}
		if !found {
			log.Printf("%s: NEW", newFile)
		} else {
			log.Printf("%s: UPDATED", newFile)
		}
		changes = true
	}
	for currentFile := range currentFiles {
		if _, found := newFiles[currentFile]; !found {
			log.Printf("%s: DELETED", currentFile)
			changes = true
			toDelete = append(toDelete, filepath.Join(fnAPIPath, currentFile))
		}
	}
	if !changes {
		log.Println("no changes")
		return nil
	}
	for _, deleted := range toDelete {
		os.Remove(filepath.Join(fnAPIPath, deleted))
	}
	newDoc, err := os.Create(filepath.Join(fnAPIPath, "doc.go"))
	if err != nil {
		return fmt.Errorf("creating updated doc.go: %v", err)
	}
	docTmpl.Execute(newDoc, releaseTag)
	commitMsg := fmt.Sprintf("updated proto defs from pulsar release tag %s (%s)", releaseTag.Tag, releaseTag.Commit.SHA)
	log.Printf("protos changed, please test and commit with:\ngit add internal/fnapi\ngit commit -m \"%s\"", commitMsg)
	return nil
}

func runCmd(cmd ...string) (string, error) {
	c := exec.Command(cmd[0], cmd[1:]...)
	output, err := c.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s:\n********\n%s", c.String(), string(output))
	}
	return strings.TrimSpace(string(output)), nil
}

func main() {
	log.SetFlags(0)
	repoRoot, err := runCmd("git", "rev-parse", "--show-toplevel")
	if err != nil {
		log.Fatalf("get repo root: %v", err)
	}
	os.Chdir(repoRoot)
	if len(os.Args) != 2 || !isRelease.MatchString(os.Args[1]) {
		log.Fatal("usage: go run generate.go v<release number>")
	}
	requireTools("git", "protoc", "protoc-gen-go")

	release := os.Args[1]
	if err := updateProto(release); err != nil {
		log.Fatal(err)
	}
}

const docTmplSrc = `// Package pb provides the protocol buffer messages that Pulsar
// uses for the client/broker wire protocol.
// See "Pulsar binary protocol specification" for more information.
// https://pulsar.incubator.apache.org/docs/latest/project/BinaryProtocol/
//
// The protocol definition files are part of the main Pulsar source,
// located within the Pulsar repository at:
// https://github.com/apache/pulsar/tree/master/pulsar-functions/proto/src/main/proto
//
// The generated Go code was created from the source Pulsar files at git:
//    tag:      {{.Tag}}
//    commit:   {{.Commit.SHA}}
//
// Files generated by the protoc-gen-go program should not be modified.

// Code generated by generate.go. DO NOT EDIT.
package fnapi`
