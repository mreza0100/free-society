package domain

import (
	pb "freeSociety/proto/generated/post"
	"freeSociety/services/post/models"
	"freeSociety/utils"
	"freeSociety/utils/files/costume"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *service) GetPost(requestorId uint64, postIds []string) ([]*pb.Post, error) {
	if len(postIds) == 0 {
		return nil, nil
	}

	var (
		rawPosts []*models.Post
		posts    []*pb.Post

		ownerIds       []uint64
		users          map[uint64]*pb.User
		likeCount      map[string]uint64
		likedGroup     map[string]*emptypb.Empty
		followingGroup map[uint64]bool

		hasRequestorId bool
		errCh          chan error
	)
	{
		followingGroup = make(map[uint64]bool)
		hasRequestorId = requestorId != 0
		errCh = make(chan error)
	}

	{
		// get posts from internal db
		var err error
		rawPosts, err = s.repo.Read.GetPost(postIds)
		if err != nil {
			return nil, err
		}
	}
	{
		// collecting owner ids from posts
		// and make theme unique
		// there maight be several posts from one owner
		// I dont want to get same user
		// over and over again in the same request to user service
		notUniqueOwnerIds := make([]uint64, len(rawPosts))
		for idx, i := range rawPosts {
			notUniqueOwnerIds[idx] = i.OwnerId
		}
		ownerIds = utils.UniqueIds(notUniqueOwnerIds)
	}
	// $ concurrency =))) getting all data in the same time
	go func() {
		var err error
		users, err = s.publishers.GetUsers(ownerIds)
		errCh <- err
	}()
	go func() {
		var err error
		likeCount, err = s.publishers.GetCounts(postIds)
		errCh <- err
	}()
	go func() {
		if hasRequestorId {
			var err error
			likedGroup, err = s.publishers.IsLikedGroup(requestorId, postIds)
			errCh <- err
			return
		}
		errCh <- nil
	}()
	go func() {
		if hasRequestorId {
			var err error
			followingGroup, err = s.publishers.IsFollowingGroup(requestorId, ownerIds)
			errCh <- err
			return
		}
		errCh <- nil
	}()

	{
		for i := 0; i < 4; i++ {
			if err := <-errCh; err != nil {
				return nil, err
			}
		}
	}
	{
		posts = make([]*pb.Post, 0, len(rawPosts))
		for _, rawPost := range rawPosts {
			converted := &pb.Post{
				Title:       rawPost.Title,
				Body:        rawPost.Body,
				Id:          rawPost.ID,
				OwnerId:     rawPost.OwnerId,
				Likes:       likeCount[rawPost.ID],
				IsFollowing: followingGroup[rawPost.OwnerId],
				User:        users[rawPost.OwnerId],
			}

			{
				pictureUrls := make([]string, 0, len(rawPost.PicturesName))
				for i := 0; i < len(rawPost.PicturesName); i++ {
					if rawPost.PicturesName[i] != "" {
						pictureUrls = append(pictureUrls, costume.ExportPicture(rawPost.PicturesName[i]))
					}
				}
				converted.PictureUrls = pictureUrls
			}
			{
				_, found := likedGroup[converted.Id]
				converted.IsLiked = found
			}

			posts = append(posts, converted)
		}
	}

	return posts, nil
}
