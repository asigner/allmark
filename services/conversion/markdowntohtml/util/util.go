// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"github.com/andreaskoch/allmark2/model"
	"io"
	"strings"
)

func GetFallbackLink(title, path string) string {
	return fmt.Sprintf(`<a href="%s" target="_blank" title="%s">%s</a>`, path, title, title)
}

func WriteFileContent(file *model.File, writer io.Writer) (string, error) {

	// get the file content
	contentProvider := file.ContentProvider()
	if err := contentProvider.Data(writer); err != nil {
		return "", err
	}

	// get the mime type
	contentType, err := contentProvider.MimeType()
	if err != nil {
		return "", err
	}

	return contentType, nil
}

func IsImageFile(file *model.File) bool {
	mimetype, err := GetMimeType(file)
	if err != nil {
		return false
	}

	return strings.HasPrefix(mimetype, "image/")
}

func IsTextFile(file *model.File) bool {
	mimetype, err := GetMimeType(file)
	if err != nil {
		return false
	}

	return strings.HasPrefix(mimetype, "text/")
}

func IsAudioFile(file *model.File) bool {
	mimetype, err := GetMimeType(file)
	if err != nil {
		return false
	}

	return strings.HasPrefix(mimetype, "audio/")
}

func IsVideoFile(file *model.File) bool {
	mimetype, err := GetMimeType(file)
	if err != nil {
		return false
	}

	return strings.HasPrefix(mimetype, "video/")
}

func IsPDFFile(file *model.File) bool {
	mimetype, err := GetMimeType(file)
	if err != nil {
		return false
	}

	return mimetype == "application/pdf"
}

func GetMimeType(file *model.File) (string, error) {
	contentProvider := file.ContentProvider()
	mimetype, err := contentProvider.MimeType()
	if err != nil {
		return "", err
	}

	return mimetype, nil
}
