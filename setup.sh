#!/bin/bash

# रङहरू (Colors)
CYAN='\033[0;36m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
RED='\033[0;31m'
RESET='\033[0m'
CLEARLN='\033[2K'

clear
# तिम्रो च्यानलको नामसहितको सिम्पल र सफा ब्यानर
echo -e "${CYAN}==================================================${RESET}"
echo -e "${GREEN}  🚀 GoBuilder APK Engine // cyber hack tech${RESET}"
echo -e "${CYAN}==================================================${RESET}"

echo -e "${YELLOW}[*] Updating system repositories...${RESET}"
apt-get update -y > /dev/null 2>&1 || pkg update -y > /dev/null 2>&1

# आवश्यक टुलहरू चेक र इन्स्टल गर्ने
for tool in javac go aapt dx zip curl; do
    echo -n -e "${YELLOW}[*] Checking: ${tool}...${RESET}"
    if ! command -v $tool &> /dev/null; then
        echo -e "\r${CLEARLN}${RED}[!] Installing missing tool: '${tool}'...${RESET}"
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
        echo -e "\r${CLEARLN}${GREEN}[✔] Verified: ${tool}${RESET}"
    fi
done

# एन्ड्रोइड लाइब्रेरी (android.jar) फिक्स गर्ने बलियो लजिक
TARGET_DIR="/data/data/com.termux/files/usr/share/aapt"
mkdir -p "$TARGET_DIR"

# यदि पहिले नै बिग्रेको फाइल छ भने त्यसलाई हटाउने
if [ -f "$TARGET_DIR/android.jar" ]; then
    rm -f "$TARGET_DIR/android.jar"
fi

echo -e "${YELLOW}[*] Downloading core Android SDK component (android.jar)...${RESET}"
# यस पटक -L -f --retry फ्ल्याग थपिएको छ ताकि डाउनलोड बीचमा नबिग्रियोस्
curl -L -f --retry 3 -o "$TARGET_DIR/android.jar" "https://raw.githubusercontent.com/xeffyr/termux-apk-repository/master/packages/android.jar"

echo -e "${CYAN}--------------------------------------------------${RESET}"
echo -e "${GREEN}[✔] Environment Setup Complete! Running main engine...${RESET}"
echo -e "${CYAN}--------------------------------------------------${RESET}"
sleep 1

# मेन प्रोग्राम रन गर्ने
go run main.go
