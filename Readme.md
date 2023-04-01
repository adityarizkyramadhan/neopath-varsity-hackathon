# Neopath API Versity Hackathon
# <strong>Overview</strong><br>

- Language ```Go``` <br>
- Framework ```Gin``` and ```Gorm``` <br>
- Database ```Postgres```<br>

# <strong>Architecture</strong><br>
- Model <br>
    => Entity data <br>
- Config <br>
    => Configuration to another service like database, etc
- Repositories <br>
    => Service to interact with database <br>
- Controller <br>
    => Service to interact with user<br>
- Usecase <br>
    => Service for logic or bussiness <br>
- Middleware <br>
    => Authentication <br>
    => Authorization to check if the user is allowed to do the action<br>
- Utils <br>
    => Helper function to make response<br>

# <strong>How to run the app</strong><br>
- Copy .env.example to env <br>
    => ``` cp .env.example .env``` <br>
- Complete .env <br>
- Docker compose <br>
    => ```docker-compose up --build -d``` <br>






