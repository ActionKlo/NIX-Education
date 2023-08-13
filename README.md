# NIX Education

This project is created as part of the training at NIX company. Your task is to create a simple REST API service using Golang to retrieve data from [JSONPlaceholder](https://jsonplaceholder.typicode.com/) and store specific data in a MySQL database.

## Task 

1. Fetch the posts and display them in the console using the standard input/output library - fmt. The correct result of solving this task will be the following output to the console: ![beginer.task1.result1.png](https://github.com/ActionKlo/NIX-Education/assets/19651281/960455cc-c301-4d15-a69c-a81a438e5187)
2. Thanks to the implemented multithreading using goroutines, our code is supposed to send 100 requests to fetch posts "simultaneously," and upon receiving the responses, display the query results in the console.
3. Using the code from point 2, instead of displaying the post information in the console, create a file for each post in the "./storage/posts" directory at the project's root and write the post's content there using the standard input/output library for file system operations - [ioutil](https://golang.org/pkg/io/ioutil/). The content of each file will have the following structure: ![beginer.task5.pic2.png](https://github.com/ActionKlo/NIX-Education/assets/19651281/19aea08a-81b3-4fcc-913b-fc935c695244)

    * It's worth noting that when writing the structure to a file, the keys are not preserved; only the values are written separated by a delimiter. Is it possible to write the file structure while preserving the keys?
4. At JSONPlaceholder, there are users who have posts (/posts?userId=7), and these posts have comments (/comments?postId={postId}). Your task is:
   1. to retrieve posts for the user with id=7, and for each obtained post, concurrently retrieve comments. All received comments should be concurrently written to a database.
   2. Use the goroutines and channels that we're familiar with.
   3. Write the posts and comments to a PostgreSQL database using [go-sql-driver/mysql](https://godoc.org/github.com/go-sql-driver/mysql).
   4. You're expected to query comments as soon as you receive information about a post - feel free to initiate routines within other routines. Here's an approximate structure of how the entire process should take place: ![beginer.task6.pic1.jpg](https://github.com/ActionKlo/NIX-Education/assets/19651281/ab769e28-947f-4d44-9a1f-c22ae4405c93)

    As a result, the database should look like this: 
   - Posts: ![beginer.task6.pic2.png](https://github.com/ActionKlo/NIX-Education/assets/19651281/c1917851-a151-4788-be23-cc9cd004c207)
   - Comments: ![beginer.task6.pic3.png](https://github.com/ActionKlo/NIX-Education/assets/19651281/fc8f3992-dc1b-4f8a-ae8b-e009fdbf5a7d)

