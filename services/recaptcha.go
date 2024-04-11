package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type RecaptchaService struct {
	apiKey  string
	siteKey string
}

func NewRecaptchaService(apiKey, siteKey string) *RecaptchaService {
	return &RecaptchaService{
		apiKey:  apiKey,
		siteKey: siteKey,
	}
}

type googleCaptchaResponse struct {
	Success     bool     `json:"success"`
	ChallengeTs string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes,omitempty"`
}

func (s *RecaptchaService) ValidateCaptcha(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return fmt.Errorf("parsing form: %w", err)
	}
	token := r.Form.Get("g-recaptcha-response")
	if token == "" {
		return fmt.Errorf("missing captcha token")
	}
	resp, err := http.PostForm(fmt.Sprintf("https://www.google.com/recaptcha/api/siteverify?%s", s.apiKey), url.Values{
		"siteKey":  {s.siteKey},
		"response": {token},
	})
	if err != nil {
		return fmt.Errorf("sending captcha request: %w", err)
	}
	defer resp.Body.Close()
	var response googleCaptchaResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return fmt.Errorf("decoding captcha response: %w", err)
	}
	return nil
}

func (s *RecaptchaService) Key() string {
	return s.siteKey
}
