package serializers

type LoginResponse struct {
	AuthToken string `json:"auth_token"`
	// RefreshToken string `json:"refresh_token"`
}
