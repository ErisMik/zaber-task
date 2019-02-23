# zaber-task

## Description:
```
The program takes filepath as an command line argument.
During execution the program loads the file and evaluates the tree in the file. Then it prints the result.
For example for a tree in the attachment:

./tree_eval my_tree.data
# Result: -2
```

## Assumptions:
- Operations are limited to `+`, `-`, `*` and `/`
- Values are integer numbers
- Serialized trees are valid
- Serialized trees are serialized such taht they folllow order of operations


## Serialization Examples:
Serialized trees exist as a preorder traversal of the tree.
They are on a single line, with a space seperating values.
Null / Empty nodes are denoted by a `_`
```
tree1.data: 1+(5-6)*3 = -2
tree2.data: 1*(8+3)*2 = 22
tree3.data: 2*3+(5-10)*6+12 = -12

tree4.data: 7*2+7*2+7*2 = 42
tree5.data: 7*3*2 = 42
tree6.data: 14*3 = 42
```
