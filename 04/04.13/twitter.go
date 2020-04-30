type Twitter struct {
    Tweets []int    //所有推特表
    UserTweets map[int][]int //每个用户发表推特表
    Follows map[int][]int   //关注用户表
    IsFollowMy map[int]bool //是否关注自己
}


/** Initialize your data structure here. */
func Constructor() Twitter {
	// 每一次实例化的时候，都重新分配一次，这样不会造成示例重复

	var Tweets  []int

	// 某用户发的某条推特
	var UserTweets = make(map[int][]int)

	// 某用户关注了哪些用户
	var Follows = make(map[int][]int)

	var IsFollowMy = make(map[int]bool)

	t := Twitter{
		Tweets:Tweets,
		UserTweets:UserTweets,
		Follows: Follows,
		IsFollowMy: IsFollowMy,
	}
	return t
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
    //每人每次发推特都记录到一个地方
    this.Tweets = append(this.Tweets, tweetId)
    //将发的推特存到自己的列表
    this.UserTweets[userId] = append(this.UserTweets[userId], tweetId)
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
    fs := this.Follows[userId] //获取关注列表
    var allTweets []int
    for _, v := range fs {
        //获取所有关注人的tweet
        //下面的append操作相当于连接两个slice
        allTweets = append(allTweets, this.UserTweets[v]...)
    }
    //如果没关注自己，则把自己的推特也加进来
    if !this.IsFollowMy[userId] {
        allTweets = append(allTweets, this.UserTweets[userId]...)
    }
    var sortTweets []int
    aTLen := len(this.Tweets)
    s := 0
    //按照发的推特顺序进行倒序排列
    for i := aTLen-1; i >= 0; i-- {
        if s >= 10 {
            break
        }
        //数组有序的
        for _, n := range allTweets {
            if this.Tweets[i] == n && s < 10 {
                s++
                sortTweets = append(sortTweets, n)
            }
        }
    }
    return sortTweets
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
    //关注了自己，则标注一下
    if followerId == followeeId {
        this.IsFollowMy[followerId] = true
    }
    //判断是否已经关注
    var isFed bool
    for _, v := range this.Follows[followerId] {
        if v == followeeId {
            isFed = true
        }
    }
    //没有关注,进行关注
    if !isFed {
        this.Follows[followerId] = append(this.Follows[followerId], followeeId)
    }
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    //取关自己
    if followeeId == followerId {
        this.IsFollowMy[followerId] = false
    }
    //去掉关注列表里被关在的人
    var tmp []int
    for _, v := range this.Follows[followerId] {
        if v != followeeId {
            tmp = append(tmp, v)
        }
    }
    this.Follows[followerId] = tmp
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */