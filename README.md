# study5
linkedlist文件夹
1. [复杂链表复制][https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/] 文件名:copylist

   这里是深拷贝，先将每个复制的节点放在每个被复制节点的后面，然后将random指向原指针指向random节点的下一个节点，最后将两个链表分离开来就得到了复制好的链表

2. [合并K个升序链表](https://leetcode-cn.com/problems/merge-k-sorted-lists/)文件名：mergelist

   两两进行合并直到合并完k个升序链表，因为都是升序链表，都从第一个节点出发，两个节点两个节点的值进行比较，如果第一个值比第二个值小就传入第一个节点得下一个节点与第二个节点进行比较结合，一直递归直到链表走到尽头就能得到合并后的升序链表

3. [链表中倒数第k个节点](https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/) 文件名:getend

   使用递归走到最后一个节点返回当前节点及标志倒数第几个节点的count进行判断，这里不论找到与否count会不断的增加，最后会走到第一个节点返回找到的节点及count

4. [反转链表](https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/)文件名：reverselist

   使用一个中间指针、上一指针和下一指针一遍迭代完成，首先将现在节点的下一个节点保存，将现在节点的下一个节点设置为上一节点，将上一节点设置为当前这个节点，然后将当前节点设置为之前保存的下一节点，一遍递归之后就得到了反转后的链表

5. [从尾到头打印链表](https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/)文件名：reverseprint

   先迭代一遍链表得到链表长度来创建一个slice存储值，然后在迭代一遍链表，用得到的链表长度作为下标反向录入链表的值，这样就得到了一个从尾到头链表的slice
   workpool在workpool文件夹里
