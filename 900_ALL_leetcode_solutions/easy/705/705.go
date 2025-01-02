// easy 705
// 67.2 acc rate
// design hashset

// JAN 1, 2025


func doNothing() {
    return
}

func find_index_in_slice (s []int, item int) int {
    ans := -1
    for i:= 0 ; i < len(s) ; i++ {
        if s[i] == item {
          return i
        }
    }
    return ans
}


// all it has is a slice that has uniq set of integers
type MyHashSet struct {
    set []int
    
}


// just return an empty struct
func Constructor() MyHashSet {
    a := MyHashSet { set : []int{}}
    return a
    
}


// if  not found, add in the end
func (this *MyHashSet) Add(key int)  {
    // check is set already has key or not
    found := false
    for i := 0 ; i < len(this.set) ; i++ {
        if this.set[i] == key { // found
           found = true
        }
    }

    if found != true {
       this.set = append(this.set,key) 
    }
    
    
}

// if found, find the index, then use variadic call to cut the slice in two and paste
func (this *MyHashSet) Remove(key int)  {
    indx := find_index_in_slice(this.set,key) 
    if  indx == -1 {
        doNothing()
    } else {
        this.set = append(this.set[:indx],this.set[indx+1:]...)
    }
    
}

// simple
func (this *MyHashSet) Contains(key int) bool {
    ans := false
    for i := 0 ; i < len(this.set) ; i++ {
        if this.set[i] == key {
            return true
        }
    }

    return ans
    
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

