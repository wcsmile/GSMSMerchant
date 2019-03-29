import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/pages/login/login.vue'
import Menu from '@/pages/menu/menu.vue'
import OrderList from '../pages/order/list.vue'
import BranchList from '../pages/branch/branch.list.vue'
import PurchaseList from '../pages/purchase/list.vue'
import ProductCardList from '../pages/product/card.list.vue'
import ProductRechargeList from '../pages/product/recharge.list.vue'
import CapitalrecordList from '../pages/capitalrecord/capitalrecord.vue'

Vue.use(Router)

export default new Router({
    mode: 'history',
    routes: [{
        path: '/',
        name: 'menu',
        component: Menu,
        children: [{
            path: '/order/list',
            name: 'OrderList',
            component: OrderList
        },
        {
            path: '/branch/list',
            name: 'BranchList',
            component: BranchList
        },
        {
            path: '/purchase/list',
            name: 'PurchaseList',
            component: PurchaseList
        },
        {
            path: '/product/card/list',
            name: 'ProductCardList',
            component: ProductCardList
        },
        {
            path: '/product/recharge/list',
            name: 'ProductRechargeList',
            component: ProductRechargeList
        },
        {
            path: '/capitalrecord/list',
            name: 'CapitalrecordList',
            component: CapitalrecordList
        }
        ]
    },
    {
        path: '/login',
        name: 'login',
        component: Login
    }
    ]
})