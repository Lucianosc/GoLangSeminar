Trabajo Practico GoLang

Integrantes: Luciano Scaminacci, Rosario Perrotta

Consigna: El trabajo practico consta de hacer una función que dado una cadena de caracteres devuelva una estructura. 

Para cumplir con la consigna, se creo dentro del package model, la estructura Result con los atributos indicados en la consigna de Type, Length y Value, con su respectivo NewResult que lo que hace es generar una nueva instancia. 
En "main.go" se crea la funcionalidad requerida para que dada una cadena de caracteres subdividimos los distintos datos tomando los primeros 2 posiciones de la cadena para el tipo, los siguientes dos para el length y los restantes para el value. 

Luego, previo a la realizacion de un conjunto de validaciones, devuelve una instancia de Result. Esta tendra valores seteados en funcion de si cumple o no con las validaciones definidas. 

Las validaciones fueron en primer lugar chequear que el largo del value sea igual al largo definido en el length. Si da como resultado que es verdadero, verifica si es de tipo "NN" o "TX", si se encuentra dentro de estas opciones entonces valida en funcion de cada tipo, si la cadena posee solo caracteres numericos llamando a la funcion IsNumber o solo caracteres alfabeticos en el caso de TX llamando a la funcion IsLetter. 

Si todas estas validaciones dan como resultado true, crea una nueva instancia pasando el type, length y value de la cadena por parametro. Caso contrario cada validacion retornara una instancia "vacia" junto con un mensaje de error indicativo de donde fue que fallo.

En el archivo main_test.go se realizaron los tests para comprobar el correcto funcionamiento de la logica planteada. Se generaron el archivo out y el out.html para validar qué parte del código fue cubierta por el test con el comando go test ./... -count=1 -coverprofile=out.test y go tool cover -html out -o out.html. 
Pudimos verificar que el 95.0 % del código generado fue cubierto por el test.