package web

import (
	"fmt"
	"net/http"
)

//HTTPWriter http response writer
type HTTPWriter struct {
	writer http.ResponseWriter
}

//Write http response write
func (h *HTTPWriter) Write(p []byte) (n int, err error) {
	return h.writer.Write(p)
}

//WriteAt s3 writer対応
func (h *HTTPWriter) WriteAt(p []byte, _ int64) (n int, err error) {
	return h.writer.Write(p)
}

//WriteFileHeader ヘッダーにファイル名を出力
func (h HTTPWriter) WriteFileHeader(filename string) {
	h.writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	h.writer.Header().Set("Cache-Control", "no-store")
	h.writer.Header().Set("Content-Type", "application/octet-stream")
}

//NewHTTPWriter 初期化
func NewHTTPWriter(writer http.ResponseWriter) *HTTPWriter {
	return &HTTPWriter{writer: writer}
}
