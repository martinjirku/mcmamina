package services

import (
	"encoding/json"
	"fmt"
	"os"
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
	dist string
}

func NewCSS(dist string) *CSS {
	return &CSS{dist: dist}
}

func (c *CSS) GetCssPath() (string, error) {
	file, err := os.Open(path.Join(".", c.dist, ".vite", "manifest.json"))
	if err != nil {
		return "", err
	}
	defer file.Close()
	var manifest Manifest
	err = json.NewDecoder(file).Decode(&manifest)
	if err != nil {
		return "", fmt.Errorf("decode manifest.json: %w", err)
	}

	style, ok := manifest["style.css"]
	if !ok {
		return "", fmt.Errorf("get style.css from manifest: %w", err)
	}
	if style.File == "" {
		return "", fmt.Errorf("get the path for style.css")
	}

	return path.Join(style.File), nil
}
