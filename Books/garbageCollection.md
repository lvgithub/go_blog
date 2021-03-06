GC（垃圾回收）必须Stop-the-world？

并发编程的许多困难都源于对象生存期问题，当对象在线程之间传递时，要确保它们安全地释放就变得很麻烦。因此GC可以使得并发编程变得容易。但是实GC也是一个挑战，但是一次实现，就可以解决人们手动管理内存的麻烦（C语言），大大提高的开发效率和避免了许多Bug。

但是GC也是有成本的，他会影响程序的效率，GC是一个非常挑战的工作，很多计算机科学家在上面耗费了数十年不断的提升效率。

GC算法设计时，会考虑几个重要指标：
* 程序吞吐量：GC对程序效率的影响，也就花费在GC的时间和程序处理正常业务的时间比；
* GC吞吐量：单位时间内垃圾回收的数量；
* 暂停时间：Stop-the-world 的时间；
* 并发：垃圾回收机制如何使用多核；
* 等等还有很多

很多人问为什么GC的时候要暂停（Stop-the-world）整个程序，为什么不能并发的执行GC呢？GC本质上是一种权衡，Stop-the-world 是为了GC吞吐量（在给定CPU时间内多少垃圾可以被收集器清除？），便不是说GC必须STW,你也可以选择降低运行速度但是可以并发执行的收集算法，这取决于你的业务。

J9VM中[标记阶段](https://www.ibm.com/support/knowledgecenter/zh/SSYKE2_8.0.0/com.ibm.java.vm.80.doc/docs/mm_gc_mark_parallel.html)有描述，标记分为：
* 并行标记
  ```
    并行标记的目的是在不降低单处理器系统上标记性能的情况下，提高多处理器系统上典型标记的性能。

    通过增加共享使用工作包池的助手线程，可以提高对象标记的性能。例如，可选取由一个线程返回给池的完整输出包作为另一个线程的新输入包。

    并行标记仍需要一个用作主协调代理进程的应用程序线程的参与。助手线程帮助标识回收的根指针并跟踪这些根。标记位是使用不需要附加锁的主机原子原语来更新的
  ```
* 并发标记
  ```
    在堆大小增加时，并发标记能够提供缩短且一致的垃圾回收暂停时间。

    在堆满之前，GC 将启动并发标记阶段。在并发阶段，GC 扫描堆，检查根对象，比如堆栈、JNI 引用和类静态字段。通过要求每个线程扫描自己的堆栈来扫描堆栈。随后，这些根将用于并发跟踪活动对象。在线程执行堆锁分配时，跟踪由低优先级的后台线程和每个应用程序线程执行。

    当 GC 利用正在运行的应用程序线程并发标记活动对象时，必须记录对已跟踪对象的任何更改。它使用在每次更新对象中的引用时运行的写屏障。在发生对象引用更新时，写屏障将使用标志。使用该标志迫使对部分堆重新扫描。
  ```

比如:你做金融交易类的项目，分秒必争，那可以选择并行的方式。如果你是一种后台任务，比如数据处理，那你可以选择STW类型算法，使 GC 的吞吐量得到最高。

两类算法最终的权衡指标就GC效率：程序工作时间与执行收集时间的比率。

没有单一的算法在所有方面都完美，语言也不可能知道程序的业务类型，这也就是“GC调优”存在的原因。这也是科学的基础规律。


