# Test Teknis <img src="https://teknologicakrainternasional.com/assets/Home/Logo.svg" width="25" alt="Laravel Logo"> (PT. Teknologi Cakra Internationsal)

## Requirement

- Golang **(v1.19.1)**
- Docker **(v20.10.17)**

## Usage

Duplikat file **static.env.example** lalu rename menjadi **static.env**, dan isi variablenya seperti berikut berikut

```dotenv
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=nats_postgres
POSTGRES_PASS=nats_postgres
POSTGRES_DBNAME=cakra_test
POSTGRES_SSLMODE=disable
POSTGRES_MAX_OPEN_CONNS=10
POSTGRES_MAX_IDLE_CONNS=10
POSTGRES_LIFETIME_OPEN_CONNS=10
POSTGRES_LIFETIME_IDLE_CONNS=10

APP_PORT=1010

```

Jalankan database dari Docker

```bash
docker compose up -d
```

Setelah itu aplikasi sudah dapat dijalankan dengan mengetikkan

```bash
go run cmd/main.go
```
## Sql Test

### Query
``` sql
  SELECT * FROM crosstab(
	$$ 
        SELECT brand, type, CONCAT(type,':',price) as newPrice 
        FROM cars ORDER BY 1,2 
	$$
  ) AS ct(brand varchar, Type1 text, Type2 text, Type3 text);
```
### Output :

<center>
<img src="https://drive.google.com/uc?export=view&id=1gza5ecDCmWwEzMjtSMwpdgmdffk5Z0ag" width="80%" alt="Laravel Logo">
</center>
<br>
## API Spec

## Cars

### Post /cars

Request Body

```json
{
    "price": 215000000,
    "brand": "Daihatsu",
    "type": "Xenia"
}
```

Response Body

```json
{
    "status": 201,
    "success": true,
    "message": "created data success"
}
```

### Get /cars

Response Body

```json
{
    "status": 200,
    "success": true,
    "message": "get data success",
    "payload": [
        {
            "id": 1,
            "price": 100000,
            "brand": "Honda",
            "type": "Civic"
        },
        {
            "id": 2,
            "price": 236000000,
            "brand": "Honda",
            "type": "Jazz"
        },
        {
            "id": 3,
            "price": 224000000,
            "brand": "Honda",
            "type": "Mobilio"
        },
        {
            "id": 4,
            "price": 330000000,
            "brand": "Toyota",
            "type": "Yaris"
        },
        {
            "id": 5,
            "price": 249000000,
            "brand": "Toyota",
            "type": "Agya"
        },
        {
            "id": 6,
            "price": 149000000,
            "brand": "Daihatsu",
            "type": "Sigra"
        },
        {
            "id": 7,
            "price": 215000000,
            "brand": "Daihatsu",
            "type": "Xenia"
        },
    ]
}
```
## Club

### GET /club/leaguestandings

Response Body
```json
{
    "status": 200,
    "success": true,
    "message": "get data success",
    "payload": [
        {
            "id": 1,
            "clubname": "Manchester United",
            "point": 9
        },
        {
            "id": 3,
            "clubname": "Spurs",
            "point": 5
        },
        {
            "id": 2,
            "clubname": "Chelsea",
            "point": 1
        },
    ]
}
````

### GET /club/rank?clubname=Spurs

Response Body
```json
{
    "status": 200,
    "success": true,
    "message": "get data success",
    "payload": {
        "clubname": "Spurs",
        "standing": 5
    }
}
````

### POST /club/recordgame

Request Body
```json
{
    "clubhomename" : "Manchester United",
    "clubawayname" : "Spurs",
    "score": "0 : 0"
}
````

Response Body
```json
{
    "status": 200,
    "success": true,
    "message": "update data success"
}
````

## Letter

### POST /is-contain-letters

Request Body
```json
{
    "firstword": "kamu sedang apa",
    "lastword": "kamu sedang"
}
````

Response Body
```json
{
    true
}
````