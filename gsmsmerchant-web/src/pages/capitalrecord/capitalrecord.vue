<template>
  <div class="panel panel-default">
    <div class="panel-body">
      <el-form ref="form" :model="params" :inline="true" class="form-inline pull-left add-qx-bottom">
        <el-form-item>
          <span class="demonstration">交易时间</span>
          <el-date-picker size="medium" :clearable="false" :unlink-panels="true" :picker-options="pickerOptions"
            value-format="yyyy-MM-dd HH:mm:ss" :default-time="['00:00:00','23:59:59']" v-model="time" type="daterange"
            range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期">
          </el-date-picker>
        </el-form-item>
        <el-form-item v-if="user_type != 'channel'">
          <el-select filterable size="medium" v-model="params.record_type" class="input-cos" placeholder="请选择交易类型">
            <el-option value="">所有</el-option>
            <el-option v-for="(item, index) in businessType" :key="index" :value="item.business_type"
              :label="item.business_name"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="user_type == 'channel'">
          <el-select filterable size="medium" v-model="params.record_type" class="input-cos" placeholder="请选择交易类型">
            <el-option value="">所有</el-option>
            <el-option v-for="(item, index) in channelType" :key="index" :value="item.value" :label="item.name">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-input placeholder="请输入交易流水号" v-model="params.record_no" size="medium" style="width: 160px"></el-input>
        </el-form-item>
        <el-form-item>
          <button class="btn btn-success" type="button" @click="search" style="margin-bottom:3px;">查询</button>
        </el-form-item>
      </el-form>
    </div>
    <el-scrollbar style="height:100%">
      <el-table :data="dataList" border height="532px">
        <el-table-column prop="id" align="center" label="交易流水号">
        </el-table-column>
        <el-table-column prop="create_time" align="center" label="交易时间">
        </el-table-column>
        <el-table-column align="center" label="交易金额（元）">
          <template slot-scope="scope">
            <span>{{ scope.row.amount | AmountFilter)}}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" prop="channel_name" label="交易后余额（元）">
          <template slot-scope="scope">
            <span>{{ scope.row.balance | AmountFilter)}}</span>
          </template>
        </el-table-column>

        <el-table-column align="center" label="交易类型" v-if="user_type == 'channel'">
          <template slot-scope="scope">
            {{scope.row.channel_type | ChannelTypeFilter(channelType)}}
          </template>
        </el-table-column>

        <el-table-column align="center" label="交易类型" v-if="user_type != 'channel'">
          <template slot-scope="scope">
            {{scope.row.business_type | BusinesstTypeFilter(businessType)}}
          </template>
        </el-table-column>

        <el-table-column align="center" label="备注" v-if="user_type == 'channel'">
          <template slot-scope="scope">
            {{scope.row.memo}}
          </template>
        </el-table-column>

        <el-table-column prop="change_type" align="center" label="备注" v-if="user_type != 'channel'">
          <template slot-scope="scope">
            {{scope.row.change_type | EnumFilter("ChangeType")}}
          </template>
        </el-table-column>
      </el-table>
    </el-scrollbar>

    <div class="page-pagination" align="right">
      <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="params.pi"
        :page-size="params.ps" :page-sizes="pageSizes" layout="total, sizes, prev, pager, next, jumper"
        :total="totalcount">
      </el-pagination>
    </div>
  </div>
</template>

<script>
  import VueCookies from 'vue-cookies'
  export default {
    components: {

    },
    data() {
      return {
        pageSizes: [10, 20, 50, 100],
        params: {
          pi: 1,
          ps: 10,
          start_time: (new Date()).getFullYear() + "-" + ((new Date()).getMonth() + 1) + "-" + ((new Date()).getDate()),
          end_time: "",
          record_type: "", // 交易类型
          record_no: ""
        },
        totalcount: 0,
        dataList: [],
        businessType: [],
        channelType: this.EnumUtility.Get("ChannelChangeType"),
        user_type: "",

        time: [
          this.DateConvert("yyyy-MM-dd 00:00:00", new Date()),
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
      var userInfo = VueCookies.get("userinfo");
      this.user_type = userInfo.acc_type

      if (this.user_type != "channel") {
        this.getBusinessType()
      }
      this.init()
    },
    methods: {
      init() {
        this.params.pi = 1
        this.params.ps = 10
        this.queryDataList()
      },
      getBusinessType() {
        this.$post("/sys/businesstype/query", {}).then(response => {
          this.businessType = response.data
        })
      },
      queryDataList() {
        if (this.time) {
          this.params.start_time = this.time[0]
          this.params.end_time = this.time[1]
        }

        this.$get("/capitalrecord/list/query", this.params).then(response => {
          if (!response.list || response.count === 0) {
            this.dataList = []
            this.totalcount = 0
            return
          }
          this.dataList = response.list
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
    }
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
<style>
  .datetime-to-date .el-date-picker__time-header {
    display: none;
  }

</style>
