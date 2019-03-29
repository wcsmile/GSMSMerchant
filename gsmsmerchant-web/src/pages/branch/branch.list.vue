
<template>
    <div class="panel panel-default">
        <div class="panel panel-default">
            <div class="panel-body">
                <el-form ref="form" :model="params" :inline="true" class="form-inline pull-left add-qx-bottom">
                    <el-form-item>
                        <el-select  placeholder=" 请选择状态 " v-model="query.status"  size = "medium" style="width: 160px">
                            <el-option value=''>全部</el-option>
                            <el-option v-for="(item, index) in branchStatusList" :key="index" :value="item.value" :label="item.name" ></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item>
                        <el-select  filterable placeholder=" 请选择门店 " v-model="query.branch_id"  size = "medium" style="width: 160px">
                            <el-option value=''>全部</el-option>
                            <el-option v-for="(item, index) in branchList" :key="index" :value="item.branch_id" :label="item.branch_name" ></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item>
                        <a class="btn btn-success " @click="refreshDataList" style="margin-bottom:3px;">查询</a> 
                    </el-form-item >
                    <el-form-item>
                        <a class="btn btn-primary " @click="add" style="margin-bottom:3px;">创建门店</a> 
                    </el-form-item >
                </el-form>
            </div>
        </div>

        <div class="table-responsive">
            <el-table :data="dataList"  style="width:100%"  border height="532px">
                <el-table-column prop="branch_id" align="center" label="门店编号" width="100px">
                </el-table-column>
                <el-table-column prop="branch_name" align="center" label="门店名称" width="250px" >
                </el-table-column>
                <el-table-column prop="contact_name" align="center" label="联系人姓名" >
                    <template slot-scope="scope">
                        <span style="text-align:center">{{scope.row.contact_name | StringFilter(scope.row.contact_name) }}</span>
                    </template>
                </el-table-column>
                <el-table-column prop="contact_tel" align="center" label="联系人电话" >
                    <template slot-scope="scope">
                        <span style="text-align:center">{{scope.row.contact_tel | StringFilter(scope.row.contact_tel) }}</span>
                    </template>
                </el-table-column>
                <el-table-column prop="chipcard_count" align="center" label="油卡数量" >
                    <template slot-scope="scope">
                        <span style="text-align:center">{{scope.row.chipcard_count | StringFilter(scope.row.chipcard_count) }}</span>
                    </template>
                </el-table-column>
                <el-table-column prop="recharge_count" align="center" label="复充值油卡数量" >
                    <template slot-scope="scope">
                        <span style="text-align:center">{{scope.row.recharge_count | StringFilter(scope.row.recharge_count) }}</span>
                    </template>
                </el-table-column>
                <el-table-column prop="cumulative_standard" align="center" label="截止昨日累计充值金额（元）" width="280px">
                    <template slot-scope="scope">
                        <span style="text-align:center">{{scope.row.cumulative_standard | StringFilter(scope.row.cumulative_standard) }}</span>
                    </template>
                </el-table-column>
                <el-table-column prop="status" align="center" label="门店状态" >
                    <template slot-scope="scope">
                       <el-button :disabled="scope.row.status == '0'" type="text" class="text-info" @click="statusChange(scope.row)">启用</el-button> 
                       <br><el-button :disabled="scope.row.status == '1'" type="text" class="text-info" @click="statusChange(scope.row)">禁用</el-button>
                    </template>
                </el-table-column>
                <el-table-column fixed="right" align="center" label= "操作" >
                    <template slot-scope="scope">
                        <el-button type="text" class="text-info" @click="lookup(scope.row.branch_id)">查看</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>


        <add-modal ref='AddModal' v-on:refresh-data="queryList"></add-modal>
        <detail-modal ref='DetailModal' v-on:refresh-data="queryList"></detail-modal>
        
        <div class="page-pagination" align="right">
            <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page="query.pi"
                :page-size="query.ps"
                :page-sizes="pageSizes"
                layout="total, sizes, prev, pager, next, jumper"
                :total="count">
        </el-pagination>
        </div>
    </div> 
</template>

<script>
import BranchAdd from './branch.add.vue'
import BranchDetail from './branch.detail.vue'

export default {
    components: {
        'add-modal': BranchAdd,
        'detail-modal': BranchDetail,
    },
    data() {
        
        return {
            branchStatusList: this.EnumUtility.Get("BranchStatus"),// 门店状态列表
            branchTypeList: this.EnumUtility.Get("BranchType"),// 门店类型列表
            pageSizes: [10, 20, 50, 100],
            count: 0,
            dataList: [], //数据列表
            totaldata: [], //合计数据列表
            branchList: [], //门店数据列表

            query:{
                pi:1,
                ps:10,
                status: "",
                branch_id:'',
            },
            statusName: "",
        };
    },
    mounted() {
        this.refreshDataList()
        this.queryBranchList()
    },
    methods: {
        handleSizeChange(val) {
            this.query.ps = val
            this.refreshDataList()
        },
        handleCurrentChange(val) {
            this.query.pi = val
            this.queryList()
        },
        refreshDataList(){
            this.query.pi = 1
            this.queryList()
        },
        queryList: function() {
            this.$get("/branch/queryall",this.query
            ).then(res=>{
                this.dataList = res.data
                this.count = res.count
                this.queryBranchList()
                return res.data
            })
            .catch(error=>{
                if (error.response.status == 404){
                    this.$notify.error({
                        title: '操作错误',
                        message: '未找到服务'
                    });
                }else{
                    this.$notify.error({
                        title: '查询门店信息',
                        message: '获取门店信息失败'
                    });
                }
            })    
        },
        queryBranchList(){
            this.$get("/branch/getallbranchbyaccid",{}).then(res=>{
                this.branchList = res.data
            })
            .catch(error=>{
                this.$notify.error({
                    title: '操作错误',
                    message: '查找门店信息错误'
                });
            })    
        },
        statusChange(item){
            var params = { 
                branch_id: item.branch_id,
                status: item.status
            }
            if(item.status == '1'){
                this.statusName = "启用"
                params.status = '0'
            }else{
                this.statusName = "禁用"
                params.status = '1'
            }
            this.$confirm("是否"+this.statusName+"门店："+item.branch_name, '提示', {
                    cancelButtonText: '取消',
                    confirmButtonText: '确定',
                    type: 'warning',
                }).then(()=>{this.$post("/branch/editstatus",params).then(res=>{
                        this.$notify({
                            title:'编辑门店状态',
                            message: this.statusName+'成功',
                            type:'success'
                        })
                        this.queryList()//刷新页面查询的响应    
                    }).catch(error=>{
                        if (error.response.status == 403) {
                            autoLogin(this);
                        } else if (error.response.status == 401) {
                            autoLogin(this);
                        }else {
                            this.$notify.error({
                                title: this.statusName+'失败',
                            })
                        }
                    })
                }).catch(error=>{
                    this.$message({
                        type: 'info',
                        message: "已经取消"
                    });
            })
        },
        add(item){
            this.$refs.AddModal.setModal(item)
        },
        lookup(branch_id){
            console.log("查询页面branch_id:",branch_id)
            this.$refs.DetailModal.setModal(branch_id)
        },
    },
};
</script>

<style>
.el-dialog__body {
  padding: 30px 20px;
  color: #dddfe4;
  font-size: 14px;
  border-top-width: 1px;
  border-top-style: solid;
  border-bottom-style: solid;
  border-bottom-width: 1px;
}
</style>
