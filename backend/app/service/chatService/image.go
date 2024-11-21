package chatService

import (
	"angrymiao-ai/app/tools/aliyun"
	"angrymiao-ai/config"
	"fmt"
	"mime/multipart"
)

func (s *Service) AnswerImageDeal(image *multipart.FileHeader) (string, error) {
	ossConfig := aliyun.OSSConfig{
		Endpoint:        config.Conf.OSS.Endpoint,
		AccessKeyID:     config.Conf.OSS.AccessKeyID,
		AccessKeySecret: config.Conf.OSS.AccessKeySecret,
		BucketName:      config.Conf.OSS.Bucket,
	}

	// 创建上传器
	uploader, err := aliyun.NewUploader(ossConfig)
	if err != nil {
		return "", err
	}

	// 验证文件
	if err = uploader.Validator.ValidateFile(image); err != nil {
		return "", err
	}

	file, err := image.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 构建OSS存储路径
	objectKey := image.Filename

	// 上传到OSS
	if err = uploader.UploadContent(file, objectKey); err != nil {
		return "", err
	}

	// 构建文件访问URL
	fileURL := fmt.Sprintf("%s/%s", config.Conf.OSS.Host, objectKey)

	return fileURL, err
}
