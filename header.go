package rpm

//go:generate stringer -type SignatureTag,TagType -output header_string.go

type SignatureTag int32
type TagType int32

const (
	Size SignatureTag = iota + 1000
	LEMD5_1
	PGP
	LEMD5_2
	MD5
	GPG
	PGP5
	PayloadSize
	ReservedSpace

	BadSHA1_1       SignatureTag = 264
	BadSHA1_2       SignatureTag = 265
	SHA1            SignatureTag = 269
	DSA             SignatureTag = 267
	RSA             SignatureTag = 268
	LongSize        SignatureTag = 270
	LongArchiveSize SignatureTag = 271
)

const (
	Null TagType = iota
	Char
	Int8
	Int16
	Int32
	Int64
	String
	Bin
	StringArray
	I18nString

	Min       TagType = 0
	Max       TagType = 9
	ForceFree TagType = 0xff
	MaskType  TagType = 0x000ffff
)
