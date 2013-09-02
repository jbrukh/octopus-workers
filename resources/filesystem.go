package resources

import (
	"os"
)

func GetFile(path string) (r io.Reader, err error) {
	return os.Open(path)
}
