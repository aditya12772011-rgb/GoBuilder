package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func parseToTokens(prompt string) string {
	hasLogin, hasDB, hasAPI := "0", "0", "0"
	words := strings.Fields(strings.ToLower(prompt))
	for _, word := range words {
		if word == "login" || word == "auth" || word == "signin" { hasLogin = "1" }
		if word == "database" || word == "storage" || word == "db" { hasDB = "1" }
		if word == "api" || word == "internet" || word == "network" { hasAPI = "1" }
	}
	return fmt.Sprintf("login=%s;db=%s;api=%s", hasLogin, hasDB, hasAPI)
}

func main() {
	const (
		Reset   = "\033[0m"
		Green   = "\033[32m"
		Cyan    = "\033[36m"
		Yellow  = "\033[33m"
		Red     = "\033[31m"
		ClearLN = "\033[2K"
		Hide    = "\033[?25l"
		Show    = "\033[?25h"
	)

	fmt.Println()
	fmt.Printf("%s┌────────────────────────────────────────────────────────┐%s\n", Cyan, Reset)
	fmt.Printf("%s│ %s🤖 GOBUILDER AI v2.0 // DEEP INTENT ARTIFACT ENGINE%s  %s│%s\n", Cyan, Green, Reset, Cyan, Reset)
	fmt.Printf("%s└────────────────────────────────────────────────────────┘%s\n", Cyan, Reset)
	fmt.Printf("%s[info]%s Core architecture initialized. Operational status: READY.\n\n", Green, Reset)

	// Beautiful Prompt Input UI
	fmt.Printf("%s💻 GoBuilder@Console%s:%s~$%s Describe your app features:\n👉 ", Cyan, Red, Yellow, Reset)
	
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() { return }
	userPrompt := strings.TrimSpace(scanner.Text())

	if userPrompt == "" {
		fmt.Printf("\n%s[✗] Null prompt parameter intercepted. Execution aborted.%s\n", Red, Reset)
		return
	}
	tokenPayload := parseToTokens(userPrompt)

	// Premium Terminal Spinner Animation
	fmt.Print(Hide)
	done := make(chan bool)
	go func() {
		frames := []string{"◢", "◣", "◤", "◥"}
		stages := []string{
			"Tokenizing semantic prompt expressions",
			"Compiling functional UI layout matrices",
			"Running resource mapping assets compilation via AAPT",
			"Synthesizing safe native Java Activity bytecode",
			"Executing DX cross-compiler core: Converting into classes.dex",
			"Packaging deployment wrappers into clean APK signature archive",
		}
		i, stageIdx := 0, 0
		ticker := time.NewTicker(70 * time.Millisecond)
		stageTicker := time.NewTicker(1100 * time.Millisecond)
		defer ticker.Stop()
		defer stageTicker.Stop()

		for {
			select {
			case <-done: return
			case <-ticker.C:
				fmt.Print("\r", ClearLN)
				fmt.Printf("%s[%s]%s Processing: %s%s...", Cyan, frames[i%len(frames)], Reset, Yellow, stages[stageIdx], Reset)
				i++
			case <-stageTicker.C:
				if stageIdx < len(stages)-1 { stageIdx++ }
			}
		}
	}()

	cmd := exec.Command("go", "run", "builder.go", tokenPayload)
	outputBytes, _ := cmd.CombinedOutput()
	
	done <- true
	fmt.Print("\r", ClearLN)
	fmt.Print(Show)

	// Rendering compiler logs from builder.go
	fmt.Print(string(outputBytes))
}
