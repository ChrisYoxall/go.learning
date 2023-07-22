/*
Some of the benefits of using the Observer pattern are as follows:

 - Loose coupling: Creates a system where a change in one component results in an automatic change in other components.

 - Broadcast communication: When an event happens to a component, all its dependents get notified without needing to
	know who these dependents are.

 - Dynamic relationships: Allows the addition or removal of observers dynamically, offering greater flexibility.

 - Event management: Excellent for event monitoring and handling situations where an action must trigger additional
	actions.
*/

package main

import "log"

type Observer interface {
	Update(message string)
}

type ConcreteObserver struct {
	ID int
}

func (co *ConcreteObserver) Update(message string) {
	log.Println("Observer", co.ID, "updated with message:", message)
}

type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	Notify(message string)
}

type ConcreteSubject struct {
	Observers []Observer
}

func (cs *ConcreteSubject) Register(observer Observer) {
	cs.Observers = append(cs.Observers, observer)
}

func (cs *ConcreteSubject) Deregister(observer Observer) {
	for i, ob := range cs.Observers {
		if ob == observer {
			cs.Observers = append(cs.Observers[:i], cs.Observers[i+1:]...)
			return
		}
	}
}
func (cs *ConcreteSubject) Notify(message string) {
	for _, observer := range cs.Observers {
		observer.Update(message)
	}
}

// This line is a compile-time check to ensure that the ConcreteSubject struct implements the Subject interface. If
// ConcreteSubject does not satisfy all the methods declared in Subject, the code will fail to compile. This is useful
// to catch any inadvertent changes that may cause ConcreteSubject to no longer implement the Subject interface. The
// variable is discarded (hence the underscore) because we don't need it for anything.
// In practice wouldn't use this as unit tests should confirm functionality.
var _ Subject = (*ConcreteSubject)(nil)

func main() {
	s := &ConcreteSubject{}

	o1 := &ConcreteObserver{1}
	o2 := &ConcreteObserver{2}

	s.Register(o1)
	s.Register(o2)
	s.Notify("Hello, observers!")

	s.Deregister(o1)
	s.Notify("Observer 1 has been removed.")
}
