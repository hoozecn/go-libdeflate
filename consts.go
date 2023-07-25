package libdeflate

import "github.com/hoozecn/go-libdeflate/v2/native"

// These constants specify several special compression levels
const (
	MinCompressionLevel        = native.MinCompressionLevel
	MaxStdZlibCompressionLevel = native.MaxStdZlibCompressionLevel
	MaxCompressionLevel        = native.MaxCompressionLevel
	DefaultCompressionLevel    = native.DefaultCompressionLevel
)
