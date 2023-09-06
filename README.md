# Pega

O proxecto Pega ten o obxectivo de extraer e amosar datos electorais.

## Ferramentas

### inebase

Importa datos base do INE á base de datos.

```
inebase [OPCIONS] CONSUNTO_DATOS FICHEIRO
```

`CONXUNTO_DATOS`: Conxunto de datos a importar. Valores soportados: concellos.

`FICHEIRO`: Ruta ao ficheiro cos datos INE.

#### Conxuntos de datos soportados

##### Concellos

Ficheiro xslx co listado de concellos do INE.

### infoelectoral

Importa datos da web infoelectoral do Ministerio de Interior á base de datos.

```
infoelectoral [OPCIONS] FICHEIRO
```

`FICHEIRO`: Ruta ao ficheiro cos datos electorais.
