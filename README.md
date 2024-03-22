 # Definición de Challenge
- Desarrollar un servicio que exponga una API REST para una lista de tareas (TODO) con endpoints para la creación, lectura, actualización y eliminación de tareas.

    - El almacenamiento de tareas debe ser en memoria o en alguna base de datos como SQLite.  
    - La salida de cada endpoint debe ser en formato JSON.

- Armar el Dockerfile

- Opcional: Documentación (OpenAPI), casos de prueba y colección de Postman.

 

Restricciones       
* No utilizar frameworks ni ORMs.     
* No utilizar bases de datos NoSQL


# Desarrollo
Se el proyecto enta desarrollado en la version `1.22.1` de Golang

## Ejecucion
Para ejecutar en entorno local utilizar el siguiente comando:
```sh
go run main.go
```
### Dockerfile

```sh
docker build -t vemo-sv .
docker run -p 8080:8080 vemo-sv:latest
```
### Test unitarios
```sh
go test ./... -v
```
### Postman
https://api.postman.com/collections/10090771-9a9c75e7-da72-4638-8eb5-8c0981ed7b6a?access_key=PMAT-01HSHB83TM7A5117QATJ9CE438

## Capas de software
Se definicion 3 capas de software:
- **Handler**: Donde se define cada manipulador de los difererentes path y metodos que se necesitan exponer para realizar el CRUD de ToDos
- **DataStore**: Capa encargada de dar almacenamiento a los ToDos, utilizando un **map[string] model.ToDo**
- **Model**: Capa de representación de datos y reglas los ToDos

> **Nota**: No se definicio una capa service ya que es una sola entidas en un futuro si se llega a agregar la entidad usuario por ejemplo estaria correcto agreagr esta capa de servicios 

