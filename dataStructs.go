package dataStructs // Does this file need to be a golang package?

type User struct {
    userID   	     int // 1 - auto-imcremented by sqlite // 2 //
    name 	     string // "Alice" // "Bob" //
    email	     string // "alice.bloggs@gmail.com" // bob.smith@yahoo.com //
    password         string // "MyPa$$word" // "mYsECR£t //
    hashedPass	     string // 
    postsIDs         []int // [1, 3, 5] // Alice posted postID 1, 3 and 5 // [2, 4] //Bob posted postID, 2 and 4 
}

type Session struct {
    sessionID        string // the session unique ID: UUID
    token  	     string  // The Cookie idetifier?
    expiry           time.Time // need to import "time"
    userUUID 	     string // ???
    IsAuthorized     bool // True // ???
}
type Post struct {
    postID           int // 1 // auto-incremented or decremented >= 0
    titlePost        string // "Welcome to the Forum"
    contentPost      string // "Hi, thanks for visiting our forum. This is the first post. We... "
    userID  	     int // 1 // Alice's userID
    isAuthor         bool // true // Boolean check if the viewer is author // Why check? Why? Can Author can edit or delete their posts?
    postDate         string? // Import "time", what data type is there for time? TBC
    likesCount	     int // 5 // auto-incremented or auto-decremented >= 0 // Can a user like a post more than once?
    dislikesCount    int // 1 // auto-incremented or auto-decremented >= 0 // Can a user both like and dislike a post?
    commentsCount    int // 2 // auto-incremented or auto-decremented >= 0
    commentIDs	     []int // [1, 3] // array of commentID linked to each post >= 0
    categoryIDs	     []int // [0, 2] // array of one or more categories selected for each post // 3+ categories. I forgot names but in notes somewhere.
}

type Category struct {
    categoryID       int // 0 // 1 // 2 // If I remeber correctly, we decided on three categories. 0 and/or 1 and/or 2. Can I select all categories for a post?
    titleCategory    string // "Sex" // "Money" // "Rock'n'Roll" // <--Please add correct titles
    postsIDs         []int // [3, 5] // postID 3 and 5 tagged with categoryID 2 "Rock'n'Roll"  // [1] postID 1 tagged with categoryID 1 "Money"
}
	
type Comment struct {
    commentID  	     int // 1 // auto-incremented or decremented >= 0
    postID     	     int // 1 // Alice's original post which the comment is linked to
    userID     	     int // 2 // userID 2 , Bob commented on Alice's postID 1 "Welcome to the Forum" 
    contentComment   string // "Hi Alice, I really like your post. Can I suggest you... "
    commentDate      string // import package "time", what data type is there for time? TBC
    likesCount	     int // 0 // 1 // auto-incremented or auto-decremented >= 0 // Can a user like a comment more than once?
    dislikesCount    int // 1 // 3 // auto-incremented or auto-decremented >= 0 // Can a user both like and dislike a comment?
    isAuthor         bool // true // Boolean check if the viewer is author // Author of what? The post? Why check? Can Author can edit or delete their comments?
}
	
type Error struct {
    errorMessage     string // "The Username or Password is incorrect" // "You are now logged out" // 
}

type Log struct {
    logID	    int // 1 // 2 // auto-increment >= 1
    logFileName	    string // "forum-log-2023-07-11.txt" // forum-log-2023-07-12.txt
}
