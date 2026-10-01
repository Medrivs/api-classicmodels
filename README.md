\# API REST en Go - Base de datos Classic Models



Este proyecto es una API REST construida en Go que se conecta a una base de datos MySQL local para exponer datos en formato JSON.



\## Archivos del proyecto

\* `main.go`: Contiene el código fuente de la API y la conexión a la base de datos.

\* `base\_de\_datos.sql`: Respaldo completo de la base de datos utilizada.

\* `go.mod` y `go.sum`: Archivos de configuración de dependencias de Go.



\## Historial de comandos utilizados en CMD (Paso a paso)



1\. \*\*Navegar a la carpeta de descargas e importar la base de datos original:\*\*

&#x20;  `cd "C:\\Users\\alnmd\\Documents\\Universidad\\Semestre 7\\Topicos de Despliegue\\mysqlsampledatabase"`

&#x20;  `"C:\\Program Files\\MySQL\\MySQL Server 8.0\\bin\\mysql.exe" -u root -p < mysqlsampledatabase.sql`



2\. \*\*Verificar instalación de Go:\*\*

&#x20;  `go version`



3\. \*\*Crear carpeta del proyecto y acceder a ella:\*\*

&#x20;  `mkdir api-classicmodels`

&#x20;  `cd api-classicmodels`



4\. \*\*Inicializar módulo de Go y descargar driver de MySQL:\*\*

&#x20;  `go mod init api-classicmodels`

&#x20;  `go get -u github.com/go-sql-driver/mysql`



5\. \*\*Exportar respaldo de la base de datos al proyecto:\*\*

&#x20;  `"C:\\Program Files\\MySQL\\MySQL Server 8.0\\bin\\mysqldump.exe" -u root -p classicmodels > base\_de\_datos.sql`



6\. \*\*Ejecutar la API:\*\*

&#x20;  `go run main.go`

