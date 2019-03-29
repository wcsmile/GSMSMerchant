<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
  import loginWithUp from 'login-with-up'; // 引入
  export default {
    name: 'App',
    data() {
      return {
        host: "http://api.test.cn/login",
        sys_name: "用户权限系统",
        copyright: "版权信息",
        ident: "al" //系统英文名称
      }
    },
    components: { //注册插件
      loginWithUp
    },
    methods: {
      call(e) {
        console.log(e)
      }
    }
  }

</script>

<style>

</style>
