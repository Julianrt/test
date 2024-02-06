# Testing Calculator #

Este es un repositorio sobre un test donde se agragan una serie de pruebas unitarias a un conjunto de métodos que contiene la estructura _Calculator_.

Se realizaron 10 pruebas por cada método de _Calcularor_ para poder analizar el comportamiento de estas, en diferentes casos, como por ejemplo, a veces escogiendo numeros donde _a_ era mayor a _b_, o viceversa, tambien probando con números iguales. Además utilizando numeros negativos.
Se hizo una consideranción especial con las pruebas del metodo _Divide_ ya que este, debe regresar un error cuando el valor de _b_ es 0, entonces se analiza que cuando se ejecuta y el valor de _b_ es 0, __si__ se está obteniendo el error deseado.
Se cambió el tipo de dato _int_ que maneja _Divide_ por _float64_ simplemente para tener números mas exactos al momento de realizar divisiones. 

#### Hecho en Go versión 1.21.1 ####

Para ejecutar las pruebas antes mencionada, es necesario clonar el proyecto y desde una terminal, en la raiz del proyecto, ejecutar el siguiente comando:
```
go test -v ./calculator/
```
