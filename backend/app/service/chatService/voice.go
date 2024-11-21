package chatService

import (
	"angrymiao-ai/app/serializer/chatSerializer"
	"angrymiao-ai/app/tools/openAiSdk"
	"angrymiao-ai/app/tools/pool"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (s *Service) HandleVoiceChat(args *chatSerializer.VoiceChatBodyRequest, gin *gin.Context) (string, error) {
	var (
		ctx           = context.Background()
		client        = openAiSdk.NewClient()
		chatReplyResp = &chatSerializer.ChatReply{}
	)

	voiceFilePath, err := s.voiceCheckDeal(args, gin)
	if err != nil {
		return "", err
	}
	defer os.Remove(voiceFilePath)

	whisperResp, err := client.OpenAiClient.CreateTranscription(ctx, openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: voiceFilePath,
	})

	if err != nil {
		return "", errors.New(fmt.Sprintf("voice to text err:%v\n", err))
	}

	poolCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resultChan := make(chan *chatSerializer.ChatReplyResp, 1)
	err = pool.Submit(func() {
		response, err := GetChatService().Answer(ctx, &chatSerializer.ChatBodyRequest{
			DeviceCode: args.DeviceCode,
			Message:    &whisperResp.Text,
		})
		if err != nil {
			resultChan <- &chatSerializer.ChatReplyResp{
				Error: err.Error(),
			}
			return
		}
		resultChan <- &chatSerializer.ChatReplyResp{
			ChatReply: response,
		}
	})

	if err != nil {
		return "", err
	}

	select {
	case result := <-resultChan:
		if result.Error != "" {
			return "", errors.New(result.Error)
		}

		chatReplyResp = result.ChatReply
	case <-poolCtx.Done():
		return "", errors.New("timeout")
	}

	if chatReplyResp.ReplyMsg == "" {
		chatReplyResp.ReplyMsg = "还未能理解你的意思，请再重复一次"
	}
	ttsFilePath, err := s.textToVoice(chatReplyResp.ReplyMsg, args.DeviceCode, client)
	if err != nil {
		return "", err
	}

	return ttsFilePath, nil
}

func (s *Service) textToVoice(message, deviceCode string, client *openAiSdk.Client) (string, error) {
	res, err := client.OpenAiClient.CreateSpeech(context.Background(), openai.CreateSpeechRequest{
		Model: openai.TTSModel1,
		Input: message,
		Voice: openai.VoiceAlloy,
	})
	defer res.Close()

	if err != nil {
		return "", err
	}

	buf, err := io.ReadAll(res)
	if err != nil {
		return "", err
	}

	// 指定文件夹路径
	dir := fmt.Sprintf("voice_output/%s", deviceCode)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	// 生成唯一文件名
	fileName := fmt.Sprintf("%s.mp3", uuid.New().String())
	filePath := filepath.Join(dir, fileName)

	// save buf to file
	err = os.WriteFile(filePath, buf, 0644)
	if err != nil {
		return "", nil
	}

	return filePath, nil
}

func (s *Service) voiceCheckDeal(args *chatSerializer.VoiceChatBodyRequest, gin *gin.Context) (string, error) {
	if args.AudioFile == nil || args.AudioFile.Size == 0 {
		return "", errors.New("音频文件为空")
	}

	fileExt := strings.ToLower(filepath.Ext(args.AudioFile.Filename))
	supportedFormats := map[string]bool{
		".mp3": true, ".opus": true, ".aac": true, ".flac": true, ".wav": true, ".pcm": true,
	}
	if !supportedFormats[fileExt] {
		return "", errors.New("不支持的音频格式")
	}

	tempFileName := fmt.Sprintf("%s-%s%s", time.Now().Format("20060102-150405"), uuid.New().String(), fileExt)
	tempFilePath := filepath.Join(os.TempDir(), tempFileName)

	if err := gin.SaveUploadedFile(args.AudioFile, tempFilePath); err != nil {
		return "", fmt.Errorf("保存音频文件失败: %v", err)
	}

	return tempFilePath, nil
}
