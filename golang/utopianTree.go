package golang

func utopianTree(n int32) int32 {
    if n == 0 {
        return 1
    } else if n%2==0 {
        return utopianTree(n-1) + 1
    } else {
		return 2 * utopianTree(n-1)
    }

    return 1

}