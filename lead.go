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
	LeadTypeBinary LeadType = iota // 0
	LeadTypeSource                 // 1
)

const (
	_                     LeadArch = iota
	LeadArchI386_x86_64            // 1
	LeadArchAlpha_Sparc64          // 2
	LeadArchSparc                  // 3
	LeadArchMIPS                   // 4
	LeadArchPowerPC                // 5
	LeadArch68000                  // 6
	LeadArchSGI                    // 7
	LeadArchRS6000                 // 8
	LeadArchIA64                   // 9
	LeadArchSparc64                // 10
	LeadArchMIPSel                 // 11
	LeadArchARM                    // 12
	LeadArchMiNT                   // 13
	LeadArchS_390                  // 14
	LeadArchS_390x                 // 15
	LeadArchPowerPC64              // 16
	LeadArchSuperH                 // 17
	LeadArchXtensa                 // 18
	LeadArchNoarch        LeadArch = 255
)

const (
	_             LeadOS = iota
	LeadOSLinux          // 1
	LeadOSIrix           // 2
	LeadOSSolaris        // 3
	LeadOSSunOS          // 4
	AmigaOS              // 5
	HP_UX                // 6
	OSF1                 // 7
	FreeBSD              // 8
	SCO_SV               // 9
	Irix64               // 10
	NextStep             // 11
	BSD_OS               // 12
	machten              // 13
	CYGWIN32_NT          // 14
	CYGWIN32_95          // 15
	UNIX_SV              // 16
	MiNT                 // 17
	OS_390               // 18
	VM_ESA               // 19
	Linux_390            // 20
	Darwin               //	21
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
