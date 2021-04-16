package configStruct

type CustomConfigStruct struct {
	Auth0Enabled bool	`valid:"-" toml:"AUTH0_ENABLED" json:"auth0_enabled"`
	Domain       string	`valid:"-" toml:"DOMAIN"json:"domain"`
	ClientId	 string	`valid:"-" toml:"CLIENT_ID"json:"client_id"`
	Audience     string	`valid:"-" toml:"AUDIENCE"json:"audience"`
	// -- separation -- //
	DelegatedEnabled bool	`valid:"-" toml:"DELEGATED_ENABLED" json:"delegated_enabled"`
	DelegatedSigningUrl	 string	`valid:"-" toml:"DELEGATED_SIGNING_URL"json:"delegated_signing_url"`
	Callback     string	`valid:"-" toml:"CALLBACK_URL"json:"callback_url"`
}

// type AuthConfigStruct struct {
	
// }
// type DelegatedConfigStruct struct {
	
// }