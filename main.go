package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"recipes-api/handlers"
	"time"
)

type Recipe struct {
	//ID           string             `json:"id"`
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name"`
	Tags         []string           `json:"tags"`
	Ingredients  []string           `json:"ingredients"`
	Instructions []string           `json:"instructions"`
	PublishedAt  time.Time          `json:"publishedAt"`
}

//var recipes []Recipe

//var ctx context.Context

//var err error

//var client *mongo.Client

//var collection *mongo.Collection
var recipesHandler *handlers.RecipesHandler

func init() {
	//recipes = make([]Recipe, 0)
	//file, _ := ioutil.ReadFile("recipes.json")
	//_ = json.Unmarshal([]byte(file), &recipes)

	// 上下文
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_DOCKER_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("connected to MongoDB")
	collection := client.Database("test").Collection("recipes")
	recipesHandler = handlers.NewRecipesHandler(ctx, collection)
	//var listOfRecipes []interface{}
	//for _, recipe := range recipes {
	//	listOfRecipes = append(listOfRecipes, recipe)
	//}
	//collection := client.Database("test").Collection("recipes")
	//insertManyResult, err := collection.InsertMany(ctx, listOfRecipes)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("insert recipes:", len(insertManyResult.InsertedIDs))
}

//func NewRecipeHandler(c *gin.Context) {
//	//var recipe Recipe
//	//if err := c.ShouldBindJSON(&recipe); err != nil {
//	//	c.JSON(http.StatusBadRequest, gin.H{
//	//		"error": err.Error()})
//	//	return
//	//}
//	//recipe.ID = xid.New().String()
//	//recipe.PublishedAt = time.Now()
//	//recipes = append(recipes, recipe)
//	//c.JSON(http.StatusOK, recipe)
//	var recipe Recipe
//	if err := c.ShouldBindJSON(&recipe); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	recipe.ID = primitive.NewObjectID()
//	recipe.PublishedAt = time.Now()
//	_, err = collection.InsertOne(ctx, recipe)
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while insert data"})
//		return
//	}
//	c.JSON(http.StatusOK, recipe)
//}
//
//func ListRecipesHandler(c *gin.Context) {
//	//use file
//	//c.JSON(http.StatusOK, recipes)
//	cur, err := collection.Find(ctx, bson.M{})
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	defer cur.Close(ctx)
//	recipes := make([]Recipe, 0)
//	for cur.Next(ctx) {
//		var recipe Recipe
//		cur.Decode(&recipe)
//		recipes = append(recipes, recipe)
//	}
//	c.JSON(http.StatusOK, recipes)
//
//}
//
//func UpdateRecipeHandler(c *gin.Context) {
//	id := c.Param("id")
//	var recipe Recipe
//	if err := c.ShouldBindJSON(&recipe); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": err.Error()})
//		return
//	}
//	objectId, _ := primitive.ObjectIDFromHex(id)
//	_, err = collection.UpdateOne(ctx,
//		bson.M{"_id": objectId},
//		bson.D{{"$set", bson.D{
//			{"name", recipe.Name},
//			{"instructions", recipe.Instructions},
//			{"ingredients", recipe.Ingredients},
//			{"tags", recipe.Tags},
//		}}})
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"message": "update message"})
//}
//
//func DeleteRecipeHandler(c *gin.Context) {
//	id := c.Param("id")
//
//	objectId, _ := primitive.ObjectIDFromHex(id)
//	_, err := collection.DeleteOne(ctx, bson.M{"_id": objectId})
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"message": "Recipe has been deleted"})
//}

//func SearchRecipesHandler(c *gin.Context) {
//	tag := c.Query("tag")
//	listOfRecipes := make([]Recipe, 0)
//	for i := 0; i < len(recipes); i++ {
//		found := false
//		for _, t := range recipes[i].Tags {
//			// 不去分大小写比较
//			if strings.EqualFold(t, tag) {
//				found = true
//			}
//		}
//		if found {
//			listOfRecipes = append(listOfRecipes, recipes[i])
//		}
//	}
//	c.JSON(http.StatusOK, listOfRecipes)
//}
func main() {
	router := gin.Default()
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
	router.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	router.Run()
}
