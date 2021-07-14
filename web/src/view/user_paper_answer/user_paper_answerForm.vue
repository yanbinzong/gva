<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="用户ID:"><el-input v-model.number="formData.user_id" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="试卷ID:"><el-input v-model.number="formData.paper_id" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="问题ID:"><el-input v-model.number="formData.question_id" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="回答内容:">
                <el-input v-model="formData.answer_content" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="分数:">
                  <el-input-number v-model="formData.score" :precision="2" clearable></el-input-number>
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
    createuser_paper_answer,
    updateuser_paper_answer,
    finduser_paper_answer
} from "@/api/user_paper_answer";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "user_paper_answer",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            user_id:0,
            paper_id:0,
            question_id:0,
            answer_content:"",
            score:0,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createuser_paper_answer(this.formData);
          break;
        case "update":
          res = await updateuser_paper_answer(this.formData);
          break;
        default:
          res = await createuser_paper_answer(this.formData);
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
    const res = await finduser_paper_answer({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reuser_paper_answer
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