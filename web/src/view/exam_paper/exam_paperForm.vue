<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="试卷标题:">
                <el-input v-model="formData.paper_title" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="试卷副标题:">
                <el-input v-model="formData.paperfu_title" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="试卷说明:">
                <el-input v-model="formData.paper_intro" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="试卷状态:">
                 <el-select v-model="formData.paper_status" placeholder="请选择" clearable>
                     <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                 </el-select>
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
    createExam_paper,
    updateExam_paper,
    findExam_paper
} from "@/api/exam_paper";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Exam_paper",
  mixins: [infoList],
  data() {
    return {
      type: "",
      intOptions:[],
          formData: {
            paper_title:"",
            paperfu_title:"",
            paper_intro:"",
            paper_status:0,
            
      }
    };
  },
  methods: {
    async save() {
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
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findExam_paper({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reexam_paper
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
    await this.getDict("int");
    
}
};
</script>

<style>
</style>