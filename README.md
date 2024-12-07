# Documentación de la API

## `GET /fields`
Este endpoint trae todas las canchas. EJ:

```JSON
[
  {
    "id": 1,
    "name": "Aguada Fútbol 5",
    "address": "Av. Gral. San Martín 2261",
    "neighborhood": "Aguada",
    "phone": "2201 0927",
    "latitude": 0,
    "longitude": 0,
    "type": "5",
    "price": 1500,
    "description": "Aguada Fútbol 5 es un complejo con dos canchas de césped artificial con caucho de última generación.Tiene las dimensiones reglamentarias (15,5 m X 30 m ).Escuela de fútbol de 5 a 13 años y espacio para que festejar tu cumpleaños o evento.El complejo cuenta con vestuarios, baños y duchas; bebidas frías.Excelente atención.",
    "logo_url": "https://canchea.com/uy/wp-content/uploads/sites/2/2013/05/aguada.png",
    "average_rating": 0,
    "services": [
      {
        "id": 0,
        "name": "Bebidas",
        "icon": ""
      },
      {
        "id": 0,
        "name": "Cantina",
        "icon": ""
      }
    ],
    "creation_date": "2024-10-23",
    "photos": [
      "https://canchea.com/uy/wp-content/uploads/sites/2/2013/05/Aguada-Fútbol-5-Montevideo-1.jpg",
      "https://canchea.com/uy/wp-content/uploads/sites/2/2013/05/Aguada-Fútbol-5-Montevideo.jpg"
    ],
    "available_days": [
      "1",
      "2",
      "3",
      "4",
      "5",
      "6",
      "7"
    ],
    "unvailable_dates": [
      {
        "FromDate": "2024-09-23 00:00:00",
        "ToDate": "2024-12-24 00:00:00"
      }
    ],
    "reservations": null
  },
  {
    "id": 2,
    "name": "Aguada Fútbol 5",
    "address": "Av. Gral. San Martín 2261",
    "neighborhood": "Aguada",
    "phone": "2201 0927",
    "latitude": 0,
    "longitude": 0,
    "type": "5",
    "price": 1500,
    "description": "Aguada Fútbol 5 es un complejo con dos canchas de césped artificial con caucho de última generación.Tiene las dimensiones reglamentarias (15,5 m X 30 m ).Escuela de fútbol de 5 a 13 años y espacio para que festejar tu cumpleaños o evento.El complejo cuenta con vestuarios, baños y duchas; bebidas frías.Excelente atención.",
    "logo_url": "https://canchea.com/uy/wp-content/uploads/sites/2/2013/05/aguada.png",
    "average_rating": 0,
    "services": [
      {
        "id": 0,
        "name": "Bebidas",
        "icon": ""
      },
      {
        "id": 0,
        "name": "Cantina",
        "icon": ""
      }
    ],
    "creation_date": "2024-10-23",
    "photos": [
      "https://canchea.com/uy/wp-content/uploads/sites/2/2013/05/Aguada-Fútbol-5-Montevideo-1.jpg",
      "https://canchea.com/uy/wp-content/uploads/sites/2/2013/05/Aguada-Fútbol-5-Montevideo.jpg",
      "https://canchea.com/uy/wp-content/uploads/sites/2/2013/05/Aguada-Fútbol-5-Montevideo1.jpg"
    ],
    "available_days": [
      "1",
      "2",
      "3",
      "4",
      "5",
      "6",
      "7"
    ],
    "unvailable_dates": null,
    "reservations": [
      {
        "Date": "2024-10-15",
        "StartTime": "19:00",
        "EndTime": "23:00"
      },
      {
        "Date": "2024-12-15",
        "StartTime": "19:00",
        "EndTime": "23:00"
      }
    ]
  },
  {
    "id": 3,
    "name": "Cancha Test",
    "address": "",
    "neighborhood": "",
    "phone": "",
    "latitude": 0,
    "longitude": 0,
    "type": "5",
    "price": 0,
    "description": "",
    "logo_url": "",
    "average_rating": 0,
    "services": [
      {
        "id": 0,
        "name": "Bebidas",
        "icon": ""
      },
      {
        "id": 0,
        "name": "Cantina",
        "icon": ""
      }
    ],
    "creation_date": "2024-10-24",
    "photos": [
      "photo1.jpg",
      "photo2.jpg"
    ],
    "available_days": [
      "1",
      "4",
      "7"
    ],
    "unvailable_dates": null,
    "reservations": null
  }
]
```
#### Parámetros

- `/fields?month=2024-10`: Trae todas las canchas, pero al traer las fechas no disponibles, y las reservas, solo traerá aquellas que sean posteriores al mes especificado. </br></br>
Esto permite que si una cancha tiene mil reservas, no se traerá las mil, sino solo aquellas que sean a futuro del mes especificado y así desahilitar esas fechas en el front. </br></br>
Por defecto, si no se especifíca nada, el `month` se setea en el mes actual.\n

- `/fields?limit=10`: Trae tantas canchas como las puestas en el parametro `limit`. </br></br>
Por defecto, esta seteado en 10. </br></br>
<i>NO es recomendable traes mas de 10 por performance</i> 

- `/fields?offset=10`: Sirve para paginar resultados. Es decir si pongo 10, esto no me traerá los primeros 10 resultados sino los que le siguen.</br></br>
Por defecto esta en 0

## `GET /fields/:id` 
Trae toda la info de una sola cancha. EJ: 

```JSON
{
  "id": 2,
  "name": "Aguada Fútbol 5",
  "address": "Av. Gral. San Martín 2261",
  "neighborhood": "Aguada",
  "phone": "2201 0927",
  "latitude": 0,
  "longitude": 0,
  "type": "5",
  "price": 1500,
  "description": "Aguada Fútbol 5 es un complejo con dos canchas de césped artificial con caucho de última generación.Tiene las dimensiones reglamentarias (15,5 m X 30 m ).Escuela de fútbol de 5 a 13 años y espacio para que festejar tu cumpleaños o evento.El complejo cuenta con vestuarios, baños y duchas; bebidas frías.Excelente atención.",
  "logo_url": "https://canchea.com/uy/wp-content/uploads/sites/2/2013/05/aguada.png",
  "average_rating": 0,
  "services": [
    {
      "id": 0,
      "name": "Bebidas",
      "icon": ""
    },
    {
      "id": 0,
      "name": "Cantina",
      "icon": ""
    }
  ],
  "creation_date": "2024-10-23",
  "photos": [
    "https://canchea.com/uy/wp-content/uploads/sites/2/2013/05/Aguada-Fútbol-5-Montevideo-1.jpg",
    "https://canchea.com/uy/wp-content/uploads/sites/2/2013/05/Aguada-Fútbol-5-Montevideo.jpg",
    "https://canchea.com/uy/wp-content/uploads/sites/2/2013/05/Aguada-Fútbol-5-Montevideo1.jpg"
  ],
  "available_days": [
    "1",
    "2",
    "3",
    "4",
    "5",
    "6",
    "7"
  ],
  "unvailable_dates": null,
  "reservations": [
    {
      "Date": "2024-10-15",
      "StartTime": "19:00",
      "EndTime": "23:00"
    },
    {
      "Date": "2024-12-15",
      "StartTime": "19:00",
      "EndTime": "23:00"
    }
  ]
}
```

#### Parámetros

-`/fields/:id?month=2024-10`: Al igual que la anterior traera las reservas y fechas no disponibles posteriores a lo especificado. </br></br>
Por defecto es siempre el mes actual

## Create Field API Endpoint

### Endpoint
`POST /api/fields`

### Authentication
- Requires Bearer Token in Authorization Header

### Request Parameters

#### Form Data Fields

| Parameter | Type | Required | Description | Example |
|-----------|------|----------|-------------|---------|
| `name` | string | Yes | Field name | `"TestFieldName"` |
| `user_id` | integer | Yes | User ID | `1` |
| `address` | string | No | address | `Calle 123` |
| `neighborhood` | string | No | neighborhood | `Barrio Peñarol💛​🖤​` |
| `phone` | string | No | phone | `097 777 777` |
| `latitude` | float64 | No | latitude | `-34.901112` |
| `longitude` | string | No | longitude | `-56.164532` |
| `type` | string | No | type of field | `5  \ 7 \ 11` |
| `price` | float64 | No | price per hour | `1200` |
| `description` | string | No | description | `A beautiful description` |
| `services[0].id` | integer | Optional | Service IDs | `1` |
| `available_days` | string[] | Optional | Available days | `"1"`, `"2"` |
| `creation_date` | string (YYYY-MM-DD) | Optional | Creation date | `"2024-12-02"` |
| `fieldImages` | file[] | Optional | Field images | Multiple image files |
| `unvailable_dates[0].fromDate` | string (YYYY-MM-DD) | Optional | Unavailable from date | `"2024-12-15"` |
| `unvailable_dates[0].toDate` | string (YYYY-MM-DD) | Optional | Unavailable to date | `"2024-12-20"` |

### CURL Example

```bash
curl -X POST \
  'http://localhost:8080/api/fields' \
  --header 'Authorization: Bearer YOUR_TOKEN_HERE' \
  --form 'name="Sports Complex"' \
  --form 'user_id=1' \
  --form 'services[0].id=1' \
  --form 'services[1].id=2' \
  --form 'available_days="1"' \
  --form 'available_days="2"' \
  --form 'creation_date="2024-12-02"' \
  --form 'unvailable_dates[0].fromDate="2024-12-15"' \
  --form 'unvailable_dates[0].toDate="2024-12-20"' \
  --form 'unvailable_dates[1].fromDate="2025-01-10"' \
  --form 'unvailable_dates[1].toDate="2025-01-15"' \
  --form 'fieldImages=@/path/to/image1.png' \
  --form 'fieldImages=@/path/to/image2.jpg'
```
### JavaScript (Fetch) Example

```js
const formData = new FormData();
formData.append('name', 'Sports Complex');
formData.append('user_id', 1);
formData.append('services[0].id', 1);
formData.append('services[1].id', 2);
formData.append('available_days', '1');
formData.append('available_days', '2');
formData.append('creation_date', '2024-12-02');
formData.append('unvailable_dates[0].fromDate', '2024-12-15');
formData.append('unvailable_dates[0].toDate', '2024-12-20');
formData.append('unvailable_dates[1].fromDate', '2025-01-10');
formData.append('unvailable_dates[1].toDate', '2025-01-15');
formData.append('fieldImages', imageFile1);
formData.append('fieldImages', imageFile2);


fetch('http://localhost:8080/api/fields', {
  method: 'POST',
  headers: {
    'Authorization': 'Bearer YOUR_TOKEN_HERE'
  },
  body: formData
});
```

#### Notes

- Multiple services can be added by incrementing index (`services[0].id, services[1].id`)
- Multiple unavailable dates can be added similarly
- Images can be added by repeating fieldImages parameter

## `GET /api/fields/per-owner`
Mandar token por authorization
```JSON
[
    {
        "field_name": "Field Name",
        "field_address": "123 Example Street",
        "field_type": "5",
        "field_phone": "123-456-7890",
        "field_status": true
    },
    {
        "field_name": "Field Name",
        "field_address": "123 Example Street",
        "field_type": "5",
        "field_phone": "123-456-7890",
        "field_status": true
    }
]
```

Te devuelve un ```Response 200```

## `UPDATE /api/fields/:id`
Mismo Body que Save

## `PATCH /api/fields/:id`
El body traera SOLO los campos que se quieran actualizar
 
## `POST /users/register`
Este endpoint crea el usuario 
phone, age, profile_picture_url no es obligatorio, pero en "role" siempre hay que poner uno de los 3 valores: admin, client, field
para la password tiene minímo 8 caracteres
```JSON
[
  {
    "first_name": "Javier",
    "last_name": "Moreno",
    "email": "javier.moreno@example.com",
    "password": "javierPass321",
    "phone": "654987321",
    "role": "client",
    "accepted_terms": true
  }
]
```

Retorna un Json con el usuario y un ```Response: 201```

## `GET /users/check-info`
El ID del usuario lo obtiene solo gracias al token que hay que mandarle en Authorization en el Header
No hay que mandarle nada en el body
Retorna un 
```JSON
{
    "id": 16,
    "first_name": "Juan",
    "last_name": "Pérezzila",
    "email": "simonpintos771@gmail.com",
    "phone": "+123456789",
    "profile_picture_url": "https://example.com/profile/juan.jpg",
    "role": "client",
    "position_player": "https://example.com/profile/juan.jpg",
    "age": 25,
    "isVerified": false
}
```

```Response: 200```

## `POST /users/login`
Este endpoint te loguea con esta entrada
```JSON
{
  "email": "javier.moreno@example.com",
  "password": "javierPass321"
}
```

Retorna un json con el usuario y un ```Response: 200``` 

## `GET /users/auth/token`
Este endpoint te devuelve un:
```JSON
{
    "message": "Ok"
} 
```
y ```Response: 200``` si tu token no ha expirado aún o un:
{
  Message: "El token ha expirado"
}
y ```Response: 401``` si el token ya expiró

## Explicación de la lógica detrás de 'forgot password'
Primero, hay que hacerle un POST a la siguiente ruta con solamente el email del usuario, lo cual activará la función de enviarle 
al usuario el mail que contiene el token y la url para reestablecer su contraseña. El usuario presionará en el botón que le redirigirá a
la ruta con método PUT (las instrucciones para esta ruta están colocadas después del método POST) y ahí pondrá su email, token y su nueva contraseña.
## `POST /users/password-recovery` 
```JSON
{
    "email": "usuario@example.com"
} 
``` 
Al hacer este post te devolverá un ```Response: 200``` acompañado de un 
```JSON
{
    "message": "Recovery email sent"
}
```

## `PUT /users/reset-password` 
```JSON
{
  "email": "usuario@example.com",
  "token": "064126",
  "new_password": "elejemplo12345"
}
``` 
Al hacer este put te devolverá un ```Response: 200``` acompañado de un 
```JSON
{
    "message": "Password successfully updated"
}
```

## `PUT /users/update-user-profile`
Este es el template para hacer el PUT, el ID se saca del token, pero es importante que se envíen todos los campos

```JSON
{
  "first_name": "Juan",
  "last_name": "Pérez",
  "email": "example@gmail.com",
  "phone": "+123456789",
  "position_player": "forward",
  "team_name": "Los Guerreros",
  "age": 25,
  "profile_picture_url": "https://example.com/profile/juan.jpg",
  "id": 8
}
```

Devuelve un status 200 con
```JSON
{
  "message": "User updated successfully"
}
```

# Explicación de la lógica detrás de verify user account
Primero hay que enviar el correo, esto se hace haciéndole un post a la siguiente URL con el email del usuario.

## `POST /users/verify-user-email` 
```JSON
{
  "email":"example@gmail.com"
}
```
El cual va a retornar un JSON con un código de 200 OK
```JSON
{
    "message": "Recovery email sent"
}
```
Después de que llegue el mail, apretas en el botón de "Verificar cuenta" y te va a llevar a la url http://localhost:8080/api/users/verify?token=034527

Al darle al botón en el mail te va a hacer un GET a la ruta: localhost:8080/api/users/verify?token=034527 Y automáticamente el usuario va a ser verificado y ya va a estar todo listo.


## `POST /users/add-favourites`
No olvidarse de mandar el token por Authorization
Para agregar a favoritos una cancha hay que enviar este JSON al endpoint
```JSON
{
    "field_id": 12
}
```

Lo cual retorna un ```Response: 200```

## `GET /users/favourites-per-user`
No olvidarse de mandar el token por Authorization
Retorna 
```JSON
[
    {
        "field_name": "Aguada Fútbol 5",
        "field_address": "Av. Gral. San Martín 2261",
        "field_phone": "2201 0927",
        "field_logo_url": "https://canchea.com/uy/wp-content/uploads/sites/2/2013/05/aguada.png"
    },
    {
        "field_name": "Cancha Test",
        "field_address": "",
        "field_phone": "",
        "field_logo_url": ""
    },
    {
        "field_name": "0 Stress",
        "field_address": "Guaviyu 3013 esq. Gral. Flores",
        "field_phone": "2711 1332",
        "field_logo_url": ""
    },
    {
        "field_name": "Campo Grande",
        "field_address": "Av. Gral. Garibaldi 1892",
        "field_phone": "2200 3129",
        "field_logo_url": "https://canchea.com/uy/wp-content/uploads/sites/2/2013/11/campogrande.png"
    }
]
``` 

con un ```Response: 200```

# Documentación de la API de Reservas

## `GET /reservations`
Trae todas las reservas. EJ:

```JSON
[
  {
    "id": 1,
    "field_id": 1,
    "date": "2024-10-15",
    "start_time": "19:00",
    "end_time": "23:00",
    "user_id": 1,
    "status": "pending"
  },
  {
    "id": 2,
    "field_id": 1,
    "date": "2024-10-15",
    "start_time": "19:00",
    "end_time": "23:00",
    "user_id": 1,
    "status": "pending"
  },
  {
    "id": 3,
    "field_id": 1,
    "date": "2024-10-15",
    "start_time": "19:00",
    "end_time": "23:00",
    "user_id": 1,
    "status": "pending"
  }
]
```

## `GET /reservations/:id`
Trae una reserva en particular. EJ:

```JSON
{
  "id": 1,
  "field_id": 1,
  "date": "2024-10-15",
  "start_time": "19:00",
  "end_time": "23:00",
  "user_id": 1,
  "status": "pending"
}
```

## `POST /reservations`
Sirve para insertar una nueva reserva: </br></br>
Advertencia: el usuario debe tener el rol de cliente, debe estar con sesión iniciada y con el Authorization Bearer {token}, sino, no te va a funcionar
aunque le pagues
<b>Ejemplo de Request Body: </b>

```JSON
{  
  "field_id": 3, 
 "date": "2024-10-15",  
 "start_time": "19:00:00",  
 "end_time": "23:00:00"
}

```

`Response: 201`



## `GET /reservations/reservations-per-owner?MonthsAgo=1&Hour=14&format=csv`
/reservations/reservations-per-owner?MonthsAgo=1&format=csv
/reservations/reservations-per-owner?Hour=14&format=csv
/reservations/reservations-per-owner?format=csv
/reservations/reservations-per-owner
La idea es que la url vaya cambiando, le sacas y le pones los params
Explicación: cuando se quiera filtrar por MonthsAgo se le envía a partir del 1 en adelante, esto hace que filtre si le mando un 3, de los últimos 3 meses, si le mando Hour filtrará a partir de la hora 14 por ejemplo, eso me va a traer todas las reservas hechas a las 14. 
Si no hay ningún filtro que aplicar, se usará solo la ruta /reservations/reservations-per-owner, los filtros se añaden con forme se le añaden queryParams.
Ahora es importante que cuando se haga la consulta en la tabla sea haga un get sin el Params format, ya que este está hecho para descargar ya sea csv o pdf en minúscula.
`Response: 200`


## `GET /reservations/reservations-per-user`
No olvidarse de enviar el token por Authorization: Bearer {token}
Devuelve un array con todas las reservas del usuario y su estado
```JSON
[
    {
        "EmailUser": "simonpintos771@gmail.com",
        "ReservationDate": "2024-10-15T00:00:00Z",
        "StartTime": "19:00:00",
        "EndTime": "23:00:00",
        "FieldName": "Cancha Test",
        "StatusReservation": "reserved"
    },
    {
        "EmailUser": "simonpintos771@gmail.com",
        "ReservationDate": "2024-11-30T00:00:00Z",
        "StartTime": "14:00:00",
        "EndTime": "18:00:00",
        "FieldName": "2 a 1 Fútbol 5",
        "StatusReservation": "reserved"
    }
]
```

`Response: 200`

## `PUT /reservations/:id`
Sirve para actualizar una reserva: </br></br>
<b>Ejemplo de Request Body: </b>

```JSON
{
  "status": "approved"
}
```

`Response: 200`

## `DELETE /reservations/:id`
Sirve para eliminar una reserva: </br></br>

`Response: 204`


