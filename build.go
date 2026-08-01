go run build.go

git add build.go
git commit -m "Add build script in Go"
git push origin main

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Helper function jo files ko copy karti hai (Deno.copyFile ki jagah)
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Destination folder agar nahi hai toh create karo
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

// Helper function jo text file write karti hai (Deno.writeTextFileSync ki jagah)
func writeTextFile(path, content string) error {
	dstDir := filepath.Dir(path)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), 0644)
}

func main() {
	fmt.Println("🚀 Starting Go build release process...")

	// Go modules ke liye distribution directories prepare kar rahe hain
	dirs := []string{"dist", "npm"}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			return
		}
	}

	// README aur LICENSE copy karna (jaise JS script karti thi)
	err := copyFile("README.md", "npm/README.md")
	if err != nil {
		fmt.Printf("⚠️ Warning: Could not copy README.md: %v\n", err)
	} else {
		fmt.Println("✅ README.md copied to npm/")
	}

	err = copyFile("LICENSE", "npm/LICENSE")
	if err != nil {
		fmt.Printf("⚠️ Warning: Could not copy LICENSE: %v\n", err)
	} else {
		fmt.Println("✅ LICENSE copied to npm/")
	}

	// Go packages ke metadata/headers write karna
	buildNotice := "// This release build is managed automatically for Go module tinycolor.\n"
	_ = writeTextFile("dist/BUILD_INFO.txt", buildNotice)

	fmt.Println("🎉 Build completed successfully!")
}

