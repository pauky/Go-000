## 学习笔记

### goroutine避免泄漏
1. 要管控go的生命周期：
    * 使用go需要知道它啥时候退出
	* 用超时、chan、context保证它可能堵塞过久
	* 并发行为需要交给调用者控制

注: log.Fatal建议只在main函数里使用，因为它底层用了os.exit，不会触发defers，有点像panic，相当于退出
### 内存模型
1. 内存重排，CPU、编译器可能会使指令不按程序编写的顺序执行，且CPU缓存的存在，多核心的缓存是独立的，所以可能只写到缓存，没有写到公共的内存上，产生内存重排现象。
2. 多协程并发修改共享变量的时候，需要避免这种内存重排导致的问题，使用同步事件（sync, chan）保证happen-before。
3. 单个machine word才是原子操作，即一条汇编指令就是一个原子操作，比如在32bit机器赋值一个64位的值，需要两条指令，所以它不是原子的。

### sync
1. 如何判断race：是否在goroutine中进行并发使用非原子操作共享数据；
2. 使用同步语义解决race：Mutex, RWMutex, Atomic
    * Mutex比Atomic重，因为涉及goroutine的切换
	* copy-on-write，可以解决读多写少的内存共享问题，可以用atomic实现它。
	* Mutex锁的实现：
	    * Barking：锁释放时唤醒第一个等待的人，提高吞吐量
		* Handsoff：均衡锁分配，解决饥饿问题；
		* Spinning：自旋尝试获得锁，加快加锁效率；
        * 1.8前使用Barking、 Spinning，有饥饿问题，之后使用Handsoff，并在触发handsoff后取消Spinning；
	* errGroup：goroutine并发控制流；
	* sync.Pool：适合复用临时对象的场景，减少内存分配次数

### channel
1. 安全消息队列，官方建议使用channel进行内存共享而不是同步语义；
2. 种类
    * 缓冲通道，有缓冲长度，发送是异步执行的，发先于收；超大的缓冲长度，超小保障能使数据到达
	* 无缓冲通道，发和收是同步的，收先于发；
### context
1. 用于goroutine管理嵌套其goroutine的生命周期，例如一个操作的goroutine需要多个rpc请求，它开了多个goroutine去处理请求，如果操作取消，需要将其下的goroutine也取消；
2. context能够记录嵌套goroutine的关系，以便支持统一管理生命周期；
3. 使用姿势：
    * 作为方法的第一参数或可选参数
    * 不要将其放在结构体中，对使用者很不友好
    * 数据应该是只读的，有新值加入的话，用copy-on-wirte的方式；
4. 原理：WithValue，沿当前节点到父节点方向递归查找；

### 毛老师的总结：
goroutine go关键字一定要关注:
1、go 生命周期（结束、终止）
2、go panic
3、把并行扔给调用者

内存模型：
1、搞清楚 原子性、可见性
2、go memory model（了解happen-before）
3、底层的 memory reordering（可以挖一挖 cpu cacline、锁总线、mesi、memory barrier）