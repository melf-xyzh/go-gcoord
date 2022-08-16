/**
 * @Time    :2022/8/16 7:52
 * @Author  :Xiaoyu.Zhang
 */

package gcoord

// GpsToAmap
/**
 *  @Description: GPS转高德
 *  @param lon 原经度
 *  @param lat 原纬度
 *  @return mgLon 转换后的经度
 *  @return mgLat 转换后的纬度
 */
func GpsToAmap(lon, lat float64) (mgLon, mgLat float64) {
	return WGS84ToGCJ02(lon, lat)
}

// GpsToBaiduMap
/**
 *  @Description: GPS转百度
 *  @param lon 原经度
 *  @param lat 原纬度
 *  @return mgLon 转换后的经度
 *  @return mgLat 转换后的纬度
 */
func GpsToBaiduMap(lon, lat float64) (mgLon, mgLat float64) {
	lon1, lat1 := WGS84ToGCJ02(lon, lat)
	return GCJ02ToBD09(lon1, lat1)
}
