package handlers

import (
	"fmt"
	"math"

	"github.com/nozaq/fibo-grpc/fiborpc"

	"golang.org/x/net/context"
)

// FiboImpl implements LeonardoClient.
type FiboImpl struct {
}

// Fibo returns a value in the fibonacci sequence at the provided index.
func (s *FiboImpl) Fibo(ctx context.Context, r *fiborpc.FiboRequest) (*fiborpc.FiboReply, error) {
	i := r.Index

	if i <= 0 {
		return nil, fmt.Errorf("index must be positive: %d", i)
	}

	v := fibo(int(i))

	res := &fiborpc.FiboReply{
		Value: int32(v),
	}

	return res, nil
}

func fibo(i int) int {
	fi := float64(i)

	sqrtFive := math.Sqrt(5.0)
	numerator := math.Pow(1.0+sqrtFive, fi) - math.Pow(1.0-sqrtFive, fi)
	denominator := math.Pow(2.0, fi) * sqrtFive

	return int(math.Floor(numerator / denominator))
}
