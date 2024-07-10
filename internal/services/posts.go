package services

import (
	"context"
	"fmt"

	"github.com/lai0xn/squid-tech/pkg/logger"
	"github.com/lai0xn/squid-tech/pkg/types"
	"github.com/lai0xn/squid-tech/prisma"
	"github.com/lai0xn/squid-tech/prisma/db"
)

type PostService struct{}

func NewPostService() *PostService {
	return &PostService{}
}

func (s *PostService) GetPost(id string) (*db.PostModel, error) {
	ctx := context.Background()
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "get event",
		"params": id,
	}).Msg("DB Query")

	result, err := prisma.Client.Post.FindUnique(
		db.Post.ID.Equals(id),
	).With(db.Post.Author.Fetch(), db.Post.Comments.Fetch()).Exec(ctx)
	result.Author().Password = ""
	result.Author().EventsIds = nil
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *PostService) UploadImage(id string, path string) (*db.PostModel, error) {
	ctx := context.Background()
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "upload image post",
		"params": id,
	}).Msg("DB Query")

	result, err := prisma.Client.Post.FindUnique(
		db.Post.ID.Equals(id),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *PostService) GetPosts(page int) ([]db.PostModel, error) {
	ctx := context.Background()
	limit := 10
	result, err := prisma.Client.Post.FindMany().Skip((page-1)*limit).Take(limit).With(db.Post.Author.Fetch(), db.Post.Comments.Fetch()).Exec(ctx)
	if err != nil {
		return nil, err
	}
	for _, r := range result {
		r.Author().Password = ""
		r.Author().EventsIds = nil

	}

	return result, nil
}

func (s *PostService) GetComment(id string) (*db.PostCommentModel, error) {
	ctx := context.Background()
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "get event",
		"params": id,
	}).Msg("DB Query")

	result, err := prisma.Client.PostComment.FindUnique(
		db.PostComment.ID.Equals(id),
	).With(db.PostComment.Author.Fetch()).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *PostService) SearchPost(name string) ([]db.PostModel, error) {
	ctx := context.Background()
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "search event",
		"params": name,
	}).Msg("DB Query")
	result, err := prisma.Client.Post.FindMany(
		db.Post.Or(
			db.Post.Content.Contains(name),
			db.Post.Description.Contains(name),
		),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *PostService) UpdatePost(id string, payload types.PostPayload) (*db.PostModel, error) {
	ctx := context.Background()
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "update org",
		"id":     id,
		"params": payload,
	}).Msg("DB Query")
	results, err := prisma.Client.Post.FindUnique(
		db.Post.ID.Equals(id),
	).Update(
		db.Post.Content.Set(payload.Content),
		db.Post.Description.Set(payload.Description),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *PostService) CreatePost(userId string, content string, description string) (*db.PostModel, error) {

	ctx := context.Background()
	results, err := prisma.Client.Post.CreateOne(
		db.Post.Content.Set(content),
		db.Post.Description.Set(description),
		// TODO: Add image
		db.Post.Author.Link(db.User.ID.Equals(userId)),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *PostService) CreateComment(userId string, postID, content string) (*db.PostCommentModel, error) {

	ctx := context.Background()
	results, err := prisma.Client.PostComment.CreateOne(
		db.PostComment.Content.Set(content),
		db.PostComment.Author.Link(db.User.ID.Equals(userId)),
		db.PostComment.Post.Link(db.Post.ID.Equals(postID)),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *PostService) DeleteComment(id string) (*db.PostCommentModel, error) {
	ctx := context.Background()
	results, err := prisma.Client.PostComment.FindUnique(
		db.PostComment.ID.Equals(id),
	).Delete().Exec(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *PostService) CommentPost(postid string, userid string, content string) (*db.PostCommentModel, error) {
	logger.LogInfo().Fields(map[string]interface{}{
		"query": "create event",
	}).Msg("DB Query")
	ctx := context.Background()
	results, err := prisma.Client.PostComment.CreateOne(
		db.PostComment.Content.Set(content),
		db.PostComment.Author.Link(db.User.ID.Equals(userid)),
		db.PostComment.Post.Link(db.Post.ID.Equals(postid)),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *PostService) DeletePost(id string) (string, error) {
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "delete org",
		"params": id,
	}).Msg("DB Query")
	ctx := context.Background()
	deleted, err := prisma.Client.Post.FindUnique(
		db.Post.ID.Equals(id),
	).Delete().Exec(ctx)
	if err != nil {
		return "", nil
	}
	fmt.Println(deleted.ID)
	return deleted.ID, nil
}
