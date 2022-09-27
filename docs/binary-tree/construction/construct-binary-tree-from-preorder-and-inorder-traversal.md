---
title: Construct Binary Tree from Preorder and Inorder Traversal
---

### 描述

Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

### 分析

1 build the value of the root of the tree according to the value of root node in the preorder
2 traveral the inorder:
2.1 build the left subnode of the root according to the slice value in the preorder and inorder
2.2 the index in the slice should not include the index of the current layer's root node
2.3 build the right subnode of the root according to the slice value in the preorder and inorder

### 代码

import Tabs from "@theme/Tabs";
import TabItem from "@theme/TabItem";

<Tabs
defaultValue="java"
values={[
{ label: 'Java', value: 'java', },
{ label: 'C++', value: 'cpp', },
]
}>
<TabItem value="java">

```java
// Construct Binary Tree from Preorder and Inorder Traversal
// 递归，时间复杂度O(n)，空间复杂度O(\logn)
public class Solution {
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        return buildTree(preorder, 0, preorder.length,
                inorder, 0, inorder.length);
    }

    TreeNode buildTree(int[] preorder, int begin1, int end1,
                       int[] inorder, int begin2, int end2) {
        if (begin1 == end1) return null;
        if (begin2 == end2) return null;

        TreeNode root = new TreeNode(preorder[begin1]);

        int inRootPos = find(inorder, begin2, end2, preorder[begin1]);
        int leftSize = inRootPos - begin2;

        root.left = buildTree(preorder, begin1 + 1, begin1 + leftSize + 1,
                inorder, begin2, begin2 + leftSize);
        root.right = buildTree(preorder, begin1 + leftSize + 1, end1,
                inorder, inRootPos + 1, end2);

        return root;
    }
    private static int find(int[] array, int begin, int end, int val) {
        for (int i = begin; i < end; ++i) {
            if (array[i] == val) return i;
        }
        return -1;
    }
}
```

</TabItem>
<TabItem value="cpp">

```cpp
// Construct Binary Tree from Preorder and Inorder Traversal
// 递归，时间复杂度O(n)，空间复杂度O(\logn)
class Solution {
public:
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        return buildTree(begin(preorder), end(preorder),
                begin(inorder), end(inorder));
    }

    template<typename InputIterator>
    TreeNode* buildTree(InputIterator pre_first, InputIterator pre_last,
            InputIterator in_first, InputIterator in_last) {
        if (pre_first == pre_last) return nullptr;
        if (in_first == in_last) return nullptr;

        auto root = new TreeNode(*pre_first);

        auto inRootPos = find(in_first, in_last, *pre_first);
        auto leftSize = distance(in_first, inRootPos);

        root->left = buildTree(next(pre_first), next(pre_first,
                leftSize + 1), in_first, next(in_first, leftSize));
        root->right = buildTree(next(pre_first, leftSize + 1), pre_last,
                next(inRootPos), in_last);

        return root;
    }
};
```

</TabItem>
</Tabs>

### 相关题目

- [Construct Binary Tree from Inorder and Postorder Traversal](construct-binary-tree-from-inorder-and-postorder-traversal.md)
