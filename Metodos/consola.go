package Metodos

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"tutorial/Estructuras"
)

func crearArchivo() {
	//para saber si existe:
	archivo, err := os.OpenFile("disco.bin", os.O_RDWR, 0777)
	//si no existe:
	if err != nil {
		//creacion de archivo
		archivo, err := os.OpenFile("disco.bin", os.O_CREATE, 0777)
		if err != nil {
			fmt.Println(err)
		}
		/*
			creacion de estructura indice, esta es la primera estructura que se guarda
			debido a que como su mismo nombre lo dice sera un indice para saber la direccion
			del puntero de maestros y alumnos y la posicion vacía en donde se puede almacenar uno nuevo:
		*/
		var indice Estructuras.Indice
		indice.P_primeralumno = -1
		indice.P_primerprofesor = -1
		indice.P_ultimoalumno = -1
		indice.P_ultimoprofesor = -1
		indice.P_libre = int32(binary.Size(indice))
		/*
			ya creado el archivo
			se procede a abrir para guardar el struct indice
		*/
		archivo2, err2 := os.OpenFile("disco.bin", os.O_RDWR, 0777)
		if err2 != nil {
			fmt.Println(err)
		}
		archivo2.Seek(0, io.SeekStart)
		//en golang para escribir un struct en un archivo se debe guardar previamente en un buffer para la conversion a binary
		var buff bytes.Buffer
		binary.Write(&buff, binary.BigEndian, &indice)
		//este metodo se usa para guardar el buffer en el archivo binario
		escribirBytes(archivo2, buff.Bytes())
		defer archivo2.Close()
		defer archivo.Close()
	}
	defer archivo.Close()

}

func Ejecutar() {
	crearArchivo()
	menu()
}

// ****************************************************************************
// ***************************** CREAR PROFESOR *******************************
// ****************************************************************************

func crearprofesor() {
	fmt.Println(" **** REGISTRO DE PROFESOR **** ")
	//creacion de struct profesor
	var profe Estructuras.Profesor
	//obtencion de datos para profesor
	fmt.Print("ID_PROFESOR: ")
	fmt.Scanf("%d", &profe.Id_profesor) //aqui se puede agregar esto: n, err := fmt.Scanf("%d", &number) para validar el tipo
	fmt.Println("")

	fmt.Print("CUI: ")
	var cui string
	fmt.Scanf("%s", &cui)
	copy(profe.Cui[:], cui)
	fmt.Println("")

	fmt.Print("NOMBRE: ")
	var nombre string
	fmt.Scanf("%s", &nombre)
	copy(profe.Nombre[:], []byte(nombre))
	fmt.Println("")

	fmt.Print("CURSO: ")
	var curso string
	fmt.Scanf("%s", &curso)
	copy(profe.Curso[:], []byte(curso))
	fmt.Println("")
	//se procede a abrir archivo en donde se guardan los alumnos y profesores
	archivo, err := os.OpenFile("disco.bin", os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}

	//muevo puntero al inicio del archivo para obtener el struct indice
	archivo.Seek(0, io.SeekStart)
	//se obtiene struct indice de archivo
	indice := getIndice(archivo, err)
	//valido si ya existe un profesor
	if indice.P_primerprofesor == -1 {
		// si no existe entonces utilizo P_libre, esta posicion es la que se almacena al crear el archivo
		// por lo que esta inmediatamente despues de la estructura indice y no sobreescribe
		// por lo tanto en esa posicion se almacena el primer profesor / alumno
		indice.P_primerprofesor = indice.P_libre

	} else {
		/*
			en caso contrario que ya existe un profesor
			se debe actualizar el struct siguiente
			colocando como posicion donde se guardara
			el profesor
		*/

		archivo.Seek(int64(indice.P_ultimoprofesor), io.SeekStart)
		siguiente := getNext(archivo, err)
		//actualizo apuntador, el valor que se actualiza sera la posicion del cursor
		//en donde se almacenara el nuevo profesor
		siguiente.Siguiente = indice.P_libre
		//almaceno struct actualizado:
		//para ello muevo el cursor nuevamente a la posicion de ultimo profesor que es donde esta el struct siguiente
		archivo.Seek(int64(indice.P_ultimoprofesor), io.SeekStart)
		//en golang para escribir un struct en un archivo se debe guardar previamente en un buffer para la conversion a binary
		var buff bytes.Buffer
		binary.Write(&buff, binary.BigEndian, &siguiente)
		//este metodo se usa para guardar el buffer en el archivo binario
		escribirBytes(archivo, buff.Bytes())
	}

	//muevo puntero a posicion libre para almacenar struct de profesor
	archivo.Seek(int64(indice.P_libre), io.SeekStart)
	//almaceno nuevo profesor en archivo
	var buff bytes.Buffer
	binary.Write(&buff, binary.BigEndian, &profe)
	//este metodo se usa para guardar el buffer en el archivo binario
	escribirBytes(archivo, buff.Bytes())
	//estructura que debe ir siempre contigua al nuevo profesor, la cual almacenara el valor de la posicion del siguiente profesor o estudiante
	var siguienteprof Estructuras.Apuntador
	//el valor es -1 debido a que en este momento es el ultimo profesor ingresado
	siguienteprof.Siguiente = -1
	//Se procede a guardar el struct apuntador (siguiente) y para ello se debe colocar proximo
	//a el profesor que se acaba de almacenar
	//el calculo seria: plibre+tamañodeprofesor
	indice.P_ultimoprofesor = indice.P_libre + int32(binary.Size(profe))
	archivo.Seek(int64(indice.P_ultimoprofesor), io.SeekStart)
	//se procede a almacenar
	var buff2 bytes.Buffer
	binary.Write(&buff2, binary.BigEndian, &siguienteprof)
	//este metodo se usa para guardar el buffer en el archivo binario
	escribirBytes(archivo, buff2.Bytes())
	//lo ultimo sera actualizar la posicion libre y guardar el struct actualizado
	indice.P_libre += int32(binary.Size(profe)) + int32(binary.Size(siguienteprof))
	//se mueve cursor al inicio, ya que alli estara siempre el struct indice
	archivo.Seek(0, io.SeekStart)
	var buff3 bytes.Buffer
	binary.Write(&buff3, binary.BigEndian, &indice)
	//este metodo se usa para guardar el buffer en el archivo binario
	escribirBytes(archivo, buff3.Bytes())
	fmt.Println("SE HA INGRESADO UN NUEVO PROFESOR, CON EXITO")
	defer archivo.Close()

}

// ****************************************************************************
// *************************** CREAR ESTUDIANTE *******************************
// ****************************************************************************

func crearEstudiante() {
	fmt.Println(" **** REGISTRO DE ESTUDIANTE **** ")
	//creacion de struct profesor
	var alumno Estructuras.Estudiante
	//obtencion de datos para profesor
	fmt.Print("ID_ESTUDIANTE: ")
	fmt.Scanf("%d", &alumno.Id_estudiante) //aqui se puede agregar esto: n, err := fmt.Scanf("%d", &number) para validar el tipo
	fmt.Println("")

	var cui string
	fmt.Print("CUI: ")
	fmt.Scanf("%s", &cui)
	copy(alumno.Cui[:], []byte(cui))
	fmt.Println("")

	var nombre string
	fmt.Print("NOMBRE: ")
	fmt.Scanf("%s", &nombre)
	copy(alumno.Nombre[:], []byte(nombre))
	fmt.Println("")

	var carnet string
	fmt.Print("CARNET: ")
	fmt.Scanf("%s", &carnet)
	copy(alumno.Carnet[:], []byte(carnet))
	fmt.Println("")
	//se procede a abrir archivo en donde se guardan los alumnos y profesores
	archivo, err := os.OpenFile("disco.bin", os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}

	//muevo puntero al inicio del archivo para obtener el struct indice
	archivo.Seek(0, io.SeekStart)
	//se obtiene struct indice de archivo
	indice := getIndice(archivo, err)
	//valido si ya existe un profesor
	if indice.P_primeralumno == -1 {
		// si no existe entonces utilizo P_libre, esta posicion es la que se almacena al crear el archivo
		// por lo que esta inmediatamente despues de la estructura indice y no sobreescribe
		// por lo tanto en esa posicion se almacena el primer profesor / alumno
		indice.P_primeralumno = indice.P_libre

	} else {
		/*
			en caso contrario que ya existe un profesor
			se debe actualizar el struct siguiente
			colocando como posicion donde se guardara
			el profesor
		*/

		archivo.Seek(int64(indice.P_ultimoalumno), io.SeekStart)
		siguiente := getNext(archivo, err)
		//actualizo apuntador, el valor que se actualiza sera la posicion del cursor
		//en donde se almacenara el nuevo profesor
		siguiente.Siguiente = indice.P_libre
		//almaceno struct actualizado:
		//para ello muevo el cursor nuevamente a la posicion de ultimo profesor que es donde esta el struct siguiente
		archivo.Seek(int64(indice.P_ultimoalumno), io.SeekStart)
		//en golang para escribir un struct en un archivo se debe guardar previamente en un buffer para la conversion a binary
		var buff bytes.Buffer
		binary.Write(&buff, binary.BigEndian, &siguiente)
		//este metodo se usa para guardar el buffer en el archivo binario
		escribirBytes(archivo, buff.Bytes())
	}

	//muevo puntero a posicion libre para almacenar struct de alumno
	archivo.Seek(int64(indice.P_libre), io.SeekStart)
	//almaceno nuevo alumno en archivo
	var buff bytes.Buffer
	binary.Write(&buff, binary.BigEndian, &alumno)
	//este metodo se usa para guardar el buffer en el archivo binario
	escribirBytes(archivo, buff.Bytes())
	//estructura que debe ir siempre contigua al nuevo alumno, la cual almacenara el valor de la posicion del siguiente alumno o estudiante
	var siguienteest Estructuras.Apuntador
	//el valor es -1 debido a que en este momento es el ultimo alumno ingresado
	siguienteest.Siguiente = -1
	//Se procede a guardar el strucscanf("%i", &opmenu);t apuntador (siguiente) y para ello se debe colocar proximo
	//a el alumno que se acaba de almacenar
	//el calculo seria: plibre+tamañodealumno
	indice.P_ultimoalumno = indice.P_libre + int32(binary.Size(alumno))
	archivo.Seek(int64(indice.P_ultimoalumno), io.SeekStart)
	//se procede a almacenar
	var buff2 bytes.Buffer
	binary.Write(&buff2, binary.BigEndian, &siguienteest)
	//este metodo se usa para guardar el buffer en el archivo binario
	escribirBytes(archivo, buff2.Bytes())
	//lo ultimo sera actualizar la posicion libre y guardar el struct actualizado
	indice.P_libre += int32(binary.Size(alumno)) + int32(binary.Size(siguienteest))
	//se mueve cursor al inicio, ya que alli estara siempre el struct indice
	archivo.Seek(0, io.SeekStart)
	var buff3 bytes.Buffer
	binary.Write(&buff3, binary.BigEndian, &indice)
	//este metodo se usa para guardar el buffer en el archivo binario
	escribirBytes(archivo, buff3.Bytes())
	fmt.Println("SE HA INGRESADO UN NUEVO ALUMNO, CON EXITO")
	defer archivo.Close()

}

// ****************************************************************************
// ******************************* MENU ***************************************
// ****************************************************************************
func menu() {
	var opmenu int = 0
	for opmenu != 4 {
		switch opmenu {
		case 0:
			fmt.Println(" ****** MENU ****** ")
			fmt.Println("1. REGISTRO DE PROFESOR  ")
			fmt.Println("2. REGISTRO DE ESTUDIANTE")
			fmt.Println("3. VER REGISTROS         ")
			fmt.Println("4. SALIR                 ")
			fmt.Println("\n\n INGRESE UNA OPCION: ")
			fmt.Scanf("%d", &opmenu)
		case 1:
			crearprofesor()
			opmenu = 0
		case 2:
			crearEstudiante()
			opmenu = 0
		case 3:
			verRegistros()
			opmenu = 0
		case 4:
			opmenu = 4
		default:
			fmt.Println("INGRESE UNA OPCION CORRECTA!!!")
			opmenu = 0
		}

	}
}

func verRegistros() {

	//leo archivo
	archivo, err := os.OpenFile("disco.bin", os.O_RDWR, 0777)
	//si no existe:
	if err != nil {
		fmt.Println(err)
	}
	//se mueve el puntero al inicio para poder leer struct indice
	archivo.Seek(0, io.SeekStart)
	indice := getIndice(archivo, err)
	// PROFESORES
	if indice.P_primerprofesor != -1 {
		//leo el primer profesor
		//pare ello se mueve el puntero hacia el primer profesor
		archivo.Seek(int64(indice.P_primerprofesor), io.SeekStart)
		//se lee struct
		profe := getProfe(archivo, err)

		//profesores
		fmt.Println(" ****** PROFESORES ****** ")
		fmt.Println("     ID: ", profe.Id_profesor)
		fmt.Println("    CUI: ", string(profe.Cui[:]))
		fmt.Println(" NOMBRE: ", string(profe.Nombre[:]))
		fmt.Println("  CURSO: ", string(profe.Curso[:]))
		fmt.Println("")

		//ahora empiezo a recorrer siguientes en busca de mas profesores, si no existen el siguiente sera -1
		archivo.Seek(int64(indice.P_primerprofesor)+int64(binary.Size(profe)), io.SeekStart)
		siguiente := getNext(archivo, err)

		for siguiente.Siguiente != -1 {
			archivo.Seek(int64(siguiente.Siguiente), io.SeekStart)
			profe := getProfe(archivo, err)
			fmt.Println(" ****** PROFESORES ****** ")
			fmt.Println("     ID: ", profe.Id_profesor)
			fmt.Println("    CUI: ", string(profe.Cui[:]))
			fmt.Println(" NOMBRE: ", string(profe.Nombre[:]))
			fmt.Println("  CURSO: ", string(profe.Curso[:]))
			fmt.Println("")
			archivo.Seek(int64(siguiente.Siguiente)+int64(binary.Size(profe)), io.SeekStart)
			siguiente = getNext(archivo, err)
		}
	} else {
		fmt.Println("AUN NO SE HAN INGRESADO PROFESORES")
	}

	// ESTUDIANTES
	if indice.P_primeralumno != -1 {
		archivo.Seek(int64(indice.P_primeralumno), io.SeekStart)
		alum := getEstudiante(archivo, err)

		//profesores
		fmt.Println(" ****** ESTUDIANTES ****** ")
		fmt.Println("     ID: ", alum.Id_estudiante)
		fmt.Println("    CUI: ", string(alum.Cui[:]))
		fmt.Println(" NOMBRE: ", string(alum.Nombre[:]))
		fmt.Println(" CARNET: ", string(alum.Carnet[:]))
		fmt.Println("")

		//ahora empiezo a recorrer siguientes en busca de mas profesores, si no existen el siguiente sera -1
		archivo.Seek(int64(indice.P_primeralumno)+int64(binary.Size(alum)), io.SeekStart)
		siguiente := getNext(archivo, err)

		for siguiente.Siguiente != -1 {
			archivo.Seek(int64(siguiente.Siguiente), io.SeekStart)
			alum := getEstudiante(archivo, err)
			fmt.Println(" ****** ESTUDIANTES ****** ")
			fmt.Println("     ID: ", alum.Id_estudiante)
			fmt.Println("    CUI: ", string(alum.Cui[:]))
			fmt.Println(" NOMBRE: ", string(alum.Nombre[:]))
			fmt.Println(" CARNET: ", string(alum.Carnet[:]))
			fmt.Println("")
			archivo.Seek(int64(siguiente.Siguiente)+int64(binary.Size(alum)), io.SeekStart)
			siguiente = getNext(archivo, err)
		}
	} else {
		fmt.Println("AUN NO SE HAN INGRESADO ESTUDIANTES")
	}
}

// *************************************************************************************************
// **************************************  METODOS AUXILIARES  *************************************
// *************************************************************************************************

//retorna struct de indice en archivo especificado
func getIndice(archivox *os.File, err error) Estructuras.Indice {
	var ind Estructuras.Indice
	tamindice := binary.Size(ind)
	info := leerBytes(archivox, tamindice)
	buff := bytes.NewBuffer(info)
	err = binary.Read(buff, binary.BigEndian, &ind)
	if err != nil {
		fmt.Println("lectura de binarios fallo: ", err)
	}
	return ind
}

//retorna struct apuntador
func getNext(archivox *os.File, err error) Estructuras.Apuntador {
	var sig Estructuras.Apuntador
	tamindice := binary.Size(sig)
	info := leerBytes(archivox, tamindice)
	buff := bytes.NewBuffer(info)
	err = binary.Read(buff, binary.BigEndian, &sig)
	if err != nil {
		fmt.Println("lectura de binarios fallo: ", err)
	}
	return sig
}

//retorna struct profesor
func getProfe(archivox *os.File, err error) Estructuras.Profesor {
	var sig Estructuras.Profesor
	tamindice := binary.Size(sig)
	info := leerBytes(archivox, tamindice)
	buff := bytes.NewBuffer(info)
	err = binary.Read(buff, binary.BigEndian, &sig)
	if err != nil {
		fmt.Println("lectura de binarios fallo: ", err)
	}
	return sig
}

//retorna struct estudiante
func getEstudiante(archivox *os.File, err error) Estructuras.Estudiante {
	var sig Estructuras.Estudiante
	tamindice := binary.Size(sig)
	info := leerBytes(archivox, tamindice)
	buff := bytes.NewBuffer(info)
	err = binary.Read(buff, binary.BigEndian, &sig)
	if err != nil {
		fmt.Println("lectura de binarios fallo: ", err)
	}
	return sig
}

//lee los bytes correspondientes al tamaño que se mando
func leerBytes(de *os.File, noBytes int) []byte {
	infobytes := make([]byte, noBytes) //crea coleccion de bytes
	_, err := de.Read(infobytes)
	if err != nil {
		fmt.Println("ERROR, ", err)
	}
	return infobytes
}

//escribe directamente en el archivo
func escribirBytes(archivo *os.File, bytes []byte) {
	_, err := archivo.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}

}
