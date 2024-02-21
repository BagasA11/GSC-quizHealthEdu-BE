
# Healthy Quiz - Backend Application

Healthy Quiz is a health-themed quiz application to answer solutions in 2024. Healthy quiz was created as an implementation of the 3rd UN SDG, by providing health education to the public.
 


## Badges



![MariaDB](https://img.shields.io/badge/MariaDB-003545?style=for-the-badge&logo=mariadb&logoColor=white)![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)
![Apache](https://img.shields.io/badge/apache-%23D42029.svg?style=for-the-badge&logo=apache&logoColor=white)
![Canva](https://img.shields.io/badge/Canva-%2300C4CC.svg?&style=for-the-badge&logo=Canva&logoColor=white)
![jwt](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=JSON%20web%20tokens&logoColor=white)


## Authors

- [@bagasa11](https://www.github.com/bagasa11)
- [@RAYNF](https://github.com/RAYNF)




## License
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)




## Features

### for User
- Authentication (register, login, logout)
- Grouping quiz
- Leaderboard
- jwt auth
### for Admin
- Authentication (login, logout)
- Create new admin account
- Create, edit, delete, and validate for quiz object
- Manage users
## Run Locally

Clone the project

```bash
  git clone https://github.com/BagasA11/GSC-quizHealthEdu-BE.git
```

install go 1.21 or latest 
```bash
  https://go.dev/dl/
```
move project into go directory
```bash
  local:\go\{GSC-quizHealthEdu-BE.git}
```
go to clone project
```bash
  cd GSC-quizHealthEdu-BE.git
```
run web server and mysql 

Start the server: open cmd 
```bash
    go run main.go
```


## Configure

### mysql
- Import SQL files from SQL Project files to MySQL client application
- open project with code editor
- go to db config file at:
```bash
  GSC-quizHealthEdu-BE/configs/db.config.go
```
- change database name at line this:
```bash
  var Dsn = "root:@tcp(127.0.0.1:3306)/golang-gsc2024?charset=utf8mb4&parseTime=True&loc=Local"
```
- into:
```bash
  var Dsn = "root:@tcp(127.0.0.1:3306)/{db-name}?charset=utf8mb4&parseTime=True&loc=Local"
```

### postman
- find postman collection file in project
```bash
    GSC-quizHealthEdu-BE/GOOGLE SOLUTION CHALLANGE.postman_collection.json
```
- import this file to postman workspace
## 🔗 Mobile apps repository
https://github.com/RAYNF/HealtyQuizz

