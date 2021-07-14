<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="试卷标题">
          <el-input placeholder="搜索条件" v-model="searchInfo.paper_title"></el-input>
        </el-form-item>        
        <el-form-item label="试卷状态">
          <el-input placeholder="搜索条件" v-model="searchInfo.paper_status"></el-input>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增试卷</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="40"></el-table-column>
      <el-table-column label="ID" prop="ID" width="40"></el-table-column>
      <el-table-column label="分支机构" prop="branch_office_id" width="100"></el-table-column>
      <el-table-column label="试卷标题" prop="paper_title" width="100"></el-table-column>
    <el-table-column label="试卷副标题" prop="paperfu_title" width="120"></el-table-column> 
    
    <el-table-column label="试卷说明" prop="paper_intro" width="120"></el-table-column>

    <el-table-column label="开始时间" width="180">
      <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>

    <el-table-column label="结束时间" width="180">
        <template slot-scope="scope">{{scope.row.end_at|formatDate1}}</template>
    </el-table-column>

      <el-table-column label="试卷状态" prop="paper_status" width="120">
        <template slot-scope="scope">
          <el-tag type="success" v-if="scope.row.paper_status">已发布</el-tag>
          <el-tag type="danger" v-if="!scope.row.paper_status">已截止</el-tag>
        </template>
      </el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="goQuestionUrl(scope.row)">题库</el-button>
          <el-button class="table-button" @click="updateExam_paper(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-button  icon="el-icon-search"  size="mini" type="primary"   @click="goPreviewQuestion(scope.row)">预览</el-button>
          <el-button  icon="el-icon-search"  size="mini" type="primary"   @click="goPreview(scope.row)">预览api</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>

        </template>
      </el-table-column>

    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="试卷标题:">
            <el-input v-model="formData.paper_title" clearable placeholder="请输入" ></el-input>
      </el-form-item>

         <el-form-item label="试卷副标题:">
            <el-input v-model="formData.paperfu_title" clearable placeholder="请输入" ></el-input>
      </el-form-item>

         <el-form-item label="试卷说明:">
            <el-input type="textarea" v-model="formData.paper_intro" clearable placeholder="请输入" ></el-input>
      </el-form-item>

        <el-form-item label="开始时间:">
          <el-date-picker
              v-model="formData.CreatedAt"
              type="datetime"
              placeholder="选择开始日期时间">
          </el-date-picker>
        </el-form-item>

        <el-form-item label="结束时间:">
          <el-date-picker
              v-model="formData.end_at"
              type="datetime"
              placeholder="选择结束日期时间">
          </el-date-picker>
        </el-form-item>

        <el-form-item label="试卷状态:">
         <el-switch v-model="formData.paper_status"
                     :active-value="1"
                     :inactive-value="0"
          >
          </el-switch>
        </el-form-item>

       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createExam_paper,
    deleteExam_paper,
    deleteExam_paperByIds,
    updateExam_paper,
    findExam_paper,
    getExam_paperList
} from "@/api/exam_paper";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Exam_paper",
  mixins: [infoList],
  data() {
    return {
      listApi: getExam_paperList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      intOptions:[],
          formData: {
            paper_title:"",
            paperfu_title:"",
            paper_intro:"",
            paper_status:0,
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatDate1: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10        
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      deleteRow(row){
        this.$confirm('确定要删除吗?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
           this.deleteExam_paper(row);
        });
      },
    goQuestionUrl(row){
      this.$confirm('你确定要查看题库吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$router.push({path: '/layout/questionP/questionTable', query: {exam_paper_id: row.ID}});
      });
    },
    goPreview(row){
        var paper_url ="http://localhost:8888/paper?paperid="+row.ID
      window.open(paper_url,'_blank') // 在新窗口打开外链接
     // this.$router.push({path: 'http://localhost:8888/app/paper', query: {paperid: row.ID}});
      this.$router.push({path:'/layout/exam_paperP/exam_paperPreview',query:{exam_paper_id:row.ID}})
    },
    goPreviewQuestion(row){
        this.$router.push({path:'/layout/exam_paperP/exam_paperPreview',query:{exam_paper_id:row.ID}})
    },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteExam_paperByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          if (this.tableData.length == ids.length && this.page > 1) {
              this.page--;
          }
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateExam_paper(row) {
      const res = await findExam_paper({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reexam_paper;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          paper_title:"",
          paperfu_title:"",
          paper_intro:"",
          paper_status:0,
          
      };
    },
    async deleteExam_paper(row) {
      const res = await deleteExam_paper({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1 && this.page > 1 ) {
            this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createExam_paper(this.formData);
          break;
        case "update":
          res = await updateExam_paper(this.formData);
          break;
        default:
          res = await createExam_paper(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    },

  },
  async created() {
    await this.getTableData();
  
    await this.getDict("int");
    
}
};
</script>

<style>
</style>
