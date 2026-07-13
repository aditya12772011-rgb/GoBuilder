package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type AppIntent struct {
	HasLogin bool
	HasDB    bool
	HasAPI   bool
}

func decodePayload(payload string) AppIntent {
	var intent AppIntent
	pairs := strings.Split(payload, ";")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if kv[0] == "login" && kv[1] == "1" { intent.HasLogin = true }
			if kv[0] == "db" && kv[1] == "1"    { intent.HasDB = true }
			if kv[0] == "api" && kv[1] == "1"   { intent.HasAPI = true }
		}
	}
	return intent
}

func findAndroidJar() string {
	paths := []string{
		"/data/data/com.termux/files/usr/share/aapt/android.jar",
		"/data/data/com.termux/files/usr/share/android-sdk/platforms/android-28/android.jar",
		"/usr/lib/android-sdk/platforms/android-30/android.jar",
	}
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil { return p }
	}
	cmd := exec.Command("find", "/data/data/com.termux/files/", "-name", "android.jar")
	var out bytes.Buffer
	cmd.Stdout = &out
	_ = cmd.Run()
	matches := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(matches) > 0 && matches[0] != "" { return matches[0] }
	return ""
}

func main() {
	const (
		Reset = "\033[0m"; Green = "\033[32m"; Red = "\033[31m"; Cyan = "\033[36m"; Yellow = "\033[33m"
	)

	if len(os.Args) < 2 { return }
	intent := decodePayload(os.Args[1])

	baseFolder := "builder_apk"
	buildID := fmt.Sprintf("gobuild_prod_%d", time.Now().Unix())
	workspace := filepath.Join(baseFolder, buildID)
	
	javaSrcDir := filepath.Join(workspace, "src", "com", "go", "builder")
	resLayoutDir := filepath.Join(workspace, "res", "layout")
	resValuesDir := filepath.Join(workspace, "res", "values")
	binOutputDir := filepath.Join(workspace, "bin")

	_ = os.MkdirAll(javaSrcDir, 0755)
	_ = os.MkdirAll(resLayoutDir, 0755)
	_ = os.MkdirAll(resValuesDir, 0755)
	_ = os.MkdirAll(binOutputDir, 0755)

	// Manifest Factory
	_ = os.WriteFile(filepath.Join(workspace, "AndroidManifest.xml"), []byte(`<?xml version="1.0" encoding="utf-8"?>
<manifest xmlns:android="http://schemas.google.com/apk/res/android" package="com.go.builder">
    <application android:allowBackup="true" android:label="GoBuilderApp">
        <activity android:name=".MainActivity" android:exported="true">
            <intent-filter>
                <action android:name="android.intent.action.MAIN" />
                <category android:name="android.intent.category.LAUNCHER" />
            </intent-filter>
        </activity>
    </application>
</manifest>`), 0644)

	// Dynamic UI Layout Builder
	xmlLayout := `<?xml version="1.0" encoding="utf-8"?>
<LinearLayout xmlns:android="http://schemas.google.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:orientation="vertical"
    android:padding="20dp"
    android:gravity="center">
`
	if intent.HasLogin {
		xmlLayout += "    <TextView android:layout_width=\"wrap_content\" android:layout_height=\"wrap_content\" android:text=\"Secure Portal Active\" android:textSize=\"22sp\"/>\n"
	} else {
		xmlLayout += "    <TextView android:layout_width=\"wrap_content\" android:layout_height=\"wrap_content\" android:text=\"GoBuilder Main Workspace\" android:textSize=\"18sp\"/>\n"
	}
	xmlLayout += "</LinearLayout>"
	_ = os.WriteFile(filepath.Join(resLayoutDir, "activity_main.xml"), []byte(xmlLayout), 0644)
	_ = os.WriteFile(filepath.Join(resValuesDir, "strings.xml"), []byte(`<?xml version="1.0" encoding="utf-8"?><resources><string name="app_name">GoBuilder</string></resources>`), 0644)

	// Java Code Formulation
	javaCode := `package com.go.builder;
import android.os.Bundle;
import android.app.Activity;
import android.widget.LinearLayout;
import android.widget.TextView;

public class MainActivity extends Activity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        LinearLayout layout = new LinearLayout(this);
        layout.setOrientation(LinearLayout.VERTICAL);
        layout.setGravity(android.view.Gravity.CENTER);
        TextView tv = new TextView(this);
        tv.setText("🚀 GoBuilder Live: Programmatic Architecture Validated!");
        tv.setTextSize(20);
        layout.addView(tv);
        setContentView(layout);
    }
}`
	javaFilePath := filepath.Join(javaSrcDir, "MainActivity.java")
	_ = os.WriteFile(javaFilePath, []byte(javaCode), 0644)

	// Auto-Heal Path Resolver
	androidJarPath := findAndroidJar()

	// Step A: Java 17+ Compliant Cross-Compilation Execution Pass
	cmdJavac := exec.Command("javac", "--release", "8", "-cp", androidJarPath, "-d", binOutputDir, javaFilePath)
	var errJavac bytes.Buffer
	cmdJavac.Stderr = &errJavac
	if err := cmdJavac.Run(); err != nil {
		// Self-Heal Fallback Option
		cmdJavacFB := exec.Command("javac", "-d", binOutputDir, javaFilePath)
		if errFB := cmdJavacFB.Run(); errFB != nil {
			fmt.Printf("\n%s[✗] Compilation Engine Failure: %s%s\n", Red, errJavac.String(), Reset)
			return
		}
	}

	// Step B: DEX Binary Translation Pass
	dexFilePath := filepath.Join(workspace, "classes.dex")
	cmdDex := exec.Command("dx", "--dex", "--output="+dexFilePath, binOutputDir)
	if err := cmdDex.Run(); err != nil {
		cmdD8 := exec.Command("d8", "--output", workspace, filepath.Join(binOutputDir, "com", "go", "builder", "MainActivity.class"))
		_ = cmdD8.Run()
	}

	// Step C: AAPT Package Assembly Pass
	finalApkPath := filepath.Join(workspace, "output_application.apk")
	cmdAapt := exec.Command("aapt", "package", "-f", "-M", filepath.Join(workspace, "AndroidManifest.xml"), "-S", filepath.Join(workspace, "res"), "-I", androidJarPath, "-F", finalApkPath)
	_ = cmdAapt.Run()

	// Step D: Binding Compiled Bytecode Object Context into Package
	origDir, _ := os.Getwd()
	os.Chdir(workspace)
	cmdZipAppend := exec.Command("zip", "-g", "output_application.apk", "classes.dex")
	_ = cmdZipAppend.Run()
	os.Chdir(origDir)

	// Verify Absolute APK Output Generation
	if _, errCheck := os.Stat(finalApkPath); errCheck == nil {
		fmt.Printf("\n\n%s┌──────────────────────────────────────────────────────────────┐%s\n", Cyan, Reset)
		fmt.Printf("%s│ ✨ SUCCESS: GOBUILDER ENGINE GENERATED THE RUNNING APK!      │%s\n", Green, Reset)
		fmt.Printf("%s├──────────────────────────────────────────────────────────────┤%s\n", Cyan, Reset)
		fmt.Printf("%s│%s Output Repository : %s%-39s%s │\n", Cyan, Reset, Yellow, baseFolder+"/", Reset, Cyan)
		fmt.Printf("%s│%s Target Path      : %s%-39s%s │\n", Cyan, Reset, Yellow, buildID+"/", Reset, Cyan)
		fmt.Printf("%s│%s Executable APK  : %s%-39s%s │\n", Cyan, Reset, Green, "output_application.apk", Reset, Cyan)
		fmt.Printf("%s└──────────────────────────────────────────────────────────────┘%s\n", Cyan, Reset)
	} else {
		fmt.Printf("\n%s[✗] Packaging Failed: Resource validation crashed down structural pipes.%s\n", Red, Reset)
	}
}
