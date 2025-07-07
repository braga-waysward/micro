/*
Package main implements a snippet manager for the micro editor.

This tool allows users to manage code snippets that can be easily inserted
into their editing workflow. Snippets are stored in JSON format in micro's
configuration directory and can be listed, added, viewed, and deleted.

The manager follows the UNIX philosophy of doing one thing well and integrates
seamlessly with micro's ecosystem.
*/
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// SnippetManager handles the storage and management of code snippets.
// It maintains snippets in memory and persists them to a JSON file.
type SnippetManager struct {
	Filepath string            // Path to the snippets JSON file
	Snippets map[string]string // In-memory storage of snippets (name -> content)
}

/*
NewSnippetManager creates a new SnippetManager instance.

It initialises the configuration directory at ~/.config/micro if it doesn't exist
and sets up the path for the snippets JSON file. The in-memory snippet storage
is initialised as an empty map.
*/
func NewSnippetManager() *SnippetManager {
	home := os.Getenv("HOME")
	configDir := filepath.Join(home, ".config", "micro")
	
	// Create config directory if it doesn't exist
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, 0755)
	}
	
	return &SnippetManager{
		Filepath: filepath.Join(configDir, "snippets.json"),
		Snippets: make(map[string]string),
	}
}

/*
Load reads snippets from the JSON file into memory.

Returns nil if the file doesn't exist (initial empty snippets), or an error
if there are issues reading or parsing the file.
*/
func (sm *SnippetManager) Load() error {
	data, err := ioutil.ReadFile(sm.Filepath)
	if err != nil {
		if os.IsNotExist(err) {
			sm.Snippets = make(map[string]string)
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &sm.Snippets)
}

/*
Save writes the in-memory snippets to the JSON file.

The file is created with 0644 permissions (read/write for owner, read for others).
The JSON output is pretty-printed with 2-space indentation.
*/
func (sm *SnippetManager) Save() error {
	data, err := json.MarshalIndent(sm.Snippets, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(sm.Filepath, data, 0644)
}

/*
List displays all stored snippets in alphabetical order.

If no snippets are stored, it displays an appropriate message.
The output format shows each snippet name preceded by a bullet point.
*/
func (sm *SnippetManager) List() {
	if len(sm.Snippets) == 0 {
		fmt.Println("No snippets saved.")
		return
	}
	
	// Sort snippet names alphabetically
	keys := make([]string, 0, len(sm.Snippets))
	for k := range sm.Snippets {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	
	fmt.Println("Saved snippets:")
	for _, k := range keys {
		fmt.Println("- " + k)
	}
}

/*
Add creates a new snippet with the given name.

The snippet content is read from stdin until an empty line is encountered.
The new snippet is added to memory and immediately persisted to disk.
*/
func (sm *SnippetManager) Add(name string) error {
	fmt.Println("Paste your snippet. End input with an empty line:")
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	
	if err := scanner.Err(); err != nil {
		return err
	}
	
	snippet := strings.Join(lines, "\n")
	sm.Snippets[name] = snippet
	return sm.Save()
}

/*
Show displays the content of a snippet with the given name.

If the snippet doesn't exist, it displays an appropriate error message.
The content is printed exactly as stored, including all formatting.
*/
func (sm *SnippetManager) Show(name string) {
	snippet, ok := sm.Snippets[name]
	if !ok {
		fmt.Printf("Snippet '%s' not found.\n", name)
		return
	}
	fmt.Println(snippet)
}

/*
Delete removes a snippet with the given name.

Returns an error if the snippet doesn't exist. On success, the change is
immediately persisted to disk.
*/
func (sm *SnippetManager) Delete(name string) error {
	if _, ok := sm.Snippets[name]; !ok {
		return fmt.Errorf("Snippet '%s' not found", name)
	}
	delete(sm.Snippets, name)
	return sm.Save()
}

/*
printUsage displays the command-line interface usage instructions.

This includes all available commands and their expected arguments.
*/
func printUsage() {
	fmt.Println("Micro Snippet Manager - Manage code snippets for micro editor")
	fmt.Println("Usage:")
	fmt.Println("  snippetmanager list                 # List all snippets")
	fmt.Println("  snippetmanager add <name>           # Add snippet (input from stdin)")
	fmt.Println("  snippetmanager show <name>          # Show snippet content")
	fmt.Println("  snippetmanager delete <name>        # Delete snippet")
	fmt.Println("\nSnippets are stored in ~/.config/micro/snippets.json")
}

/*
main is the entry point for the snippet manager.

It handles command-line arguments and delegates to the appropriate SnippetManager
methods. Invalid commands or missing arguments result in usage instructions.
*/
func main() {
	sm := NewSnippetManager()
	err := sm.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading snippets: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "list":
		sm.List()
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide snippet name.")
			os.Exit(1)
		}
		name := os.Args[2]
		err := sm.Add(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error adding snippet: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Snippet '%s' added.\n", name)
	case "show":
		if len(os.Args) < 3 {
			fmt.Println("Please provide snippet name.")
			os.Exit(1)
		}
		sm.Show(os.Args[2])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide snippet name.")
			os.Exit(1)
		}
		err := sm.Delete(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting snippet: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Snippet '%s' deleted.\n", os.Args[2])
	default:
		printUsage()
		os.Exit(1)
	}
}