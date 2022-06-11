package constants

import util "whm-api/utils"

const SessionUserIdKey = "session_user_id"

func BasePath() string {
	return util.GodotEnv("BASE_PATH")
}
