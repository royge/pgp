version=1.0.0
provider_path=registry.terraform.io/royge/pgp/$(version)

build:
	go build -o terraform-provider-pgp_$(version)

# MacOS installation
install:
	mkdir -p ~/Library/Application\ Support/io.terraform/plugins/$(provider_path)/darwin_arm64
	cp terraform-provider-pgp_$(version) ~/Library/Application\ Support/io.terraform/plugins/$(provider_path)/darwin_arm64
