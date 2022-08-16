/**
 * @Time    :2022/8/5 10:04
 * @Author  :Xiaoyu.Zhang
 */

package main

import (
	"fmt"
	"github.com/melf-xyzh/go-gcoord/gcoord"
)

func main() {
	// 定义经纬度
	lon := 114.878074
	lat := 40.837259
	// 高德转百度
	mgLon, mgLat := gcoord.AmapToBaiduMap(lon, lat)
	fmt.Println(mgLon, mgLat)
	// 高德转GPS
	mgLon, mgLat = gcoord.AmapToGPS(lon, lat)
	fmt.Println(mgLon, mgLat)

	// 百度转高德
	mgLon, mgLat = gcoord.BaiduMapToAmap(lon, lat)
	fmt.Println(mgLon, mgLat)

	// 百度转GPS
	mgLon, mgLat = gcoord.BaiduMapToGPS(lon, lat)
	fmt.Println(mgLon, mgLat)

	// GPS转高德
	mgLon, mgLat = gcoord.GpsToAmap(lon, lat)
	fmt.Println(mgLon, mgLat)
	// GPS转百度
	mgLon, mgLat = gcoord.GpsToBaiduMap(lon, lat)
	fmt.Println(mgLon, mgLat)

	// 图吧转高德
	mgLon, mgLat = gcoord.MapBarToAmap(lon, lat)
	fmt.Println(mgLon, mgLat)
	// 图吧转百度
	mgLon, mgLat = gcoord.MapBarToBaiduMap(lon, lat)
	fmt.Println(mgLon, mgLat)
	// 图吧转GPS
	mgLon, mgLat = gcoord.MapBarToGps(lon, lat)
	fmt.Println(mgLon, mgLat)
}
