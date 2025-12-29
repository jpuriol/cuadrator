# Cuadrator

Cuadrator es una herramienta escrita en Go para gestionar y generar cuadrantes de turnos de forma sencilla. Permite definir participantes, turnos y ocupaciones mediante archivos YAML y generar un documento PDF con el resultado.

## Comandos Disponibles

La herramienta ofrece varios comandos para interactuar con los datos del cuadrante:

- `check`: Verifica si la información contenida en los archivos YAML es válida y consistente.
- `pdf`: Genera un documento PDF con la información del cuadrante.
- `show`: Muestra la ocupación que tiene un participante en un turno específico.
- `stats`: Muestra estadísticas sobre cuántos participantes hay por cada turno.

Para ver la ayuda detallada de cualquier comando, puedes ejecutar:
```bash
./cuadrator [comando] --help
```

## Configuración

El proyecto utiliza tres archivos principales de configuración en formato YAML:

### `participants.yaml`
Contiene la lista de nombres de los participantes que pueden ser asignados al cuadrante.

### `schema.yaml`
Define la estructura del cuadrante:
- **name**: Nombre del cuadrante (ej. "Marzo 2023").
- **shifts**: Definición de los horarios de los turnos.
- **occupations**: Lista de las diferentes tareas u ocupaciones disponibles.

### `quadrant.yaml`
Contiene la asignación real de los participantes a los turnos y ocupaciones definidos en el esquema.

## Instalación y Uso

Si tienes Go instalado, puedes compilar el proyecto ejecutando:

```bash
go build -o cuadrator ./cmd/cuadrator/main.go
```

Luego, puedes ejecutar los comandos mencionados anteriormente:

```bash
./cuadrator check
./cuadrator pdf
```

## Autor
Desarrollado por jpuriol.
