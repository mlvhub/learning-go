package mongo_test

import (
	"testing"

	"github.com/mlvhub/learning-go/mongodb-rest-api/pkg"
	"github.com/mlvhub/learning-go/mongodb-rest-api/pkg/mock"
	"github.com/mlvhub/learning-go/mongodb-rest-api/pkg/mongo"
)

var mongoConfig = &root.MongoConfig{
	URL:       "mongodb://test:test@testcluster-shard-00-00-viokg.mongodb.net:27017,testcluster-shard-00-01-viokg.mongodb.net:27017,testcluster-shard-00-02-viokg.mongodb.net:27017/test?ssl=true&replicaSet=TestCluster-shard-0&authSource=admin",
	DBName:    "test",
	UserTable: "users",
}

func Test_UserService(t *testing.T) {
	t.Run("CreateUser", createUser_should_insert_user_into_mongo)
}

func createUser_should_insert_user_into_mongo(t *testing.T) {
	//Arrange
	session, err := mongo.NewSession(mongoConfig.URL)
	if err != nil {
		t.Errorf("Unable to connect to mongo: %s", err)
	}
	defer func() {
		session.DropDatabase(mongoConfig.DBName)
		session.Close()
	}()
	cryptoMock := &mock.Crypto{}
	userService := mongo.NewUserService(session.Copy(), mongoConfig, cryptoMock)

	testUsername := "integration_test_user"
	testPassword := "integration_test_password"
	user := root.User{
		Username: testUsername,
		Password: testPassword,
	}

	//Act
	err = userService.Create(&user)

	//Assert
	if err != nil {
		t.Errorf("Unable to create user: %s", err)
	}

	var results []root.User
	session.GetCollection(mongoConfig.DBName, mongoConfig.UserTable).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Errorf("Incorrect number of results. Expected `1`, got: `%d`", count)
	}
	if results[0].Username != user.Username {
		t.Errorf("Incorrect Username. Expected `%s`, Got: `%s`", testUsername, results[0].Username)
	}
}
