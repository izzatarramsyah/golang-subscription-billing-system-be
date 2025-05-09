// handlers/user_handlers.go
package handlers

import (
	"net/http"
	"subscription-billing-system/models"
	apiModels "subscription-billing-system/api/v1/models"

	"subscription-billing-system/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
	"os"
    "path/filepath"
)

type ProductAPI struct {
	productUseCase models.ProductUseCase
	ebookUsecase models.EbookUseCase
}

func NewProductAPI(productUseCase models.ProductUseCase, ebookUsecase models.EbookUseCase) *ProductAPI {
	return &ProductAPI{productUseCase, ebookUsecase}
}

func (h *ProductAPI) Create(c *gin.Context) {
	
	Name := c.PostForm("Name")
    Description := c.PostForm("Description")
    OwnerID := c.PostForm("OwnerID")

	// Mengambil file yang diupload
	file, err := c.FormFile("File")
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "File is required"))
		return
	}
	
	// Membuat direktori jika belum ada
	uploadPath := "./uploads/ebooks/"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.MkdirAll(uploadPath, os.ModePerm)
	}

	// Menyimpan file yang diupload
	filePath := filepath.Join(uploadPath, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to save file"))
		return
	}

	productId,err := h.productUseCase.CreateProduct(Name, Description, OwnerID);
	if  err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to save product"))
		return
	}

	// Menyimpan data ebook ke database atau sistem lain
	h.ebookUsecase.UploadEbook(Name, filePath, productId)

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(true, "Product created successfully"))
}

func (h *ProductAPI) Update(c *gin.Context) {
	productIdStr := c.PostForm("ProductId")
    productId, err := uuid.Parse(productIdStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid productId"))
        return
    }

	name := c.PostForm("Name")
	description := c.PostForm("Description")

	ownerIDStr := c.PostForm("OwnerID")
    ownerID, err := uuid.Parse(ownerIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid ownerId"))
        return
    }

	if name == "" || description == "" {
        c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Missing required fields"))
        return
    }

	input := models.Product{
		ID: 	     productId,
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
	}
	
	file, err := c.FormFile("File")
	if err != nil && file != nil { 
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "File is required"))
		return
	}

	if file != nil {
		uploadPath := "./uploads/ebooks/"
		if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
			if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
				c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to create directory"))
				return
			}
		}

		filePath := filepath.Join(uploadPath, file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to save file"))
			return
		}

		if err := h.ebookUsecase.UpdateEbook(filePath, productId); err != nil {
			c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to update ebook"))
			return
		}
	}

	if err := h.productUseCase.UpdateProduct(&input); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to update product"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(true, "Product updated successfully"))
}


func (h *ProductAPI) List(c *gin.Context) {

	products, err := h.productUseCase.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to retrive products"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(products, "Products retrieved successfully"))
}

func (h *ProductAPI) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	productId, err := uuid.Parse(idParam)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid Product"))
		return
	}

	product, err := h.productUseCase.GetProductByID(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest,"Failed to retrive product"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(product, "Product retrieved successfully"))
}

func (h *ProductAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.productUseCase.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to delete product"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(true, "Product deleted successfully"))
}

func (h *ProductAPI) GetByOwnerID(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Unauthorized"))
		return
	}

	products, err := h.productUseCase.GetProductByOwnerID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest,"Failed to retrive products"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(products, "Product retrieved successfully"))
}

func (h *ProductAPI) UpdateStatusProduct(c *gin.Context) {
	idParam := c.Param("id")
	productId, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid Product ID"))
		return
	}

	var req struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid request"))
		return
	}

	if err := h.productUseCase.UpdateStatusProduct(productId, req.Status); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to update product status"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(true, "Product status updated successfully"))
}