package controller

import (
    "net/http"

    "github.com/andersonlira/item-api/gateway/txtdb"
    "github.com/andersonlira/item-api/domain"
    "github.com/andersonlira/item-api/usecase"
	"github.com/labstack/echo/v4"

)


//GetItemList return all objects 
func GetItemList(c echo.Context) error {

    list := txtdb.GetItemList()

	return c.JSON(http.StatusOK, list)
}

func GetItemByID(c echo.Context) error {
    ID := c.Param("id")
    it, err := txtdb.GetItemByID(ID)
    if err != nil {
        return c.JSON(http.StatusNotFound,it)
    }
    return c.JSON(http.StatusOK, it)
}

func SaveItem(c echo.Context) error {
    it := domain.Item{}
    c.Bind(&it)
    it = txtdb.SaveItem(it)
    return c.JSON(http.StatusCreated, it)
}

func UpdateItem(c echo.Context) error {
    ID := c.Param("id")
    it := domain.Item{}
    c.Bind(&it)
    it = txtdb.UpdateItem(ID,it)
    return c.JSON(http.StatusOK, it)
}

func DeleteItem(c echo.Context) error {
    ID := c.Param("id")
    result := txtdb.DeleteItem(ID)
    return c.JSON(http.StatusOK, result)
}


func LowestPrice(c echo.Context) error {
    ID := c.Param("id")
    it := domain.Price{}
    c.Bind(&it)
    return usecase.LastPrice(ID, it,usecase.Low)
}

func HighestPrice(c echo.Context) error {
    ID := c.Param("id")
    it := domain.Price{}
    c.Bind(&it)
    return usecase.LastPrice(ID, it,usecase.High)
}


func LastPrice(c echo.Context) error {
    ID := c.Param("id")
    it := domain.Price{}
    c.Bind(&it)
    return usecase.LastPrice(ID, it,usecase.Last)
}
