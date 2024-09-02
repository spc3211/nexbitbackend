package util

import (
	"encoding/json"
	"math/rand"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberBytes = "1234567890"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // of letter indices fitting in 63 bits
)

// GenerateRandomString generates and returns a random alphanumeric string (with mixed case) of size n
func GenerateRandomString(n int, includeLetters, includeNumbers bool) string {
	charIncluded := ""
	if includeLetters {
		charIncluded += letterBytes
	}
	if includeNumbers {
		charIncluded += numberBytes
	}

	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(charIncluded) {
			sb.WriteByte(charIncluded[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return sb.String()
}

func GetStringPtr(s string) *string {
	return &s
}

func UnmarshalJson(data []byte, v any) error {
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	return nil
}

func GetOidcProviderID(clientIDPs map[string]map[string]string, client, idp string) string {
	if idps, ok := clientIDPs[client]; ok {
		if id, ok := idps[idp]; ok {
			return id
		}
	}
	return ""
}

func GetAppNameConfig(appNameConfigs map[string]string, appName string) string {
	if url, ok := appNameConfigs[appName]; ok {
		return url
	}
	return ""
}
