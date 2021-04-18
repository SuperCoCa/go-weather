package cmd

import (
	"errors"
	"fmt"
	"go-weather/server"
	"log"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/spf13/cobra"
)

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "查询天气",
	Long:  "查询天气",
}

var realtimeCmd = &cobra.Command{
	Use:   "realtime",
	Short: "获取地区实时天气",
	Long:  "获取地区实时天气",
	Args: func(cmd *cobra.Command, args []string) error {
		return checkArgs(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		city, _ := cmd.Flags().GetString("city")

		// 读取xlsx文件获取城市代码
		adcode, err := getAdcode(city)
		if err != nil {
			log.Fatalf(err.Error())
			os.Exit(1)
		}

		if adcode == 0 {
			log.Fatalf("找不到该地区，请确认输入地区名称是否正确")
			os.Exit(1)
		}

		// 获取天气信息
		info, err := server.GetRealtimeWeather(adcode)
		if err != nil {
			log.Fatalf(err.Error())
			os.Exit(1)
		}

		fmt.Printf("省份：%v\n", info.Lives[0].Province)
		fmt.Printf("城市：%v\n", info.Lives[0].City)
		fmt.Printf("天气现象：%v\n", info.Lives[0].Weather)
		fmt.Printf("实时气温：%v\n", info.Lives[0].Temperature)
		fmt.Printf("风向：%v\n", info.Lives[0].Winddirection)
		fmt.Printf("风力：%v\n", info.Lives[0].Windpower)
		fmt.Printf("空气湿度：%v\n", info.Lives[0].Humidity)
	},
}

var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "获取地区天气预报",
	Long:  "获取地区天气预报",
	Args: func(cmd *cobra.Command, args []string) error {
		return checkArgs(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		city, _ := cmd.Flags().GetString("city")

		// 读取xlsx文件获取城市代码
		adcode, err := getAdcode(city)
		if err != nil {
			log.Fatalf(err.Error())
			os.Exit(1)
		}

		if adcode == 0 {
			log.Fatalf("找不到该地区，请确认输入地区名称是否正确")
			os.Exit(1)
		}

		// 获取天气信息
		info, err := server.GetForecastWeather(adcode)
		if err != nil {
			log.Fatalf(err.Error())
			os.Exit(1)
		}

		weekStr := map[string]string{"1": "一", "2": "二", "3": "三", "4": "四", "5": "五", "6": "六", "7": "日"}

		fmt.Printf("省份：%v\n", info.Forecasts[0].Province)
		fmt.Printf("城市：%v\n", info.Forecasts[0].City)
		for _, cast := range info.Forecasts[0].Casts {
			fmt.Println("-----------------------")
			fmt.Printf("时间：%v 星期%v\n", cast.Date, weekStr[cast.Week])
			fmt.Printf("日间天气现象：%v\n", cast.Dayweather)
			fmt.Printf("日间预测气温：%v\n", cast.Daytemp)
			fmt.Printf("日间风向：%v\n", cast.Daywind)
			fmt.Printf("日间风力：%v\n", cast.Daypower)
			fmt.Printf("夜间天气现象：%v\n", cast.Nightweather)
			fmt.Printf("夜间预测气温：%v\n", cast.Nighttemp)
			fmt.Printf("夜间风向：%v\n", cast.Nightwind)
			fmt.Printf("夜间风力：%v\n", cast.Nightpower)
		}
	},
}

// 校验参数
func checkArgs(cmd *cobra.Command, args []string) error {
	city, err := cmd.Flags().GetString("city")
	if err != nil {
		log.Printf("%v", err)
		return errors.New("请输入城市名称或ID")
	}
	if len(city) == 0 {
		return errors.New("请携带参数 -c 或 --city")
	}

	return nil
}

// 获取地区代码
func getAdcode(city string) (int, error) {
	var adcode int

	// 读取地区代码xlsx文件
	xlsx, err := excelize.OpenFile("adcode.xlsx")
	if err != nil {
		return 0, err
	}

	// 读取xlsx文件中的表Sheet1
	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
		return 0, err
	}

	// 循环Sheet1表中的行并与输入的地区名进行判断
	for _, row := range rows {
		areaRune := []rune(row[0])
		if city == row[0] || city == string(areaRune[:len(areaRune)-1]) {
			adcode, _ = strconv.Atoi(row[1])
		}
	}

	return adcode, nil
}

func init() {
	weatherCmd.AddCommand(realtimeCmd)
	weatherCmd.AddCommand(forecastCmd)

	realtimeCmd.Flags().StringP("city", "c", "广州市", "输入地区全称（e.g. 广州市）")
	forecastCmd.Flags().StringP("city", "c", "广州市", "输入地区全称（e.g. 广州市）")
}
