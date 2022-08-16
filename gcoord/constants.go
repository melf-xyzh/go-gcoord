/**
 * @Time    :2022/8/5 11:12
 * @Author  :Xiaoyu.Zhang
 */

package gcoord

const (
	PI  = 3.1415926535897932384626 // 圆周率
	EE  = 0.00669342162296594323   // 椭球的偏心率
	A   = 6378245.0                // 卫星椭球坐标投影到平面地图坐标系的投影因子
	X_PI = PI * 3000.0 / 180.0
)


// 参考文档
// https://blog.csdn.net/u011439796/article/details/90050390/
// https://blog.csdn.net/machao0_0/article/details/120440208
// https://blog.csdn.net/xiaoxiaovbb/article/details/123111209