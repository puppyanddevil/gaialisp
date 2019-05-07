package gaialisp

import (
	"fmt"
)

type VM struct {
	root *Node
}

func NewVM(root *Node) *VM {
	vm := &VM{}
	vm.root = root

	return vm

}

func (self *VM) Eval() {
	self.evalNode(self.root)
}

func (self *VM) evalNode(node *Node) *Node {
	//fmt.Printf("eval node: %d sval %s ival %f\n", node.nodeType, node.sval, node.ival)
	switch node.nodeType {
	case NTSEXPR:
		return self.evalSExpr(node)
	case NTID:
		panic("vm error: variable is not supported currently")
	case NTNUM:
		return node
	case NTLITERAL:
		return node
	default:
		panic(fmt.Sprintf("vm error: unknown nodeType %d", node.nodeType))
	}
}

func (self *VM) evalSExpr(node *Node) *Node {
	if node.nodeType != NTSEXPR {
		panic("vm error: eval s-expr is not s-expr")
	} else {
		subs := node.subs
		if len(subs) <= 0 {
			panic("vm error: not support empty s-expr yet")
		}
		firstNode := subs[0]
		// eval first node #begin
		if firstNode.nodeType == NTID {
			// todo support variable
			// current only support internal functions
			return self.callInternalFunction(firstNode.sval, subs[1:])
		} else {
			// todo support normal node eval
			panic("only identifier is support in s-expr's first node currently")
		}

		// eval first node #end
	}
}

func (self *VM) callInternalFunction(functionName string, args []*Node) *Node {
	if functionName == "print" {
		// todo support more args
		if len(args) > 0 {
			argOrg := args[0]
			arg := self.evalNode(argOrg)
			if arg.nodeType == NTNUM {
				fmt.Printf("%f\n", arg.ival)
			} else if arg.nodeType == NTLITERAL {
				fmt.Printf("%s\n", arg.sval)
			} else {
				fmt.Println("print only can print number or literal")
			}
		} else {
			fmt.Println("print requires at least 1 arg")
		}
		//fixme return a nil node instead
		return &Node{nodeType: NTNUM, ival: 0}
	} else if functionName == "+" {
		var sumNum float64 = 0
		if len(args) > 0 {
			for _, arg := range args {
				node := self.evalNode(arg)
				if node.nodeType == NTNUM {
					sumNum += node.ival
				} else {
					fmt.Println("sum only support number")
					break
				}
			}

		} else {
			fmt.Println("sum requires at least 1 arg")
		}
		return &Node{nodeType: NTNUM, ival: sumNum}
	} else if functionName == "progn" {
		//sequence execution
		node := &Node{nodeType: NTNUM, ival: 0}

		if len(args) > 0 {
			//fixme default return node should be nil node
			for _, arg := range args {
				node = self.evalNode(arg)
			}
		} else {
			fmt.Println("progn requires at least 1 arg")
		}
		return node

	} else {
		panic(fmt.Sprintf("unknown function name: %s", functionName))
	}

}