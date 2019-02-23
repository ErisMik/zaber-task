package main

import "fmt"
import "os"
import "strconv"
import "flag"
import "bufio"

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

    // fmt.Printf("%d %s %d\n", left_value, operation, right_value)
}

func deserialize_tree(node *TreeNode, s *bufio.Scanner) *TreeNode {
    if node == nil {
        node = &TreeNode{}
    }

    new_token := s.Scan()

    if !new_token {
        node.Value = "_"
    } else {
        node.Value = s.Text()
    }

    if node.Value != "_" {
        node.LeftChild = deserialize_tree(node.LeftChild, s)
        node.RightChild = deserialize_tree(node.RightChild, s)
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
    scanner := bufio.NewScanner(bufio.NewReader(f))
    scanner.Split(bufio.ScanWords)

    root := TreeNode{}
    deserialize_tree(&root, scanner)
    f.Close()

    result := evaluate_tree(&root)

    fmt.Printf("%d\n", result)
}
