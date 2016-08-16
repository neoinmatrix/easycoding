#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define MAXSIZE 1000
typedef int ElemType;
#define RED    0
#define BLACK    1
typedef struct RBTNode
{
    char    color;
    ElemType    data;
    struct RBTNode *    p;
    struct RBTNode *    left;
    struct RBTNode *    right;
}RBTNode, * PRBTNode;

typedef struct RBTree
{
    PRBTNode    root;
    PRBTNode    nil;        //统一的空节点，该节点是黑的
}RBTree, * PRBTree;

int leftRotate (PRBTree tree, PRBTNode t);
int rightRotate (PRBTree tree, PRBTNode t);
PRBTNode insertRB (PRBTree tree, ElemType d);
int insertRB_fixup (PRBTree tree, PRBTNode t);
int createRBTree (PRBTree tree, ElemType d[], int n);
int initRB (PRBTree tree);
PRBTNode maximum (PRBTree tree, PRBTNode t);
PRBTNode minimum (PRBTree tree, PRBTNode t);
PRBTNode next (PRBTree tree, PRBTNode t);
PRBTNode precursor (PRBTree tree, PRBTNode t);
int walkNext (PRBTree tree);
int inOrderWalk (PRBTree tree, PRBTNode t);
int deleteRB_fixup (PRBTree tree, PRBTNode c);
PRBTNode deleteRB (PRBTree tree, PRBTNode t);
int main ()
{
    PRBTNode p;
    int d[MAXSIZE];
    int n = 0;
    int i;
    RBTree tree;
    initRB(&tree);
    /*
    insertRB(&tree, 11);
    insertRB(&tree, 2);
    insertRB(&tree, 14);
    insertRB(&tree, 1);
    insertRB(&tree, 7);
    insertRB(&tree, 15);
    insertRB(&tree, 5);
    insertRB(&tree, 8);
    insertRB(&tree, 4);
    */
    p=    insertRB(&tree, 26);
    insertRB(&tree, 17);
    insertRB(&tree, 41);
    insertRB(&tree, 14);
    insertRB(&tree, 21);
    insertRB(&tree, 30);
    insertRB(&tree, 47);
    insertRB(&tree, 10);
    insertRB(&tree, 16);
    insertRB(&tree, 19);
    insertRB(&tree, 23);
    insertRB(&tree, 28);
    insertRB(&tree, 38);
    insertRB(&tree, 7);
    insertRB(&tree, 12);
    insertRB(&tree, 15);
    insertRB(&tree, 20);
    insertRB(&tree, 3);
    insertRB(&tree, 35);
    insertRB(&tree, 39);
     

    srand(time(NULL));

    /*
    puts("请输入数据的个数：");
    scanf("%d",&n);
    printf("随机生成的%d个数据是：\n",n);
    for (i = 0; i < n; i++)
    {
        d[i] = rand()%1000;
        printf("%d  ",d[i]);
    }
    puts("");
    puts("建树开始");
    createRBTree(&tree, d, n);
    */

    inOrderWalk(&tree,tree.root);
    puts("");
    printf("根是%d \n",tree.root->data);
    
    printf("删除%d后：",p->data);
    deleteRB(&tree, p);
    
    inOrderWalk(&tree,tree.root);
    puts("");
    printf("根是%d \n",tree.root->data);
    return 0;
}
PRBTNode insertRB (PRBTree tree, ElemType d)
{//插入元素
//!!!记得插入的元素的初始化，p指向为父母节点，left和right赋值为NULL。
    PRBTNode t = NULL;
    PRBTNode p = NULL;
    int flag = 0;        //用来表示插入在左边的树还是右边的树
    t = tree->root;
    
    //插入的节点是root,并做相应的初始化
    if (tree->root == tree->nil)
    {
        tree->root = (PRBTNode)malloc(sizeof(RBTNode));
        tree->root->data = d;
        tree->root->color = BLACK;
        tree->root->p = tree->root->left =tree->root->right = tree->nil;
        
        return tree->root;
    }

    while (t != tree->nil)
    {
        p = t;
        if (d < t->data)
        {
            flag = 0;    
            t = t->left;
        }
        else 
        {
            if (d > t->data)
            {
                flag = 1; 
                t = t->right;
            }
            else
            {
                if ( (flag=rand()%2) == 0)
                    t = t->left;
                else
                    t = t->right;
            }
        }
    }//while

    //将t指向带插入节点的地址，并初始化
    t = (PRBTNode)malloc(sizeof(RBTNode));
    t->data = d;
    t->color = RED;
    t->p = p;
    t->left = t->right = tree->nil;
    
    if (!flag)
        p->left = t;
    else
        p->right = t;

    insertRB_fixup(tree, t);
    return t;
}

int insertRB_fixup (PRBTree tree, PRBTNode t)
{//插入的节点可能破坏红黑树的性质。该函数检测插入的节点是否破坏了红黑树的性质。如果破坏了，就对树进行调整，使其满足红黑树的性质
    while (t->p->color == RED)    //只有插入节点的父亲是红色的才会破坏红黑树的性质（4.如果一个结点是红的，那么它的俩个儿子都是黑的）
    {
        if (t->p->p->left == t->p)    //插入节点的父节点本身是left
        {
            if (t->p->p->right->color == RED)            //case 1
            {                                    
                t = t->p->p;
                t->left->color = t->right->color = BLACK;
                t->color = RED;
            }
            else
            {
                if (t->p->right == t)            //case 2
                {//将case 2转换为了case 3        
                    t = t->p;        //这步赋值是为了在转换为case 3时，t指向的是下面的红节点，和case 3的情况相一致
                    leftRotate(tree, t);
                }
                //case 3
                t->p->color = BLACK;
                t->p->p->color = RED;
                rightRotate(tree, t->p->p);
            }
        }//if
        else    //插入节点的父节点本身是right
        {
            if (t->p->p->left->color == RED)            //case 1
            {                                    
                t = t->p->p;
                t->left->color = t->right->color = BLACK;
                t->color = RED;
            }
            else
            {
                if (t->p->left == t)            //case 2
                {//将case 2转换为了case 3        
                    t = t->p;        //这步赋值是为了在转换为case 3时，t指向的是下面的红节点，和case 3的情况相一致
                    rightRotate(tree, t);
                }
                //case 3
                
                t->p->color = BLACK;
                t->p->p->color = RED;
                leftRotate(tree, t->p->p);
            }
        }//else
    }//while
    tree->root->color = BLACK;
    return 0;
}
int leftRotate (PRBTree tree, PRBTNode t)
{
    PRBTNode c;        //左旋，c指向t的right
    c = t->right;
    if (t->right == tree->nil) //左旋，t的right不能为空
        return 1;

    //这个if-else用于将t的父亲节点的left或right点指向c，如果t的父节点为不存在，则树的root指向c
    if (t->p != tree->nil)        //判断t是否为root
    {
        if (t->p->left == t)    //判断t是t的父节点的left还是right
            t->p->left = c;
        else
            t->p->right = c;
    }
    else
        tree->root = c;

    c->p = t->p;    //更新c的父节点

    t->right = c->left;
    if (c->left != tree->nil)
        c->left->p = t;
    c->left = t;
    t->p = c;
    return 0;
}
int rightRotate (PRBTree tree, PRBTNode t)
{
    PRBTNode c;        //右旋，c指向t的left
    c = t->left;
    if (t->left == tree->nil) //右旋，t的left不能为空
        return 1;

    //这个if-else用于将t的父亲节点的left或right点指向c，如果t的父节点为不存在，则树的root指向c
    if (t->p != tree->nil)        //判断t是否为root
    {
        if (t->p->left == t)    //判断t是t的父节点的left还是right
            t->p->left = c;
        else
            t->p->right = c;
    }
    else
        tree->root = c;

    c->p = t->p;    //更新c的父节点

    t->left = c->right;
    if (c->right != tree->nil)
        c->right->p = t;
    c->right = t;
    t->p = c;
    return 0;
}
int createRBTree (PRBTree tree, ElemType d[], int n)
{//用元素的插入建树
    int index = -1; 
    int tmp = -1;

    srand(time(NULL));
    
    while (n--)
    {
        index =(int) rand()%(n+1);//此时共有n+1个数据
        tmp = d[index];
        d[index] = d[n];
        d[n] = tmp;
        insertRB(tree, d[n]);
        printf("插入%d\t",d[n]);
    }
    puts("");
    return 0;
}//createRBTree

int initRB (PRBTree tree)
{//红黑树的初始化
    if (tree == NULL)
        return 0;
    tree->nil = (PRBTNode)malloc(sizeof(RBTNode));
    tree->nil->color = BLACK;
    tree->root = tree->nil;
    return 0;
}//initRB

PRBTNode minimum (PRBTree tree, PRBTNode t)
{//返回最小值，如果t是NULL返回NULL

    if (t == tree->nil)
        return NULL;
    while (t->left != tree->nil)
        t = t->left;
    return t;
}//minimum

PRBTNode maximum (PRBTree tree, PRBTNode t)
{//返回最大值，如果t是NULL返回NULL
    if (t == tree->nil)
        return NULL;
    while (t->right != tree->nil)
        t = t->right;
    return t;
}//maximum

PRBTNode next (PRBTree tree, PRBTNode t)
{//给出t的后继的节点。如果没有后继，就返回NULL
    PRBTNode p;        //指示父节点
    if (t->right == tree->nil)
    {
        p = t->p;
        while (p != tree->nil && p->right == t)
        {
            t = p;
            p = t->p;
        }
        return p;    //如果是最后一个元素，p的值为NULL
    }
    else
        return minimum(tree, t->right);
}//next

PRBTNode precursor (PRBTree tree, PRBTNode t)
{//返回节点t前驱，如果没有前驱，就返回NULL
    PRBTNode p;
    if (t->left == tree->nil)
    {
        p = t->p;
        while (p != tree->nil && p->left == t)
        {
            t = p;
            p = t->p;
        }
        return p;
    }
    else
        return maximum(tree, t->left);
}//precusor

int walkNext (PRBTree tree)
{//遍历二叉搜索树。先找到最小的元素，再通过用next()求后继来遍历树
    PRBTNode t;
    t = minimum(tree,tree->root);
    while (t != tree->nil)
    {
        printf("%d ",t->data);
        if (t->color == BLACK)
            printf("B\t");
        else
            printf("R\t");
        t = next(tree,t);
    }
    return 0;
}//walkNext

PRBTNode deleteRB (PRBTree tree, PRBTNode t)
{//删除数据。要求给处数据节点的指针
    PRBTNode c = NULL;        //c指向要取代被删除节点的子节点
    PRBTNode d = NULL;
    ElemType tmp;
    if (t == tree->nil)
        return NULL;

    //d指向真正要删除的元素的下标。如果t的left和right都有值，则转化为删除t的后继节点，并把后继节点的内容复制给t指向的节点。
    //而其他情况则直接删除t指向的节点
    if (t->left != tree->nil && t->right != tree->nil)
    {
        d = next(tree, t);
        //因为实际操作要删除的是d指向的节点，所以先交换data
        tmp = d->data;
        d->data = t->data;
        t->data = tmp;
    }
    else
        d = t;

    //确定c的指向
    if (d->left == tree->nil)
        c = d->right;
    else
        c = d->left;

    //将c的父亲指针设为d的父亲指针,c不会为空（因为存在nil节点）
    c->p = d->p;
    if (d->p != tree->nil)
    {
        if (d->p->left == d)
            d->p->left = c;
        else
            d->p->right = c;
    }
    else
        tree->root = c;

    if (d->color == BLACK)
        deleteRB_fixup(tree, c);
    return d;
}//deleteRB

int deleteRB_fixup (PRBTree tree, PRBTNode c)
{
    PRBTNode b;        //兄弟节点

    while (c != tree->root && c->color == BLACK)
    {
        if (c == c->p->left)
        {
            b = c->p->right;

            if (b->color == RED) //case 1  
            {//b节点是红的，可以说明c和b的父亲节点是黑的。通过以下的操作可以吧case 1转换为case 2,3,4中的一个
                b->color = BLACK;
                c->p->color = RED;
                leftRotate(tree, c->p);
                b = c->p->right;    //新的兄弟节点，这个节点一定是黑色的。这个节点之前是红色节点的儿子
            }
            if (b->right->color == BLACK && b->left->color == BLACK) //case 2
            {
                b->color = RED;        //将c的父节点的另一颗子树黑节点减少1
                c = c->p;            //将c上移。上移之后，c的黑高度相同了（因为另一颗子树的根节点有黑边为红）
            }
            else    //case 3或case 4
            {
                if (b->right->color == BLACK)        //case 3    通过以下操作将case 3 转化为case 4
                {
                    b->color = RED;
                    b->left->color = BLACK;
                    rightRotate(tree, b);
                    b = c->p->right;
                }
                //case 4
                //通过下面的操作，红黑树的性质恢复
                b->color = b->p->color;
                b->p->color = BLACK;
                b->right->color = BLACK;
                leftRotate(tree, c->p);
                c = tree->root;        //红黑树性质恢复，结束循环。不用break，是因为while结束后还要执行c->color = BLACK;
            }
        }//if (c == c->p->left)
        else
        {
            b = c->p->left;
            if (b->color == RED) //case 1  
            {//b节点是红的，可以说明c和b的父亲节点是黑的。通过以下的操作可以吧case 1转换为case 2,3,4中的一个
                b->color = BLACK;
                c->p->color = RED;
                rightRotate(tree, c->p);
                b = c->p->left;    //新的兄弟节点，这个节点一定是黑色的。这个节点之前是红色节点的儿子
            }
            if (b->right->color == BLACK && b->left->color == BLACK) //case 2
            {
                b->color = RED;        //将c的父节点的另一颗子树黑节点减少1
                c = c->p;            //将c上移。上移之后，c的黑高度相同了（因为另一颗子树的根节点有黑边为红）
            }
            else    //case 3或case 4
            {
                if (b->left->color == BLACK)        //case 3    通过以下操作将case 3 转化为case 4
                {
                    b->color = RED;
                    b->right->color = BLACK;
                    leftRotate(tree, b);
                    b = c->p->left;
                }
                //case 4
                //通过下面的操作，红黑树的性质恢复
                b->color = b->p->color;
                b->p->color = BLACK;
                b->left->color = BLACK;
                rightRotate(tree, c->p);
                c = tree->root;        //红黑树性质恢复，结束循环。不用break，是因为while结束后还要执行c->color = BLACK;
            }
        }//else
        
    }
    c->color = BLACK;
    return 0;
}//deleteRB_fixup

int inOrderWalk (PRBTree tree, PRBTNode t)
{//中序遍历
    if (t == tree->nil)
        return 0;
    putchar('(');
    inOrderWalk (tree, t->left);
    putchar(')');

    printf(" %d ",t->data);
    if (t->color == BLACK)
        printf("B");
    else
        printf("R");

    putchar('(');
    inOrderWalk (tree, t->right);
    putchar(')');
    return 0;
}