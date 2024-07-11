package auth

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGenerateToken(t *testing.T) {
	tk := GenerateToken("strect", "eric")
	spew.Dump(tk)
	panic(1)
}
