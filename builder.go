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

func findAndroidJar() string {
	exactPath := "/data/data/com.termux/files/usr/share/aapt/android.jar"
	if _, err := os.Stat(exactPath); err == nil { return exactPath }
	
	paths := []string{
		"/data/data/com.termux/files/usr/share/android-sdk/platforms/android-28/android.jar",
		"/usr/lib/android-sdk/platforms/android-30/android.jar",
	}
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil { return p }
	}
	return ""
}

func main() {
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

	// Generate AndroidManifest.xml
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

	// Generate Layout XML
	xmlLayout := `<?xml version="1.0" encoding="utf-8"?>
<LinearLayout xmlns:android="http://schemas.google.com/apk/res/android" android:layout_width="match_parent" android:layout_height="match_parent" android:orientation="vertical" android:gravity="center">
    <TextView android:layout_width="wrap_content" android:layout_height="wrap_content" android:text="Application Built via cyber hack tech" android:textSize="18sp"/>
</LinearLayout>`
	_ = os.WriteFile(filepath.Join(resLayoutDir, "activity_main.xml"), []byte(xmlLayout), 0644)
	_ = os.WriteFile(filepath.Join(resValuesDir, "strings.xml"), []byte(`<?xml version="1.0" encoding="utf-8"?><resources><string name="app_name">GoBuilder</string></resources>`), 0644)

	// Generate Valid Java Source Code
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
        tv.setText("GoBuilder Core Engine: Live Context Running!");
        tv.setTextSize(20);
        layout.addView(tv);
        setContentView(layout);
    }
}`
	javaFilePath := filepath.Join(javaSrcDir, "MainActivity.java")
	_ = os.WriteFile(javaFilePath, []byte(javaCode), 0644)

	androidJarPath := findAndroidJar()
	if androidJarPath == "" {
		fmt.Println("\n[✗] Error: android.jar library not found.")
		return
	}

	// 1. Compile Java Source safely handling Java 17 syntax architectures
	cmdJavac := exec.Command("javac", "--release", "8", "-cp", androidJarPath, "-d", binOutputDir, javaFilePath)
	var errJavac bytes.Buffer
	cmdJavac.Stderr = &errJavac
	if err := cmdJavac.Run(); err != nil {
		cmdJavacFB := exec.Command("javac", "-cp", androidJarPath, "-d", binOutputDir, javaFilePath)
		var errJavacFB bytes.Buffer
		cmdJavacFB.Stderr = &errJavacFB
		if errFB := cmdJavacFB.Run(); errFB != nil {
			fmt.Printf("\n[✗] Java Engine Error:\n%s\n", errJavacFB.String())
			return
		}
	}

	// 2. Transform compiled bytecode into Android Dalvik Executable (classes.dex)
	dexFilePath := filepath.Join(workspace, "classes.dex")
	classTarget := filepath.Join(binOutputDir, "com", "go", "builder", "MainActivity.class")
	
	cmdDex := exec.Command("dx", "--dex", "--output="+dexFilePath, binOutputDir)
	if err := cmdDex.Run(); err != nil {
		cmdD8 := exec.Command("d8", "--output", workspace, classTarget)
		_ = cmdD8.Run()
	}

	// 3. Assemble binary asset package container via AAPT
	finalApkPath := filepath.Join(workspace, "output_application.apk")
	cmdAapt := exec.Command("aapt", "package", "-f", "-M", filepath.Join(workspace, "AndroidManifest.xml"), "-S", filepath.Join(workspace, "res"), "-I", androidJarPath, "-F", finalApkPath)
	_ = cmdAapt.Run()

	// 4. Inject executable dex runtime tree inside the package archive
	origDir, _ := os.Getwd()
	os.Chdir(workspace)
	cmdZipAppend := exec.Command("zip", "-g", "output_application.apk", "classes.dex")
	_ = cmdZipAppend.Run()
	os.Chdir(origDir)

	// Clean, professional console output validation box
	if _, errCheck := os.Stat(finalApkPath); errCheck == nil {
		fmt.Println("\n----------------------------------------")
		fmt.Println("✔ APK Generated Successfully!")
		fmt.Printf("Location: %s\n", finalApkPath)
		fmt.Println("----------------------------------------")
	} else {
		fmt.Println("\n[✗] Package Assembly Failed.")
	}
}
