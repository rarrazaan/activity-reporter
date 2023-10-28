package dependency

import (
	"activity-reporter/shared/helper"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func NewCloudinary() *cloudinary.Cloudinary {
	helper.LoadEnv()
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("CLOUD_KEY"), os.Getenv("CLOUD_SECRET"))
	if err != nil {
		return nil
	}
	return cld
}
