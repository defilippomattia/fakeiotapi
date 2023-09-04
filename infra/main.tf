provider "aws" {
    region = "eu-central-1"
    #AWS_ACCESS_KEY and AWS_SECRET_KEY are set as environment variables
}

resource "aws_vpc" "fake_iot" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = "vpc-prod-fake-iot"
    Region = "eu-central-1"
    Environment = "prod"
    Project = "fake-iot"
    CreatedBy = "Terraform"
  }
}

resource "aws_subnet" "fake_iot_1" {
  cidr_block = "10.0.0.0/24"
  vpc_id = aws_vpc.fake_iot.id
  availability_zone = "eu-central-1a"
  map_public_ip_on_launch = true

  tags = {
    Name = "subnet-1-prod-fake-iot"
    Region = "eu-central-1"
    Environment = "prod"
    Project = "fake-iot"
    CreatedBy = "Terraform"
    Public = "true"
    AvailabilityZone = "eu-central-1a"
  }
}

resource "aws_internet_gateway" "fake_iot" {
    vpc_id = aws_vpc.fake_iot.id

  tags = {
    Name = "igw-prod-fake-iot"
    Region = "eu-central-1"
    Environment = "prod"
    Project = "fake-iot"
    CreatedBy = "Terraform"
  }
}

resource "aws_route_table" "fake_iot" {
    vpc_id = aws_vpc.fake_iot.id


  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.fake_iot.id
  }

  tags = {
    Name = "rt-prod-fake-iot"
    Region = "eu-central-1"
    Environment = "prod"
    Project = "fake-iot"
    CreatedBy = "Terraform"
  }
}

resource "aws_route_table_association" "subnet_association" {
  subnet_id      = aws_subnet.fake_iot_1.id
  route_table_id = aws_route_table.fake_iot.id
}

resource "aws_security_group" "ssh" {
  name        = "ssh-security-group"
  description = "Allow SSH traffic"
    vpc_id = aws_vpc.fake_iot.id


  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

    tags = {
    Name = "ssh-security-group"
  }

}

resource "aws_security_group" "http" {
  name        = "http-security-group"
  description = "Allow HTTP traffic"
  vpc_id      = aws_vpc.fake_iot.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "http-security-group"
  }
}

resource "aws_security_group" "https" {
  name        = "https-security-group"
  description = "Allow HTTPS traffic"
  vpc_id      = aws_vpc.fake_iot.id

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "https-security-group"
  }
}

resource "aws_instance" "fake_iot" {
  ami           = "ami-04e601abe3e1a910f" 
  instance_type = "t2.small" #try with micro
  subnet_id     = aws_subnet.fake_iot_1.id
  key_name      = "fake-iot-key"
  security_groups = [
    aws_security_group.ssh.id,
    aws_security_group.http.id,
    aws_security_group.https.id
  ]

  root_block_device {
    volume_size = 12 
    delete_on_termination = true  
  }

    user_data = <<-EOF
              #!/bin/bash
              sudo apt update
              sudo apt upgrade -y
              sudo apt install -y docker.io docker-compose
              sudo snap install core
              sudo snap refresh core
              sudo snap install --classic certbot
              sudo -u ubuntu mkdir /home/ubuntu/myapp
              sudo -u ubuntu mkdir /home/ubuntu/mycerts
              EOF

  tags = {
    Name = "ec2-prod-fake-iot"
    Region = "eu-central-1"
    Environment = "prod"
    Project = "fake-iot"
    CreatedBy = "Terraform"
  }
}


resource "aws_eip" "fake_iot" {
  instance = aws_instance.fake_iot.id
  domain   = "vpc"

  tags = {
    Name = "eip-prod-fake-iot"
    Region = "eu-central-1"
    Environment = "prod"
    Project = "fake-iot"
    CreatedBy = "Terraform"
  }
}


resource "aws_route53_zone" "fake_iot" {
  name = "fakeiotapi.xyz"
  tags = {
    Name = "fakeiotapi.xyz"
    Region = "eu-central-1"
    Environment = "prod"
    Project = "fake-iot"
    CreatedBy = "Terraform"
  }
}

resource "aws_route53_record" "fake_iot_root" {
  zone_id = aws_route53_zone.fake_iot.zone_id
  name    = aws_route53_zone.fake_iot.name # fakeiotapi.xyz
  type    = "A"
  ttl     = "300"
  records = [aws_eip.fake_iot.public_ip]
}

resource "aws_route53_record" "fake_iot_www" {
  zone_id = aws_route53_zone.fake_iot.zone_id
  name    = "www.${aws_route53_zone.fake_iot.name}" # www.fakeiotapi.xyz
  type    = "CNAME"
  ttl     = "300"
  records = [aws_eip.fake_iot.public_ip]
}