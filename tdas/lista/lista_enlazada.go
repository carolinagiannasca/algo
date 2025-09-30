package lista

const (
	MSG_LISTA_VACIA    = "La lista esta vacia"
	MSG_ITER_TERMINADA = "El iterador termino de iterar"
)

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	anterior *nodoLista[T]
	actual   *nodoLista[T]
	lista    *listaEnlazada[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	return lista
}

func crearNodoLista[T any](dato T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = dato
	return nodo
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoPrimero := crearNodoLista(dato)
	if lista.EstaVacia() {
		lista.ultimo = nuevoPrimero
	}
	nuevoPrimero.siguiente = lista.primero
	lista.primero = nuevoPrimero
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoUltimo := crearNodoLista(dato)
	if lista.EstaVacia() {
		lista.primero = nuevoUltimo
	} else {
		lista.ultimo.siguiente = nuevoUltimo
	}
	lista.ultimo = nuevoUltimo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(MSG_LISTA_VACIA)
	}
	nodo := lista.primero
	lista.primero = nodo.siguiente
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return nodo.dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(MSG_LISTA_VACIA)
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(MSG_LISTA_VACIA)
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil {
		if !visitar(actual.dato) {
			break
		}
		actual = actual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := new(iterListaEnlazada[T])
	iterador.lista = lista
	iterador.actual = lista.primero
	return iterador
}

func (iter *iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic(MSG_ITER_TERMINADA)
	}
	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(MSG_ITER_TERMINADA)
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iterListaEnlazada[T]) Insertar(dato T) {
	nuevoNodo := crearNodoLista(dato)
	nuevoNodo.siguiente = iter.actual
	if iter.anterior == nil {
		iter.lista.primero = nuevoNodo
	} else {
		iter.anterior.siguiente = nuevoNodo
	}
	if iter.actual == nil {
		iter.lista.ultimo = nuevoNodo
	}
	iter.actual = nuevoNodo
	iter.lista.largo++
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic(MSG_ITER_TERMINADA)
	}
	nodoActual := iter.actual
	if iter.anterior == nil {
		iter.lista.primero = nodoActual.siguiente
	} else {
		iter.anterior.siguiente = nodoActual.siguiente
	}
	if iter.actual.siguiente == nil {
		iter.lista.ultimo = iter.anterior
	}
	iter.actual = iter.actual.siguiente
	iter.lista.largo--
	return nodoActual.dato
}
