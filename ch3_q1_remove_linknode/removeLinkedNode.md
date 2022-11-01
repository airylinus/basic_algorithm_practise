# Remove specific from linked list

## Question

- Input will be a linked list which contains may Nodes.
  Node may look like this:
  ```
  type LinkNode struct {
      Val   int
      Next  *LinkNode
  }
  ```
- We assumed every node contained in input list has different
  Val property.

- Write code remove the node which has the exactly Val property
  as input target.

  ```
  func removeLinkNode(nodes *LinkNode, target int) {
      //your code here
  }
  ```
- For example:
  While target is `3` and Input list like: `1 -> 2 -> 3 -> 4`, the correct 
  answer should be `1 -> 2 -> 4`
  
## Thinking

- To remove target `3`, the first step is found it. Then we make the previous
  Node's Next pointed to the target's Next.
  ```
  previousNodeOf(target).Next = target.Next
  ```
  Target node will be deleted as the whole linked list jump over it.
  We can implement this approach as blow:
  ```
  func removeLinkNode(head *common.LinkNode, targetVal int) (newHead *common.LinkNode) {
      p := head
      for p != nil && p.Next != nil {
          p.Next.Val == targetVal {
              ext = p.Next.Next
          }
          p.Next
      }
      return head
  }
  ```

## Disadvantage and improvement
  
  If target is the first / last node of linked node, the early approach will fail.
  Problem is it's hard to refer to previous node when we found target node by value, cuz 
  there is no property such as 'Previous'.

  As the value of node is easy to reset, we can easily swap two node's values. We may
  solve the problem by these steps:
  - swap values of target node and target's next node

    ``` targetNode.Val = targetNode.Next.Value```

  - delete target's next node

    ```targetNode.Next = targetNode.Next.Next```

  
  
  