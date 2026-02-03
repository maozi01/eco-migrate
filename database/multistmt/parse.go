// Package multistmt provides methods for parsing multi-statement database migrations
package multistmt

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// StartBufSize is the default starting size of the buffer used to scan and parse multi-statement migrations
var StartBufSize = 4096

// Handler handles a single migration parsed from a multi-statement migration.
// It's given the single migration to handle and returns whether or not further statements
// from the multi-statement migration should be parsed and handled.
type Handler func(migration []byte) bool

// Parse parses the given multi-statement migration
func Parse(reader io.Reader, delimiter []byte, maxMigrationSize int, h Handler) error {
	bufReader := bufio.NewReaderSize(reader, 10*1024*1024)
	var buf bytes.Buffer
	buf.Grow(8 * 1024)
	for {
		part, err := bufReader.ReadSlice(delimiter[len(delimiter)-1])
		if bytes.HasSuffix(part, delimiter) {
			buf.Write(part)
			sql := bytes.TrimSpace(buf.Bytes())
			buf.Reset()
			if len(sql) > 0 {
				if !h(sql) {
					return nil
				}
			}
		} else {
			buf.Write(part)
		}
		if err == bufio.ErrBufferFull {
			continue
		}
		if err == io.EOF {
			if buf.Len() > 0 {
				sql := bytes.TrimSpace(buf.Bytes())
				if len(sql) > 0 {
					h(sql)
				}
			}
			return nil
		}
		if err != nil {
			return err
		}
		if buf.Len() > maxMigrationSize {
			return fmt.Errorf("migration statement exceeds max size %d", maxMigrationSize)
		}
	}
}
