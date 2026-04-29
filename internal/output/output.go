package output

import (
	"encoding/json"
	"fmt"
	"os"
)

type FriendlyTexter interface {
	FriendlyText() string
}

func PrintItems[T FriendlyTexter](items []T, asJSON bool) {
	if asJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(items)
		return
	}
	for _, item := range items {
		fmt.Println(item.FriendlyText())
		fmt.Println()
	}
}

func PrintItem[T FriendlyTexter](item T, asJSON bool) {
	if asJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(item)
		return
	}
	fmt.Println(item.FriendlyText())
}

func PrintSuccess(msg string, asJSON bool) {
	if asJSON {
		fmt.Printf(`{"ok": true, "message": %q}`+"\n", msg)
		return
	}
	fmt.Println("Success:", msg)
}

func PrintError(err error) {
	fmt.Fprintln(os.Stderr, "Error:", err)
}
