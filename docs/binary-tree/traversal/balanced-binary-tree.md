---
title: Balanced Binary Tree
---

### 描述

Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

### 分析

1 postorder traversal
2 depth function is the process that is plus one count before trace back to previous layer
3 judge condition is the differences between the left subtree node and right subtree by no more than one

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
// Balanced Binary Tree
// 时间复杂度O(n)，空间复杂度O(logn)
public class Solution {
    public boolean isBalanced (TreeNode root) {
        return balancedHeight (root) >= 0;
    }

    /**
     * Returns the height of `root` if `root` is a balanced tree,
     * otherwise, returns `-1`.
     */
    private static int balancedHeight (TreeNode root) {
        if (root == null) return 0;  // 终止条件

        int left = balancedHeight (root.left);
        int right = balancedHeight (root.right);

        if (left < 0 || right < 0 || Math.abs(left - right) > 1) return -1;  // 剪枝

        return Math.max(left, right) + 1; // 三方合并
    }
}
```

</TabItem>
<TabItem value="cpp">

```cpp
// Balanced Binary Tree
// 时间复杂度O(n)，空间复杂度O(logn)
class Solution {
public:
    bool isBalanced (TreeNode* root) {
        return balancedHeight (root) >= 0;
    }

    /**
     * Returns the height of `root` if `root` is a balanced tree,
     * otherwise, returns `-1`.
     */
    int balancedHeight (TreeNode* root) {
        if (root == nullptr) return 0;  // 终止条件

        int left = balancedHeight (root->left);
        int right = balancedHeight (root->right);

        if (left < 0 || right < 0 || abs(left - right) > 1) return -1;  // 剪枝

        return max(left, right) + 1; // 三方合并
    }
};
```

</TabItem>
</Tabs>
