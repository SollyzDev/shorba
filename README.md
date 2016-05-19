#Shorba

Shorba is a small tool to generate dummy/fake data in Go based on struct (serving as model) and insert those data into mongodb.

let's suppose you have a struct like that:

	type Post struct {
		ID string `bson:"_id,omitempty"`
		Title string `bson:"title"`
		Content string `bson:"content"`
		PublishDate time.Time `bson:"publishDate"`
	}


and you want to populate 200 fake posts for testing and develpoment purposes,
you can use shorba to do that for you, like this:

	postModel := &Post{}
	shorba.populate("posts", &postModel, 200)


this should generate 200 different posts and insert them into collection 'posts'


####Usage

first install shorba

	go get github.com/eslammostafa/shorba

import it

	import (
		...
		"github.com/eslammostafa/shorba"
	)

connect to your mongodb

	shorba.Connect("localhost", "username", "password", "dbname")

 if there is no username nor password write "" instead of each one

 	shorba.Connect("localhost", "", "", "dbname")


populate data

	postModel := &Post{}
	shorba.populate("posts", &postModel, 135)


#### TODO

* Support arrays
* Support Embedded Documents
