# ----------------------------------------------------------------
#
# Deploy AWS instances
#
# Code assumes that a key pair called "mimosa" exists in the region already
#
# Instance lifetime is 1 day but can be changed below
#
# ----------------------------------------------------------------

# ----------------------------------------------------------------
#
# DEPLOY - Specify the region/total of your choice:
#
# terraform apply -auto-approve -var total=2 -var awsregion=eu-west-1 -state=eu-west-1.tfstate
#
# MAKE SURE YOU SET THE STATE FILE TO MATCH YOUR REGION!
#
# ----------------------------------------------------------------

# ----------------------------------------------------------------
#
# TEAR DOWN - Specify the region/total of your choice:
#
# terraform destroy -auto-approve -var total=2 -var awsregion=eu-west-1 -state=eu-west-1.tfstate
#
# TRULY, MAKE SURE YOU SET THE STATE FILE TO MATCH YOUR REGION!
#
# ----------------------------------------------------------------

provider "aws" {
  region = "eu-west-1"
}

provider "hiera" {
}

variable "workspace" {
  type = string
}

data "hiera_lookup" "ws1" {
  workspace = var.workspace
}

resource "aws_instance" "test" {
  ami           = "ami-005af4c3162f495fa"
  instance_type = data.hiera_lookup.ws1.instance_type
  key_name      = "mimosa"
  tags = {
    mimosa   = "true"
    lifetime = "1d"
    Name     = "mimosa-${data.hiera_lookup.ws1.instance_name}"
  }
}
