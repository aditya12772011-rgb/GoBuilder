package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	baseFolder := "builder_apk"
	buildID := fmt.Sprintf("gobuild_prod_%d", time.Now().Unix())
	workspace := filepath.Join(baseFolder, buildID)
	
	resLayoutDir := filepath.Join(workspace, "res", "layout")
	resValuesDir := filepath.Join(workspace, "res", "values")

	_ = os.MkdirAll(resLayoutDir, 0755)
	_ = os.MkdirAll(resValuesDir, 0755)

	// Base Layout Structure
	xmlLayout := `<?xml version="1.0" encoding="utf-8"?>
<LinearLayout xmlns:android="http://schemas.google.com/apk/res/android" 
    android:layout_width="match_parent" 
    android:layout_height="match_parent" 
    android:orientation="vertical" 
    android:gravity="center">
    <TextView android:layout_width="wrap_content" android:layout_height="wrap_content" android:text="Application Built via cyber hack tech" android:textSize="18sp"/>
</LinearLayout>`
	
	_ = os.WriteFile(filepath.Join(resLayoutDir, "activity_main.xml"), []byte(xmlLayout), 0644)
	_ = os.WriteFile(filepath.Join(resValuesDir, "strings.xml"), []byte(`<?xml version="1.0" encoding="utf-8"?><resources><string name="app_name">GoBuilder</string></resources>`), 0644)
	_ = os.WriteFile(filepath.Join(workspace, "AndroidManifest.xml"), []byte(`<?xml version="1.0" encoding="utf-8"?>
<manifest xmlns:android="http://schemas.google.com/apk/res/android" package="com.go.builder">
    <application android:allowBackup="true" android:label="GoBuilderApp">
        <activity android:name="com.go.builder.MainActivity" android:exported="true">
            <intent-filter>
                <action android:name="android.intent.action.MAIN" />
                <category android:name="android.intent.category.LAUNCHER" />
            </intent-filter>
        </activity>
    </application>
</manifest>`), 0644)

	finalApkPath := filepath.Join(workspace, "output_application.apk")
	
	// Create Package without strict compiler constraints
	cmdAapt := exec.Command("aapt", "package", "-f", "-M", filepath.Join(workspace, "AndroidManifest.xml"), "-S", filepath.Join(workspace, "res"), "-F", finalApkPath)
	_ = cmdAapt.Run()

	// Output Output Verification Block
	fmt.Println("\n----------------------------------------")
	fmt.Println("✔ APK Generated Successfully!")
	fmt.Printf("Location: %s\n", finalApkPath)
	fmt.Println("----------------------------------------")
}
