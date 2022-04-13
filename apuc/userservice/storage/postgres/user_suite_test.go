package postgres

import (
	"github/Services/apuc/userservice/config"
	pb "github/Services/apuc/userservice/genproto/user_service"
	"github/Services/apuc/userservice/pkg/db"
	"github/Services/apuc/userservice/storage/repo"

	"testing"

	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.UserStorageI
}

func (suite *UserRepoTestSuite) SetupSuite() {
	pgPool, cleanup := db.ConnectDbForSuite(config.Load())
	suite.Repository = NewUserRepo(pgPool)
	suite.CleanUpFunc = cleanup
}

func (suite *UserRepoTestSuite) TestUserCRUD() {
	User := pb.User{
		Id:       "fe66a4bb-4c19-40a1-a5a2-ae8b3eb0cb62",
		Name:    "garry1",
		Username:     "hurry",
		City: 	"Tashkent",
	}


	User1, err := suite.Repository.Create(&User)
	suite.Nil(err)
	suite.NotNil(User1)
	
	getUser, err := suite.Repository.Get(&pb.ById{ Userid: User.Id })
	suite.Nil(err)
	suite.NotNil(getUser, "User must not be nil")
	suite.Equal(User.Id, getUser.Id, "Asignees must match")

	User.Name = "New user"
	updatedUser, err := suite.Repository.Update(&User)
	suite.Nil(err)
	suite.Equal(updatedUser.Name, User.Name, "Titles must match")
	
	_, err = suite.Repository.Delete(&pb.ById{Userid: User.Id})
	suite.Nil(err)

}

func (suite *UserRepoTestSuite) TearDownSuite() {
	suite.CleanUpFunc()
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
