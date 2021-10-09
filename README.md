# Task  | Instagram Backend API


Name: Avaneesh Kanshi   
Registration Number: 19BCT0072  
University: Vellore Institute of Technology, Vellore

## Steps to Test



### 1. Run Main File main.go

```
go run main.go
```

### 2. Use curl to simulate the requests

    
* To create a new Instagram user
```
curl --location --request POST 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"test_1",
    "email":"testing.main@mail.com",
    "password":"passkey"
}'
```
![](test_images/createuser.png?raw=true "Title") 

* To get User details from User ID
```
curl --location --request GET 'http://localhost:8080/users/6161952a782330154b5daca4'
```
![](test_images/getuser.png?raw=true "Title") 
* To create a new post using User ID
```
curl --location --request POST 'http://localhost:8080/posts' \
--header 'Content-Type: application/json' \
--data-raw '{
    "caption": "Sample Caption",
    "imageURL": "https://image.shutterstock.com/image-vector/sample-stamp-grunge-texture-vector-600w-1389188327.jpg",
    "time": "12:45 PM Saturday",
    "userID": "6161952a782330154b5daca4"
    
}'
```
![](test_images/createpost.png?raw=true "Title") 

* To get a post using Post ID
```
curl --location --request GET 'http://localhost:8080/posts/6161c3a4b5b588a9094ad0c9'
```
![](test_images/getpostid.png?raw=true "Title")  
and if you want to move through pages then use the following method
```
curl --location --request GET 'http://localhost:8080/posts/6161c3a4b5b588a9094ad0c9/2'
```
![](test_images/paginate.png?raw=true "Title")  





## Version History

* 0.1
    * Initial Release

