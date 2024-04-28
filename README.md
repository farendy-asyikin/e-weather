# e-weather

## Built With

- [Go 1.22.1](https://go.dev/)
- [Gin](https://gin-gonic.com/)
- [Gorm](https://gorm.io/)

## Getting Started

- Clone this repository
  ```
  git clone git@github.com:farendy-asyikin/e-weather.git
  ```
- Copy `.env.example` to `.env`
  ```
  cp .env.example .env
  ```
- Install dependencies
  ```
  go mod download
  ```
- Run the server
  ```
  go run main.go
  ```
- Run the data seeder

- Database using Postgre SQL

## Contributing

1. Create new branch with the following format
   ```
   feat/your-feature
   fix/your-fix
   ```
2. Commit your changes according to the rules of [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
3. Push your branch
4. Create a pull request with the following format
   ```
   [Feature/Fix]: your pull request title
   ```
5. Add description to your pull request like cURL example and the expected response
6. Ask for review
7. After approved, then can be merged

## Postman Collection
    https://www.postman.com/grey-crescent-340741/workspace/m-cineplex/collection/33809558-f55c7e37-3f6b-474b-90b6-6a1a61c2f31d?action=share&creator=33809558&active-environment=33809558-c744b13b-0219-4117-aca3-a49964673b81