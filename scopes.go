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
)
