/**
 * @Time    :2022/8/5 11:11
 * @Author  :Xiaoyu.Zhang
 */

package gcoord

import (
	"fmt"
	"math"
	"strconv"
)

// isOutOfChina
/**
 *  @Description: 判断是否国外坐标
 *  @param lon 经度
 *  @param lat 纬度
 *  @return isOutChina 是否在国外
 */
func isOutOfChina(lon, lat float64) (isOutChina bool) {
	if lon < 72.004 || lon > 137.8347 {
		return true
	} else {
		return lat < 0.8293 || lat > 55.8271
	}
}

// transformLat
/**
 *  @Description: 纬度转换
 *  @param x
 *  @param y
 *  @return lAt
 */
func transformLat(x, y float64) (lAt float64) {
	num := -100.0 + 2.0*x + 3.0*y + 0.2*y*y + 0.1*x*y + 0.2*math.Sqrt(math.Abs(x))
	num += (20.0*math.Sin(6.0*x*PI) + 20.0*math.Sin(2.0*x*PI)) * 2.0 / 3.0
	num += (20.0*math.Sin(y*PI) + 40.0*math.Sin(y/3.0*PI)) * 2.0 / 3.0
	num += (160.0*math.Sin(y/12.0*PI) + 320*math.Sin(y*PI/30.0)) * 2.0 / 3.0
	return num
}

// transformLon
/**
 *  @Description: 经度转换
 *  @param x
 *  @param y
 *  @return lon
 */
func transformLon(x, y float64) (lon float64) {
	num := 300.0 + x + 2.0*y + 0.1*x*x + 0.1*x*y + 0.1*math.Sqrt(math.Abs(x))
	num += (20.0*math.Sin(6.0*x*PI) + 20.0*math.Sin(2.0*x*PI)) * 2.0 / 3.0
	num += (20.0*math.Sin(x*PI) + 40.0*math.Sin(x/3.0*PI)) * 2.0 / 3.0
	num += (150.0*math.Sin(x/12.0*PI) + 300.0*math.Sin(x/30.0*PI)) * 2.0 / 3.0
	return num
}

// calculation
/**
 *  @Description: 坐标转换
 *  @param longitude
 *  @param latitude
 *  @return lng
 *  @return lAt
 */
func calculation(lon, lat float64) (mgLon, mgLat float64) {
	if isOutOfChina(lon, lat) {
		return lon, lat
	}
	dLat := transformLat(lon-105.0, lat-35.0)
	dLon := transformLon(lon-105.0, lat-35.0)
	radLat := lat / 180.0 * PI
	magic := math.Sin(radLat)
	magic = 1 - EE*magic*magic
	sqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((A * (1 - EE)) / (magic * sqrtMagic) * PI)
	dLon = (dLon * 180.0) / (A / sqrtMagic * math.Cos(radLat) * PI)
	mgLon = lon + dLon
	mgLat = lat + dLat
	mgLon, _ = strconv.ParseFloat(fmt.Sprintf("%.6f", mgLon), 64)
	mgLat, _ = strconv.ParseFloat(fmt.Sprintf("%.6f", mgLat), 64)
	return
}
