//无重复的排列
//dfs
func permutation(S string) []string {
    n := len(S)
    ret := []string{}
    visited := map[int]bool{}
    var br func(now string)

    br = func(now string) {
        if len(now) == n {
            ret = append(ret, now)
             return
        }
        for i := 0; i < n; i++ {
            if visited[i] {
                continue
            }
            visited[i] = true
            br(now + string(S[i]))
            visited[i] = false
        }
    }
    br("")
    return ret
}