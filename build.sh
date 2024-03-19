#!/usr/bin/env bash

check_os() {
    . /etc/os-release
    case $ID in
        ubuntu)
        sudo apt install wget curl docker docker.io -y
        ;;
        arch)
        sudo pacman -S docker wget curl
        ;;
        debian)
        sudo apt install wget curl docker docker.io -y
        ;;
        centos)
        sudo yum install -y yum-utils && sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo && sudo yum install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
        ;;
    esac
}


check_os
docker compose up -d postgres
docker compose up -d

