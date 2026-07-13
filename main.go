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

	clearCmd := exec.Command("clear")
	clearCmd.Stdout = os.Stdout
	_ = clearCmd.Run()

	fmt.Println()
	fmt.Printf("%s┌────────────────────────────────────────────────────────┐%s\n", Cyan, Reset)
	fmt.Printf("%s│ %s⚡ GOBUILDER AI v3.0 : NEXT-GEN COMPILER TOOLCHAIN%s    %s│%s\n", Cyan, Green, Reset, Cyan, Reset)
	fmt.Printf("%s└────────────────────────────────────────────────────────┘%s\n", Cyan, Reset)
	fmt.Printf("%s[status]%s Core Engine Engaged. Ready for intent injection.\n\n", Green, Reset)

	fmt.Printf("%s GoBuilder@Terminal%s:%s~$%s Enter App Concept:\n ➔ ", Cyan, Red, Yellow, Reset)
	
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() { return }
	userPrompt := strings.TrimSpace(scanner.Text())

	if userPrompt == "" {
		fmt.Printf("\n%s[!] Prompt cannot be empty.%s\n", Red, Reset)
		return
	}
	tokenPayload := parseToTokens(userPrompt)

	fmt.Print(Hide)
	done := make(chan bool)
	go func() {
		// Premium Cyberpunk Loader
		blocks := []string{"■□□□□□□□□□ 10%", "■■□□□□□□□□ 20%", "■■■■□□□□□□ 40%", "■■■■■■□□□□ 60%", "■■■■■■■■□□ 80%", "■■■■■■■■■■ 100%"}
		stages := []string{
			"ANALYZING INTENT TOKENS  ",
			"GENERATING XML STRUCTURES ",
			"COMPILING JAVA BYTECODE   ",
			"CONVERTING TO DEX BINARY  ",
			"PACKAGING WITH AAPT LAYER ",
			"FINALIZING APK CONTAINER  ",
		}
		
		idx := 0
		ticker := time.NewTicker(350 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-done: return
			case <-ticker.C:
				fmt.Print("\r", ClearLN)
				fmt.Printf("%s[⚙] %s %s[%s]%s", Yellow, stages[idx%len(stages)], Cyan, blocks[idx%len(blocks)], Reset)
				idx++
			}
		}
	}()

	cmd := exec.Command("go", "run", "builder.go", tokenPayload)
	outputBytes, _ := cmd.CombinedOutput()
	
	done <- true
	fmt.Print("\r", ClearLN)
	fmt.Print(Show)

	fmt.Print(string(outputBytes))
}
