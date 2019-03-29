<template>
  <div id="app">
    <login-with-up
      :copyright="copyright"
      :systemName="systemName"
      :conf="conf"
      :call="call"
      ref="loginItem"
      :err-msg.sync="errMsg"
      :requireCode="requireCode"
      :getCodeCall="getCodeCall"
    >

    </login-with-up>
  </div>
</template>

<script>
  import VueCookies from 'vue-cookies'
  import loginWithUp from 'login-with-up'; // 引入
  export default {
    name: 'login',
    data () {
      return {
        errMsg:{},
        systemName: "芯片卡商户系统",
        copyright:"四川千行你我科技有限公司Copyright© 2019 版权所有",
        conf:{loginNameType:"请输入用户名",pwd:"请输入密码", validateCode:"请输入微信验证码"},   //输入框提示信息配置
        requireCode: false,
      }
    },
    components:{ //注册插件
      loginWithUp
    },
    mounted(){
      document.title = "用户登录"
      VueCookies.remove("__jwt__")
      sessionStorage.removeItem("__jwt__")
    },
    methods:{
      call(e){
        //在这里获取数据进行登录
        this.errMsg = {message:"登录中。。。。",timestamp:  Date.parse(new Date())};

        this.$get("/member/login",e)
            .then(res=>{
              res.current_time = this.getCurrentTime()
              VueCookies.set("userinfo", res)
              setTimeout(() => {
                this.$router.push("/")
              }, 500);

            })
            .catch(error=>{
              console.log("error:",error)
              if (error.response.status == 406){
                this.errMsg = {message:"用户账号不存在",timestamp:  Date.parse(new Date())}
              }else if (error.response.status == 900){
                this.errMsg = {message:"用户密码错误",timestamp:  Date.parse(new Date())}
              }else{
                this.$refs.loginItem.showMsg("系统繁忙")
              }
            })
      },
      // getCodeCall(e){
      //    e.ident="al"
      //    this.$refs.loginItem.showMsg("发送中.....");
      //    this.$get("/member/getcode", e)
      //       .then(res=>{
      //         this.$refs.loginItem.showMsg("微信验证码发送成功");

      //       })
      //       .catch(error=>{
      //         if (error.response.status == 403 || error.response.status == 406){
      //           this.$refs.loginItem.showMsg("用户信息获取失败");
      //         }else{
      //           this.$refs.loginItem.showMsg("系统繁忙")
      //         }
      //       })
      // },
      //毫秒数转换成时间
      getCurrentTime(milliseconds){
          var myDate = new Date();
          var year = myDate.getFullYear();
          var month = myDate.getMonth() + 1;
          var day = myDate.getDate()
          var hour = myDate.getHours();
          var minute = myDate.getMinutes();
          var second = myDate.getSeconds();

          month = this.checkTime(month).toString();
          day = this.checkTime(day).toString();
          hour = this.checkTime(hour).toString();
          minute = this.checkTime(minute).toString();
          second = this.checkTime(second).toString();

          return year+"-"+month+"-"+day+" "+hour+":"+minute+":"+second;
      },
      // 只有一位数字时添加“0”
      checkTime(i){
          if(i < 10){
              i = "0" + i;
          }
          return i;
      }
    }
  }
</script>
