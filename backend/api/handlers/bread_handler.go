package handlers

import (
	"BakeMeets/api/models"
	"BakeMeets/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BreadHandler struct {
	Br *repository.BreadRepository
}

type FavoriteRequest struct {
	BreadID    int  `json:"breadId"`
	IsFavorite bool `json:"isFavorite"`
}

func (h *BreadHandler) GetAllBreads(c echo.Context) error {
	breads, err := h.Br.GetAllBreads()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, breads)
}

func (h *BreadHandler) GetBreadByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bread, err := h.Br.GetBreadByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, bread)
}

func (h *BreadHandler) CreateBread(c echo.Context) error {
	var bread models.Bread
	if err := c.Bind(&bread); err != nil {
		return err
	}
	if bread.ImageURL == "" {
		bread.ImageURL = "noimage.webp"
	}
	createdBread, err := h.Br.CreateBread(bread)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, createdBread)
}

func (h *BreadHandler) UpdateBread(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var bread models.Bread
	if err := c.Bind(&bread); err != nil {
		return err
	}
	bread.ID = uint(id)
	if bread.ImageURL == "" {
		bread.ImageURL = "noimage.webp"
	}
	updatedBread, err := h.Br.UpdateBread(bread)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedBread)
}

func (h *BreadHandler) DeleteBread(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bread, err := h.Br.GetBreadByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := h.Br.DeleteBread(bread); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{"result": "success"})
}

func (h *BreadHandler) UpdateFavorite(c echo.Context) error {
	req := new(FavoriteRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	bread, err := h.Br.GetBreadByID(uint(req.BreadID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// データベースまたはその他のストレージで状態を更新する
	if req.IsFavorite == true {
		bread.FavoriteCount++
	} else {
		if bread.FavoriteCount > 0 {
			bread.FavoriteCount--
		}
	}
	updatedBread, err := h.Br.UpdateBread(bread)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedBread)
}
