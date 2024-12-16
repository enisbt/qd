package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const aliasFile = ".qd_aliases.json"

type AliasMap map[string]string

func loadAliases() (AliasMap, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	aliasPath := filepath.Join(homeDir, aliasFile)

	if _, err := os.Stat(aliasPath); os.IsNotExist(err) {
		return AliasMap{}, nil
	}

	data, err := os.ReadFile(aliasPath)
	if err != nil {
		return nil, err
	}

	var aliases AliasMap
	if err := json.Unmarshal(data, &aliases); err != nil {
		return nil, err
	}

	return aliases, nil
}

func saveAliases(aliases AliasMap) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	filePath := filepath.Join(homeDir, aliasFile)

	data, err := json.MarshalIndent(aliases, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

func saveAlias(alias string) error {
	aliases, err := loadAliases()
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	aliases[alias] = cwd
	if err := saveAliases(aliases); err != nil {
		return err
	}

	fmt.Printf("Saved alias '%s' for directory '%s'\n", alias, cwd)
	return nil
}

func listAliases() error {
	aliases, err := loadAliases()
	if err != nil {
		return err
	}

	if len(aliases) == 0 {
		fmt.Println("No aliases saved.")
		return nil
	}

	fmt.Println("Saved aliases:")
	for alias, path := range aliases {
		fmt.Printf("  %s -> %s\n", alias, path)
	}
	return nil
}

func deleteAlias(alias string) error {
	aliases, err := loadAliases()
	if err != nil {
		return err
	}

	if len(aliases) == 0 {
		fmt.Println("No aliases to delete.")
	}
	_, ok := aliases[alias]
	if !ok {
		return fmt.Errorf("alias '%s' not found", alias)
	}

	delete(aliases, alias)
	if err := saveAliases(aliases); err != nil {
		return err
	}

	fmt.Printf("Deleted alias '%s'\n", alias)
	return nil
}

func gotoCommand(alias string) error {
	aliases, err := loadAliases()
	if err != nil {
		return err
	}

	dir, exists := aliases[alias]
	if !exists {
		return fmt.Errorf("alias '%s' not found", alias)
	}

	fmt.Println(dir)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: qd <command> [<args>]")
		return
	}

	switch os.Args[1] {
	case "save":
		if len(os.Args) < 3 {
			fmt.Println("Usage: qd save <alias>")
			return
		}
		if err := saveAlias(os.Args[2]); err != nil {
			fmt.Println("Error: ", err)
		}
	case "list":
		if err := listAliases(); err != nil {
			fmt.Println("Error: ", err)
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: qd delete <alias>")
			return
		}
		if err := deleteAlias(os.Args[2]); err != nil {
			fmt.Println("Error: ", err)
		}
	default:
		if len(os.Args) < 2 {
			fmt.Println("Usage: qd <alias>")
			return
		}
		if err := gotoCommand(os.Args[1]); err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
