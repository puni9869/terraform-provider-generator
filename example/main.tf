terraform {
  required_providers {
    generator = {
      version = "1.0"
      source  = "github.com/puni9869/generator"
    }
  }
}


provider "generator" {}

resource "generator" "random" {
  number = "10"
}

data "generator" "random" {
  number = "10"
}

output "val" {
  value = data.generator.random
}
