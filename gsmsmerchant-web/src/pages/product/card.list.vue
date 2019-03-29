<template>
    <div class="panel panel-default">
        <div class="panel-body">
            <el-form ref="form" :model="params" :inline="true" class="form-inline pull-left add-qx-bottom">
                <div class="form-group">
                    <el-form-item>
                        <span class="demonstration">下发时间</span>
                        <el-date-picker size="medium" :clearable="false" :unlink-panels="true" :picker-options="pickerOptions" value-format="yyyy-MM-dd HH:mm:ss" :default-time="['00:00:00','23:59:59']" v-model="time" type="daterange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期">
                        </el-date-picker>
                    </el-form-item>
                    <!-- <el-form-item>
                        <el-input v-model="params.branch" size="medium" placeholder="请输入门店名称"></el-input>
                    </el-form-item> -->

                    <el-form-item>
                        <el-select filterable size="medium" v-model="params.order_source" class="input-cos" filterable placeholder="请选择门店">
                            <el-option value="">所有</el-option>
                            <el-option v-for="(item, index) in branchList" :key="index" :value="item.branch_id" :label="item.branch_name"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item>
                        <el-input v-model="params.card_no" size="medium" placeholder="请输入卡号"></el-input>
                    </el-form-item>
                    <button class="btn btn-success" type="button" @click="search">查询</button>
                </div>
            </el-form>
        </div>
        <el-scrollbar style="height:100%">
            <el-table :data="dataList" border height="532px">
                <el-table-column prop="recharge_account_id" align="center" label="卡号">
                </el-table-column>
                <el-table-column align="center" prop="create_time" label="下发时间">
                </el-table-column>
                <el-table-column align="center" label="卡片类型">
                    <template slot-scope="scope">
                        <span>{{ scope.row.business_type | BusinesstTypeFilter(businessType)}}</span>
                    </template>
                </el-table-column>
                <el-table-column align="center" prop="channel_name" label="初始余额（元）">
                    <template slot-scope="scope">
                        <span v-if="scope.row.business_type=='6801'">0</span>
                        <span v-if="scope.row.business_type!='6801'">{{ scope.row.face | AmountFilter)}}</span>
                    </template>
                </el-table-column>
                <el-table-column align="center" label="所属门店">
                    <template slot-scope="scope">
                        <span>{{ scope.row.order_source | BranchFilter(branchList)}}</span>
                    </template>
                </el-table-column>
                <el-table-column align="center" label="复充值">
                    <template slot-scope="scope">
                        <span v-if="scope.row.has_first_recharge=='1'">否</span>
                        <span v-if="scope.row.has_first_recharge=='0'">是</span>
                    </template>
                </el-table-column>
                <el-table-column align="center" label="激活状态">
                    <template slot-scope="scope">
                        <div :class="orderStatusClassFilter(scope.row.status)">{{scope.row.status | EnumFilter('OilCardStatus')}}</div>
                    </template>
                </el-table-column>
                <el-table-column align="center" label="卡片状态">
                    <template slot-scope="scope">
                        <span v-if="scope.row.has_first_recharge=='1'">已下发</span>
                        <span v-if="scope.row.has_first_recharge=='0'">使用中</span>
                    </template>
                </el-table-column>
            </el-table>
        </el-scrollbar>

        <div class="page-pagination" align="right">
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="params.pi" :page-size="params.ps" :page-sizes="pageSizes" layout="total, sizes, prev, pager, next, jumper" :total="totalcount">
            </el-pagination>
        </div>
    </div>
</template>

<script>
import Notification from 'element-ui'

export default {
  components: {
  },
  data() {
    return {
      heightMax: '100%',
      colorClass: [
        'text-success',
        'text-danger',
        'text-muted',
        'text-primary',
      ],
      tableheader: {
        'text-align:center': true
      },
      pageSizes: [10, 20, 50, 100],
      params: {
        pi: 1,
        ps: 10,
        start_time: "",
        end_time: "",
        branch: "",
        card_no: "",// 发货类型
      },
      businessType: [],
      branchList: [],
      totalcount: 0,
      dataList: [],
      needLogisticsList: this.EnumUtility.Get("NeedLogistics"),
      time: [
        this.DateConvert("yyyy-MM-dd 00:00:00", new Date(new Date().getTime() - 3600 * 1000 * 24 * 90)),
        this.DateConvert("yyyy-MM-dd 23:59:59", new Date())
      ],
      pickerOptions: {
        shortcuts: [{
          text: '最近一周',
          onClick(picker) {
            const end = new Date();
            const start = new Date();
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
            picker.$emit('pick', [start, end]);
          }
        }, {
          text: '最近一个月',
          onClick(picker) {
            const end = new Date();
            const start = new Date();
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
            picker.$emit('pick', [start, end]);
          }
        }, {
          text: '最近三个月',
          onClick(picker) {
            const end = new Date();
            const start = new Date();
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
            picker.$emit('pick', [start, end]);
          }
        }]
      },
    }
  },
  mounted() {
    this.init()
  },
  methods: {
    init() {
      this.getBusinessType()
      this.getAllBranch()
      this.params.pi = 1
      this.params.ps = 10
      this.queryDataList()
    },
    getBusinessType() {
      this.$get("/sys/businesstype/query", {}).then(response => {
        this.businessType = response.data
      })
    },
    getAllBranch() {
      this.$get("/branch/getallbranch", {}).then(response => {
        this.branchList = response.data
      })
    },
    queryDataList() {
      if (this.time) {
        this.params.start_time = this.time[0]
        this.params.end_time = this.time[1]
      }

      this.$get("/product/cardquery", this.params).then(response => {
        if (!response.data || response.count === 0) {
          this.dataList = []
          this.totalcount = 0
          return
        }
        this.dataList = response.data
        this.totalcount = parseInt(response.count)
      })
    },
    search() {
      this.params.pi = 1
      this.queryDataList()
    },
    handleSizeChange(val) {
      this.params.ps = val
      this.queryDataList()
    },
    handleCurrentChange(val) {
      this.params.pi = val
      this.queryDataList()
    },
    orderStatusClassFilter(item) {
      switch (item) {
        case '0':
          return this.colorClass[0]
          break;
        case '90':
          return this.colorClass[1]
          break;
        case '99':
          return this.colorClass[2]
          break;
        default:
          return this.colorClass[3]
          break;
      }
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