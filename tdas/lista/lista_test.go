package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_VOLUMEN = 100000
)

type Persona struct {
	nombre string
	edad   int
}

// Tests de lista

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, TDALista.MSG_LISTA_VACIA, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, TDALista.MSG_LISTA_VACIA, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, TDALista.MSG_LISTA_VACIA, func() { lista.BorrarPrimero() })
}

func TestInsertarPrimeroUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(23)
	require.EqualValues(t, 23, lista.VerPrimero(), "Si la lista tiene un sólo elemento, el primero es ese elemento")
	require.EqualValues(t, 23, lista.VerUltimo(), "Si la lista tiene un sólo elemento, el último es ese elemento")
	require.EqualValues(t, 1, lista.Largo(), "Después de agregarle un elemento a una lista vacía, tiene largo 1")
	require.False(t, lista.EstaVacia(), "Después de agregarle un elemento a una lista vacía, deja de estar vacía")
}

func TestInsertarUltimoUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("hola")
	require.EqualValues(t, "hola", lista.VerPrimero(), "Si la lista tiene un sólo elemento, el primero es ese elemento")
	require.EqualValues(t, "hola", lista.VerUltimo(), "Si la lista tiene un sólo elemento, el último es ese elemento")
	require.EqualValues(t, 1, lista.Largo(), "Después de agregarle un elemento a una lista vacía, tiene largo 1")
	require.False(t, lista.EstaVacia(), "Después de agregarle un elemento a una lista vacía, deja de estar vacía")
}

func TestBorrarPrimeroUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[bool]()
	lista.InsertarPrimero(true)
	require.EqualValues(t, true, lista.BorrarPrimero(), "El elemento eliminado es el único de la lista")
	require.True(t, lista.EstaVacia(), "Después de eliminar el único elemento de la lista, queda vacía")
	require.EqualValues(t, 0, lista.Largo(), "Después de eliminar el único elemento de la lista, tiene largo 0")
	require.PanicsWithValue(t, TDALista.MSG_LISTA_VACIA, func() { lista.VerPrimero() }, "Después de eliminar el único elemento de la lista, no se puede ver el primero")
	require.PanicsWithValue(t, TDALista.MSG_LISTA_VACIA, func() { lista.VerUltimo() }, "Después de eliminar el único elemento de la lista no se puede ver el último")
}

func insertarPrimeroVarios[T any](t *testing.T, lista TDALista.Lista[T], valores []T) {
	for i, val := range valores {
		lista.InsertarPrimero(val)
		require.EqualValues(t, valores[0], lista.VerUltimo(), "El último elemento de la lista es el primero en ser agregado al principio")
		require.EqualValues(t, val, lista.VerPrimero(), "El primer elemento de la lista es el último en ser agregado al principio")
		require.EqualValues(t, i+1, lista.Largo(), "El largo de la lista es correcto")
	}
	require.False(t, lista.EstaVacia(), "La lista no está vacía después de insertar varios elementos al principio")
}

func insertarUltimoVarios[T any](t *testing.T, lista TDALista.Lista[T], valores []T) {
	for i, val := range valores {
		lista.InsertarUltimo(val)
		require.EqualValues(t, val, lista.VerUltimo(), "El último elemento de la lista es el último en ser agregado al final")
		require.EqualValues(t, valores[0], lista.VerPrimero(), "El primer elemento de la lista es el primero en ser agregado al final")
		require.EqualValues(t, i+1, lista.Largo(), "El largo de la lista es correcto")
	}
	require.False(t, lista.EstaVacia(), "La lista no está vacía después de insertar varios elementos al principio")
}

func borrarPrimeroVarios[T any](t *testing.T, lista TDALista.Lista[T], valores []T) {
	for i, val := range valores {
		require.EqualValues(t, val, lista.VerPrimero(), "El primer elemento de la lista es el correcto")
		require.EqualValues(t, val, lista.BorrarPrimero(), "El elemento borrado es el primero de la lista")
		require.EqualValues(t, len(valores)-i-1, lista.Largo(), "El largo de la lista es correcto")
	}
	require.True(t, lista.EstaVacia(), "La lista está vacía después de borrar todos sus elementos")
	require.PanicsWithValue(t, TDALista.MSG_LISTA_VACIA, func() { lista.VerPrimero() }, "No se puede acceder al primero de la lista después de haber borrado todos sus elementos")
	require.PanicsWithValue(t, TDALista.MSG_LISTA_VACIA, func() { lista.BorrarPrimero() }, "No se puede borrar el primero después de ya haber borrado todos los elementos de la lista")
}

func TestInsertarPrimeroEnteros(t *testing.T) {
	numeros := []int{3, 9, 27, 81}
	lista := TDALista.CrearListaEnlazada[int]()
	insertarPrimeroVarios(t, lista, numeros)
}

func TestInsertarPrimeroCadenas(t *testing.T) {
	cadenas := []string{"Jijo", "Debu"}
	lista := TDALista.CrearListaEnlazada[string]()
	insertarPrimeroVarios(t, lista, cadenas)
}

func TestInsertarUltimoBooleanos(t *testing.T) {
	booleanos := []bool{true, true, false, true, false, false}
	lista := TDALista.CrearListaEnlazada[bool]()
	insertarUltimoVarios(t, lista, booleanos)
}

func TestInsertarUltimoStructs(t *testing.T) {
	structs := []Persona{{nombre: "Gustav", edad: 60}, {nombre: "Nils", edad: 20}}
	lista := TDALista.CrearListaEnlazada[Persona]()
	insertarUltimoVarios(t, lista, structs)
}

func TestBorrarPrimeroFloats(t *testing.T) {
	numeros := []float64{10.1, 3.14, 4.67, 6.01}
	lista := TDALista.CrearListaEnlazada[float64]()
	insertarUltimoVarios(t, lista, numeros)
	borrarPrimeroVarios(t, lista, numeros)
}

func TestBorrarPrimeroCadenas(t *testing.T) {
	cadenas := []string{"raúl", "ricardo", "roberto", "ramón"}
	lista := TDALista.CrearListaEnlazada[string]()
	insertarUltimoVarios(t, lista, cadenas)
	borrarPrimeroVarios(t, lista, cadenas)
}

func TestBorrarPrimeroBooleanos(t *testing.T) {
	booleanos := []bool{true, false, true, false}
	lista := TDALista.CrearListaEnlazada[bool]()
	insertarUltimoVarios(t, lista, booleanos)
	borrarPrimeroVarios(t, lista, booleanos)
}

func TestBorrarPrimeroStructs(t *testing.T) {
	structs := []Persona{{nombre: "Juan", edad: 12}, {nombre: "Pedro", edad: 13}}
	lista := TDALista.CrearListaEnlazada[Persona]()
	insertarUltimoVarios(t, lista, structs)
	borrarPrimeroVarios(t, lista, structs)
}

func TestAlternado(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := range 30 {
		lista.InsertarUltimo(i)
		require.EqualValues(t, i, lista.VerPrimero(), "El primer elemento de la lista es el primero en ser agregado")
		require.EqualValues(t, i, lista.BorrarPrimero(), "El elemento borrado es el primero de la lista")
		require.True(t, lista.EstaVacia(), "Después de insertar y borrar un elemento, la lista queda vacía")
		require.PanicsWithValue(t, TDALista.MSG_LISTA_VACIA, func() { lista.VerPrimero() }, "No se puede ver el primero de la lista después de haber desencolado su único elemento")
		require.PanicsWithValue(t, TDALista.MSG_LISTA_VACIA, func() { lista.BorrarPrimero() }, "No se puede borrar después de haber borrado el único elemento de la lista")
	}
}

func TestVolumenInsertandoPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := range _VOLUMEN {
		lista.InsertarPrimero(i)
		require.EqualValues(t, i, lista.VerPrimero(), "Después de insertar un elemento al principio de la lista, el primero es ese elemento")
		require.EqualValues(t, 0, lista.VerUltimo(), "Después de insertar un elemento al final de la lista, el último se mantiene igual")
	}
	for i := range _VOLUMEN {
		require.EqualValues(t, _VOLUMEN-1-i, lista.VerPrimero(), "Después de borrar un elemento de la lista, el primero pasa a ser el siguiente elemento")
		lista.BorrarPrimero()
	}
}

func TestVolumenInsertandoUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	numeros := make([]int, _VOLUMEN)
	for i := range _VOLUMEN {
		numeros[i] = i
	}
	insertarUltimoVarios(t, lista, numeros)
	borrarPrimeroVarios(t, lista, numeros)
}

func TestListaMantieneOrden(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(1487)
	require.Equal(t, 1, lista.VerPrimero(), "El primer elemento es correcto")
	require.Equal(t, 1487, lista.VerUltimo(), "El último elemento es correcto")
	lista.InsertarPrimero(588)
	lista.InsertarPrimero(700)
	lista.BorrarPrimero()
	require.Equal(t, 588, lista.VerPrimero(), "El primer elemento es correcto")
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.Equal(t, 3, lista.VerPrimero(), "El primer elemento es correcto")
	require.Equal(t, 1487, lista.VerUltimo(), "El último elemento es correcto")
}

// Tests de iterador interno

func TestIterar(t *testing.T) {
	numeros := []int{1, 2, 3, 4, 5}
	lista := TDALista.CrearListaEnlazada[int]()
	insertarUltimoVarios(t, lista, numeros)
	suma := 0
	lista.Iterar(func(num int) bool {
		suma += num
		return true
	})
	require.EqualValues(t, 15, suma, "El resultado de la iteración es el esperado")
}

func TestIterarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	pares := 0
	lista.Iterar(func(num int) bool {
		if num%2 == 0 {
			pares++
		}
		return true
	})
	require.EqualValues(t, 0, pares, "Al iterar una lista vacía, no se obtiene resultado")
}

func TestIterarConCorte(t *testing.T) {
	numeros := []int{1, 2, 3, 4, 5}
	lista := TDALista.CrearListaEnlazada[int]()
	insertarUltimoVarios(t, lista, numeros)
	elemBuscado := 3
	pos := 0
	encontrado := false
	lista.Iterar(func(num int) bool {
		if num == elemBuscado {
			encontrado = true
			return false
		}
		pos++
		return true
	})
	require.EqualValues(t, 2, pos, "La iteración se corta correctamente y el resultado es el esperado")
	require.True(t, encontrado)
}

// Tests de iterador externo

func TestIterListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, TDALista.MSG_ITER_TERMINADA, func() { iter.VerActual() }, "Un iterador recién actúa como si hubiera terminado de iterar si la lista está vacía")
	require.PanicsWithValue(t, TDALista.MSG_ITER_TERMINADA, func() { iter.Siguiente() }, "No se puede ver el siguiente de una lista vacía con el iterador")
	require.False(t, iter.HaySiguiente(), "No hay siguiente si la lista está vacía")
}

func TestRecorrerConIterador(t *testing.T) {
	cadenas := []string{"otorrinolaringologo", "terremoto", "lavarropas", "acido ribonucleico"}
	lista := TDALista.CrearListaEnlazada[string]()
	insertarUltimoVarios(t, lista, cadenas)
	iter := lista.Iterador()
	for _, val := range cadenas {
		require.EqualValues(t, val, iter.VerActual(), "El iterador apunta al elemento correcto")
		iter.Siguiente()
	}
	require.PanicsWithValue(t, TDALista.MSG_ITER_TERMINADA, func() { iter.VerActual() }, "Una vez iterada toda la lista no se puede acceder al elemento actual")
}

func TestInsertarPrincipio(t *testing.T) {
	numeros := []int{2, 3, 5, 7, 11, 13}
	lista := TDALista.CrearListaEnlazada[int]()
	insertarUltimoVarios(t, lista, numeros)
	iter := lista.Iterador()
	iter.Insertar(1)
	require.EqualValues(t, 1, lista.VerPrimero(), "Al insertar un elemento en la posición en la que se crea el iterador, se inserta al principio")
	require.EqualValues(t, 1, iter.VerActual(), "El iterador apunta al elemento recién agregado")
	require.EqualValues(t, 7, lista.Largo(), "El largo se actualiza luego de insertar un elemento")
}

func TestInsertarFinal(t *testing.T) {
	cadenas := []string{"otorrinolaringologo", "terremoto", "lavarropas", "acido ribonucleico"}
	lista := TDALista.CrearListaEnlazada[string]()
	insertarUltimoVarios(t, lista, cadenas)
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter.Insertar("murciélago")
	require.EqualValues(t, "murciélago", lista.VerUltimo(), "Si el iterador apunta al final, el elemento se inserta en la última posición")
	require.EqualValues(t, "murciélago", iter.VerActual(), "El iterador apunta al elemento recién agregado")
	require.EqualValues(t, 5, lista.Largo(), "El largo se actualiza luego de insertar un elemento")
}

func TestInsertarMedio(t *testing.T) {
	booleanos := []bool{true, true, true, true}
	lista := TDALista.CrearListaEnlazada[bool]()
	insertarUltimoVarios(t, lista, booleanos)
	iter := lista.Iterador()
	posicion := 2
	for i := 0; i < posicion; i++ {
		iter.Siguiente()
	}
	iter.Insertar(false)
	require.EqualValues(t, 5, lista.Largo(), "El largo se actualiza luego de insertar un elemento")
	require.EqualValues(t, false, iter.VerActual(), "El iterador apunta al elemento recién agregado")
	for i := 0; i < posicion; i++ {
		lista.BorrarPrimero()
	}
	require.EqualValues(t, false, lista.VerPrimero(), "El elemento se inserta en la posición correcta")
}

func TestInsertarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(3)
	require.EqualValues(t, 3, iter.VerActual(), "El elemento actual es el recién insertado")
	require.EqualValues(t, 3, lista.VerPrimero(), "El primer elemento es el único de la lista")
	require.EqualValues(t, 3, lista.VerUltimo(), "El último elemento es el único de la lista")
	require.EqualValues(t, 1, lista.Largo(), "El largo de la lista es correcto")
}

func TestBorrarPrincipio(t *testing.T) {
	numeros := []int{2, 4, 8, 16, 32, 64}
	lista := TDALista.CrearListaEnlazada[int]()
	insertarUltimoVarios(t, lista, numeros)
	iter := lista.Iterador()
	require.EqualValues(t, 2, iter.Borrar(), "Al borrar un elemento en la posición en la que se crea el iterador, se borra el primer elemento de la lista")
	require.EqualValues(t, 4, iter.VerActual(), "El iterador apunta al elemento recién agregado")
	require.EqualValues(t, 4, lista.VerPrimero(), "El primer elemento de la lista es correcto")
	require.EqualValues(t, 5, lista.Largo(), "El largo es correcto luego de eliminar un elemento")
}

func TestBorrarFinal(t *testing.T) {
	cadenas := []string{"hola", "mundo", "!!!"}
	lista := TDALista.CrearListaEnlazada[string]()
	insertarUltimoVarios(t, lista, cadenas)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	require.EqualValues(t, "!!!", iter.Borrar(), "Si el iterador apunta al último elemento, se elimina el último de la lista")
	require.PanicsWithValue(t, TDALista.MSG_ITER_TERMINADA, func() { iter.VerActual() }, "No se puede ver el elemento actual después de borrar el último de la lista")
	require.False(t, iter.HaySiguiente(), "No hay siguiente después de borrar el último elemento")
	require.EqualValues(t, "mundo", lista.VerUltimo(), "El último elemento de la lista es correcto")
	require.EqualValues(t, 2, lista.Largo(), "El largo es correcto luego de eliminar un elemento")
}

func TestBorrarMedio(t *testing.T) {
	booleanos := []bool{true, true, false, true}
	lista := TDALista.CrearListaEnlazada[bool]()
	insertarUltimoVarios(t, lista, booleanos)
	iter := lista.Iterador()
	posicion := 2
	for i := 0; i < posicion; i++ {
		iter.Siguiente()
	}
	require.EqualValues(t, false, iter.Borrar(), "El elemento borrado es correcto")
	require.EqualValues(t, true, iter.VerActual(), "El elemento actual es correcto")
	require.EqualValues(t, 3, lista.Largo(), "El largo es correcto luego de eliminar un elemento")
}

func TestBorrarListaConUnicoElem(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	iter := lista.Iterador()
	iter.Insertar("hola")
	require.EqualValues(t, "hola", iter.Borrar(), "El elemento borrado es el único de la lista")
	require.True(t, lista.EstaVacia(), "Después de borrar el único elemento de la lista, queda vacía")
}

func TestBorrarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, TDALista.MSG_ITER_TERMINADA, func() { iter.Borrar() }, "No se puede borrar con el iterador en una lista vacía")
}
