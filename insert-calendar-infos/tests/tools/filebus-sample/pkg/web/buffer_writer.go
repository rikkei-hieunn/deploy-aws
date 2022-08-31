package web

import (
	"bytes"
)

//BufferWriter バッファライター構成
type BufferWriter struct {
	buff bytes.Buffer
}

func (b *BufferWriter) Write(p []byte) (n int, err error) {
	return b.buff.Write(p)
}

//GetBytes Get bytes buffer
func (b *BufferWriter) GetBytes() []byte {
	return b.buff.Bytes()
}

//WriteAt s3 writer対応
func (b *BufferWriter) WriteAt(p []byte, _ int64) (n int, err error) {
	return b.buff.Write(p)
}

//NewBufferWriter 初期化
func NewBufferWriter() *BufferWriter {
	var buff bytes.Buffer
	return &BufferWriter{buff: buff}
}
