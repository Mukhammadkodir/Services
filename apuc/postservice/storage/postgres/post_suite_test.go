package postgres

import (
	"github/Services/apuc/postservice/config"
	pb "github/Services/apuc/postservice/genproto/post_service"
	"github/Services/apuc/postservice/pkg/db"
	"github/Services/apuc/postservice/storage/repo"

	"testing"

	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.PostStorageI
}

func (suite *UserRepoTestSuite) SetupSuite() {
	pgPool, cleanup := db.ConnectDbForSuite(config.Load())
	suite.Repository = NewPostRepo(pgPool)
	suite.CleanUpFunc = cleanup
}

func (suite *UserRepoTestSuite) TestUserCRUD() {
	a := []string{"yuio ghjk", "ghjk tyui bnm", "werty dffgghj bnmuyi"}
	Post := pb.Post{
		Id:      "fe66a4bb-4c19-40a1-a5a2-ae8b3eb0cb62",
		Title:   "tyuib jkjkl",
		Comment: "pariatur",
		UserId:  "fe66a4bb-4c19-40a1-a5a2-ae8b3eb0cb62",
	}
	for _, i := range a {
		Post.Image = append(Post.Image, &pb.Photo{Image: i})
	}

	Post1, err := suite.Repository.Create(&Post)
	suite.Nil(err)
	suite.NotNil(Post1)

	getPost, err := suite.Repository.Get(&pb.ById{Userid: Post.Id})
	suite.Nil(err)
	suite.NotNil(getPost, "User must not be nil")
	suite.Equal(Post.Id, getPost.Id, "Asignees must match")

	Post.Title = "New Post"
	updatedPost, err := suite.Repository.Update(&Post)
	suite.Nil(err)
	suite.Equal(updatedPost.Title, Post.Title, "Titles must match")

	_, err = suite.Repository.Delete(&pb.ById{Userid: Post.Id})
	suite.Nil(err)

}

func (suite *UserRepoTestSuite) TearDownSuite() {
	suite.CleanUpFunc()
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
