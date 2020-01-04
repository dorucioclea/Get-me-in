[1mdiff --git a/account-service/internal/repo.go b/account-service/internal/repo.go[m
[1mindex fe2adcc..c7e43ee 100644[m
[1m--- a/account-service/internal/repo.go[m
[1m+++ b/account-service/internal/repo.go[m
[36m@@ -1,10 +1,11 @@[m
 package internal[m
 [m
 import ([m
[31m-	"github.com/aws/aws-sdk-go/aws/credentials"[m
 	"github/Get-me-in/account-service/configs"[m
 	"github/Get-me-in/pkg/dynamodb"[m
 	"net/http"[m
[32m+[m
[32m+[m	[32m"github.com/aws/aws-sdk-go/aws/credentials"[m
 )[m
 [m
 func TestFunc(w http.ResponseWriter, r *http.Request) {[m
[36m@@ -23,11 +24,19 @@[m [mfunc CreateUser(w http.ResponseWriter, r *http.Request) {[m
 	dynamodb.CreateItem(w, DecodeToDynamoAttribute(w, r))[m
 }[m
 [m
[31m-func Login(w http.ResponseWriter, r *http.Request){[m
[32m+[m[32m// func GetUser(w http.ResponseWriter, r *http.Request) {[m
[32m+[m[32m// 	return[m
[32m+[m[32m// }[m
[32m+[m
[32m+[m[32mfunc DeleteUser(w http.ResponseWriter, r *http.Request) {[m
[32m+[m	[32mdynamo.DeleteUser()[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc Login(w http.ResponseWriter, r *http.Request) {[m
 [m
 	body, found := dynamodb.GetItem(w, configs.FIND_BY_EMAIL, DecodeToJSON(w, r.Body).Email)[m
 [m
 	if found {[m
 		w.Write([]byte(body))[m
 	}[m
[31m-}[m
\ No newline at end of file[m
[32m+[m[32m}[m
[1mdiff --git a/pkg/dynamodb/delete_item.go b/pkg/dynamodb/delete_item.go[m
[1mindex 0f07a9d..41541bd 100644[m
[1m--- a/pkg/dynamodb/delete_item.go[m
[1m+++ b/pkg/dynamodb/delete_item.go[m
[36m@@ -1,33 +1,22 @@[m
 package dynamodb[m
 [m
[31m-/*[m
 import ([m
[31m-	"encoding/json"[m
 	"fmt"[m
[32m+[m	[32m"net/http"[m
[32m+[m
 	"github.com/aws/aws-sdk-go/aws"[m
 	"github.com/aws/aws-sdk-go/service/dynamodb"[m
[31m-	"net/http"[m
 )[m
 [m
[32m+[m[32m// Delete User ..[m
 func DeleteUser(w http.ResponseWriter, r *http.Request) {[m
 [m
[31m-	var u internal.User[m
[31m-[m
[31m-	// Try to decode the request body into the struct. If there is an error,[m
[31m-	// respond to the client with the error message and a 400 status code.[m
[31m-	err := json.NewDecoder(r.Body).Decode(&u)[m
[31m-[m
[31m-	if err != nil {[m
[31m-		http.Error(w, err.Error(), http.StatusBadRequest)[m
[31m-		return[m
[31m-	}[m
[31m-[m
[31m-	user_id := u.uuid[m
[32m+[m	[32mu := DecodeToJSON(w, r.Body)[m
 [m
 	input := &dynamodb.DeleteItemInput{[m
 		Key: map[string]*dynamodb.AttributeValue{[m
[31m-			"uuid": {[m
[31m-				N: aws.String(movieYear),[m
[32m+[m			[32m"email": {[m
[32m+[m				[32mN: aws.String(u.Email),[m
 			},[m
 		},[m
 		TableName: aws.String("Users"),[m
[36m@@ -42,4 +31,4 @@[m [mfunc DeleteUser(w http.ResponseWriter, r *http.Request) {[m
 	}[m
 [m
 	fmt.Println("Deleted User from table.")[m
[31m-}*/[m
\ No newline at end of file[m
[32m+[m[32m}[m
[1mdiff --git a/pkg/dynamodb/get_item.go b/pkg/dynamodb/get_item.go[m
[1mindex 30a07f5..505f117 100644[m
[1m--- a/pkg/dynamodb/get_item.go[m
[1m+++ b/pkg/dynamodb/get_item.go[m
[36m@@ -1,35 +1,24 @@[m
 package dynamodb[m
 [m
[31m-import "net/http"[m
[31m-[m
[31m-/*[m
 import ([m
[31m-	"encoding/json"[m
 	"fmt"[m
[32m+[m	[32m"net/http"[m
[32m+[m	[32m"github/Get-me-in/account-service/internal"[m
 	"github.com/aws/aws-sdk-go/aws"[m
 	"github.com/aws/aws-sdk-go/service/dynamodb"[m
 	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"[m
[31m-	"net/http"[m
 )[m
 [m
[32m+[m[32m// Retrieve user[m
 func GetUser(w http.ResponseWriter, r *http.Request) {[m
 [m
[31m-	var u internal.User[m
[31m-[m
[31m-	// Try to decode the request body into the struct. If there is an error,[m
[31m-	// respond to the client with the error message and a 400 status code.[m
[31m-	err := json.NewDecoder(r.Body).Decode(&u)[m
[31m-[m
[31m-	if err != nil {[m
[31m-		http.Error(w, err.Error(), http.StatusBadRequest)[m
[31m-		return[m
[31m-	}[m
[32m+[m	[32mvar u DecodeToJSON(w, r.Body)[m
 [m
 	result, err := svc.GetItem(&dynamodb.GetItemInput{[m
 		TableName: aws.String("Users"),[m
 		Key: map[string]*dynamodb.AttributeValue{[m
[31m-			"uuid": {[m
[31m-				N: aws.String(u.SearchParam),[m
[32m+[m			[32m"email": {[m
[32m+[m				[32mN: aws.String(u.Email),[m
 			},[m
 		},[m
 	})[m
[36m@@ -58,7 +47,3 @@[m [mfunc GetUser(w http.ResponseWriter, r *http.Request) {[m
 	w.WriteHeader(http.StatusOK)[m
 	return item[m
 }[m
[31m-*/[m
[31m-func GetItem(w http.ResponseWriter, s string, v string) (string, bool){[m
[31m-	return "userdetails",true[m
[31m-}[m
\ No newline at end of file[m
[1mdiff --git a/pkg/dynamodb/update_item.go b/pkg/dynamodb/update_item.go[m
[1mindex dc21954..df15220 100644[m
[1m--- a/pkg/dynamodb/update_item.go[m
[1m+++ b/pkg/dynamodb/update_item.go[m
[36m@@ -1,10 +1,34 @@[m
 package dynamodb[m
 [m
[31m-import ([m
[31m-	"fmt"[m
[31m-	"net/http"[m
[31m-)[m
[31m-[m
[31m-func UpdateUser(w http.ResponseWriter, r *http.Request) {[m
[31m-	fmt.Println("")[m
[31m-}[m
\ No newline at end of file[m
[32m+[m[32m// //Need to do[m
[32m+[m[32m// func UpdateUser(w http.ResponseWriter, r *http.Request) {[m
[32m+[m
[32m+[m[32m// 	tableName := "Movies"[m
[32m+[m[32m// 	movieName := "The Big New Movie"[m
[32m+[m[32m// 	movieYear := "2015"[m
[32m+[m[32m// 	movieRating := "0.5"[m
[32m+[m
[32m+[m[32m// 	input := &dynamodb.UpdateItemInput{[m
[32m+[m[32m// 		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{[m
[32m+[m[32m// 			":r": {[m
[32m+[m[32m// 				N: aws.String(movieRating),[m
[32m+[m[32m// 			},[m
[32m+[m[32m// 		},[m
[32m+[m[32m// 		TableName: aws.String(tableName),[m
[32m+[m[32m// 		Key: map[string]*dynamodb.AttributeValue{[m
[32m+[m[32m// 			"Email": {[m
[32m+[m[32m// 				N: aws.String(movieYear),[m
[32m+[m[32m// 			},[m
[32m+[m[32m// 		},[m
[32m+[m[32m// 		ReturnValues:     aws.String("UPDATED_NEW"),[m
[32m+[m[32m// 		UpdateExpression: aws.String("set Rating = :r"),[m
[32m+[m[32m// 	}[m
[32m+[m
[32m+[m[32m// 	_, err := svc.UpdateItem(input)[m
[32m+[m[32m// 	if err != nil {[m
[32m+[m[32m// 		fmt.Println(err.Error())[m
[32m+[m[32m// 		return[m
[32m+[m[32m// 	}[m
[32m+[m
[32m+[m[32m// 	fmt.Println("Successfully updated '" + movieName + "' (" + movieYear + ") rating to " + movieRating)[m
[32m+[m[32m// }[m
