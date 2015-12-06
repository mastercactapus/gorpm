package rpm

import (
	"archive/zip"
)

type RPM struct {
	Lead      Lead
	Signature Header
	Header    Header
}
