package apiv1

// GenerateKubernetesAuthTokenRPC is the path for the GenerateKubernetesAuthToken RPC.
const GenerateKubernetesAuthTokenRPC = "kubernetes/auth/tokens"

// GenerateKubernetesAuthTokenRequest is the request message for the GenerateKubernetesAuthToken RPC.
type GenerateKubernetesAuthTokenRequest struct {
	Username string   `json:"username"`
	Groups   []string `json:"groups"`
}

// GenerateKubernetesAuthTokenResponse is the response message for the GenerateKubernetesAuthToken RPC.
type GenerateKubernetesAuthTokenResponse struct {
	Token string `json:"token"`
}
