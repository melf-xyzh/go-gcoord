/**
 * @Time    :2022/8/16 7:52
 * @Author  :Xiaoyu.Zhang
 */

package gcoord

// BaiduMapToAmap
/**
 *  @Description: 百度转高德
 *  @param lon 原经度
 *  @param lat 原纬度
 *  @return mgLon 转换后的经度
 *  @return mgLat 转换后的纬度
 */
func BaiduMapToAmap(lon, lat float64) (mgLon, mgLat float64) {
	return BD09ToGCJ02(lon, lat)
}

// BaiduMapToGPS
/**
 *  @Description: 百度地图转GPS
 *  @param lon 原经度
 *  @param lat 原纬度
 *  @return mgLon 转换后的经度
 *  @return mgLat 转换后的纬度
 */
func BaiduMapToGPS(lon, lat float64) (mgLon, mgLat float64) {
	lon1, lat1 := BD09ToGCJ02(lon, lat)
	return GCJ02ToWGS84(lon1, lat1)
}
