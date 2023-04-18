#!/bin/bash

# Build program
go build -o program

# Copy file binary ke dalam EC2 instance
scp -i C:\devi\MSIB\go_Devi-Amalia\26_Compute-Service\Praktikum\devi.pem program ubuntu@ec2-3-106-126-85.ap-southeast-2.compute.amazonaws.com:/home/user/program

# Copy file konfigurasi ke dalam EC2 instance
scp -i  C:\devi\MSIB\go_Devi-Amalia\26_Compute-Service\Praktikum\devi.pem config.yml ubuntu@ec2-3-106-126-85.ap-southeast-2.compute.amazonaws.com:/home/user/program


# Restart service di dalam EC2 instance
ssh -i C:\devi\MSIB\go_Devi-Amalia\26_Compute-Service\Praktikum\devi.pem ubuntu@ec2-3-106-126-85.ap-southeast-2.compute.amazonaws.com:/home/user/program
 'sudo systemctl restart myservice.service'
