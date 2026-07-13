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

echo -e "${YELLOW}[*] Updating package repositories...${RESET}"
apt-get update -y > /dev/null 2>&1 || pkg update -y > /dev/null 2>&1

for tool in javac go aapt dx zip curl; do
    echo -n -e "${YELLOW}[*] Checking module: ${tool}...${RESET}"
    if ! command -v $tool &> /dev/null; then
        echo -e "\r${CLEARLN}${RED}[!] Core tool '${tool}' missing. Auto-deploying...${RESET}"
        if [ "$tool" == "javac" ]; then
            apt-get install default-jdk -y || pkg install openjdk-17 -y
        elif [ "$tool" == "go" ]; then
            apt-get install golang -y || pkg install golang -y
        elif [ "$tool" == "aapt" ] || [ "$tool" == "dx" ]; then
            apt-get install aapt axmlrpc -y || pkg install aapt dx -y
        else
            apt-get install $tool -y || pkg install $tool -y
        fi
    else
        echo -e "\r${CLEARLN}${GREEN}[✔] Module '${tool}' verified.${RESET}"
    fi
done

# Critical Dependency Setup
mkdir -p /data/data/com.termux/files/usr/share/aapt/
if [ ! -f /data/data/com.termux/files/usr/share/aapt/android.jar ]; then
    echo -e "${YELLOW}[*] Injecting missing android.jar development framework...${RESET}"
    curl -L -o /data/data/com.termux/files/usr/share/aapt/android.jar https://raw.githubusercontent.com/xeffyr/termux-apk-repository/master/packages/android.jar > /dev/null 2>&1
fi

echo -e "${CYAN}────────────────────────────────────────────────────${RESET}"
echo -e "${GREEN}[✔] Pipeline Environments Initialized Successfully!${RESET}"
echo -e "${CYAN}────────────────────────────────────────────────────${RESET}"
sleep 1

go run main.go
