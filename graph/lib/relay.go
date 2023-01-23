package lib

import (
	"encoding/base64"
	"strings"
)

type GlobalId struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func ToGlobalID(nodeType string, id string) string {
	str := nodeType + ":" + id
	globalId := base64.StdEncoding.EncodeToString([]byte(str))
	return globalId
}

func FromGlobalId(globalID string) *GlobalId {
	str := ""
	gid, _ := base64.StdEncoding.DecodeString(globalID)
	str = string(gid)
	tokens := strings.Split(str, ":")
	return &GlobalId{
		Type: tokens[0],
		ID:   tokens[1],
	}
}
