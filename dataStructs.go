package dataScucts

type User struct {
    Id   	int
    Name 	string
    Email	string
    Password string
}

type Post struct {
    id        	  int
    titlePost     string
    contentPost   string
    authorId  	  int
    postDate      string
    likesCount	  int
    dislikesCount int
    commentsCount int
    categories	  []string
    isAuthor      bool //check if the viewer is author
}

type Comment struct {
    Id          	int
    PostId     	  int
    AuthorId      int
    Content   	  string
    commentDate   string
    LikesCount	  int
    DislikesCount int
    isAuthor      bool //check if the viewer is author
}

type Session struct {
    Id 	   int
    Token  string
    Expire int
    UserId int
}
