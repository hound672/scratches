package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type AuthProfileType string

const (
	AuthProfileTypeHTTPBasic     AuthProfileType = "http_basic"
	AuthProfileTypeHTMLAutoForm  AuthProfileType = "html_auto_form"
	AuthProfileTypeHTMLFormBased AuthProfileType = "html_form_based"
	AuthProfileTypeRawCookie     AuthProfileType = "raw_cookie"
	AuthProfileTypeAPIKey        AuthProfileType = "api_key"
	AuthProfileTypeBearer        AuthProfileType = "bearer"
)

type AuthProfile struct {
	UUID     uuid.UUID
	Name     string
	Type     AuthProfileType
	Settings interface{}
}
type AuthProfileSettings interface {
	AuthProfileSettingsTypeHTTPBasic |
	AuthProfileSettingsTypeHTMLAutoForm |
	AuthProfileSettingsTypeHTMLFormBased |
	AuthProfileSettingsTypeRawCookie |
	AuthProfileSettingsTypeBearer |
	AuthProfileSettingsTypeAPIKey
}

func GetAuthProfileSettings[T AuthProfileSettings](profile *AuthProfile) (*T, error) {
	obj, ok := profile.Settings.(*T)
	if !ok {
		return nil, fmt.Errorf("can't cast settings to type %T", *new(T))
	}

	return obj, nil
}

type AuthProfileCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthProfileSettingsTypeHTTPBasic struct {
	Credentials AuthProfileCredentials `json:"credentials"`
	TestString  string                 `json:"test_string"`
}

type AuthProfileSettingsTypeHTMLAutoForm struct {
	Credentials   AuthProfileCredentials `json:"credentials"`
	FormURL       string                 `json:"form_url"`
	SuccessString string                 `json:"success_string"`
}

type AuthProfileSettingsTypeHTMLFormBased struct {
	Credentials     AuthProfileCredentials `json:"credentials"`
	FormURL         string                 `json:"form_url"`
	FormXPath       string                 `json:"form_x_path"`
	UsernameField   string                 `json:"username_field"`
	PasswordField   string                 `json:"password_field"`
	RegexpOfSuccess string                 `json:"regexp_of_success"`
	SubmitValue     string                 `json:"submit_value"`
}

type AuthProfileSettingsTypeRawCookie struct {
	Cookies         []string `json:"cookies"`
	SuccessURL      string   `json:"success_url"`
	RegexpOfSuccess string   `json:"regexp_of_success"`
}

type AuthProfileSettingsTypeBearer struct {
	Token         string `json:"token"`
	SuccessURL    string `json:"success_url"`
	SuccessString string `json:"success_string"`
}

type AuthProfileAPIKeyPlace string

const (
	AuthProfileAPIKeyPlaceCookie AuthProfileAPIKeyPlace = "cookie"
	AuthProfileAPIKeyPlaceHeader AuthProfileAPIKeyPlace = "header"
	AuthProfileAPIKeyPlaceQuery  AuthProfileAPIKeyPlace = "query"
)

type AuthProfileSettingsTypeAPIKey struct {
	Place         AuthProfileAPIKeyPlace `json:"place"`
	APIKey        string                 `json:"api_key"`
	SuccessURL    string                 `json:"success_url"`
	SuccessString string                 `json:"success_string"`
}


func main() {
	log.Printf("11")

	auth_profile_settings := "{\"UUID\":\"fda0725d-8935-46db-b1d3-4a915c9f7213\",\"ProductUUID\":\"4d4a37dc-aed5-4f45-9455-376fd2e623df\",\"Name\":\"DVWA Auto form\",\"Type\":\"html_auto_form\",\"Settings\":{\"credentials\":{\"username\":\"user\",\"password\":\"user\"},\"form_url\":\"http://we-dvwa.rd.ptsecurity.ru/login.php\",\"success_string\":\"http://we-dvwa.rd.ptsecurity.ru/index.php\"},\"CreatorUUID\":\"4373e2d0-250d-42ab-893e-4839dc3f90bc\",\"CreatorUsername\":\"asoc_user\",\"CreatedAt\":\"2024-06-06T06:06:53.777801Z\",\"DeletedAt\":null}"

	var authProfile *AuthProfile
	log.Printf("Before: %v", auth_profile_settings)

	msg := json.RawMessage{}
	authProfile = &AuthProfile{Settings: &msg}

	if err := json.Unmarshal([]byte(auth_profile_settings), &authProfile); err != nil {
		log.Fatalf("Error unmarshal res: %v", err)
	}

	log.Printf("After: %v", authProfile)

	switch authProfile.Type {
	case AuthProfileTypeHTMLAutoForm:
		var autoForm *AuthProfileSettingsTypeHTMLAutoForm
		if err := json.Unmarshal(msg, &autoForm); err != nil {
			log.Fatal(err)
		}
		authProfile.Settings = autoForm
	}

	autoForm, err := GetAuthProfileSettings[AuthProfileSettingsTypeHTMLAutoForm](authProfile)
	if err != nil {
		log.Fatalf(
			"activityTypes.GetAuthProfileSettings[activityTypes.AuthProfileSettingsTypeHTMLAutoForm]: %w", err,
		)
	}

	_ = autoForm
}
