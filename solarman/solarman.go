package solarman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/Coolknight/transmission-telegram-bot/config"
)

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Uid          int    `json:"uid"`
	Msg          string `json:"msg"`
	Success      bool   `json:"success"`
	RequestId    string `json:"requestId"`
}

type DataList struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Unit  string `json:"unit"`
	Value string `json:"value"`
}

type DeviceDataResponse struct {
	Code        string     `json:"code"`
	DataList    []DataList `json:"dataList"`
	DeviceId    int        `json:"deviceId"`
	DeviceSn    string     `json:"deviceSn"`
	DeviceState int        `json:"deviceState"`
	DeviceType  string     `json:"deviceType"`
	Msg         string     `json:"msg"`
	Success     bool       `json:"success"`
	RequestId   string     `json:"requestId"`
}

func getAuthToken(appId, appSecret, email, password, authURL string) (string, error) {
	payload := map[string]interface{}{
		"appSecret": appSecret,
		"email":     email,
		"password":  password,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", authURL+"?appId="+appId, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var authResponse AuthResponse
	if err := json.Unmarshal(body, &authResponse); err != nil {
		return "", err
	}

	if !authResponse.Success {
		return "", fmt.Errorf("failed to get auth token: %s", authResponse.Msg)
	}

	return authResponse.AccessToken, nil
}

func pollAPI(deviceSn, accessToken, apiURL string) (int, error) {
	payload := map[string]interface{}{
		"deviceSn": deviceSn,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var deviceDataResponse DeviceDataResponse
	if err := json.Unmarshal(body, &deviceDataResponse); err != nil {
		return 0, err
	}

	if !deviceDataResponse.Success {
		if deviceDataResponse.Msg == "auth invalid token" {
			return 0, fmt.Errorf("invalid token")
		}
		return 0, fmt.Errorf("failed to get device data: %s", deviceDataResponse.Msg)
	}

	return deviceDataResponse.DeviceState, nil
}

func sendAlert(telegramBotToken, telegramChatID, message string) error {
	telegramURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramBotToken)
	resp, err := http.PostForm(telegramURL, url.Values{
		"chat_id": {telegramChatID},
		"text":    {message},
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	return nil
}

func deviceStateMessage(state int) string {
	switch state {
	case 1:
		return "Inverter is online." // All OK
	case 2:
		return "Inverter is alerting." // Something is bad (no AC)
	case 3:
		return "Inverter is offline." // There is no Sun
	default:
		return "Unknown inverter state."
	}
}

func ApiAlert(cfg *config.Config) {
	token, err := getAuthToken(cfg.Solarman.AppId, cfg.Solarman.AppSecret, cfg.Solarman.Email,
		cfg.Solarman.Password, cfg.API.AuthURL)
	if err != nil {
		log.Fatalf("Error getting initial auth token: %v", err)
	}

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		deviceState, err := pollAPI(cfg.Device.DeviceSn, token, cfg.API.ApiURL)
		if err != nil {
			if err.Error() == "invalid token" {
				log.Println("Token expired, fetching a new one.")
				token, err = getAuthToken(cfg.Solarman.AppId, cfg.Solarman.AppSecret, cfg.Solarman.Email,
					cfg.Solarman.Password, cfg.API.AuthURL)
				if err != nil {
					log.Printf("Error getting new auth token: %v", err)
					continue
				}
				deviceState, err = pollAPI(cfg.Device.DeviceSn, token, cfg.API.ApiURL)
				if err != nil {
					log.Printf("Error polling API with new token: %v", err)
					continue
				}
			} else {
				log.Printf("Error polling API: %v", err)
				continue
			}
		}

		message := deviceStateMessage(deviceState)
		if deviceState == 2 {
			err = sendAlert(cfg.Telegram.BotToken, cfg.Telegram.ChatID, fmt.Sprintf("Alert! %s", message))
			if err != nil {
				log.Printf("Error sending alert: %v", err)
			}
		}
	}
}
