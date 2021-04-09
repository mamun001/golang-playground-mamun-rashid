
func defangIPaddr(address string) string {
    res1 := strings.Replace(address, ".", "[.]", 3)
    return res1
}
