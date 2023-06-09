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

func (m *mongoDB) CreateMeasureDetail(ctx context.Context, data *pb.MeasureDetail) (*pb.MeasureDetail, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	_, err := m.CMeasureDetail().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetMeasureDetail(ctx context.Context, id string) (data *pb.MeasureDetail, err error) {
	data = &pb.MeasureDetail{}
	filter := bson.M{"id": id}
	err = m.CMeasureDetail().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) UpdateMeasureDetail(ctx context.Context, id string, data *pb.MeasureDetail) (*pb.MeasureDetail, error) {
	newData := &pb.MeasureDetail{}
	filter := bson.M{"id": id}
	data.Id = id
	update := bson.M{"$set": data}
	err := m.CMeasureDetail().FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(newData)
	if err != nil {
		return newData, err
	}
	return newData, nil
}

func (m *mongoDB) DeleteMeasureDetail(ctx context.Context, id string) (data *pb.MeasureDetail, err error) {
	data = &pb.MeasureDetail{}
	filter := bson.M{"id": id}
	err = m.CMeasureDetail().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetMeasureDetails(ctx context.Context, limit, skip int64, query string) (count int64, datas []*pb.MeasureDetail, err error) {
	datas = []*pb.MeasureDetail{}

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
	cursor, err := m.CMeasureDetail().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.MeasureDetail{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("MeasureDetail Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CMeasureDetail().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("MeasureDetail Count Documents Error")
		count = int64(0)
	}
	return
}
