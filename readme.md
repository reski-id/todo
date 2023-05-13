

<h3 align="center">Todo Rest Api <br>
<h5 align="center" >Golang Gin Gorm<h5>
<br>
</h4>
<p align="left">
<h2>
  Content <br></h2>
  • Key Features <br>
  • Installing Using Github<br>
  • Installing Using Docker<br>
  • End Point<br>
  • Technology<br>
  • Contact me<br>
</p>

## 📱 Features

* Activity
* Todo


## ⚙️ Installing from Github

installing and running the app from github repository <br>
To clone and run this application, you'll need [Git](https://git-scm.com) and [Golang](https://go.dev/dl/) installed on your computer. From your command line:

```bash
# Clone this repository
$ git clone https://github.com/reski-id/todolist.git

# Go into the repository
$ cd todolist

# Install dependencies
$ go get

# Run the app
$ go run main.go

# if you have problem while running you can use bash cmd and type this..
$ source .env #then type 
$ go run main.go 
```

> **Note**
> Make sure you allready create database mysql `todolist` for this app.more info in local `.env` file.


## ⚙️ Installing and Runing with Docker
if you are using docker or aws/google cloud server you can run this application by creating a container. <br>

```bash
# Pull this latest app from dockerhub 
$ docker pull programmerreski/todolist-api

# if you have mysql container you can skip this
$ docker pull mysql

$ docker run --name mysqlku -p 3306:3306 -d -e MYSQL_ROOT_PASSWORD="yourmysqlpassword" mysql 

# create app container
$ docker run --name todo -p 80:8000 -d --link mysqlku -e SECRET="secr3t" -e SERVERPORT=8000 -e Name="todo" -e Address=mysqlku -e Port=3306 -e Username="root" -e Password="yourmysqlpassword" programmerreski/todolist-api

# Run the app
$ docker logs todo
```

## 📜 End Point  

Activity
| Methode       | End Point      | used for            
| ------------- | -------------  | -----------                  
| `GET`         | /activity-groups            | Get all Activity      
| `GET`         | /activity-groups/:id         | Get One Activity          
| `POST`        | /activity-groups              | Insert Activity 
| `PUT`         | /activity-groups/:id         | Update data Activity
| `DELETE`      | /activity-groups/:id         | Delete Activity  

Todo
| Methode       | End Point      | used for            
| ------------- | -------------  | -----------                  
| `GET`         | /todo-items           | Get all todo      
| `GET`         | /todo-items/:id         | Get One todo          
| `POST`        | /todo-items              | Insert todo 
| `PUT`         | /todo-items/:id         | Update data todo
| `DELETE`      | /todo-items/:id        | Delete todo  


## 🛠️ Technology

This software uses the following Tech:

- [Golang](https://go.dev/dl/)
- [Gin Framework](https://gin-gonic.com/)
- [Gorm](https://gorm.io/index.html)
- [mysql](https://www.mysql.com/)
- [Linux](https://www.linux.com/)
- [Docker](https://www.docker.com/)
- [Dockerhub](https://hub.docker.com/u/programmerreski)
- [Git Repository](https://github.com/reski-id)
- [Trunk Base Development](https://trunkbaseddevelopment.com/)


## 📱 Contact me
feel free to contact me ... 
- Email programmer.reski@gmail.com 
- [Linkedin](https://www.linkedin.com/in/reski-id)
- [Github](https://github.com/reski-id)
- Whatsapp <a href="https://wa.me/+6281261478432?text=Hello">Send WhatsApp Message</a>
