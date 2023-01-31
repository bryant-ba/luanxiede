package main

import (
	"github.com/minio/minio-go/v6"
	"log"
)

func main() {
	endpoint := "10.3.1.24:9000"
	accessKeyID := "admin123"
	secretAccessKey := "admin123"
	useSSL := false //注意没有安装证书的填false

	// 初使化minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("初始化完成！")
	// 链接一个叫df1024的存储桶。
	//下面注释部分是创建一个叫df1024的存储桶。
	bucketName := "df1024"
	location := "cn-north-1"
	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		log.Println("创建bucket失败！")
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Println("打印失败！")
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)
	//
	//// 上传一个zip文件。
	//objectName := "main.zip"
	//filePath := "C:\\Users\\11\\Desktop\\GO\\project\\src\\github.com\\DF1024\\Minio对象存储\\文件上传\\main.zip"
	//contentType := "application/zip"
	//
	//// 使用FPutObject上传一个zip文件。
	//n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}
