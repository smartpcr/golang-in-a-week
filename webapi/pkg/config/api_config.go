package config

type AuthType string

const (
	AuthTypeNone        AuthType = "none"
	AuthTypeJwt         AuthType = "jwt"
	AuthTypeCertificate AuthType = "certificate"
)

var authTypeMap = map[string]AuthType{
	"none":        AuthTypeNone,
	"jwt":         AuthTypeJwt,
	"certificate": AuthTypeCertificate,
}

type ApiConfig struct {
	Port     int
	AuthType AuthType
}

func GetApiConfig() *ApiConfig {
	authType := getEnv("API_AUTH_TYPE", "none")
	apiAuthType, ok := authTypeMap[authType]
	if !ok {
		apiAuthType = AuthTypeNone
	}
	return &ApiConfig{
		Port:     getEnv("API_PORT", 8080),
		AuthType: apiAuthType,
	}
}
