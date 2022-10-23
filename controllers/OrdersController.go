package controllers

import (
	_"fmt"
	"net/http"

	"rest-api-project/database"
	"rest-api-project/models"
	"rest-api-project/structs"

	"github.com/gin-gonic/gin"
)

func CreateOrders(ctx *gin.Context) {
	db := database.GetDB()
	var items models.Items
	var orders models.Orders
	body := structs.Order{}

	
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	items.Item_Code = body.Items[0].ItemCode
	items.Description = body.Items[0].Description
	items.Quantity = uint(body.Items[0].Quantity)
	orders.Customer_name = body.CustomerName
	orders.OrderedAt = body.OrderedAt
	items.Order_Id = orders.Order_Id
	
	err := db.Create(&items).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err = db.Create(&orders).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	// debugging
	
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success creating data",
		"order data" : orders,
		"item data" : items,
	})
}

func GetOrders(ctx *gin.Context) {
	db := database.GetDB()
	orders := models.Orders{}

	err := db.Find(&orders).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	ctx.JSON(http.StatusAccepted, orders)
	
}

func UpdateOrders(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	db := database.GetDB()
	orders := models.Orders{}
	items := models.Items{}
	body := structs.Order{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_message": "cannot bind data",
			"error" : err,
		})
	}

	bodyItems := body.Items[0]

	err := db.Model(&orders).Where("order_id = ?", orderId).
	Updates(models.Orders{
		Customer_name: body.CustomerName,
		OrderedAt: body.OrderedAt,
	}).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_message" : "update order data failed",
			"error" : err,
		})
	}

	err = db.Model(&items).Where("order_id=?", orderId).
	Updates(models.Items{
		Item_Code: bodyItems.ItemCode,
		Description: bodyItems.Description,
		Quantity: uint(bodyItems.Quantity),
	}).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_message":"update items data failed",
			"error":err,
		})
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message":"data successfully updated",
	})
}

func DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()
	orderId := ctx.Param("orderId")
	orders := models.Orders{}

	err := db.Where("order_id=?",orderId).Delete(&orders).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error_message" : "failed to delete order",
			"error" : err,
		})
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message" : "successfully deleted order",
	})
}
