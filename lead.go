package rpm

//go:generate stringer -type LeadType,LeadArch,LeadOS,LeadSignatureType -output lead_string.go

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

var ErrBadMagic = errors.New("invalid magic")

var LeadMagic = []byte{0xED, 0xAB, 0xEE, 0xDB}

type Lead struct {
	Magic         [4]byte
	Major, Minor  byte
	Type          LeadType
	Arch          LeadArch
	Name          [66]byte
	OS            LeadOS
	SignatureType LeadSignatureType
	Reserved      [16]byte
}

type LeadType int16
type LeadArch int16
type LeadOS int16
type LeadSignatureType int16

const (
	LeadSignatureTypeHeaderSignature LeadSignatureType = 5
)

const (
	LeadTypeBinary LeadType = 0
	LeadTypeSource LeadType = 1
)

const (
	LeadArchI386_x86_64   LeadArch = 1
	LeadArchAlpha_Sparc64 LeadArch = 2
	LeadArchSparc         LeadArch = 3
	LeadArchMIPS          LeadArch = 4
	LeadArchPowerPC       LeadArch = 5
	LeadArch68000         LeadArch = 6
	LeadArchSGI           LeadArch = 7
	LeadArchRS6000        LeadArch = 8
	LeadArchIA64          LeadArch = 9
	LeadArchSparc64       LeadArch = 10
	LeadArchMIPSel        LeadArch = 11
	LeadArchARM           LeadArch = 12
	LeadArchMiNT          LeadArch = 13
	LeadArchS_390         LeadArch = 14
	LeadArchS_390x        LeadArch = 15
	LeadArchPowerPC64     LeadArch = 16
	LeadArchSuperH        LeadArch = 17
	LeadArchXtensa        LeadArch = 18
	LeadArchNoarch        LeadArch = 255
)

const (
	LeadOSLinux   LeadOS = 1
	LeadOSIrix    LeadOS = 2
	LeadOSSolaris LeadOS = 3
	LeadOSSunOS   LeadOS = 4
	AmigaOS       LeadOS = 5
	HP_UX         LeadOS = 6
	OSF1          LeadOS = 7
	FreeBSD       LeadOS = 8
	SCO_SV        LeadOS = 9
	Irix64        LeadOS = 10
	NextStep      LeadOS = 11
	BSD_OS        LeadOS = 12
	machten       LeadOS = 13
	CYGWIN32_NT   LeadOS = 14
	CYGWIN32_95   LeadOS = 15
	UNIX_SV       LeadOS = 16
	MiNT          LeadOS = 17
	OS_390        LeadOS = 18
	VM_ESA        LeadOS = 19
	Linux_390     LeadOS = 20
	Darwin        LeadOS = 21
)

func (l *Lead) UnmarshalBinary(p []byte) error {
	if len(p) < 96 {
		return io.ErrShortBuffer
	}
	if !bytes.Equal(p[:4], LeadMagic) {
		return ErrBadMagic
	}
	b := bytes.NewBuffer(p[:96])
	return binary.Read(b, binary.BigEndian, l)
}
func (l Lead) GetName() string {
	i := bytes.IndexByte(l.Name[:], 0)
	return string(l.Name[:i])
}
