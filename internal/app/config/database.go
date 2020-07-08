package config

type Database struct {
	Firebase Firebase `yaml:"firebase"`
}

type Firebase struct {
	FirebaseAdmin FirebaseAdminSDKCred `yaml:"firebase_admin"`
	DatabaseURL   string               `yaml:"database_url"`
}

type FirebaseAdminSDKCred struct {
	Type          string `yaml:"type" json:"type" env:"FIREBASE_ADMIN_TYPE"`
	ProjectID     string `yaml:"project_id" json:"project_id" env:"FIREBASE_ADMIN_PROJECT_ID"`
	PrivateKeyID  string `yaml:"private_key_id" json:"private_key_id" env:"FIREBASE_ADMIN_PRIVATE_KEY_ID"`
	PrivateKey    string `yaml:"private_key" json:"private_key" env:"FIREBASE_ADMIN_PRIVATE_KEY"`
	ClientEmail   string `yaml:"client_email" json:"client_email" env:"FIREBASE_ADMIN_CLIENT_EMAIL"`
	ClientID      string `yaml:"client_id" json:"client_id" env:"FIREBASE_ADMIN_CLIENT_ID"`
	AuthURI       string `yaml:"auth_uri" json:"auth_uri" env:"FIREBASE_ADMIN_AUTH_URI"`
	TokenURL      string `yaml:"token_uri" json:"token_uri" env:"FIREBASE_ADMIN_TOKEN_URL"`
	AuthProvider  string `yaml:"auth_provider_x509_cert_url" json:"auth_provider_x509_cert_url" env:"FIREBASE_ADMIN_AUTH_PROVIDER"`
	ClientCertURL string `yaml:"client_x509_cert_url" json:"client_x509_cert_url" env:"FIREBASE_ADMIN_CLIENT_CERT_URL"`
}
