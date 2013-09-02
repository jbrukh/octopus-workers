package resources

import (
	"io"
	"os"
)

func GetFile(path string) (r io.Reader, err error) {
	return os.Open(path)
}
