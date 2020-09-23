# resource "aws_vpc" "tVPC" {
#   cidr_block           = "10.1.0.0/16"
#   instance_tenancy     = "default"
#   enable_dns_support   = "true"
#   enable_dns_hostnames = "false"
#   tags                 = { Name = var.project_name }
# }

# resource "aws_internet_gateway" "tGW" {
#   vpc_id = aws_vpc.tVPC.id
# }

# resource "aws_subnet" "public-a" {
#   vpc_id            = aws_vpc.tVPC.id
#   cidr_block        = "10.1.1.0/24"
#   availability_zone = "ap-northeast-1a"
# }

# resource "aws_route_table" "public-route" {
#   vpc_id = aws_vpc.tVPC.id
#   route {
#     cidr_block = "0.0.0.0/0"
#     gateway_id = aws_internet_gateway.tGW.id
#   }
# }

# resource "aws_route_table_association" "puclic-a" {
#   subnet_id      = aws_subnet.public-a.id
#   route_table_id = aws_route_table.public-route.id
# }

# resource "aws_security_group" "tSecurityGroup" {
#   name        = var.security_group_name
#   description = "Allow SSH inbound traffic"
#   vpc_id      = aws_vpc.tVPC.id
#   ingress {
#     from_port   = 22
#     to_port     = 22
#     protocol    = "tcp"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
#   ingress {
#     from_port   = 80
#     to_port     = 80
#     protocol    = "tcp"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
#   egress {
#     from_port   = 0
#     to_port     = 0
#     protocol    = "-1"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
# }

# data aws_ssm_parameter amzn2_ami {
#   name = "/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2"
# }

# resource "aws_instance" "tEC2" {
#   ami           = data.aws_ssm_parameter.amzn2_ami.value
#   instance_type = "t2.micro"
#   key_name      = var.aws_pem
#   vpc_security_group_ids = [
#     aws_security_group.tSecurityGroup.id
#   ]
#   subnet_id                   = aws_subnet.public-a.id
#   associate_public_ip_address = "true"
#   tags                        = { Name = var.instance_name }
# }
