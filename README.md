Trabajo de Lucas Guerrero.

## Solucion

Para este trabajo, se creo un struct "Result" en base a las estipulaciones del TP, funcionando como base para el parser.

Se inicio la funcion "parse(input string) (Result, error) declarando un Result vacio llamado "parsedInput". Debido a que seria imposible parsear el input correctamente si este tuviera menos de 5 caracteres, se chequea el largo de input antes de continuar con la operacion; de ser asi, se retora un error.

Luego se toman los primeros dos caracteres y se declaran como en el "Type" de parsedInput. Antes de declarar los siguientes caracteres estos son convertidos en int con la funcion "Atoi": Si retorna un error, quiere decir que el 3er y 4to caracter no representan numeros y como resultado, es imposible parsear el input. En este caso tambien se retorna un error, de lo contrario, el int retornado por Atoi se declara como el "Lenght" de parsedInput.

Por ultimo, todos los caracteres faltantes se declaran como el "Value" de parsedInput.

Si no se retorno ningun error en toda la operacion, se imprime y retorna parsedInput. 

## Testing

Para el testing se uso como base el ejemplo de un archivo test demostrado en el trabajo practico. Se hicieron clausulas de testing tomando en cuenta los datos de prueba entregados (por ejemplo, si el testData espera un error se testea por la presencia de un error. Si el testData tiene un value vacio, se testea por un value vacio).

Con los datos de prueba siguientes: 
```go
}{
	{"TX02AB", true, "TX", "AB", 2},
	{"NN100987654321", true, "NN", "0987654321", 10},
	{"TX06ABCDE", false, "", "", 0},
	{"NN04000A", false, "", "", 0},
}
```
Se obtuvo el siguiente resultado:
```
{Type:TX Value:AB Length:2}
{Type:NN Value:0987654321 Length:10}
{Type:TX Value:ABCDE Length:6}
{Type:NN Value:000A Length:4}
PASS
coverage: 83.3% of statements
ok      tudai2021.com   0.075s
```
 