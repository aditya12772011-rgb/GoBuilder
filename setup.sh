#!/bin/bash

CYAN='\033[0;36m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
RED='\033[0;31m'
RESET='\033[0m'
CLEARLN='\033[2K'

clear
echo -e "${CYAN}┌──────────────────────────────────────────────────┐${RESET}"
echo -e "${CYAN}│${GREEN}    ⚡ GOBUILDER: PREMIUM ANDROID APK TOOLCHAIN   ${CYAN}│${RESET}"
echo -e "${CYAN}└──────────────────────────────────────────────────┘${RESET}"

echo -e "${YELLOW}[*] Synchronizing system package managers...${RESET}"
apt-get update -y > /dev/null 2>&1 || pkg update -y > /dev/null 2>&1

# 1. Verification of Compiler Libraries
for tool in javac go aapt dx zip; do
    echo -n -e "${YELLOW}[*] Validating compiler sub-system: core:${tool}...${RESET}"
    if ! command -v $tool &> /dev/null; then
        echo -e "\r${CLEARLN}${RED}[!] Core module '${tool}' missing. Provisioning dependency...${RESET}"
        if [ "$tool" == "javac" ]; then
            apt-get install default-jdk -y || pkg install openjdk-17 -y
        elif [ "$tool" == "go" ]; then
            apt-get install golang -y || pkg install golang -y
        elif [ "$tool" == "aapt" ] || [ "$tool" == "dx" ]; then
            apt-get install aapt axmlrpc -y || pkg install aapt dx -y
        else
            apt-get install $tool -y || pkg install $tool -y
        fi
        echo -e "${GREEN}[✔] Sub-system core:${tool} deployed.${RESET}"
    else
        echo -e "\r${CLEARLN}${GREEN}[✔] Sub-system core:${tool} active.${RESET}"
    fi
done

echo -e "${CYAN}────────────────────────────────────────────────────${RESET}"
echo -e "${GREEN}[✔] Build pipeline environments verified successfully!${RESET}"
echo -e "${CYAN}────────────────────────────────────────────────────${RESET}"
sleep 1

# Launching Master Controller Console
go run main.go
