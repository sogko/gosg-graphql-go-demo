package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/sogko/gosg-graphql-go-demo/server/data"
	"github.com/sogko/gosg-graphql-go-demo/server/microblog"
)

var Root graphql.Schema
var PostType *graphql.Object
var UserType *graphql.Object

func init() {

	UserType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "User object",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id": &graphql.Field{
					Type:        graphql.Int,
					Description: "User ID",
				},
				"name": &graphql.Field{
					Type:        graphql.String,
					Description: "User's full name",
				},
				"handle": &graphql.Field{
					Type:        graphql.String,
					Description: "User's handle",
				},
				"avatar_url": &graphql.Field{
					Type:        graphql.String,
					Description: "Relative URL to user's avatar",
				},
				"total_followers": &graphql.Field{
					Type:        graphql.Int,
					Description: "User's total followers count",
				},
				"total_friends": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "User's total friends count",
				},
				"followers": &graphql.Field{
					Type:        graphql.NewList(UserType),
					Description: "User's list of followers",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if user, ok := p.Source.(*microblog.User); ok {
							// retrieve list of friends for given user id
							return data.GetFollowersForUser(user.ID)
						}
						return nil, nil
					},
				},
				"friends": &graphql.Field{
					Type:        graphql.NewList(UserType),
					Description: "User's list of friends",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if user, ok := p.Source.(*microblog.User); ok {
							// retrieve list of friends for given user id
							return data.GetFriendsForUser(user.ID)
						}
						return nil, nil
					},
				},
				"posts": &graphql.Field{
					Type:        graphql.NewList(PostType),
					Description: "User's list of posts",
					Args: graphql.FieldConfigArgument{
						"limit": &graphql.ArgumentConfig{
							Type:         graphql.Int,
							DefaultValue: 10,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// get value of `limit` arg
						limit := 10
						if v, ok := p.Args["limit"].(int); ok {
							limit = v
						}
						if user, ok := p.Source.(*microblog.User); ok {
							// retrieve list of posts for given user id
							return data.GetPostsForUser(user.ID, limit)
						}
						return nil, nil
					},
				},
			}
		}),
	})

	PostType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Post",
		Description: "Post object",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id": &graphql.Field{
					Type:        graphql.Int,
					Description: "Post ID",
				},
				"content": &graphql.Field{
					Type:        graphql.String,
					Description: "Text content of the post",
				},
				"author": &graphql.Field{
					Type:        UserType,
					Description: "Author of the post",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if post, ok := p.Source.(*microblog.Post); ok {
							// retrieve list of user for given id
							return data.GetUser(post.AuthorID)
						}
						return nil, nil
					},
				},
				"total_likes": &graphql.Field{
					Type:        graphql.Int,
					Description: "Number of users who liked the post",
				},
				"liked_by": &graphql.Field{
					Type:        graphql.NewList(UserType),
					Description: "List of users who liked the post",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if post, ok := p.Source.(*microblog.Post); ok {
							// retrieve list of users that like the post for given post id
							return data.GetLikedByForPost(post.ID)
						}
						return nil, nil
					},
				},
			}
		}),
	})

	var err error
	Root, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "RootQuery",
			Description: "Root for all query objects on our GraphQL server",
			Fields: graphql.Fields{
				"posts": &graphql.Field{
					Type:        graphql.NewList(PostType),
					Description: "Get list of posts",
					Args: graphql.FieldConfigArgument{
						"limit": &graphql.ArgumentConfig{
							Type:         graphql.Int,
							DefaultValue: 10,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// get value of `limit` arg
						limit := 10
						if v, ok := p.Args["limit"].(int); ok {
							limit = v
						}
						return data.GetPosts(limit)
					},
				},
				"post": &graphql.Field{
					Type:        PostType,
					Description: "Get a single post",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// get value of `id` arg
						id := 0
						if v, ok := p.Args["id"].(int); ok {
							id = v
						}
						return data.GetPost(id)
					},
				},
				"user": &graphql.Field{
					Type:        UserType,
					Description: "Get a single user",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// get value of `id` arg
						id := 0
						if v, ok := p.Args["id"].(int); ok {
							id = v
						}
						// retrieve user for given id from data source
						return data.GetUser(id)
					},
				},
			},
		}),
	})
	if err != nil {
		panic(err)
	}
}
