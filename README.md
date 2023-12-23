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

### User Data
```sh 
    #!/bin/bash
    
    echo -e \nCreating volume...
    mount_path=/mnt/graduation
    mkfs -t ext4 /dev/sdb
    mkdir $mount_path
    mount /dev/sdb /mnt/graduation
    
    echo -e \nDownloading App...
    archive_name=app.tar.gz
    wget https://github.com/afifurrohman-id/graduation-management/releases/download/v0.2.0/graduation-management_0.2.0_linux_amd64.tar.gz -O $archive_name
    tar xzf $archive_name
    
    echo -e \nStarting App...
    export POSTGRES_HOST=<rds_endpoint>
    export POSTGRES_USER=<username>
    export POSTGRES_PASSWORD=<password>
    export POSTGRES_DB=graduation
    export PORT=3000
    export APP_ENV=production
    export LOG_PATH=$mount_path/app.log
    ./graduation-management &
```
