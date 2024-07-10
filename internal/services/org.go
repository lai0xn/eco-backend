package services

import (
	"context"
	"fmt"

	"github.com/lai0xn/squid-tech/pkg/logger"
	"github.com/lai0xn/squid-tech/pkg/redis"
	"github.com/lai0xn/squid-tech/pkg/types"
	"github.com/lai0xn/squid-tech/prisma"
	"github.com/lai0xn/squid-tech/prisma/db"
)

type OrgService struct{}

func NewOrgService() *OrgService {
	return &OrgService{}
}

func (s *OrgService) GetOrg(id string) (*db.OrganizationModel, error) {
	ctx := context.Background()
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "get org",
		"params": id,
	}).Msg("DB Query")

	user, err := prisma.Client.Organization.FindUnique(
		db.Organization.ID.Equals(id),
	).With(db.Organization.Owner.Fetch(),db.Organization.Followers.Fetch()).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *OrgService) GetUserOrgs(id string) ([]db.OrganizationModel, error) {
	ctx := context.Background()
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "get user orgs",
		"params": id,
	}).Msg("DB Query")
	user, err := prisma.Client.Organization.FindMany(
		db.Organization.OwnerID.Equals(id),
	).With(db.Organization.Followers.Fetch()).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *OrgService) SearchOrg(name string) ([]db.OrganizationModel, error) {
	ctx := context.Background()
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "search org",
		"params": name,
	}).Msg("DB Query")
	users, err := prisma.Client.Organization.FindMany(
		db.Organization.Name.Contains(name),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *OrgService) UpdateOrg(id string, payload types.OrgPayload) (*db.OrganizationModel, error) {
	ctx := context.Background()
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "update org",
		"id":     id,
		"params": payload,
	}).Msg("DB Query")
	users, err := prisma.Client.Organization.FindUnique(
		db.Organization.ID.Equals(id),
	).Update(
		db.Organization.Name.Set(payload.Name),
		db.Organization.Description.Set(payload.Description),
			).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *OrgService) CreateOrg(id string, payload types.OrgPayload) (*db.OrganizationModel, error) {
	logger.LogInfo().Fields(map[string]interface{}{
		"query":   "create org",
		"ownerId": id,
		"params":  payload,
	}).Msg("DB Query")
	ctx := context.Background()
	users, err := prisma.Client.Organization.CreateOne(
		db.Organization.Name.Set(payload.Name),
		db.Organization.Description.Set(payload.Description),
		db.Organization.Image.Set("uplodas/profiles/default.jpg"),
		db.Organization.BgImage.Set("uplodas/bgs/default.jpg"),
		db.Organization.Owner.Link(db.User.ID.Equals(id)),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *OrgService) UpdateOrgImage(id string, path string) (string, error) {
	fmt.Println(id)
	ctx := context.Background()
	user, err := prisma.Client.Organization.FindUnique(
		db.Organization.ID.Equals(id),
	).Update(
		db.Organization.Image.Set(path),
	).Exec(ctx)
	if err != nil {
		return "", err
	}
	return user.Image, nil
}

func (s *OrgService) UpdateOrgBg(id string, path string) (string, error) {
	fmt.Println(id)
	ctx := context.Background()
	user, err := prisma.Client.Organization.FindUnique(
		db.Organization.ID.Equals(id),
	).Update(
		db.Organization.BgImage.Set(path),
	).Exec(ctx)
	if err != nil {
		return "", err
	}
	return user.Image, nil
}

func (s *OrgService) DeleteOrg(id string) (string, error) {
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "delete org",
		"params": id,
	}).Msg("DB Query")
	ctx := context.Background()
	deleted, err := prisma.Client.Organization.FindUnique(
		db.Organization.ID.Equals(id),
	).Delete().Exec(ctx)
	if err != nil {
		return "", nil
	}
	fmt.Println(deleted.ID)
	return deleted.ID, nil
}

func (s *OrgService) Follow(userID string,orgId string) (string, error) {
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "follow org",

	}).Msg("DB Query")
	ctx := context.Background()
	org, err := prisma.Client.Organization.FindUnique(
		db.Organization.ID.Equals(orgId),
	).Update(db.Organization.Followers.Link(db.User.ID.Equals(userID))).Exec(ctx)
	if err != nil {
		return "", nil
	}
  client := redis.GetClient()
  key := fmt.Sprintf("notifs:%s",org.OwnerID)
  client.Publish(context.Background(),key,fmt.Sprintf("Your org %s gained a new follower",org.Name))
	return org.ID, nil
  
}


func (s *OrgService) Unfollow(userID string,orgId string) (string, error) {
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "follow org",

	}).Msg("DB Query")
	ctx := context.Background()
	deleted, err := prisma.Client.Organization.FindUnique(
		db.Organization.ID.Equals(orgId),
	).Update(db.Organization.Followers.Unlink(db.User.ID.Equals(userID))).Exec(ctx)
	if err != nil {
		return "", nil
	}
	fmt.Println(deleted.ID)
	return deleted.ID, nil
}
