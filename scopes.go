package scopes

type (
	Scope string
)

func (s Scope) String() string {
	return string(s)
}

// Scopes, add your scopes here
const (
	// Auth scopes
	ScopeAuthClientRead  Scope = "auth.client:read"
	ScopeAuthClientWrite Scope = "auth.client:write"

	// User scopes
	ScopeUserRead           Scope = "user:read"
	ScopeUserWrite          Scope = "user:write"
	ScopeUserWorkspaceRead  Scope = "user.workspace:read"
	ScopeUserWorkspaceWrite Scope = "user.workspace:write"

	// Wallet scopes
	ScopeWalletRead  Scope = "wallet:read"
	ScopeWalletWrite Scope = "wallet:write"
	// Wallet User scopes
	ScopeWalletUserRead  Scope = "wallet.user:read"
	ScopeWalletUserWrite Scope = "wallet.user:write"
	// Wallet Settings scopes
	ScopeWalletSettingRead  Scope = "wallet.setting:read"
	ScopeWalletSettingWrite Scope = "wallet.setting:write"
	// Wallet Transaction scopes
	ScopeWalletTransactionRead  Scope = "wallet.transaction:read"
	ScopeWalletTransactionWrite Scope = "wallet.transaction:write"
	// Wallet Webhook scopes
	ScopeWalletWebhookRead  Scope = "wallet.webhook:read"
	ScopeWalletWebhookWrite Scope = "wallet.webhook:write"
)
