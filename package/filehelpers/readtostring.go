package filehelpers

import (
	"github.com/spf13/afero"
	"strings"
)

func ReadToString(fs afero.Fs, path string) (string, error) {
	file, err := fs.Open(path)
	if err != nil {
		return "", err
	}
	raw := []byte{}
	_, err = file.Read(raw)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

func ReadFirstLineString(fs afero.Fs, path string) (string, error) {
	raw, err := ReadToString(fs, path)
	if err != nil {
		return "", err
	}
	return strings.Split(string(raw), "\n")[0], nil
}
