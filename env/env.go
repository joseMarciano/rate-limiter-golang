package env

import (
	"fmt"
	"math"
	"os"
	"rate-limiter/internal/model"
	"strconv"
)

var configMap = make(map[string]any)

const ipConfigLimitRate = "IP_CONFIG_LIMIT_RATE"
const ipLockedTime = "IP_LOCKED_TIME"

const apiKeyConfigLimitRate = "API_KEY_CONFIG_LIMIT_RATE"
const apiKeyLockedTime = "API_KEY_LOCKED_TIME"

func GetConfigLimitRate(typeClient model.TypeClient) int {
	if model.Ip == typeClient {
		return getIpConfigLimitRate()
	}

	return getApiKeyConfigLimitRate()
}

func GetLockedTimeLimitRate(typeClient model.TypeClient) int {
	if model.Ip == typeClient {
		return getIpLockedTime()
	}

	return getApiKeyLockedTime()
}

func getIpConfigLimitRate() int {

	if value, ok := configMap[ipConfigLimitRate]; ok {
		return value.(int)
	}

	number, err := strconv.Atoi(os.Getenv(ipConfigLimitRate))
	if err != nil {
		return math.MaxInt
	}

	configMap[ipConfigLimitRate] = number
	fmt.Printf("getting number %s", number)
	return number
}

func getIpLockedTime() int {

	if value, ok := configMap[ipLockedTime]; ok {
		return value.(int)
	}

	number, err := strconv.Atoi(os.Getenv(ipLockedTime))
	if err != nil {
		return math.MinInt
	}

	configMap[ipLockedTime] = number
	return number
}

func getApiKeyConfigLimitRate() int {

	if value, ok := configMap[apiKeyConfigLimitRate]; ok {
		return value.(int)
	}

	number, err := strconv.Atoi(os.Getenv(apiKeyConfigLimitRate))
	if err != nil {
		return math.MaxInt
	}

	configMap[apiKeyConfigLimitRate] = number
	return number
}

func getApiKeyLockedTime() int {

	if value, ok := configMap[apiKeyLockedTime]; ok {
		return value.(int)
	}

	number, err := strconv.Atoi(os.Getenv(apiKeyLockedTime))
	if err != nil {
		return math.MinInt
	}

	configMap[apiKeyLockedTime] = number
	return number
}
