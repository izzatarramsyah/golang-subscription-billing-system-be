package handlers

import (
	// "fmt"
    "net/http"
    "os"
    "path/filepath"
    // "strconv"

    "github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"subscription-billing-system/models"
    apiModels "subscription-billing-system/api/v1/models"
)

type EbookAPI struct {
	usecase models.EbookUseCase
	subscriptionUseCase models.SubscriptionUseCase
}

func NewEbookAPI(usecase models.EbookUseCase, subscriptionUseCase models.SubscriptionUseCase) *EbookAPI {
	return &EbookAPI{usecase, subscriptionUseCase}
}

func (h *EbookAPI) UploadEbook(c *gin.Context) {

    idParam := c.PostForm("productID")
    productId, err := uuid.Parse(idParam)
    title := c.PostForm("title")
    file, err := c.FormFile("file")
    if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "File is required"))
        return
    }

    uploadPath := "./uploads/ebooks/"
    if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
        os.MkdirAll(uploadPath, os.ModePerm)
    }

    filePath := filepath.Join(uploadPath, file.Filename)
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "File to upload file"))
        return
    }

    h.usecase.UploadEbook(title, filePath, productId)
	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(true, "File uploaded successfully"))
}

func (h *EbookAPI) ListEbooks(c *gin.Context) {
	ebooks, err := h.usecase.ListEbooks()

	if err != nil {
        c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to retrive Ebooks"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(ebooks, "Ebooks retrive successfully"))
}

func (h *EbookAPI) DownloadEbook(c *gin.Context) {
    idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid Ebook ID"))
        return
    }

    ebook, err := h.usecase.GetEbook(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed retrive ebook"))
        return
    }

    c.FileAttachment(ebook.FilePath, filepath.Base(ebook.FilePath))
}

func (h *EbookAPI) GetEbookAccess(c *gin.Context) {
    idParam := c.Param("id")
    id, err := uuid.Parse(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid Ebook ID"))
        return
    }

    // Ambil userID dari token JWT (anggap kamu simpan di context)
    // userIDStr, exists := c.Get("userID")
    // if !exists {
    //     c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
    //     return
    // }
    // userID, err := uuid.Parse(userIDStr.(string))
    // if err != nil {
    //     c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
    //     return
    // }

    // Cek apakah user punya subscription yang aktif
    // active, err := h.subscriptionUseCase.IsSubscriptionActive(userID)
    // if err != nil {
    //     c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify subscription"})
    //     return
    // }
    // if !active {
    //     c.JSON(http.StatusForbidden, gin.H{"error": "Subscription inactive or expired"})
    //     return
    // }

    ebook, err := h.usecase.GetEbook(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed retrive ebook"))
        return
    }

    c.JSON(http.StatusOK, apiModels.NewSuccessResponse(ebook.FilePath, "Ebooks retrive successfully"))

}


func (h *EbookAPI) ServeEbook(c *gin.Context) { 
    var req struct {
        FilePath string `json:"fileUrl"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid request"))
        return
    }

    if req.FilePath == "" {
        c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "File path is required"))
        return
    }

    if _, err := os.Stat(req.FilePath); os.IsNotExist(err) {
        c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "File not found"))
        return
    }

    c.Header("Content-Type", "application/pdf")

    fileData, err := os.ReadFile(req.FilePath )
    if err != nil {
        c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "File to read file"))
        return
    }

    c.Data(http.StatusOK, "application/pdf", fileData)
}
