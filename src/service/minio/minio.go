package minio

import (
	"bytes"
	"fmt"
	"log"

	"github.com/minio/minio-go"
)

type dao struct {
	client *minio.Client
}

func NewService(addr, accessKey, secretKey string, useSSL bool) MinioService {
	endpoint := addr
	accessKeyID := accessKey
	secretAccessKey := secretKey
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &dao{client: minioClient}
}

func (d *dao) GetFile(bucket, path string) ([]byte, error) {
	obj, err := d.client.GetObject(bucket, path, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(obj)
	return buf.Bytes(), nil
}

func (d *dao) PutFile(bucket, path string, data []byte) error {
	file := bytes.NewReader(data)
	info, err := d.client.PutObject(bucket, path, file, int64(len(data)), minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	log.Println("Put file success", info)
	return nil
}

func (d *dao) CreateBucket(name string) error {
	err := d.client.MakeBucket(name, "us-east-1")
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Create bucket success", name)
	return nil
}

func (d *dao) BucketExists(name string) bool {
	found, err := d.client.BucketExists(name)
	if err != nil {
		log.Println(err)
		return false
	}
	return found
}

func (d *dao) RemoveFile(bucket, path string) error {
	return d.client.RemoveObject(bucket, path)
}
