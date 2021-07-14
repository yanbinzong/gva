<template>
  <div>
    <el-row>
      <el-col :span="2">
        <el-upload
            :action="`${path}/excel/importExcel`"
            :headers="{'x-token':token}"
            :on-success="loadExcel"
            :show-file-list="false"
        >
          <el-button size="small" type="primary" icon="el-icon-upload2">导入</el-button>
        </el-upload>
      </el-col>
      <el-col :span="2">
        <el-button size="small" type="primary" icon="el-icon-download" @click="handleExcelExport('ExcelExport.xlsx')">导出</el-button>
      </el-col>
      <el-col :span="2">
        <el-button size="small" type="success" icon="el-icon-download" @click="downloadExcelTemplate()">下载模板</el-button>
      </el-col>
    </el-row>


    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="问题编号">
          <el-input placeholder="搜索条件" v-model="searchInfo.question_id"></el-input>
        </el-form-item>    
        <el-form-item label="问题标题">
          <el-input placeholder="搜索条件" v-model="searchInfo.question_name"></el-input>
        </el-form-item>      
<!--        <el-form-item label="是否必答">-->
<!--          <el-input placeholder="搜索条件" v-model="searchInfo.answer"></el-input>-->
<!--        </el-form-item>    -->
        <el-form-item label="题型编号">
          <el-input placeholder="搜索条件" v-model="searchInfo.topic_id"></el-input>
        </el-form-item>
        <el-form-item label="发题人">
          <el-input placeholder="搜索条件" v-model="searchInfo.author"></el-input>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增问题</el-button>
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
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="日期" width="155">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
<!--    -->
<!--    <el-table-column label="问题编号" prop="question_id" width="120"></el-table-column>-->
      <el-table-column label="分支机构" prop="branch_office_id" width="88"></el-table-column>
      <el-table-column label="问题标题" prop="question_name" width="120"></el-table-column>
    
    <el-table-column label="说明" prop="direction" width="120"></el-table-column> 

    <el-table-column label="必答" prop="answer" width="88"></el-table-column>

    <el-table-column label="题型编号" width="100">
      <template slot-scope="scope">
        <li v-for="(typelist,index) in question_typeList" :key="index">
          <em v-if="typelist.ID==scope.row.topic_id" >
          {{ typelist.qtype_name }}
          </em>
        </li>
      </template>
    </el-table-column>
    <el-table-column label="排序 " prop="sort" width="55"></el-table-column>
    <el-table-column label="发题人" prop="author" width="120"></el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateQuestion(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
      @current-change="handleCurrentChangeQestion"
      @size-change="handleSizeChangeQuestion"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
<!--         <el-form-item label="问题编号:"><el-input v-model.number="formData.question_id" clearable placeholder="请输入"></el-input>-->
<!--      </el-form-item>-->

        <el-form-item label="分支机构:">
          <el-input v-model.number="formData.branch_office_id" clearable placeholder="请输入" type="number" ></el-input>
        </el-form-item>

        <el-form-item label="试卷:">
          <el-input v-model.number="formData.exam_paper_id" clearable placeholder="请输入" type="number" ></el-input>
        </el-form-item>

      <el-form-item label="问题标题:">
            <el-input v-model="formData.question_name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
      <el-form-item label="说明:">
            <el-input v-model="formData.direction" clearable placeholder="请输入" ></el-input>
      </el-form-item>
        <el-form-item label="排序:">
          <el-input v-model.number="formData.sort" clearable placeholder="请输入" type="number" ></el-input>
        </el-form-item>
         <el-form-item label="是否必答:">
<!--             <el-switch v-model="formData.answer" active-value=1 inactive-value=0 ></el-switch>-->
           <el-switch v-model="formData.answer"
                      :active-value="1"
                      :inactive-value="0"
                     >
           </el-switch>
          </el-form-item>
       
         <el-form-item label="题型编号:">
           <el-select v-model="formData.topic_id" placeholder="请选择">
             <el-option
                 v-for="item in question_typeList"
                 :key="item.ID"
                 :label="item.qtype_name"
                 :value="item.ID.toString()">
             </el-option>
           </el-select>
<!--            <el-input v-model="formData.topic_id" clearable placeholder="请输入" ></el-input>-->
         </el-form-item>

         <el-form-item label="发题人:">
            <el-input v-model="formData.author" clearable placeholder="请输入" :disabled="true" ></el-input>
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
const path = process.env.VUE_APP_BASE_API;
import { mapGetters } from 'vuex';

import { wenTiExportExcel, loadExcelData, downloadTemplate } from "@/api/excel";

import {
    createQuestion,
    deleteQuestion,
    deleteQuestionByIds,
    updateQuestion,
    findQuestion,
    getQuestionList
} from "@/api/question";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";


import {
  getQuestion_typeList
} from "@/api/questiontype";  //  此处请自行替换地址
export default {
  name: "Question",
  mixins: [infoList],
  data() {
    return {
      listApi: getQuestionList,
      // listApiQuestionType:getQuestion_typeList,
      path: path,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            question_id:0,
            question_name:"",
            direction:"",
            answer:0,
            topic_id:"",
            author:"",
            branch_office_id:0,
            exam_paper_id:0
      },
      question_typeList:"",
    };
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'token'])
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
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
    handleExcelExport(fileName) {
      if (!fileName || typeof fileName !== "string") {
        fileName = "ExcelExport.xlsx";
      }
      console.log(this.tableData)
      wenTiExportExcel(this.tableData, fileName);
    },
    loadExcel() {
      this.listApi = loadExcelData;
      this.getTableDataSelect();
    },
    downloadExcelTemplate() {
      downloadTemplate('ExcelTemplate.xlsx')
    },
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10
        this.getTableDataSelect()
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
           this.deleteQuestion(row);
        });
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
        const res = await deleteQuestionByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          if (this.tableData.length == ids.length && this.page > 1) {
              this.page--;
          }
          this.deleteVisible = false
          this.getTableDataSelect()
        }
      },
    async updateQuestion(row) {
      const res = await findQuestion({ ID: row.ID });

      const questionTypeList =  await getQuestion_typeList();
      ///console.log(questionTypeList)
      if (questionTypeList.code==0){
        this.question_typeList = questionTypeList.data.list;
        console.log( this.question_typeList)
      }

      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.requestion;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          question_id:0,
          question_name:"",
          direction:"",
          answer:0,
          topic_id:"",
          author:"",
         exam_paper_id:0
          
      };
    },
    async deleteQuestion(row) {
      const res = await deleteQuestion({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1 && this.page > 1 ) {
            this.page--;
        }
        this.getTableDataSelect();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          var  exam_paper_id = this.$route.query.exam_paper_id
          this.formData.exam_paper_id = exam_paper_id

          res = await createQuestion(this.formData);
          break;
        case "update":
          res = await updateQuestion(this.formData);
          break;
        default:
          res = await createQuestion(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        var  exam_paper_id = this.$route.query.exam_paper_id
        this.formData.exam_paper_id = exam_paper_id
        this.getTableDataSelect();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;

      // 获取用户信息
      var getLocalData =localStorage.getItem('vuex')
      console.log(typeof(getLocalData)); // string
      var jsonObj = JSON.parse(getLocalData);
      console.log(jsonObj.user.userInfo.uuid)
      this.formData.author= jsonObj.user.userInfo.nickName;
    }
  },
  async created() {
    var  exam_paper_id = this.$route.query.exam_paper_id
    this.formData.exam_paper_id = exam_papeformCreater_id
    console.log(exam_paper_id)
    await this.getTableDataSelect();

    // 获取问题类型
    const questionTypeList =  await getQuestion_typeList();
    ///console.log(questionTypeList)
    if (questionTypeList.code==0){
      this.question_typeList = questionTypeList.data.list;
      console.log( this.question_typeList)
    }

  }
};
</script>
<style>
</style>
