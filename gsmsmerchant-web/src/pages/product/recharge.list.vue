<template>
    <div class="panel panel-default">
        <div class="panel-body">
        </div>
        <el-scrollbar style="height:100%">
            <el-table :data="dataList" border height="532px">
                <el-table-column align="center" label="业务类型">
                    <template slot-scope="scope">
                        <span>{{ scope.row.business_type | BusinesstTypeFilter(businessType)}}</span>
                    </template>
                </el-table-column>
                <el-table-column align="center" prop="province_name" label="省份">
                </el-table-column>
                <el-table-column align="center" prop="city_name" label="城市">
                </el-table-column>>
                <el-table-column align="center"  prop="face" label="面值（元）">
                     <!-- <template slot-scope="scope">
                        <span>{{ scope.row.face | AmountFilter}}</span>
                    </template> -->
                </el-table-column>
                <el-table-column align="center" prop="deduct_discount" label="折扣扣款">
                </el-table-column>
                <el-table-column align="center" label="状态">
                    <template slot-scope="scope">
                        <div :class="orderStatusClassFilter(scope.row.status)">{{scope.row.status | EnumFilter('DownProductStatus')}}</div>
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
      },
      businessType: [],
      totalcount: 0,
      dataList: [],
    }
  },
  mounted() {
    this.init()
  },
  methods: {
    init() {
      this.getBusinessType()
      this.params.pi = 1
      this.params.ps = 10
      this.queryDataList()
    },
    getBusinessType() {
      this.$get("/sys/businesstype/query", {}).then(response => {
        this.businessType = response.data
      })
    },
    queryDataList() {
      this.$get("/product/rechargequery", this.params).then(response => {
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