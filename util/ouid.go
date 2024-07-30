package util

import (
	"github.com/google/uuid"
	"strings"
)

// GenerateOuid 生成不带-的uuid
func GenerateOuid() string {
	uuidWithDashes := uuid.New().String()
	return strings.ReplaceAll(uuidWithDashes, "-", "")
}
