package domain

import (
	"fmt"
	"freeSociety/configs"
	pb "freeSociety/proto/generated/post"
	"freeSociety/services/post/instances"
	"freeSociety/services/post/models"
	"freeSociety/services/post/repository"
	"freeSociety/utils"
	"freeSociety/utils/files"
	"freeSociety/utils/hash"
	"path"
	"strings"

	"github.com/mreza0100/golog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NewOpts struct {
	Lgr        *golog.Core
	Publishers instances.Publishers
}

func New(opts *NewOpts) instances.Sevice {
	return &service{
		lgr:        opts.Lgr.With("In domain->"),
		repo:       repository.NewRepo(opts.Lgr),
		publishers: opts.Publishers,
	}
}

type service struct {
	lgr        *golog.Core
	repo       *instances.Repository
	publishers instances.Publishers
}

func (s *service) NewPost(title, body string, userId uint64, pictures []*pb.Picture) (uint64, error) {
	var (
		picturesPath = make([]string, len(pictures))
		postId       uint64
	)

	{
		if len(pictures) > configs.Max_picture_per_post {
			return 0, fmt.Errorf("more then %v pictures", configs.Max_picture_per_post)
		}
		for i := 0; i < len(pictures); i++ {
			format := files.GetFileFormat(pictures[i].Name)
			hash := hash.Hash512(pictures[i].Content)

			picturesPath[i] = hash + format
		}
	}
	{
		var err error
		postId, err = s.repo.Write.NewPost(title, body, userId, picturesPath)
		if err != nil {
			return 0, err
		}
	}
	{
		for i := 0; i < len(pictures); i++ {
			p := path.Join(configs.PicturesPath, picturesPath[i])

			if files.FileExist(p) {
				continue
			}
			err := files.CreateAndWriteFile(p, pictures[i].Content)
			if err != nil {
				return 0, err
			}
		}
	}
	return postId, nil
}

func (s *service) DeletePost(postId, userId uint64) error {
	rawPicturesPaths, err := s.repo.Write.DeletePost(postId, userId)
	if err != nil {
		return err
	}

	picturesPaths := strings.Split(rawPicturesPaths, configs.DB_picture_sep)
	for i := 0; i < len(picturesPaths); i++ {
		exist, err := s.repo.Read.IsPictureExist(picturesPaths[i])
		s.lgr.InfoLog(exist, "---", err)
		if err != nil {
			return err
		}
		if exist {
			continue
		}

		if files.DeleteFile(path.Join(configs.PicturesPath, picturesPaths[i])) != nil {
			return err
		}
	}

	return nil
}

func (s *service) DeleteUserPosts(userId uint64) error {
	return s.repo.Write.DeleteUserPosts(userId)
}

func (s *service) IsPostsExists(postIds []uint64) ([]uint64, error) {
	return s.repo.Read.IsExists(postIds)
}

func (s *service) GetPost(requestorId uint64, postIds []uint64) ([]*pb.Post, error) {
	if len(postIds) == 0 {
		return nil, nil
	}

	var (
		rawPosts []*models.Post
		posts    []*pb.Post

		ownerIds       []uint64
		users          map[uint64]*pb.User
		likeCount      map[uint64]uint64
		likedGroup     map[uint64]*emptypb.Empty
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
			id := rawPost.ID
			ownerId := rawPost.OwnerId

			converted := &pb.Post{
				Title:       rawPost.Title,
				Body:        rawPost.Body,
				Id:          id,
				OwnerId:     ownerId,
				Likes:       likeCount[id],
				IsFollowing: followingGroup[id],
				User:        users[ownerId],
				PictureUrls: strings.Split(rawPost.PicturesPath, configs.DB_picture_sep),
			}

			_, found := likedGroup[converted.Id]
			converted.IsLiked = found

			posts = append(posts, converted)
		}
	}

	return posts, nil
}
