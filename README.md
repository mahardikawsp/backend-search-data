# Go-Lang Clean Architecture
Reference : 
- [uncle-bob blog](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [betterprogramming](https://betterprogramming.pub/the-clean-architecture-beginners-guide-e4b7058c1165)
- [hackernoon](https://hackernoon.com/creating-clean-architecture-using-golang-9h5i3wgr)
- [dev.to](https://dev.to/bmf_san/dive-to-clean-architecture-with-golang-cd4)
- Medium 
  - [golangid](https://medium.com/golangid/mencoba-golang-clean-architecture-c2462f355f41)
  - [jfeng45](https://medium.com/@jfeng45/go-microservice-with-clean-architecture-a08fa916a5db)
  - ...
- ...
 
Reference Repo :
- [Khannedy](https://github.com/khannedy/golang-clean-architecture)
- [bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch) :
  - [v1](https://github.com/bxcodec/go-clean-arch/tree/v1)
  - [v2](https://github.com/bxcodec/go-clean-arch/tree/v2)
  - [master](https://github.com/bxcodec/go-clean-arch)

## Library Framework
Beberapa recommend library framework yg bisa dipakai :
- Web/Engine : 
  - (*alternatif*) GoFiber ( [godocs](https://docs.gofiber.io/) | [github](https://github.com/gofiber/fiber) )
  - (*alternatif*) Echo ( [github](https://github.com/labstack/echo) | [echo](https://echo.labstack.com/) )
- Validation : 
  - (*alternatif*) [Go-Ozzo](https://github.com/go-ozzo)
  - (*alternatif*) [asaskevich/govalidator](https://github.com/asaskevich/govalidator)
  - (*alternatif*) [thedevsaddam/govalidator](https://github.com/thedevsaddam/govalidator)
- Configuration : 
  - (*alternatif*) Viper | [github](https://github.com/spf13/viper)
  - (*alternatif*) GoDotEnv | [pkg.go](https://pkg.go.dev/github.com/joho/godotenv) | [github](https://github.com/joho/godotenv)
- Database : 
  - Postgres | [github](https://github.com/go-pg/pg) | [uptrace](https://pg.uptrace.dev/)
  - Redis | [github](https://github.com/go-redis/redis) | [uptrace](https://redis.uptrace.dev/)
  
## Main Architecture:
- Model : layer yg menyimpan sekumpulan object/data/entity structure
- Router : layer yg berkaitan dengan handler/controller/router/delivery/presenter serta validation request
- Service : layer yg berkaitan dengan business rules/business logic/business process suatu service
- Repository : layer yg menghubungkan suatu service dengan datastore/database

Extra Structure :
- config : layer ini digunakan untuk menyimpan dan mengambil sebuah variable konfigurasi atau proses setup suatu konfigurasi
- thirdparty : layer yg berkaitan dengan vendor/pihak ke 3/service lain
- helper : sebuah package/fungsi bantuan yang dapat digunakan secara global
- middleware : layer yg berada di antara interface adapter layer (handler) dengan application business layer (usecase)
- ...
