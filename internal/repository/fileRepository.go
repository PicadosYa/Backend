package repository

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"

	// "github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type IFileRepository interface {
	UploadFile(file interface{}, publicID string) (string, error)
}

type Cloudinary struct {
	Cld *cloudinary.Cloudinary
	Ctx context.Context
}

func NewCloudinary() *Cloudinary {
	// Add your Cloudinary credentials, set configuration parameter
	// Secure=true to return "https" URLs, and create a context
	//===================
	cld, _ := cloudinary.New()
	cld.Config.URL.Secure = true
	ctx := context.Background()
	return &Cloudinary{cld, ctx}
}

func (c *Cloudinary) UploadFile(file interface{}, publicID string) (string, error) {

	// Upload the image.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err := c.Cld.Upload.Upload(c.Ctx, file, uploader.UploadParams{
		PublicID:       publicID,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true)})
	if err != nil {
		log.Panic(err)
		return "", err
	}

	// Log the delivery URL
	log.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL)
	return resp.SecureURL, nil
}
