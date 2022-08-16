/**
 * @Time    :2022/8/16 7:07
 * @Author  :Xiaoyu.Zhang
 */

package gcoord

import "math"

// WGS84ToGCJ02
/**
 *  @Description: WGS-84 转 GCJ-02
 *  @param lon
 *  @param lat
 *  @return mgLon
 *  @return mgLat
 */
func WGS84ToGCJ02(lon, lat float64) (mgLon, mgLat float64) {
	return calculation(lon, lat)
}

// GCJ02ToWGS84
/**
 *  @Description: GCJ-02 转 WGS-84
 *  @param lon
 *  @param lat
 *  @return mgLon
 *  @return mgLat
 */
func GCJ02ToWGS84(lon, lat float64) (mgLon, mgLat float64) {
	dLat := transformLat(lon-105.0, lat-35.0)
	dLng := transformLon(lon-105.0, lat-35.0)
	radLat := lat / 180.0 * PI
	magic := math.Sin(radLat)
	magic = 1 - EE*magic*magic
	sqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((A * (1 - EE)) / (magic * sqrtMagic) * PI)
	dLng = (dLng * 180.0) / (A / sqrtMagic * math.Cos(radLat) * PI)
	mgLon, mgLat = lon-dLng, lat-dLat
	return
}

// BD09ToGCJ02
/**
 *  @Description: BD-09坐标 转 GCJ-02坐标
 *  @param lon
 *  @param lat
 *  @return mgLon
 *  @return mgLat
 */
func BD09ToGCJ02(lon, lat float64) (mgLon, mgLat float64) {
	x := lon - 0.0065
	y := lat - 0.006
	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*X_PI)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*X_PI)
	return z * math.Cos(theta), z * math.Sin(theta)
}

// GCJ02ToBD09
/**
 *  @Description: GCJ-02坐标 转 BD-09坐标
 *  @param lon
 *  @param lat
 *  @return mgLon
 *  @return mgLat
 */
func GCJ02ToBD09(lon, lat float64) (mgLon, mgLat float64) {
	z := math.Sqrt(lon*lon+lat*lat) + 0.00002*math.Sin(lat*X_PI)
	theta := math.Atan2(lat, lon) + 0.000003*math.Cos(lon*X_PI)
	mgLon = z*math.Cos(theta) + 0.0065
	mgLat = z*math.Sin(theta) + 0.006
	return
}

// MapBarToWGS84
/**
 *  @Description: MapBar坐标 转 WGS-84坐标
 *  @param lon
 *  @param lat
 *  @return mgLon
 *  @return mgLat
 */
func MapBarToWGS84(lon, lat float64) (mgLon, mgLat float64) {
	lon0 := float64(int(lon*100000.0) % 36000000)
	lat0 := float64(int(lat*100000.0) % 36000000)

	lon1 := lon0 - math.Cos(lat0/100000.0)*lon0/18000.0 - math.Sin(lon0/100000.0)*lat0/9000.0
	lat1 := lat0 - math.Sin(lat0/100000.0)*lon0/18000.0 - math.Cos(lon0/100000.0)*lat0/9000.0

	var lonX float64 = 1
	if lon0 <= 0 {
		lonX = -1
	}
	var latX float64 = 1
	if lon0 <= 0 {
		latX = -1
	}
	lon2 := lon0 - math.Cos(lat1/100000.0)*lon1/18000.0 - math.Sin(lon1/100000.0)*lat1/9000.0 + lonX
	lat2 := lat0 - math.Sin(lat1/100000.0)*lon1/18000.0 - math.Cos(lon1/100000.0)*lat1/9000.0 + latX
	mgLon, mgLat = lon2/100000.0, lat2/100000.0
	return
}

// MapBarToGCJ02
/**
 *  @Description: MapBar坐标 转 GCJ-02坐标
 *  @param lon
 *  @param lat
 *  @return mgLon
 *  @return mgLat
 */
func MapBarToGCJ02(lon, lat float64) (mgLon, mgLat float64) {
	lon1, lat1 := MapBarToWGS84(lon, lat)
	mgLon, mgLat = WGS84ToGCJ02(lon1, lat1)
	return
}
