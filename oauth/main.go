package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const htmlIndex = `<html><body>
<a href="/GoogleLogin">Log in with Google</a>
</body></html>
`

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/GoogleCallback",
		ClientID:     os.Getenv("DWN_OAUTH_GOOGLE_KEY"),
		ClientSecret: os.Getenv("DWN_OAUTH_GOOGLE_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
	// Some random string, random for each request
	oauthStateString = "random"
)

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/GoogleLogin", handleGoogleLogin)
	http.HandleFunc("/GoogleCallback", handleGoogleCallback)
	fmt.Println("Now serving on port 3000!")
	fmt.Println(http.ListenAndServe(":3000", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()
	claims := &GoogleClaims{}
	err = json.NewDecoder(response.Body).Decode(claims)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, "Content: %+v\n", claims)
}

type GoogleClaims struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Link          string `json:"link"`
	Picture       string `json:"picture"`
	Gender        string `json:"gender"`
	Locale        string `json:"locale"`
}
