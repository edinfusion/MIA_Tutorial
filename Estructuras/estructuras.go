package Estructuras

/*
en golang es recomendable usar int32 o int64
ya que si solo se usa int, este tiene
un tamaño variable por lo cual el struct
nunca tendra un tamaño fijo, cosa que no
beneficia para el sistema de archivos
*/

type Indice struct {
	P_libre          int32 //4 bytes cada int32
	P_primerprofesor int32
	P_ultimoprofesor int32
	P_primeralumno   int32
	P_ultimoalumno   int32
} //tamaño de struct 20 bytes

type Apuntador struct {
	Siguiente int32
}

type Profesor struct {
	Id_profesor int32
	Cui         [13]byte
	Nombre      [25]byte
	Curso       [25]byte
}

type Estudiante struct {
	Id_estudiante int32
	Cui           [13]byte
	Nombre        [25]byte
	Carnet        [10]byte
}
