version=1.0.0
provider_path=registry.terraform.io/royge/age/$(version)

build:
	go build -o terraform-provider-age_$(version)

# MacOS installation
install:
	mkdir -p ~/Library/Application\ Support/io.terraform/plugins/$(provider_path)/darwin_arm64
	cp terraform-provider-age_$(version) ~/Library/Application\ Support/io.terraform/plugins/$(provider_path)/darwin_arm64

encrypt:
	age -a -o secret.txt.age -R recipients.txt secret.txt
