// https://www.mta.ca/~rrosebru/oldcourse/263114/Dsa.pdf

O(1) constant:
    the operation doesn't depend on the size of its input, e.g. adding
    a node to the tail of a linked list where we always maintain a pointer to
    the tail node.

O(n) linear:
    the run time complexity is proportionate to the size of n.

O(log n) logarithmic:
    normally associated with algorithms that break the problem
    into smaller chunks per each invocation, e.g. searching a binary search
    tree.

O(n log n) just n log n:
    usually associated with an algorithm that breaks the problem
    into smaller chunks per each invocation, and then takes the results

O(n2) quadratic: e.g. bubble sort.

O(n3) cubic: very rare.

O(2n) exponential: incredibly rare.

链表 (Linked List)
  单链表(Singly Linked List)
    插入 (Insertion)
    查找 (Searching)
    删除 (Deletion)
    遍历 (Traversing the list)
    反向遍历 (Traversing the list in reverse order)
  双向链表 (Doubly Linked List)
    插入 (Insertion)
    删除 (Deletion)
    反向遍历 (Reverse Traversal)

二叉查找树 (Binary Search Tree, BST)
  左子树的节点小于根结点，右子树的节点大于或等于跟节点
  最坏情况二叉查找树会退化成链表
  递归 (Recursive)
  插入 (Insertion)
  查找 (Searching)
  删除 (Deletion)
  查找父节点 (Finding the parent of a given node)
  获得节点的引用 (Attaing a reference to a node)
  查找最大值和最小值 (Finding the smallest and largest values in the BST)
  遍历 (Tree Traversals)
      深度优先 (Depth First)
      先序 (Preorder)
      后序 (Postorder)
      中序 (Inorder)
    广度优先 (Breath Fisrt) / 队列 (Queue)

堆 (Heap)
  二叉堆 (Binary Heap) / 数组 (Array)
  小堆 (Min Heap) / 大堆 (Max Heap)
  插入 (Insertion)
  删除 (Deletion)
  查找 (Searching)
  遍历 (Traversal)

集合 (Set)
  无序集合 (Unordered)
    哈希表 (hash table)
  有序集合 (Ordered)
    BST / AVL Tree

队列 (Queue)
  先进先出 (First In First Out, FIFO)
  入队 (Enqueue): places an item at the back of the queue;
  出队 (Dequeue): retrieves the item at the front of the queue, and removes it from the queue;
  Peek: retrieves the item at the front of the queue without removing it from the queue

  标准队列 (standard queue)
    单向链表 (singly list)
  优先队列 (Priority Queue)
    堆 (heap)
  双端队列 (Double Ended Queue, deque)
    双向链表 (doubly list)

AVL Tree
  // https://en.wikipedia.org/wiki/AVL_tree
  // https://zh.wikipedia.org/wiki/AVL%E6%A0%91
  // https://medium.com/@sarahzhao25/avl-trees-where-to-find-rotate-them-7b062e0a30f8
  // https://www.tutorialspoint.com/data_structures_algorithms/avl_tree_algorithm.htm
  平衡因子(balance factor): -1, 0, 1
  树的旋转(Tree Rotations):
    右旋(Right Rotation)
    左旋(Left Rotation)
    左-右旋(Left Right Rotation)
    右-左旋(Right left Rotation)
  自平衡(Tree Rebalancing): LeftRotation, RightRotation, LeftAndRightRotation, RightAndLeftRotation
  插入(Insertion)
  删除(Deletion)

树 (Tree)
  https://www.quora.com/What-is-the-height-size-and-depth-of-a-binary-tree
  https://stackoverflow.com/questions/2603692/what-is-the-difference-between-tree-depth-and-height
  节点的高(height): 节点到叶子节点的路径的最大边的数目
    只有一个节点的高为 0，空节点的高为 -1
  树的高(height of binary tree): 根结点的高
  节点的深度(depth): 根结点到节点路径的边的数目
  树的深度(depth of the root node): 根结点的深度
  层级(level)

// https://codefarm.me/2019/03/21/mathematical-induction-and-recusion/
