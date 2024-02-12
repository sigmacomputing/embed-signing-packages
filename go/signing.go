package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

// Replace with your own values
const (
	embedPath   = "your path here"
	clientID    = "your clientid here"
	embedSecret = "your secret here"
)

func myQuote(val string) string {
	val = strings.ReplaceAll(val, " ", "%20")
	return val
}

func urlencode(pairs map[string]interface{}) string {
	var encodedParams []string
	for k, v := range pairs {
		encodedParams = append(encodedParams, myQuote(k)+"="+myQuote(fmt.Sprint(v)))
	}
	return strings.Join(encodedParams, "&")
}

func main() {
	params := map[string]interface{}{
		":nonce":              fmt.Sprintf("%x", rand.Int63()),
		":email":              "test@test.com",
		":external_user_id":   "test@test.com",
		":client_id":          clientID,
		":time":               time.Now().Unix(),
		":session_length":     3600,
		":mode":               "userbacked",
		":external_user_team": "Embedded Users,EmbeddingTown",
		":account_type":       "embedUser",
		// custom controls/parameters
		// "Store-Region": "West",
	}

	urlWithParams := embedPath + "?" + urlencode(params)

	hmacHash := hmac.New(sha256.New, []byte(embedSecret))
	io.WriteString(hmacHash, urlWithParams)
	signature := hex.EncodeToString(hmacHash.Sum(nil))

	urlWithSignature := urlWithParams + "&" + urlencode(map[string]interface{}{":signature": signature})

	fmt.Println(urlWithSignature)
}
