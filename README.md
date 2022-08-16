## go-gcoord

WGS84 / BD09 / GCJ02 / MapBar 经纬度坐标互转

### 使用

```bash
go get github.com/melf-xyzh/go-gcoord
```

用例

```go
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
```

#### 参考文档

> https://blog.csdn.net/u011439796/article/details/90050390/
> https://blog.csdn.net/machao0_0/article/details/120440208
> https://blog.csdn.net/xiaoxiaovbb/article/details/123111209