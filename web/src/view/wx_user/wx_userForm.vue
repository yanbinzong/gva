<template>
<div>
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
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createWx_user,
    updateWx_user,
    findWx_user
} from "@/api/wx_user";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Wx_user",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
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
  methods: {
    async save() {
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
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findWx_user({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.rewx_user
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
}
};
</script>

<style>
</style>