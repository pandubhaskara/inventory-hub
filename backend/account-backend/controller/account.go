package controller

import (
	"account-backend/config"
	"account-backend/database"
	"account-backend/helper"
	"account-backend/model"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AccountController struct{}

func (ctlr *AccountController) Index(c *gin.Context) {
	data := c.MustGet("account").(model.Account)

	// Return a success message to the client
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (ctlr *AccountController) Update(c *gin.Context) {
	data := c.MustGet("account").(model.Account)

	// Get the data off req body
	type Body struct {
		Name                   string `json:"name" binding:"required" `
		MobileNumber           string `json:"mobile_number" binding:"required,min=10,max=16"`
		Gender                 string `json:"gender" binding:"required"`
		MaritalStatus          string `json:"marital_status" binding:"required"`
		PlaceOfBirth           string `json:"place_of_birth" binding:"required"`
		DateOfBirth            string `json:"date_of_birth" binding:"required"`
		NationalIdentityNumber string `json:"national_identity_number" binding:"required,len=16"`
	}

	var body Body

	if err := c.ShouldBindJSON(&body); err != nil {
		helper.Logger.Error(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	// Update account data
	dateOfBirth, _ := time.Parse("2006-01-02", body.DateOfBirth)

	result := database.Db.Where("email=?", data.Email).Updates(&model.Account{
		Name:                   body.Name,
		MobileNumber:           &body.MobileNumber,
		Gender:                 &body.Gender,
		MaritalStatus:          &body.MaritalStatus,
		PlaceOfBirth:           &body.PlaceOfBirth,
		DateOfBirth:            &dateOfBirth,
		NationalIdentityNumber: &body.NationalIdentityNumber,
	})
	if result.Error != nil {
		helper.Logger.Error(result.Error)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": result.Error,
		})
		c.Abort()
		return
	}

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"message": "Account has been updated successfully!",
	})
}

func (ctlr *AccountController) UploadProfilePicture(c *gin.Context) {
	data := c.MustGet("account").(model.Account)

	// Load the AWS SDK configuration
	cfg, err := awsConfig.LoadDefaultConfig(context.Background())
	if err != nil {
		helper.Logger.Error(err)
		c.JSON(http.StatusPreconditionFailed, gin.H{
			"message": err,
		})
		c.Abort()
		return
	}

	// Create an Amazon S3 service client
	svc := s3.NewFromConfig(cfg)

	// get the file
	file, err := c.FormFile("file")
	if err != nil {
		helper.Logger.Error(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	// Get the file content type and access the file extension
	fileType := strings.Split(file.Header.Get("Content-Type"), "/")[1]
	if !(fileType == "jpeg" || fileType == "png") {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please upload image with jpeg/png format!",
		})
		c.Abort()
		return
	}

	f, err := file.Open()
	if err != nil {
		helper.Logger.Error(err)
		c.JSON(http.StatusPreconditionFailed, gin.H{
			"message": err,
		})
		c.Abort()
		return
	}
	defer f.Close()

	uploader := manager.NewUploader(svc)

	folder := uuid.New()

	// upload file to aws
	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:             aws.String(config.Aws.Bucket),
		Key:                aws.String(folder.String() + "/" + file.Filename),
		Body:               f,
		ACL:                "public-read",
		ContentDisposition: aws.String("inline"),
		ContentType:        aws.String(fileType),
	})
	if uploadErr != nil {
		helper.Logger.Error(uploadErr)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Failed to upload file",
		})
		c.Abort()
		return
	}

	update := database.Db.Model(&model.Account{}).Where("email=?", data.Email).Update("profile_picture", result.Location)
	if update.Error != nil {
		helper.Logger.Error(update.Error)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": update.Error,
		})
		c.Abort()
		return
	}

	// Return a success message to the client
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully update profile picture!",
	})
}
