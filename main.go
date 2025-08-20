package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/wii-tools/wadlib"
)

const OriginalURL = "http://news.wapp.wii.com/v2/%d/%03d/news.bin"
const NewURL = "http://wii.rauln.com/news/%d/%03d/news.bin"
const ExpectedOffset = 0x1AC37C // The offset in the decrypted data where the URL is located

func verifyWAD(wadPath string) error {
	// Load the WAD file
	wad, err := wadlib.LoadWADFromFile(wadPath)
	if err != nil {
		return err
	}

	// Verify that the WAD contains the updated URL in the content at index 1
	content, err := wad.GetContent(1)
	if err != nil {
		return err
	}

	// Check if the content contains the new URL in the specific offset
	offset := bytes.Index(content, []byte(NewURL))
	if offset == -1 {
		return fmt.Errorf("new URL not found in content")
	}

	if offset != ExpectedOffset {
		return fmt.Errorf("new URL found at unexpected offset: 0x%X, expected: 0x%X", offset, ExpectedOffset)
	}

	return nil
}

func patchNewsURL(wadPath string, outputPath string) error {
	wad, err := wadlib.LoadWADFromFile(wadPath)
	if err != nil {
		return err
	}

	// Get decrypted content at index 1 (record with ID "0000000b")
	content, err := wad.GetContent(1)
	if err != nil {
		return err
	}

	// Create container for the updatedContent
	updatedContent := make([]byte, len(content))

	// 44 bytes
	originalContent := []byte(OriginalURL)

	// 43 bytes
	newContent := []byte(NewURL)

	// Pad the new URL to match original length (44 bytes)
	paddedURL := make([]byte, len(originalContent))
	copy(paddedURL, newContent)

	// Find and replace the URL
	offset := bytes.Index(content, originalContent)
	if offset == -1 {
		return fmt.Errorf("original URL not found in content")
	}
	fmt.Printf("Found original URL at hex offset: 0x%X\n", offset)

	// Alternatively:
	// .... Patch 0x2C (44) bytes at offset 0x1AC37C in the decrypted data

	copy(updatedContent[offset:offset+len(originalContent)], paddedURL)

	err = wad.UpdateContent(1, updatedContent)
	if err != nil {
		return err
	}

	wadBytes, err := wad.GetWAD(wadlib.WADTypeCommon)
	if err != nil {
		return err
	}

	return os.WriteFile(outputPath, wadBytes, 0644)
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: wiinewspr-patcher <path_to_news_wad_file> <output_path>")
		fmt.Println("Example: wiinewspr-patcher news.wad patched_news.wad")
		return
	}

	// Undocumented command to verify patched WAD locally
	if os.Args[1] == "verify" {
		if err := verifyWAD(os.Args[2]); err != nil {
			fmt.Printf("Verification failed: %v\n", err)
			return
		}

		fmt.Println("WAD verification successful, new URL found in content at expected offset")
		return
	}

	patchNewsURL(os.Args[1], os.Args[2])
	fmt.Printf("News URL patched successfully, saved to '%s'", os.Args[2])
}
