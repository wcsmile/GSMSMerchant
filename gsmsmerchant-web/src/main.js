// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import "jquery"
import "bootstrap"
import Vue from 'vue'
import App from './App'
import router from './router'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import VueCookies from 'vue-cookies'
import dateCovert from './util/date'
import {
    get,
    post,
    patch,
    put,
    del
} from './util/http'
Vue.use(ElementUI)
Vue.use(VueCookies)

//定义全局变量
Vue.prototype.$get = get
Vue.prototype.$post = post;
Vue.prototype.$patch = patch;
Vue.prototype.$put = put;
Vue.prototype.$del = del;

import { EnumUtility } from '@/util/enum'
import EnumFilter from './util/filters'

Vue.prototype.EnumUtility = new EnumUtility() // 枚举字典
Vue.prototype.DateConvert = dateCovert // 日期格式转换
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
    el: '#app',
    router,
    components: { App },
    template: '<App/>'
})