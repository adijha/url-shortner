package helpers

import (
	"strings"
)

// func GetRateLimit(r *RedisClient) int {
// 	return r.Get(database.Ctx, "rate_limit").Val().(int)
// }
// func GetRateLimitReset(r *RedisClient) time.Duration {
// 	return r.Get(database.Ctx, "rate_limit_reset").Val().(time.Duration)
// }

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == "localhost" {
		return false
	}

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	if newURL == "localhost" {
		return false
	}

	return true
}
