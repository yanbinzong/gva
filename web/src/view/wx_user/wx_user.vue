<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="openid">
          <el-input placeholder="搜索条件" v-model="searchInfo.openid"></el-input>
        </el-form-item>    
        <el-form-item label="用户昵称">
          <el-input placeholder="搜索条件" v-model="searchInfo.nickname"></el-input>
        </el-form-item>                  
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增微信用户</el-button>
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
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="openid" prop="openid" width="120"></el-table-column> 
    
    <el-table-column label="用户昵称" prop="nickname" width="120"></el-table-column> 
    
    <el-table-column label="用户头像" prop="avatarurl" width="120"></el-table-column> 
    
    <el-table-column label="性别" prop="gender" width="120"></el-table-column> 
    
    <el-table-column label="所在国家" prop="country" width="120"></el-table-column> 
    
    <el-table-column label="省份" prop="province" width="120"></el-table-column> 
    
    <el-table-column label="城市" prop="city" width="120"></el-table-column> 
    
    <el-table-column label="语言" prop="language" width="120"></el-table-column> 
    
    <el-table-column label="手机号码" prop="mobile" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateWx_user(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
         <el-form-item label="openid:">
            <el-input v-model="formData.openid" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="用户昵称:">
            <el-input v-model="formData.nickname" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="用户头像:">
            <el-input v-model="formData.avatarurl" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="性别:">
            <el-input v-model="formData.gender" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="所在国家:">
            <el-input v-model="formData.country" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="省份:">
            <el-input v-model="formData.province" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="城市:">
            <el-input v-model="formData.city" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="语言:">
            <el-input v-model="formData.language" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="手机号码:">
            <el-input v-model="formData.mobile" clearable placeholder="请输入" ></el-input>
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
    createWx_user,
    deleteWx_user,
    deleteWx_userByIds,
    updateWx_user,
    findWx_user,
    getWx_userList
} from "@/api/wx_user";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Wx_user",
  mixins: [infoList],
  data() {
    return {
      listApi: getWx_userList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            openid:"",
            nickname:"",
            avatarurl:"",
            gender:"",
            country:"",
            province:"",
            city:"",
            language:"",
            mobile:"",
            
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
           this.deleteWx_user(row);
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
        const res = await deleteWx_userByIds({ ids })
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
    async updateWx_user(row) {
      const res = await findWx_user({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rewx_user;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          openid:"",
          nickname:"",
          avatarurl:"",
          gender:"",
          country:"",
          province:"",
          city:"",
          language:"",
          mobile:"",
          
      };
    },
    async deleteWx_user(row) {
      const res = await deleteWx_user({ ID: row.ID });
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
          res = await createWx_user(this.formData);
          break;
        case "update":
          res = await updateWx_user(this.formData);
          break;
        default:
          res = await createWx_user(this.formData);
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
    }
  },
  async created() {
    await this.getTableData();
  
}
};
</script>

<style>
</style>
