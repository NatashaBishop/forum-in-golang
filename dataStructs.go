package dataStructs

type User struct {
    userID   	 string
    name 	 string
    email	 string
    password     string
    postsIDs     []int
}

type Post struct {
    postID      string
    titlePost     string
    contentPost   string
    userID  	  int
    isAuthor      bool //check if the viewer is author
    postDate      string //have a look @ the Go package "time", what data type is there for time?
    likesCount	  int
    dislikesCount int
    commentsCount int
    categories	  []string // one to many

}

type Comment struct {
    commentID  	string
    postID     		int
    userID     		int
    commentContent   	string
    commentDate   	string //have a look @ the Go package "time", what data type is there for time?
    likesCount	  	int
    dislikesCount 	int
    isAuthor      	bool //check if the viewer is author
}

type Session struct {
	sessionID    	string
	token  		string
        expiry   	time.Time // need to import "time"
	userUUID 	string
	IsAuthorized 	bool
}


type Error struct {
	errorMessage string
}
