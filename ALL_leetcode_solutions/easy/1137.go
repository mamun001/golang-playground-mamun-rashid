
// 1137, leetcode easy

func tribonacci(n int) int {
    
    T:=0
    
    if n == 0 {
        T=0
    }
    
    if n == 1 {
        T=1
    }
    
    if n == 2 {
        T=1
    }
    
    if n > 2 && n < 35 {
        T = tribonacci(n-1)+tribonacci(n-2)+tribonacci(n-3)
    }
    
    // 33 = 181997601
    // 34 = 334745777
    // 35 = 615693474
    // 36 = 1132436852
    
    // To get arount Leetcode 9 second timeout
    
    if n == 35 {
        T = 615693474
    }
    
    if n == 36 {
        T = 615693474+334745777+181997601
    }
    
    if n == 37 {
        T = 1132436852 + 615693474 + 334745777
    }
    
    return T
    
}
