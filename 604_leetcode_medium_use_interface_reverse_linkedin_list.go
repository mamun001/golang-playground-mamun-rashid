/*   Below is the interface for ImmutableListNode, which is already defined for you.
 *
 *   type ImmutableListNode struct {
 *       
 *   }
 *
 *   func (this *ImmutableListNode) getNext() ImmutableListNode {
 *		// return the next node.
 *   }
 *
 *   func (this *ImmutableListNode) printValue() {
 *		// print the value of this node.
 *   }
 */



func printLinkedListInReverse(head ImmutableListNode) {
    
    //a := make([]int, 0)
    
    //head.printValue()
    //head.getNext().printValue()
    
    // works
    //if head.getNext() == nil {
        //fmt.Println("nil")
    //}
    
    // works
    //b := head // make  a copy
    //b.printValue()
    
    // find length of the linkedlist
    b := head
    length := 0
    for b != nil {
        //b.printValue()
        b = b.getNext()
        length++
    }
    //fmt.Println(length)
    
    
    
    // run length times
    // first time, go all the way to the end
    // each time, go ONE less step forward
    for i:=1; i < length+1 ; i++ {
      b = head
      for j:=0; j<length-i; j++ {
          b = b.getNext()
      }
      b.printValue()
    }
    
    return
}
