terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.25.2"
    }
  }
  backend "s3" {
    bucket = "ftf-tf-state"
    key    = "user-subgraph/terraform.tfstate"
    region = "us-east-1"
  }
}

variable "do_token" {
  description = "digital ocean access token"
  type        = string
}
variable "env" {
  description = "environment name"
  type        = string
  default     = "main"
}

provider "digitalocean" {
  # Configuration options
  token = var.do_token
}

resource "digitalocean_database_cluster" "mongodb-example" {
  name       = "ftf-${var.env}-cluster"
  engine     = "mongodb"
  version    = "5"
  size       = "db-s-1vcpu-1gb"
  region     = "nyc3"
  node_count = 1
}
