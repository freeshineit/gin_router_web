package models

import "mime/multipart"

// ChunkFile file chunk
type ChunkFile struct {
	Name   string                `json:"name" form:"name"`
	Chunk  int                   `json:"chunk" form:"chunk"`
	Chunks int                   `json:"chunks" form:"chunks"`
	File   *multipart.FileHeader `json:"file" form:"file"`
}
