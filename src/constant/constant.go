package constant

import "time"

// middleware.
const (
	DefaultMdwHeaderToken      = "Authorization"
	DefaultMdwHeaderBearer     = "Bearer"
	DefaultMdwRateLimiter      = 20
	DefaultMdwSentryDebug      = true
	DefaultMdwSentrySampleRate = 1.0
	DefaultMdwTimeout          = 10 * time.Second
)

// pagination.
const (
	DefaultOrder = "created_at DESC"
	DefaultPage  = 1
	DefaultLimit = 10
)

const DefaultCacheExpireDuration = 24 * time.Hour

// file type.
var (
	FileTypeDocument = []string{".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".csv", ".json"}
	FileTypeImage    = []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".tif", ".raw", ".svg", ".webp", ".ico", ".eps", ".psd", ".heif", ".heic", ".hdr"}
	FileTypeVideo    = []string{".mp4", ".avi", ".mkv", ".wmv", ".mov", ".flv", ".webm", ".3gp", ".mpeg", ".mpg", ".divx", ".flv"}
	FileTypeAudio    = []string{".mp3", ".wav", ".aac", ".flac", ".ogg", ".wma", ".m4a", ".aiff", ".mp2", ".midi"}
)
