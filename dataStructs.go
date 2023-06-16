package dataStructs

type User struct {
    userID   	 int
    Name 	 string
    Email	 string
    Password     string
    PostsIDs     []int
}

type Post struct {
    postID        int
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
    commentID  		int
    postID     		int
    userID     		int
    commentContent   	string
    commentDate   	string //have a look @ the Go package "time", what data type is there for time?
    likesCount	  	int
    dislikesCount 	int
    isAuthor      	bool //check if the viewer is author
}

type Session struct {
	sessionID    	int
	token  		string
        expiry   	time.Time // need to import "time"
	userID 		int
}


type Error struct {
	ErrorMessage string
}
