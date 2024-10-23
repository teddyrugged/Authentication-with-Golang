package controllers

import (
	"golang/database"
	"net/http"
	"time"
	"fmt"
	"log"
	"context"
	"strconv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-playground/validator/v10"
	helper "golang/helpers"
	"golang/models"
	"golang/helpers"
	"github.com/gin-gonic/gin"
)





var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()
 
func HashPassword(password string) string {



}



func VerifyPassword(userPassword string, providedPassword string) (bool, string) {


}


func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
		}
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email is already in use"})
		}

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserId = user.ID.Hex()
		// refresh_token, _ := helper.GenerateAllToken(*user)
		// user.Token = refresh_token.AccessToken
		Token, RefreshToken, _ := helper.GenerateAllToken(*user.Email, *user.FirstName, *user.LastName)
		user.Token = &token
		user.RefreshToken = &RefreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr !=nul{
			msg :=fmt.Sprint("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error" :msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)





		password := HashPassword(*user.Password)
		user.Password = &password


	}

	


}

func Login() gin.HandlerFunc {


}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helper.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))


	}


}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User  
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)

	}

}
