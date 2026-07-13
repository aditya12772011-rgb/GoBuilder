#!/bin/bash

CYAN='\033[0;36m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
RED='\033[0;31m'
RESET='\033[0m'
CLEARLN='\033[2K'

clear
echo -e "${CYAN}[ go-builder v3.7 ] // cyber hack tech${RESET}\n"

echo -e "${YELLOW}[*] Validating repositories...${RESET}"
apt-get update -y > /dev/null 2>&1 || pkg update -y > /dev/null 2>&1

for tool in javac go aapt dx zip curl; do
    if ! command -v $tool &> /dev/null; then
        echo -e "${RED}[!] Installing core module: ${tool}...${RESET}"
        if [ "$tool" == "javac" ]; then
            apt-get install default-jdk -y || pkg install openjdk-17 -y
        elif [ "$tool" == "go" ]; then
            apt-get install golang -y || pkg install golang -y
        elif [ "$tool" == "aapt" ] || [ "$tool" == "dx" ]; then
            apt-get install aapt axmlrpc -y || pkg install aapt dx -y
        else
            apt-get install $tool -y || pkg install $tool -y
        fi
    fi
done

TARGET_DIR="/data/data/com.termux/files/usr/share/aapt"
mkdir -p "$TARGET_DIR"
if [ ! -f "$TARGET_DIR/android.jar" ]; then
    echo -e "${YELLOW}[*] Downloading android.jar architecture...${RESET}"
    curl -L -f --retry 3 -o "$TARGET_DIR/android.jar" "https://raw.githubusercontent.com/xeffyr/termux-apk-repository/master/packages/android.jar" > /dev/null 2>&1
fi

echo -e "${GREEN}[✔] System environment verified.${RESET}"
sleep 1
go run main.go
