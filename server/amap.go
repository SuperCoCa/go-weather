package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-weather/entity"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	Key = "这里填写个人的key值"
	Uri = "https://restapi.amap.com/v3/weather/weatherInfo"
)

func GetRealtimeWeather(code int) (*entity.Response, error) {
	info, err := WeatherRequest(code, "base")
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func GetForecastWeather(code int) (*entity.Response, error) {
	info, err := WeatherRequest(code, "all")
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func WeatherRequest(code int, extensions string) (entity.Response, error) {
	info := entity.Response{}
	client := &http.Client{Timeout: 5 * time.Second}
	url := fmt.Sprintf(Uri+"?key=%s&city=%d&extensions=%s", Key, code, extensions)
	res, err := client.Get(url)
	if err != nil {
		log.Fatalf("查询错误：%v", err)
		return info, errors.New("查询失败")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("查询错误：%v", err)
		return info, errors.New("查询失败")
	}
	if err := json.Unmarshal(body, &info); err != nil {
		fmt.Printf("查询错误：%v", err)
		return info, errors.New("查询失败")
	}

	// 1：成功，0：失败
	if info.Status != "1" {
		return info, errors.New(info.Info)
	}

	return info, nil
}
