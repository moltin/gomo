package core

import (
	"bytes"

	"github.com/moltin/gomo/form"
)

// FileUploadRequest represents a request to upload a file to Moltin
type FileUploadRequest struct {
	File   bytes.Buffer `json:"file"`
	Public bool         `json:"public,omitempty"`
}

// File is a Moltin File - https://docs.moltin.com/advanced/files
type File struct {
	ID       string    `json:"id,omitempty" form:"-"`
	Type     string    `json:"type" form:"-"`
	FileName string    `json:"file_name" form:"file_name,omitempty"`
	Public   bool      `json:"public" form:"public"`
	MimeType string    `json:"mime_type" form:"mime_type,omitempty"`
	FileSize int       `json:"file_size" form:"file_size,omitempty"`
	Meta     *FileMeta `json:"meta,omitempty" form:"-"`
	Link     *struct {
		Href string `json:"href"`
	} `json:"link" form:"-"`
	Links *Links     `json:"links,omitempty" form:"-"`
	File  *form.File `json:"-" form:"file"`
}

// FileMeta represents the meta object for a moltin file entity
type FileMeta struct {
	Dimensions struct {
		Width  int32 `json:"width"`
		Height int32 `json:"height"`
	} `json:"dimensions"`
	Timestamps Timestamps `json:"timestamps"`
}
