package mongodb

import (
	"context"
	"encoding/json"

	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/pkg/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongoDB) CreateMovePrescription(ctx context.Context, data *pb.MovePrescription) (*pb.MovePrescription, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	_, err := m.CMovePrescription().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetMovePrescription(ctx context.Context, id string) (data *pb.MovePrescription, err error) {
	data = &pb.MovePrescription{}
	filter := bson.M{"id": id}
	err = m.CMovePrescription().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) UpdateMovePrescription(ctx context.Context, id string, data *pb.MovePrescription) (*pb.MovePrescription, error) {
	newData := &pb.MovePrescription{}
	filter := bson.M{"id": id}
	data.Id = id
	update := bson.M{"$set": data}
	err := m.CMovePrescription().FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(newData)
	if err != nil {
		return newData, err
	}
	return newData, nil
}

func (m *mongoDB) DeleteMovePrescription(ctx context.Context, id string) (data *pb.MovePrescription, err error) {
	data = &pb.MovePrescription{}
	filter := bson.M{"id": id}
	err = m.CMovePrescription().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetMovePrescriptions(ctx context.Context, limit, skip int64, query string) (count int64, datas []*pb.MovePrescription, err error) {
	datas = []*pb.MovePrescription{}

	filter := bson.M{}
	if query != "" {
		err = json.Unmarshal([]byte(query), &filter)
		if err != nil {
			return
		}
	}

	findOption := &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}
	cursor, err := m.CMovePrescription().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.MovePrescription{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("MovePrescription Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CMovePrescription().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("MovePrescription Count Documents Error")
		count = int64(0)
	}
	return
}
