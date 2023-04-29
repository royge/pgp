terraform {
  required_providers {
    age = {
      version = "~> 1.0.0"
      source  = "royge/age"
    }
  }
}
resource "age_cipher" "secret_cipher" {
  filename = "secret.txt.age"
}

output "result" {
  value     = age_cipher.secret_cipher.result
  sensitive = true
}
