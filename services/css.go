package services

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log/slog"
	"path"
)

type File struct {
	File    string   `json:"file,omitempty"`
	Imports []string `json:"imports,omitempty"`
	IsEntry bool     `json:"isEntry,omitempty"`
	Src     string   `json:"src,omitempty"`
}

type Manifest = map[string]File

type CSS struct {
	log *slog.Logger
	fs  fs.FS
}

func NewCSS(fs fs.FS, log *slog.Logger) *CSS {
	return &CSS{fs: fs, log: log}
}

func (c *CSS) GetCssPath() (string, error) {
	manifestPath := path.Join("./", ".vite", "manifest.json")
	c.log.Info(fmt.Sprintf("reading manifest.json: %s", manifestPath))
	file, err := c.fs.Open(manifestPath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var manifest Manifest
	c.log.Info(fmt.Sprintf("decoding manifest.json: %s", manifestPath))
	err = json.NewDecoder(file).Decode(&manifest)
	if err != nil {
		return "", fmt.Errorf("decode manifest.json: %w", err)
	}
	c.log.Info(fmt.Sprintf("accessing from manifest.json: %s", manifestPath))
	style, ok := manifest["style.css"]
	if !ok {
		return "", fmt.Errorf("get style.css from manifest: %w", err)
	}
	if style.File == "" {
		return "", fmt.Errorf("get the path for style.css")
	}

	return path.Join("/", style.File), nil
}
