package dto

import "mime/multipart"

type CreateFileRequest struct {
	NamespaceID int64                 `form:"namespaceID"`
	FilePath    string                `form:"filePath"`
	File        *multipart.FileHeader `form:"file"`
}
