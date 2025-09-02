package howtousecomposeinterfaces

import "fmt"

// Downloader Interface
type Downloader interface {
	Download(file string) ([]byte, error)
}

// Uploader interface
type Uploader interface {
	Upload(file string, data []byte) error
}

// CloudStorage is a composed interface
type CloudStorage interface {
	Downloader
	Uploader
}

// AWS S3 implementation
type S3 struct{}

func (s *S3) Download(file string) ([]byte, error) {
	return []byte("file data from S3"), nil
}

func (s *S3) Upload(file string, data []byte) error {
	fmt.Println("Uploaded to s3:", file)
	return nil
}

// Google Cloud Storage implementation
type GCS struct{}

func (g *GCS) Download(file string) ([]byte, error) {
	return []byte("file data from GCS"), nil
}

func (g *GCS) Upload(file string, data []byte) error {
	fmt.Println("Uploaded to GCS:", file)
	return nil
}

func Cloud_Storage_Client() {
	fmt.Println("\n===> Composed interface Cloud Storage Sample")
	var storage CloudStorage

	storage = &S3{}
	storage.Upload("report.pdf", []byte("data"))
	data, _ := storage.Download("report.pdf")
	fmt.Println("Downloaded:", string(data))

	storage = &GCS{}
	storage.Upload("image.png", []byte("pic data"))
}
