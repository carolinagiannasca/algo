package lista

type Lista[T any] interface {
	// EstaVacia devuelve true si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento al principio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento de la lista y lo devuelve en caso de que ésta tenga elementos.
	// Si está vacía, entra en pánico con el mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primer elemento de la lista. Si está vacía, entra en pánico con el mensaje
	// "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del último elemento de la lista. Si está vacía, entra en pánico con el mensaje
	// "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos que contiene la lista.
	Largo() int

	// Iterar aplica la función visitar a cada uno de los datos de la lista, hasta que ésta llegue a su fin
	// ó hasta que visitar devuelva false.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador externo que permite recorrer, insertar y eliminar elementos de la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	// VerActual devuelve el elemento de la lista al que apunta actualmente el iterador.
	// Si ya se han iterado todos los elementos entra en pánico con el mensaje "El iterador termino de iterar".
	VerActual() T

	// HaySiguiente devuelve true si quedan elementos por iterar desde la posición actual, false en caso de que se haya
	// recorrido la lista por completo.
	HaySiguiente() bool

	// Siguiente avanza el iterador al elemento siguiente de la lista en caso de que lo haya.
	// Si ya se han iterado todos los elementos entra en pánico con el mensaje "El iterador termino de iterar".
	Siguiente()

	// Insertar agrega un nuevo elemento a la lista, el cual toma la posición a la que apunta el iterador actualmente.
	// Una vez realizada la operación, el iterador apunta al elemento recién agregado.
	Insertar(T)

	// Borrar elimina de la lista el elemento al que apunta el iterador y lo devuelve. Al finalizar, el iterador apunta al elemento siguiente.
	// Si ya se han iterado todos los elementos entra en pánico con el mensaje "El iterador termino de iterar".
	Borrar() T
}
