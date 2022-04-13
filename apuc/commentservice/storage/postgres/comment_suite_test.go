package postgres

import (
	"github/Services/apuc/commentservice/config"
	pb "github/Services/apuc/commentservice/genproto/comment_service"
	"github/Services/apuc/commentservice/pkg/db"
	"github/Services/apuc/commentservice/storage/repo"

	"testing"

	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.CommentStorageI
}

func (suite *UserRepoTestSuite) SetupSuite() {
	pgPool, cleanup := db.ConnectDbForSuite(config.Load())
	suite.Repository = NewCommentRepo(pgPool)
	suite.CleanUpFunc = cleanup
}

func (suite *UserRepoTestSuite) TestUserCRUD() {
	User := pb.Comment{
		Id:       "fe66a4bb-4c19-40a1-a5a2-ae8b3eb0cb62",
		UserId:    "fe66a4bb-4c19-40a1-a5a2-ae8b3eb0cb62",
		PostId:     "fe66a4bb-4c19-40a1-a5a2-ae8b3eb0cb62",
		Text: 	"Tashkent",
	}


	User1, err := suite.Repository.Create(&User)
	suite.Nil(err)
	suite.NotNil(User1)
	
	getUser, err := suite.Repository.Get(&pb.ById{ Id: User.Id })
	suite.Nil(err)
	suite.NotNil(getUser, "User must not be nil")
	suite.Equal(User.Id, getUser.Id, "Asignees must match")

	User.Text = "New user"
	updatedUser, err := suite.Repository.Update(&User)
	suite.Nil(err)
	suite.Equal(updatedUser.Text, User.Text, "Titles must match")
	
	_, err = suite.Repository.Delete(&pb.ById{Id: User.Id})
	suite.Nil(err)

}

func (suite *UserRepoTestSuite) TearDownSuite() {
	suite.CleanUpFunc()
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
