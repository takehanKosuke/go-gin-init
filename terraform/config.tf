variable "aws_access_key" {}
variable "aws_secret_key" {}

provider "aws" {
  version    = "~> 3.0"
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region     = "ap-northeast-1"
}
