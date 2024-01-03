package src

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwk"
)

func ValidateToken() {
	// Replace with your actual JWKS URL
	jwksURL := "https://cognito-idp.ap-southeast-2.amazonaws.com/ap-southeast-2_EozV71Rz4/.well-known/jwks.json"

	// Sample token string (replace with your actual token)
	tokenString := "eyJraWQiOiJacjQxbmRIeWE0OUJseU5ta0hYeWRvU0N2SmowckNsWUxvUXkzelNtRWlrPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiJiYTQyZTkzYi1iNzA3LTRhN2QtYjI2MS0wN2ZlM2JkZGI4ZmIiLCJjb2duaXRvOmdyb3VwcyI6WyJMdGpscWYxVmFHIiwiOEF4ejVTYVRjcSIsIjRMbW9aQUxkOWkiLCJkZXYiXSwiZW1haWxfdmVyaWZpZWQiOnRydWUsImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC5hcC1zb3V0aGVhc3QtMi5hbWF6b25hd3MuY29tXC9hcC1zb3V0aGVhc3QtMl9Fb3pWNzFSejQiLCJjdXN0b206dXNlcl9pZCI6IjhmZjQzNzI1LWQ0NDItNDZmZS05ZmMzLTk2YzA1ODU2ZGY1MyIsInBob25lX251bWJlcl92ZXJpZmllZCI6ZmFsc2UsImNvZ25pdG86dXNlcm5hbWUiOiJiYTQyZTkzYi1iNzA3LTRhN2QtYjI2MS0wN2ZlM2JkZGI4ZmIiLCJvcmlnaW5fanRpIjoiNjRkYTU5YTQtM2M5NC00OGJmLTk4YzItODEyMjk4NmI4OWVmIiwiYXVkIjoiNGxzbzQ5MTB1aml0OHJwbXNyaDQxa2MybmYiLCJldmVudF9pZCI6ImY3MWViMzg2LTk2MTUtNGJhMS1hNmIzLTg0YmZhNjgwYzEwNiIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNzAzMDUyMDUyLCJuYW1lIjoibGFtbXYiLCJwaG9uZV9udW1iZXIiOiIrODQzNTg0NDMyNTQiLCJleHAiOjE3MDMwNjM0NjcsImlhdCI6MTcwMzA1OTg2NywianRpIjoiMmJiZjllMzItZDkyNS00MWQ0LWFlZWItMjk4NzQ3YjQ0MzYwIiwiZW1haWwiOiJsYW1tdiszQGtvem8tamFwYW4uY29tIn0.mZD3O9iKD0tIszfWg0kGb1NJbUY0stoAGvNQqG74TiFDaxJ-QNDDmei4S2JGNxj-eapHCXJwSxSkZHi59fdAn0zJxJnaGMVJ0asqDhtWW4D_d3xnjak76vL55Iu8pC9BqohWTFkkDzpybtZ5tYeb4eFTwFVbfM8ZtSrW94OFBMdzGmvePcnt1fuu8p6Gy3V3WR50RaPsHZ1Nkr7wAP6ZwZatnn0saTTJ_E4MYre6BJ4sUEieorrjYHuf3QdbRetO-6MMIfYaENJrYT22K5OmYk1Zs6QBjolQkbBfZMpw4vPHQU3ha3q9JNbmgldjxUwmcO7glXJJf-B39tF0lQnmLg"
	// tokenString := "eyJraWQiOiJJTXNyOVRhaHlwanhzVTVlb3lQdlpYUUlhZEFaOFQzMVk1QjJtWkVkaFBjPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiJjNzI0NmEwOC03MGIxLTcwMmEtYmE1YS01ZjU5YzE4ZDY1MTUiLCJjb2duaXRvOmdyb3VwcyI6WyJ0ZXN0YWRtaW4iXSwiZW1haWxfdmVyaWZpZWQiOnRydWUsImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC5hcC1ub3J0aGVhc3QtMS5hbWF6b25hd3MuY29tXC9hcC1ub3J0aGVhc3QtMV85S2RpQzl1VWkiLCJjdXN0b206dXNlcl9pZCI6ImYwZjY4MTQ4LTBhZjQtNDBhOS05N2MxLWIzOTQyZTExNTNjMSIsImNvZ25pdG86dXNlcm5hbWUiOiJjNzI0NmEwOC03MGIxLTcwMmEtYmE1YS01ZjU5YzE4ZDY1MTUiLCJvcmlnaW5fanRpIjoiNDE2NDE5ZjAtNzgyNS00NzE5LWFkNmUtZGY1Y2NiN2ZhMjk0IiwiYXVkIjoiNm9lbXQ4OGF2MW45ZnN1b25mOW1laGZ0dDciLCJldmVudF9pZCI6ImM0ZDFmYmFiLWUxMTEtNGEzNy04OGQ5LTgyM2UxYjcwOGE3ZiIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNzAzMDYwMDQ4LCJuYW1lIjoibGFtbXYiLCJleHAiOjE3MDMwNjM2NDgsImlhdCI6MTcwMzA2MDA0OCwianRpIjoiMGRiZGRlMjctZDFiNS00NDgwLWI3MTEtNjdkNjBkMjY1ZjMwIiwiZW1haWwiOiJsYW1tdisyQGtvem8tamFwYW4uY29tIn0.eP_cHkX7WoA8zqZ1p9P_xuZ-twQ2Iikp4OdVO4b8UhJwvY7U_LXGUW6Hkg7eYT8fgpv5hIPAUyUUhBbl_gyabdlVZszzUedKY_p3u5h5NIRoNTdTyClmQeEzlSnJ-LvBAjIjBpu1tKre7GGNOAFn1MDyUjLvsnngyyrEjouykU_MLNoTvEjodrjLFj08JVQrIPCDdevBYPe5rH8EWSJ9DCf4sw_ZzImdBRaZUf0B1vjONXc9KJbS1KBvihzWzdYFseHXYpkorMx4Gsk9P-VavyyemtThRpqz3Jm-bJqvSBo8QBIIA6tZJ0hZbKG3XKJRhxjJcBvFA5brCboXqWQ1yQ"
	ctx := context.Background()
	// Fetch JWKS
	keySet, err := jwk.Fetch(ctx, jwksURL)
	if err != nil {
		fmt.Println("Error fetching JWKS:", err)
		return
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("kid header not found")
		}
		key, ok := keySet.LookupKeyID(kid)
		if !ok {
			return nil, fmt.Errorf("key with specified kid is not present in JWKS")
		}

		var publicKey interface{}
		err := key.Raw(&publicKey)
		if err != nil {
			return nil, fmt.Errorf("could not parse public key")
		}
		return publicKey, nil
	})

	if err != nil {
		fmt.Println("Error validating token:", err)
		return
	}

	fmt.Println("Token is valid:", token.Valid)

	// Access decoded claims
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println("Claims:", claims)
}
