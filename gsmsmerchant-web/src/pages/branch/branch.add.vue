<template>
        <bootstrap-modal ref="addModal" :need-header="true" :need-footer="true" :opened="showModalSetting" :closed="closeModal" size = "large">
            <div slot="title">
                <h4>创建门店</h4>
            </div>
            <div slot="body" >
                <el-form label-position="right" label-width="160px" status-icon :model="addInfo" ref="addForm" :rules="rules"  size ="mini" class="demo-ruleForm">
                    <el-row>
                        <el-col :span="21">
                            <div class="grid-content bg-purple">
                                <el-form-item prop="branch_name" label="门店名称:">
                                    <el-input placeholder="请输入门店名称"  v-model="addInfo.branch_name"  maxlength="20" style="width:300px"></el-input>
                                </el-form-item>
                                <el-form-item prop="contact_name" label="联系人姓名:">
                                    <el-input  placeholder="请输入联系人姓名"  v-model="addInfo.contact_name" maxlength="20" style="width:300px"></el-input>
                                </el-form-item>
                                <el-form-item prop="contact_tel" label="联系人电话:">
                                    <el-input  placeholder="请输入联系人电话"  v-model="addInfo.contact_tel" maxlength="20" style="width:300px"></el-input>
                                </el-form-item>
                                <el-form-item  label="所在区域:" required>
                                     <el-col :span="7">
                                        <el-form-item prop="province" style="margin-right:10px">
                                            <el-select  placeholder=" 请选择省 " v-model="addInfo.province"  size = "medium" @change="provinceSelect()">
                                                <el-option value=''>全部</el-option>
                                                <el-option v-for="(item, index) in provinceList" :key="index" :value="item.value" :label="item.name" ></el-option>
                                            </el-select>
                                        </el-form-item>
                                    </el-col>
                                    <el-col :span="7">
                                        <el-form-item prop="city" style="margin-right:10px">
                                            <el-select  placeholder="请选择市" v-model="addInfo.city"  size = "medium" @change="citySelect()">
                                                <el-option value=''>全部</el-option>
                                                <el-option v-for="(item, index) in cityList" :key="index" :value="item.value" :label="item.name" ></el-option>
                                            </el-select>
                                        </el-form-item>
                                    </el-col>
                                    <el-col :span="7">
                                        <el-form-item prop="district" style="margin-right:10px">
                                            <el-select  placeholder=" 请选择区 " v-model="addInfo.district"  size = "medium">
                                                <el-option value=''>全部</el-option>
                                                <el-option v-for="(item, index) in districtList" :key="index" :value="item.value" :label="item.name" ></el-option>
                                            </el-select>
                                        </el-form-item>
                                    </el-col>
                                </el-form-item>
                                <el-form-item prop="address" label="详细地址:">
                                    <el-input placeholder="请输入详细地址"  v-model="addInfo.address" type="textarea"  maxlength="256" :rows="3" style="width:517px"></el-input>
                                </el-form-item> 
                            </div>
                        </el-col>
                    </el-row>
                </el-form>
            </div>
            <div slot="footer">
                <el-button @click="cancelSubmit()">取消</el-button>
                <el-button type="primary" @click="submit()">确定</el-button>
            </div>
        
        </bootstrap-modal> 
</template>

<script>
export default {
    components:{
        'bootstrap-modal': require('vue2-bootstrap-modal')        
    },
    data(){
        let reg = /^[0-9]+([.]{1}[0-9]+){0,1}$/; //电话号码只能输入数字
        // 电话号码验证
        var validatorContactTel = (rule, value, callback)=>{
            if (value == "") {
                return callback(new Error('联系人电话不能为空'))
            }else if (!reg.test(value)) {
                return callback(new Error('只能输入整数'))
            }
            return callback()
        }
        return{
            branchStatusList: this.EnumUtility.Get("BranchStatus"),// 门店状态列表
            provinceList: [],// 省列表
            cityList: [],// 市列表
            districtList: [],// 区列表
            rules:{
                branch_name: [
                        { required: true, message: '门店名称不能为空', trigger: 'change' },
                    ],
                contact_name: [
                        { required: true, message: '联系人姓名不能为空', trigger: 'change' },
                    ],
                contact_tel: [
                        {required: true,validator:validatorContactTel, trigger: 'change' },
                    ],
                contact_name: [
                        { required: true, message: '联系人姓名不能为空', trigger: 'change' },
                    ],
                province: [
                        { required: true, message: '所在区域省不能为空', trigger: 'change' },
                    ],
                city: [
                        { required: true, message: '所在区域市不能为空', trigger: 'change' },
                    ],
                district: [
                        { required: true, message: '所在区域区不能为空', trigger: 'change' },
                    ],
                address: [
                        { required: true, message: '地址不能为空', trigger: 'change' },
                    ],
            },
            addInfo: {
                branch_name: "",
                contact_name:"",
                contact_tel: "",
                province: "",
                city: "",
                district: "",
                address: "",
            },   
        }
    },
    methods:{
        setModal(item){
            for(var key in this.addInfo){
                this.$set(this.addInfo,key,item[key])
            }

            this.$refs.addModal.open()
            this.getProvinceInfo()
        },
        showModalSetting(){
            this.$refs.addForm.clearValidate() //因为表单每次都会绑定值所以这里不用消除验证
        },
        closeModal(){
            for(var key in this.addInfo){
                this.addInfo[key] = ''              
            }
        },
        submit(){
            this.$refs['addForm'].validate((valid) => {
            if (!valid) {
                return false
            } else {
                this.$post('/branch/addbranch', this.addInfo).then(res=>{
                    if (res.data == "SUCCESS") {
                        this.$notify({
                            title:'创建门店',
                            message:'创建门店成功',
                            type:'success'
                        })
                        this.$emit('refresh-data')//刷新页面查询的响应
                        this.$refs.addModal.close()  
                    }else{
                        return
                    }
                })
                .catch(error=>{
                     if (error.response.status == 403) {
                        autoLogin(this);
                    } else if (error.response.status == 401) {
                        autoLogin(this);
                    } else if (error.response.status == 901) {
                        if (error.response.data.data) {
                            var err = ''
                            err = error.response.data.data.substr(6, error.response.data.data.length)
                            this.$notify.error({
                                title: '创建门店失败',
                                message: err
                            });
                        }
                    }else {
                        this.$notify.error({
                            title:'操作失败',
                        })
                    }
                })
                }
            })
        },
        cancelSubmit(){
            this.$refs.addModal.close()
        },
        // 获取省列表
        getProvinceInfo(){
            this.$post('/branch/queryprovince', this.addInfo).then(res=>{
                if (res.data != "") {
                    this.provinceList = res.data
                }else{
                    return
                }
            }).catch(error=>{
                this.$notify.error({
                    title:'获取省列表信息失败',  
                })
            })
        },
        // 根据省选择市
        provinceSelect(){
            
            this.cityList = []
            this.districtList = []
            this.addInfo.city =''    
            this.addInfo.district =''    
        
            this.$post('/branch/querycitybyprovince', {province:this.addInfo.province}).then(res=>{
                if (res.data!= "") {
                    this.cityList = res.data
                }else{
                    return
                }
            }).catch(error=>{
                this.$notify.error({
                    title:'获取市列表信息失败',  
                })
            })
            
        },
        // 根据市选择区
        citySelect(){
           
            this.districtList = []
            this.addInfo.district ='' 
        
            this.$post('/branch/querydistrictbycity',{city:this.addInfo.city}).then(res=>{
                if (res.data!= "") {
                    this.districtList = res.data
                }else{
                    return
                }
            }).catch(error=>{
                this.$notify.error({
                    title:'获取区列表信息失败',  
                })
            })
        
        },
    }
}
</script>

<style scoped>
.modal-lg {
	width: 650px;
}
</style>
