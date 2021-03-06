# Manipulación de estructuras (structs) en archivos binarios

### El propósito de este repositorio es poder ayudar a comprender como guardar, obtener y modificar estructuras luego de que han sido almacenadas en un archivo, ya que es fundamental para el curso de manejo e implementación de archivos tener claros estos conceptos. Por lo que en este repositorio se desarrolla el siguiente ejemplo.

------

------



#### **Enunciado**

Se tendrán 2 estructuras con los siguientes atributos:

### Profesor

| ATRIBUTO    | DESCRIPCIÓN |
| ----------- | ----------- |
| id_profesor | int         |
| CUI         | char[13]    |
| Nombre      | char[25]    |
| Curso       | char[25]    |

### Estudiante

| ATRIBUTO      | DESCRIPCIÓN |
| ------------- | ----------- |
| id_Estudiante | int         |
| CUI           | char[13]    |
| Nombre        | char[25]    |
| Carnet        | char[10]    |

Al iniciar la aplicación esta debe de crear el archivo binario en caso no exista, si el archivo
ha sido creado anteriormente no se debe de sobrescribir. La aplicación debe de contar con
un menú principal con las siguientes opciones:
1. Registro de Profesor
2. Registro de Estudiante
3. Ver Registros
4. Salir




1. Registro Profesor
     Si es seleccionada esta opción se le debe de solicitar por medio de consola el ingreso
       de cada uno de los datos que componen al struct Profesor, una vez realizado el ingreso
       de los datos el struct debe de escribirse en un archivo binario.
2. Registro Estudiante
     Si es seleccionada esta opción se le debe de solicitar por medio de consola el ingreso
       de cada uno de los datos que componen al struct Estudiante, una vez realizado el
       ingreso de los datos el struct debe de escribirse en un archivo binario.
3. Ver Registros
     Si es seleccionada esta opción se debe de leer el archivo binario y mostrar a discreción
       de cada uno todos aquellos registros de profesor y estudiantes registrados con
       anterioridad.
4. Salir
     Si es seleccionada esta opción se debe de finalizar la aplicación.



### Explicación de como se realizo este ejemplo:

**El ingreso de profesores y estudiantes tiene un orden aleatorio**, esto quiere decir que se puede ingresar un estudiante y luego un profesor o viceversa, también ingresar dos estructuras del mismo tipo seguido, en fin no se puede prever el comportamiento del ingreso de estas estructuras, p**or lo que es necesario tener estructuras auxiliares para saber en que posición del archivo (posición del puntero) se encuentra cada estructura**. Para ello se tiene esta imagen que es el esquema general de lo que se hizo (esquema general - figura 1).

------

![Esquema general - elaboracion propia](https://github.com/edinfusion/MIA_Tutorial/blob/8e695bc0668e7a327b93088fc10a24b90531f5ce/images/esquemageneral%20.png "Figura 1")  Figura 1

------

Como se puede observar en la imagen anterior, **se tienen 4 tipos de structs, de los cuales uno es de tipo profesor, otro de tipo estudiante y los dos restantes son estructuras auxiliares**. **La estructura con el nombre de "Siguiente", como se observa siempre que se guarda un profesor o estudiante se almacena en forma contigua esta estructura**, esto con el fin de poder saber en que parte del archivo (posición del cursor) esta almacenado el siguiente profesor o estudiante respectivamente **y así poder simular una lista enlazada de profesores y una lista enlazada de estudiantes** los cuales, como se  explicaba anteriormente no tienen por que estar de forma continua, se ingresan en orden aleatorio (ver nuevamente imagen de esquema general y ver las flechas para entender el funcionamiento de estas listas).

**La siguiente estructura auxiliar tiene el nombre de "Indice"** la cual como su nombre lo indica nos dará la dirección de posiciones clave a través de 5 atributos (ver figura 2). **Esta estructura siempre se guardara al inicio del archivo**, con la idea de que haga alusión a su nombre, c**uando aun no se han ingresado profesores o estudiantes 4 de los 5 atributos tienen valor de -1 y el atributo restante contendrá el valor numérico del tamaño del struct indice**, en este struct cada vez que se almacena un profesor o estudiante siempre se actualizan estos valores, como se explica mas adelante. (ver figura 3)

------

![Estructura Indice](https://github.com/edinfusion/MIA_Tutorial/blob/master/images/struct_indice%20.png "Figura 2") Figura 2

------

Para la comprensión de como se usan estas estructuras, a la hora de almacenar se tienen dos casos, uno cuando no existe ningún profesor/alumno y el segundo donde ya existe un profesor/alumno como se detalla, a continuación:

------

![imagen de caso 1](https://github.com/edinfusion/MIA_Tutorial/blob/104dcf3d3e509798ac302ad2d932a6468c916cbe/images/caso1.png)

------

![imagen de caso 2](https://github.com/edinfusion/MIA_Tutorial/blob/104dcf3d3e509798ac302ad2d932a6468c916cbe/images/caso2%20.png)

------

