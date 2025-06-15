package metafile

import (
	"errors"
	"fmt"
	"os"

	toml "github.com/pelletier/go-toml/v2"
)

type TomlFileHandler[T any] struct {
	path    string
	Content *T
}

func (fh *TomlFileHandler[T]) Read() error {
	fileContent, err := os.ReadFile(fh.path)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		fmt.Println("file does not exist, start with empty values: ", fh.path)
		fileContent = []byte{}
	} else if err != nil {
		return fmt.Errorf("failed to read toml file: %w", err)
	}

	if err := toml.Unmarshal(fileContent, fh.Content); err != nil {
		return fmt.Errorf("failed to unmarshal toml file: %w", err)
	}
	return nil
}

func (fh *TomlFileHandler[T]) Write() error {
	content, err := toml.Marshal(fh.Content)
	if err != nil {
		return fmt.Errorf("failed to marshal toml file: %w", err)
	}
	return os.WriteFile(fh.path, content, 0644)
}

func NewTomlFileHandler[T any](path string, content *T) *TomlFileHandler[T] {
	return &TomlFileHandler[T]{
		path:    path,
		Content: content,
	}
}
