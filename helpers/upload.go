package helpers

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"

	"syncspend/config"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
)

func UploadProfileImage(file multipart.File, header *multipart.FileHeader, id int) (string, error) {
	client := config.BlobClient
	containerName := config.ENV.AzureContainerName
	folder := "profile"

	// Get file extension and create a unique blob name
	ext := strings.ToLower(filepath.Ext(header.Filename))
	blobName := fmt.Sprintf("%s/profile_%d%s", folder, id, ext)

	// Read file content into byte slice
	fileContent := make([]byte, header.Size)
	_, err := file.Read(fileContent)
	if err != nil {
		return "", fmt.Errorf("failed to read file content: %v", err)
	}

	// Determine the content type based on the file extension
	contentType := getContentTypeByExtension(ext)

	// Prepare context for the upload
	ctx := context.Background()

	// Set upload options
	settings := blockblob.UploadBufferOptions{
		Concurrency: 1, // Number of concurrent uploads
		HTTPHeaders: &blob.HTTPHeaders{
			BlobContentType: to.Ptr(contentType), // Convert contentType to *string
		},
	}

	// Upload the buffer to Azure Blob Storage
	_, err = client.UploadBuffer(ctx, containerName, blobName, fileContent, &settings)
	if err != nil {
		return "", fmt.Errorf("failed to upload blob: %v", err)
	}

	// Generate the URL of the uploaded image
	url := fmt.Sprintf("%s/%s/%s", strings.TrimRight(config.ENV.AzureStorageEndpoint, "/"), containerName, blobName)
	log.Printf("Image uploaded successfully to: %s", url)

	return url, nil
}

// Helper function to get content type based on file extension
func getContentTypeByExtension(ext string) string {
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".bmp":
		return "image/bmp"
	case ".webp":
		return "image/webp"
	default:
		return "application/octet-stream"
	}
}
