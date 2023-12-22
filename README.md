# Graduation Management

## Usage

- Clone the repository

```sh
git clone https://github.com/afifurrohman-id/graduation-management.git
```

- Go to Project Directory

```sh
cd graduation-management
```

- Run Docker Compose

```sh
docker compose -f deployments/compose.yaml up -d
```

- Insert environment variables to `.env` file

```sh
cat <<EOENV > deployments/.env

POSTGRES_HOST=example.com
POSTGRES_USER=user
POSTGRES_PASSWORD=example
POSTGRES_DB=graduation
PORT=3000
LOG_PATH=path/to/app.log

EOENV
```

- Run App
    
```sh 
    go run main.go
```