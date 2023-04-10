terraform {
  required_providers {
    cause = {
      version = "1.0"
      source  = "github.com/puni9869/generator"
    }
  }
}


provider "generator" {}

resource "generator" "random" {
  number = "10"
}
