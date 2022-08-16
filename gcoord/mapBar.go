/**
 * @Time    :2022/8/16 7:58
 * @Author  :Xiaoyu.Zhang
 */

package gcoord

// MapBarToGps
/**
 *  @Description: MapBar 转 GPS
 *  @param lon
 *  @param lat
 *  @return mgLon
 *  @return mgLat
 */
func MapBarToGps(lon, lat float64) (mgLon, mgLat float64) {
	return MapBarToWGS84(lon, lat)
}

// MapBarToAmap
/**
 *  @Description: MapBar 转 高德
 *  @param lon
 *  @param lat
 *  @return mgLon
 *  @return mgLat
 */
func MapBarToAmap(lon, lat float64) (mgLon, mgLat float64) {
	return MapBarToGCJ02(lon, lat)
}

// MapBarToBaiduMap
/**
 *  @Description: MapBar 转 百度
 *  @param lon
 *  @param lat
 *  @return mgLon
 *  @return mgLat
 */
func MapBarToBaiduMap(lon, lat float64) (mgLon, mgLat float64) {
	lon1,lat1 :=MapBarToGCJ02(lon, lat)
	return GCJ02ToBD09(lon1,lat1)
}