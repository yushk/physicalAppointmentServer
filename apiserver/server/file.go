package server

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/apiserver/restapi/operations/monitor"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
	"github.com/yushk/sport_backend/pkg/pb"
)

// DownloadFile GET /file/download
func DownloadFile(params monitor.DownloadFileParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	steam, err := client.File().DownloadFile(ctx, &pb.DownloadFileRequest{
		Id: *params.ID,
	})

	if err != nil {
		return Error(err)
	}
	return middleware.ResponderFunc(func(rw http.ResponseWriter, producer runtime.Producer) {
		metaData, err := steam.Header()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		for k, v := range metaData {
			if len(v) > 0 {
				if k == "content-type" {
					// ignore application/grpc
					continue
				}
				if k == "file-content-type" {
					rw.Header().Set("Content-Type", v[0])
				} else {
					rw.Header().Set(k, v[0])
				}
			}
		}
		for {
			resp, err := steam.Recv()
			if err == io.EOF {
				logrus.Warnln("steam is eof")
				rw.WriteHeader(http.StatusOK)
				break
			}
			if err != nil {
				logrus.Errorln("steam error", err)
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			_, err = rw.Write(resp.File)
			if err != nil {
				logrus.Errorln("write error", err)
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	})
}

// GetClubFileZip
func GetClubFileZip(params monitor.GetClubFileZipParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	reply, err := client.File().GetClubFileZip(ctx, &pb.GetClubFileZipRequest{
		Club: *params.Club,
	})

	if err != nil {
		return Error(err)
	}
	zipFileName := *params.Club + ".zip"
	// 创建zip文件
	zipFile, err := createZipFile(zipFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer zipFile.Close() // 创建zip.Writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 遍历文件列表，将每个文件写入zip文件中
	for _, file := range reply.Files {
		if err := writeFileToZip(zipWriter, file.Path+"/"+file.Name, file.File); err != nil {
			logrus.Errorln(err)
		}
	}
	return middleware.ResponderFunc(func(rw http.ResponseWriter, producer runtime.Producer) {
		// 读取zip文件
		zipBytes, err := ioutil.ReadFile(zipFileName)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Header().Set("Content-Type", "application/octet-stream")
		rw.Header().Set("Content-Disposition", "attachment;filename="+filepath.Base(zipFileName))
		if _, err := rw.Write(zipBytes); err != nil {
			http.Error(rw, "Failed to write data to response", http.StatusInternalServerError)
			return
		}
	})
}

// createZipFile 创建zip文件
func createZipFile(filename string) (*os.File, error) {
	zipFile, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return zipFile, nil
}

// writeFileToZip 将文件内容写入zip文件中
func writeFileToZip(zipFile *zip.Writer, filename string, content []byte) error {
	fileWriter, err := zipFile.Create(filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(fileWriter, ioutil.NopCloser(bytes.NewReader(content)))
	if err != nil {
		return err
	}
	return nil
}

// UploadFile POST /file/upload
func UploadFile(params monitor.UploadFileParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	buf := make([]byte, 1024*1024)
	stream, err := client.File().UploadFile(ctx)
	if err != nil {
		return Error(err)
	}

	rtfile := params.File.(*runtime.File)
	filename := rtfile.Header.Filename
	defer rtfile.Close()
	filePath := ""
	if params.Path != nil {
		filePath = *params.Path
	}

	total := 0
	for {
		n, err := rtfile.Read(buf)
		if err == io.EOF {
			logrus.Infoln("Upload file EOF", total)
			break
		}
		if err != nil {
			_, err = stream.CloseAndRecv()
			return Error(err)
		}
		request := &pb.UploadFileRequest{
			Name: filename,
			Path: filePath,
			File: buf[:n],
		}
		err = stream.Send(request)
		if err != nil {
			stream.CloseAndRecv()
			return Error(err)
		}
		total += n
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		return Error(err)
	}
	payload := reply.Id
	return monitor.NewUploadFileOK().WithPayload(payload)
}

func DeleteFile(params monitor.DeleteFileParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	_, err := client.File().DeleteFile(ctx, &pb.DeleteFileRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteFileOK()
}
