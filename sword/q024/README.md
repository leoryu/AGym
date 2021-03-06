# 二叉搜索树的后序遍历序列

## 问题

>输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则输出Yes,否则输出No。假设输入的数组的任意两个数字都互不相同。

## 思路

>结合二叉搜索树的特性 步骤如下：
>
>- 如果数组L长度小于等于2，直接返回true
>- 把数组L按照原来顺序切分成 A = [...], B = [...], C = [len(L)-1]。从倒数第二个数开始，向前找，从第一个小于C的元素开始，往前都是A。中间是B，肯定都是大于C的数。
>- 如果A中，发现还有元素大于C，那么说明L不可能是某二叉搜索树的后序遍历的结果，直接返回false。
>- 反之，递归处理切分A和B

## 特殊测试用例

- 空树
