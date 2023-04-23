# Pre-required knowledge
1. HCL (HashiCorp Configuration Language)
2. Terraform basic level understanding.
3. Modules in Terraform.
4. Golang

# What is terraform?
Terraform is an infrastructure as code tool that lets you build, change, and version cloud and on-prem resources safely and efficiently.

# What is Infrastructure as Code (IaC)?
Infrastructure as Code (IaC) is the managing and provisioning of infrastructure through code instead of through manual processes.

# How terraform is related to IaC?
Terraform is a tool which gives us the flexibility to manage and provision the resources that are required to spin up any service.

# Component of terraform?
It has two important component
1. Terraform core --> Manages resource plan, graph, states, configuration.
2. Terraform plugin --> Responsible for defining resources for specific services.


# On what we concentrate today
Terraform plugin


# About Terraform SDK (for creating plugins)
1. Written in golang.
2. Manages the schema of resource, data and import.
3. Terraform provider also known as Terraform Plugin.
4. End format is executable binaries as a provider.
5. Terraform registry is the place from where it gets downloaded.


# How to provider looks like in terraform code.
```hcl
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }
  required_version = ">= 1.2.0"
}
```

Inside the `required_providers` block we see all the `providers` a.k.a plugins. Here `aws` is a hashicorp terraform provider for `AWS Cloud`.
So today we will learn how to create our own provider like `aws`.



# Terraform blocks
1. Resources:-
   Use `resource` blocks to define components of your infrastructure. A `resource` might be a physical or virtual component such as an EC2 instance, or it can be a logical resource such as a Heroku application. Resource blocks contain arguments which you use to configure the resource. Arguments can include things like machine sizes, disk image names, or VPC IDs. Our providers reference lists the required and optional arguments for each resource.

    ```hcl
    resource "aws_instance" "web" {
      ami           = "ami-a1b2c3d4"
      instance_type = "t2.micro"
    }
    ```


2. DataSources:-
   Use `data` block to define component of infrastructure. A `data` is already created resources, contains the arguments like config, filter_conditions, depends_on etc. Data sources enable Terraform to use the information determined outside of Terraform, represented by different separate Terraform states or code, or changed by functions. In other words, Cloud infrastructure, applications, and services transmit data, which Terraform can query and perform managing data sources.

    ```hcl
    data "aws_ami" "example" {
      most_recent = true
    
      owners = ["self"]
      tags = {
        Name   = "app-server"
        Tested = "true"
      }
    }
    ```

3. Imports:-
   Terraform can import existing infrastructure resources. This functionality lets you bring existing resources under Terraform management.
   Terraform import can only import resources into the state. Importing does not generate configuration. Before you run terraform import you must manually write a resource configuration block for the resource. The resource block describes where Terraform should map the imported object.

    ```
    terraform import aws_instance.foo i-abcd1234
    ```

# Lets build the provider

We ganna be building a random generator provider.
1. Clone the repository
    ```sh
    git clone https://github.com/puni9869/terraform-provider-generator.git
    ```

2. Install golang from https://go.dev/doc/install

3. Change into the cloned repository.
    ```sh
    cd terraform-provider-generator
    ```

4. The directory should have the following structure.

    ```sh
    tree -L 3
    
    
    .
    |-- LICENSE
    |-- Makefile
    |-- README.md
    |-- generator
    |   |-- data_randomgenerator.go
    |   |-- provider.go
    |   `-- resource_randomgenerator.go
    |-- go.mod
    |-- go.sum
    |-- main.go
    `-- example
        `-- main.tf
    
    3 directories, 11 files
    
    ``` 

5. Download the dependencies.
    ```sh
    go mod tidy
    go mod verify
    ```

6. Build the binary
   We are using make to build the binary of our provider.
    ```sh
    make build
    ```

7. Install the provider in local system.
    ```sh
    make install
    
    # Output
    go build -o terraform-provider-generator
    mkdir -p ~/.terraform.d/plugins/github.com/puni9869/generator/1.0/darwin_amd64
    mv terraform-provider-generator ~/.terraform.d/plugins/github.com/puni9869/generator/1.0/darwin_amd64
    ```

   Note:- We are not publishing the provider on terraform registry. Terraform query the local cache directory `~/.terraform.d/plugins` to download the provider.


8. Go to `example` folder:
    ```bash
    cd example
    terraform init
    
    # Output
    
    Initializing the backend...
    
    Initializing provider plugins...
    - Finding github.com/puni9869/generator versions matching "1.0.0"...
    - Installing github.com/puni9869/generator v1.0.0...
    - Installed github.com/puni9869/generator v1.0.0 (unauthenticated)
    
    Terraform has created a lock file .terraform.lock.hcl to record the provider
    selections it made above. Include this file in your version control repository
    so that Terraform can guarantee to make the same selections by default when
    you run "terraform init" in the future.
    
    Terraform has been successfully initialized!
    
    You may now begin working with Terraform. Try running "terraform plan" to see
    any changes that are required for your infrastructure. All Terraform commands
    should now work.
    
    If you ever set or change modules or backend configuration for Terraform,
    rerun this command to reinitialize your working directory. If you forget, other
    commands will detect it and remind you to do so if necessary.
    
    ```

   Directory structure:
    ```bash
    tree -a
    .
    |-- .terraform
    |   `-- providers
    |       `-- github.com
    |           `-- puni9869
    |               `-- generator
    |                   `-- 1.0.0
    |                       `-- darwin_amd64 -> /Users/puni9869/.terraform.d/plugins/github.com/puni9869/generator/1.0/darwin_amd64
    |-- .terraform.lock.hcl
    |-- main.tf
    
    
    8 directories, 2 files
    
    ```


9. Terraform plan
    ```bash
    terraform plan
    
    #o/p
    Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following
    symbols:
      + create
    
    Terraform will perform the following actions:
    
      # generator.random will be created
      + resource "generator" "random" {
          + id     = (known after apply)
          + number = "10"
        }
    
    Plan: 1 to add, 0 to change, 0 to destroy.
    
    Changes to Outputs:
      + val = {
          + id     = "1681202668"
          + number = "10"
        }
    
    ─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
    
    Note: You didn't use the -out option to save this plan, so Terraform can't guarantee to take exactly these actions if you run "terraform
    apply" now.
    
    ```

10. Terraform apply:-
    ```bash
    terraform apply
    
    # o/p
    Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following
    symbols:
      + create
    
    Terraform will perform the following actions:
    
      # generator.random will be created
      + resource "generator" "random" {
          + id     = (known after apply)
          + number = "10"
        }
    
    Plan: 1 to add, 0 to change, 0 to destroy.
    
    Changes to Outputs:
      + val = {
          + id     = "1681202723"
          + number = "10"
        }
    
    Do you want to perform these actions?
      Terraform will perform the actions described above.
      Only 'yes' will be accepted to approve.
    
      Enter a value: yes
    
    generator.random: Creating...
    generator.random: Creation complete after 0s [id=["3419777e-d845-11ed-906b-325096b39f47","34197d00-d845-11ed-a0e5-325096b39f47","34197d96-d845-11ed-8b36-325096b39f47","34197e04-d845-11ed-87b6-325096b39f47","34197e68-d845-11ed-b566-325096b39f47","34197ecc-d845-11ed-843c-325096b39f47","34197f30-d845-11ed-9cee-325096b39f47","34197f94-d845-11ed-b9f0-325096b39f47","34197fee-d845-11ed-b907-325096b39f47","3419805c-d845-11ed-88dc-325096b39f47"]]
    
    Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
    
    Outputs:
    
    val = {
      "id" = "1681202723"
      "number" = "10"
    }
    
    ```

11. Checking the state  `terraform.tfstate`:
    ```json
    {
      "version": 4,
      "terraform_version": "1.0.1",
      "serial": 1,
      "lineage": "b616fc98-f56c-8d24-fdb3-e856c1956cc8",
      "outputs": {
        "val": {
          "value": {
            "id": "1681202723",
            "number": "10"
          },
          "type": [
            "object",
            {
              "id": "string",
              "number": "string"
            }
          ]
        }
      },
      "resources": [
        {
          "mode": "data",
          "type": "generator",
          "name": "random",
          "provider": "provider[\"github.com/puni9869/generator\"]",
          "instances": [
            {
              "schema_version": 0,
              "attributes": {
                "id": "1681202723",
                "number": "10"
              },
              "sensitive_attributes": []
            }
          ]
        },
        {
          "mode": "managed",
          "type": "generator",
          "name": "random",
          "provider": "provider[\"github.com/puni9869/generator\"]",
          "instances": [
            {
              "schema_version": 0,
              "attributes": {
                "id": "[\"3419777e-d845-11ed-906b-325096b39f47\",\"34197d00-d845-11ed-a0e5-325096b39f47\",\"34197d96-d845-11ed-8b36-325096b39f47\",\"34197e04-d845-11ed-87b6-325096b39f47\",\"34197e68-d845-11ed-b566-325096b39f47\",\"34197ecc-d845-11ed-843c-325096b39f47\",\"34197f30-d845-11ed-9cee-325096b39f47\",\"34197f94-d845-11ed-b9f0-325096b39f47\",\"34197fee-d845-11ed-b907-325096b39f47\",\"3419805c-d845-11ed-88dc-325096b39f47\"]",
                "number": "10"
              },
              "sensitive_attributes": [],
              "private": "bnVsbA=="
            }
          ]
        }
      ]
    }
    
    ```
    At level `resources.[*].instances.[*].id`  these are generated values from the provider.


# Let's understand the plugin code in `generator` directory.

1. `provider.go` This file contains the provider defination of `ResourceBlock`, `DataSoures`, `Imports` and provider's name. In our case `generate` is the name of our provider.

   Here we can define in our terraform block
    ```hcl
    terraform {
      required_providers {
        generator = {
          version = "1.0"
          source  = "github.com/puni9869/generator"
        }
      }
    }
    
    
    provider "generator" {}
    ```

2. `data_randomgenerator.go` This file contains the logic for `data` sources block.
   As below, we can define in terraform
    ```hcl
    data "generator" "random" {
      number = "10"
    }
    ```

3. `resource_randomgenerator.go` this file contains the logic for `resource` block. Resource block having create, read, update and destroy functions which are responsible for managing resource.

   As below,  we can define in terraform
    ```hcl
    resource "generator" "random" {
      number = "10"
    }
    ```

   Inside terraform we have `number` property which is an attribute in terraform `resource` and `data` block.
   In go code this is mapped with `Schema: map[string]*schema.Schema ` where we define all the attributes and their characteristics like type, default value etc. [More](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider).


