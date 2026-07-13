# 🛠️ GoBuilder: Automated AI-Intent Android APK Toolchain

GoBuilder is a lightweight, ultra-fast, and automated build pipeline written completely in **Go** that parses user feature requirements into structural matrices and compiles them into real, installable Android APK artifacts directly from your terminal. 

No heavy Android Studio IDE or complex Gradle configurations required. Designed to work flawlessly in standard Linux environments and non-root **Termux** modules.

---

## 🚀 Key Features

* **Instant Dependency Management:** Automatically provisions Java JDK, Go, AAPT2, and archiving tools in the background without requiring `sudo` privileges.
* **Semantic Intent Parsing:** Converts natural English prompts (e.g., "secure login system with database") into optimized `Key-Value` status tokens (`login=1;db=1;api=0`).
* **Dynamic Resource Synthesis:** Automatically writes structural `AndroidManifest.xml`, responsive UI `activity_main.xml`, and functional `MainActivity.java` dynamically based on prompt intents.
* **Smooth Single-Line UI:** Features a fluid, colored CLI loading spinner that cycles through system compilation stages on a single line without flooding your terminal history.
* **Self-Contained Workspace Storage:** Automatically stacks all builds inside a central `builder_apk/` directory, isolating unique timestamps for each prompt execution.

---

## 🏗️ Architecture Pipeline

1. **`setup.sh`** ➡️ Validates system paths, installs runtime tools, and auto-executes the frontend console.
2. **`main.go`** ➡️ Intercepts user prompt strings, decodes intents into tokens, and operates the live CLI status animation.
3. **`builder.go`** ➡️ Generates XML assets/Java layouts, runs background compiler passes, and compresses packages into the final installable `.apk`.

---

## 🛠️ Quick Start & Installation

To deploy and boot up the GoBuilder engine, open your terminal and run the following commands to clone the repository and execute the pipeline:

`
# Clone the official repository
git clone [https://github.com/aditya12772011-rgb/GoBuilder.git](https://github.com/aditya12772011-rgb/GoBuilder.git)

# Navigate into the project directory
cd GoBuilder

# Grant execution permission to the deployment tool
chmod +x setup.sh

# Run the automated pipeline setup
./setup.sh
subscribe me on YouTube: https://youtube.com/@CyberhackTech-i6e
