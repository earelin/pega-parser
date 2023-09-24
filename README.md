# Pega

O proxecto Pega ten o obxectivo de extraer e amosar datos electorais.

Implementa una API REST para consular los datos de resultados
electorais almacenados no sistema.

## API

### Datos xerais

#### GET `/comunidades-autonomas`

*Valor Retornado*

```
[
  {
    id: long
    nome: string
  }
]
```

#### GET `/comunidade-autonoma/{id}/provincias`

*Valor Retornado*

```
[
  {
    id: long
    nome: string
  }
]
```

#### GET `/provincias`

*Valor Retornado*

```
[
  {
    id: long
    nome: string
  }
]
```

#### GET `/provincia/{id}/concellos`

*Valor Retornado*

```
[
  {
    id: long
    nome: string
  }
]
```

#### GET `/concellos/pescuda/{search}`

*Valor Retornado*

```
[
  {
    id: long
    nome: string
  }
]
```

### Datos de Procesos Electorais

#### GET `/procesos-electorais`

Lista de procesos electorais.

*Parámetros*

`tipo`: Tipo de proceso electoral:
```
 1: Referéndum
 2: Xerais
 3: Congreso
 4: Senado
 5: Municipais
 6: Autonómicas
 7: Cabildos Insulares
10: Parlamento Europeu
15: Partidos Xudiciais e Diputacións Provinciais
```
`ambito`: Ámbito territorial do proceso electoral.
Código INE.

*Valor Retornado*

```
[
  {
    id: long
    tipo: int
    ambito: int
    data: iso string
  }
]
```

#### GET `/procesos-electorais/{id}/datos-xerais`

Datos xerais do proceso electoral.

*Valor Retornado*

```
{
  censoIne: number
  censoCera: number
  primeiroAvanceParticipacion: iso string
  segundoAvanceParticipacion: iso string
}
```

#### GET `/procesos-electorais/{id}/resultados`

Resultados xerais do proceso electoral.

*Valor Retornado*

```
{
  votantesPrimeiroAvanceParticipacion: number;
  votantesSegundoAvanceParticipacion: number;
  votantesCere: number;
  votosEnBranco: number;
  votosNulos: number;
  votosACandaturas: number;
  votosPorCandidatura: [
    {
      candidatura: Candidatura;
      representantesEleitos: number;
      votos: number;
    }
  ];
}
```

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
