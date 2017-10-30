package main

import (
	"errors"
	"net/rpc"
)

func main() {

	math:=new(Math)
	rpc.Register(math)

}

//作为一个服务必须要有一个类型

//func () NAME (ARGS TYPE,REPLY *TYPE) error

type Args struct {
	A,B int
}

type Math int

func (m *Math)multiply (args *Args,reply *int) {
*reply*args.A*args.B
	return nil
}


type Quotient struct {
	Quo,Rem int
}

func(m *Math)Divide(args *Args,quo *Quotient)error{
	if args.B==0{
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem=args.A % args.B
}



