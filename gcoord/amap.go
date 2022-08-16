/**
 * @Time    :2022/8/5 10:05
 * @Author  :XiAoyu.ZhAng
 */

package gcoord

// AmapToGPS
/**
 *  @Description: 高德转GPS
 *  @param lon 原经度
 *  @param lat 原纬度
 *  @return mgLon 转换后的经度
 *  @return mgLat 转换后的纬度
 */
func AmapToGPS(lon, lat float64) (mgLon, mgLat float64) {
	mgLon, mgLat = GCJ02ToWGS84(lon, lat)
	return
}

// AmapToBaiduMap
/**
 *  @Description: 高德转百度
 *  @param lon 原经度
 *  @param lat 原纬度
 *  @return mgLon 转换后的经度
 *  @return mgLat 转换后的纬度
 */
func AmapToBaiduMap(lon, lat float64) (mgLon, mgLat float64) {
	mgLon, mgLat = GCJ02ToBD09(lon, lat)
	return
}
