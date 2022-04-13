package service

import (
	"context"

	p "github/Services/apuc/postservice/genproto/comment_service"
	pb "github/Services/apuc/postservice/genproto/post_service"
	l "github/Services/apuc/postservice/pkg/logger"
	grpcClient "github/Services/apuc/postservice/service/grpc_client"
	"github/Services/apuc/postservice/storage"

	"github.com/jmoiron/sqlx"
)

//PostService ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.GrpcClientI
}

//NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger, client grpcClient.GrpcClientI) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *PostService) Create(ctx context.Context, req *pb.Post) (*pb.Post, error) {

	post, err := s.storage.Post().Create(req)
	if err != nil {
		s.logger.Error("Failed create post", l.Error(err))
		return nil, nil
	}

	comment, err := s.client.CommentService().Create(context.Background(), &p.Comment{
		UserId: post.UserId,
		PostId: post.Id,
		Text:   req.Comment,
	})

	if err != nil {
		s.logger.Error("Error creating comment from comment_service", l.Error(err))
		return nil, err
	}
	post.Comment = comment.Text

	return post, nil
}

func (s *PostService) Delete(ctx context.Context, req *pb.ById) (*pb.Empty, error) {

	_, err := s.storage.Post().Delete(req)
	if err != nil {
		s.logger.Error("Failed Delete post", l.Error(err))
		return nil, nil
	}

	comment, err := s.client.CommentService().Delete(context.Background(), &p.ById{Id: req.Userid})

	if err != nil {
		s.logger.Error("Error delete comment from comment_service", l.Error(err))
		return nil, err
	}

	return (*pb.Empty)(comment), nil
}

func (s *PostService) Update(ctx context.Context, req *pb.Post) (*pb.Post, error) {

	post, err := s.storage.Post().Update(req)
	if err != nil {
		s.logger.Error("Failed update post", l.Error(err))
		return nil, nil
	}

	comment, err := s.client.CommentService().Update(context.Background(), &p.Comment{
		UserId: post.UserId,
		PostId: post.Id,
		Text:   req.Comment,
	})

	if err != nil {
		s.logger.Error("Error updating comment from comment_service", l.Error(err))
		return nil, err
	}
	post.Comment = comment.Text

	return post, nil
}

func (s *PostService) List(ctx context.Context, req *pb.ListReq) (*pb.ListResp, error) {

	resp, err := s.storage.Post().List(req)
	if err != nil {
		s.logger.Error("Error getting list Posts", l.Error(err))
		return nil, err
	}

	list := pb.ListResp{}
	for _, i := range resp.Posts {

		comment, err := s.client.CommentService().Get(context.Background(), &p.ById{Id: i.Id})
		if err != nil {
			s.logger.Error("Error getting comment from comment_service", l.Error(err))
			return nil, err
		}
		i.Comment = comment.Text

		list.Posts = append(list.Posts, i)
	}

	list.Count = resp.Count

	return &list, nil
}

func (s *PostService) GetById(ctx context.Context, req *pb.ByUId) (*pb.ListResp, error) {

	posts, err := s.storage.Post().GetById(req)
	if err != nil {
		s.logger.Error("Failed to getting post's list with user_id", l.Error(err))
		return nil, nil
	}

	list := pb.ListResp{}
	for _, i := range posts.Posts {

		comment, err := s.client.CommentService().Get(context.Background(), &p.ById{Id: i.Id})
		if err != nil {
			s.logger.Error("Error getting comment from comment_service", l.Error(err))
			return nil, err
		}
		i.Comment = comment.Text

		list.Posts = append(list.Posts, i)
	}

	list.Count = posts.Count

	return &list, nil
}

func (s *PostService) Get(ctx context.Context, req *pb.ById) (*pb.Post, error) {

	post, err := s.storage.Post().Get(req)
	if err != nil {
		s.logger.Error("Failed to get post with id", l.Error(err))
		return nil, nil
	}

	comment, err := s.client.CommentService().Get(context.Background(), &p.ById{Id: post.Id})

	if err != nil {
		s.logger.Error("Error getting comment from comment_service", l.Error(err))
		return nil, err
	}

	post.Comment = comment.Text

	return post, nil
}

func (s *PostService) DeleteByUser(ctx context.Context, req *pb.ById) (*pb.Empty, error) {

	_, err := s.storage.Post().DeleteByUser(req)
	if err != nil {
		s.logger.Error("Failed Delete post", l.Error(err))
		return nil, nil
	}

	comment, err := s.client.CommentService().DeleteByUser(context.Background(), &p.ById{Id: req.Userid})

	if err != nil {
		s.logger.Error("Error delete comment from comment_service", l.Error(err))
		return nil, err
	}

	return (*pb.Empty)(comment), nil
}