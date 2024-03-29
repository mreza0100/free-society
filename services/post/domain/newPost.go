package domain

import (
	"fmt"
	"freeSociety/configs"
	pb "freeSociety/proto/generated/post"
	"freeSociety/utils"
	"freeSociety/utils/files"
	"freeSociety/utils/files/costume"
)

func (s *service) NewPost(title, body string, userId uint64, pictures []*pb.Picture) (string, error) {
	var (
		picturesNames = make([]string, len(pictures))
		postId        string
	)

	{
		if len(pictures) > configs.Max_picture_per_post {
			return "", fmt.Errorf("more then %v pictures", configs.Max_picture_per_post)
		}
		for i := 0; i < len(pictures); i++ {
			format := files.GetFileFormat(pictures[i].Name)
			id := utils.GenerateUuid()

			picturesNames[i] = id + format
		}
	}
	{
		var err error
		postId, err = s.repo.Write.NewPost(title, body, userId, picturesNames)
		if err != nil {
			return "", err
		}
	}
	{
		for i := 0; i < len(pictures); i++ {
			p := costume.GetFullPathPicture(picturesNames[i])

			err := files.CreateAndWriteFile(p, pictures[i].Content)
			if err != nil {
				return "", err
			}
		}
	}
	return postId, nil
}
