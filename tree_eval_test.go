package main

import (
        "testing"
        "os"
        "bufio"
        "strings"
)

func assertNotPanic(t *testing.T, f func()) {
    defer func() {
        if r := recover(); r != nil {
            t.Errorf("Code paniced where not expected")
        }
    }()
    f()
}

func assertPanic(t *testing.T, f func()) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("Code did not panic where it should have")
        }
    }()
    f()
}

// "Integration" Tests
func TestTrees (t *testing.T) {
    testCases := []struct {
        Filename string
        Expected int
    }{
        {"trees/tree1.data", -2},
        {"trees/tree2.data", 22},
        {"trees/tree3.data", -12},
        {"trees/tree4.data", 42},
        {"trees/tree5.data", 42},
        {"trees/tree6.data", 42},
    }

    for _, testCase := range testCases {
        // Create a scanner to parse through the serialized tree
        f, err := os.Open(testCase.Filename)
        if err != nil {
            t.Errorf("Filename %s is missing", testCase.Filename)
            continue
        }
        scanner := bufio.NewScanner(bufio.NewReader(f))
        scanner.Split(bufio.ScanWords)

        // Deserialize the tree into the root TreeNode object
        root := TreeNode{}
        deserialize_tree(&root, scanner)
        f.Close()

        // Evaluate the tree and print the result
        result := evaluate_tree(&root)
        if result != testCase.Expected {
            t.Errorf("Serialized tree was not evaluated correctly, expected: '%d', got:  '%d'", testCase.Expected, result)
        }
    }
}

// Deserializer unit tests
func TestDeserializer (t *testing.T) {
    testCases := []struct {
        InputString string
    }{
        {"+ 1 _ _ * - 5 _ _ 6 _ _ 3 _ _"},  // Ok case
        {"+ 1 _ _ * - 5 _ _ 6 _ _ 3 _ _ \n \n"},

        {""},  // Empty
        {"\n\n\n\t\n"},

        {"+ 1"},  // Not complete tree
        {"+ + + + _ _"},
        {" _ _ 1 2 3 4 5"},
        {"1  2  3  4"},
    }

    for _, testCase := range testCases {
        scanner := bufio.NewScanner(strings.NewReader(testCase.InputString))
        scanner.Split(bufio.ScanWords)

        root := TreeNode{}
        assertNotPanic(t, func(){ deserialize_tree(&root, scanner) })
    }
}


// Evaluator unit tests
func TestEvaluator (t *testing.T) {
    testCases := []struct {
        InputString string
        Valid bool
    }{
        {"+ 1 _ _ * - 5 _ _ 6 _ _ 3 _ _", true},
        {"+ 1 _ _ * - 5 _ _ 6 _ _ 3", true},
        {"1 2 3 4", true},

        {"+ 1 _ _ * - 5 _ _ 6 _ 3 _ _", false},
        {"+ 1", false},
        {"+ + + + _ _", false},
        {" _ _ 1 2 3 4 5", false},

    }

    for _, testCase := range testCases {
        scanner := bufio.NewScanner(strings.NewReader(testCase.InputString))
        scanner.Split(bufio.ScanWords)

        root := TreeNode{}
        deserialize_tree(&root, scanner)

        if testCase.Valid {
            assertNotPanic(t, func(){ evaluate_tree(&root) })
        } else {
            assertPanic(t, func(){ evaluate_tree(&root) })
        }
    }
}
