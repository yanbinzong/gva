<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="类别:">
                <el-input v-model="formData.type" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="url地址:">
                <el-input v-model="formData.url" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="状态:"><el-input v-model.number="formData.status" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="文件扩展名:">
                <el-input v-model="formData.exts" clearable placeholder="请输入" ></el-input>
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
    createAttachment,
    updateAttachment,
    findAttachment
} from "@/api/attachment";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Attachment",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            type:"",
            url:"",
            status:0,
            exts:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createAttachment(this.formData);
          break;
        case "update":
          res = await updateAttachment(this.formData);
          break;
        default:
          res = await createAttachment(this.formData);
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
    const res = await findAttachment({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reattachment
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