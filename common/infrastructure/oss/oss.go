package oss

import (
	"fmt"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var objStorage *Oss

type OssService interface {
	UploadPicture(des string, picture io.Reader) (url string, err error)
}

type Oss struct {
	Client *oss.Client
	Bucket *oss.Bucket

	BucketName        string
	EndPoint          string
	RootDir           string
	PictureFolderName string
}

func Init(cfg *Config) error {
	// new client
	cli, err := oss.New(cfg.Endpoint, cfg.Id, cfg.Secret,
		oss.EnableCRC(true),
		oss.Timeout(cfg.ConnTimeout, cfg.ReadWriteTimeout),
	)
	if err != nil {
		return err
	}

	// new bucket
	bucket, err := cli.Bucket(cfg.Bucket)
	if err != nil {
		return err
	}

	objStorage = &Oss{
		Client:            cli,
		Bucket:            bucket,
		BucketName:        cfg.Bucket,
		EndPoint:          cfg.Endpoint,
		RootDir:           cfg.RootDir,
		PictureFolderName: cfg.PictureFolderName,
	}

	return nil
}

func Client() OssService {
	return objStorage
}

func (o *Oss) UploadPicture(position string, file io.Reader) (string, error) {
	return o.uploadFile(o.genObjectName(position), file)
}

func (o *Oss) uploadFile(objectName string, file io.Reader) (string, error) {
	return o.genVisitUrl(objectName), o.Bucket.PutObject(objectName, file)
}

// genObjectName generate full object name with position("username/filename")
func (o *Oss) genObjectName(position string) string {
	return fmt.Sprintf("%s/%s", o.getPictureRootDir(), position)
}

func (o *Oss) getPictureRootDir() string {
	return fmt.Sprintf("%s/%s", o.RootDir, o.PictureFolderName)
}

func (o *Oss) genVisitUrl(objectName string) string {
	return fmt.Sprintf("https://%s.%s/%s", o.BucketName, o.EndPoint, objectName)
}
