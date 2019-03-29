
<template>
    <div class="panel panel-default">
        <div class="panel panel-default">
            <div class="panel-body">
                <el-form ref="form" :model="params" :inline="true" class="form-inline pull-left add-qx-bottom">
                    <el-form-item>
                    <el-form-item>
                       <el-select placeholder="请选择订单来源" v-model="query.channel_no">
                            <el-option value=''>全部</el-option>
                            <el-option v-for="(item, index) in channelTypeList" :key="index" :value="item.channel_no" :label="item.channel_name" ></el-option>
                        </el-select>
                    </el-form-item>
                        <el-select  placeholder="付款状态" v-model="query.payment_status" style="width: 160px">
                            <el-option value=''>全部</el-option>
                            <el-option v-for="(item, index) in payStatusList" :key="index" :value="item.value" :label="item.name" ></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item>
                        <el-input v-model="query.order_no"   placeholder="请输入订单编号"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <a class="btn btn-success " @click="refreshDataList" style="margin-bottom:3px;">搜索</a> 
                    </el-form-item >
                </el-form>
            </div>
        </div>

        <div class="table-responsive">
            <el-table :data="dataList"  style="width:100%"  border height="532px">
                <el-table-column prop="order_no" align="center" label="订单编号">
                </el-table-column>
                <el-table-column prop="branch_type" align="center" label="订单归属" width="180px">
                  <!-- <template slot-scope="scope">
                    <span style="text-align:center">{{scope.row.branch_type | EnumFilter("BranchType") }}</span>
                  </template> -->
                </el-table-column>
                <el-table-column  prop="channel_name" align="center" label="订单来源">
                </el-table-column>
                <el-table-column prop="recharge_account_id" align="center" label="交易卡号" >
                </el-table-column>
                <el-table-column prop="business_type" align="center" label="交易类型" >
                    <template slot-scope="scope">
                        {{scope.row.business_type | BusinesstTypeFilter(businessType)}}
                    </template>
                </el-table-column>
                <el-table-column prop="total_face" align="center" label="交易金额（元）" >
                </el-table-column>
                <el-table-column prop="product_face" align="center" label="面值（元）" >
                </el-table-column>
                <el-table-column prop="payment_status" align="center" label="付款状态" >
                    <template slot-scope="scope">
                        <span :class="scope.row.payment_status == '0'?'text-success':'text-danger'" style="text-align:center">{{scope.row.payment_status | EnumFilter("PayStatus") }}</span>
                    </template>
                </el-table-column>
                <el-table-column prop="recharge_status" align="center" label="充值状态" >
                    <template slot-scope="scope">
                        <span :class="scope.row.recharge_status == '0'?'text-success':'text-danger'" style="text-align:center">{{scope.row.recharge_status | EnumFilter("RechargeStatus") }}</span>
                    </template>
                </el-table-column>
            </el-table>
        </div>
        
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
export default {
  data() {
    return {
      pageSizes: [10, 20, 50, 100],
      count: 0,
      dataList: [], //数据列表
      businessType: [],
      channelTypeList:[], //订单来源
      payStatusList:this.EnumUtility.Get("PayStatus"),// 支付状态列表
      query:{
        pi:1,
        ps:10,
        order_no:'',
        payment_status: "",
        channel_no:'',
      },
    }
  },
  mounted() {
    this.queryChannelList()
    this.queryList()
    this.getBusinessType()
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
    queryList() {
        this.$get("/order/query",this.query
        ).then(res=>{
            this.dataList = res.data
            this.count = res.count
            return res.data
        })
        .catch(error=>{
            console.log(error)
        })    
    },
    queryChannelList() {
        this.$get("/order/querychannel",{}).then(res=>{
            this.channelTypeList = res.data
            return res.data
        })
        .catch(error=>{
            console.log(error)
        })    
    },
    getBusinessType() {
        this.$get("/sys/businesstype/query", {}).then(response => {
          this.businessType = response.data
        })
      },
  },
}
</script>

<style scoped>
.input-cos {
  width: 150px;
}

.el-scrollbar__wrap {
  overflow-y: hidden;
}

.el-button--medium {
  padding: 0;
}
</style>
<style >
.datetime-to-date .el-date-picker__time-header {
  display: none;
}
</style>