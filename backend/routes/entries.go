package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ImArnav19/go-react-calorieTracker/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var entryCollection *mongo.Collection = OpenCollection(Client, "calories")

func AddEntry(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100) //context model
	var entry models.Entry

	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Binderror": err.Error()}) //Binding JSON from BODY
		fmt.Println(err)
		return
	}

	validateErr := validate.Struct(entry) //validating err

	if validateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Valerror": validateErr.Error()})
		fmt.Println(validateErr)
		return
	}

	entry.ID = primitive.NewObjectID()                          //new id creation
	results, insertErr := entryCollection.InsertOne(ctx, entry) //saving
	if insertErr != nil {
		msg := "Insert Problem in DB!"
		c.JSON(http.StatusInternalServerError, gin.H{"Inserror": msg})
		fmt.Println(insertErr)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, results)

}

func AllEntry(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second) //contex setup for timeout

	var entries []bson.M                               //bson marshal object
	cursor, err := entryCollection.Find(ctx, bson.M{}) //find all entries of bson.M{} interface

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //error handling
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &entries); err != nil { //assign it to entries
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	fmt.Println(entries)
	c.JSON(http.StatusOK, entries) //ok

}

func GetEntry(c *gin.Context) {
	entryID := c.Params.ByName("id") //params
	docID, _ := primitive.ObjectIDFromHex(entryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second) //defined by ctx

	var entry bson.M

	if err := entryCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	fmt.Println(entry)
	c.JSON(http.StatusOK, entry)

}

func GetEntriesByIngredient(c *gin.Context) {

	ingredients := c.Params.ByName("id")

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entries []bson.M

	cursor, err := entryCollection.Find(ctx, bson.M{"ingredients": ingredients})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //error handling
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &entries); err != nil { //assign it to entries
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(entries)
	c.JSON(http.StatusOK, entries)

}

func UpdateEntry(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entry models.Entry
	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validateErr := validate.Struct(entry)
	if validateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validateErr.Error()})
		fmt.Println(validateErr)
		return
	}

	result, UpdateErr := entryCollection.ReplaceOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{
			"Dish":        entry.Dish,
			"Fat":         entry.Fat,
			"Ingredients": entry.Ingredients,
			"Calories":    entry.Calories,
		},
	)

	if UpdateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": UpdateErr.Error()})
		fmt.Println(UpdateErr)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)

}

func UpdateIngredient(c *gin.Context) {

	id := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(id)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	type Ingredient struct {
		Ingredients *string `json:"ingredients"`
	}
	var ingredient Ingredient

	if err := c.BindJSON(&ingredient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	result, err := entryCollection.UpdateOne(ctx, bson.M{"_id": docID},
		bson.D{{"$set", bson.D{{"ingredients", ingredient.Ingredients}}}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)

}

func DeleteEntry(c *gin.Context) {
	entryID := c.Params.ByName("id")               //with context access to id
	docID, _ := primitive.ObjectIDFromHex(entryID) //converting to hex code

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second) //assigning a timeout

	result, err := entryCollection.DeleteOne(ctx, bson.M{"_id": docID}) //execution, bson.M -> marshal

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //post error
		fmt.Println(err)
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, result.DeletedCount)

}
