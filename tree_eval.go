package main

import (
        "fmt"
        "os"
        "strconv"
        "flag"
        "bufio"
)

// Class to represent a node of an arithmitic binary tree
type TreeNode struct {
    Value       string      // The value of the node, either a number "6" or an operation "+"
    LeftChild   *TreeNode   // Pointer to the left child node of this node
    RightChild  *TreeNode   // Pointer to the right child node of this node
}

// Exit the program if an error is detected
func check_err(e error) {
    if e != nil {
        panic(e)
    }
}

// Check what operation to do and perform the operation on the provided values
func perform_arthmitic(result *int, operation string, left_value int, right_value int) {
    if operation == "+" {
        *result = left_value + right_value

    } else if operation == "-" {
        *result = left_value - right_value

    } else if operation == "*" {
        *result = left_value * right_value

    } else if operation == "." {
        *result = left_value / right_value
    }
}

func deserialize_tree(node *TreeNode, s *bufio.Scanner) *TreeNode {
    // Create a TreeNode object if it doesn't exist already
    if node == nil {
        node = &TreeNode{}
    }

    // Check if another token exists in the file
    new_token := s.Scan()

    // Treat this node as "nil" if no token exists otherwise parse the token
    if !new_token {
        node.Value = "_"
    } else {
        node.Value = s.Text()
    }

    // If node isn't "nil, recurse and find it's children
    if node.Value != "_" {
        node.LeftChild = deserialize_tree(node.LeftChild, s)
        node.RightChild = deserialize_tree(node.RightChild, s)
    }

    // Return the pointer to the node
    return node
}

func evaluate_tree(node *TreeNode) int {
    // Try to convert the string value of the node into an integer
    node_value, err := strconv.Atoi(node.Value)
    var result int

    // If the node_value is an integer, result is the node value.
    if err == nil {
        result = node_value

    // Otherwise it is an operation, so recurse and solve the operation to find the node value
    } else {
        left_value := evaluate_tree(node.LeftChild)
        right_value := evaluate_tree(node.RightChild)

        perform_arthmitic(&result, node.Value, left_value, right_value)
    }

    // Finally return the result
    return result
}

func main() {
    // Ensure that the filename argument exist
    if len(os.Args) != 2 {
        fmt.Printf("Usage: %s [filename]\n", os.Args[0])
        os.Exit(1)
    }

    // Get the filename from the command line args
    flag.Parse()
    filename := flag.Arg(0)

    // Create a scanner to parse through the serialized tree
    f, err := os.Open(filename)
    check_err(err)
    scanner := bufio.NewScanner(bufio.NewReader(f))
    scanner.Split(bufio.ScanWords)

    // Deserialize the tree into the root TreeNode object
    root := TreeNode{}
    deserialize_tree(&root, scanner)
    f.Close()

    // Evaluate the tree and print the result
    result := evaluate_tree(&root)
    fmt.Printf("%d\n", result)
}
