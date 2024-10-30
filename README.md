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

 
## `POST /users/register`
Este endpoint crea el usuario 
telefono, edad, posicion, foto no es obligatorio
```JSON
[
  {
    "first_name": "Javier",
    "last_name": "Moreno",
    "email": "javier.moreno@example.com",
    "password": "javierPass321",
    "phone": "654987321",
    "profile_picture_url": "https://example.com/javier_pic.jpg",
    "role": "client",
    "position_player": "defensa",
    "age": 22
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
