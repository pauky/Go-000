## 学习笔记
### 作业
#### 题目
参考 Hystrix 实现一个滑动窗口计数器

#### 作答
参考[kratos-metrics](https://github.com/go-kratos/kratos/tree/3cc8d9412681de1e1a1d99629c08c0115ed768e5/pkg/stat/metric)，去除了gauge的支持及其它计数无关的方法，实现了滑动窗口的计数器。
