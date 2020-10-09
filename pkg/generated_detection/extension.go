package generated_detection

import (
	"strings"
)

func (d detector) hasExtension(ext string) bool {
	return strings.ToLower(d.FileName.Ext) == ext
}
