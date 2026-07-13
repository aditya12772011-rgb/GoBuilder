l#!/bin/bash

CYAN='\033[0;36m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
RESET='\033[0m'

clear
echo -e "${CYAN}==================================================${RESET}"
echo -e "${GREEN}  go-builder engine // cyber hack tech${RESET}"
echo -e "${CYAN}==================================================${RESET}"

echo -e "${YELLOW}[*] Preparing system dependencies...${RESET}"
apt-get update -y > /dev/null 2>&1 || pkg update -y > /dev/null 2>&1
apt-get install golang zip aapt -y > /dev/null 2>&1 || pkg install golang zip aapt -y > /dev/null 2>&1

echo -e "${GREEN}[✔] Setup finished. Launching main engine...${RESET}"
sleep 1

go run main.go
