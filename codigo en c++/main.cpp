#include <iostream>
#include <string.h>
#include <cstdio>

using namespace std;

// STRUCTS

typedef struct punteros {
    int now;        //direccion donde se guardo el ultimo registro y se guarda el siguiente
    int root_prof;  //posicion del primer profesor
    int end_prof;   //posicion del ultimo profesor guardado, para actualizar su siguiente
    int root_alum;  //posicion del primer alumno
    int end_alum;   //posicion del ultimo alumno guardado, para actualizar su siguiente
};

typedef struct siguente {
    int next; //guarda la direccion del siguiente profesor o alumno
};

typedef struct Profesor {
    int id_profesor;// 4 bytes
    char cui[13];//   13 bytes
    char nombre[25];// 25 bytes
    char curso[25];//  25 bytes
    //                 sizeof (67)
};

typedef struct Estudiante {
    int id_estudiante;
    char cui[13];
    char nombre[25];
    char carnet[10];
    //                sizeof (52)
};


//metodo para crear archivo:
void crearArchivo()
{
    FILE *archivo;
    archivo = fopen("HT1.bin", "rb"); //Apertura en modo de sólo lectura. El archivo debe existir.
    if (archivo == NULL)              //se verifica si ya existe el archivo
    {
        punteros in_punteros;
        archivo = fopen("HT1.bin", "wb+");//Apertura en modo de lectura/escritura. Si el archivoexiste, se reescribirá (pierde el contenido anterior).
        fseek(archivo, 0, SEEK_SET);
        in_punteros.end_alum = -1;
        in_punteros.end_prof = -1;
        in_punteros.root_alum = -1;
        in_punteros.root_prof = -1;
        in_punteros.now = sizeof(in_punteros);//coloco now hasta el final del struct indice
        fwrite(&in_punteros, sizeof(in_punteros), 1, archivo);
        fclose(archivo);
    }
}

void crearProfesor()
{
    //limpiar pantalla
    system("clear");
    //creacion de struct profe
    Profesor profe;
    cout << " ****** REGISTRO DE PROFESOR ******" << endl;
    cout << endl;
    cout << " ID_PROFESOR: ";
    scanf("%i", &profe.id_profesor);
    cout << endl;
    cout << " CUI: ";
    scanf("%s", &profe.cui);
    cout << endl;
    cout << " NOMBRE: ";
    scanf("%s", &profe.nombre);
    cout << endl;
    cout << " CURSO: ";
    scanf("%s", &profe.curso);
    cout << endl;

    FILE *archivo;
    //se abre archivo binario
    archivo = fopen("HT1.bin", "rb+");//Apertura en modo de lectura/escritura. El archivo debe existir.
    punteros indice;
    siguente sig;
    fseek(archivo, 0, SEEK_SET);//muevo puntero al inicio de archivo
    //leo direcion de indice
    fread(&indice, sizeof(indice), 1, archivo);
    //verifico si raiz profesor esta vacía
    if (indice.root_prof == -1)
    {
        indice.root_prof = indice.now; // si esta vacia actualizo la direccion
    }
    else //si ya existe el primer profesor, entoces agrego la direccion en donde se ubicara el siguiente profesor
    {
        fseek(archivo, indice.end_prof, SEEK_SET);
        fread(&sig, sizeof(sig), 1, archivo);
        //cout <<"posicion sig " <<indice.end_prof << endl;
        sig.next = indice.now;//queda al inicio, aqui esta la direccion donde empieza el profesor
        fseek(archivo, indice.end_prof, SEEK_SET);// regreso puntero
        fwrite(&sig, sizeof(sig), 1, archivo);//guardo siguiente
    }

    fclose(archivo);
    archivo = fopen("HT1.bin", "ab+");//ab+ sirve para agregar al final de archivo binario
    fseek(archivo, indice.now, SEEK_SET);// movimiento de posicion de guardar struct profesor
    sig.next = -1;//se deja el siguiente en -1 ya que es el ultimo guardado
    fwrite(&profe, sizeof(profe), 1, archivo);//ALMACENO STRUCT PROFESOR
    fclose(archivo);
    archivo = fopen("HT1.bin", "rb+");//Apertura en modo de lectura/escritura. El archivo debe existir.
    fseek(archivo, indice.now + sizeof(profe), SEEK_SET);//mover puntero al final de struct profesor + now
    //cout <<"posicion sig " <<indice.now << " + "<< sizeof (profe) << endl;
    fwrite(&sig, sizeof(sig), 1, archivo);//guardo siguiente -1 ya que no hay siguiente aun es el ultimo
    indice.end_prof = indice.now + sizeof(profe);//actualizo ubicacion de puntero de el ultimo profesor
    fseek(archivo, 0, SEEK_SET);
    indice.now = sizeof(profe) + sizeof(sig) + indice.now;//muevo la posicion de now, para que este lista para ingresar un nuevo profesor
    fwrite(&indice, sizeof(indice), 1, archivo);
    fclose(archivo);
    cout << "SE HA REGISTRADO UN NUEVO PROFESOR CON EXITO" << endl;
}

void crearEstudiante()
{
    //limpiar pantalla
    system("clear");
    //creacion de struct profe
    Estudiante estud;
    cout << " ****** REGISTRO DE ESTUDIANTE ******" << endl;
    cout << endl;
    cout << " ID_ESTUDIANTE: ";
    scanf("%i", &estud.id_estudiante);
    cout << endl;
    cout << " CUI: ";
    scanf("%s", &estud.cui);
    cout << endl;
    cout << " NOMBRE: ";
    scanf("%s", &estud.nombre);
    cout << endl;
    cout << " CARNET: ";
    scanf("%s",&estud.carnet);
    cout << endl;

    FILE *archivo;
    //se abre archivo binario
    archivo = fopen("HT1.bin", "rb+");//Apertura en modo de lectura/escritura. El archivo debe existir.
    punteros indice;
    siguente sig;
    fseek(archivo, 0, SEEK_SET);//muevo puntero al inicio de archivo
    //leo direcion de indice
    fread(&indice, sizeof(indice), 1, archivo);
    //verifico si raiz profesor esta vacía
    if (indice.root_alum == -1)
    {
        indice.root_alum = indice.now; // si esta vacia actualizo la direccion
    }
    else //si ya existe el primer profesor, entoces agrego la direccion en donde se ubicara el siguiente profesor
    {
        fseek(archivo, indice.end_alum, SEEK_SET);
        fread(&sig, sizeof(sig), 1, archivo);
        //cout <<"posicion sig " <<indice.end_prof << endl;
        sig.next = indice.now;//queda al inicio, aqui esta la direccion donde empieza el profesor
        fseek(archivo, indice.end_alum, SEEK_SET);// regreso puntero
        fwrite(&sig, sizeof(sig), 1, archivo);//guardo siguiente
    }

    fclose(archivo);
    archivo = fopen("HT1.bin", "ab+");//ab+ sirve para agregar al final de archivo binario
    fseek(archivo, indice.now, SEEK_SET);// movimiento de posicion de guardar struct profesor
    sig.next = -1;//se deja el siguiente en -1 ya que es el ultimo guardado
    fwrite(&estud, sizeof(estud), 1, archivo);//ALMACENO STRUCT PROFESOR
    fclose(archivo);
    archivo = fopen("HT1.bin", "rb+");//Apertura en modo de lectura/escritura. El archivo debe existir.
    fseek(archivo, indice.now + sizeof(estud), SEEK_SET);//mover puntero al final de struct profesor + now
    //cout <<"posicion sig " <<indice.now << " + "<< sizeof (profe) << endl;
    fwrite(&sig, sizeof(sig), 1, archivo);//guardo siguiente -1 ya que no hay siguiente aun es el ultimo
    indice.end_alum = indice.now + sizeof(estud);//actualizo ubicacion de puntero de el ultimo profesor
    fseek(archivo, 0, SEEK_SET);
    indice.now = sizeof(estud) + sizeof(sig) + indice.now;//muevo la posicion de now, para que este lista para ingresar un nuevo profesor
    fwrite(&indice, sizeof(indice), 1, archivo);
    fclose(archivo);
    cout << "SE HA REGISTRADO UN NUEVO ESTUDIANTE CON EXITO" << endl;
}

void verRegistros()
{
    system("clear");
    FILE *file;
    //se crea un struct para reporte de profesores
    Profesor profesreport;
    //se crea un struct para reporte de ESTUDIANTES
    Estudiante studreport;
    //struct de indice
    punteros indice;
    //struct de siguientes
    siguente sig;
    file = fopen("HT1.bin", "rb");//Apertura en modo de sólo lectura. El archivo debe existir.
    if (file == NULL)
    {
        cout << "NO SE PUEDE ACCEDER AL ARCHIVO";
    }
    else
    {
        //muevo puntero al inicio de archivo
        fseek(file, 0, SEEK_SET);

        //leo de archivo el struct de indices
        fread(&indice,sizeof (indice),1,file);

        //MUEVO EL PUNTERO HASTA EL PRIMER REGISTRO PROFESOR
        fseek(file,indice.root_prof,SEEK_SET);

        //IMPRESIÓN DE PRIMER PROFESOR
        fread(&profesreport,sizeof (profesreport),1,file);
        cout << "****** PROFESOR ******" << endl;
        cout << "    ID: " << profesreport.id_profesor << endl;
        cout << "   CUI: " << profesreport.cui << endl;
        cout << "NOMBRE: " << profesreport.nombre << endl;
        cout << " CURSO: " << profesreport.curso << endl;

        //verifico si hay siguientes para seguir imprimiendo
        fseek(file,indice.root_prof+sizeof (profesreport),SEEK_SET);
        //cout << "valor: " << indice.root_prof+sizeof (profesreport)<<endl;
        fread(&sig, sizeof(sig),1,file);
        while (sig.next != -1)
        {
            fseek(file,sig.next,SEEK_SET);
            fread(&profesreport,sizeof (profesreport),1,file);
            cout << "****** PROFESOR ******" << endl;
            cout << "    ID: " << profesreport.id_profesor << endl;
            cout << "   CUI: " << profesreport.cui << endl;
            cout << "NOMBRE: " << profesreport.nombre << endl;
            cout << " CURSO: " << profesreport.curso << endl;
            fseek(file,sig.next+sizeof (profesreport),SEEK_SET);
            fread(&sig, sizeof(sig),1,file);
        }


        //muevo puntero al inicio de archivo
        fseek(file, 0, SEEK_SET);

        //leo de archivo el struct de indices
        fread(&indice,sizeof (indice),1,file);

        //MUEVO EL PUNTERO HASTA EL PRIMER REGISTRO ESTUDIANTE
        fseek(file,indice.root_alum,SEEK_SET);

        //IMPRESIÓN DE PRIMER ESTUDIANTE
        fread(&studreport,sizeof (studreport),1,file);
        cout << "****** ESTUDIANTE ******" << endl;
        cout << "    ID: " << studreport.id_estudiante << endl;
        cout << "   CUI: " << studreport.cui << endl;
        cout << "NOMBRE: " << studreport.nombre << endl;
        cout << "CARNET: " << studreport.carnet << endl;

        //verifico si hay siguientes para seguir imprimiendo
        fseek(file,indice.root_alum+sizeof (studreport),SEEK_SET);
        //cout << "valor: " << indice.root_prof+sizeof (profesreport)<<endl;
        fread(&sig, sizeof(sig),1,file);
        while (sig.next != -1)
        {
            fseek(file,sig.next,SEEK_SET);
            fread(&studreport,sizeof (studreport),1,file);
            cout << "****** ESTUDIANTE ******" << endl;
            cout << "    ID: " << studreport.id_estudiante << endl;
            cout << "   CUI: " << studreport.cui << endl;
            cout << "NOMBRE: " << studreport.nombre << endl;
            cout << "CARNET: " << studreport.carnet << endl;
            fseek(file,sig.next+sizeof (studreport),SEEK_SET);
            fread(&sig, sizeof(sig),1,file);
        }

        fclose(file);
    }
}
int main()
{
    crearArchivo();
    int opmenu=0;
    while (opmenu != 4)
    {
        switch (opmenu)
        {
            case 0:
            {
                //limpiar pantalla
                //system("clear");
                cout << "   ******** MENU ********" << endl;
                cout << endl;
                cout << "1. REGISTRO DE PROFESOR  " << endl;
                cout << "2. REGISTRO DE ESTUDIANTE" << endl;
                cout << "3. VER REGISTROS         " << endl;
                cout << "4. SALIR                 " << endl;
                cout << "\n\n INGRESE UNA OPCION: ";
                scanf("%i", &opmenu);
            }
                break;
            case 1:
            {
                crearProfesor();
                opmenu=0;
            }
                break;
            case 2:
            {
                crearEstudiante();
                opmenu=0;
            }
                break;
            case 3:
                verRegistros();
                opmenu = 0;
                break;

            default:
                cout << "INGRESE UNA OPCION CORRECTA" << endl;
                opmenu=0;
                break;
        }
    }
    return 0;
}