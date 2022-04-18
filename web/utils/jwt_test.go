package utils

import "testing"

func TestGenerateToken(t *testing.T) {
	t.Run("generate token", func(t *testing.T) {
		GenerateToken("c9emmgvkobjld07qlm0g")
	})
}
