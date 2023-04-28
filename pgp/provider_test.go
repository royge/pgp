package pgp_test

import (
	"testing"

	"github.com/royge/terraform-provider-pgp/pgp"
)

func TestProvider(t *testing.T) {
	t.Parallel()

	provider := pgp.Provider()

	if err := provider.InternalValidate(); err != nil {
		t.Fatalf("error validating provider: %v", err)
	}
}
