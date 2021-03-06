// generated by stringer -type SignatureTag,TagType -output header_string.go; DO NOT EDIT

package rpm

import "fmt"

const (
	_SignatureTag_name_0 = "BadSHA1_1BadSHA1_2"
	_SignatureTag_name_1 = "DSARSASHA1LongSizeLongArchiveSize"
	_SignatureTag_name_2 = "SizeLEMD5_1PGPLEMD5_2MD5GPGPGP5PayloadSizeReservedSpace"
)

var (
	_SignatureTag_index_0 = [...]uint8{0, 9, 18}
	_SignatureTag_index_1 = [...]uint8{0, 3, 6, 10, 18, 33}
	_SignatureTag_index_2 = [...]uint8{0, 4, 11, 14, 21, 24, 27, 31, 42, 55}
)

func (i SignatureTag) String() string {
	switch {
	case 264 <= i && i <= 265:
		i -= 264
		return _SignatureTag_name_0[_SignatureTag_index_0[i]:_SignatureTag_index_0[i+1]]
	case 267 <= i && i <= 271:
		i -= 267
		return _SignatureTag_name_1[_SignatureTag_index_1[i]:_SignatureTag_index_1[i+1]]
	case 1000 <= i && i <= 1008:
		i -= 1000
		return _SignatureTag_name_2[_SignatureTag_index_2[i]:_SignatureTag_index_2[i+1]]
	default:
		return fmt.Sprintf("SignatureTag(%d)", i)
	}
}

const (
	_TagType_name_0 = "NullCharInt8Int16Int32Int64StringBinStringArrayI18nString"
	_TagType_name_1 = "ForceFree"
	_TagType_name_2 = "MaskType"
)

var (
	_TagType_index_0 = [...]uint8{0, 4, 8, 12, 17, 22, 27, 33, 36, 47, 57}
	_TagType_index_1 = [...]uint8{0, 9}
	_TagType_index_2 = [...]uint8{0, 8}
)

func (i TagType) String() string {
	switch {
	case 0 <= i && i <= 9:
		return _TagType_name_0[_TagType_index_0[i]:_TagType_index_0[i+1]]
	case i == 255:
		return _TagType_name_1
	case i == 65535:
		return _TagType_name_2
	default:
		return fmt.Sprintf("TagType(%d)", i)
	}
}
