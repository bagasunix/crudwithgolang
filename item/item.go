package item

import (
	"context"
	"crudwithgolang/db"
	"crudwithgolang/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"go.mongodb.org/mongo-driver/bson"
)

type Item struct {
	ItemId      string    `bson:"item_id,omitempty" json:"item_id"`
	Name        string    `form:"name" bson:"name,omitempty" json:"name"`
	Price       int       `bson:"price,omitempty" json:"price"`
	Type        string    `form:"type" bson:"type" json:"type"`
	Description string    `form:"description" bson:"description,omitempty" json:"description"`
	Slug        string    `bson:"slug,omitempty" json:"slug"`
	CreatedAt   time.Time `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at,omitempty" json:"updated_at"`
}

var itemColl = db.DB.Collection("item")

func AddItem(c *gin.Context) {
	var itemData Item

	if err := c.BindJSON(&itemData); err != nil {
		utils.SendBadRequest(c, err)
		return
	}

	// create and uuid string to be stored at db
	newItemId, err := utils.CreateUUIDStr()
	if err != nil {
		utils.SendInternalServerError(c, err)
		return
	}

	newItem := bson.M{
		"item_id":     newItemId,
		"name":        itemData.Name,
		"price":       itemData.Price,
		"type":        itemData.Type,
		"description": itemData.Description,
		"slug":        slug.MakeLang(itemData.Name, "en"),
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
	}

	_, err = itemColl.InsertOne(context.Background(), newItem)
	if err != nil {
		utils.SendInternalServerError(c, err)
		return
	}

	utils.SendSuccess(c, "new item added successfully")
}

func GetItem(c *gin.Context) {
	cursor, err := itemColl.Find(context.Background(), bson.M{})
	if err != nil {
		utils.SendInternalServerError(c, err)
		return
	}

	item := []Item{}
	if err = cursor.All(db.Ctx, &item); err != nil {
		utils.SendInternalServerError(c, err)
		return
	}

	utils.SendData(c, item)
}

func GetItemBySlug(c *gin.Context) {
	slug := c.Param("id")
	var item Item

	if err := itemColl.FindOne(context.Background(), bson.M{"slug": slug}).Decode(&item); err != nil {
		utils.SendInternalServerError(c, err)
		return
	}

	utils.SendData(c, item)
}

func EditItemBySlug(c *gin.Context) {
	id := c.Param("id")
	var itemData Item

	if err := c.BindJSON(&itemData); err != nil {
		utils.SendBadRequest(c, err)
		return
	}

	// add updated_at, updated_by
	itemData.Slug = slug.MakeLang(itemData.Name, "en")
	itemData.UpdatedAt = time.Now()

	// update process
	pByte, err := bson.Marshal(itemData)
	if err != nil {
		utils.SendInternalServerError(c, err)
		return
	}

	var update bson.M
	err = bson.Unmarshal(pByte, &update)
	if err != nil {
		utils.SendInternalServerError(c, err)
		return
	}

	_, err = itemColl.UpdateOne(
		context.Background(),
		bson.M{"slug": id},
		bson.M{"$set": update},
	)
	if err != nil {
		utils.SendInternalServerError(c, err)
		return
	}

	utils.SendSuccess(c, "item edited successfully")
}

func DeleteItemBySlug(c *gin.Context) {
	id := c.Param("id")

	_, err := itemColl.DeleteOne(context.Background(), bson.M{"slug": id})
	if err != nil {
		utils.SendInternalServerError(c, err)
		return
	}

	utils.SendSuccess(c, "career deleted successfully")
}
