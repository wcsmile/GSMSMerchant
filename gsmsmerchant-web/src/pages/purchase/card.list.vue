<template>
    <div>
        <bootstrap-modal ref="showModal" :need-header="true" :need-footer="true" size="medium">
            <div slot="title">
                <h3>卡号列表</h3>
            </div>
            <div slot="body">
                <div clas="table-responsive">
                    <el-scrollbar style="height:100%">
                        <el-table ref="table" :data="dataList" border height="432px" size="medium">
                            <el-table-column align="center" label="卡号" prop="recharge_account_id">
                            </el-table-column>
                        </el-table>
                    </el-scrollbar>
                    <div class="page-pagination" align="right">
                        <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="params.pi" :page-size="params.ps" :page-sizes="pageSizes" layout="total, sizes, prev, pager, next, jumper" :total="totalcount">
                        </el-pagination>
                    </div>
                </div>
            </div>

        </bootstrap-modal>

    </div>
</template>

<script>
export default {
  components: {
    "bootstrap-modal": require("vue2-bootstrap-modal")
  },
  data() {
    return {
      colorClass: ["text-success", "text-danger", "text-muted", "text-primary"],

      pageSizes: [10, 20, 50, 100],
      params: {
        pi: 1,
        ps: 10,
        order_no: ""
      },
      totalcount: 0,
      dataList: [],
    };
  },
  mounted() { },
  methods: {
    setModal(item) {
      this.QueryCardList(item.order_no);
      this.$refs.showModal.open();
    },
    closeModal() {
      this.$refs.showModal.close();
    },
    handleSizeChange(val) {
      this.params.ps = val
      this.queryDataList()
    },
    handleCurrentChange(val) {
      this.params.pi = val
      this.queryDataList()
    },
    QueryCardList(order_no) {
      this.params.order_no = order_no
      this.$get("/product/purchase/cards", this.params).then(response => {
        this.dataList = response.data
        this.totalcount = response.count
      })
    },
  }
};
</script>

<style >
.m-b-none {
  margin-bottom: 0 !important;
}

.table-responsive {
  overflow-y: hidden;
}

.table > tbody > tr > td,
.table > tfoot > tr > td {
  padding: 14px 15px;
  border-top: 1px solid #eaeff0;
  vertical-align: middle;
}

.table > thead > tr > th {
  padding: 8px 15px;
  border-bottom: 1px solid #eaeff0;
}

.table-bordered {
  border-color: #eaeff0;
}

.table-bordered > tbody > tr > td {
  border-color: #eaeff0;
}

.table-bordered > thead > tr > th {
  border-color: #eaeff0;
}

.table-striped > tbody > tr:nth-child(odd) > td,
.table-striped > tbody > tr:nth-child(odd) > th {
  background: #fafbfc;
}

.table-striped > thead > th {
  background: #fafbfc;
  border-right: 1px solid #eaeff0;
}

.table-striped > thead > th:last-child {
  border-right: none;
}
</style>