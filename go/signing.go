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

	"github.com/golang-jwt/jwt/v5"
)

// Replace with your own values
const (
	embedPath   = "your path here"
	clientID    = "your clientid here"
	embedSecret = "your secret here"
)

func generateJWTEmbedURl() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":          "xyz@xyz.com",
		"jti":          fmt.Sprintf("%x", rand.Int63()),
		"iat":          time.Now().Unix(),
		"exp":          time.Now().Add(time.Hour * 1).Unix(),
		"iss":          clientID,
		"ver":          "1.1",
		"aud":          "sigmacomputing",
		"teams":        [...]string{"EmbedTeam"},
		"account_type": "Pro",
	})

	tokenString, err := token.SignedString([]byte(embedSecret))

	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("https://app.sigmacomputing.com/<your org>/<your workbook>?:embed=true&:jwt=%s", tokenString)
	return url
}

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

func secureEmbedUrl() string {
	params := map[string]interface{}{
		":nonce":              fmt.Sprintf("%x", rand.Int63()),
		":email":              "xyz@xyz.com",
		":external_user_id":   "xyz@xyz.com",
		":client_id":          clientID,
		":time":               time.Now().Unix(),
		":session_length":     3600,
		":mode":               "userbacked",
		":external_user_team": "EmbedTeam",
		":account_type":       "embedUser",
		// custom controls/parameters
		// "Store-Region": "West",
	}

	urlWithParams := embedPath + "?" + urlencode(params)

	hmacHash := hmac.New(sha256.New, []byte(embedSecret))
	io.WriteString(hmacHash, urlWithParams)
	signature := hex.EncodeToString(hmacHash.Sum(nil))

	urlWithSignature := urlWithParams + "&" + urlencode(map[string]interface{}{":signature": signature})

	return urlWithSignature
}

func main() {
	fmt.Println("=========JWT Embed URL=========")
	fmt.Println(generateJWTEmbedURl())
	fmt.Println("==================================")
	fmt.Println("=========Secure Embed URL=========")
	fmt.Println(secureEmbedUrl())
	fmt.Println("==================================")
}
