package main

import "fmt"
import "os"
import "strconv"
import "flag"

type TreeNode struct {
    Value  string
    LeftChild  *TreeNode
    RightChild  *TreeNode
}

func check_null(e error) {
    if e != nil {
        panic(e)
    }
}

func perform_arthmitic(result *int, operation string, left_value int, right_value int) {
    if operation == "+" {
        *result = left_value + right_value
    }
    if operation == "-" {
        *result = left_value - right_value
    }
    if operation == "*" {
        *result = left_value * right_value
    }
    if operation == "." {
        *result = left_value / right_value
    }
}

func deserialize_tree(node *TreeNode, f *os.File) *TreeNode {
    if node == nil {
        node = &TreeNode{}
    }

    bytes := make([]byte, 2)
    n, err := f.Read(bytes)
    check_null(err)

    if n < 1 {
        node.Value = "_"
    } else {
        node.Value = string(bytes[0])
    }

    if node.Value != "_" {
        node.LeftChild = deserialize_tree(node.LeftChild, f)
        node.RightChild = deserialize_tree(node.RightChild, f)
    }

    return node
}

func evaluate_tree(node *TreeNode) int {
    value, err := strconv.Atoi(node.Value)
    var result int

    if err == nil {
        result = value
    } else {
        left_value := evaluate_tree(node.LeftChild)
        right_value := evaluate_tree(node.RightChild)

        perform_arthmitic(&result, node.Value, left_value, right_value)
    }

    return result
}

func main() {
    flag.Parse()
    filename := flag.Arg(0)

    f, err := os.Open(filename)
    check_null(err)

    root := TreeNode{}
    deserialize_tree(&root, f)
    f.Close()

    result := evaluate_tree(&root)

    fmt.Printf("%d\n", result)
}
