#Shorba

Shorba is a small tool to generate dummy data in Go based on struct (serving as model) and insert these data into mongodb with any quantity needed.

for example:

	shorba.populate("posts", &postModel, 200)


this should generate 200 different posts and insert them into collection 'posts'
