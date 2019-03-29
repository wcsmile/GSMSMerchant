<template>
  <div>
    <nav-menu
      :menus="menus"
      :copyright="copyright"
      :themes="themes"
      :logo="logo"
      :systemName="systemName"
      :headpic="headpic"
      ref="NewTap"
      :userinfo="userinfo"
      :pwd="pwd"
    >

    <keep-alive>
      <router-view v-if="$route.meta.keepAlive" @addTab="addTab" @close="close" @setTab="setTab" ></router-view>
    </keep-alive>
    <router-view v-if="!$route.meta.keepAlive" @addTab="addTab" @close="close" @setTab="setTab"></router-view>

      <!-- <router-view  @addTab="addTab" /> -->
    </nav-menu>

    <el-popover
      placement="left"
      title=""
      width="260"
      trigger="hover">
      <el-table :data="gridData">
        <el-table-column property="content" label="账户信息"></el-table-column>
      </el-table>
     
      <el-button class="btn-username" slot="reference">{{userinfo.login_account}}</el-button>
    </el-popover>

    <el-dialog title="修改密码" width="30%" :visible.sync="dialogAddVisible">
      <el-form :model="updateInfo" :rules="rules" ref="addForm">

        <el-form-item label="请输入原密码" prop="password_old">
          <el-input type="password" v-model="updateInfo.password_old"  ></el-input>
        </el-form-item>

        <el-form-item label="请输入新密码" prop="password">
          <el-input type="password" v-model="updateInfo.password"  ></el-input>
        </el-form-item>

        <el-form-item label="请确认密码" prop="checkPass">
          <el-input type="password" v-model="updateInfo.checkPass"  ></el-input>
        </el-form-item>

      </el-form>
      <div slot="footer" >
        <button class="btn btn-sm btn-primary" @click="resetForm(`addForm`)">取 消</button>
        <button class="btn btn-sm btn-danger"  @click="add(`addForm`)">确 定</button>
      </div>

    </el-dialog>

  </div>
</template>

<script>
  
  import navMenu from 'nav-menu'; // 引入
  import VueCookies from 'vue-cookies'
  export default {
    name: 'app',
    data () {

      var validatePass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入密码'));
        } else {
          if (this.updateInfo.checkPass !== '') {
            this.$refs.addForm.validateField('checkPass');
          }
          callback();
        }
      };
      var validatePass2 = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请再次输入密码'));
        } else if (value !== this.updateInfo.password) {
          callback(new Error('两次输入密码不一致!'));
        } else {
          callback();
        }
      };

      return {
        headpic:"http://sso.sinopecscsy.com/static/img/a0.jpg",
        logo:"http://sso.sinopecscsy.com/static/img/43612a9fe1f92658cc3bc6e3edc0766e.png",
        copyright:"2019 四川千行你我科技有限公司", //版权信息
        themes:"bg-primary|bg-primary|bg-dark", //顶部左侧背景颜色,顶部右侧背景颜色,右边菜单背景颜色
        menus:[{}],  //菜单数据
        systemName:"芯片卡商户系统",  //系统名称
        userinfo: {},
        dialogAddVisible:false,     //添加表单显示隐藏
        updateInfo:{
          password_old: "",
          password: "",
          checkPass: "",
        },
        gridData: [],
        rules: {                    //数据验证规则
          password_old: [
            { required: true, message: "请输入原密码", trigger: "blur" }
          ],
          password: [
            { required: true, message: "请输入新密码", trigger: "blur" },
            { validator: validatePass, trigger: 'change' }
          ],
          checkPass: [
            { required: true, message: "请确认密码", trigger: "blur" },
            { validator: validatePass2, trigger: 'change' }
          ],
        },

      }
    },
    components:{ //注册插件
      navMenu
    },
    created(){
      this.queryDictionaryList();
      this.getMenu();
    },
    mounted(){
      this.getSysInfo()
      this.userinfo = VueCookies.get("userinfo");
      this.gridData = [
        {"content": "账号名称：" + this.userinfo.login_account},
        {"content": "本次登录：" + this.userinfo.current_time},
        {"content": "登录地址：" + (this.userinfo.last_login_ip ? this.userinfo.last_login_ip : "---")},
        {"content": "上次登录：" + (this.userinfo.last_login_time ? this.userinfo.last_login_time : "---")},
      ]                          
      //向组件添加一个自定义标签，点击标签会路由到一个页面
      this.$refs.NewTap.add("采购管理","/purchase/list", {});   //设置默认页面
    },
    methods:{
      getSysInfo(){
        this.$get("/sys/get")
            .then(res=>{
              this.headpic = res.default_pic,
              this.logo = res.sys_logo,
              this.copyright = res.sys_copy, //版权信息
              this.themes = res.sys_themes, //顶部左侧背景颜色,顶部右侧背景颜色,右边菜单背景颜色
              this.systemName = res.sys_name,  //系统名称
              document.title = res.sys_name
            })
            .catch(error=>{
               console.log(error)
            })
      },
      pwd(val){
        this.dialogAddVisible = val;
      },
      resetForm(formName) {
        this.dialogAddVisible = false;
        this.$refs[formName].resetFields();
      },
      getMenu(){
          this.$get("/member/menu/get")
            .then(res=>{
              console.log("res:",res)
              this.menus = res
            })
            .catch(error=>{
               console.log(error)
            })
      },
      add(formName){
        // console.log(this.addData)
        var params = {
          password_old : this.updateInfo.password_old,
          password : this.updateInfo.password,
          passwords: this.updateInfo.checkPass
        }
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.$post("/member/update/pwd", params).then(res=>{
              this.$notify({
                title:'成功',
                message:'修改操作完成',
                type:'success'
              });
              this.dialogAddVisible = false;
              this.$refs[formName].resetFields();
            }).catch(errro=>{
              this.$notify({
                title:'失败',
                message:"原密码错误或密码修改次数超过限制",
                type:'error'
              });
              this.$refs[formName].resetFields();
            })
          } else {
            console.log(`error submit!!`);
            return false;
          }
        });
      },
      //@name 标签名称
      //@path 路由
      //@obj  路由参数 类型：Object
      addTab(name,path,obj){
        this.$refs.NewTap.add(name,path,obj);   //调用组件方法，添加一个页面
      },
      close(v){
         this.$refs.NewTap.closeTab(v);   
      },
      setTab(name,path,obj){
        console.log("outer",name,path,obj);
        this.$refs.NewTap.set(name,path,obj);
      },

      queryDictionaryList(){
        var serviceApi = '/dictionary/query'
        this.$get(serviceApi, {})
            .then(res=>{
                this.EnumUtility.Set(res.list)
            })
            .catch(error=>{
                this.$notify.error({
                    title:'操作失败',
                })
            })
      }
    }
  }
</script>
<style>
#app>div>span{
  position: absolute;
  top: 6px;
  right: 70px;
  z-index: 99999;
}
.btn-username{
  background: transparent;
  border-color: transparent;
  color: #fff;
}
</style>
