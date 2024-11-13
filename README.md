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

## `POST /fields`
Sirve para insertar una nueva cancha: </br></br>
<b>Ejemplo de Request Body: </b>

```JSON
{
  "name": "Name",
  "address": "Adress",
  "neighborhood": "neighborhood",
  "phone": "097 777 777",
  "latitude": 0,
  "longitude": 0,
  "type": "5",
  "price": 1500,
  "description": "description",
  "logo_url": "https://exmaple.com/name.jpg",
  "services": [
    {
      "id": 1
    },
    {
      "id": 1
    }
  ],
  "creation_date": "2024-10-23",
  "photos": [
    "https://example.com/uy/wp-content/uploads/sites/2/2013/05/name-1.jpg",
    "https://example.com/uy/wp-content/uploads/sites/2/2013/05/name-2.jpg",
    "https://example.com/uy/wp-content/uploads/sites/2/2013/05/name-3.jpg"
  ],
  "available_days": [
    "1",
    "2",
    "3",
    "4",
    "5",
    "6",
    "7"
  ]
}
```

`Response: 201`

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
<b>Ejemplo de Request Body: </b>

```JSON
{
  "field_id": 1,
  "date": "2024-10-15",
  "start_time": "19:00",
  "end_time": "23:00",
  "user_id": 1,
  "status": "pending"
}
```

`Response: 201`

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


