package server

import (
	"context"
	"io"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// UploadFile 上传文件
func (s *Server) UploadFile(srv pb.FileManager_UploadFileServer) error {
	reply := &pb.UploadFileReply{}
	file := []byte{}
	filename := ""
	filePath := ""
	for {
		request, err := srv.Recv()
		if err == io.EOF {
			id, err := s.db.UploadFile(srv.Context(), filename, filePath, file)
			if err != nil {
				return err
			}
			reply.Id = id
			return srv.SendAndClose(reply)
		}
		if err != nil {
			logrus.Errorln("srv revc error:", err)
			return err
		}
		if filename == "" {
			filename = request.Name
		}
		if filePath == "" {
			filePath = request.Path
		}
		file = append(file, request.File...)
	}
}

// DownloadFile 下载文件
func (s *Server) DownloadFile(req *pb.DownloadFileRequest, srv pb.FileManager_DownloadFileServer) error {

	filename, file, err := s.db.DownloadFile(srv.Context(), req.Id)
	if err != nil {
		return err
	}
	fileSuffix := strings.ToLower(path.Ext(filename))
	contentType := "application/octet-stream"
	switch fileSuffix {
	case ".svg":
		contentType = "image/svg+xml;charset=UTF-8"
	case ".png":
		contentType = "image/png"
	case ".bmp":
		contentType = "image/bmp"
	case ".jpeg":
		contentType = "image/jpeg"
	case ".jpg":
		contentType = "image/jpeg"
	}
	srv.SetHeader(metadata.New(map[string]string{
		"Content-Disposition": "attachment; filename=" + url.QueryEscape(filename),
		"Accept-Ranges":       "bytes",
		"File-Content-Type":   contentType,
		"Content-Length":      strconv.Itoa(len(file)),
	}))
	err = srv.Send(&pb.DownloadFileReply{
		File: file,
	})
	return err
}

// GetClubFileZip ...
func (s *Server) GetClubFileZip(ctx context.Context, req *pb.GetClubFileZipRequest) (reply *pb.GetClubFileZipReply, err error) {
	files, err := s.db.GetClubFileZip(ctx, req.Club)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetClubFileZipReply{
		Files: files,
	}
	return
}

// DeleteFile ...
func (s *Server) DeleteFile(ctx context.Context, req *pb.DeleteFileRequest) (reply *pb.DeleteFileReply, err error) {
	err = s.db.DeleteFile(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.DeleteFileReply{}
	return
}
