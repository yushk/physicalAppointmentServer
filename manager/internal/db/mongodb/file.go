package mongodb

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/pkg/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FileInfo 文件属性
type FileInfo struct {
	Name string `bson:"name"`
	Path string `bson:"path"`
	File []byte `bson:"file"`
}

func (m *mongoDB) UploadFile(ctx context.Context, filename, path string, file []byte) (string, error) {
	info := &FileInfo{
		Name: filename,
		Path: path,
		File: file,
	}
	result, err := m.CFile().InsertOne(ctx, info)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID)
	return id.Hex(), err
}

func (m *mongoDB) DownloadFile(ctx context.Context, id string) (string, []byte, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", []byte{}, err
	}
	info := &FileInfo{}
	err = m.CFile().FindOne(ctx, bson.M{"_id": objectID}).Decode(&info)
	return info.Name, info.File, err
}

func (m *mongoDB) GetClubFileZip(ctx context.Context, club string) ([]*pb.File, error) {
	files := []*pb.File{}
	cursor, err := m.CFile().Find(ctx, bson.M{"path": bson.M{"$regex": "^" + club + "/"}})
	if err != nil {
		logrus.Errorln(err)
	}
	for cursor.Next(context.Background()) {
		file := &pb.File{}
		if err := cursor.Decode(&file); err != nil {
			logrus.Errorln(err)
		}
		files = append(files, file)
	}
	return files, err
}

func (m *mongoDB) DeleteFile(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = m.CFile().DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	return nil
}
