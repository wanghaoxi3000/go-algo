## 排序算法

![sort](./asset/sort.png)

### 关于时间复杂度：
1. 平方阶 (O(n2)) 排序 各类简单排序：直接插入、直接选择和冒泡排序。
2. 线性对数阶 (O(nlog2n)) 排序 快速排序、堆排序和归并排序；
3. O(n^(1+§)) 排序，§ 是介于 0 和 1 之间的常数。 希尔排序
4. 线性阶 (O(n)) 排序 基数排序，此外还有桶、箱排序。

### 关于稳定性：
1. 稳定的排序算法：冒泡排序、插入排序、归并排序和基数排序。
2. 不是稳定的排序算法：选择排序、快速排序、希尔排序、堆排序。

### 思考
> 选择排序和插入排序的时间复杂度相同，都是O(n^2)，在实际的软件开发中，为什么更倾向于使用插入排序呢？

从代码实现上来看，冒泡排序的数据交换要比插入排序的数据移动要复杂，冒泡排序需要3个赋值操作，而插入排序只需要1个，所以在对相同数组进行排序时，冒泡排序的运行时间理论上要长于插入排序。


### 动画演示
#### 冒泡排序
![bubble](./asset/bubble.gif)

#### 选择排序
![selection](./asset/selection.gif)

#### 插入排序
![selection](./asset/insertion.gif)

#### 快速排序
![quick](./asset/quicksort.gif)
