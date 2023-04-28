terraform {
  required_providers {
    pgp = {
      version = "~> 1.0.0"
      source  = "royge/pgp"
    }
  }
}
resource "pgp_cipher" "secret_cipher" {
  filename = "secret.txt.gpg"
}
