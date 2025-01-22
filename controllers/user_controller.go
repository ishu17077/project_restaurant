package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"project_restaurant/database"
	helpers "project_restaurant/helpers"
	"project_restaurant/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
		projectStage := bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "total_count", Value: 1},
			{Key: "user_items", Value: bson.D{{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}}}},
		}}}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage,
			projectStage,
		})
		defer cancel()
		//either pass an error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Listing Users!"})
			return
		}

		var allUsers []bson.M
		if err := result.All(ctx, &allUsers); err != nil {
			log.Fatal(err)

		}
		//ideally want to return all the users based on verious query parameter
		c.JSON(http.StatusOK, allUsers)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		userId := c.Param("user_id")
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding the user"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		//* Convert the JSON data coming from postman to something go lang would understand
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//* validate the above data based on user struct
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}
		//* check if email is already used by another user
		// count,err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email}) //* working
		count, err := userCollection.CountDocuments(ctx, bson.D{{Key: "email", Value: user.Email}, {Key: "phone", Value: user.Password}})
		//! let's see if this works instead of querying two times
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking email or phone"})
			log.Panic(err)

		}
		//* hash password
		password := HashPassword(*user.Password)
		user.Password = &password
		//* check if phone no is already used by another user
		// count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone}) //* working
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking email or phone"})
		// 	log.Panic(err)
		// }
		if count > 0 {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Email or Phone already in use"})
			return
		}
		//* get some extra details for user object like created_at, updated_at,ID

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()

		//* generate token and refresh tokens (generate all tokens from helper)

		token, refeshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, user.User_id)
		user.Token = &token
		user.Refresh_Token = &refeshToken
		//* if all okay, insert this new user to user collection
		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		//* return statusOk and send the result back
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User
		//* Convert the Login JSON data coming from postman to something go lang would understand
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			defer cancel()
			return
		}
		//* find a user with that email and check if exist

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found!"})
			return
		}
		//* if user exists, verify the password
		passwordIsValid, msg := verifyPassword(*foundUser.Password, *user.Password)
		defer cancel()
		if !passwordIsValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}
		//* if all goes well, then you'll generate tokens
		token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *&foundUser.User_id)

		//* update tokens - token and refresh tokens
		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		//* return StatusOK
		c.JSON(http.StatusOK, foundUser)
	}
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func verifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	check := false
	msg := "Password Incorrect"
	if err == nil {
		check = true
		msg = fmt.Sprintf("Login Success!")
	}
	return check, msg
}
