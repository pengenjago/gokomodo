# Gokomodo Backend Technical Assessment

Start Development
1. Change config database in :
   - application.properties
   - docker.compose.yml 
     > If running in Docker
   - auth_test.go
     > Change dsn same as in **application.properties** if want to run unit test

2. Install swaggo
     > $ go get -u github.com/swaggo/swag/cmd/swag

3. Generate swagger if api change
     > $ swag init

4. Run App
     > go run main.go

5. After Run App, system will generate Tabel. For dummy data can execute **t_gokomodo.sql** 
     
     ````
     Account Seller :
          email : deb@gmail.com
          email : clothing@gmail.com
          password : Seller123!
          
     Account Buyer :
     1.
          email : firmanelpri@gmail.com
          password : Dev1234!
     2.
          email : budiman@gmail.com
          password : Buyer123!
     ````

6. Open Swagger using browser
     > http://localhost:8080/swagger/index.html