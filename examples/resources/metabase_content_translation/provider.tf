terraform {
  required_providers {
    metabase = {
      source  = "flovouin/metabase"
      version = "~> 0.9"
    }
  }
}

provider "metabase" {
  endpoint = var.metabase_endpoint
  username = var.metabase_username
  password = var.metabase_password
}

variable "metabase_endpoint" {
  description = "Metabase endpoint URL"
  type        = string
}

variable "metabase_username" {
  description = "Metabase username"
  type        = string
}

variable "metabase_password" {
  description = "Metabase password"
  type        = string
  sensitive   = true
}
