package utils_file

import (
	"path/filepath"
	"slices"
	"strings"
)

// FileTypeMap 定义文件类型及其对应的扩展名
var FileTypeMap = map[string][]string{
	"image":      {"jpg", "jpeg", "png", "gif", "bmp", "tiff", "ico", "webp"},
	"video":      {"mp4", "avi", "mov", "wmv", "flv", "mkv", "webm"},
	"audio":      {"mp3", "wav", "aac", "flac", "m4a", "ogg", "wma"},
	"file":       {"doc", "docx", "xls", "xlsx", "ppt", "pptx", "pdf", "txt"},
	"archive":    {"zip", "rar", "7z", "tar", "gz"},
	"executable": {"exe", "apk", "app", "deb", "rpm", "pkg", "msi"},
	"disk":       {"iso", "dmg"},
}

// GetFileTypes 根据媒体类型返回对应的文件扩展名列表
func GetFileTypes(mediaType string) []string {
	if extensions, ok := FileTypeMap[mediaType]; ok {
		return extensions
	}
	return []string{}
}

// GetMediaType 根据文件扩展名判断媒体类型
func GetMediaType(extension string) string {
	ext := strings.ToLower(extension)
	for mediaType, extensions := range FileTypeMap {
		if slices.Contains(extensions, ext) {
			return mediaType
		}
	}
	return "other"
}

// IsValidFileType 检查文件类型是否合法
func IsValidFileType(filename string, allowedTypes []string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	if ext == "" {
		return false
	}
	ext = strings.TrimPrefix(ext, ".")

	for _, mediaType := range allowedTypes {
		if extensions, ok := FileTypeMap[mediaType]; ok {
			if slices.Contains(extensions, ext) {
				return true
			}
		}
	}
	return false
}
