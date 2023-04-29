package age_test

import (
	"os"
	"testing"

	"github.com/royge/terraform-provider-age/age"
)

func TestProvider(t *testing.T) {
	t.Parallel()

	provider := age.Provider()

	if err := provider.InternalValidate(); err != nil {
		t.Fatalf("error validating provider: %v", err)
	}
}

func TestDecrypt(t *testing.T) {
	textFile := "../secret.txt"
	cipherFile := "../secret.txt.age"
	privateKey := os.Getenv("AGE_PRIVATE_KEY")

	res, err := age.Decrypt(cipherFile, privateKey)
	if err != nil {
		t.Fatalf("error decrypting file: %v", err)
	}

	b, err := os.ReadFile(textFile)
	if err != nil {
		t.Fatalf("error reading original text file: %v", err)
	}
	
	if string(b) != res {
		t.Errorf("want decrypt() = '%v', got '%v'", string(b), res)
	}
}
