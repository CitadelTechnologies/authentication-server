package security

import(
    "testing"
)

func TestGenerateRandomToken(t *testing.T) {
    token := GenerateRandomToken(12)

    if len(token) != 24 {
        t.Errorf("Token length was incorrect, got %d bytes, want %d bytes", len(token), 24)
    }
}
