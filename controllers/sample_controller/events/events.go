package events

import "fmt"

func AdditionEvent(obj any) {
	fmt.Println("Add event detected:", obj)
}

func UpdationEvent(oldObj, newObj any){
	fmt.Println("Update event detected:", newObj)
}

func DeletionEvent(obj any){
	fmt.Println("Delete event detected:", obj)
}