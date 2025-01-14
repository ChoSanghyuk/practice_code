package log

type Level int8

const ModuleKey = "module"
const ReqID = "reqID"

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
	NoLevel
	Disabled
)

const (
	TimeFormatUnix      = ""
	TimeFormatUnixMs    = "UNIXMS"
	TimeFormatUnixMicro = "UNIXMICRO"
	TimeFormatUnixNano  = "UNIXNANO"
	RFC3339             = "2006-01-02T15:04:05Z07:00"
)

const (
	DefaultFilepath   = "./"
	DefaultFilename   = "./log"
	DefaultMaxSize    = 100
	DefaultMaxBackups = 10
	DefaultLocalTime  = true
	DefaultCompress   = false
	DefaultRotate     = 60 * 60 * 24
)
