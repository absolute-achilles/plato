package storage

// TODO: implement storage here
type S3Storage struct {
	Bucket     string
	Region     string
	AccessKey  string
	SecretKey  string
	Endpoint   string // For S3-compatible services like Cloudflare R2
	CDNBaseURL string
}
