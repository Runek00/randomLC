/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public boolean isBalanced(TreeNode root) {
        return isBalancedAndHeight(root).balanced();
    }
    
    record BH(boolean balanced, int height){}

    BH isBalancedAndHeight(TreeNode root) {
        if (root == null) {
            return new BH(true, 0);
        }
        var lBH = isBalancedAndHeight(root.left);
        var rBH = isBalancedAndHeight(root.right);
        if (!lBH.balanced() || !rBH.balanced() || Math.abs(lBH.height() - rBH.height()) > 1){
            return new BH(false, 1);
        }
        var h = Math.max(lBH.height(), rBH.height()) + 1;
        return new BH(true, h);
   }



}