package common

import (
	"cmp"
	"os"
	"strconv"
)

func convertStrToInt(reqAllowedEnv string) int {
	reqAllowed, err := strconv.Atoi(reqAllowedEnv)
	if err != nil || reqAllowed < 0 {
		return 0
	}

	return reqAllowed
}

var MaxNoOfRequestAllowed = cmp.Or(convertStrToInt(os.Getenv("MAX_REQUEST_ALLOWED")), 50)
const CampaignActiveStatus = "ACTIVE"