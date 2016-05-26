package data

import (
	"fmt"

	"github.com/sogko/gosg-graphql-go-demo/server/microblog"
)

var Users = []*microblog.User{}
var Posts = []*microblog.Post{}

var UserFollowersMap = map[int][]int{}
var UserFriendsMap = map[int][]int{}
var PostLikesMap = map[int][]int{}

func init() {
	UserFollowersMap = map[int][]int{
		1:  {2, 3, 4, 5, 6, 7, 8, 9, 10},
		2:  {1, 5, 4, 10},
		3:  {1, 2, 6, 8, 10},
		4:  {1, 3, 6, 8, 10},
		5:  {3, 2, 6, 7, 10},
		6:  {4, 2, 5, 8, 10},
		7:  {4, 2, 6, 1, 10},
		8:  {5, 2, 6, 10},
		9:  {3, 2, 6, 8, 10},
		10: {1, 2, 6, 8},
	}
	UserFriendsMap = map[int][]int{
		1:  {2, 3, 4, 5, 6, 7, 8, 9, 10},
		2:  {1, 5, 4, 10},
		3:  {1, 2, 6, 8, 10},
		4:  {1, 3, 6, 8, 10},
		5:  {3, 2, 6, 7, 10},
		6:  {4, 2, 5, 8, 10},
		7:  {4, 2, 6, 1, 10},
		8:  {5, 2, 6, 10},
		9:  {3, 2, 6, 8, 10},
		10: {1, 2, 6, 8},
	}
	PostLikesMap = map[int][]int{
		1:  {2, 3, 7, 8, 9, 10},
		2:  {1, 2, 7},
		3:  {2, 6, 3},
		4:  {2, 9, 4},
		5:  {1, 8, 3},
		6:  {3, 2, 9},
		7:  {2, 4},
		9:  {2, 9, 10},
		10: {1, 3, 6},
		11: {1, 2, 3},
		12: {1, 2, 8},
		13: {5, 7, 4},
		14: {9, 8, 3},
		15: {3, 2, 7},
	}
	user1 := &microblog.User{
		ID:             1,
		Name:           "Hafiz Ismail",
		Handle:         "sogko",
		AvatarURL:      "https://pbs.twimg.com/profile_images/2952455687/c6bb8a5ca72135f4a225d7e5921db6fc_bigger.jpeg",
		TotalFollowers: 385,
		TotalFriends:   119,
	}
	user2 := &microblog.User{
		ID:             2,
		Name:           "B L A I S E",
		Handle:         "BlaiseGabriella",
		AvatarURL:      "https://pbs.twimg.com/profile_images/421478457218850816/-Xq4aGI3_bigger.jpeg",
		TotalFollowers: 190,
		TotalFriends:   80,
	}
	user3 := &microblog.User{
		ID:             3,
		Name:           "EFF",
		Handle:         "EFF",
		AvatarURL:      "https://pbs.twimg.com/profile_images/712790912007479296/qBtsMbMr_bigger.jpg",
		TotalFollowers: 24,
		TotalFriends:   12,
	}
	user4 := &microblog.User{
		ID:             4,
		Name:           "EFF",
		Handle:         "EFF",
		AvatarURL:      "https://pbs.twimg.com/profile_images/712790912007479296/qBtsMbMr_bigger.jpg",
		TotalFollowers: 24,
		TotalFriends:   12,
	}
	user5 := &microblog.User{
		ID:             5,
		Name:           "Eric Meyer",
		Handle:         "meyerweb",
		AvatarURL:      "https://pbs.twimg.com/profile_images/671335262459334657/O6oMcV8M_bigger.png",
		TotalFollowers: 9623,
		TotalFriends:   428,
	}
	user6 := &microblog.User{
		ID:             6,
		Name:           "Michel",
		Handle:         "neumino",
		AvatarURL:      "https://pbs.twimg.com/profile_images/449431509699543041/sqRIwp2c_bigger.png",
		TotalFollowers: 9623,
		TotalFriends:   428,
	}
	user7 := &microblog.User{
		ID:             7,
		Name:           "Mick Pollard ",
		Handle:         "aussielunix",
		AvatarURL:      "https://pbs.twimg.com/profile_images/432269414042320896/_W320BfK_bigger.jpeg",
		TotalFollowers: 9623,
		TotalFriends:   428,
	}

	user8 := &microblog.User{
		ID:             8,
		Name:           "Damian Gryski",
		Handle:         "dgryski",
		AvatarURL:      "https://pbs.twimg.com/profile_images/59859794/dmg-simpsons-head_bigger.png",
		TotalFollowers: 9623,
		TotalFriends:   428,
	}

	user9 := &microblog.User{
		ID:             9,
		Name:           "Neil Patrick Harris",
		Handle:         "ActuallyNPH",
		AvatarURL:      "https://pbs.twimg.com/profile_images/378800000601987374/898bea5682c583597c4f471bdbb87a4b_bigger.jpeg",
		TotalFollowers: 9623,
		TotalFriends:   428,
	}
	user10 := &microblog.User{
		ID:             10,
		Name:           "Amanda Katz‏",
		Handle:         "katzish",
		AvatarURL:      "https://pbs.twimg.com/profile_images/732875606/A_laughing_b_w_square_bigger.jpg",
		TotalFollowers: 9623,
		TotalFriends:   428,
	}

	post1 := &microblog.Post{
		ID:         1,
		Content:    "Hello World",
		AuthorID:   1,
		TotalLikes: 1023,
	}
	post2 := &microblog.Post{
		ID:         2,
		Content:    "Oh. You need a little dummy text for your mockup? How quaint.",
		AuthorID:   2,
		TotalLikes: 123,
	}
	post3 := &microblog.Post{
		ID:         3,
		Content:    "I bet you’re still using Bootstrap too...",
		AuthorID:   3,
		TotalLikes: 12,
	}
	post4 := &microblog.Post{
		ID:         4,
		Content:    "Irony asymmetrical cold-pressed fap farm-to-table distillery +1 bicycle rights shabby chic cronut VHS venmo mustache.",
		AuthorID:   10,
		TotalLikes: 23,
	}
	post5 := &microblog.Post{
		ID:         5,
		Content:    "Trust fund 3 wolf moon meditation, austin ramps asymmetrical keytar brunch. ",
		AuthorID:   1,
		TotalLikes: 234,
	}
	post6 := &microblog.Post{
		ID:         6,
		Content:    "Beard man braid bushwick, neutra affogato plaid brooklyn literally chambray offal",
		AuthorID:   8,
		TotalLikes: 234,
	}
	post7 := &microblog.Post{
		ID:         7,
		Content:    "Vice brunch polaroid, banh mi paleo cray kinfolk letterpress sriracha echo park bushwick.",
		AuthorID:   4,
		TotalLikes: 234,
	}
	post8 := &microblog.Post{
		ID:         8,
		Content:    "Skateboard meditation salvia, kogi twee aesthetic williamsburg cred sriracha bespoke ugh PBR&B pug bicycle rights master cleanse.",
		AuthorID:   5,
		TotalLikes: 234,
	}
	post9 := &microblog.Post{
		ID:         9,
		Content:    "Meh aesthetic echo park literally hella, fap mlkshk shoreditch wolf viral kogi tacos +1 kombucha.",
		AuthorID:   6,
		TotalLikes: 234,
	}
	post10 := &microblog.Post{
		ID:         10,
		Content:    "Marfa single-origin coffee vice PBR&B organic.",
		AuthorID:   7,
		TotalLikes: 234,
	}
	post11 := &microblog.Post{
		ID:         11,
		Content:    "Gluten-free disrupt biodiesel trust fund occupy.",
		AuthorID:   5,
		TotalLikes: 234,
	}
	post12 := &microblog.Post{
		ID:         12,
		Content:    "Tofu narwhal cronut, tattooed asymmetrical fap butcher biodiesel portland chartreuse.",
		AuthorID:   10,
		TotalLikes: 234,
	}
	post13 := &microblog.Post{
		ID:         13,
		Content:    "Waistcoat tattooed mumblecore you probably haven't heard of them, kickstarter twee chillwave chia cornhole cray keffiyeh pork belly mustache.",
		AuthorID:   1,
		TotalLikes: 234,
	}
	post14 := &microblog.Post{
		ID:         14,
		Content:    "Street art 8-bit pork belly, everyday carry listicle gentrify swag gluten-free fap readymade franzen meditation helvetica.",
		AuthorID:   8,
		TotalLikes: 234,
	}
	post15 := &microblog.Post{
		ID:         15,
		Content:    "Trust fund retro hammock lomo chillwave quinoa",
		AuthorID:   9,
		TotalLikes: 234,
	}

	Posts = []*microblog.Post{
		post1,
		post2,
		post3,
		post4,
		post5,
		post6,
		post7,
		post8,
		post9,
		post10,
		post11,
		post12,
		post13,
		post14,
		post15,
	}

	Users = []*microblog.User{
		user1,
		user2,
		user3,
		user4,
		user5,
		user6,
		user7,
		user8,
		user9,
		user10,
	}
}

// GetUser retrieves User for given ID
func GetUser(id int) (*microblog.User, error) {
	for _, user := range Users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("User (id: %v) not found", id)
}

// GetFriendsForUser retrieves Friends for User of given ID
func GetFriendsForUser(userID int) ([]*microblog.User, error) {
	friends := []*microblog.User{}
	for _, userID := range UserFriendsMap[userID] {
		user, _ := GetUser(userID)
		friends = append(friends, user)
	}
	return friends, nil
}

// GetFollowersForUser retrieves Followers for User of given ID
func GetFollowersForUser(userID int) ([]*microblog.User, error) {
	followers := []*microblog.User{}
	for _, userID := range UserFollowersMap[userID] {
		user, _ := GetUser(userID)
		followers = append(followers, user)
	}
	return followers, nil
}

// GetLikedByForPost retrieves Users who liked Post of given ID
func GetLikedByForPost(postID int) ([]*microblog.User, error) {
	likedBy := []*microblog.User{}
	for _, userID := range PostLikesMap[postID] {
		user, _ := GetUser(userID)
		likedBy = append(likedBy, user)
	}
	return likedBy, nil
}

// GetPostsForUser retrieves Posts for User of given ID
func GetPostsForUser(userID int, limit int) ([]*microblog.Post, error) {
	posts := []*microblog.Post{}
	for _, post := range Posts {
		if post.AuthorID == userID {
			posts = append(posts, post)
		}
	}
	if limit > len(posts) {
		limit = len(posts)
	}
	return posts[0:limit], nil
}

// GetPost retrieves Post for given ID
func GetPost(id int) (*microblog.Post, error) {
	for _, post := range Posts {
		if post.ID == id {
			return post, nil
		}
	}
	return nil, fmt.Errorf("Post (id: %v) not found", id)
}

// GetPosts retrieves list of Posts
func GetPosts(limit int) ([]*microblog.Post, error) {
	if limit > len(Posts) {
		limit = len(Posts)
	}
	return Posts[0:limit], nil
}
