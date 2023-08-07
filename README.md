# NIX Education

This project is created as part of the training at NIX company. Your task is to create a simple REST API service using Golang to retrieve data from [JSONPlaceholder](https://jsonplaceholder.typicode.com/) and store specific data in a MySQL database.

## Task 

1. Fetch the posts and display them in the console using the standard input/output library - fmt. The correct result of solving this task will be the following output to the console: ![beginer.task1.result1.png](..%2F..%2FDesktop%2Fbeginer.task1.result1.png)
2. Thanks to the implemented multithreading using goroutines, our code is supposed to send 100 requests to fetch posts "simultaneously," and upon receiving the responses, display the query results in the console.
3. Using the code from point 2, instead of displaying the post information in the console, create a file for each post in the "./storage/posts" directory at the project's root and write the post's content there using the standard input/output library for file system operations - [ioutil](https://golang.org/pkg/io/ioutil/). The content of each file will have the following structure: ![beginer.task5.pic2.png](..%2F..%2FDesktop%2Fbeginer.task5.pic2.png)
    * It's worth noting that when writing the structure to a file, the keys are not preserved; only the values are written separated by a delimiter. Is it possible to write the file structure while preserving the keys?
4. At JSONPlaceholder, there are users who have posts (/posts?userId=7), and these posts have comments (/comments?postId={postId}). Your task is:
   1. to retrieve posts for the user with id=7, and for each obtained post, concurrently retrieve comments. All received comments should be concurrently written to a database.
   2. Use the goroutines and channels that we're familiar with.
   3. Write the posts and comments to a MySQL database using [go-sql-driver/mysql](https://godoc.org/github.com/go-sql-driver/mysql).
   4. You're expected to query comments as soon as you receive information about a post - feel free to initiate routines within other routines. Here's an approximate structure of how the entire process should take place: ![beginer.task6.pic1.jpg](..%2F..%2FDesktop%2Fbeginer.task6.pic1.jpg)

    As a result, the database should look like this: 
   - Posts: ![beginer.task6.pic2.png](..%2F..%2FDesktop%2Fbeginer.task6.pic2.png)
   - Comments: ![beginer.task6.pic3.png](..%2F..%2FDesktop%2Fbeginer.task6.pic3.png)

