provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key      = var.private_key
  region           = var.region
}

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}


variable "private_key" {}
