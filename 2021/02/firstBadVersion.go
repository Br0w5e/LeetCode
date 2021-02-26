package main

//278. 第一个错误的版本

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

//二分法
func firstBadVersion(n int) int {
	l, r := 0, n
	for l < r {
		mid := l + (r-l)>>1
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
