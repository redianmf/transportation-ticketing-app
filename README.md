# Transportation Ticketing App

This app intended to be used in mass transport system e.g. MRT, Bus, Train, etc. This app identifies user identity which in application could be parsed from QR Code that scanned when user passed through the gate

### Auth

This app uses JWT for authentication

#### Login

```http
POST /login
```

| Parameter  | Type     | Description  |
| :--------- | :------- | :----------- |
| `email`    | `string` | **Required** |
| `password` | `string` | **Required** |

#### Register

```http
POST /register
```

| Parameter    | Type     | Description       |
| :----------- | :------- | :---------------- |
| `email`      | `string` | **Required**      |
| `password`   | `string` | **Required**      |
| `username`   | `string` |                   |
| `birth_date` | `string` | ISO String format |

### Wallet

API for handle user wallet

#### Get User Wallet

```http
GET /wallet
```

#### Update Wallet

```http
PATCH /wallet
```

| Parameter | Type      | Description |
| :-------- | :-------- | :---------- |
| `amount`  | `integer` |             |

### Transportation Modes

API for handle transportation modes

#### Get All Transportation Modes

```http
GET /transportation-modes
```

#### Get Transportation Modes By Id

```http
GET /transportation-modes/:id
```

| URL Param | Type      | Description            |
| :-------- | :-------- | :--------------------- |
| `id`      | `integer` | Transportation mode id |

#### Insert New Transportation Modes

```http
POST /transportation-modes
```

| Parameter           | Type     | Description                                                                      |
| :------------------ | :------- | :------------------------------------------------------------------------------- |
| `name`              | `string` | Transportation mode name                                                         |
| `base_price`        | `int`    | Transportation mode fare                                                         |
| `additional_price`  | `int`    | Transportation mode additional price (used when price_calculation is not "FLAT") |
| `price_calculation` | `string` | "**FLAT**" OR "**POINT**"                                                        |

#### Update Transportation Modes

```http
PATCH /transportation-modes/:id
```

| URL Param | Type      | Description            |
| :-------- | :-------- | :--------------------- |
| `id`      | `integer` | Transportation mode id |

| Parameter           | Type     | Description                                                                      |
| :------------------ | :------- | :------------------------------------------------------------------------------- |
| `name`              | `string` | Transportation mode name                                                         |
| `base_price`        | `int`    | Transportation mode fare                                                         |
| `additional_price`  | `int`    | Transportation mode additional price (used when price_calculation is not "FLAT") |
| `price_calculation` | `string` | "**FLAT**" OR "**POINT**"                                                        |

#### Delete Transportation Modes

```http
DELETE /transportation-modes/:id
```

| URL Param | Type      | Description            |
| :-------- | :-------- | :--------------------- |
| `id`      | `integer` | Transportation mode id |

### Transaction

API for handle transaction

#### Transaction Gate (in and out)

```http
POST /transactions/:transportationModeId/:pointId
```

| URL Param              | Type      | Description                              |
| :--------------------- | :-------- | :--------------------------------------- |
| `transportationModeId` | `integer` | Transportation mode id                   |
| `pointId`              | `integer` | The bus stop, train station, etc. number |
