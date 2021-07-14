import Vue from 'vue'
export default new Vue()

/**
 * 全局事件通信总线
 * 呼叫中心
 * 作用： 可以让任何组件之间互相通信
 *
 * import Vue from 'vue'
 *
 * export defautl new Vue()
 *
 *  假设  a 组件要给 b 组件发送数据
 *
 *  b 注册通信的事件
 *  a 发布通信事件
 *
 *  各自加载 import global Bus from '@/utils/global-bus'
 *  注册globalBus.$on('自定义名称',(data) => { 处理函数 })
 *
 *  a 发布通信事件
 *  import globalBus from '@/utils/global-bus'
 *  a->呼叫中心->b
 *  globalBus.$emit('自定义名称',可选参数)
 *  a->b 呼叫自定义名称必须一致
 *  注意： 通信两端所使用的事件名称必须一致，否则是无效的
 *
 *   反之， 如果 b 要给 a 发送数据8
 *   那肯定就是 a on 注册时间， b emit 来发布事件
 */