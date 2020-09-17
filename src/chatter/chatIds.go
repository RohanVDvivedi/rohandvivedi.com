package chatter

import (
	"rohandvivedi.com/src/randstring"
	"strings"
)

const CHAT_USER_ID_PREFIX  = "CHAT_USER-"
const CHAT_GROUP_ID_PREFIX = "CHAT_GRUP-"
const CHAT_CONNECTION_ID_PREFIX = "CHAT_CONN-"

const unique_id_part_length = 64;

func GetNewId(prefix string) string {
	return prefix + randstring.GetRandomString(unique_id_part_length)
}

/* functions for connections id */
func GetNewChatConnectionId() string {
	return GetNewId(CHAT_CONNECTION_ID_PREFIX)
}

func IsChatConnectionId(Id string) bool {
	return strings.HasPrefix(Id, CHAT_CONNECTION_ID_PREFIX)
}

/* functions for user id */
func GetNewChatUserId() string {
	return GetNewId(CHAT_USER_ID_PREFIX)
}

func IsChatUserId(Id string) bool {
	return strings.HasPrefix(Id, CHAT_USER_ID_PREFIX)
}

/* functions for group id */
func GetNewChatGroupId() string {
	return GetNewId(CHAT_GROUP_ID_PREFIX)
}

func IsChatGroupId(Id string) bool {
	return strings.HasPrefix(Id, CHAT_GROUP_ID_PREFIX)
}

/* functions for server id */
func IsChatManagerId(Id string) bool {
	return strings.HasPrefix(Id, "server")
}