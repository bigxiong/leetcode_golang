package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"time"
)

type Twitter struct {
	// 用户列表
	users map[int]*User
}

type User struct {
	id int
	// 该用户关注用户列表
	followed map[int]*User
	// 该用户发送推文列表
	head *list.List
}

func NewUser(id int) *User {
	u := User{id: id, head: list.New(), followed: make(map[int]*User)}
	u.followed[u.id] = &u
	return &u
}

func (u *User) PostTweet(tweetId int)  {
	tweet := NewTweet(tweetId, "")
	u.head.PushFront(tweet)
}

func (u *User) Follow(followee *User) {
	u.followed[followee.id] = followee
}

func (u *User) UnFollow(followeeId int) {
	delete(u.followed, followeeId)
}

type Tweet struct {
	id int
	time int64
	text string
}

func NewTweet(id int, text string) *Tweet {
	return &Tweet{id: id, text: text, time: time.Now().UnixNano()}
}


/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{users: make(map[int]*User)}
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	_, ok := this.users[userId]
	if !ok {
		this.users[userId] = NewUser(userId)
	}
	this.users[userId].PostTweet(tweetId)
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	u, ok := this.users[userId]
	if !ok {
		return nil
	}
	var queue PriorityQueue
	heap.Init(&queue)
	for _, v := range u.followed {
		if v.head.Len() > 0 {
			heap.Push(&queue, v.head.Front())
		}
	}
	var res []int
	for queue.Len() > 0  {
		if len(res) == 10 {
			break
		}
		top := heap.Pop(&queue).(*list.Element)
		res = append(res, top.Value.(*Tweet).id)
		if top.Next() != nil {
			heap.Push(&queue, top.Next())
		}
	}
	return res
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	_, ok := this.users[followerId]
	if !ok {
		this.users[followerId] = NewUser(followerId)
	}
	_, ok = this.users[followeeId]
	if !ok {
		this.users[followeeId] = NewUser(followeeId)
	}

	this.users[followerId].Follow(this.users[followeeId])
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	if followeeId == followerId {
		return
	}
	follower, ok := this.users[followerId]
	if !ok {
		return
	}
	follower.UnFollow(followeeId)
}


type PriorityQueue []*list.Element

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Value.(*Tweet).time > pq[j].Value.(*Tweet).time
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*list.Element)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func main()  {
	/*
	["Twitter","postTweet","unfollow","getNewsFeed"]
	[[],[1,5],[1,1],[1]]

	*/
	twitter := Constructor()
	twitter.PostTweet(1,1)
	feeds := twitter.GetNewsFeed(1)
	fmt.Println(feeds)

	twitter.Follow(2,1)
	feeds = twitter.GetNewsFeed(2)
	fmt.Println(feeds)
	twitter.Unfollow(2,1)
	feeds = twitter.GetNewsFeed(2)
	fmt.Println(feeds)
}