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
