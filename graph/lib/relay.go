package lib

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////
// Global IDs
// ////////////////////////////////////////////////////////////////////////////

type GlobalId struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func ToGlobalID(nodeType string, id string) string {
	str := nodeType + ":" + id
	globalId := base64.StdEncoding.EncodeToString([]byte(str))
	return globalId
}

func FromGlobalId(globalID string) (*GlobalId, error) {
	str := ""
	gid, _ := base64.StdEncoding.DecodeString(globalID)
	str = string(gid)
	tokens := strings.Split(str, ":")
	if len(tokens) != 2 {
		return nil, fmt.Errorf("invalid global id %s", globalID)
	}
	return &GlobalId{
		Type: tokens[0],
		ID:   tokens[1],
	}, nil
}
