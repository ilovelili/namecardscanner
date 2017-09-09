package naturallanguage

// DocumentType Document Type
type DocumentType string

const (
	// Unspeified type unspecified
	Unspeified DocumentType = "TYPE_UNSPECIFIED"
	// PlainText plain text
	PlainText DocumentType = "PLAIN_TEXT"
	// HTML html
	HTML DocumentType = "HTML"
)

// EncodingType encoding type
type EncodingType string

const (
	// None no encoding
	None EncodingType = "NONE"
	// Utf8 Utf8
	Utf8 EncodingType = "UTF8"
	// Utf16 Utf16
	Utf16 EncodingType = "UTF16"
	// Utf32 Utf32
	Utf32 EncodingType = "UTF32"
)
