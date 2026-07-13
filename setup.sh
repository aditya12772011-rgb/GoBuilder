#!/bin/bash

CYAN='\033[0;36m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
RED='\033[0;31m'
RESET='\033[0m'
CLEARLN='\033[2K'

clear
echo -e "${CYAN}====================================================${RESET}"
echo -e "${GREEN} 🛠️  GOBUILDER: AUTOMATED APK PRODUCTION TOOLCHAIN   ${RESET}"
echo -e "${CYAN}====================================================${RESET}"

echo -e "${YELLOW}[*] Updating system package repositories...${RESET}"
apt-get update -y > /dev/null 2>&1 || pkg update -y > /dev/null 2>&1

# 1. Java JDK Installation Check
echo -n -e "${YELLOW}[*] Verifying Java Compiler (javac)...${RESET}"
if ! command -v javac &> /dev/null; then
    echo -e "\r${CLEARLN}${RED}[!] Java JDK missing. Installing in background...${RESET}"
    apt-get install default-jdk -y > /dev/null 2>&1 || pkg install openjdk-17 -y > /dev/null 2>&1
    echo -e "${GREEN}[✔] Java JDK configured successfully.${RESET}"
else
    echo -e "\r${CLEARLN}${GREEN}[✔] Java JDK already verified.${RESET}"
fi

# 2. Go Installation Check
echo -n -e "${YELLOW}[*] Verifying Go Language Runtime...${RESET}"
if ! command -v go &> /dev/null; then
    echo -e "\r${CLEARLN}${RED}[!] Installing Golang...${RESET}"
    apt-get install golang -y > /dev/null 2>&1 || pkg install golang -y > /dev/null 2>&1
    echo -e "${GREEN}[✔] Golang configured successfully.${RESET}"
else
    echo -e "\r${CLEARLN}${GREEN}[✔] Go Language Runtime verified.${RESET}"
fi

# 3. Android AAPT2 Check
echo -n -e "${YELLOW}[*] Verifying Android AAPT2 Tools...${RESET}"
if ! command -v aapt2 &> /dev/null; then
    echo -e "\r${CLEARLN}${RED}[!] Installing AAPT2 Build Asset tools...${RESET}"
    apt-get install aapt2 -y > /dev/null 2>&1 || pkg install aapt2 -y > /dev/null 2>&1
    echo -e "${GREEN}[✔] AAPT2 tools configured successfully.${RESET}"
else
    echo -e "\r${CLEARLN}${GREEN}[✔] AAPT2 tools already verified.${RESET}"
fi

# 4. Zip Tool Check
if ! command -v zip &> /dev/null; then
    apt-get install zip -y > /dev/null 2>&1 || pkg install zip -y > /dev/null 2>&1
fi

echo -e "${CYAN}----------------------------------------------------${RESET}"
echo -e "${GREEN}[✔] Environment Ready! Launching GoBuilder Engine...${RESET}"
echo -e "${CYAN}====================================================${RESET}"
sleep 1

# Automatically launch the main interactive module
go run main.go
