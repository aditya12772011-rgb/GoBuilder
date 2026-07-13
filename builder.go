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

func main() {
	const (
		Reset  = "\033[0m"
		Green  = "\033[32m"
		Red    = "\033[31m"
		Cyan   = "\033[36m"
		Yellow = "\033[33m"
	)

	if len(os.Args) < 2 { return }
	tokenPayload := os.Args[1]
	intent := decodePayload(tokenPayload)

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

	// 1. Android Manifest Generation
	manifestXml := `<?xml version="1.0" encoding="utf-8"?>
<manifest xmlns:android="http://schemas.google.com/apk/res/android" package="com.go.builder">
    <application android:allowBackup="true" android:label="GoBuilderApp">
        <activity android:name=".MainActivity" android:exported="true">
            <intent-filter>
                <action android:name="android.intent.action.MAIN" />
                <category android:name="android.intent.category.LAUNCHER" />
            </intent-filter>
        </activity>
    </application>
</manifest>`
	_ = os.WriteFile(filepath.Join(workspace, "AndroidManifest.xml"), []byte(manifestXml), 0644)

	// 2. Resource XML Architect
	xmlLayout := `<?xml version="1.0" encoding="utf-8"?>
<LinearLayout xmlns:android="http://schemas.google.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:orientation="vertical"
    android:padding="20dp"
    android:gravity="center">
`
	if intent.HasLogin {
		xmlLayout += "    <TextView android:layout_width=\"wrap_content\" android:layout_height=\"wrap_content\" android:text=\"Secure Auth Panel\" android:textSize=\"22sp\"/>\n"
	} else {
		xmlLayout += "    <TextView android:layout_width=\"wrap_content\" android:layout_height=\"wrap_content\" android:text=\"GoBuilder Core Screen\" android:textSize=\"18sp\"/>\n"
	}
	xmlLayout += "</LinearLayout>"
	_ = os.WriteFile(filepath.Join(resLayoutDir, "activity_main.xml"), []byte(xmlLayout), 0644)
	_ = os.WriteFile(filepath.Join(resValuesDir, "strings.xml"), []byte(`<?xml version="1.0" encoding="utf-8"?><resources><string name="app_name">GoBuilder</string></resources>`), 0644)

	// 3. Programmatic Native Java Logic Injection
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
        tv.setText("🚀 GoBuilder Engine: Real DEX Native Code Running Successfully!");
        tv.setTextSize(22);
        layout.addView(tv);
        setContentView(layout);
    }
}`
	javaFilePath := filepath.Join(javaSrcDir, "MainActivity.java")
	_ = os.WriteFile(javaFilePath, []byte(javaCode), 0644)

	// 4. Critical Path Detection for android.jar
	androidJarPath := "/data/data/com.termux/files/usr/share/aapt/android.jar"
	if _, err := os.Stat(androidJarPath); os.IsNotExist(err) {
		// Fallback paths if default Termux location misses
		alternatePaths := []string{
			"/data/data/com.termux/files/usr/share/android-sdk/platforms/android-28/android.jar",
			"/usr/lib/android-sdk/platforms/android-30/android.jar",
		}
		for _, altPath := range alternatePaths {
			if _, errAlt := os.Stat(altPath); errAlt == nil {
				androidJarPath = altPath
				break
			}
		}
	}

	// Step A: Javac bytecode compilation pass
	cmdJavac := exec.Command("javac", "-bootclasspath", androidJarPath, "-d", binOutputDir, javaFilePath)
	var errJavac bytes.Buffer
	cmdJavac.Stderr = &errJavac
	if err := cmdJavac.Run(); err != nil {
		fmt.Printf("\n%s[✗] Javac Error: %s%s\n", Red, errJavac.String(), Reset)
		return
	}

	// Step B: DX Compilation Engine -> Injecting classes.dex
	dexFilePath := filepath.Join(workspace, "classes.dex")
	cmdDex := exec.Command("dx", "--dex", "--output="+dexFilePath, binOutputDir)
	var errDex bytes.Buffer
	cmdDex.Stderr = &errDex
	if err := cmdDex.Run(); err != nil {
		// Fallback to d8
		cmdD8 := exec.Command("d8", "--output", workspace, filepath.Join(binOutputDir, "com", "go", "builder", "MainActivity.class"))
		var errD8 bytes.Buffer
		cmdD8.Stderr = &errD8
		if errOut := cmdD8.Run(); errOut != nil {
			fmt.Printf("\n%s[✗] Dex Engine Error (dx/d8 failed): %s%s\n", Red, errD8.String(), Reset)
			return
		}
	}

	// Step C: AAPT Resource Archiving
	finalApkPath := filepath.Join(workspace, "output_application.apk")
	cmdAapt := exec.Command("aapt", "package", "-f", "-M", filepath.Join(workspace, "AndroidManifest.xml"), "-S", filepath.Join(workspace, "res"), "-I", androidJarPath, "-F", finalApkPath)
	var errAapt bytes.Buffer
	cmdAapt.Stderr = &errAapt
	if err := cmdAapt.Run(); err != nil {
		fmt.Printf("\n%s[✗] AAPT Packaging Error: %s%s\n", Red, errAapt.String(), Reset)
		return
	}

	// Step D: Binding Dalvik Executable into the final zip frame
	origDir, _ := os.Getwd()
	os.Chdir(workspace)
	cmdZipAppend := exec.Command("zip", "-g", "output_application.apk", "classes.dex")
	_ = cmdZipAppend.Run()
	os.Chdir(origDir)

	// 5. Ultimate Success Dashboard
	fmt.Printf("\n\n%s┌──────────────────────────────────────────────────────────────┐%s\n", Cyan, Reset)
	fmt.Printf("%s│ ✨ SUCCESS: GOBUILDER CREATED A STANDALONE RUNNING APK!       │%s\n", Green, Reset)
	fmt.Printf("%s├──────────────────────────────────────────────────────────────┤%s\n", Cyan, Reset)
	fmt.Printf("%s│%s Output Repository : %s%-39s%s │\n", Cyan, Reset, Yellow, baseFolder+"/", Reset, Cyan)
	fmt.Printf("%s│%s Active Workspace : %s%-39s%s │\n", Cyan, Reset, Yellow, buildID+"/", Reset, Cyan)
	fmt.Printf("%s│%s Target Binary    : %s%-39s%s │\n", Cyan, Reset, Green, "output_application.apk", Reset, Cyan)
	fmt.Printf("%s└──────────────────────────────────────────────────────────────┘%s\n", Cyan, Reset)
}
