package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// OSSConfig 存储 OSS 配置信息
type OSSConfig struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
}

// Uploader OSS上传器结构体
type Uploader struct {
	client    *oss.Client
	bucket    *oss.Bucket
	config    OSSConfig
	Validator *FileValidator
}

// NewUploader 创建新的上传器实例
func NewUploader(config OSSConfig) (*Uploader, error) {
	client, err := oss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		return nil, fmt.Errorf("创建 OSS 客户端失败: %v", err)
	}

	bucket, err := client.Bucket(config.BucketName)
	if err != nil {
		return nil, fmt.Errorf("获取 Bucket 失败: %v", err)
	}

	return &Uploader{
		client:    client,
		bucket:    bucket,
		config:    config,
		Validator: NewFileValidator(),
	}, nil
}

// UploadFile 上传本地文件到 OSS
func (u *Uploader) UploadFile(localFilePath, ossPath string) error {
	// 确保本地文件存在
	if _, err := os.Stat(localFilePath); os.IsNotExist(err) {
		return fmt.Errorf("本地文件不存在: %s", localFilePath)
	}

	// 如果没有指定 OSS 路径，使用文件名
	if ossPath == "" {
		ossPath = filepath.Base(localFilePath)
	}

	// 执行上传
	err := u.bucket.PutObjectFromFile(ossPath, localFilePath)
	if err != nil {
		return fmt.Errorf("上传文件失败: %v", err)
	}

	return nil
}

// UploadContent 上传内容到 OSS
func (u *Uploader) UploadContent(reader io.Reader, ossPath string) error {
	if ossPath == "" {
		return fmt.Errorf("OSS 路径不能为空")
	}

	err := u.bucket.PutObject(ossPath, reader)
	if err != nil {
		return fmt.Errorf("上传内容失败: %v", err)
	}

	return nil
}

// GenerateSignedURL 生成文件的签名URL
func (u *Uploader) GenerateSignedURL(ossPath string, expireSeconds int64) (string, error) {
	if ossPath == "" {
		return "", fmt.Errorf("OSS 路径不能为空")
	}

	url, err := u.bucket.SignURL(ossPath, oss.HTTPGet, expireSeconds)
	if err != nil {
		return "", fmt.Errorf("生成签名URL失败: %v", err)
	}

	return url, nil
}

// DeleteFile 删除 OSS 上的文件
func (u *Uploader) DeleteFile(ossPath string) error {
	if ossPath == "" {
		return fmt.Errorf("OSS 路径不能为空")
	}

	err := u.bucket.DeleteObject(ossPath)
	if err != nil {
		return fmt.Errorf("删除文件失败: %v", err)
	}

	return nil
}

// FileValidator 文件验证器结构体
type FileValidator struct {
	MaxSize      int64               // 最大文件大小（字节）
	AllowedTypes map[string]struct{} // 允许的MIME类型
	AllowedExts  map[string]struct{} // 允许的文件扩展名
}

// NewFileValidator 创建新的文件验证器
func NewFileValidator() *FileValidator {
	return &FileValidator{
		MaxSize: 10 << 20, // 默认10MB
		AllowedTypes: map[string]struct{}{
			"image/jpeg": {},
			"image/png":  {},
			"image/gif":  {},
		},
		AllowedExts: map[string]struct{}{
			".jpg":  {},
			".jpeg": {},
			".png":  {},
			".gif":  {},
		},
	}
}

// ValidateFile 验证文件
func (v *FileValidator) ValidateFile(file *multipart.FileHeader) error {
	// 检查文件大小
	if file.Size > v.MaxSize {
		return fmt.Errorf("文件大小超过限制：%d bytes", v.MaxSize)
	}

	// 检查文件类型
	contentType := file.Header.Get("Content-Type")
	if _, ok := v.AllowedTypes[contentType]; !ok {
		return fmt.Errorf("不支持的文件类型：%s", contentType)
	}

	// 检查文件扩展名
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if _, ok := v.AllowedExts[ext]; !ok {
		return fmt.Errorf("不支持的文件扩展名：%s", ext)
	}

	return nil
}
