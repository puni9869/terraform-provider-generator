# Terraform Provider Generator

Run the following command to build the provider

```shell
go get github.com/hashicorp/terraform-plugin-sdk/v2
go build -o terraform-provider-generator
```

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
export TF_LOG=VERBOSE # optional
cd terraform && terraform init && terraform apply
```

[More](https://developer.hashicorp.com/terraform/plugin)
