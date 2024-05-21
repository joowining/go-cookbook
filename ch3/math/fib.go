package math

import "math/big"

var memoize map[int]*big.Int

func init() {
	memoize = make(map[int]*big.Int)
}

func Fib(n int) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}

	if n < 2 {
		memoize[n] = big.NewInt(1)
	}

	if val, ok := memoize[n]; ok {
		return val
	}

	memoize[n] = big.NewInt(0)
	memoize[n].Add(memoize[n], Fib(n-1)) // Add는 bigInt의 메소드이다. 
	memoize[n].Add(memoize[n], Fib(n-2)) // 최초로 접근하는 피보나치 값은 0이기에 그전 값 Fib(n-1), Fib(n-2)에 접근해 둘을 더하는 재귀적 호출이 이루어진다. 

	return memoize[n]
}
