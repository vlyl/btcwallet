package ord

import (
	"errors"
	"path/filepath"
	"strings"
)

type Media int

const (
	Audio Media = iota
	Iframe
	Image
	Pdf
	Text
	Unknown
	Video
)

var TABLE = []struct {
	contentType string
	media       Media
	extensions  []string
}{
	{"application/json", Text, []string{"json"}},
	{"application/pdf", Pdf, []string{"pdf"}},
	{"application/pgp-signature", Text, []string{"asc"}},
	{"application/yaml", Text, []string{"yaml", "yml"}},
	{"audio/flac", Audio, []string{"flac"}},
	{"audio/mpeg", Audio, []string{"mp3"}},
	{"audio/wav", Audio, []string{"wav"}},
	{"image/apng", Image, []string{"apng"}},
	{"image/avif", Image, []string{}},
	{"image/gif", Image, []string{"gif"}},
	{"image/jpeg", Image, []string{"jpg", "jpeg"}},
	{"image/png", Image, []string{"png"}},
	{"image/svg+xml", Iframe, []string{"svg"}},
	{"image/webp", Image, []string{"webp"}},
	{"model/gltf-binary", Unknown, []string{"glb"}},
	{"model/stl", Unknown, []string{"stl"}},
	{"text/html;charset=utf-8", Iframe, []string{"html"}},
	{"text/plain;charset=utf-8", Text, []string{"txt"}},
	{"video/mp4", Video, []string{"mp4"}},
	{"video/webm", Video, []string{"webm"}},
}

func ContentTypeForPath(path string) (string, error) {
	extension := filepath.Ext(path)
	if len(extension) == 0 {
		return "", errors.New("file must have extension")
	}
	extension = strings.TrimPrefix(extension, ".")
	extension = strings.ToLower(extension)

	// todo mp4

	for _, entry := range TABLE {
		for _, ext := range entry.extensions {
			if ext == extension {
				return entry.contentType, nil
			}
		}
	}

	var extensions []string
	for _, entry := range TABLE {
		extensions = append(extensions, entry.extensions...)
	}
	return "", errors.New("unsupported file extension `." + extension + "`, supported extensions: " + strings.Join(extensions, " "))
}
